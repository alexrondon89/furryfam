package srv

type CloudInstance interface {
	CreateVirtualMachineInstances()
}

type CloudSrv struct {
	Srv CloudInstance
}

func NewCloudSrv(cloudSrv CloudInstance) CloudSrv {
	return CloudSrv{
		Srv: cloudSrv,
	}
}

func (inst CloudSrv) CreateVirtualMachines() {
	inst.Srv.CreateVirtualMachineInstances()
}
