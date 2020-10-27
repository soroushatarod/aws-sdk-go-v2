// Code generated by smithy-go-codegen DO NOT EDIT.

package wafv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from the
// prior release, see the AWS WAF Developer Guide
// (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html). Gets
// detailed information about a specified number of requests--a sample--that AWS
// WAF randomly selects from among the first 5,000 requests that your AWS resource
// received during a time range that you choose. You can specify a sample size of
// up to 500 requests, and you can specify any time range in the previous three
// hours. GetSampledRequests returns a time range, which is usually the time range
// that you specified. However, if your resource (such as a CloudFront
// distribution) received 5,000 requests before the specified time range elapsed,
// GetSampledRequests returns an updated time range. This new time range indicates
// the actual period during which AWS WAF selected the requests in the sample.
func (c *Client) GetSampledRequests(ctx context.Context, params *GetSampledRequestsInput, optFns ...func(*Options)) (*GetSampledRequestsOutput, error) {
	if params == nil {
		params = &GetSampledRequestsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSampledRequests", params, optFns, addOperationGetSampledRequestsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSampledRequestsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSampledRequestsInput struct {

	// The number of requests that you want AWS WAF to return from among the first
	// 5,000 requests that your AWS resource received during the time range. If your
	// resource received fewer requests than the value of MaxItems, GetSampledRequests
	// returns information about all of them.
	//
	// This member is required.
	MaxItems *int64

	// The metric name assigned to the Rule or RuleGroup for which you want a sample of
	// requests.
	//
	// This member is required.
	RuleMetricName *string

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB),
	// an API Gateway REST API, or an AppSync GraphQL API. To work with CloudFront, you
	// must also specify the Region US East (N. Virginia) as follows:
	//
	//     * CLI -
	// Specify the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	// --region=us-east-1.
	//
	//     * API and SDKs - For all calls, use the Region endpoint
	// us-east-1.
	//
	// This member is required.
	Scope types.Scope

	// The start date and time and the end date and time of the range for which you
	// want GetSampledRequests to return a sample of requests. You must specify the
	// times in Coordinated Universal Time (UTC) format. UTC format includes the
	// special designator, Z. For example, "2016-09-27T14:50Z". You can specify any
	// time range in the previous three hours.
	//
	// This member is required.
	TimeWindow *types.TimeWindow

	// The Amazon resource name (ARN) of the WebACL for which you want a sample of
	// requests.
	//
	// This member is required.
	WebAclArn *string
}

type GetSampledRequestsOutput struct {

	// The total number of requests from which GetSampledRequests got a sample of
	// MaxItems requests. If PopulationSize is less than MaxItems, the sample includes
	// every request that your AWS resource received during the specified time range.
	PopulationSize *int64

	// A complex type that contains detailed information about each of the requests in
	// the sample.
	SampledRequests []*types.SampledHTTPRequest

	// Usually, TimeWindow is the time range that you specified in the
	// GetSampledRequests request. However, if your AWS resource received more than
	// 5,000 requests during the time range that you specified in the request,
	// GetSampledRequests returns the time range for the first 5,000 requests. Times
	// are in Coordinated Universal Time (UTC) format.
	TimeWindow *types.TimeWindow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetSampledRequestsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetSampledRequests{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetSampledRequests{}, middleware.After)
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
	addOpGetSampledRequestsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSampledRequests(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetSampledRequests(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wafv2",
		OperationName: "GetSampledRequests",
	}
}
