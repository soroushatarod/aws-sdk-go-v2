// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates an account with an Amazon Macie master account.
func (c *Client) CreateMember(ctx context.Context, params *CreateMemberInput, optFns ...func(*Options)) (*CreateMemberOutput, error) {
	if params == nil {
		params = &CreateMemberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateMember", params, optFns, addOperationCreateMemberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateMemberInput struct {

	// The details for the account to associate with the master account.
	//
	// This member is required.
	Account *types.AccountDetail

	// A map of key-value pairs that specifies the tags to associate with the account
	// in Amazon Macie. An account can have a maximum of 50 tags. Each tag consists of
	// a tag key and an associated tag value. The maximum length of a tag key is 128
	// characters. The maximum length of a tag value is 256 characters.
	Tags map[string]*string
}

type CreateMemberOutput struct {

	// The Amazon Resource Name (ARN) of the account that was associated with the
	// master account.
	Arn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateMemberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateMember{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateMember{}, middleware.After)
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
	addOpCreateMemberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateMember(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateMember(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "CreateMember",
	}
}
