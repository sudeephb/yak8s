package main

import (
	"fmt"
	"yak8s/pkg/cli"
)

func main() {
	fmt.Println("Starting yak8s...")

	// Call CLI command handler (provision in this case)
	if err := cli.RunProvisionCommand(3); err != nil {
		fmt.Printf("Error provisioning VMs: %v\n", err)
		return
	}

	fmt.Println("Successfully provisioned 3 VMs")
}
