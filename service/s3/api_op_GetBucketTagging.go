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

// Returns the tag set associated with the bucket. To use this operation, you must
// have permission to perform the s3:GetBucketTagging action. By default, the
// bucket owner has this permission and can grant this permission to others.
// GetBucketTagging has the following special error:
//
//     * Error code:
// NoSuchTagSetError
//
//         * Description: There is no tag set associated with
// the bucket.
//
// The following operations are related to GetBucketTagging:
//
//     *
// PutBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html)
//
//
// * DeleteBucketTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html)
func (c *Client) GetBucketTagging(ctx context.Context, params *GetBucketTaggingInput, optFns ...func(*Options)) (*GetBucketTaggingOutput, error) {
	if params == nil {
		params = &GetBucketTaggingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketTagging", params, optFns, addOperationGetBucketTaggingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketTaggingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketTaggingInput struct {

	// The name of the bucket for which to get the tagging information.
	//
	// This member is required.
	Bucket *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type GetBucketTaggingOutput struct {

	// Contains the tag set.
	//
	// This member is required.
	TagSet []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketTaggingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketTagging{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketTagging{}, middleware.After)
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
	addOpGetBucketTaggingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketTagging(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetBucketTagging(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketTagging",
	}
}
