// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm/api/json/models"
)

// ChangeRemoteNodeReader is a Reader for the ChangeRemoteNode structure.
type ChangeRemoteNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeRemoteNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeRemoteNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChangeRemoteNodeOK creates a ChangeRemoteNodeOK with default headers values
func NewChangeRemoteNodeOK() *ChangeRemoteNodeOK {
	return &ChangeRemoteNodeOK{}
}

/*ChangeRemoteNodeOK handles this case with default header values.

(empty)
*/
type ChangeRemoteNodeOK struct {
	Payload *models.InventoryChangeRemoteNodeResponse
}

func (o *ChangeRemoteNodeOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Nodes/ChangeRemote][%d] changeRemoteNodeOK  %+v", 200, o.Payload)
}

func (o *ChangeRemoteNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryChangeRemoteNodeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
