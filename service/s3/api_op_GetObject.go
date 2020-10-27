// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"time"
)

// Retrieves objects from Amazon S3. To use GET, you must have READ access to the
// object. If you grant READ access to the anonymous user, you can return the
// object without using an authorization header. An Amazon S3 bucket has no
// directory hierarchy such as you would find in a typical computer file system.
// You can, however, create a logical hierarchy by using object key names that
// imply a folder structure. For example, instead of naming an object sample.jpg,
// you can name it photos/2006/February/sample.jpg. To get an object from such a
// logical hierarchy, specify the full key name for the object in the GET
// operation. For a virtual hosted-style request example, if you have the object
// photos/2006/February/sample.jpg, specify the resource as
// /photos/2006/February/sample.jpg. For a path-style request example, if you have
// the object photos/2006/February/sample.jpg in the bucket named examplebucket,
// specify the resource as /examplebucket/photos/2006/February/sample.jpg. For more
// information about request types, see HTTP Host Header Bucket Specification
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html#VirtualHostingSpecifyBucket).
// To distribute large files to many people, you can save bandwidth costs by using
// BitTorrent. For more information, see Amazon S3 Torrent
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3Torrent.html). For more
// information about returning the ACL of an object, see GetObjectAcl
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html). If the
// object you are retrieving is stored in the GLACIER or DEEP_ARCHIVE storage
// classes, before you can retrieve the object you must first restore a copy using
// RestoreObject
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html).
// Otherwise, this operation returns an InvalidObjectStateError error. For
// information about restoring archived objects, see Restoring Archived Objects
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/restoring-objects.html).
// Encryption request headers, like x-amz-server-side-encryption, should not be
// sent for GET requests if your object uses server-side encryption with CMKs
// stored in AWS KMS (SSE-KMS) or server-side encryption with Amazon S3–managed
// encryption keys (SSE-S3). If your object does use these types of keys, you’ll
// get an HTTP 400 BadRequest error. If you encrypt an object by using server-side
// encryption with customer-provided encryption keys (SSE-C) when you store the
// object in Amazon S3, then when you GET the object, you must use the following
// headers:
//
//     * x-amz-server-side-encryption-customer-algorithm
//
//     *
// x-amz-server-side-encryption-customer-key
//
//     *
// x-amz-server-side-encryption-customer-key-MD5
//
// For more information about SSE-C,
// see Server-Side Encryption (Using Customer-Provided Encryption Keys)
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html).
// Assuming you have permission to read object tags (permission for the
// s3:GetObjectVersionTagging action), the response also returns the
// x-amz-tagging-count header that provides the count of number of tags associated
// with the object. You can use GetObjectTagging
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html) to
// retrieve the tag set associated with an object. Permissions You need the
// s3:GetObject permission for this operation. For more information, see Specifying
// Permissions in a Policy
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html). If
// the object you request does not exist, the error Amazon S3 returns depends on
// whether you also have the s3:ListBucket permission.
//
//     * If you have the
// s3:ListBucket permission on the bucket, Amazon S3 will return an HTTP status
// code 404 ("no such key") error.
//
//     * If you don’t have the s3:ListBucket
// permission, Amazon S3 will return an HTTP status code 403 ("access denied")
// error.
//
// Versioning By default, the GET operation returns the current version of
// an object. To return a different version, use the versionId subresource. If the
// current version of the object is a delete marker, Amazon S3 behaves as if the
// object was deleted and includes x-amz-delete-marker: true in the response. For
// more information about versioning, see PutBucketVersioning
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketVersioning.html).
// Overriding Response Header Values There are times when you want to override
// certain response header values in a GET response. For example, you might
// override the Content-Disposition response header value in your GET request. You
// can override values for a set of response headers using the following query
// parameters. These response header values are sent only on a successful request,
// that is, when status code 200 OK is returned. The set of headers you can
// override using these parameters is a subset of the headers that Amazon S3
// accepts when you create an object. The response headers that you can override
// for the GET response are Content-Type, Content-Language, Expires, Cache-Control,
// Content-Disposition, and Content-Encoding. To override these header values in
// the GET response, you use the following request parameters. You must sign the
// request, either using an Authorization header or a presigned URL, when using
// these parameters. They cannot be used with an unsigned (anonymous) request.
//
//
// * response-content-type
//
//     * response-content-language
//
//     *
// response-expires
//
//     * response-cache-control
//
//     *
// response-content-disposition
//
//     * response-content-encoding
//
// Additional
// Considerations about Request Headers If both of the If-Match and
// If-Unmodified-Since headers are present in the request as follows: If-Match
// condition evaluates to true, and; If-Unmodified-Since condition evaluates to
// false; then, S3 returns 200 OK and the data requested. If both of the
// If-None-Match and If-Modified-Since headers are present in the request as
// follows: If-None-Match condition evaluates to false, and; If-Modified-Since
// condition evaluates to true; then, S3 returns 304 Not Modified response code.
// For more information about conditional requests, see RFC 7232
// (https://tools.ietf.org/html/rfc7232). The following operations are related to
// GetObject:
//
//     * ListBuckets
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html)
//
//     *
// GetObjectAcl
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html)
func (c *Client) GetObject(ctx context.Context, params *GetObjectInput, optFns ...func(*Options)) (*GetObjectOutput, error) {
	if params == nil {
		params = &GetObjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetObject", params, optFns, addOperationGetObjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetObjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetObjectInput struct {

	// The bucket name containing the object. When using this API with an access point,
	// you must direct requests to the access point hostname. The access point hostname
	// takes the form AccessPointName-AccountId.s3-accesspoint.Region.amazonaws.com.
	// When using this operation with an access point through the AWS SDKs, you provide
	// the access point ARN in place of the bucket name. For more information about
	// access point ARNs, see Using Access Points
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-access-points.html) in
	// the Amazon Simple Storage Service Developer Guide. When using this API with
	// Amazon S3 on Outposts, you must direct requests to the S3 on Outposts hostname.
	// The S3 on Outposts hostname takes the form
	// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com. When using
	// this operation using S3 on Outposts through the AWS SDKs, you provide the
	// Outposts bucket ARN in place of the bucket name. For more information about S3
	// on Outposts ARNs, see Using S3 on Outposts
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/S3onOutposts.html) in the
	// Amazon Simple Storage Service Developer Guide.
	//
	// This member is required.
	Bucket *string

	// Key of the object to get.
	//
	// This member is required.
	Key *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string

	// Return the object only if its entity tag (ETag) is the same as the one
	// specified, otherwise return a 412 (precondition failed).
	IfMatch *string

	// Return the object only if it has been modified since the specified time,
	// otherwise return a 304 (not modified).
	IfModifiedSince *time.Time

	// Return the object only if its entity tag (ETag) is different from the one
	// specified, otherwise return a 304 (not modified).
	IfNoneMatch *string

	// Return the object only if it has not been modified since the specified time,
	// otherwise return a 412 (precondition failed).
	IfUnmodifiedSince *time.Time

	// Part number of the object being read. This is a positive integer between 1 and
	// 10,000. Effectively performs a 'ranged' GET request for the part specified.
	// Useful for downloading just a part of an object.
	PartNumber *int32

	// Downloads the specified range bytes of an object. For more information about the
	// HTTP Range header, see
	// https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35
	// (https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.35). Amazon S3
	// doesn't support retrieving multiple ranges of data per GET request.
	Range *string

	// Confirms that the requester knows that they will be charged for the request.
	// Bucket owners need not specify this parameter in their requests. For information
	// about downloading objects from requester pays buckets, see Downloading Objects
	// in Requestor Pays Buckets
	// (https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectsinRequesterPaysBuckets.html)
	// in the Amazon S3 Developer Guide.
	RequestPayer types.RequestPayer

	// Sets the Cache-Control header of the response.
	ResponseCacheControl *string

	// Sets the Content-Disposition header of the response
	ResponseContentDisposition *string

	// Sets the Content-Encoding header of the response.
	ResponseContentEncoding *string

	// Sets the Content-Language header of the response.
	ResponseContentLanguage *string

	// Sets the Content-Type header of the response.
	ResponseContentType *string

	// Sets the Expires header of the response.
	ResponseExpires *time.Time

	// Specifies the algorithm to use to when encrypting the object (for example,
	// AES256).
	SSECustomerAlgorithm *string

	// Specifies the customer-provided encryption key for Amazon S3 to use in
	// encrypting data. This value is used to store the object and then it is
	// discarded; Amazon S3 does not store the encryption key. The key must be
	// appropriate for use with the algorithm specified in the
	// x-amz-server-side-encryption-customer-algorithm header.
	SSECustomerKey *string

	// Specifies the 128-bit MD5 digest of the encryption key according to RFC 1321.
	// Amazon S3 uses this header for a message integrity check to ensure that the
	// encryption key was transmitted without error.
	SSECustomerKeyMD5 *string

	// VersionId used to reference a specific version of the object.
	VersionId *string
}

