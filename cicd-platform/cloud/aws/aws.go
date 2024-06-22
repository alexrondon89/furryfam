package aws

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/spf13/viper"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"

	platfConfig "github.com/alexrondon89/furryfam/cicd-platform/config"
)

type AwsInstance struct {
	platformConfig platfConfig.PlatformConfig
	awsConfig      aws.Config
}

var InstanceTypes = map[string]types.InstanceType{"t2.micro": types.InstanceTypeT2Micro}

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

func (awsInst *AwsInstance) CreateVirtualMachineInstances() {
	ec2Client := ec2.NewFromConfig(awsInst.awsConfig)
	for _, item := range awsInst.platformConfig.Aws.Ec2 {
		input := &ec2.RunInstancesInput{
			ImageId:      aws.String(item.ImageId),
			InstanceType: InstanceTypes[item.InstanceType],
			MinCount:     aws.Int32(item.MinCount),
			MaxCount:     aws.Int32(item.MaxCount),
			TagSpecifications: []types.TagSpecification{
				{
					ResourceType: types.ResourceTypeInstance,
					Tags: []types.Tag{
						{
							Key:   aws.String("Name"),
							Value: aws.String(item.Name),
						},
					},
				},
			},
		}
		result, err := ec2Client.RunInstances(context.TODO(), input)
		if err != nil {
			log.Fatalf("Unable to create instance, %v", err)
		}
		fmt.Printf("Created instance with ID: %s\n", *result.Instances[0].InstanceId)
	}
}
