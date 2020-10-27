// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Lists the consumers registered to receive data from a stream using enhanced
// fan-out, and provides information about each consumer. This operation has a
// limit of 5 transactions per second per stream.
func (c *Client) ListStreamConsumers(ctx context.Context, params *ListStreamConsumersInput, optFns ...func(*Options)) (*ListStreamConsumersOutput, error) {
	if params == nil {
		params = &ListStreamConsumersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreamConsumers", params, optFns, addOperationListStreamConsumersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamConsumersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamConsumersInput struct {

	// The ARN of the Kinesis data stream for which you want to list the registered
	// consumers. For more information, see Amazon Resource Names (ARNs) and AWS
	// Service Namespaces
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#arn-syntax-kinesis-streams).
	//
	// This member is required.
	StreamARN *string

	// The maximum number of consumers that you want a single call of
	// ListStreamConsumers to return.
	MaxResults *int32

	// When the number of consumers that are registered with the data stream is greater
	// than the default value for the MaxResults parameter, or if you explicitly
	// specify a value for MaxResults that is less than the number of consumers that
	// are registered with the data stream, the response includes a pagination token
	// named NextToken. You can specify this NextToken value in a subsequent call to
	// ListStreamConsumers to list the next set of registered consumers. Don't specify
	// StreamName or StreamCreationTimestamp if you specify NextToken because the
	// latter unambiguously identifies the stream. You can optionally specify a value
	// for the MaxResults parameter when you specify NextToken. If you specify a
	// MaxResults value that is less than the number of consumers that the operation
	// returns if you don't specify MaxResults, the response will contain a new
	// NextToken value. You can use the new NextToken value in a subsequent call to the
	// ListStreamConsumers operation to list the next set of consumers. Tokens expire
	// after 300 seconds. When you obtain a value for NextToken in the response to a
	// call to ListStreamConsumers, you have 300 seconds to use that value. If you
	// specify an expired token in a call to ListStreamConsumers, you get
	// ExpiredNextTokenException.
	NextToken *string

	// Specify this input parameter to distinguish data streams that have the same
	// name. For example, if you create a data stream and then delete it, and you later
	// create another data stream with the same name, you can use this input parameter
	// to specify which of the two streams you want to list the consumers for. You
	// can't specify this parameter if you specify the NextToken parameter.
	StreamCreationTimestamp *time.Time
}

type ListStreamConsumersOutput struct {

	// An array of JSON objects. Each object represents one registered consumer.
	Consumers []*types.Consumer

	// When the number of consumers that are registered with the data stream is greater
	// than the default value for the MaxResults parameter, or if you explicitly
	// specify a value for MaxResults that is less than the number of registered
	// consumers, the response includes a pagination token named NextToken. You can
	// specify this NextToken value in a subsequent call to ListStreamConsumers to list
	// the next set of registered consumers. For more information about the use of this
	// pagination token when calling the ListStreamConsumers operation, see
	// ListStreamConsumersInput$NextToken. Tokens expire after 300 seconds. When you
	// obtain a value for NextToken in the response to a call to ListStreamConsumers,
	// you have 300 seconds to use that value. If you specify an expired token in a
	// call to ListStreamConsumers, you get ExpiredNextTokenException.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListStreamConsumersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListStreamConsumers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListStreamConsumers{}, middleware.After)
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
	addOpListStreamConsumersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListStreamConsumers(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListStreamConsumers(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "ListStreamConsumers",
	}
}
