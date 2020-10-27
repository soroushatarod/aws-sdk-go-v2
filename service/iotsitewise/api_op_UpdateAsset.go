// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an asset's name. For more information, see Updating assets and models
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/update-assets-and-models.html)
// in the AWS IoT SiteWise User Guide.
func (c *Client) UpdateAsset(ctx context.Context, params *UpdateAssetInput, optFns ...func(*Options)) (*UpdateAssetOutput, error) {
	if params == nil {
		params = &UpdateAssetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateAsset", params, optFns, addOperationUpdateAssetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateAssetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateAssetInput struct {

	// The ID of the asset to update.
	//
	// This member is required.
	AssetId *string

	// A unique, friendly name for the asset.
	//
	// This member is required.
	AssetName *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string
}

type UpdateAssetOutput struct {

	// The status of the asset, which contains a state (UPDATING after successfully
	// calling this operation) and any error message.
	//
	// This member is required.
	AssetStatus *types.AssetStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateAssetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateAsset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateAsset{}, middleware.After)
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
	addIdempotencyToken_opUpdateAssetMiddleware(stack, options)
	addOpUpdateAssetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateAsset(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpUpdateAsset struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateAsset) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateAsset) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateAssetInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateAssetInput ")
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
func addIdempotencyToken_opUpdateAssetMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpUpdateAsset{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateAsset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "UpdateAsset",
	}
}
