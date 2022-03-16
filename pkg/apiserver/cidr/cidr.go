package cidr

import (
	"fmt"
	"net/netip"
)

func FindAvailableIP(cidr netip.Prefix, allocated []string) (string, error) {
	allocatedMap := toMap(allocated)
	ips, _ := cidrIPs(cidr)
	for _, ip := range ips {
		if _, found := allocatedMap[ip]; !found {
			return ip, nil
		}
	}
	return "", fmt.Errorf("no available IPs in range %v", cidr)
}

func toMap(strings []string) map[string]struct{} {
	m := make(map[string]struct{})
	for _, s := range strings {
		m[s] = struct{}{}
	}
	return m
}

func cidrIPs(cidr netip.Prefix) ([]string, error) {
	var ips []string
	addr := cidr.Addr()
	for ip := addr; cidr.Contains(ip); ip = ip.Next() {
		ips = append(ips, ip.String())
	}

	if cidr.Bits() == 32 {
		return ips, nil
	} else {
		// remove network address and broadcast address
		return ips[1 : len(ips)-1], nil
	}
}
