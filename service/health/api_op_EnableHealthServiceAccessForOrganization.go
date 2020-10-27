// Code generated by smithy-go-codegen DO NOT EDIT.

package health

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Calling this operation enables AWS Health to work with AWS Organizations. This
// applies a service-linked role (SLR) to the master account in the organization.
// To call this operation, you must sign in as an IAM user, assume an IAM role, or
// sign in as the root user (not recommended) in the organization's master account.
// For more information, see Aggregating AWS Health events
// (https://docs.aws.amazon.com/health/latest/ug/aggregate-events.html) in the AWS
// Health User Guide.
func (c *Client) EnableHealthServiceAccessForOrganization(ctx context.Context, params *EnableHealthServiceAccessForOrganizationInput, optFns ...func(*Options)) (*EnableHealthServiceAccessForOrganizationOutput, error) {
	if params == nil {
		params = &EnableHealthServiceAccessForOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableHealthServiceAccessForOrganization", params, optFns, addOperationEnableHealthServiceAccessForOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableHealthServiceAccessForOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableHealthServiceAccessForOrganizationInput struct {
}

type EnableHealthServiceAccessForOrganizationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationEnableHealthServiceAccessForOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpEnableHealthServiceAccessForOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableHealthServiceAccessForOrganization{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableHealthServiceAccessForOrganization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opEnableHealthServiceAccessForOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "health",
		OperationName: "EnableHealthServiceAccessForOrganization",
	}
}
