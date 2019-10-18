// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type RestoreClusterFromSnapshotInput struct {
	_ struct{} `type:"structure"`

	// Cluster is a required field
	Cluster *RestoreClusterFromSnapshotRequestCluster `locationName:"cluster" type:"structure" required:"true"`

	// SnapshotName is a required field
	SnapshotName *string `location:"uri" locationName:"SnapshotName" type:"string" required:"true"`
}

// String returns the string representation
func (s RestoreClusterFromSnapshotInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RestoreClusterFromSnapshotInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RestoreClusterFromSnapshotInput"}

	if s.Cluster == nil {
		invalidParams.Add(aws.NewErrParamRequired("Cluster"))
	}

	if s.SnapshotName == nil {
		invalidParams.Add(aws.NewErrParamRequired("SnapshotName"))
	}
	if s.Cluster != nil {
		if err := s.Cluster.Validate(); err != nil {
			invalidParams.AddNested("Cluster", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RestoreClusterFromSnapshotInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.Cluster != nil {
		v := s.Cluster

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cluster", v, metadata)
	}
	if s.SnapshotName != nil {
		v := *s.SnapshotName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "SnapshotName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type RestoreClusterFromSnapshotOutput struct {
	_ struct{} `type:"structure"`

	Cluster *Cluster `locationName:"cluster" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s RestoreClusterFromSnapshotOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s RestoreClusterFromSnapshotOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Cluster != nil {
		v := s.Cluster

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cluster", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opRestoreClusterFromSnapshot = "RestoreClusterFromSnapshot"

// RestoreClusterFromSnapshotRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using RestoreClusterFromSnapshotRequest.
//    req := client.RestoreClusterFromSnapshotRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/hatoba-2019-03-27/RestoreClusterFromSnapshot
func (c *Client) RestoreClusterFromSnapshotRequest(input *RestoreClusterFromSnapshotInput) RestoreClusterFromSnapshotRequest {
	op := &aws.Operation{
		Name:       opRestoreClusterFromSnapshot,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/snapshots/{SnapshotName}:restore",
	}

	if input == nil {
		input = &RestoreClusterFromSnapshotInput{}
	}

	req := c.newRequest(op, input, &RestoreClusterFromSnapshotOutput{})
	return RestoreClusterFromSnapshotRequest{Request: req, Input: input, Copy: c.RestoreClusterFromSnapshotRequest}
}

// RestoreClusterFromSnapshotRequest is the request type for the
// RestoreClusterFromSnapshot API operation.
type RestoreClusterFromSnapshotRequest struct {
	*aws.Request
	Input *RestoreClusterFromSnapshotInput
	Copy  func(*RestoreClusterFromSnapshotInput) RestoreClusterFromSnapshotRequest
}

// Send marshals and sends the RestoreClusterFromSnapshot API request.
func (r RestoreClusterFromSnapshotRequest) Send(ctx context.Context) (*RestoreClusterFromSnapshotResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &RestoreClusterFromSnapshotResponse{
		RestoreClusterFromSnapshotOutput: r.Request.Data.(*RestoreClusterFromSnapshotOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// RestoreClusterFromSnapshotResponse is the response type for the
// RestoreClusterFromSnapshot API operation.
type RestoreClusterFromSnapshotResponse struct {
	*RestoreClusterFromSnapshotOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// RestoreClusterFromSnapshot request.
func (r *RestoreClusterFromSnapshotResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
