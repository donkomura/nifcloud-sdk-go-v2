// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeResourcesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeResourcesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeResourcesOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	ResourceInfo *ResourceInfo `locationName:"resourceInfo" type:"structure"`
}

// String returns the string representation
func (s DescribeResourcesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeResources = "DescribeResources"

// DescribeResourcesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeResourcesRequest.
//    req := client.DescribeResourcesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeResources
func (c *Client) DescribeResourcesRequest(input *DescribeResourcesInput) DescribeResourcesRequest {
	op := &aws.Operation{
		Name:       opDescribeResources,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeResourcesInput{}
	}

	req := c.newRequest(op, input, &DescribeResourcesOutput{})
	return DescribeResourcesRequest{Request: req, Input: input, Copy: c.DescribeResourcesRequest}
}

// DescribeResourcesRequest is the request type for the
// DescribeResources API operation.
type DescribeResourcesRequest struct {
	*aws.Request
	Input *DescribeResourcesInput
	Copy  func(*DescribeResourcesInput) DescribeResourcesRequest
}

// Send marshals and sends the DescribeResources API request.
func (r DescribeResourcesRequest) Send(ctx context.Context) (*DescribeResourcesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeResourcesResponse{
		DescribeResourcesOutput: r.Request.Data.(*DescribeResourcesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeResourcesResponse is the response type for the
// DescribeResources API operation.
type DescribeResourcesResponse struct {
	*DescribeResourcesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeResources request.
func (r *DescribeResourcesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
