// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Releases the specified Elastic IP address. [EC2-Classic, default VPC] Releasing
// an Elastic IP address automatically disassociates it from any instance that it's
// associated with. To disassociate an Elastic IP address without releasing it, use
// DisassociateAddress. [Nondefault VPC] You must use DisassociateAddress to
// disassociate the Elastic IP address before you can release it. Otherwise, Amazon
// EC2 returns an error (InvalidIPAddress.InUse). After releasing an Elastic IP
// address, it is released to the IP address pool. Be sure to update your DNS
// records and any servers or devices that communicate with the address. If you
// attempt to release an Elastic IP address that you already released, you'll get
// an AuthFailure error if the address is already allocated to another AWS account.
// [EC2-VPC] After you release an Elastic IP address for use in a VPC, you might be
// able to recover it. For more information, see AllocateAddress.
func (c *Client) ReleaseAddress(ctx context.Context, params *ReleaseAddressInput, optFns ...func(*Options)) (*ReleaseAddressOutput, error) {
	if params == nil {
		params = &ReleaseAddressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ReleaseAddress", params, optFns, addOperationReleaseAddressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ReleaseAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ReleaseAddressInput struct {

	// [EC2-VPC] The allocation ID. Required for EC2-VPC.
	AllocationId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The set of Availability Zones, Local Zones, or Wavelength Zones from which AWS
	// advertises IP addresses. If you provide an incorrect network border group, you
	// will receive an InvalidAddress.NotFound error. For more information, see Error
	// Codes
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html).
	// You cannot use a network border group with EC2 Classic. If you attempt this
	// operation on EC2 classic, you will receive an InvalidParameterCombination error.
	// For more information, see Error Codes
	// (https://docs.aws.amazon.com/AWSEC2/latest/APIReference/errors-overview.html).
	NetworkBorderGroup *string

	// [EC2-Classic] The Elastic IP address. Required for EC2-Classic.
	PublicIp *string
}

type ReleaseAddressOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationReleaseAddressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpReleaseAddress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpReleaseAddress{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opReleaseAddress(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opReleaseAddress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ReleaseAddress",
	}
}
