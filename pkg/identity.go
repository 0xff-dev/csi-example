package pkg

import (
	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
	"k8s.io/klog"
)

const (
	driverName    = "csi.example.golang"
	driverVersion = "beta"
)

// Identidy Impl
type Identity struct {
}

func NewIdentity() *Identity {
	return &Identity{}
}

// GetPluginInfo return Plugin detail
func (i *Identity) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
	klog.V(4).Infof("Identity.GetPluginInfo: args: %+v", *req)

	return &csi.GetPluginInfoResponse{
		Name:          driverName,
		VendorVersion: driverVersion,
	}, nil
}

// Probe health check.
func (i *Identity) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
	klog.V(4).Infof("Identidy.Probe: args: %+v", *req)

	return &csi.ProbeResponse{}, nil
}

// GetPluginCapabilities Returns the functions supported by the plugin
func (i *Identity) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
	klog.V(4).Infof("Identity.GetPluginCapabilities: args: %+v", *req)

	return &csi.GetPluginCapabilitiesResponse{
		Capabilities: []*csi.PluginCapability{
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
					},
				},
			},
			{
				Type: &csi.PluginCapability_Service_{
					Service: &csi.PluginCapability_Service{
						Type: csi.PluginCapability_Service_VOLUME_ACCESSIBILITY_CONSTRAINTS,
					},
				},
			},
		},
	}, nil
}
