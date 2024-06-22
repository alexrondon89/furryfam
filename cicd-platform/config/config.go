package config

type PlatformConfig struct {
	Service string
	Aws     *aws
}

type aws struct {
	Region string
	Ec2    []ec2
}

type ec2 struct {
	Name         string
	ImageId      string
	InstanceType string
	MinCount     int32
	MaxCount     int32
}
