package incus

import (
	"fmt"
	"yak8s/internal/incus"

	client "github.com/lxc/incus/client"
	"github.com/lxc/incus/shared/api"
)

// VMManager is responsible for managing Incus VMs
type VMManager struct {
	client client.InstanceServer
}

// NewVMManager initializes a new VMManager
func NewVMManager() (*VMManager, error) {
	c, err := incus.ConnectIncus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Incus: %w", err)
	}

	return &VMManager{client: c}, nil
}

// ProvisionVMs provisions the specified number of VMs
func (vm *VMManager) ProvisionVMs(count int) error {
	for i := 1; i <= count; i++ {
		name := fmt.Sprintf("vm-%d", i)
		config := VMConfig{
			Name:       name,
			ImageAlias: "ubuntu/jammy/cloud", // Image alias for Ubuntu 22.04 server
		}

		if err := vm.createVM(config); err != nil {
			return fmt.Errorf("failed to create VM %s: %w", name, err)
		}
		fmt.Printf("VM %s created successfully\n", name)
	}

	return nil
}

// createVM creates an individual VM based on the given configuration
func (vm *VMManager) createVM(config VMConfig) error {
	req := api.InstancesPost{
		Name: config.Name,
		Source: api.InstanceSource{
			Type:     "image",
			Alias:    config.ImageAlias,
			Server:   "https://images.linuxcontainers.org",
			Protocol: "simplestreams",
		},
	}

	// Create the VM
	op, err := vm.client.CreateInstance(req)
	if err != nil {
		return fmt.Errorf("error creating instance: %w", err)
	}

	// Wait for the VM creation operation to complete
	err = op.Wait()
	if err != nil {
		return fmt.Errorf("error waiting for VM creation: %w", err)
	}

	// Start the VM
	op, err = vm.client.UpdateInstanceState(config.Name, api.InstanceStatePut{
		Action:  "start",
		Timeout: -1,
	}, "")
	if err != nil {
		return fmt.Errorf("error starting VM %s: %w", config.Name, err)
	}

	err = op.Wait()
	if err != nil {
		return fmt.Errorf("error waiting for VM %s to start: %w", config.Name, err)
	}

	return nil
}
