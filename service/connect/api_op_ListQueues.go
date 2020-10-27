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

// Provides information about the queues for the specified Amazon Connect instance.
// For more information about queues, see Queues: Standard and Agent
// (https://docs.aws.amazon.com/connect/latest/adminguide/concepts-queues-standard-and-agent.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) ListQueues(ctx context.Context, params *ListQueuesInput, optFns ...func(*Options)) (*ListQueuesOutput, error) {
	if params == nil {
		params = &ListQueuesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListQueues", params, optFns, addOperationListQueuesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListQueuesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListQueuesInput struct {

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The maximimum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The type of queue.
	QueueTypes []types.QueueType
}

type ListQueuesOutput struct {

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the queues.
	QueueSummaryList []*types.QueueSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListQueuesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListQueues{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListQueues{}, middleware.After)
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
	addOpListQueuesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListQueues(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListQueues(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "ListQueues",
	}
}
