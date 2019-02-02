package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/extraordy/ocplab-install/pkg/cluster"
	"github.com/extraordy/ocplab-install/pkg/helpers"
	"github.com/urfave/cli"
)

var sourceImage string
var TfCreateFunc = cli.Command{
	Name:        "create",
	Aliases:     []string{"c"},
	Usage:       "Create a new lab environment",
	Description: "Create a new lab environment",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "image, img",
			Usage:       "Define custom image path",
			Destination: &sourceImage,
		},
	},
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

		cl := &cluster.Cluster{}
		args := make([]string, 0)

		cyan.Println("OpenShift lab provisioning for libvirt Terraform provider")
		green.Printf("Number of masters nodes: ")
		fmt.Scanf("%d", &cl.MasterCount)
		green.Printf("Number of infra nodes: ")
		fmt.Scanf("%d", &cl.InfraCount)
		green.Printf("Number of worker nodes: ")
		fmt.Scanf("%d", &cl.WorkerCount)
		err = cl.VerifyClusterSize()
		if err != nil {
			return err
		}
		cl.SetLb()

		masterStr := strconv.Itoa(cl.MasterCount)
		infraStr := strconv.Itoa(cl.InfraCount)
		workerStr := strconv.Itoa(cl.WorkerCount)
		lbStr := strconv.Itoa(cl.LbCount)

		args = append(args, "apply", "-auto-approve",
			"-var", "master-count="+masterStr,
			"-var", "infra-count="+infraStr,
			"-var", "worker-count="+workerStr,
			"-var", "lb-count="+lbStr)

		if sourceImage != "" {
			err = helpers.CheckSourceImage(sourceImage)
			if err != nil {
				return err
			}
			args = append(args, "-var", "source-image="+sourceImage)
		}

		cmd := exec.Command(tfPath, args...)

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

		cyan.Printf("\nOpenShift lab setup finished.\n")
		return nil
	},
}
