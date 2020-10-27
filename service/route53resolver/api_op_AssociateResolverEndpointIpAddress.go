// Code generated by smithy-go-codegen DO NOT EDIT.

package route53resolver

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53resolver/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds IP addresses to an inbound or an outbound Resolver endpoint. If you want to
// add more than one IP address, submit one AssociateResolverEndpointIpAddress
// request for each IP address. To remove an IP address from an endpoint, see
// DisassociateResolverEndpointIpAddress
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_DisassociateResolverEndpointIpAddress.html).
func (c *Client) AssociateResolverEndpointIpAddress(ctx context.Context, params *AssociateResolverEndpointIpAddressInput, optFns ...func(*Options)) (*AssociateResolverEndpointIpAddressOutput, error) {
	if params == nil {
		params = &AssociateResolverEndpointIpAddressInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateResolverEndpointIpAddress", params, optFns, addOperationAssociateResolverEndpointIpAddressMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateResolverEndpointIpAddressOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResolverEndpointIpAddressInput struct {

	// Either the IPv4 address that you want to add to a Resolver endpoint or a subnet
	// ID. If you specify a subnet ID, Resolver chooses an IP address for you from the
	// available IPs in the specified subnet.
	//
	// This member is required.
	IpAddress *types.IpAddressUpdate

	// The ID of the Resolver endpoint that you want to associate IP addresses with.
	//
	// This member is required.
	ResolverEndpointId *string
}

type AssociateResolverEndpointIpAddressOutput struct {

	// The response to an AssociateResolverEndpointIpAddress request.
	ResolverEndpoint *types.ResolverEndpoint

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateResolverEndpointIpAddressMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateResolverEndpointIpAddress{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateResolverEndpointIpAddress{}, middleware.After)
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
	addOpAssociateResolverEndpointIpAddressValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResolverEndpointIpAddress(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opAssociateResolverEndpointIpAddress(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53resolver",
		OperationName: "AssociateResolverEndpointIpAddress",
	}
}
