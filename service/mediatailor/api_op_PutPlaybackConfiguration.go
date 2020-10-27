// Code generated by smithy-go-codegen DO NOT EDIT.

package mediatailor

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediatailor/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Adds a new playback configuration to AWS Elemental MediaTailor.
func (c *Client) PutPlaybackConfiguration(ctx context.Context, params *PutPlaybackConfigurationInput, optFns ...func(*Options)) (*PutPlaybackConfigurationOutput, error) {
	if params == nil {
		params = &PutPlaybackConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutPlaybackConfiguration", params, optFns, addOperationPutPlaybackConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutPlaybackConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPlaybackConfigurationInput struct {

	// The URL for the ad decision server (ADS). This includes the specification of
	// static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing you can provide a static
	// VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string

	// The configuration for Avail Suppression. Ad suppression can be used to turn off
	// ad personalization in a long manifest, or if a viewer joins mid-break.
	AvailSuppression *types.AvailSuppression

	// The configuration for bumpers. Bumpers are short audio or video clips that play
	// at the start or before the end of an ad break.
	Bumper *types.Bumper

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *types.CdnConfiguration

	// The configuration for DASH content.
	DashConfiguration *types.DashConfigurationForPut

	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *types.LivePreRollConfiguration

	// The configuration for manifest processing rules. Manifest processing rules
	// enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules *types.ManifestProcessingRules

	// The identifier for the playback configuration.
	Name *string

	// The maximum duration of underfilled ad time (in seconds) allowed in an ad break.
	PersonalizationThresholdSeconds *int32

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in
	// gaps in media content. Configuring the slate is optional for non-VPAID
	// configurations. For VPAID, the slate is required because MediaTailor provides it
	// in the slots that are designated for dynamic ad content. The slate must be a
	// high-quality asset that contains both audio and video.
	SlateAdUrl *string

	// The tags to assign to the playback configuration.
	Tags map[string]*string

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of
	// MediaTailor. Use this only if you have already set up custom profiles with the
	// help of AWS Support.
	TranscodeProfileName *string

	// The URL prefix for the master playlist for the stream, minus the asset ID. The
	// maximum length is 512 characters.
	VideoContentSourceUrl *string
}

type PutPlaybackConfigurationOutput struct {

	// The URL for the ad decision server (ADS). This includes the specification of
	// static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing, you can provide a static
	// VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string

	// The configuration for Avail Suppression. Ad suppression can be used to turn off
	// ad personalization in a long manifest, or if a viewer joins mid-break.
	AvailSuppression *types.AvailSuppression

	// The configuration for bumpers. Bumpers are short audio or video clips that play
	// at the start or before the end of an ad break.
	Bumper *types.Bumper

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *types.CdnConfiguration

	// The configuration for DASH content.
	DashConfiguration *types.DashConfiguration

	// The configuration for HLS content.
	HlsConfiguration *types.HlsConfiguration

	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *types.LivePreRollConfiguration

	// The configuration for manifest processing rules. Manifest processing rules
	// enable customization of the personalized manifests created by MediaTailor.
	ManifestProcessingRules *types.ManifestProcessingRules

	// The identifier for the playback configuration.
	Name *string

	// The maximum duration of underfilled ad time (in seconds) allowed in an ad break.
	PersonalizationThresholdSeconds *int32

	// The Amazon Resource Name (ARN) for the playback configuration.
	PlaybackConfigurationArn *string

	// The URL that the player accesses to get a manifest from AWS Elemental
	// MediaTailor. This session will use server-side reporting.
	PlaybackEndpointPrefix *string

	// The URL that the player uses to initialize a session that uses client-side
	// reporting.
	SessionInitializationEndpointPrefix *string

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill in
	// gaps in media content. Configuring the slate is optional for non-VPAID playback
	// configurations. For VPAID, the slate is required because MediaTailor provides it
	// in the slots designated for dynamic ad content. The slate must be a high-quality
	// asset that contains both audio and video.
	SlateAdUrl *string

	// The tags assigned to the playback configuration.
	Tags map[string]*string

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of
	// MediaTailor. Use this only if you have already set up custom profiles with the
	// help of AWS Support.
	TranscodeProfileName *string

	// The URL prefix for the master playlist for the stream, minus the asset ID. The
	// maximum length is 512 characters.
	VideoContentSourceUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutPlaybackConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutPlaybackConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutPlaybackConfiguration{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutPlaybackConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutPlaybackConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediatailor",
		OperationName: "PutPlaybackConfiguration",
	}
}
