package cli

import (
	"fmt"
	"yak8s/pkg/incus"
)

// CreateNetworkCommand creates a custom network for yak8s
func RunCreateNetworkCommand(networkName string) error {
	vmManager, err := incus.NewVMManager()
	if err != nil {
		return fmt.Errorf("error initializing VM manager: %w", err)
	}
	if err := vmManager.CreateNetwork(networkName); err != nil {
		return fmt.Errorf("error creating network: %w", err)
	}
	return nil
}

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

// RunDeletionCommand deletes the VMs
func RunDeletionCommand(vmCount int) error {
	// Initialize the VM Manager
	vmManager, err := incus.NewVMManager()
	if err != nil {
		return fmt.Errorf("error initializing VM manager: %w", err)
	}

	// Stop the VMs
	if err := vmManager.RemoveVMs(vmCount); err != nil {
		return fmt.Errorf("error removing VMs: %w", err)
	}
	return nil
}

// Help command
func RunHelpCommand() {
	fmt.Println("Usage: yak8s <command>")
	fmt.Println("Commands:")
	fmt.Println("  up       : Spin up VMs to deploy the cluster.")
	fmt.Println("  down     : Gracefully stop and remove the VMs.")
	fmt.Println("  help     : Show this help message.")
}
