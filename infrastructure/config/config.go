package config

type PlatformConfig struct {
	Service string
	Aws     *Aws
}

type Aws struct {
	Region string
	Ec2    map[string]VM
}

type VM struct {
	User         string
	Name         string
	ImageId      string
	InstanceType string
	MinCount     int32
	MaxCount     int32
	KeyName      string
	KeyLocation  string
}
