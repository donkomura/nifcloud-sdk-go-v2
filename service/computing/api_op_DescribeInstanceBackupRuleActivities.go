// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package computing

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DescribeInstanceBackupRuleActivitiesInput struct {
	_ struct{} `type:"structure"`

	Duration *int64 `locationName:"Duration" type:"integer"`

	EndDateTime *string `locationName:"EndDateTime" type:"string"`

	InstanceBackupRuleId *string `locationName:"InstanceBackupRuleId" type:"string"`

	MaxRecords *int64 `locationName:"MaxRecords" type:"integer"`
}

// String returns the string representation
func (s DescribeInstanceBackupRuleActivitiesInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DescribeInstanceBackupRuleActivitiesOutput struct {
	_ struct{} `type:"structure"`

	ActivitiesSet []ActivitiesSetItem `locationName:"activitiesSet" locationNameList:"item" type:"list"`

	InstanceBackupRuleId *string `locationName:"instanceBackupRuleId" type:"string"`

	InstanceBackupRuleName *string `locationName:"instanceBackupRuleName" type:"string"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s DescribeInstanceBackupRuleActivitiesOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDescribeInstanceBackupRuleActivities = "DescribeInstanceBackupRuleActivities"

// DescribeInstanceBackupRuleActivitiesRequest returns a request value for making API operation for
// NIFCLOUD Computing.
//
//    // Example sending a request using DescribeInstanceBackupRuleActivitiesRequest.
//    req := client.DescribeInstanceBackupRuleActivitiesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/computing-2016-11-15/DescribeInstanceBackupRuleActivities
func (c *Client) DescribeInstanceBackupRuleActivitiesRequest(input *DescribeInstanceBackupRuleActivitiesInput) DescribeInstanceBackupRuleActivitiesRequest {
	op := &aws.Operation{
		Name:       opDescribeInstanceBackupRuleActivities,
		HTTPMethod: "POST",
		HTTPPath:   "/api/",
	}

	if input == nil {
		input = &DescribeInstanceBackupRuleActivitiesInput{}
	}

	req := c.newRequest(op, input, &DescribeInstanceBackupRuleActivitiesOutput{})
	return DescribeInstanceBackupRuleActivitiesRequest{Request: req, Input: input, Copy: c.DescribeInstanceBackupRuleActivitiesRequest}
}

// DescribeInstanceBackupRuleActivitiesRequest is the request type for the
// DescribeInstanceBackupRuleActivities API operation.
type DescribeInstanceBackupRuleActivitiesRequest struct {
	*aws.Request
	Input *DescribeInstanceBackupRuleActivitiesInput
	Copy  func(*DescribeInstanceBackupRuleActivitiesInput) DescribeInstanceBackupRuleActivitiesRequest
}

// Send marshals and sends the DescribeInstanceBackupRuleActivities API request.
func (r DescribeInstanceBackupRuleActivitiesRequest) Send(ctx context.Context) (*DescribeInstanceBackupRuleActivitiesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeInstanceBackupRuleActivitiesResponse{
		DescribeInstanceBackupRuleActivitiesOutput: r.Request.Data.(*DescribeInstanceBackupRuleActivitiesOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeInstanceBackupRuleActivitiesResponse is the response type for the
// DescribeInstanceBackupRuleActivities API operation.
type DescribeInstanceBackupRuleActivitiesResponse struct {
	*DescribeInstanceBackupRuleActivitiesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeInstanceBackupRuleActivities request.
func (r *DescribeInstanceBackupRuleActivitiesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
