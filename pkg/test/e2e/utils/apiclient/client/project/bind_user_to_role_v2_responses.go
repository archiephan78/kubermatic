// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// BindUserToRoleV2Reader is a Reader for the BindUserToRoleV2 structure.
type BindUserToRoleV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BindUserToRoleV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBindUserToRoleV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewBindUserToRoleV2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewBindUserToRoleV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewBindUserToRoleV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBindUserToRoleV2OK creates a BindUserToRoleV2OK with default headers values
func NewBindUserToRoleV2OK() *BindUserToRoleV2OK {
	return &BindUserToRoleV2OK{}
}

/* BindUserToRoleV2OK describes a response with status code 200, with default header values.

RoleBinding
*/
type BindUserToRoleV2OK struct {
	Payload *models.RoleBinding
}

func (o *BindUserToRoleV2OK) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings][%d] bindUserToRoleV2OK  %+v", 200, o.Payload)
}
func (o *BindUserToRoleV2OK) GetPayload() *models.RoleBinding {
	return o.Payload
}

func (o *BindUserToRoleV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RoleBinding)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBindUserToRoleV2Unauthorized creates a BindUserToRoleV2Unauthorized with default headers values
func NewBindUserToRoleV2Unauthorized() *BindUserToRoleV2Unauthorized {
	return &BindUserToRoleV2Unauthorized{}
}

/* BindUserToRoleV2Unauthorized describes a response with status code 401, with default header values.

EmptyResponse is a empty response
*/
type BindUserToRoleV2Unauthorized struct {
}

func (o *BindUserToRoleV2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings][%d] bindUserToRoleV2Unauthorized ", 401)
}

func (o *BindUserToRoleV2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToRoleV2Forbidden creates a BindUserToRoleV2Forbidden with default headers values
func NewBindUserToRoleV2Forbidden() *BindUserToRoleV2Forbidden {
	return &BindUserToRoleV2Forbidden{}
}

/* BindUserToRoleV2Forbidden describes a response with status code 403, with default header values.

EmptyResponse is a empty response
*/
type BindUserToRoleV2Forbidden struct {
}

func (o *BindUserToRoleV2Forbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings][%d] bindUserToRoleV2Forbidden ", 403)
}

func (o *BindUserToRoleV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewBindUserToRoleV2Default creates a BindUserToRoleV2Default with default headers values
func NewBindUserToRoleV2Default(code int) *BindUserToRoleV2Default {
	return &BindUserToRoleV2Default{
		_statusCode: code,
	}
}

/* BindUserToRoleV2Default describes a response with status code -1, with default header values.

errorResponse
*/
type BindUserToRoleV2Default struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the bind user to role v2 default response
func (o *BindUserToRoleV2Default) Code() int {
	return o._statusCode
}

func (o *BindUserToRoleV2Default) Error() string {
	return fmt.Sprintf("[POST /api/v2/projects/{project_id}/clusters/{cluster_id}/roles/{namespace}/{role_id}/bindings][%d] bindUserToRoleV2 default  %+v", o._statusCode, o.Payload)
}
func (o *BindUserToRoleV2Default) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *BindUserToRoleV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
