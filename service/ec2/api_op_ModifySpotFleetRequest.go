// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the specified Spot Fleet request. You can only modify a Spot Fleet
// request of type maintain. While the Spot Fleet request is being modified, it is
// in the modifying state. To scale up your Spot Fleet, increase its target
// capacity. The Spot Fleet launches the additional Spot Instances according to the
// allocation strategy for the Spot Fleet request. If the allocation strategy is
// lowestPrice, the Spot Fleet launches instances using the Spot Instance pool with
// the lowest price. If the allocation strategy is diversified, the Spot Fleet
// distributes the instances across the Spot Instance pools. If the allocation
// strategy is capacityOptimized, Spot Fleet launches instances from Spot Instance
// pools with optimal capacity for the number of instances that are launching. To
// scale down your Spot Fleet, decrease its target capacity. First, the Spot Fleet
// cancels any open requests that exceed the new target capacity. You can request
// that the Spot Fleet terminate Spot Instances until the size of the fleet no
// longer exceeds the new target capacity. If the allocation strategy is
// lowestPrice, the Spot Fleet terminates the instances with the highest price per
// unit. If the allocation strategy is capacityOptimized, the Spot Fleet terminates
// the instances in the Spot Instance pools that have the least available Spot
// Instance capacity. If the allocation strategy is diversified, the Spot Fleet
// terminates instances across the Spot Instance pools. Alternatively, you can
// request that the Spot Fleet keep the fleet at its current size, but not replace
// any Spot Instances that are interrupted or that you terminate manually. If you
// are finished with your Spot Fleet for now, but will use it again later, you can
// set the target capacity to 0.
func (c *Client) ModifySpotFleetRequest(ctx context.Context, params *ModifySpotFleetRequestInput, optFns ...func(*Options)) (*ModifySpotFleetRequestOutput, error) {
	if params == nil {
		params = &ModifySpotFleetRequestInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifySpotFleetRequest", params, optFns, addOperationModifySpotFleetRequestMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifySpotFleetRequestOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for ModifySpotFleetRequest.
type ModifySpotFleetRequestInput struct {

	// The ID of the Spot Fleet request.
	//
	// This member is required.
	SpotFleetRequestId *string

	// Indicates whether running Spot Instances should be terminated if the target
	// capacity of the Spot Fleet request is decreased below the current size of the
	// Spot Fleet.
	ExcessCapacityTerminationPolicy types.ExcessCapacityTerminationPolicy

	// The launch template and overrides. You can only use this parameter if you
	// specified a launch template (LaunchTemplateConfigs) in your Spot Fleet request.
	// If you specified LaunchSpecifications in your Spot Fleet request, then omit this
	// parameter.
	LaunchTemplateConfigs []*types.LaunchTemplateConfig

	// The number of On-Demand Instances in the fleet.
	OnDemandTargetCapacity *int32

	// The size of the fleet.
	TargetCapacity *int32
}

// Contains the output of ModifySpotFleetRequest.
type ModifySpotFleetRequestOutput struct {

	// Is true if the request succeeds, and an error otherwise.
	Return *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifySpotFleetRequestMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifySpotFleetRequest{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifySpotFleetRequest{}, middleware.After)
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
	addOpModifySpotFleetRequestValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifySpotFleetRequest(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifySpotFleetRequest(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifySpotFleetRequest",
	}
}
