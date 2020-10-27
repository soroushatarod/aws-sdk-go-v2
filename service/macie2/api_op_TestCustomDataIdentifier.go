// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Tests a custom data identifier.
func (c *Client) TestCustomDataIdentifier(ctx context.Context, params *TestCustomDataIdentifierInput, optFns ...func(*Options)) (*TestCustomDataIdentifierOutput, error) {
	if params == nil {
		params = &TestCustomDataIdentifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestCustomDataIdentifier", params, optFns, addOperationTestCustomDataIdentifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestCustomDataIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestCustomDataIdentifierInput struct {

	// The regular expression (regex) that defines the pattern to match. The expression
	// can contain as many as 512 characters.
	//
	// This member is required.
	Regex *string

	// The sample text to inspect by using the custom data identifier. The text can
	// contain as many as 1,000 characters.
	//
	// This member is required.
	SampleText *string

	// An array that lists specific character sequences (ignore words) to exclude from
	// the results. If the text matched by the regular expression is the same as any
	// string in this array, Amazon Macie ignores it. The array can contain as many as
	// 10 ignore words. Each ignore word can contain 4 - 90 characters. Ignore words
	// are case sensitive.
	IgnoreWords []*string

	// An array that lists specific character sequences (keywords), one of which must
	// be within proximity (maximumMatchDistance) of the regular expression to match.
	// The array can contain as many as 50 keywords. Each keyword can contain 4 - 90
	// characters. Keywords aren't case sensitive.
	Keywords []*string

	// The maximum number of characters that can exist between text that matches the
	// regex pattern and the character sequences specified by the keywords array. Macie
	// includes or excludes a result based on the proximity of a keyword to text that
	// matches the regex pattern. The distance can be 1 - 300 characters. The default
	// value is 50.
	MaximumMatchDistance *int32
}

type TestCustomDataIdentifierOutput struct {

	// The number of instances of sample text that matched the detection criteria
	// specified in the custom data identifier.
	MatchCount *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationTestCustomDataIdentifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestCustomDataIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestCustomDataIdentifier{}, middleware.After)
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
	addOpTestCustomDataIdentifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestCustomDataIdentifier(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opTestCustomDataIdentifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "TestCustomDataIdentifier",
	}
}
