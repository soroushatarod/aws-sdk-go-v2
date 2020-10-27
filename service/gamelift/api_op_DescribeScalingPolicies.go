// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves all scaling policies applied to a fleet. To get a fleet's scaling
// policies, specify the fleet ID. You can filter this request by policy status,
// such as to retrieve only active scaling policies. Use the pagination parameters
// to retrieve results as a set of sequential pages. If successful, set of
// ScalingPolicy objects is returned for the fleet. A fleet may have all of its
// scaling policies suspended (StopFleetActions). This operation does not affect
// the status of the scaling policies, which remains ACTIVE. To see whether a
// fleet's scaling policies are in force or suspended, call DescribeFleetAttributes
// and check the stopped actions.
//
//     * DescribeFleetCapacity
//
//     *
// UpdateFleetCapacity
//
//     * DescribeEC2InstanceLimits
//
//     * Manage scaling
// policies:
//
//         * PutScalingPolicy (auto-scaling)
//
//         *
// DescribeScalingPolicies (auto-scaling)
//
//         * DeleteScalingPolicy
// (auto-scaling)
//
//     * Manage fleet actions:
//
//         * StartFleetActions
//
//
// * StopFleetActions
func (c *Client) DescribeScalingPolicies(ctx context.Context, params *DescribeScalingPoliciesInput, optFns ...func(*Options)) (*DescribeScalingPoliciesOutput, error) {
	if params == nil {
		params = &DescribeScalingPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeScalingPolicies", params, optFns, addOperationDescribeScalingPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeScalingPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeScalingPoliciesInput struct {

	// A unique identifier for a fleet to retrieve scaling policies for. You can use
	// either the fleet ID or ARN value.
	//
	// This member is required.
	FleetId *string

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// Scaling policy status to filter results on. A scaling policy is only in force
	// when in an ACTIVE status.
	//
	//     * ACTIVE -- The scaling policy is currently in
	// force.
	//
	//     * UPDATEREQUESTED -- A request to update the scaling policy has been
	// received.
	//
	//     * UPDATING -- A change is being made to the scaling policy.
	//
	//
	// * DELETEREQUESTED -- A request to delete the scaling policy has been received.
	//
	//
	// * DELETING -- The scaling policy is being deleted.
	//
	//     * DELETED -- The scaling
	// policy has been deleted.
	//
	//     * ERROR -- An error occurred in creating the
	// policy. It should be removed and recreated.
	StatusFilter types.ScalingStatusType
}

// Represents the returned data in response to a request operation.
type DescribeScalingPoliciesOutput struct {

	// Token that indicates where to resume retrieving results on the next call to this
	// operation. If no token is returned, these results represent the end of the list.
	NextToken *string

	// A collection of objects containing the scaling policies matching the request.
	ScalingPolicies []*types.ScalingPolicy

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeScalingPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeScalingPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeScalingPolicies{}, middleware.After)
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
	addOpDescribeScalingPoliciesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeScalingPolicies(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeScalingPolicies(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeScalingPolicies",
	}
}
