// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Validates the syntax of a matchmaking rule or rule set. This operation checks
// that the rule set is using syntactically correct JSON and that it conforms to
// allowed property expressions. To validate syntax, provide a rule set JSON
// string. Learn more
//
//     * Build a Rule Set
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/match-rulesets.html)
//
// Related
// operations
//
//     * CreateMatchmakingConfiguration
//
//     *
// DescribeMatchmakingConfigurations
//
//     * UpdateMatchmakingConfiguration
//
//     *
// DeleteMatchmakingConfiguration
//
//     * CreateMatchmakingRuleSet
//
//     *
// DescribeMatchmakingRuleSets
//
//     * ValidateMatchmakingRuleSet
//
//     *
// DeleteMatchmakingRuleSet
func (c *Client) ValidateMatchmakingRuleSet(ctx context.Context, params *ValidateMatchmakingRuleSetInput, optFns ...func(*Options)) (*ValidateMatchmakingRuleSetOutput, error) {
	if params == nil {
		params = &ValidateMatchmakingRuleSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ValidateMatchmakingRuleSet", params, optFns, addOperationValidateMatchmakingRuleSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ValidateMatchmakingRuleSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type ValidateMatchmakingRuleSetInput struct {

	// A collection of matchmaking rules to validate, formatted as a JSON string.
	//
	// This member is required.
	RuleSetBody *string
}

// Represents the returned data in response to a request operation.
type ValidateMatchmakingRuleSetOutput struct {

	// A response indicating whether the rule set is valid.
	Valid *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationValidateMatchmakingRuleSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpValidateMatchmakingRuleSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpValidateMatchmakingRuleSet{}, middleware.After)
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
	addOpValidateMatchmakingRuleSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opValidateMatchmakingRuleSet(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opValidateMatchmakingRuleSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ValidateMatchmakingRuleSet",
	}
}
