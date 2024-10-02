package incus

import (
	"fmt"

	client "github.com/lxc/incus/client"
)

// ConnectIncus connects to the Incus Unix socket
func ConnectIncus() (client.InstanceServer, error) {
	c, err := client.ConnectIncusUnix("", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Incus socket: %w", err)
	}
	return c, nil
}
