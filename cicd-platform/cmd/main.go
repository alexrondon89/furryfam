package main

import (
	"fmt"
	"github.com/alexrondon89/furryfam/cicd-platform/cloud"
	"github.com/alexrondon89/furryfam/cicd-platform/internal/srv"
)

func main() {
	cloudConfig := cloud.GetCloudConfig("aws_config", "aws", "json")
	cloudInstan := cloud.GetCloudInstance(cloudConfig)
	cloudSrv := srv.NewCloudSrv(cloudInstan.AwsInst)
	cloudSrv.CreateVirtualMachines()
	fmt.Println("Created instance")
}
