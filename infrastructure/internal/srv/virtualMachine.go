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
	//	InstanceID: "i-0c1458da9208b80ca",
	//	PublicIP:   "54.146.253.45",
	//	PrivateIP:  "172.31.16.47",
	//}
	sshClient := inst.Srv.ConnectToVirtualMachine(vmConfig, vmInfo)
	defer sshClient.Close()
	inst.Srv.InstallDocker(sshClient)
	inst.Srv.CopyFilesToEC2(vmConfig, vmInfo)
	// a new ssh client created to avoid error with docker and the new docker user rights
	sshClientTwo := inst.Srv.ConnectToVirtualMachine(vmConfig, vmInfo)
	defer sshClient.Close()
	inst.Srv.CreateJenkinsContainer(sshClientTwo, vmConfig, vmInfo)
}
