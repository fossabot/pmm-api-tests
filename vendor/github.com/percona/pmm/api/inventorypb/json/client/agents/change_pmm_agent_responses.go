// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangePMMAgentReader is a Reader for the ChangePMMAgent structure.
type ChangePMMAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangePMMAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangePMMAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewChangePMMAgentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangePMMAgentOK creates a ChangePMMAgentOK with default headers values
func NewChangePMMAgentOK() *ChangePMMAgentOK {
	return &ChangePMMAgentOK{}
}

/*ChangePMMAgentOK handles this case with default header values.

A successful response.
*/
type ChangePMMAgentOK struct {
	Payload *ChangePMMAgentOKBody
}

func (o *ChangePMMAgentOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/ChangePMMAgent][%d] changePmmAgentOk  %+v", 200, o.Payload)
}

func (o *ChangePMMAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangePMMAgentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangePMMAgentDefault creates a ChangePMMAgentDefault with default headers values
func NewChangePMMAgentDefault(code int) *ChangePMMAgentDefault {
	return &ChangePMMAgentDefault{
		_statusCode: code,
	}
}

/*ChangePMMAgentDefault handles this case with default header values.

An error response.
*/
type ChangePMMAgentDefault struct {
	_statusCode int

	Payload *ChangePMMAgentDefaultBody
}

// Code gets the status code for the change PMM agent default response
func (o *ChangePMMAgentDefault) Code() int {
	return o._statusCode
}

func (o *ChangePMMAgentDefault) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/ChangePMMAgent][%d] ChangePMMAgent default  %+v", o._statusCode, o.Payload)
}

func (o *ChangePMMAgentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangePMMAgentDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangePMMAgentDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ChangePMMAgentDefaultBody
*/
type ChangePMMAgentDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this change PMM agent default body
func (o *ChangePMMAgentDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangePMMAgentDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePMMAgentDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangePMMAgentDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePMMAgentOKBody change PMM agent OK body
swagger:model ChangePMMAgentOKBody
*/
type ChangePMMAgentOKBody struct {

	// pmm agent
	PMMAgent *ChangePMMAgentOKBodyPMMAgent `json:"pmm_agent,omitempty"`
}

// Validate validates this change PMM agent OK body
func (o *ChangePMMAgentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePMMAgent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangePMMAgentOKBody) validatePMMAgent(formats strfmt.Registry) error {

	if swag.IsZero(o.PMMAgent) { // not required
		return nil
	}

	if o.PMMAgent != nil {
		if err := o.PMMAgent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changePmmAgentOk" + "." + "pmm_agent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangePMMAgentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePMMAgentOKBody) UnmarshalBinary(b []byte) error {
	var res ChangePMMAgentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangePMMAgentOKBodyPMMAgent PMMAgent runs on Generic on Container Node.
swagger:model ChangePMMAgentOKBodyPMMAgent
*/
type ChangePMMAgentOKBodyPMMAgent struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// True if Agent is running and connected to pmm-managed.
	Connected bool `json:"connected,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Node identifier where this instance runs.
	RunsOnNodeID string `json:"runs_on_node_id,omitempty"`
}

// Validate validates this change PMM agent OK body PMM agent
func (o *ChangePMMAgentOKBodyPMMAgent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangePMMAgentOKBodyPMMAgent) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangePMMAgentOKBodyPMMAgent) UnmarshalBinary(b []byte) error {
	var res ChangePMMAgentOKBodyPMMAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}