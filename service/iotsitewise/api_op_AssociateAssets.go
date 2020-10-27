// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Associates a child asset with the given parent asset through a hierarchy defined
// in the parent asset's model. For more information, see Associating assets
// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/add-associated-assets.html)
// in the AWS IoT SiteWise User Guide.
func (c *Client) AssociateAssets(ctx context.Context, params *AssociateAssetsInput, optFns ...func(*Options)) (*AssociateAssetsOutput, error) {
	if params == nil {
		params = &AssociateAssetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateAssets", params, optFns, addOperationAssociateAssetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateAssetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateAssetsInput struct {

	// The ID of the parent asset.
	//
	// This member is required.
	AssetId *string

	// The ID of the child asset to be associated.
	//
	// This member is required.
	ChildAssetId *string

	// The ID of a hierarchy in the parent asset's model. Hierarchies allow different
	// groupings of assets to be formed that all come from the same asset model. For
	// more information, see Asset hierarchies
	// (https://docs.aws.amazon.com/iot-sitewise/latest/userguide/asset-hierarchies.html)
	// in the AWS IoT SiteWise User Guide.
	//
	// This member is required.
	HierarchyId *string

	// A unique case-sensitive identifier that you can provide to ensure the
	// idempotency of the request. Don't reuse this client token if a new idempotent
	// request is required.
	ClientToken *string
}

type AssociateAssetsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationAssociateAssetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateAssets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateAssets{}, middleware.After)
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
	addIdempotencyToken_opAssociateAssetsMiddleware(stack, options)
	addOpAssociateAssetsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateAssets(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type idempotencyToken_initializeOpAssociateAssets struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpAssociateAssets) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpAssociateAssets) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*AssociateAssetsInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *AssociateAssetsInput ")
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
func addIdempotencyToken_opAssociateAssetsMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpAssociateAssets{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opAssociateAssets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "AssociateAssets",
	}
}
