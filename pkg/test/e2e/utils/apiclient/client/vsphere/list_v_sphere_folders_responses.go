// Code generated by go-swagger; DO NOT EDIT.

package vsphere

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"k8c.io/kubermatic/v2/pkg/test/e2e/utils/apiclient/models"
)

// ListVSphereFoldersReader is a Reader for the ListVSphereFolders structure.
type ListVSphereFoldersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVSphereFoldersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVSphereFoldersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVSphereFoldersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVSphereFoldersOK creates a ListVSphereFoldersOK with default headers values
func NewListVSphereFoldersOK() *ListVSphereFoldersOK {
	return &ListVSphereFoldersOK{}
}

/* ListVSphereFoldersOK describes a response with status code 200, with default header values.

VSphereFolder
*/
type ListVSphereFoldersOK struct {
	Payload []*models.VSphereFolder
}

func (o *ListVSphereFoldersOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/vsphere/folders][%d] listVSphereFoldersOK  %+v", 200, o.Payload)
}
func (o *ListVSphereFoldersOK) GetPayload() []*models.VSphereFolder {
	return o.Payload
}

func (o *ListVSphereFoldersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVSphereFoldersDefault creates a ListVSphereFoldersDefault with default headers values
func NewListVSphereFoldersDefault(code int) *ListVSphereFoldersDefault {
	return &ListVSphereFoldersDefault{
		_statusCode: code,
	}
}

/* ListVSphereFoldersDefault describes a response with status code -1, with default header values.

errorResponse
*/
type ListVSphereFoldersDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the list v sphere folders default response
func (o *ListVSphereFoldersDefault) Code() int {
	return o._statusCode
}

func (o *ListVSphereFoldersDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/providers/vsphere/folders][%d] listVSphereFolders default  %+v", o._statusCode, o.Payload)
}
func (o *ListVSphereFoldersDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ListVSphereFoldersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
