package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/extraordy/ocplab-install/pkg/helpers"
	"github.com/urfave/cli"
)

var initPath string
var TfInitFunc = cli.Command{
	Name:        "init",
	Aliases:     []string{"i"},
	Usage:       "Initialize the lab environment",
	Description: "Initialize the lab environment",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "path",
			Usage:       "Define custom init path",
			Destination: &initPath,
		},
	},
	Action: func(c *cli.Context) error {
		tfPath, err := helpers.CheckTerraformBin()
		if err != nil {
			fmt.Println("Terraform binary lookup error: %v", err)
			os.Exit(1)
		}

		err = helpers.CheckLibvirtPlugin()
		if err != nil {
			fmt.Println("Terraform libvirt plugin lookup error: %v", err)
			os.Exit(1)
		}

		if initPath == "" {
			initPath, err = os.Getwd()
			if err != nil {
				fmt.Println("Unable to get current working directory")
				os.Exit(1)
			}
		}
		fileName := fmt.Sprintf("%s/main.tf", initPath)
		err = helpers.GenerateResource(fileName)
		if err != nil {
			fmt.Println("Unable to create the main.tf file in the init directory")
			os.Exit(1)
		}

		cmd := exec.Command(tfPath, "init")

		// Terraform init works by default in the current working directory.
		// Set the Dir field to pass it the custom cwd path.
		cmd.Dir = initPath
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
