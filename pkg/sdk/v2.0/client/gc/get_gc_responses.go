// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetGCReader is a Reader for the GetGC structure.
type GetGCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGCUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGCForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGCNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGCInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGCOK creates a GetGCOK with default headers values
func NewGetGCOK() *GetGCOK {
	return &GetGCOK{}
}

/*
GetGCOK describes a response with status code 200, with default header values.

Get gc results successfully.
*/
type GetGCOK struct {
	Payload *models.GCHistory
}

// IsSuccess returns true when this get Gc o k response has a 2xx status code
func (o *GetGCOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Gc o k response has a 3xx status code
func (o *GetGCOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc o k response has a 4xx status code
func (o *GetGCOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Gc o k response has a 5xx status code
func (o *GetGCOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc o k response a status code equal to that given
func (o *GetGCOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGCOK) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcOK  %+v", 200, o.Payload)
}

func (o *GetGCOK) String() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcOK  %+v", 200, o.Payload)
}

func (o *GetGCOK) GetPayload() *models.GCHistory {
	return o.Payload
}

func (o *GetGCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GCHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCUnauthorized creates a GetGCUnauthorized with default headers values
func NewGetGCUnauthorized() *GetGCUnauthorized {
	return &GetGCUnauthorized{}
}

/*
GetGCUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetGCUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc unauthorized response has a 2xx status code
func (o *GetGCUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc unauthorized response has a 3xx status code
func (o *GetGCUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc unauthorized response has a 4xx status code
func (o *GetGCUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Gc unauthorized response has a 5xx status code
func (o *GetGCUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc unauthorized response a status code equal to that given
func (o *GetGCUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGCUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGCUnauthorized) String() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGCUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCForbidden creates a GetGCForbidden with default headers values
func NewGetGCForbidden() *GetGCForbidden {
	return &GetGCForbidden{}
}

/*
GetGCForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGCForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc forbidden response has a 2xx status code
func (o *GetGCForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc forbidden response has a 3xx status code
func (o *GetGCForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc forbidden response has a 4xx status code
func (o *GetGCForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Gc forbidden response has a 5xx status code
func (o *GetGCForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc forbidden response a status code equal to that given
func (o *GetGCForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGCForbidden) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcForbidden  %+v", 403, o.Payload)
}

func (o *GetGCForbidden) String() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcForbidden  %+v", 403, o.Payload)
}

func (o *GetGCForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCNotFound creates a GetGCNotFound with default headers values
func NewGetGCNotFound() *GetGCNotFound {
	return &GetGCNotFound{}
}

/*
GetGCNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetGCNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc not found response has a 2xx status code
func (o *GetGCNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc not found response has a 3xx status code
func (o *GetGCNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc not found response has a 4xx status code
func (o *GetGCNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Gc not found response has a 5xx status code
func (o *GetGCNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc not found response a status code equal to that given
func (o *GetGCNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGCNotFound) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcNotFound  %+v", 404, o.Payload)
}

func (o *GetGCNotFound) String() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcNotFound  %+v", 404, o.Payload)
}

func (o *GetGCNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCInternalServerError creates a GetGCInternalServerError with default headers values
func NewGetGCInternalServerError() *GetGCInternalServerError {
	return &GetGCInternalServerError{}
}

/*
GetGCInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetGCInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc internal server error response has a 2xx status code
func (o *GetGCInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc internal server error response has a 3xx status code
func (o *GetGCInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc internal server error response has a 4xx status code
func (o *GetGCInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Gc internal server error response has a 5xx status code
func (o *GetGCInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Gc internal server error response a status code equal to that given
func (o *GetGCInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGCInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGCInternalServerError) String() string {
	return fmt.Sprintf("[GET /system/gc/{gc_id}][%d] getGcInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGCInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
