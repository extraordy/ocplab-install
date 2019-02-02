package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/extraordy/ocplab-install/pkg/helpers"
	"github.com/urfave/cli"
)

var TfDestroyFunc = cli.Command{
	Name:        "destroy",
	Aliases:     []string{"d"},
	Usage:       "Destroy the lab environment",
	Description: "Destroy the lab environment",
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

		cmd := exec.Command(tfPath, "destroy", "-auto-approve")

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

		cyan.Printf("\nOpenShift lab destroy finished.\n")
		return nil
	},
}
