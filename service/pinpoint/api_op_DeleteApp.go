// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteAppRequest
type DeleteAppInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteAppInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteAppInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteAppInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAppInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.ApplicationId != nil {
		v := *s.ApplicationId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "application-id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteAppResponse
type DeleteAppOutput struct {
	_ struct{} `type:"structure" payload:"ApplicationResponse"`

	// Provides information about an application.
	//
	// ApplicationResponse is a required field
	ApplicationResponse *ApplicationResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteAppOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DeleteAppOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.ApplicationResponse != nil {
		v := s.ApplicationResponse

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "ApplicationResponse", v, metadata)
	}
	return nil
}

const opDeleteApp = "DeleteApp"

// DeleteAppRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Deletes an application.
//
//    // Example sending a request using DeleteAppRequest.
//    req := client.DeleteAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteApp
func (c *Client) DeleteAppRequest(input *DeleteAppInput) DeleteAppRequest {
	op := &aws.Operation{
		Name:       opDeleteApp,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}",
	}

	if input == nil {
		input = &DeleteAppInput{}
	}

	req := c.newRequest(op, input, &DeleteAppOutput{})
	return DeleteAppRequest{Request: req, Input: input, Copy: c.DeleteAppRequest}
}

// DeleteAppRequest is the request type for the
// DeleteApp API operation.
type DeleteAppRequest struct {
	*aws.Request
	Input *DeleteAppInput
	Copy  func(*DeleteAppInput) DeleteAppRequest
}

// Send marshals and sends the DeleteApp API request.
func (r DeleteAppRequest) Send(ctx context.Context) (*DeleteAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAppResponse{
		DeleteAppOutput: r.Request.Data.(*DeleteAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAppResponse is the response type for the
// DeleteApp API operation.
type DeleteAppResponse struct {
	*DeleteAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApp request.
func (r *DeleteAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}