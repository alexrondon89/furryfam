package srv

import (
	platfConfig "github.com/alexrondon89/furryfam/infrastructure/config"
	"golang.org/x/crypto/ssh"
)

type CloudInstance interface {
	GetVirtualMachineConfiguration(string) platfConfig.VM
	CreateVirtualMachineInstance(platfConfig.VM) VMInfo
	ConnectToVirtualMachine(platfConfig.VM, VMInfo) *ssh.Client
	InstallDocker(session *ssh.Client)
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
	//	InstanceID: "i-0f3f8972f7a3d7f12",
	//	PublicIP:   "54.172.154.55",
	//	PrivateIP:  "172.31.56.72",
	//}
	sshClient := inst.Srv.ConnectToVirtualMachine(vmConfig, vmInfo)
	defer sshClient.Close()
	inst.Srv.InstallDocker(sshClient)
}
