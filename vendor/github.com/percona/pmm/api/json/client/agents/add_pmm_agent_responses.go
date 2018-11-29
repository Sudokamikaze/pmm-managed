// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/percona/pmm/api/json/models"
)

// AddPMMAgentReader is a Reader for the AddPMMAgent structure.
type AddPMMAgentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPMMAgentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddPMMAgentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddPMMAgentOK creates a AddPMMAgentOK with default headers values
func NewAddPMMAgentOK() *AddPMMAgentOK {
	return &AddPMMAgentOK{}
}

/*AddPMMAgentOK handles this case with default header values.

(empty)
*/
type AddPMMAgentOK struct {
	Payload *models.InventoryAddPMMAgentResponse
}

func (o *AddPMMAgentOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/AddPMMAgent][%d] addPmmAgentOK  %+v", 200, o.Payload)
}

func (o *AddPMMAgentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryAddPMMAgentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
