// +build linux darwin

package main

import (
	"context"
	"os"

	"github.com/nais/device/device-agent/runtimeconfig"
)

func runHelper(rc *runtimeconfig.RuntimeConfig, ctx context.Context) error {
	cmd := adminCommandContext(ctx, "./bin/device-agent-helper",
		"--interface", rc.Config.Interface,
		"--device-ip", rc.BootstrapConfig.DeviceIP,
		"--wireguard-binary", rc.Config.WireGuardBinary,
		"--wireguard-go-binary", rc.Config.WireGuardGoBinary,
		"--wireguard-config-path", rc.Config.WireGuardConfigPath,
		"--log-level", rc.Config.LogLevel,
	)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Start()
}