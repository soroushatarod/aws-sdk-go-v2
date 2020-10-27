// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the application launch configuration associated with the specified
// application.
func (c *Client) GetAppLaunchConfiguration(ctx context.Context, params *GetAppLaunchConfigurationInput, optFns ...func(*Options)) (*GetAppLaunchConfigurationOutput, error) {
	if params == nil {
		params = &GetAppLaunchConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAppLaunchConfiguration", params, optFns, addOperationGetAppLaunchConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAppLaunchConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAppLaunchConfigurationInput struct {

	// The ID of the application.
	AppId *string
}

type GetAppLaunchConfigurationOutput struct {

	// The ID of the application.
	AppId *string

	// Indicates whether the application is configured to launch automatically after
	// replication is complete.
	AutoLaunch *bool

	// The name of the service role in the customer's account that AWS CloudFormation
	// uses to launch the application.
	RoleName *string

	// The launch configurations for server groups in this application.
	ServerGroupLaunchConfigurations []*types.ServerGroupLaunchConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAppLaunchConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAppLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAppLaunchConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAppLaunchConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAppLaunchConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "GetAppLaunchConfiguration",
	}
}
