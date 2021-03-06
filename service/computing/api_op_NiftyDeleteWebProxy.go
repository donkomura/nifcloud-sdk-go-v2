// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyDeleteWebProxyInput struct {
	_ struct{} `type:"structure"`

	Agreement *bool `locationName:"Agreement" type:"boolean"`

	RouterId *string `locationName:"RouterId" type:"string"`

	RouterName *string `locationName:"RouterName" type:"string"`
}

// String returns the string representation
func (s NiftyDeleteWebProxyInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyDeleteWebProxyOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyDeleteWebProxyOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyDeleteWebProxy = "NiftyDeleteWebProxy"

// NiftyDeleteWebProxyRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyDeleteWebProxyRequest.
//    req := client.NiftyDeleteWebProxyRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyDeleteWebProxy
func (c *Client) NiftyDeleteWebProxyRequest(input *NiftyDeleteWebProxyInput) NiftyDeleteWebProxyRequest {
	op := &aws.Operation{
		Name:       opNiftyDeleteWebProxy,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyDeleteWebProxyInput{}
	}

	req := c.newRequest(op, input, &NiftyDeleteWebProxyOutput{})
	return NiftyDeleteWebProxyRequest{Request: req, Input: input, Copy: c.NiftyDeleteWebProxyRequest}
}

// NiftyDeleteWebProxyRequest is the request type for the
// NiftyDeleteWebProxy API operation.
type NiftyDeleteWebProxyRequest struct {
	*aws.Request
	Input *NiftyDeleteWebProxyInput
	Copy  func(*NiftyDeleteWebProxyInput) NiftyDeleteWebProxyRequest
}

// Send marshals and sends the NiftyDeleteWebProxy API request.
func (r NiftyDeleteWebProxyRequest) Send(ctx context.Context) (*NiftyDeleteWebProxyResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyDeleteWebProxyResponse{
		NiftyDeleteWebProxyOutput: r.Request.Data.(*NiftyDeleteWebProxyOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyDeleteWebProxyResponse is the response type for the
// NiftyDeleteWebProxy API operation.
type NiftyDeleteWebProxyResponse struct {
	*NiftyDeleteWebProxyOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyDeleteWebProxy request.
func (r *NiftyDeleteWebProxyResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