type GetObjectOutput struct {

	// Indicates that a range of bytes was specified.
	AcceptRanges *string

	// Object data.
	Body io.ReadCloser

	// Specifies caching behavior along the request/reply chain.
	CacheControl *string

	// Specifies presentational information for the object.
	ContentDisposition *string

	// Specifies what content encodings have been applied to the object and thus what
	// decoding mechanisms must be applied to obtain the media-type referenced by the
	// Content-Type header field.
	ContentEncoding *string

	// The language the content is in.
	ContentLanguage *string

	// Size of the body in bytes.
	ContentLength *int64

	// The portion of the object returned in the response.
	ContentRange *string

	// A standard MIME type describing the format of the object data.
	ContentType *string

	// Specifies whether the object retrieved was (true) or was not (false) a Delete
	// Marker. If false, this response header does not appear in the response.
	DeleteMarker *bool

	// An ETag is an opaque identifier assigned by a web server to a specific version
	// of a resource found at a URL.
	ETag *string

	// If the object expiration is configured (see PUT Bucket lifecycle), the response
	// includes this header. It includes the expiry-date and rule-id key-value pairs
	// providing object expiration information. The value of the rule-id is URL
	// encoded.
	Expiration *string

	// The date and time at which the object is no longer cacheable.
	Expires *time.Time

	// Last modified date of the object
	LastModified *time.Time

	// A map of metadata to store with the object in S3.
	Metadata map[string]*string

	// This is set to the number of metadata entries not returned in x-amz-meta
	// headers. This can happen if you create metadata using an API like SOAP that
	// supports more flexible metadata than the REST API. For example, using SOAP, you
	// can create metadata whose values are not legal HTTP headers.
	MissingMeta *int32

	// Indicates whether this object has an active legal hold. This field is only
	// returned if you have permission to view an object's legal hold status.
	ObjectLockLegalHoldStatus types.ObjectLockLegalHoldStatus

	// The Object Lock mode currently in place for this object.
	ObjectLockMode types.ObjectLockMode

	// The date and time when this object's Object Lock will expire.
	ObjectLockRetainUntilDate *time.Time

	// The count of parts this object has.
	PartsCount *int32

	// Amazon S3 can return this if your request involves a bucket that is either a
	// source or destination in a replication rule.
	ReplicationStatus types.ReplicationStatus

	// If present, indicates that the requester was successfully charged for the
	// request.
	RequestCharged types.RequestCharged

	// Provides information about object restoration operation and expiration time of
	// the restored object copy.
	Restore *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header confirming the encryption algorithm used.
	SSECustomerAlgorithm *string

	// If server-side encryption with a customer-provided encryption key was requested,
	// the response will include this header to provide round-trip message integrity
	// verification of the customer-provided encryption key.
	SSECustomerKeyMD5 *string

	// If present, specifies the ID of the AWS Key Management Service (AWS KMS)
	// symmetric customer managed customer master key (CMK) that was used for the
	// object.
	SSEKMSKeyId *string

	// The server-side encryption algorithm used when storing this object in Amazon S3
	// (for example, AES256, aws:kms).
	ServerSideEncryption types.ServerSideEncryption

	// Provides storage class information of the object. Amazon S3 returns this header
	// for all objects except for S3 Standard storage class objects.
	StorageClass types.StorageClass

	// The number of tags, if any, on the object.
	TagCount *int32

	// Version of the object.
	VersionId *string

	// If the bucket is configured as a website, redirects requests for this object to
	// another object in the same bucket or to an external URL. Amazon S3 stores the
	// value of this header in the object metadata.
	WebsiteRedirectLocation *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetObjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetObject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetObject{}, middleware.After)
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
	addOpGetObjectValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetObject(options.Region), middleware.Before)
	addMetadataRetrieverMiddleware(stack)
	addUpdateEndpointMiddleware(stack, options)
	addResponseErrorMiddleware(stack)
	v4.AddContentSHA256HeaderMiddleware(stack)
	disableAcceptEncodingGzip(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetObject(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetObject",
	}
}
