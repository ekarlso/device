package main

import (
	"context"
	"fmt"
	"net/netip"
	"os"
	"os/signal"
	"syscall"

	mdns "github.com/miekg/dns"
	"github.com/nais/device/pkg/dns"
	nbdns "github.com/netbirdio/netbird/dns"
	"github.com/sirupsen/logrus"
)

func main() {
	ns, err := dns.DefaultNameServers()
	if err != nil {
		panic(err)
	}

	if len(ns) == 0 {
		fmt.Println("No nameservers found, defaulting to Cloudflare DNS")
		ns = []nbdns.NameServer{
			{NSType: nbdns.UDPNameServerType, IP: netip.MustParseAddr("8.8.8.8"), Port: 53},
		}
	}

	ctx := context.Background()

	logrus.SetLevel(logrus.TraceLevel)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill, os.Signal(syscall.SIGTERM))
	defer cancel()

	server, err := dns.NewDefaultServer(ctx, "utun69", "")
	if err != nil {
		panic(err)
	}

	err = server.UpdateDNSServer(3, nbdns.Config{
		ServiceEnable: true,
		NameServerGroups: []*nbdns.NameServerGroup{
			{
				Primary:     true,
				NameServers: ns,
				Domains:     []string{},
			},
		},
		CustomZones: []nbdns.CustomZone{
			{
				Domain: "dev.intern.nav.no",
				Records: []nbdns.SimpleRecord{
					{Name: "dev.intern.nav.no", Type: int(mdns.TypeA), Class: nbdns.DefaultClass, TTL: 60, RData: "127.0.0.1"},
					{Name: "*.dev.intern.nav.no", Type: int(mdns.TypeA), Class: nbdns.DefaultClass, TTL: 60, RData: "127.0.0.2"},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("START")

	<-ctx.Done()

	server.Stop()
}
