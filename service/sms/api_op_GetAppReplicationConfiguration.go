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

// Retrieves the application replication configuration associated with the
// specified application.
func (c *Client) GetAppReplicationConfiguration(ctx context.Context, params *GetAppReplicationConfigurationInput, optFns ...func(*Options)) (*GetAppReplicationConfigurationOutput, error) {
	if params == nil {
		params = &GetAppReplicationConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetAppReplicationConfiguration", params, optFns, addOperationGetAppReplicationConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetAppReplicationConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetAppReplicationConfigurationInput struct {

	// The ID of the application.
	AppId *string
}

type GetAppReplicationConfigurationOutput struct {

	// The replication configurations associated with server groups in this
	// application.
	ServerGroupReplicationConfigurations []*types.ServerGroupReplicationConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetAppReplicationConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetAppReplicationConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetAppReplicationConfiguration{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetAppReplicationConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetAppReplicationConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "GetAppReplicationConfiguration",
	}
}
