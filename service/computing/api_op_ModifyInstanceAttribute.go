// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type ModifyInstanceAttributeInput struct {
	_ struct{} `type:"structure"`

	Attribute *string `locationName:"Attribute" type:"string"`

	Force *bool `locationName:"Force" type:"boolean"`

	InstanceId *string `locationName:"InstanceId" type:"string"`

	NiftyReboot *string `locationName:"NiftyReboot" type:"string"`

	Tenancy *string `locationName:"Tenancy" type:"string"`

	Value *string `locationName:"Value" type:"string"`
}

// String returns the string representation
func (s ModifyInstanceAttributeInput) String() string {
	return nifcloudutil.Prettify(s)
}

type ModifyInstanceAttributeOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s ModifyInstanceAttributeOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opModifyInstanceAttribute = "ModifyInstanceAttribute"

// ModifyInstanceAttributeRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using ModifyInstanceAttributeRequest.
//    req := client.ModifyInstanceAttributeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/ModifyInstanceAttribute
func (c *Client) ModifyInstanceAttributeRequest(input *ModifyInstanceAttributeInput) ModifyInstanceAttributeRequest {
	op := &aws.Operation{
		Name:       opModifyInstanceAttribute,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &ModifyInstanceAttributeInput{}
	}

	req := c.newRequest(op, input, &ModifyInstanceAttributeOutput{})
	return ModifyInstanceAttributeRequest{Request: req, Input: input, Copy: c.ModifyInstanceAttributeRequest}
}

// ModifyInstanceAttributeRequest is the request type for the
// ModifyInstanceAttribute API operation.
type ModifyInstanceAttributeRequest struct {
	*aws.Request
	Input *ModifyInstanceAttributeInput
	Copy  func(*ModifyInstanceAttributeInput) ModifyInstanceAttributeRequest
}

// Send marshals and sends the ModifyInstanceAttribute API request.
func (r ModifyInstanceAttributeRequest) Send(ctx context.Context) (*ModifyInstanceAttributeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyInstanceAttributeResponse{
		ModifyInstanceAttributeOutput: r.Request.Data.(*ModifyInstanceAttributeOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyInstanceAttributeResponse is the response type for the
// ModifyInstanceAttribute API operation.
type ModifyInstanceAttributeResponse struct {
	*ModifyInstanceAttributeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyInstanceAttribute request.
func (r *ModifyInstanceAttributeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
