// Code generated by smithy-go-codegen DO NOT EDIT.

package ivs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivs/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a specified playback authorization key pair and returns the arn and
// fingerprint. The privateKey held by the caller can be used to generate viewer
// authorization tokens, to grant viewers access to authorized channels.
func (c *Client) GetPlaybackKeyPair(ctx context.Context, params *GetPlaybackKeyPairInput, optFns ...func(*Options)) (*GetPlaybackKeyPairOutput, error) {
	if params == nil {
		params = &GetPlaybackKeyPairInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPlaybackKeyPair", params, optFns, addOperationGetPlaybackKeyPairMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPlaybackKeyPairOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPlaybackKeyPairInput struct {

	// ARN of the key pair to be returned.
	//
	// This member is required.
	Arn *string
}

type GetPlaybackKeyPairOutput struct {

	// A key pair used to sign and validate a playback authorization token.
	KeyPair *types.PlaybackKeyPair

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPlaybackKeyPairMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetPlaybackKeyPair{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetPlaybackKeyPair{}, middleware.After)
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
	addOpGetPlaybackKeyPairValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetPlaybackKeyPair(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetPlaybackKeyPair(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivs",
		OperationName: "GetPlaybackKeyPair",
	}
}
