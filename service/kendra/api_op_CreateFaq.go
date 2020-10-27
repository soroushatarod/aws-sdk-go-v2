// Code generated by smithy-go-codegen DO NOT EDIT.

package kendra

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kendra/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an new set of frequently asked question (FAQ) questions and answers.
func (c *Client) CreateFaq(ctx context.Context, params *CreateFaqInput, optFns ...func(*Options)) (*CreateFaqOutput, error) {
	if params == nil {
		params = &CreateFaqInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateFaq", params, optFns, addOperationCreateFaqMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateFaqOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateFaqInput struct {

	// The identifier of the index that contains the FAQ.
	//
	// This member is required.
	IndexId *string

	// The name that should be associated with the FAQ.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of a role with permission to access the S3 bucket
	// that contains the FAQs. For more information, see IAM Roles for Amazon Kendra
	// (https://docs.aws.amazon.com/kendra/latest/dg/iam-roles.html).
	//
	// This member is required.
	RoleArn *string

	// The S3 location of the FAQ input data.
	//
	// This member is required.
	S3Path *types.S3Path

	// A token that you provide to identify the request to create a FAQ. Multiple calls
	// to the CreateFaqRequest operation with the same client token will create only
	// one FAQ.
	ClientToken *string

	// A description of the FAQ.
	Description *string

	// The format of the input file. You can choose between a basic CSV format, a CSV
	// format that includes customs attributes in a header, and a JSON format that
	// includes custom attributes. The format must match the format of the file stored
	// in the S3 bucket identified in the S3Path parameter. For more information, see
	// Adding questions and answers
	// (https://docs.aws.amazon.com/kendra/latest/dg/in-creating-faq.html).
	FileFormat types.FaqFileFormat

	// A list of key-value pairs that identify the FAQ. You can use the tags to
	// identify and organize your resources and to control access to resources.
	Tags []*types.Tag
}

type CreateFaqOutput struct {

	// The unique identifier of the FAQ.
	Id *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateFaqMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateFaq{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateFaq{}, middleware.After)
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
	addIdempotencyToken_opCreateFaqMiddleware(stack, options)
	addOpCreateFaqValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFaq(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpCreateFaq struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateFaq) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateFaq) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateFaqInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateFaqInput ")
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
func addIdempotencyToken_opCreateFaqMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpCreateFaq{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateFaq(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kendra",
		OperationName: "CreateFaq",
	}
}
