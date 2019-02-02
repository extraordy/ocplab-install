package main

import (
	"fmt"
	"os"

	"github.com/extraordy/ocplab-install/pkg/cmd"
	"github.com/urfave/cli"
)

const appDescription = `ocplab-install is a tool to create an OpenShift cluster
						using Terraform and Libvirt provider.`

func main() {
	app := cli.NewApp()
	app.Name = "ocplab-install"
	app.Description = appDescription
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		cmd.TfCreateFunc,
		cmd.TfDestroyFunc,
		cmd.TfInitFunc,
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
