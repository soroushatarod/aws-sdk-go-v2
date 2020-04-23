// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package chime

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateMeetingInput struct {
	_ struct{} `type:"structure"`

	// The unique identifier for the client request. Use a different token for different
	// meetings.
	//
	// ClientRequestToken is a required field
	ClientRequestToken *string `min:"2" type:"string" required:"true" idempotencyToken:"true" sensitive:"true"`

	// The external meeting ID.
	ExternalMeetingId *string `min:"2" type:"string" sensitive:"true"`

	// The Region in which to create the meeting. Available values: ap-northeast-1,
	// ap-southeast-1, ap-southeast-2, ca-central-1, eu-central-1, eu-north-1, eu-west-1,
	// eu-west-2, eu-west-3, sa-east-1, us-east-1, us-east-2, us-west-1, us-west-2.
	MediaRegion *string `type:"string"`

	// Reserved.
	MeetingHostId *string `min:"2" type:"string" sensitive:"true"`

	// The configuration for resource targets to receive notifications when meeting
	// and attendee events occur.
	NotificationsConfiguration *MeetingNotificationConfiguration `type:"structure"`

	// The tag key-value pairs.
	Tags []Tag `min:"1" type:"list"`
}

// String returns the string representation
func (s CreateMeetingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMeetingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateMeetingInput"}

	if s.ClientRequestToken == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClientRequestToken"))
	}
	if s.ClientRequestToken != nil && len(*s.ClientRequestToken) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ClientRequestToken", 2))
	}
	if s.ExternalMeetingId != nil && len(*s.ExternalMeetingId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ExternalMeetingId", 2))
	}
	if s.MeetingHostId != nil && len(*s.MeetingHostId) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("MeetingHostId", 2))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.NotificationsConfiguration != nil {
		if err := s.NotificationsConfiguration.Validate(); err != nil {
			invalidParams.AddNested("NotificationsConfiguration", err.(aws.ErrInvalidParams))
		}
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

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMeetingInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	var ClientRequestToken string
	if s.ClientRequestToken != nil {
		ClientRequestToken = *s.ClientRequestToken
	} else {
		ClientRequestToken = protocol.GetIdempotencyToken()
	}
	{
		v := ClientRequestToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ClientRequestToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ExternalMeetingId != nil {
		v := *s.ExternalMeetingId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "ExternalMeetingId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MediaRegion != nil {
		v := *s.MediaRegion

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MediaRegion", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeetingHostId != nil {
		v := *s.MeetingHostId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "MeetingHostId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.NotificationsConfiguration != nil {
		v := s.NotificationsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "NotificationsConfiguration", v, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Tags", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

type CreateMeetingOutput struct {
	_ struct{} `type:"structure"`

	// The meeting information, including the meeting ID and MediaPlacement.
	Meeting *Meeting `type:"structure"`
}

// String returns the string representation
func (s CreateMeetingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateMeetingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Meeting != nil {
		v := s.Meeting

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Meeting", v, metadata)
	}
	return nil
}

const opCreateMeeting = "CreateMeeting"

// CreateMeetingRequest returns a request value for making API operation for
// Amazon Chime.
//
// Creates a new Amazon Chime SDK meeting in the specified media Region with
// no initial attendees. For more information about the Amazon Chime SDK, see
// Using the Amazon Chime SDK (https://docs.aws.amazon.com/chime/latest/dg/meetings-sdk.html)
// in the Amazon Chime Developer Guide.
//
//    // Example sending a request using CreateMeetingRequest.
//    req := client.CreateMeetingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/chime-2018-05-01/CreateMeeting
func (c *Client) CreateMeetingRequest(input *CreateMeetingInput) CreateMeetingRequest {
	op := &aws.Operation{
		Name:       opCreateMeeting,
		HTTPMethod: "POST",
		HTTPPath:   "/meetings",
	}

	if input == nil {
		input = &CreateMeetingInput{}
	}

	req := c.newRequest(op, input, &CreateMeetingOutput{})
	return CreateMeetingRequest{Request: req, Input: input, Copy: c.CreateMeetingRequest}
}

// CreateMeetingRequest is the request type for the
// CreateMeeting API operation.
type CreateMeetingRequest struct {
	*aws.Request
	Input *CreateMeetingInput
	Copy  func(*CreateMeetingInput) CreateMeetingRequest
}

// Send marshals and sends the CreateMeeting API request.
func (r CreateMeetingRequest) Send(ctx context.Context) (*CreateMeetingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateMeetingResponse{
		CreateMeetingOutput: r.Request.Data.(*CreateMeetingOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateMeetingResponse is the response type for the
// CreateMeeting API operation.
type CreateMeetingResponse struct {
	*CreateMeetingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateMeeting request.
func (r *CreateMeetingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}