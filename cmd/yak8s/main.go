package main

import (
	"fmt"
	"os"
	"yak8s/pkg/cli"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: yak8s <command>")
		fmt.Println("Exiting...")
		os.Exit(1)
	}

	numVms := 3
	command := os.Args[1]
	switch command {
	case "up":
		fmt.Println("Running yak8s up...")
		fmt.Println("Spinning up VMs to deploy the cluster.")
		if err := cli.RunProvisionCommand(numVms); err != nil {
			fmt.Printf("Error provisioning VMs: %v\n", err)
			return
		}
		fmt.Printf("Successfully provisioned %d VMs\n", numVms)

	case "down":
		fmt.Println("Running yak8s down...")
		fmt.Println("Gracefully stopping and removing the VMs.")
		if err := cli.RunDeletionCommand(numVms); err != nil {
			fmt.Printf("Error deleting VMs: %v\n", err)
			return
		}
		fmt.Printf("Successfully removed %d VMs\n", numVms)

	case "help":
		cli.RunHelpCommand()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}

}
