package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/extraordy/ocplab-install/pkg/helpers"
	"github.com/urfave/cli"
)

var TfInitFunc = cli.Command{
	Name:        "init",
	Aliases:     []string{"i"},
	Usage:       "Initialize the lab environment",
	Description: "Initialize the lab environment",
	Action: func(c *cli.Context) error {
		tfPath, err := helpers.CheckTerraformBin()
		if err != nil {
			red.Println("Terraform binary lookup error: %v", err)
			os.Exit(1)
		}

		err = helpers.CheckLibvirtPlugin()
		if err != nil {
			red.Println("Terraform libvirt plugin lookup error: %v", err)
			os.Exit(1)
		}

		cmd := exec.Command(tfPath, "init")

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("\n\n")
		err = cmd.Start()
		if err != nil {
			return err
		}
		err = cmd.Wait()
		if err != nil {
			return err
		}

		cyan.Printf("\nOpenShift lab init completed.\n")
		return nil

	},
}
