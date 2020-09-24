package serverless

import (
	"github.com/awslabs/goformation/v4/cloudformation/policies"
)

// HttpApi_HttpApiDefinition AWS CloudFormation Resource (AWS::Serverless::HttpApi.HttpApiDefinition)
// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-httpapi-httpapidefinition.html
type HttpApi_HttpApiDefinition struct {

	// Bucket AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-httpapi-httpapidefinition.html
	Bucket string `json:"Bucket,omitempty"`

	// Key AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-httpapi-httpapidefinition.html
	Key string `json:"Key,omitempty"`

	// Version AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-httpapi-httpapidefinition.html
	Version int `json:"Version,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *HttpApi_HttpApiDefinition) AWSCloudFormationType() string {
	return "AWS::Serverless::HttpApi.HttpApiDefinition"
}
