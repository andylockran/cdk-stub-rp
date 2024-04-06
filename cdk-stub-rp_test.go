package main

import (
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
)

// example tests. To run these tests, uncomment this file along with the
// example resource in cdk-stub-rp_test.go
func TestCdkStubRpStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewCdkStubRpStack(app, "MyStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack, nil)

	template.HasResourceProperties(jsii.String("AWS::DynamoDB::Table"), map[string]interface{}{
		"TimeToLiveSpecification": map[string]*string{
			"AttributeName": jsii.String("expiry"),
		},
	})
}
