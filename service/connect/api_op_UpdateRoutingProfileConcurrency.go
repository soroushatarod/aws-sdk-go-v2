// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the channels that agents can handle in the Contact Control Panel (CCP)
// for a routing profile.
func (c *Client) UpdateRoutingProfileConcurrency(ctx context.Context, params *UpdateRoutingProfileConcurrencyInput, optFns ...func(*Options)) (*UpdateRoutingProfileConcurrencyOutput, error) {
	if params == nil {
		params = &UpdateRoutingProfileConcurrencyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoutingProfileConcurrency", params, optFns, addOperationUpdateRoutingProfileConcurrencyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRoutingProfileConcurrencyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoutingProfileConcurrencyInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The channels agents can handle in the Contact Control Panel (CCP).
	//
	// This member is required.
	MediaConcurrencies []*types.MediaConcurrency

	// The identifier of the routing profile.
	//
	// This member is required.
	RoutingProfileId *string
}

type UpdateRoutingProfileConcurrencyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateRoutingProfileConcurrencyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRoutingProfileConcurrency{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRoutingProfileConcurrency{}, middleware.After)
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
	addOpUpdateRoutingProfileConcurrencyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoutingProfileConcurrency(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateRoutingProfileConcurrency(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "UpdateRoutingProfileConcurrency",
	}
}
