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

// Describes the specified replication job or all of your replication jobs.
func (c *Client) GetReplicationJobs(ctx context.Context, params *GetReplicationJobsInput, optFns ...func(*Options)) (*GetReplicationJobsOutput, error) {
	if params == nil {
		params = &GetReplicationJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReplicationJobs", params, optFns, addOperationGetReplicationJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReplicationJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReplicationJobsInput struct {

	// The maximum number of results to return in a single call. The default value is
	// 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string

	// The ID of the replication job.
	ReplicationJobId *string
}

type GetReplicationJobsOutput struct {

	// The token required to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the replication jobs.
	ReplicationJobList []*types.ReplicationJob

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetReplicationJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetReplicationJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReplicationJobs{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetReplicationJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetReplicationJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "GetReplicationJobs",
	}
}
