// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/organizations/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the AWS accounts that are designated as delegated administrators in this
// organization. This operation can be called only from the organization's
// management account or by a member account that is a delegated administrator for
// an AWS service.
func (c *Client) ListDelegatedAdministrators(ctx context.Context, params *ListDelegatedAdministratorsInput, optFns ...func(*Options)) (*ListDelegatedAdministratorsOutput, error) {
	if params == nil {
		params = &ListDelegatedAdministratorsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDelegatedAdministrators", params, optFns, addOperationListDelegatedAdministratorsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDelegatedAdministratorsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDelegatedAdministratorsInput struct {

	// The total number of results that you want included on each page of the response.
	// If you do not include this parameter, it defaults to a value that is specific to
	// the operation. If additional items exist beyond the maximum you specify, the
	// NextToken response element is present and has a value (is not null). Include
	// that value as the NextToken request parameter in the next call to the operation
	// to get the next part of the results. Note that Organizations might return fewer
	// results than the maximum even when there are more results available. You should
	// check NextToken after every operation to ensure that you receive all of the
	// results.
	MaxResults *int32

	// The parameter for receiving additional results if you receive a NextToken
	// response in a previous request. A NextToken response indicates that more output
	// is available. Set this parameter to the value of the previous call's NextToken
	// response to indicate where the output should continue from.
	NextToken *string

	// Specifies a service principal name. If specified, then the operation lists the
	// delegated administrators only for the specified service. If you don't specify a
	// service principal, the operation lists all delegated administrators for all
	// services in your organization.
	ServicePrincipal *string
}

type ListDelegatedAdministratorsOutput struct {

	// The list of delegated administrators in your organization.
	DelegatedAdministrators []*types.DelegatedAdministrator

	// If present, indicates that more output is available than is included in the
	// current response. Use this value in the NextToken request parameter in a
	// subsequent call to the operation to get the next part of the output. You should
	// repeat this until the NextToken response element comes back as null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDelegatedAdministratorsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDelegatedAdministrators{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDelegatedAdministrators{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDelegatedAdministrators(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDelegatedAdministrators(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "ListDelegatedAdministrators",
	}
}
