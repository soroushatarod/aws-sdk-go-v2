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

// Disassociates a set of queues from a routing profile.
func (c *Client) DisassociateRoutingProfileQueues(ctx context.Context, params *DisassociateRoutingProfileQueuesInput, optFns ...func(*Options)) (*DisassociateRoutingProfileQueuesOutput, error) {
	if params == nil {
		params = &DisassociateRoutingProfileQueuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateRoutingProfileQueues", params, optFns, addOperationDisassociateRoutingProfileQueuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateRoutingProfileQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateRoutingProfileQueuesInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The queues to disassociate from this routing profile.
	//
	// This member is required.
	QueueReferences []*types.RoutingProfileQueueReference

	// The identifier of the routing profile.
	//
	// This member is required.
	RoutingProfileId *string
}

type DisassociateRoutingProfileQueuesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateRoutingProfileQueuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateRoutingProfileQueues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateRoutingProfileQueues{}, middleware.After)
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
	addOpDisassociateRoutingProfileQueuesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateRoutingProfileQueues(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateRoutingProfileQueues(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "DisassociateRoutingProfileQueues",
	}
}
