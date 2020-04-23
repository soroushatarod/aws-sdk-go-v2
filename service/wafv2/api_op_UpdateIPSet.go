// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateIPSetInput struct {
	_ struct{} `type:"structure"`

	// Contains an array of strings that specify one or more IP addresses or blocks
	// of IP addresses in Classless Inter-Domain Routing (CIDR) notation. AWS WAF
	// supports all address ranges for IP versions IPv4 and IPv6.
	//
	// Examples:
	//
	//    * To configure AWS WAF to allow, block, or count requests that originated
	//    from the IP address 192.0.2.44, specify 192.0.2.44/32.
	//
	//    * To configure AWS WAF to allow, block, or count requests that originated
	//    from IP addresses from 192.0.2.0 to 192.0.2.255, specify 192.0.2.0/24.
	//
	//    * To configure AWS WAF to allow, block, or count requests that originated
	//    from the IP address 1111:0000:0000:0000:0000:0000:0000:0111, specify 1111:0000:0000:0000:0000:0000:0000:0111/128.
	//
	//    * To configure AWS WAF to allow, block, or count requests that originated
	//    from IP addresses 1111:0000:0000:0000:0000:0000:0000:0000 to 1111:0000:0000:0000:ffff:ffff:ffff:ffff,
	//    specify 1111:0000:0000:0000:0000:0000:0000:0000/64.
	//
	// For more information about CIDR notation, see the Wikipedia entry Classless
	// Inter-Domain Routing (https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing).
	//
	// Addresses is a required field
	Addresses []string `type:"list" required:"true"`

	// A description of the IP set that helps with identification. You cannot change
	// the description of an IP set after you create it.
	Description *string `min:"1" type:"string"`

	// A unique identifier for the set. This ID is returned in the responses to
	// create and list commands. You provide it to operations like update and delete.
	//
	// Id is a required field
	Id *string `min:"1" type:"string" required:"true"`

	// A token used for optimistic locking. AWS WAF returns a token to your get
	// and list requests, to mark the state of the entity at the time of the request.
	// To make changes to the entity associated with the token, you provide the
	// token to operations like update and delete. AWS WAF uses the token to ensure
	// that no changes have been made to the entity since you last retrieved it.
	// If a change has been made, the update fails with a WAFOptimisticLockException.
	// If this happens, perform another get, and use the new token returned by that
	// operation.
	//
	// LockToken is a required field
	LockToken *string `min:"1" type:"string" required:"true"`

	// The name of the IP set. You cannot change the name of an IPSet after you
	// create it.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// Specifies whether this is for an AWS CloudFront distribution or for a regional
	// application. A regional application can be an Application Load Balancer (ALB)
	// or an API Gateway stage.
	//
	// To work with CloudFront, you must also specify the Region US East (N. Virginia)
	// as follows:
	//
	//    * CLI - Specify the Region when you use the CloudFront scope: --scope=CLOUDFRONT
	//    --region=us-east-1.
	//
	//    * API and SDKs - For all calls, use the Region endpoint us-east-1.
	//
	// Scope is a required field
	Scope Scope `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s UpdateIPSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateIPSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateIPSetInput"}

	if s.Addresses == nil {
		invalidParams.Add(aws.NewErrParamRequired("Addresses"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}
	if s.Id != nil && len(*s.Id) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Id", 1))
	}

	if s.LockToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("LockToken"))
	}
	if s.LockToken != nil && len(*s.LockToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("LockToken", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if len(s.Scope) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Scope"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateIPSetOutput struct {
	_ struct{} `type:"structure"`

	// A token used for optimistic locking. AWS WAF returns this token to your update
	// requests. You use NextLockToken in the same manner as you use LockToken.
	NextLockToken *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateIPSetOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateIPSet = "UpdateIPSet"

// UpdateIPSetRequest returns a request value for making API operation for
// AWS WAFV2.
//
//
// This is the latest version of AWS WAF, named AWS WAFV2, released in November,
// 2019. For information, including how to migrate your AWS WAF resources from
// the prior release, see the AWS WAF Developer Guide (https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html).
//
// Updates the specified IPSet.
//
//    // Example sending a request using UpdateIPSetRequest.
//    req := client.UpdateIPSetRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/wafv2-2019-07-29/UpdateIPSet
func (c *Client) UpdateIPSetRequest(input *UpdateIPSetInput) UpdateIPSetRequest {
	op := &aws.Operation{
		Name:       opUpdateIPSet,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateIPSetInput{}
	}

	req := c.newRequest(op, input, &UpdateIPSetOutput{})
	return UpdateIPSetRequest{Request: req, Input: input, Copy: c.UpdateIPSetRequest}
}

// UpdateIPSetRequest is the request type for the
// UpdateIPSet API operation.
type UpdateIPSetRequest struct {
	*aws.Request
	Input *UpdateIPSetInput
	Copy  func(*UpdateIPSetInput) UpdateIPSetRequest
}

// Send marshals and sends the UpdateIPSet API request.
func (r UpdateIPSetRequest) Send(ctx context.Context) (*UpdateIPSetResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateIPSetResponse{
		UpdateIPSetOutput: r.Request.Data.(*UpdateIPSetOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateIPSetResponse is the response type for the
// UpdateIPSet API operation.
type UpdateIPSetResponse struct {
	*UpdateIPSetOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateIPSet request.
func (r *UpdateIPSetResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}