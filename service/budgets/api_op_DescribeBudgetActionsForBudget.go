// Code generated by smithy-go-codegen DO NOT EDIT.

package budgets

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/budgets/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes all of the budget actions for a budget.
func (c *Client) DescribeBudgetActionsForBudget(ctx context.Context, params *DescribeBudgetActionsForBudgetInput, optFns ...func(*Options)) (*DescribeBudgetActionsForBudgetOutput, error) {
	if params == nil {
		params = &DescribeBudgetActionsForBudgetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBudgetActionsForBudget", params, optFns, addOperationDescribeBudgetActionsForBudgetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBudgetActionsForBudgetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBudgetActionsForBudgetInput struct {

	// The account ID of the user. It should be a 12-digit number.
	//
	// This member is required.
	AccountId *string

	// A string that represents the budget name. The ":" and "\" characters aren't
	// allowed.
	//
	// This member is required.
	BudgetName *string

	// An integer that represents how many entries a paginated response contains. The
	// maximum is 100.
	MaxResults *int32

	// A generic string.
	NextToken *string
}

type DescribeBudgetActionsForBudgetOutput struct {

	// A list of the budget action resources information.
	//
	// This member is required.
	Actions []*types.Action

	// A generic string.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBudgetActionsForBudgetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeBudgetActionsForBudget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeBudgetActionsForBudget{}, middleware.After)
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
	addOpDescribeBudgetActionsForBudgetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBudgetActionsForBudget(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBudgetActionsForBudget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "budgets",
		OperationName: "DescribeBudgetActionsForBudget",
	}
}
