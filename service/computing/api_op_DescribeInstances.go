// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeInstancesInput struct {
	_ struct{} `type:"structure"`

	InstanceId []string `locationName:"InstanceId" type:"list"`

	Tenancy []string `locationName:"Tenancy" type:"list"`
}

// String returns the string representation
func (s DescribeInstancesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeInstancesOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	ReservationSet []ReservationSetItem `locationName:"reservationSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeInstancesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeInstances = "DescribeInstances"

// DescribeInstancesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeInstancesRequest.
//    req := client.DescribeInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeInstances
func (c *Client) DescribeInstancesRequest(input *DescribeInstancesInput) DescribeInstancesRequest {
	op := &aws.Operation{
		Name:       opDescribeInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeInstancesInput{}
	}

	req := c.newRequest(op, input, &DescribeInstancesOutput{})
	return DescribeInstancesRequest{Request: req, Input: input, Copy: c.DescribeInstancesRequest}
}

// DescribeInstancesRequest is the request type for the
// DescribeInstances API operation.
type DescribeInstancesRequest struct {
	*aws.Request
	Input *DescribeInstancesInput
	Copy  func(*DescribeInstancesInput) DescribeInstancesRequest
}

// Send marshals and sends the DescribeInstances API request.
func (r DescribeInstancesRequest) Send(ctx context.Context) (*DescribeInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstancesResponse{
		DescribeInstancesOutput: r.Request.Data.(*DescribeInstancesOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstancesResponse is the response type for the
// DescribeInstances API operation.
type DescribeInstancesResponse struct {
	*DescribeInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstances request.
func (r *DescribeInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
