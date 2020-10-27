// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new Amazon WorkMail organization. Optionally, you can choose to
// associate an existing AWS Directory Service directory with your organization. If
// an AWS Directory Service directory ID is specified, the organization alias must
// match the directory alias. If you choose not to associate an existing directory
// with your organization, then we create a new Amazon WorkMail directory for you.
// For more information, see Adding an organization
// (https://docs.aws.amazon.com/workmail/latest/adminguide/add_new_organization.html)
// in the Amazon WorkMail Administrator Guide. You can associate multiple email
// domains with an organization, then set your default email domain from the Amazon
// WorkMail console. You can also associate a domain that is managed in an Amazon
// Route 53 public hosted zone. For more information, see Adding a domain
// (https://docs.aws.amazon.com/workmail/latest/adminguide/add_domain.html) and
// Choosing the default domain
// (https://docs.aws.amazon.com/workmail/latest/adminguide/default_domain.html) in
// the Amazon WorkMail Administrator Guide. Optionally, you can use a customer
// managed master key from AWS Key Management Service (AWS KMS) to encrypt email
// for your organization. If you don't associate an AWS KMS key, Amazon WorkMail
// creates a default AWS managed master key for you.
func (c *Client) CreateOrganization(ctx context.Context, params *CreateOrganizationInput, optFns ...func(*Options)) (*CreateOrganizationOutput, error) {
	if params == nil {
		params = &CreateOrganizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateOrganization", params, optFns, addOperationCreateOrganizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateOrganizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateOrganizationInput struct {

	// The organization alias.
	//
	// This member is required.
	Alias *string

	// The idempotency token associated with the request.
	ClientToken *string

	// The AWS Directory Service directory ID.
	DirectoryId *string

	// The email domains to associate with the organization.
	Domains []*types.Domain

	// When true, allows organization interoperability between Amazon WorkMail and
	// Microsoft Exchange. Can only be set to true if an AD Connector directory ID is
	// included in the request.
	EnableInteroperability *bool

	// The Amazon Resource Name (ARN) of a customer managed master key from AWS KMS.
	KmsKeyArn *string
}

type CreateOrganizationOutput struct {

	// The organization ID.
	OrganizationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateOrganizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateOrganization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateOrganization{}, middleware.After)
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
	addIdempotencyToken_opCreateOrganizationMiddleware(stack, options)
	addOpCreateOrganizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateOrganization(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateOrganization struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateOrganization) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateOrganization) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateOrganizationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateOrganizationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateOrganizationMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateOrganization{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateOrganization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "CreateOrganization",
	}
}
