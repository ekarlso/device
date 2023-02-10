package dns

import (
	"strings"

	nbdns "github.com/netbirdio/netbird/dns"
)

type hostManager interface {
	applyDNSConfig(config hostDNSConfig) error
	restoreHostDNS() error
}

type hostDNSConfig struct {
	domains    []domainConfig
	routeAll   bool
	serverIP   string
	serverPort int
}

type domainConfig struct {
	domain    string
	matchOnly bool
}

func dnsConfigToHostDNSConfig(dnsConfig nbdns.Config, ip string, port int) hostDNSConfig {
	config := hostDNSConfig{
		routeAll:   false,
		serverIP:   ip,
		serverPort: port,
	}
	for _, nsConfig := range dnsConfig.NameServerGroups {
		if nsConfig.Primary {
			config.routeAll = true
		}

		for _, domain := range nsConfig.Domains {
			config.domains = append(config.domains, domainConfig{
				domain:    strings.TrimSuffix(domain, "."),
				matchOnly: true,
			})
		}
	}

	for _, customZone := range dnsConfig.CustomZones {
		config.domains = append(config.domains, domainConfig{
			domain:    strings.TrimSuffix(customZone.Domain, "."),
			matchOnly: false,
		})
	}

	return config
}
