// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of IP address ranges, specified as IPv4 CIDR blocks, that you
// can use for the network management interface when you enable Bring Your Own
// License (BYOL). This operation can be run only by AWS accounts that are enabled
// for BYOL. If your account isn't enabled for BYOL, you'll receive an
// AccessDeniedException error. The management network interface is connected to a
// secure Amazon WorkSpaces management network. It is used for interactive
// streaming of the WorkSpace desktop to Amazon WorkSpaces clients, and to allow
// Amazon WorkSpaces to manage the WorkSpace.
func (c *Client) ListAvailableManagementCidrRanges(ctx context.Context, params *ListAvailableManagementCidrRangesInput, optFns ...func(*Options)) (*ListAvailableManagementCidrRangesOutput, error) {
	if params == nil {
		params = &ListAvailableManagementCidrRangesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListAvailableManagementCidrRanges", params, optFns, addOperationListAvailableManagementCidrRangesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListAvailableManagementCidrRangesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListAvailableManagementCidrRangesInput struct {

	// The IP address range to search. Specify an IP address range that is compatible
	// with your network and in CIDR notation (that is, specify the range as an IPv4
	// CIDR block).
	//
	// This member is required.
	ManagementCidrRangeConstraint *string

	// The maximum number of items to return.
	MaxResults *int32

	// If you received a NextToken from a previous call that was paginated, provide
	// this token to receive the next set of results.
	NextToken *string
}

type ListAvailableManagementCidrRangesOutput struct {

	// The list of available IP address ranges, specified as IPv4 CIDR blocks.
	ManagementCidrRanges []*string

	// The token to use to retrieve the next set of results, or null if no more results
	// are available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListAvailableManagementCidrRangesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListAvailableManagementCidrRanges{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListAvailableManagementCidrRanges{}, middleware.After)
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
	addOpListAvailableManagementCidrRangesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListAvailableManagementCidrRanges(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListAvailableManagementCidrRanges(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ListAvailableManagementCidrRanges",
	}
}
