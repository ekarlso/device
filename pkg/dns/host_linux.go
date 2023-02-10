package dns

import (
	"bufio"
	"fmt"
	"net/netip"
	"os"
	"regexp"
	"strings"

	nbdns "github.com/netbirdio/netbird/dns"
	log "github.com/sirupsen/logrus"
)

const (
	defaultResolvConfPath = "/tmp/dns/resolv.conf"
)

const (
	netbirdManager osManagerType = iota
	fileManager
	networkManager
	systemdManager
	resolvConfManager
)

type osManagerType int

func newHostManager(wgInterface string) (hostManager, error) {
	osManager, err := getOSDNSManagerType()
	if err != nil {
		return nil, err
	}

	log.Debugf("discovered mode is: %d", osManager)
	switch osManager {
	case networkManager:
		return newNetworkManagerDbusConfigurator(wgInterface)
	case systemdManager:
		return newSystemdDbusConfigurator(wgInterface)
	case resolvConfManager:
		return newResolvConfConfigurator(wgInterface)
	default:
		return newFileConfigurator()
	}
}

func getOSDNSManagerType() (osManagerType, error) {
	file, err := os.Open(defaultResolvConfPath)
	if err != nil {
		return 0, fmt.Errorf("unable to open %s for checking owner, got error: %s", defaultResolvConfPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) == 0 {
			continue
		}
		if text[0] != '#' {
			return fileManager, nil
		}
		if strings.Contains(text, fileGeneratedResolvConfContentHeader) {
			return netbirdManager, nil
		}
		if strings.Contains(text, "NetworkManager") && isDbusListenerRunning(networkManagerDest, networkManagerDbusObjectNode) && isNetworkManagerSupported() {
			log.Debugf("is nm running on supported v? %t", isNetworkManagerSupportedVersion())
			return networkManager, nil
		}
		if strings.Contains(text, "systemd-resolved") && isDbusListenerRunning(systemdResolvedDest, systemdDbusObjectNode) {
			return systemdManager, nil
		}
		if strings.Contains(text, "resolvconf") {
			if isDbusListenerRunning(systemdResolvedDest, systemdDbusObjectNode) {
				var value string
				err = getSystemdDbusProperty(systemdDbusResolvConfModeProperty, &value)
				if err == nil {
					if value == systemdDbusResolvConfModeForeign {
						return systemdManager, nil
					}
				}
				log.Errorf("got an error while checking systemd resolv conf mode, error: %s", err)
			}
			return resolvConfManager, nil
		}
	}
	return fileManager, nil
}

var nsExpression = regexp.MustCompile(`(?m)nameserver ([0-9.])$`)

func DefaultNameServers() ([]nbdns.NameServer, error) {
	b, err := os.ReadFile(defaultResolvConfPath)
	if err != nil {
		return nil, err
	}

	var servers []nbdns.NameServer
	m := nsExpression.FindAllStringSubmatch(string(b), -1)

	for _, match := range m {
		servers = append(servers, nbdns.NameServer{NSType: nbdns.UDPNameServerType, IP: netip.MustParseAddr(match[1]), Port: 53})
	}

	return servers, nil
}
