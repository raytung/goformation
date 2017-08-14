package resources

// AWS::OpsWorks::Instance AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html
type AWSOpsWorksInstance struct {

	// AgentVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-agentversion
	AgentVersion string `json:"AgentVersion"`

	// AmiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-amiid
	AmiId string `json:"AmiId"`

	// Architecture AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-architecture
	Architecture string `json:"Architecture"`

	// AutoScalingType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-autoscalingtype
	AutoScalingType string `json:"AutoScalingType"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-availabilityzone
	AvailabilityZone string `json:"AvailabilityZone"`

	// BlockDeviceMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-blockdevicemappings
	BlockDeviceMappings []AWSOpsWorksInstanceBlockDeviceMapping `json:"BlockDeviceMappings"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-ebsoptimized
	EbsOptimized bool `json:"EbsOptimized"`

	// ElasticIps AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-elasticips
	ElasticIps []AWSOpsWorksInstancestring `json:"ElasticIps"`

	// Hostname AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-hostname
	Hostname string `json:"Hostname"`

	// InstallUpdatesOnBoot AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-installupdatesonboot
	InstallUpdatesOnBoot bool `json:"InstallUpdatesOnBoot"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-instancetype
	InstanceType string `json:"InstanceType"`

	// LayerIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-layerids
	LayerIds []AWSOpsWorksInstancestring `json:"LayerIds"`

	// Os AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-os
	Os string `json:"Os"`

	// RootDeviceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-rootdevicetype
	RootDeviceType string `json:"RootDeviceType"`

	// SshKeyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-sshkeyname
	SshKeyName string `json:"SshKeyName"`

	// StackId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-stackid
	StackId string `json:"StackId"`

	// SubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-subnetid
	SubnetId string `json:"SubnetId"`

	// Tenancy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-tenancy
	Tenancy string `json:"Tenancy"`

	// TimeBasedAutoScaling AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-timebasedautoscaling
	TimeBasedAutoScaling AWSOpsWorksInstanceTimeBasedAutoScaling `json:"TimeBasedAutoScaling"`

	// VirtualizationType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-virtualizationtype
	VirtualizationType string `json:"VirtualizationType"`

	// Volumes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-instance.html#cfn-opsworks-instance-volumes
	Volumes []AWSOpsWorksInstancestring `json:"Volumes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksInstance) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Instance"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksInstance) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
