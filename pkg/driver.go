package pkg

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/csi-lib-utils/protosanitizer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"k8s.io/klog"
)

type Driver struct {
	nodeID   string
	endpoint string
}

func NewDriver(nodeID, endPoint string) *Driver {
	klog.V(4).Infof("Driver: %v version: %v", driverName, driverVersion)

	return &Driver{
		nodeID:   nodeID,
		endpoint: endPoint,
	}
}

func (d *Driver) Run() {
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(logGRPC),
	}
	srv := grpc.NewServer(opts...)
	csi.RegisterControllerServer(srv, NewController())
	csi.RegisterIdentityServer(srv, NewIdentity())
	csi.RegisterNodeServer(srv, NewNode(d.nodeID))

	proto, addr, err := ParseEndPoint(d.endpoint)
	if err != nil {
		klog.Error(err)
		return
	}
	if proto == "unix" {
		addr = "/" + addr
		if err = os.Remove(addr); err != nil && !os.IsNotExist(err) {
			klog.Fatal(err)
		}
	}

	listen, err := net.Listen(proto, addr)
	if err != nil {
		klog.Fatal(err)
	}
	srv.Serve(listen)
}

func logGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	klog.V(4).Infof("GRPC call: %s ", info.FullMethod)
	klog.V(4).Infof("GRPC request: %v", protosanitizer.StripSecrets(req))
	resp, err := handler(ctx, req)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	klog.V(4).Infof("GRPC response: %s", protosanitizer.StripSecrets(resp))
	return resp, nil
}

func ParseEndPoint(endPoint string) (string, string, error) {
	ep := strings.ToLower(endPoint)
	if strings.HasPrefix(ep, "unix://") || strings.HasPrefix(ep, "tcp://") {
		s := strings.SplitN(ep, "://", 2)
		if s[1] != "" {
			return s[0], s[1], nil
		}
	}
	return "", "", fmt.Errorf("invalid endpoint: %v", endPoint)
}
