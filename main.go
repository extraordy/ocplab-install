package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

const appDescription = `ocplab-install is a tool to create an OpenShift cluster
						using Terraform and Libvirt provider.`

var (
	sourceImage string
	args        []string
)

// Cluster holds the informations about the machines to provide
type Cluster struct {
	masterCount int
	infraCount  int
	workerCount int
	lbCount     int
}

// Todo: print all errors and not only the first catched
// verifyClusterSize tests if the minimal cluster size is fulfilled
func (c *Cluster) verifyClusterSize() error {
	if c.masterCount == 0 {
		return fmt.Errorf("Error: The number of master nodes must be at least 1.")
	}
	if c.infraCount == 0 {
		return fmt.Errorf("Error: The number of infra nodes must be at least 1.")
	}
	return nil
}

// setLb sets the Load Balancer variable if masterCount > 1
func (c *Cluster) setLb() {
	if c.masterCount != 1 {
		c.lbCount = 1
	} else {
		c.lbCount = 0
	}
}

// checkTerraformBin checks the location of the Terrafom binary
func checkTerraformBin() (string, error) {
	tf, err := exec.LookPath("terraform")
	if err != nil {
		return "", err
	}
	return tf, nil
}

// checkLibvirtPlugin checks the location of the Terraform libvirt plugin
func checkLibvirtPlugin() error {
	home := os.Getenv("HOME")
	_, err := os.Stat(home + "/.terraform.d/plugins/terraform-provider-libvirt")
	if err != nil {
		return err
	}
	return nil
}

// checkSourceImage is used to verify custom images paths
func checkSourceImage(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Error: the provided image could not be found")
	}
	return nil
}

func main() {

	cyan := color.New(color.FgCyan)
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)

	tfPath, err := checkTerraformBin()
	if err != nil {
		red.Println("Terraform binary lookup error: %v", err)
		os.Exit(1)
	}

	err = checkLibvirtPlugin()
	if err != nil {
		red.Println("Terraform libvirt plugin lookup error: %v", err)
		os.Exit(1)
	}

	app := cli.NewApp()
	app.Name = "ocplab-install"
	app.Description = appDescription
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
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
				cl := &Cluster{}

				cyan.Println("OpenShift lab provisioning for libvirt Terraform provider")
				green.Printf("Number of masters nodes: ")
				fmt.Scanf("%d", &cl.masterCount)
				green.Printf("Number of infra nodes: ")
				fmt.Scanf("%d", &cl.infraCount)
				green.Printf("Number of worker nodes: ")
				fmt.Scanf("%d", &cl.workerCount)
				err = cl.verifyClusterSize()
				if err != nil {
					return err
				}
				cl.setLb()

				masterStr := strconv.Itoa(cl.masterCount)
				infraStr := strconv.Itoa(cl.infraCount)
				workerStr := strconv.Itoa(cl.workerCount)
				lbStr := strconv.Itoa(cl.lbCount)

				args = append(args, "apply", "-auto-approve",
					"-var", "master-count="+masterStr,
					"-var", "infra-count="+infraStr,
					"-var", "worker-count="+workerStr,
					"-var", "lb-count="+lbStr)

				if sourceImage != "" {
					err = checkSourceImage(sourceImage)
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
		},
		{
			Name:        "destroy",
			Aliases:     []string{"d"},
			Usage:       "Destroy the lab environment",
			Description: "Destroy the lab environment",
			Action: func(c *cli.Context) error {

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
		},
		{
			Name:        "init",
			Aliases:     []string{"i"},
			Usage:       "Initialize the lab environment",
			Description: "Initialize the lab environment",
			Action: func(c *cli.Context) error {

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
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
