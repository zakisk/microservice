// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateProductReader is a Reader for the UpdateProduct structure.
type UpdateProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateProductCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateProductCreated creates a UpdateProductCreated with default headers values
func NewUpdateProductCreated() *UpdateProductCreated {
	return &UpdateProductCreated{}
}

/*
UpdateProductCreated describes a response with status code 201, with default header values.

No content is returned by this API endpoint
*/
type UpdateProductCreated struct {
}

// IsSuccess returns true when this update product created response has a 2xx status code
func (o *UpdateProductCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update product created response has a 3xx status code
func (o *UpdateProductCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update product created response has a 4xx status code
func (o *UpdateProductCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this update product created response has a 5xx status code
func (o *UpdateProductCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this update product created response a status code equal to that given
func (o *UpdateProductCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the update product created response
func (o *UpdateProductCreated) Code() int {
	return 201
}

func (o *UpdateProductCreated) Error() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductCreated ", 201)
}

func (o *UpdateProductCreated) String() string {
	return fmt.Sprintf("[PUT /products][%d] updateProductCreated ", 201)
}

func (o *UpdateProductCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
