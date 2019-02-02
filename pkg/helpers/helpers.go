package helpers

import (
	"fmt"
	"os"
	"os/exec"
)

// checkLibvirtPlugin checks the location of the Terraform libvirt plugin
func CheckLibvirtPlugin() error {
	home := os.Getenv("HOME")
	_, err := os.Stat(home + "/.terraform.d/plugins/terraform-provider-libvirt")
	if err != nil {
		return err
	}
	return nil
}

// checkSourceImage is used to verify custom images paths
func CheckSourceImage(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("Error: the provided image could not be found")
	}
	return nil
}

// checkTerraformBin checks the location of the Terrafom binary
func CheckTerraformBin() (string, error) {
	tf, err := exec.LookPath("terraform")
	if err != nil {
		return "", err
	}
	return tf, nil
}
