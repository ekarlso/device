package gateway_agent

import (
	"github.com/nais/device/pkg/pb"
	log "github.com/sirupsen/logrus"
)

type noopConfigurer struct{}

func NewNoOpConfigurer() NetworkConfigurer {
	return &noopConfigurer{}
}

func (n *noopConfigurer) ApplyWireGuardConfig(devices []*pb.Device) error {
	log.Debugf("Applying WireGuard configuration with %d devices", len(devices))
	for i, device := range devices {
		log.Debugf("(%02d) %s", i+1, device.String())
	}
	return nil
}

func (n *noopConfigurer) ForwardRoutes(routes []string) error {
	log.Debugf("Applying %d forwarding routes:", len(routes))
	for i, route := range routes {
		log.Debugf("(%02d) %s", i+1, route)
	}
	return nil
}

func (n *noopConfigurer) ConnectedDeviceCount() (int, error) {
	return 0, nil
}

func (n *noopConfigurer) SetupInterface() error {
	log.Debugf("SetupInterface()")
	return nil
}

func (n *noopConfigurer) SetupIPTables() error {
	log.Debugf("SetupIPTables()")
	return nil
}

var _ NetworkConfigurer = &noopConfigurer{}
