package cluster

import (
	"fmt"
)

// Cluster holds the informations about the machines to provide
type Cluster struct {
	MasterCount int
	InfraCount  int
	WorkerCount int
	LbCount     int
}

// Todo: print all errors and not only the first catched
// verifyClusterSize tests if the minimal cluster size is fulfilled
func (c *Cluster) VerifyClusterSize() error {
	if c.MasterCount == 0 {
		return fmt.Errorf("Error: The number of master nodes must be at least 1 to correctly deploy the control plane.")
	}
	if c.MasterCount != 0 && (c.MasterCount%2 != 1) {
		return fmt.Errorf("Error: The master nodes count must always be an odd number.")
	}
	if c.InfraCount == 0 && c.MasterCount != 1 && c.WorkerCount != 0 {
		return fmt.Errorf("Error: The number of infra nodes must be at least 1, with the exception of allinone setups (1 master, 0 infra, 0 workers).")
	}
	return nil
}

// setLb sets the Load Balancer variable if masterCount > 1
func (c *Cluster) SetLb() {
	if c.MasterCount != 1 {
		c.LbCount = 1
	} else {
		c.LbCount = 0
	}
}
