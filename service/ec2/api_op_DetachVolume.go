// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Detaches an EBS volume from an instance. Make sure to unmount any file systems
// on the device within your operating system before detaching the volume. Failure
// to do so can result in the volume becoming stuck in the busy state while
// detaching. If this happens, detachment can be delayed indefinitely until you
// unmount the volume, force detachment, reboot the instance, or all three. If an
// EBS volume is the root device of an instance, it can't be detached while the
// instance is running. To detach the root volume, stop the instance first. When a
// volume with an AWS Marketplace product code is detached from an instance, the
// product code is no longer associated with the instance. For more information,
// see Detaching an Amazon EBS volume
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-detaching-volume.html)
// in the Amazon Elastic Compute Cloud User Guide.
func (c *Client) DetachVolume(ctx context.Context, params *DetachVolumeInput, optFns ...func(*Options)) (*DetachVolumeOutput, error) {
	if params == nil {
		params = &DetachVolumeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetachVolume", params, optFns, addOperationDetachVolumeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetachVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetachVolumeInput struct {

	// The ID of the volume.
	//
	// This member is required.
	VolumeId *string

	// The device name.
	Device *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// Forces detachment if the previous detachment attempt did not occur cleanly (for
	// example, logging into an instance, unmounting the volume, and detaching
	// normally). This option can lead to data loss or a corrupted file system. Use
	// this option only as a last resort to detach a volume from a failed instance. The
	// instance won't have an opportunity to flush file system caches or file system
	// metadata. If you use this option, you must perform file system check and repair
	// procedures.
	Force *bool

	// The ID of the instance. If you are detaching a Multi-Attach enabled volume, you
	// must specify an instance ID.
	InstanceId *string
}

// Describes volume attachment details.
type DetachVolumeOutput struct {

	// The time stamp when the attachment initiated.
	AttachTime *time.Time

	// Indicates whether the EBS volume is deleted on instance termination.
	DeleteOnTermination *bool

	// The device name.
	Device *string

	// The ID of the instance.
	InstanceId *string

	// The attachment state of the volume.
	State types.VolumeAttachmentState

	// The ID of the volume.
	VolumeId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDetachVolumeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDetachVolume{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDetachVolume{}, middleware.After)
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
	addOpDetachVolumeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachVolume(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDetachVolume(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DetachVolume",
	}
}
