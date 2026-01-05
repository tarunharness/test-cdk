package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type TestCdkStackProps struct {
	awscdk.StackProps
}

func NewTestCdkStack(scope constructs.Construct, id string, props *TestCdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// Define the Lambda function resource
	myFunction := awslambda.NewFunction(stack, jsii.String("HelloWorldFunction"), &awslambda.FunctionProps{
		Runtime: awslambda.Runtime_NODEJS_20_X(), // Provide any supported Node.js runtime
		Handler: jsii.String("index.handler"),
		// Role:    role,
		Code: awslambda.Code_FromInline(jsii.String(`
		  exports.handler = async function(event) {
			return {
			  statusCode: 200,
			  body: JSON.stringify('Hello Harnessians! What a lovely day'),
			};
		  };
		`)),
	})

	// Define the Lambda function URL resource
	myFunctionUrl := myFunction.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})

	// Define a CloudFormation output for your URL
	awscdk.NewCfnOutput(stack, jsii.String("myFunctionUrlOutput"), &awscdk.CfnOutputProps{
		Value: myFunctionUrl.Url(),
	})

	// The code that defines your stack goes here
	awss3.NewBucket(stack, jsii.String("MyFirstBucket"), &awss3.BucketProps{
		// Optional: Specify a globally unique bucket name. If left out, CDK generates one.
		BucketName: jsii.String("test-cdk-bucket-01"),

		// Optional: Define a removal policy. DESTROY deletes the bucket when the stack is destroyed.
		// Be careful with this in production for non-empty buckets.
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	// The code that defines your stack goes here
	awss3.NewBucket(stack, jsii.String("MySecondBucket"), &awss3.BucketProps{
		// Optional: Specify a globally unique bucket name. If left out, CDK generates one.
		BucketName: jsii.String("test-cdk-bucket-02"),

		// Optional: Define a removal policy. DESTROY deletes the bucket when the stack is destroyed.
		// Be careful with this in production for non-empty buckets.
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	// The code that defines your stack goes here
	awss3.NewBucket(stack, jsii.String("MyThirdBucket"), &awss3.BucketProps{
		// Optional: Specify a globally unique bucket name. If left out, CDK generates one.
		BucketName: jsii.String("test-cdk-bucket-03"),

		// Optional: Define a removal policy. DESTROY deletes the bucket when the stack is destroyed.
		// Be careful with this in production for non-empty buckets.
		RemovalPolicy: awscdk.RemovalPolicy_DESTROY,
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewTestCdkStack(app, "TestCdkStack", &TestCdkStackProps{
		awscdk.StackProps{
			Env: env(),
			// Synthesizer: synth,
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	// If unspecified, this stack will be "environment-agnostic".
	// Account/Region-dependent features and context lookups will not work, but a
	// single synthesized template can be deployed anywhere.
	//---------------------------------------------------------------------------
	return &awscdk.Environment{
		Account: jsii.String("665453390054"),
		Region:  jsii.String("eu-west-1"),
	}
}
