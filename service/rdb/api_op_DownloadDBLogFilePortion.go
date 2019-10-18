// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdb

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
)

type DownloadDBLogFilePortionInput struct {
	_ struct{} `type:"structure"`

	DBInstanceIdentifier *string `locationName:"DBInstanceIdentifier" type:"string"`

	LogFileName *string `locationName:"LogFileName" type:"string"`

	Marker *string `locationName:"Marker" type:"string"`

	NumberOfLines *int64 `locationName:"NumberOfLines" type:"integer"`
}

// String returns the string representation
func (s DownloadDBLogFilePortionInput) String() string {
	return nifcloudutil.Prettify(s)
}

type DownloadDBLogFilePortionOutput struct {
	_ struct{} `type:"structure"`

	AdditionalDataPending *bool `type:"boolean"`

	LogFileData *string `type:"string"`

	Marker *string `type:"string"`
}

// String returns the string representation
func (s DownloadDBLogFilePortionOutput) String() string {
	return nifcloudutil.Prettify(s)
}

const opDownloadDBLogFilePortion = "DownloadDBLogFilePortion"

// DownloadDBLogFilePortionRequest returns a request value for making API operation for
// NIFCLOUD RDB.
//
//    // Example sending a request using DownloadDBLogFilePortionRequest.
//    req := client.DownloadDBLogFilePortionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/rdb-2013-05-15N2013-12-16/DownloadDBLogFilePortion
func (c *Client) DownloadDBLogFilePortionRequest(input *DownloadDBLogFilePortionInput) DownloadDBLogFilePortionRequest {
	op := &aws.Operation{
		Name:       opDownloadDBLogFilePortion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DownloadDBLogFilePortionInput{}
	}

	req := c.newRequest(op, input, &DownloadDBLogFilePortionOutput{})
	return DownloadDBLogFilePortionRequest{Request: req, Input: input, Copy: c.DownloadDBLogFilePortionRequest}
}

// DownloadDBLogFilePortionRequest is the request type for the
// DownloadDBLogFilePortion API operation.
type DownloadDBLogFilePortionRequest struct {
	*aws.Request
	Input *DownloadDBLogFilePortionInput
	Copy  func(*DownloadDBLogFilePortionInput) DownloadDBLogFilePortionRequest
}

// Send marshals and sends the DownloadDBLogFilePortion API request.
func (r DownloadDBLogFilePortionRequest) Send(ctx context.Context) (*DownloadDBLogFilePortionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DownloadDBLogFilePortionResponse{
		DownloadDBLogFilePortionOutput: r.Request.Data.(*DownloadDBLogFilePortionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DownloadDBLogFilePortionResponse is the response type for the
// DownloadDBLogFilePortion API operation.
type DownloadDBLogFilePortionResponse struct {
	*DownloadDBLogFilePortionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DownloadDBLogFilePortion request.
func (r *DownloadDBLogFilePortionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}