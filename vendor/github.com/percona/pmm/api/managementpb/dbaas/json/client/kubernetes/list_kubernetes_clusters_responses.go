// Code generated by go-swagger; DO NOT EDIT.

package kubernetes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListKubernetesClustersReader is a Reader for the ListKubernetesClusters structure.
type ListKubernetesClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListKubernetesClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListKubernetesClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListKubernetesClustersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListKubernetesClustersOK creates a ListKubernetesClustersOK with default headers values
func NewListKubernetesClustersOK() *ListKubernetesClustersOK {
	return &ListKubernetesClustersOK{}
}

/*ListKubernetesClustersOK handles this case with default header values.

A successful response.
*/
type ListKubernetesClustersOK struct {
	Payload *ListKubernetesClustersOKBody
}

func (o *ListKubernetesClustersOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/List][%d] listKubernetesClustersOk  %+v", 200, o.Payload)
}

func (o *ListKubernetesClustersOK) GetPayload() *ListKubernetesClustersOKBody {
	return o.Payload
}

func (o *ListKubernetesClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListKubernetesClustersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListKubernetesClustersDefault creates a ListKubernetesClustersDefault with default headers values
func NewListKubernetesClustersDefault(code int) *ListKubernetesClustersDefault {
	return &ListKubernetesClustersDefault{
		_statusCode: code,
	}
}

/*ListKubernetesClustersDefault handles this case with default header values.

An unexpected error response.
*/
type ListKubernetesClustersDefault struct {
	_statusCode int

	Payload *ListKubernetesClustersDefaultBody
}

// Code gets the status code for the list kubernetes clusters default response
func (o *ListKubernetesClustersDefault) Code() int {
	return o._statusCode
}

func (o *ListKubernetesClustersDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/DBaaS/Kubernetes/List][%d] ListKubernetesClusters default  %+v", o._statusCode, o.Payload)
}

func (o *ListKubernetesClustersDefault) GetPayload() *ListKubernetesClustersDefaultBody {
	return o.Payload
}

func (o *ListKubernetesClustersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ListKubernetesClustersDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DetailsItems0 details items0
swagger:model DetailsItems0
*/
type DetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this details items0
func (o *DetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DetailsItems0) UnmarshalBinary(b []byte) error {
	var res DetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*KubernetesClustersItems0 Cluster contains public info about kubernetes cluster.
swagger:model KubernetesClustersItems0
*/
type KubernetesClustersItems0 struct {

	// Kubernetes cluster name.
	KubernetesClusterName string `json:"kubernetes_cluster_name,omitempty"`
}

// Validate validates this kubernetes clusters items0
func (o *KubernetesClustersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *KubernetesClustersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *KubernetesClustersItems0) UnmarshalBinary(b []byte) error {
	var res KubernetesClustersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListKubernetesClustersDefaultBody list kubernetes clusters default body
swagger:model ListKubernetesClustersDefaultBody
*/
type ListKubernetesClustersDefaultBody struct {

	// error
	Error string `json:"error,omitempty"`

	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*DetailsItems0 `json:"details"`
}

// Validate validates this list kubernetes clusters default body
func (o *ListKubernetesClustersDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListKubernetesClustersDefaultBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ListKubernetesClusters default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersDefaultBody) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ListKubernetesClustersOKBody list kubernetes clusters OK body
swagger:model ListKubernetesClustersOKBody
*/
type ListKubernetesClustersOKBody struct {

	// Kubernetes clusters.
	KubernetesClusters []*KubernetesClustersItems0 `json:"kubernetes_clusters"`
}

// Validate validates this list kubernetes clusters OK body
func (o *ListKubernetesClustersOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateKubernetesClusters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListKubernetesClustersOKBody) validateKubernetesClusters(formats strfmt.Registry) error {

	if swag.IsZero(o.KubernetesClusters) { // not required
		return nil
	}

	for i := 0; i < len(o.KubernetesClusters); i++ {
		if swag.IsZero(o.KubernetesClusters[i]) { // not required
			continue
		}

		if o.KubernetesClusters[i] != nil {
			if err := o.KubernetesClusters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listKubernetesClustersOk" + "." + "kubernetes_clusters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListKubernetesClustersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListKubernetesClustersOKBody) UnmarshalBinary(b []byte) error {
	var res ListKubernetesClustersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
