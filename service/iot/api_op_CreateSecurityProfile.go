// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a Device Defender security profile.
func (c *Client) CreateSecurityProfile(ctx context.Context, params *CreateSecurityProfileInput, optFns ...func(*Options)) (*CreateSecurityProfileOutput, error) {
	if params == nil {
		params = &CreateSecurityProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateSecurityProfile", params, optFns, addOperationCreateSecurityProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateSecurityProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateSecurityProfileInput struct {

	// The name you are giving to the security profile.
	//
	// This member is required.
	SecurityProfileName *string

	// Please use CreateSecurityProfileRequest$additionalMetricsToRetainV2 instead. A
	// list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here.
	AdditionalMetricsToRetain []*string

	// A list of metrics whose data is retained (stored). By default, data is retained
	// for any metric used in the profile's behaviors, but it is also retained for any
	// metric specified here.
	AdditionalMetricsToRetainV2 []*types.MetricToRetain

	// Specifies the destinations to which alerts are sent. (Alerts are always sent to
	// the console.) Alerts are generated when a device (thing) violates a behavior.
	AlertTargets map[string]*types.AlertTarget

	// Specifies the behaviors that, when violated by a device (thing), cause an alert.
	Behaviors []*types.Behavior

	// A description of the security profile.
	SecurityProfileDescription *string

	// Metadata that can be used to manage the security profile.
	Tags []*types.Tag
}

type CreateSecurityProfileOutput struct {

	// The ARN of the security profile.
	SecurityProfileArn *string

	// The name you gave to the security profile.
	SecurityProfileName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateSecurityProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateSecurityProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateSecurityProfile{}, middleware.After)
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
	addOpCreateSecurityProfileValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateSecurityProfile(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateSecurityProfile(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateSecurityProfile",
	}
}
