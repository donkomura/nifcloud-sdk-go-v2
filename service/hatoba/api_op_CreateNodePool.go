// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package hatoba

import (
	"context"

	"github.com/alice02/nifcloud-sdk-go-v2/internal/nifcloudutil"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	// ClusterName is a required field
	ClusterName *string `location:"uri" locationName:"ClusterName" type:"string" required:"true"`

	// NodePool is a required field
	NodePool *CreateClusterRequestNodePool `locationName:"nodePool" type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateNodePoolInput) String() string {
	return nifcloudutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNodePoolInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateNodePoolInput"}

	if s.ClusterName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ClusterName"))
	}

	if s.NodePool == nil {
		invalidParams.Add(aws.NewErrParamRequired("NodePool"))
	}
	if s.NodePool != nil {
		if err := s.NodePool.Validate(); err != nil {
			invalidParams.AddNested("NodePool", err.(aws.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateNodePoolInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.NodePool != nil {
		v := s.NodePool

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "nodePool", v, metadata)
	}
	if s.ClusterName != nil {
		v := *s.ClusterName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "ClusterName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateNodePoolOutput struct {
	_ struct{} `type:"structure"`

	NodePool *NodePool `locationName:"nodePool" type:"structure"`

	RequestId *string `locationName:"requestId" type:"string"`
}

// String returns the string representation
func (s CreateNodePoolOutput) String() string {
	return nifcloudutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateNodePoolOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.NodePool != nil {
		v := s.NodePool

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "nodePool", v, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "requestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opCreateNodePool = "CreateNodePool"

// CreateNodePoolRequest returns a request value for making API operation for
// NIFCLOUD Hatoba beta.
//
//    // Example sending a request using CreateNodePoolRequest.
//    req := client.CreateNodePoolRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/hatoba-2019-03-27/CreateNodePool
func (c *Client) CreateNodePoolRequest(input *CreateNodePoolInput) CreateNodePoolRequest {
	op := &aws.Operation{
		Name:       opCreateNodePool,
		HTTPMethod: "POST",
		HTTPPath:   "/v1/clusters/{ClusterName}/nodePools",
	}

	if input == nil {
		input = &CreateNodePoolInput{}
	}

	req := c.newRequest(op, input, &CreateNodePoolOutput{})
	return CreateNodePoolRequest{Request: req, Input: input, Copy: c.CreateNodePoolRequest}
}

// CreateNodePoolRequest is the request type for the
// CreateNodePool API operation.
type CreateNodePoolRequest struct {
	*aws.Request
	Input *CreateNodePoolInput
	Copy  func(*CreateNodePoolInput) CreateNodePoolRequest
}

// Send marshals and sends the CreateNodePool API request.
func (r CreateNodePoolRequest) Send(ctx context.Context) (*CreateNodePoolResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateNodePoolResponse{
		CreateNodePoolOutput: r.Request.Data.(*CreateNodePoolOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateNodePoolResponse is the response type for the
// CreateNodePool API operation.
type CreateNodePoolResponse struct {
	*CreateNodePoolOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateNodePool request.
func (r *CreateNodePoolResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}