// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DeleteDhcpOptionsInput struct {
	_ struct{} `type:"structure"`

	DhcpOptionsId *string `locationName:"DhcpOptionsId" type:"string"`
}

// String returns the string representation
func (s DeleteDhcpOptionsInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DeleteDhcpOptionsOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s DeleteDhcpOptionsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDeleteDhcpOptions = "DeleteDhcpOptions"

// DeleteDhcpOptionsRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DeleteDhcpOptionsRequest.
//    req := client.DeleteDhcpOptionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DeleteDhcpOptions
func (c *Client) DeleteDhcpOptionsRequest(input *DeleteDhcpOptionsInput) DeleteDhcpOptionsRequest {
	op := &aws.Operation{
		Name:       opDeleteDhcpOptions,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DeleteDhcpOptionsInput{}
	}

	req := c.newRequest(op, input, &DeleteDhcpOptionsOutput{})
	return DeleteDhcpOptionsRequest{Request: req, Input: input, Copy: c.DeleteDhcpOptionsRequest}
}

// DeleteDhcpOptionsRequest is the request type for the
// DeleteDhcpOptions API operation.
type DeleteDhcpOptionsRequest struct {
	*aws.Request
	Input *DeleteDhcpOptionsInput
	Copy  func(*DeleteDhcpOptionsInput) DeleteDhcpOptionsRequest
}

// Send marshals and sends the DeleteDhcpOptions API request.
func (r DeleteDhcpOptionsRequest) Send(ctx context.Context) (*DeleteDhcpOptionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteDhcpOptionsResponse{
		DeleteDhcpOptionsOutput: r.Request.Data.(*DeleteDhcpOptionsOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteDhcpOptionsResponse is the response type for the
// DeleteDhcpOptions API operation.
type DeleteDhcpOptionsResponse struct {
	*DeleteDhcpOptionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteDhcpOptions request.
func (r *DeleteDhcpOptionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
