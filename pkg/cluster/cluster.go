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
		return fmt.Errorf("Error: The number of master nodes must be at least 1.")
	}
	if c.InfraCount == 0 {
		return fmt.Errorf("Error: The number of infra nodes must be at least 1.")
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
