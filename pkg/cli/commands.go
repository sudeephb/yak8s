package cli

import (
	"fmt"
	"yak8s/pkg/incus"
)

// RunProvisionCommand provisions the specified number of VMs
func RunProvisionCommand(vmCount int) error {
	// Initialize the VM Manager
	vmManager, err := incus.NewVMManager()
	if err != nil {
		return fmt.Errorf("error initializing VM manager: %w", err)
	}

	// Provision the VMs
	if err := vmManager.ProvisionVMs(vmCount); err != nil {
		return fmt.Errorf("error provisioning VMs: %w", err)
	}

	return nil
}