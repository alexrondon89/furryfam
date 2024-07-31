package srv

import (
	platfConfig "github.com/alexrondon89/furryfam/infrastructure/config"
	"golang.org/x/crypto/ssh"
)

type CloudInstance interface {
	GetVirtualMachineConfiguration(string) platfConfig.VM
	CreateVirtualMachineInstance(platfConfig.VM) VMInfo
	ConnectToVirtualMachine(platfConfig.VM, VMInfo) *ssh.Client
	InstallDocker(*ssh.Client)
	CopyFilesToEC2(platfConfig.VM, VMInfo)
	CreateJenkinsContainer(*ssh.Client, platfConfig.VM, VMInfo)
}

type VMInfo struct {
	InstanceID string
	PublicIP   string
	PrivateIP  string
}

type CloudSrv struct {
	Srv CloudInstance
}

func NewCloudSrv(cloudSrv CloudInstance) CloudSrv {
	return CloudSrv{
		Srv: cloudSrv,
	}
}

func (inst CloudSrv) CreateVirtualMachine(purpose string) {
	vmConfig := inst.Srv.GetVirtualMachineConfiguration(purpose)
	vmInfo := inst.Srv.CreateVirtualMachineInstance(vmConfig)
	//vmInfo := VMInfo{
	//	InstanceID: "i-04125ed5f1bdaeb7b",
	//	PublicIP:   "52.91.38.142",
	//	PrivateIP:  "172.31.24.71",
	//}
	sshClient := inst.Srv.ConnectToVirtualMachine(vmConfig, vmInfo)
	defer sshClient.Close()
	inst.Srv.InstallDocker(sshClient)
	inst.Srv.CopyFilesToEC2(vmConfig, vmInfo)
	inst.Srv.CreateJenkinsContainer(sshClient, vmConfig, vmInfo)
}
