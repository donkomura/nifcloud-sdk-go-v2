// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type NiftyRebootVpnGatewaysInput struct {
	_ struct{} `type:"structure"`

	VpnGateway []RequestVpnGatewayStruct `locationName:"VpnGateway" type:"list"`
}

// String returns the string representation
func (s NiftyRebootVpnGatewaysInput) String() string {
	return nifcloudutil.Prettify(s)
}

type NiftyRebootVpnGatewaysOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	Return *bool `locationName:"return" type:"boolean"`
}

// String returns the string representation
func (s NiftyRebootVpnGatewaysOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opNiftyRebootVpnGateways = "NiftyRebootVpnGateways"

// NiftyRebootVpnGatewaysRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using NiftyRebootVpnGatewaysRequest.
//    req := client.NiftyRebootVpnGatewaysRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/NiftyRebootVpnGateways
func (c *Client) NiftyRebootVpnGatewaysRequest(input *NiftyRebootVpnGatewaysInput) NiftyRebootVpnGatewaysRequest {
	op := &aws.Operation{
		Name:       opNiftyRebootVpnGateways,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &NiftyRebootVpnGatewaysInput{}
	}

	req := c.newRequest(op, input, &NiftyRebootVpnGatewaysOutput{})
	return NiftyRebootVpnGatewaysRequest{Request: req, Input: input, Copy: c.NiftyRebootVpnGatewaysRequest}
}

// NiftyRebootVpnGatewaysRequest is the request type for the
// NiftyRebootVpnGateways API operation.
type NiftyRebootVpnGatewaysRequest struct {
	*aws.Request
	Input *NiftyRebootVpnGatewaysInput
	Copy  func(*NiftyRebootVpnGatewaysInput) NiftyRebootVpnGatewaysRequest
}

// Send marshals and sends the NiftyRebootVpnGateways API request.
func (r NiftyRebootVpnGatewaysRequest) Send(ctx context.Context) (*NiftyRebootVpnGatewaysResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &NiftyRebootVpnGatewaysResponse{
		NiftyRebootVpnGatewaysOutput: r.Request.Data.(*NiftyRebootVpnGatewaysOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NiftyRebootVpnGatewaysResponse is the response type for the
// NiftyRebootVpnGateways API operation.
type NiftyRebootVpnGatewaysResponse struct {
	*NiftyRebootVpnGatewaysOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// NiftyRebootVpnGateways request.
func (r *NiftyRebootVpnGatewaysResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
