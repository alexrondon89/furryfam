package main

import (
	"github.com/alexrondon89/furryfam/infrastructure/internal/cloud"
	"github.com/alexrondon89/furryfam/infrastructure/internal/srv"
)

func main() {
	cloudConfig := cloud.GetCloudConfig("aws_config", "aws", "json")
	cloudInstan := cloud.GetCloudInstance(cloudConfig)
	cloudSrv := srv.NewCloudSrv(cloudInstan.AwsInst)
	cloudSrv.CreateVirtualMachine("cicd")
}
