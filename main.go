package main

import (
	"flag"

	"github.com/0xff-dev/csi-example/pkg"
	"k8s.io/klog"
)

var (
	endpoint = flag.String("endpoint", "", "csi endpoint")
	nodeID   = flag.String("nodeid", "", "node id")
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	s := pkg.NewDriver(*nodeID, *endpoint)
	s.Run()
}
