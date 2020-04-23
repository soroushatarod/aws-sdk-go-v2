// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DeleteTemplateAliasInput struct {
	_ struct{} `type:"structure"`

	// The name for the template alias. If you name a specific alias, you delete
	// the version that the alias points to. You can specify the latest version
	// of the template by providing the keyword $LATEST in the AliasName parameter.
	//
	// AliasName is a required field
	AliasName *string `location:"uri" locationName:"AliasName" min:"1" type:"string" required:"true"`

	// The ID of the AWS account that contains the item to delete.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the template that the specified alias is for.
	//
	// TemplateId is a required field
	TemplateId *string `location:"uri" locationName:"TemplateId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTemplateAliasInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTemplateAliasInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTemplateAliasInput"}

	if s.AliasName == nil {
		invalidParams.Add(aws.NewErrParamRequired("AliasName"))
	}
	if s.AliasName != nil && len(*s.AliasName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AliasName", 1))
	}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.TemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TemplateId"))
	}
	if s.TemplateId != nil && len(*s.TemplateId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("TemplateId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTemplateAliasInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AliasName != nil {
		v := *s.AliasName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AliasName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DeleteTemplateAliasOutput struct {
	_ struct{} `type:"structure"`

	// The name for the template alias.
	AliasName *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of the resource.
	Arn *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// An ID for the template associated with the deletion.
	TemplateId *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DeleteTemplateAliasOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteTemplateAliasOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AliasName != nil {
		v := *s.AliasName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AliasName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.TemplateId != nil {
		v := *s.TemplateId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TemplateId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opDeleteTemplateAlias = "DeleteTemplateAlias"

// DeleteTemplateAliasRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Deletes the item that the specified template alias points to. If you provide
// a specific alias, you delete the version of the template that the alias points
// to.
//
//    // Example sending a request using DeleteTemplateAliasRequest.
//    req := client.DeleteTemplateAliasRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DeleteTemplateAlias
func (c *Client) DeleteTemplateAliasRequest(input *DeleteTemplateAliasInput) DeleteTemplateAliasRequest {
	op := &aws.Operation{
		Name:       opDeleteTemplateAlias,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}/aliases/{AliasName}",
	}

	if input == nil {
		input = &DeleteTemplateAliasInput{}
	}

	req := c.newRequest(op, input, &DeleteTemplateAliasOutput{})
	return DeleteTemplateAliasRequest{Request: req, Input: input, Copy: c.DeleteTemplateAliasRequest}
}

// DeleteTemplateAliasRequest is the request type for the
// DeleteTemplateAlias API operation.
type DeleteTemplateAliasRequest struct {
	*aws.Request
	Input *DeleteTemplateAliasInput
	Copy  func(*DeleteTemplateAliasInput) DeleteTemplateAliasRequest
}

// Send marshals and sends the DeleteTemplateAlias API request.
func (r DeleteTemplateAliasRequest) Send(ctx context.Context) (*DeleteTemplateAliasResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteTemplateAliasResponse{
		DeleteTemplateAliasOutput: r.Request.Data.(*DeleteTemplateAliasOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteTemplateAliasResponse is the response type for the
// DeleteTemplateAlias API operation.
type DeleteTemplateAliasResponse struct {
	*DeleteTemplateAliasOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteTemplateAlias request.
func (r *DeleteTemplateAliasResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}