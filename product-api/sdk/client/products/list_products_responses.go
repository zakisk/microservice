// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zakisk/microservice/product-api/sdk/models"
)

// ListProductsReader is a Reader for the ListProducts structure.
type ListProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListProductsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListProductsOK creates a ListProductsOK with default headers values
func NewListProductsOK() *ListProductsOK {
	return &ListProductsOK{}
}

/*
ListProductsOK describes a response with status code 200, with default header values.

A list of products
*/
type ListProductsOK struct {
	Payload []*models.Product
}

// IsSuccess returns true when this list products o k response has a 2xx status code
func (o *ListProductsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list products o k response has a 3xx status code
func (o *ListProductsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list products o k response has a 4xx status code
func (o *ListProductsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list products o k response has a 5xx status code
func (o *ListProductsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list products o k response a status code equal to that given
func (o *ListProductsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list products o k response
func (o *ListProductsOK) Code() int {
	return 200
}

func (o *ListProductsOK) Error() string {
	return fmt.Sprintf("[GET /products][%d] listProductsOK  %+v", 200, o.Payload)
}

func (o *ListProductsOK) String() string {
	return fmt.Sprintf("[GET /products][%d] listProductsOK  %+v", 200, o.Payload)
}

func (o *ListProductsOK) GetPayload() []*models.Product {
	return o.Payload
}

func (o *ListProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProductsInternalServerError creates a ListProductsInternalServerError with default headers values
func NewListProductsInternalServerError() *ListProductsInternalServerError {
	return &ListProductsInternalServerError{}
}

/*
ListProductsInternalServerError describes a response with status code 500, with default header values.

Internal Server error when something went wrong at server end
*/
type ListProductsInternalServerError struct {
	Payload *models.GenericError
}

// IsSuccess returns true when this list products internal server error response has a 2xx status code
func (o *ListProductsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list products internal server error response has a 3xx status code
func (o *ListProductsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list products internal server error response has a 4xx status code
func (o *ListProductsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list products internal server error response has a 5xx status code
func (o *ListProductsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list products internal server error response a status code equal to that given
func (o *ListProductsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list products internal server error response
func (o *ListProductsInternalServerError) Code() int {
	return 500
}

func (o *ListProductsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /products][%d] listProductsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProductsInternalServerError) String() string {
	return fmt.Sprintf("[GET /products][%d] listProductsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProductsInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *ListProductsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}