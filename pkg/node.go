package pkg

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/klog"
	"k8s.io/utils/exec"
	"k8s.io/utils/mount"
)

type Node struct {
	nodeID  string
	mounter mount.SafeFormatAndMount
}

func NewNode(nodeID string) *Node {
	return &Node{
		nodeID: nodeID,
		mounter: mount.SafeFormatAndMount{
			Interface: mount.New(""),
			Exec:      exec.New(),
		},
	}
}

func (n *Node) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	klog.V(4).Infof("Node.NodeStageVolue: args: %+v", *req)

	return &csi.NodeStageVolumeResponse{}, nil
}

func (n *Node) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	klog.V(4).Infof("Node.NodeUnstageVolume: args: %+v", *req)

	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (n *Node) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	klog.V(4).Infof("Node.NodePublisVolime: args: %+v", *req)

	return &csi.NodePublishVolumeResponse{}, nil
}

func (n *Node) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	klog.V(4).Infof("Node.NodeUnpublishVolume: args: %+v", *req)

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (n *Node) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	klog.V(4).Infof("Node.NodeGetVolumeStats: args: %+v", *req)

	return &csi.NodeGetVolumeStatsResponse{}, nil
}

func (n *Node) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	klog.V(4).Infof("Node.NodeExpandVolume: args: %+v", *req)

	return nil, status.Error(codes.Unimplemented, "")
}

func (n *Node) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	klog.V(4).Infof("Node.NodeGetCapabilities: args: %+v", *req)

	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
					},
				},
			},
		},
	}, nil
}

func (n *Node) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	klog.V(4).Infof("Node.NodeGetInfo: args: %+v", *req)

	return &csi.NodeGetInfoResponse{
		NodeId: n.nodeID,
	}, nil
}
