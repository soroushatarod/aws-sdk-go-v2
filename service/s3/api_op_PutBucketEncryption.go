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

// This implementation of the PUT operation uses the encryption subresource to set
// the default encryption state of an existing bucket. This implementation of the
// PUT operation sets default encryption for a bucket using server-side encryption
// with Amazon S3-managed keys SSE-S3 or AWS KMS customer master keys (CMKs)
// (SSE-KMS). For information about the Amazon S3 default encryption feature, see
// Amazon S3 Default Bucket Encryption
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html). This
// operation requires AWS Signature Version 4. For more information, see
// Authenticating Requests (AWS Signature Version 4). To use this operation, you
// must have permissions to perform the s3:PutEncryptionConfiguration action. The
// bucket owner has this permission by default. The bucket owner can grant this
// permission to others. For more information about permissions, see Permissions
// Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html) in the
// Amazon Simple Storage Service Developer Guide. Related Resources
//
//     *
// GetBucketEncryption
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html)
//
//
// * DeleteBucketEncryption
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketEncryption.html)
func (c *Client) PutBucketEncryption(ctx context.Context, params *PutBucketEncryptionInput, optFns ...func(*Options)) (*PutBucketEncryptionOutput, error) {
	if params == nil {
		params = &PutBucketEncryptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutBucketEncryption", params, optFns, addOperationPutBucketEncryptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutBucketEncryptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutBucketEncryptionInput struct {

	// Specifies default encryption for a bucket using server-side encryption with
	// Amazon S3-managed keys (SSE-S3) or customer master keys stored in AWS KMS
	// (SSE-KMS). For information about the Amazon S3 default encryption feature, see
	// Amazon S3 Default Bucket Encryption
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html) in the
	// Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Specifies the default server-side-encryption configuration.
	//
	// This member is required.
	ServerSideEncryptionConfiguration *types.ServerSideEncryptionConfiguration

	// The base64-encoded 128-bit MD5 digest of the server-side encryption
	// configuration. This parameter is auto-populated when using the command from the
	// CLI.
	ContentMD5 *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type PutBucketEncryptionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutBucketEncryptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpPutBucketEncryption{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpPutBucketEncryption{}, middleware.After)
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
	addOpPutBucketEncryptionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutBucketEncryption(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	smithyhttp.AddChecksumMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutBucketEncryption(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "PutBucketEncryption",
	}
}
