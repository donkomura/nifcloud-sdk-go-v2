// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeSecurityGroupsInput struct {
	_ struct{} `type:"structure"`

	Filter []RequestFilterStruct `locationName:"Filter" type:"list"`

	GroupName []string `locationName:"GroupName" type:"list"`
}

// String returns the string representation
func (s DescribeSecurityGroupsInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeSecurityGroupsOutput struct {
	_ struct{} `type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`

	SecurityGroupInfo []SecurityGroupInfoSetItem `locationName:"securityGroupInfo" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s DescribeSecurityGroupsOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeSecurityGroups = "DescribeSecurityGroups"

// DescribeSecurityGroupsRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeSecurityGroupsRequest.
//    req := client.DescribeSecurityGroupsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeSecurityGroups
func (c *Client) DescribeSecurityGroupsRequest(input *DescribeSecurityGroupsInput) DescribeSecurityGroupsRequest {
	op := &aws.Operation{
		Name:       opDescribeSecurityGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeSecurityGroupsInput{}
	}

	req := c.newRequest(op, input, &DescribeSecurityGroupsOutput{})
	return DescribeSecurityGroupsRequest{Request: req, Input: input, Copy: c.DescribeSecurityGroupsRequest}
}

// DescribeSecurityGroupsRequest is the request type for the
// DescribeSecurityGroups API operation.
type DescribeSecurityGroupsRequest struct {
	*aws.Request
	Input *DescribeSecurityGroupsInput
	Copy  func(*DescribeSecurityGroupsInput) DescribeSecurityGroupsRequest
}

// Send marshals and sends the DescribeSecurityGroups API request.
func (r DescribeSecurityGroupsRequest) Send(ctx context.Context) (*DescribeSecurityGroupsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeSecurityGroupsResponse{
		DescribeSecurityGroupsOutput: r.Request.Data.(*DescribeSecurityGroupsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeSecurityGroupsResponse is the response type for the
// DescribeSecurityGroups API operation.
type DescribeSecurityGroupsResponse struct {
	*DescribeSecurityGroupsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeSecurityGroups request.
func (r *DescribeSecurityGroupsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
