// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesis/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the shard count of the specified stream to the specified number of
// shards. Updating the shard count is an asynchronous operation. Upon receiving
// the request, Kinesis Data Streams returns immediately and sets the status of the
// stream to UPDATING. After the update is complete, Kinesis Data Streams sets the
// status of the stream back to ACTIVE. Depending on the size of the stream, the
// scaling action could take a few minutes to complete. You can continue to read
// and write data to your stream while its status is UPDATING. To update the shard
// count, Kinesis Data Streams performs splits or merges on individual shards. This
// can cause short-lived shards to be created, in addition to the final shards.
// These short-lived shards count towards your total shard limit for your account
// in the Region. When using this operation, we recommend that you specify a target
// shard count that is a multiple of 25% (25%, 50%, 75%, 100%). You can specify any
// target value within your shard limit. However, if you specify a target that
// isn't a multiple of 25%, the scaling action might take longer to complete. This
// operation has the following default limits. By default, you cannot do the
// following:
//
//     * Scale more than ten times per rolling 24-hour period per
// stream
//
//     * Scale up to more than double your current shard count for a
// stream
//
//     * Scale down below half your current shard count for a stream
//
//     *
// Scale up to more than 500 shards in a stream
//
//     * Scale a stream with more
// than 500 shards down unless the result is less than 500 shards
//
//     * Scale up
// to more than the shard limit for your account
//
// For the default limits for an AWS
// account, see Streams Limits
// (https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide. To request an increase in
// the call rate limit, the shard limit for this API, or your overall shard limit,
// use the limits form
// (https://console.aws.amazon.com/support/v1#/case/create?issueType=service-limit-increase&limitType=service-code-kinesis).
func (c *Client) UpdateShardCount(ctx context.Context, params *UpdateShardCountInput, optFns ...func(*Options)) (*UpdateShardCountOutput, error) {
	if params == nil {
		params = &UpdateShardCountInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateShardCount", params, optFns, addOperationUpdateShardCountMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateShardCountOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateShardCountInput struct {

	// The scaling type. Uniform scaling creates shards of equal size.
	//
	// This member is required.
	ScalingType types.ScalingType

	// The name of the stream.
	//
	// This member is required.
	StreamName *string

	// The new number of shards. This value has the following default limits. By
	// default, you cannot do the following:
	//
	//     * Set this value to more than double
	// your current shard count for a stream.
	//
	//     * Set this value below half your
	// current shard count for a stream.
	//
	//     * Set this value to more than 500 shards
	// in a stream (the default limit for shard count per stream is 500 per account per
	// region), unless you request a limit increase.
	//
	//     * Scale a stream with more
	// than 500 shards down unless you set this value to less than 500 shards.
	//
	// This member is required.
	TargetShardCount *int32
}

type UpdateShardCountOutput struct {

	// The current number of shards.
	CurrentShardCount *int32

	// The name of the stream.
	StreamName *string

	// The updated number of shards.
	TargetShardCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateShardCountMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateShardCount{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateShardCount{}, middleware.After)
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
	addOpUpdateShardCountValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateShardCount(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateShardCount(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "UpdateShardCount",
	}
}
