// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a job to import or export data between Amazon S3 and your on-premises
// data center. Your AWS account must have the right trust policies and permissions
// in place to create a job for a Snow device. If you're creating a job for a node
// in a cluster, you only need to provide the clusterId value; the other job
// attributes are inherited from the cluster.
func (c *Client) CreateJob(ctx context.Context, params *CreateJobInput, optFns ...func(*Options)) (*CreateJobOutput, error) {
	if params == nil {
		params = &CreateJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateJob", params, optFns, addOperationCreateJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobInput struct {

	// The ID for the address that you want the Snow device shipped to.
	AddressId *string

	// The ID of a cluster. If you're creating a job for a node in a cluster, you need
	// to provide only this clusterId value. The other job attributes are inherited
	// from the cluster.
	ClusterId *string

	// Defines an optional description of this specific job, for example Important
	// Photos 2016-08-11.
	Description *string

	// Defines the device configuration for an AWS Snowcone job.
	DeviceConfiguration *types.DeviceConfiguration

	// The forwarding address ID for a job. This field is not supported in most
	// regions.
	ForwardingAddressId *string

	// Defines the type of job that you're creating.
	JobType types.JobType

	// The KmsKeyARN that you want to associate with this job. KmsKeyARNs are created
	// using the CreateKey
	// (https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html) AWS Key
	// Management Service (KMS) API action.
	KmsKeyARN *string

	// Defines the Amazon Simple Notification Service (Amazon SNS) notification
	// settings for this job.
	Notification *types.Notification

	// Defines the Amazon S3 buckets associated with this job. With IMPORT jobs, you
	// specify the bucket or buckets that your transferred data will be imported into.
	// With EXPORT jobs, you specify the bucket or buckets that your transferred data
	// will be exported from. Optionally, you can also specify a KeyRange value. If you
	// choose to export a range, you define the length of the range by providing either
	// an inclusive BeginMarker value, an inclusive EndMarker value, or both. Ranges
	// are UTF-8 binary sorted.
	Resources *types.JobResource

	// The RoleARN that you want to associate with this job. RoleArns are created using
	// the CreateRole
	// (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html) AWS
	// Identity and Access Management (IAM) API action.
	RoleARN *string

	// The shipping speed for this job. This speed doesn't dictate how soon you'll get
	// the Snow device, rather it represents how quickly the Snow device moves to its
	// destination while in transit. Regional shipping speeds are as follows:
	//
	//     * In
	// Australia, you have access to express shipping. Typically, Snow devices shipped
	// express are delivered in about a day.
	//
	//     * In the European Union (EU), you
	// have access to express shipping. Typically, Snow devices shipped express are
	// delivered in about a day. In addition, most countries in the EU have access to
	// standard shipping, which typically takes less than a week, one way.
	//
	//     * In
	// India, Snow devices are delivered in one to seven days.
	//
	//     * In the US, you
	// have access to one-day shipping and two-day shipping.
	ShippingOption types.ShippingOption

	// If your job is being created in one of the US regions, you have the option of
	// specifying what size Snow device you'd like for this job. In all other regions,
	// Snowballs come with 80 TB in storage capacity.
	SnowballCapacityPreference types.SnowballCapacity

	// The type of AWS Snow Family device to use for this job. For cluster jobs, AWS
	// Snow Family currently supports only the EDGE device type. The type of AWS Snow
	// device to use for this job. Currently, the only supported device type for
	// cluster jobs is EDGE. For more information, see Snowball Edge Device Options
	// (https://docs.aws.amazon.com/snowball/latest/developer-guide/device-differences.html)
	// in the Snowball Edge Developer Guide.
	SnowballType types.SnowballType

	// The tax documents required in your AWS Region.
	TaxDocuments *types.TaxDocuments
}

type CreateJobOutput struct {

	// The automatically generated ID for a job, for example
	// JID123e4567-e89b-12d3-a456-426655440000.
	JobId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateJob{}, middleware.After)
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
	addOpCreateJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "CreateJob",
	}
}
