// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeVpnConnectionsInput struct {
	_ struct{} `type:"structure"`

	Filter []RequestFilterStruct `locationName:"Filter" type:"list"`

	VpnConnectionId []string `locationName:"VpnConnectionId" type:"list"`
}

// String returns the string representation
func (s DescribeVpnConnectionsInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeVpnConnectionsOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	VpnConnectionSet []VpnConnectionSetItem `locationName:"vpnConnectionSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeVpnConnectionsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeVpnConnections = "DescribeVpnConnections"

// DescribeVpnConnectionsRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeVpnConnectionsRequest.
//    req := client.DescribeVpnConnectionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeVpnConnections
func (c *Client) DescribeVpnConnectionsRequest(input *DescribeVpnConnectionsInput) DescribeVpnConnectionsRequest {
	op := &aws.Operation{
		Name:       opDescribeVpnConnections,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeVpnConnectionsInput{}
	}

	req := c.newRequest(op, input, &DescribeVpnConnectionsOutput{})
	return DescribeVpnConnectionsRequest{Request: req, Input: input, Copy: c.DescribeVpnConnectionsRequest}
}

// DescribeVpnConnectionsRequest is the request type for the
// DescribeVpnConnections API operation.
type DescribeVpnConnectionsRequest struct {
	*aws.Request
	Input *DescribeVpnConnectionsInput
	Copy  func(*DescribeVpnConnectionsInput) DescribeVpnConnectionsRequest
}

// Send marshals and sends the DescribeVpnConnections API request.
func (r DescribeVpnConnectionsRequest) Send(ctx context.Context) (*DescribeVpnConnectionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVpnConnectionsResponse{
		DescribeVpnConnectionsOutput: r.Request.Data.(*DescribeVpnConnectionsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVpnConnectionsResponse is the response type for the
// DescribeVpnConnections API operation.
type DescribeVpnConnectionsResponse struct {
	*DescribeVpnConnectionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVpnConnections request.
func (r *DescribeVpnConnectionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
