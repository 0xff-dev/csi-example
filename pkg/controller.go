package pkg

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
)

const (
	devicePath = "devicePath"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

// Stage create volume
func (c *Controller) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	klog.V(4).Infof("Controller.CreateVolume: args: %+vr", *req)

	// success create a volume
	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      "volume-1234",
			CapacityBytes: 20 * (1 << 30),
			VolumeContext: req.GetParameters(),
		},
	}, nil
}

func (c *Controller) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	klog.V(4).Infof("Controller.DeleteVolume: args: %+v", *req)

	return &csi.DeleteVolumeResponse{}, nil
}

// Stage attach
func (c *Controller) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	klog.V(4).Infof("Controller.ControllerPublishVolume: args: %+v", *req)

	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			devicePath: "/dev/sde",
		},
	}, nil
}

func (c *Controller) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	klog.V(4).Infof("Controller.ControllerUnpublishVolume: args: %+v", *req)

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (c *Controller) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: []*csi.ControllerServiceCapability{
			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
					},
				},
			},
			{
				Type: &csi.ControllerServiceCapability_Rpc{
					Rpc: &csi.ControllerServiceCapability_RPC{
						Type: csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
					},
				},
			},
		},
	}, nil
}

func (c *Controller) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (c *Controller) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
