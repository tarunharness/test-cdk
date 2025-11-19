package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
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

	// The code that defines your stack goes here

	// example resource
	// awssqs.NewQueue(stack, jsii.String("TestCdkQueue"), &awssqs.QueueProps{
	// VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	// })

	// 1. Define the IAM Role for the Lambda function
	// _ = awsiam.NewRole(stack, jsii.String("MyCustomLambdaRole"), &awsiam.RoleProps{
	// 	AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	// 	ManagedPolicies: &[]awsiam.IManagedPolicy{
	// 		awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("service-role/AWSLambdaBasicExecutionRole")),
	// 	},
	// })

	// Reference an existing IAM role by its ARN

	// existingRoleArn := "arn:aws:iam::665453390054:role/aws-reserved/sso.amazonaws.com/AWSReservedSSO_AWSPowerUserAccess_6b6c4ac271ad17fb" // Replace with your role ARN
	// existingRole := awsiam.Role_FromRoleArn(stack, jsii.String("ExistingLambdaRole"), jsii.String(existingRoleArn), nil)
	// existingRole.AttachInlinePolicy(awsiam.Policy_FromPolicyName(stack, jsii.String("ExistingLambdaPolicy"), nil))

	// Create role:

	// role := awsiam.NewRole(stack, jsii.String("LambdaExecRole"), &awsiam.RoleProps{
	// 	AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
	// 	ManagedPolicies: &[]awsiam.IManagedPolicy{
	// 		awsiam.ManagedPolicy_FromAwsManagedPolicyName(jsii.String("PowerUserAccess")),
	// 	},
	// })

	// Use role:
	// Role: role,

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

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	// synth := awscdk.NewDefaultStackSynthesizer(&awscdk.DefaultStackSynthesizerProps{
	// CloudFormationExecutionRole: jsii.String("arn:aws:iam::665453390054:role/aws-reserved/sso.amazonaws.com/AWSReservedSSO_AWSPowerUserAccess_6b6c4ac271ad17fb"),
	// CloudFormationExecutionRole: jsii.String("arn:aws:iam::665453390054:role/cdk-hnb659fds-deploy-role-665453390054-us-east-1"),
	// DeployRoleArn:               jsii.String("arn:aws:iam::665453390054:role/cdk-hnb659fds-deploy-role-665453390054-us-east-1"),
	// })

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
		Region:  jsii.String("us-east-1"),
	}

	// Uncomment if you know exactly what account and region you want to deploy
	// the stack to. This is the recommendation for production stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String("123456789012"),
	//  Region:  jsii.String("us-east-1"),
	// }

	// Uncomment to specialize this stack for the AWS Account and Region that are
	// implied by the current CLI configuration. This is recommended for dev
	// stacks.
	//---------------------------------------------------------------------------
	// return &awscdk.Environment{
	//  Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
	//  Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	// }
}
