// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreatePortfolioInput
type CreatePortfolioInput struct {
	_ struct{} `type:"structure"`

	// The language code.
	//
	//    * en - English (default)
	//
	//    * jp - Japanese
	//
	//    * zh - Chinese
	AcceptLanguage *string `type:"string"`

	// The description of the portfolio.
	Description *string `type:"string"`

	// The name to use for display purposes.
	//
	// DisplayName is a required field
	DisplayName *string `min:"1" type:"string" required:"true"`

	// A unique identifier that you provide to ensure idempotency. If multiple requests
	// differ only by the idempotency token, the same response is returned for each
	// repeated request.
	//
	// IdempotencyToken is a required field
	IdempotencyToken *string `min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The name of the portfolio provider.
	//
	// ProviderName is a required field
	ProviderName *string `min:"1" type:"string" required:"true"`

	// One or more tags.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreatePortfolioInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreatePortfolioInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreatePortfolioInput"}

	if s.DisplayName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DisplayName"))
	}
	if s.DisplayName != nil && len(*s.DisplayName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DisplayName", 1))
	}

	if s.IdempotencyToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("IdempotencyToken"))
	}
	if s.IdempotencyToken != nil && len(*s.IdempotencyToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IdempotencyToken", 1))
	}

	if s.ProviderName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProviderName"))
	}
	if s.ProviderName != nil && len(*s.ProviderName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProviderName", 1))
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreatePortfolioOutput
type CreatePortfolioOutput struct {
	_ struct{} `type:"structure"`

	// Information about the portfolio.
	PortfolioDetail *PortfolioDetail `type:"structure"`

	// Information about the tags associated with the portfolio.
	Tags []Tag `type:"list"`
}

// String returns the string representation
func (s CreatePortfolioOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreatePortfolio = "CreatePortfolio"

// CreatePortfolioRequest returns a request value for making API operation for
// AWS Service Catalog.
//
// Creates a portfolio.
//
//    // Example sending a request using CreatePortfolioRequest.
//    req := client.CreatePortfolioRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicecatalog-2015-12-10/CreatePortfolio
func (c *Client) CreatePortfolioRequest(input *CreatePortfolioInput) CreatePortfolioRequest {
	op := &aws.Operation{
		Name:       opCreatePortfolio,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreatePortfolioInput{}
	}

	req := c.newRequest(op, input, &CreatePortfolioOutput{})
	return CreatePortfolioRequest{Request: req, Input: input, Copy: c.CreatePortfolioRequest}
}

// CreatePortfolioRequest is the request type for the
// CreatePortfolio API operation.
type CreatePortfolioRequest struct {
	*aws.Request
	Input *CreatePortfolioInput
	Copy  func(*CreatePortfolioInput) CreatePortfolioRequest
}

// Send marshals and sends the CreatePortfolio API request.
func (r CreatePortfolioRequest) Send(ctx context.Context) (*CreatePortfolioResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreatePortfolioResponse{
		CreatePortfolioOutput: r.Request.Data.(*CreatePortfolioOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreatePortfolioResponse is the response type for the
// CreatePortfolio API operation.
type CreatePortfolioResponse struct {
	*CreatePortfolioOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreatePortfolio request.
func (r *CreatePortfolioResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}