// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyRestoreRouterPreviousVersionInput struct {
	_ struct{} `type:"structure"`

	RouterId *string `locationName:"RouterId" type:"string"`

	RouterName *string `locationName:"RouterName" type:"string"`
}

// String returns the string representation
func (s NiftyRestoreRouterPreviousVersionInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyRestoreRouterPreviousVersionOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyRestoreRouterPreviousVersionOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyRestoreRouterPreviousVersion = "NiftyRestoreRouterPreviousVersion"

// NiftyRestoreRouterPreviousVersionRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyRestoreRouterPreviousVersionRequest.
//    req := client.NiftyRestoreRouterPreviousVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyRestoreRouterPreviousVersion
func (c *Client) NiftyRestoreRouterPreviousVersionRequest(input *NiftyRestoreRouterPreviousVersionInput) NiftyRestoreRouterPreviousVersionRequest {
	op := &aws.Operation{
		Name:       opNiftyRestoreRouterPreviousVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyRestoreRouterPreviousVersionInput{}
	}

	req := c.newRequest(op, input, &NiftyRestoreRouterPreviousVersionOutput{})
	return NiftyRestoreRouterPreviousVersionRequest{Request: req, Input: input, Copy: c.NiftyRestoreRouterPreviousVersionRequest}
}

// NiftyRestoreRouterPreviousVersionRequest is the request type for the
// NiftyRestoreRouterPreviousVersion API operation.
type NiftyRestoreRouterPreviousVersionRequest struct {
	*aws.Request
	Input *NiftyRestoreRouterPreviousVersionInput
	Copy  func(*NiftyRestoreRouterPreviousVersionInput) NiftyRestoreRouterPreviousVersionRequest
}

// Send marshals and sends the NiftyRestoreRouterPreviousVersion API request.
func (r NiftyRestoreRouterPreviousVersionRequest) Send(ctx context.Context) (*NiftyRestoreRouterPreviousVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyRestoreRouterPreviousVersionResponse{
		NiftyRestoreRouterPreviousVersionOutput: r.Request.Data.(*NiftyRestoreRouterPreviousVersionOutput),
		response:                                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyRestoreRouterPreviousVersionResponse is the response type for the
// NiftyRestoreRouterPreviousVersion API operation.
type NiftyRestoreRouterPreviousVersionResponse struct {
	*NiftyRestoreRouterPreviousVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyRestoreRouterPreviousVersion request.
func (r *NiftyRestoreRouterPreviousVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
