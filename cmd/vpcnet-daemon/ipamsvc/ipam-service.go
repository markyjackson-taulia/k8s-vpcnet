package ipamsvc

import (
	"context"

	"github.com/golang/glog"
	"github.com/lstoll/k8s-vpcnet/pkg/allocator"
	"github.com/lstoll/k8s-vpcnet/pkg/config"
	"github.com/lstoll/k8s-vpcnet/pkg/vpcnetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"k8s.io/apimachinery/pkg/util/runtime"
)

var _ vpcnetpb.IPAMServer = &Service{}

// Evictor is called when the pool is empty to remove pods
type Evictor interface {
	// EvictPod removes the given pod
	EvictPod(namespace, name string) error
}

// Allocator defines out contract with the IP allocator
type Allocator interface {
	Allocate(containerID, podName, podNamspace string) (*allocator.Allocation, error)
	ReleaseByContainer(containerID string) error
}

// Service implements the IPAM gRPC service by interacting with an underying
// Allocator
type Service struct {
	Allocator Allocator
	Config    *config.Config
	Evictor   Evictor
}

// Add is called when a container is added
func (i *Service) Add(ctx context.Context, req *vpcnetpb.AddRequest) (*vpcnetpb.AddResponse, error) {
	a, err := i.Allocator.Allocate(req.ContainerID, req.PodName, req.PodNamespace)
	if err != nil {
		if err == allocator.ErrNoFreeIPs && i.Config.DeletePodWhenNoIPs {
			glog.Warningf("No free IPs while allocating Container %q Pod %s/%s, deleting pod", req.ContainerID, req.PodNamespace, req.PodName)
			if err := i.Evictor.EvictPod(req.PodNamespace, req.PodName); err != nil {
				runtime.HandleError(err)
				glog.Errorf("Error evicting pod, ignoring: [%+v]", err)
			}
		} else {
			glog.Errorf("Error calling allocator Allocate for Container %q Pod %s/%s: [%+v}", req.ContainerID, req.PodNamespace, req.PodName, err)
		}
		return nil, grpc.Errorf(codes.Internal, "Error allocating address: %q", err.Error())
	}

	return &vpcnetpb.AddResponse{
		AllocatedIP: a.IP.String(),
		ENIIP:       a.ENIIP.String(),
		SubnetCIDR:  a.ENISubnet.String(),
	}, nil
}

// Del is called when a container is removed
func (i *Service) Del(ctx context.Context, req *vpcnetpb.DelRequest) (*vpcnetpb.DelResponse, error) {
	err := i.Allocator.ReleaseByContainer(req.ContainerID)
	if err != nil {
		glog.Errorf("Error calling allocator Release for Container %q: [%+v}", req.ContainerID, err)
		return nil, grpc.Errorf(codes.Internal, "Error releasing address: %q", err.Error())
	}

	return &vpcnetpb.DelResponse{}, nil
}
