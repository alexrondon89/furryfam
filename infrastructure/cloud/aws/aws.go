package aws

import (
	"context"
	"fmt"
	"github.com/alexrondon89/furryfam/infrastructure/internal/util"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"log"
	"time"

	platfConfig "github.com/alexrondon89/furryfam/infrastructure/config"
	"github.com/alexrondon89/furryfam/infrastructure/internal/srv"
)

type AwsInstance struct {
	platformConfig platfConfig.PlatformConfig
	awsConfig      aws.Config
}

var InstanceTypes = map[string]types.InstanceType{"t2.medium": types.InstanceTypeT2Medium}

func CreateAwsInstance(platformConfig platfConfig.PlatformConfig) *AwsInstance {
	awsInst := &AwsInstance{
		platformConfig: platformConfig,
	}
	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(), func(lc *awsConfig.LoadOptions) error {
		lc.Region = platformConfig.Aws.Region
		lc.Credentials = awsInst
		return nil
	})
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	awsInst.awsConfig = cfg
	return awsInst
}

func (awsInst *AwsInstance) Retrieve(ctx context.Context) (aws.Credentials, error) {
	cred := aws.Credentials{}
	viper.BindEnv("AWS_ACCESS_KEY_ID")
	viper.BindEnv("AWS_SECRET_ACCESS_KEY")
	accessKey := viper.GetString("AWS_ACCESS_KEY_ID")
	secretKey := viper.GetString("AWS_SECRET_ACCESS_KEY")
	cred.AccessKeyID = accessKey
	cred.SecretAccessKey = secretKey
	return cred, nil
}

func (awsInst *AwsInstance) GetVirtualMachineConfiguration(purpose string) platfConfig.VM {
	ec2Obj, ok := awsInst.platformConfig.Aws.Ec2[purpose]
	if !ok {
		log.Fatalf("Unable to get instance config for purpose %s", purpose)
	}
	return ec2Obj
}

func (awsInst *AwsInstance) CreateVirtualMachineInstance(vm platfConfig.VM) srv.VMInfo {
	ec2Client := ec2.NewFromConfig(awsInst.awsConfig)
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(vm.ImageId),
		InstanceType: InstanceTypes[vm.InstanceType],
		MinCount:     aws.Int32(vm.MinCount),
		MaxCount:     aws.Int32(vm.MaxCount),
		KeyName:      aws.String(vm.KeyName),
		NetworkInterfaces: []types.InstanceNetworkInterfaceSpecification{
			{
				AssociatePublicIpAddress: aws.Bool(true),
				DeviceIndex:              aws.Int32(0),
			},
		},
		TagSpecifications: []types.TagSpecification{
			{
				ResourceType: types.ResourceTypeInstance,
				Tags: []types.Tag{
					{
						Key:   aws.String("Name"),
						Value: aws.String(vm.Name),
					},
				},
			},
		},
	}
	result, err := ec2Client.RunInstances(context.TODO(), input)
	if err != nil {
		log.Fatalf("Unable to create instance, %v", err)
	}
	instanceId := *result.Instances[0].InstanceId
	fmt.Printf("Created instance with ID: %s\n", instanceId)

	describeInstancesOutput := awsInst.getDescribeInstance(ec2Client, instanceId)
	instance := describeInstancesOutput.Reservations[0].Instances[0]
	//securityGroupId := instance.SecurityGroups[0].GroupId
	//awsInst.authorizeSecurityGroup(ec2Client, *securityGroupId)

	return srv.VMInfo{
		InstanceID: *instance.InstanceId,
		PrivateIP:  *instance.PrivateIpAddress,
		PublicIP:   *instance.PublicIpAddress,
	}
}

func (awsInst *AwsInstance) ConnectToVirtualMachine(vm platfConfig.VM, vmInfo srv.VMInfo) *ssh.Client {
	signer := util.GetSigner(vm.KeyLocation)
	sshClient := util.GetSshClient(signer, vm.User, vmInfo.PublicIP)
	return sshClient
}

func (awsInst *AwsInstance) InstallDocker(sshClient *ssh.Client) {
	// Commands to install Docker
	cmds := []string{
		"sudo apt-get update",
		"sudo apt-get install docker.io -y",
		"sudo systemctl start docker",
		"sudo systemctl enable docker",
	}

	for _, cmd := range cmds {
		session, err := sshClient.NewSession()
		if err != nil {
			log.Fatal("Failed to create session: ", err)
		}

		output, err := session.CombinedOutput(cmd)
		if err != nil {
			log.Fatalf("failed to run command: %s", err)
		}
		fmt.Printf(string(output))
		session.Close()
	}
	fmt.Printf("docker installed")
}

func (awsInst *AwsInstance) authorizeSecurityGroup(ec2Client *ec2.Client, securityGroupID string) {
	authorizeSecurityGroupIngressInput := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(securityGroupID),
		IpPermissions: []types.IpPermission{
			{
				IpProtocol: aws.String("tcp"),
				FromPort:   aws.Int32(22),
				ToPort:     aws.Int32(22),
				IpRanges: []types.IpRange{
					{
						CidrIp:      aws.String("0.0.0.0/0"),
						Description: aws.String("Allow SSH"),
					},
				},
			},
		},
	}

	_, err := ec2Client.AuthorizeSecurityGroupIngress(context.TODO(), authorizeSecurityGroupIngressInput)
	if err != nil {
		log.Fatalf("Failed to authorize security group ingress: %v", err)
	}

	fmt.Println("Successfully updated security group to allow SSH")
}

func (awsInst *AwsInstance) getDescribeInstance(ec2Client *ec2.Client, instanceId string) ec2.DescribeInstancesOutput {
	var describeInstancesOutput *ec2.DescribeInstancesOutput
	var err error
	describeInstancesInput := &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceId},
	}

	for i := 0; i <= 5; i++ {
		describeInstancesOutput, err = ec2Client.DescribeInstances(context.TODO(), describeInstancesInput)
		if err != nil {
			log.Fatalf("Failed to describe instances: %v", err)
		}
		if describeInstancesOutput.Reservations[0].Instances[0].PublicIpAddress != nil {
			break
		}
		fmt.Printf("waiting for PublicIpAddress assigned....")
		time.Sleep(3 * time.Second)
	}
	fmt.Printf("describe Instances: PublicIP %s, PrivateIP %s",
		*describeInstancesOutput.Reservations[0].Instances[0].PublicIpAddress,
		*describeInstancesOutput.Reservations[0].Instances[0].PrivateIpAddress,
	)
	return *describeInstancesOutput
}
