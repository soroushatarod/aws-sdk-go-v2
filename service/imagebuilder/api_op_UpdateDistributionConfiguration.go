// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package imagebuilder

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type UpdateDistributionConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The idempotency token of the distribution configuration.
	//
	// ClientToken is a required field
	ClientToken *string `locationName:"clientToken" min:"1" type:"string" required:"true" idempotencyToken:"true"`

	// The description of the distribution configuration.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the distribution configuration that you
	// want to update.
	//
	// DistributionConfigurationArn is a required field
	DistributionConfigurationArn *string `locationName:"distributionConfigurationArn" type:"string" required:"true"`

	// The distributions of the distribution configuration.
	//
	// Distributions is a required field
	Distributions []Distribution `locationName:"distributions" type:"list" required:"true"`
}

// String returns the string representation
func (s UpdateDistributionConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDistributionConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDistributionConfigurationInput"}

	if s.ClientToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientToken"))
	}
	if s.ClientToken != nil && len(*s.ClientToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientToken", 1))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.DistributionConfigurationArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("DistributionConfigurationArn"))
	}

	if s.Distributions == nil {
		invalidParams.Add(aws.NewErrParamRequired("Distributions"))
	}
	if s.Distributions != nil {
		for i, v := range s.Distributions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Distributions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDistributionConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientToken string
	if s.ClientToken != nil {
		ClientToken = *s.ClientToken
	} else {
		ClientToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DistributionConfigurationArn != nil {
		v := *s.DistributionConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "distributionConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Distributions != nil {
		v := s.Distributions

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "distributions", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type UpdateDistributionConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The idempotency token used to make this request idempotent.
	ClientToken *string `locationName:"clientToken" min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the distribution configuration that was
	// updated by this request.
	DistributionConfigurationArn *string `locationName:"distributionConfigurationArn" type:"string"`

	// The request ID that uniquely identifies this request.
	RequestId *string `locationName:"requestId" min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateDistributionConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s UpdateDistributionConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ClientToken != nil {
		v := *s.ClientToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "clientToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DistributionConfigurationArn != nil {
		v := *s.DistributionConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "distributionConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opUpdateDistributionConfiguration = "UpdateDistributionConfiguration"

// UpdateDistributionConfigurationRequest returns a request value for making API operation for
// EC2 Image Builder.
//
// Updates a new distribution configuration. Distribution configurations define
// and configure the outputs of your pipeline.
//
//    // Example sending a request using UpdateDistributionConfigurationRequest.
//    req := client.UpdateDistributionConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/imagebuilder-2019-12-02/UpdateDistributionConfiguration
func (c *Client) UpdateDistributionConfigurationRequest(input *UpdateDistributionConfigurationInput) UpdateDistributionConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateDistributionConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/UpdateDistributionConfiguration",
	}

	if input == nil {
		input = &UpdateDistributionConfigurationInput{}
	}

	req := c.newRequest(op, input, &UpdateDistributionConfigurationOutput{})
	return UpdateDistributionConfigurationRequest{Request: req, Input: input, Copy: c.UpdateDistributionConfigurationRequest}
}

// UpdateDistributionConfigurationRequest is the request type for the
// UpdateDistributionConfiguration API operation.
type UpdateDistributionConfigurationRequest struct {
	*aws.Request
	Input *UpdateDistributionConfigurationInput
	Copy  func(*UpdateDistributionConfigurationInput) UpdateDistributionConfigurationRequest
}

// Send marshals and sends the UpdateDistributionConfiguration API request.
func (r UpdateDistributionConfigurationRequest) Send(ctx context.Context) (*UpdateDistributionConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDistributionConfigurationResponse{
		UpdateDistributionConfigurationOutput: r.Request.Data.(*UpdateDistributionConfigurationOutput),
		response:                              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDistributionConfigurationResponse is the response type for the
// UpdateDistributionConfiguration API operation.
type UpdateDistributionConfigurationResponse struct {
	*UpdateDistributionConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDistributionConfiguration request.
func (r *UpdateDistributionConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}