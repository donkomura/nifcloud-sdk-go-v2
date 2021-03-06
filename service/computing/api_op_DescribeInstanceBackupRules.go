// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/aokumasan/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeInstanceBackupRulesInput struct {
	_ struct{} `type:"structure"`

	InstanceBackupRuleId []string `locationName:"InstanceBackupRuleId" type:"list"`
}

// String returns the string representation
func (s DescribeInstanceBackupRulesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeInstanceBackupRulesOutput struct {
	_ struct{} `type:"structure"`

	InstanceBackupRulesSet []InstanceBackupRulesSetItem `locationName:"instanceBackupRulesSet" locationNameList:"item" type:"list"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s DescribeInstanceBackupRulesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeInstanceBackupRules = "DescribeInstanceBackupRules"

// DescribeInstanceBackupRulesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeInstanceBackupRulesRequest.
//    req := client.DescribeInstanceBackupRulesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeInstanceBackupRules
func (c *Client) DescribeInstanceBackupRulesRequest(input *DescribeInstanceBackupRulesInput) DescribeInstanceBackupRulesRequest {
	op := &aws.Operation{
		Name:       opDescribeInstanceBackupRules,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeInstanceBackupRulesInput{}
	}

	req := c.newRequest(op, input, &DescribeInstanceBackupRulesOutput{})
	return DescribeInstanceBackupRulesRequest{Request: req, Input: input, Copy: c.DescribeInstanceBackupRulesRequest}
}

// DescribeInstanceBackupRulesRequest is the request type for the
// DescribeInstanceBackupRules API operation.
type DescribeInstanceBackupRulesRequest struct {
	*aws.Request
	Input *DescribeInstanceBackupRulesInput
	Copy  func(*DescribeInstanceBackupRulesInput) DescribeInstanceBackupRulesRequest
}

// Send marshals and sends the DescribeInstanceBackupRules API request.
func (r DescribeInstanceBackupRulesRequest) Send(ctx context.Context) (*DescribeInstanceBackupRulesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstanceBackupRulesResponse{
		DescribeInstanceBackupRulesOutput: r.Request.Data.(*DescribeInstanceBackupRulesOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstanceBackupRulesResponse is the response type for the
// DescribeInstanceBackupRules API operation.
type DescribeInstanceBackupRulesResponse struct {
	*DescribeInstanceBackupRulesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstanceBackupRules request.
func (r *DescribeInstanceBackupRulesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
