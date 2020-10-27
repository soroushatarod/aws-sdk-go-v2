// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the policy status for an Amazon S3 bucket, indicating whether the
// bucket is public. In order to use this operation, you must have the
// s3:GetBucketPolicyStatus permission. For more information about Amazon S3
// permissions, see Specifying Permissions in a Policy
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html).
// For more information about when Amazon S3 considers a bucket public, see The
// Meaning of "Public"
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status).
// The following operations are related to GetBucketPolicyStatus:
//
//     * Using
// Amazon S3 Block Public Access
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html)
//
//
// * GetPublicAccessBlock
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html)
//
//
// * PutPublicAccessBlock
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutPublicAccessBlock.html)
//
//
// * DeletePublicAccessBlock
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html)
func (c *Client) GetBucketPolicyStatus(ctx context.Context, params *GetBucketPolicyStatusInput, optFns ...func(*Options)) (*GetBucketPolicyStatusOutput, error) {
	if params == nil {
		params = &GetBucketPolicyStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketPolicyStatus", params, optFns, addOperationGetBucketPolicyStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketPolicyStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketPolicyStatusInput struct {

	// The name of the Amazon S3 bucket whose policy status you want to retrieve.
	//
	// This member is required.
	Bucket *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type GetBucketPolicyStatusOutput struct {

	// The policy status for the specified bucket.
	PolicyStatus *types.PolicyStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketPolicyStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketPolicyStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketPolicyStatus{}, middleware.After)
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
	addOpGetBucketPolicyStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketPolicyStatus(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBucketPolicyStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketPolicyStatus",
	}
}
