package wireguard

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/coreos/go-iptables/iptables"
	log "github.com/sirupsen/logrus"
)

type NetworkConfigurer interface {
	ApplyWireGuardConfig(peers []Peer) error
	ForwardRoutes(routes []string) error
	ConnectedDeviceCount() (int, error)
	SetupInterface() error
	SetupIPTables() error
}

type networkConfigurer struct {
	config        *Config
	ipTables      *iptables.IPTables
	interfaceName string
	interfaceIP   string
	configPath    string
	tunnelIP      string
}

func NewConfigurer(configPath, tunnelIP, privateKey, intf string, listenPort int, ipTables *iptables.IPTables) NetworkConfigurer {
	return &networkConfigurer{
		config: &Config{
			PrivateKey: privateKey,
			ListenPort: listenPort,
		},
		configPath:    configPath,
		interfaceName: intf,
		ipTables:      ipTables,
		tunnelIP:      tunnelIP,
	}
}

func (nc *networkConfigurer) SetupInterface() error {
	if err := exec.Command("ip", "link", "del", nc.interfaceName).Run(); err != nil {
		log.Infof("pre-deleting WireGuard interface (ok if this fails): %v", err)
	}

	run := func(commands [][]string) error {
		for _, s := range commands {
			cmd := exec.Command(s[0], s[1:]...)

			if out, err := cmd.CombinedOutput(); err != nil {
				return fmt.Errorf("running %v: %w: %v", cmd, err, string(out))
			} else {
				log.Debugf("%v: %v\n", cmd, string(out))
			}
		}
		return nil
	}

	commands := [][]string{
		{"ip", "link", "add", "dev", nc.interfaceName, "type", "wireguard"},
		{"ip", "link", "set", nc.interfaceName, "mtu", "1360"},
		{"ip", "address", "add", "dev", nc.interfaceName, nc.tunnelIP + "/21"},
		{"ip", "link", "set", nc.interfaceName, "up"},
	}

	return run(commands)
}

// ApplyWireGuardConfig runs syncconfig with the provided WireGuard config
func (nc *networkConfigurer) ApplyWireGuardConfig(peers []Peer) error {
	configFile, err := os.OpenFile(nc.configPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("open WireGuard config file: %w", err)
	}
	defer configFile.Close()

	nc.config.Peers = peers
	err = nc.config.MarshalINI(configFile)
	if err != nil {
		return fmt.Errorf("write WireGuard config: %w", err)
	}

	err = configFile.Close()
	if err != nil {
		return fmt.Errorf("close WireGuard config: %w", err)
	}

	cmd := exec.Command("wg", "syncconf", nc.interfaceName, nc.configPath)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("sync WireGuard config: %w", err)
	}

	log.Debugf("Actuated WireGuard config at %v", nc.configPath)

	return nil
}

func (nc *networkConfigurer) ConnectedDeviceCount() (int, error) {
	output, err := exec.Command("wg", "show", nc.interfaceName, "endpoints").Output()
	if err != nil {
		return 0, err
	}

	re := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}:\d{1,5}`)
	matches := re.FindAll(output, -1)

	return len(matches), nil
}
