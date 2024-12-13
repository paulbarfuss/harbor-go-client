// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// DeleteInstanceReader is a Reader for the DeleteInstance structure.
type DeleteInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteInstanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteInstanceOK creates a DeleteInstanceOK with default headers values
func NewDeleteInstanceOK() *DeleteInstanceOK {
	return &DeleteInstanceOK{}
}

/*
DeleteInstanceOK describes a response with status code 200, with default header values.

Success
*/
type DeleteInstanceOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this delete instance o k response has a 2xx status code
func (o *DeleteInstanceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete instance o k response has a 3xx status code
func (o *DeleteInstanceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete instance o k response has a 4xx status code
func (o *DeleteInstanceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete instance o k response has a 5xx status code
func (o *DeleteInstanceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete instance o k response a status code equal to that given
func (o *DeleteInstanceOK) IsCode(code int) bool {
	return code == 200
}

func (o *DeleteInstanceOK) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceOK ", 200)
}

func (o *DeleteInstanceOK) String() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceOK ", 200)
}

func (o *DeleteInstanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteInstanceUnauthorized creates a DeleteInstanceUnauthorized with default headers values
func NewDeleteInstanceUnauthorized() *DeleteInstanceUnauthorized {
	return &DeleteInstanceUnauthorized{}
}

/*
DeleteInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteInstanceUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete instance unauthorized response has a 2xx status code
func (o *DeleteInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete instance unauthorized response has a 3xx status code
func (o *DeleteInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete instance unauthorized response has a 4xx status code
func (o *DeleteInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete instance unauthorized response has a 5xx status code
func (o *DeleteInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete instance unauthorized response a status code equal to that given
func (o *DeleteInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *DeleteInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteInstanceUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteInstanceUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteInstanceForbidden creates a DeleteInstanceForbidden with default headers values
func NewDeleteInstanceForbidden() *DeleteInstanceForbidden {
	return &DeleteInstanceForbidden{}
}

/*
DeleteInstanceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteInstanceForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete instance forbidden response has a 2xx status code
func (o *DeleteInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete instance forbidden response has a 3xx status code
func (o *DeleteInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete instance forbidden response has a 4xx status code
func (o *DeleteInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete instance forbidden response has a 5xx status code
func (o *DeleteInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete instance forbidden response a status code equal to that given
func (o *DeleteInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *DeleteInstanceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteInstanceForbidden) String() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceForbidden  %+v", 403, o.Payload)
}

func (o *DeleteInstanceForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteInstanceNotFound creates a DeleteInstanceNotFound with default headers values
func NewDeleteInstanceNotFound() *DeleteInstanceNotFound {
	return &DeleteInstanceNotFound{}
}

/*
DeleteInstanceNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteInstanceNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete instance not found response has a 2xx status code
func (o *DeleteInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete instance not found response has a 3xx status code
func (o *DeleteInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete instance not found response has a 4xx status code
func (o *DeleteInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete instance not found response has a 5xx status code
func (o *DeleteInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete instance not found response a status code equal to that given
func (o *DeleteInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeleteInstanceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteInstanceNotFound) String() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceNotFound  %+v", 404, o.Payload)
}

func (o *DeleteInstanceNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteInstanceInternalServerError creates a DeleteInstanceInternalServerError with default headers values
func NewDeleteInstanceInternalServerError() *DeleteInstanceInternalServerError {
	return &DeleteInstanceInternalServerError{}
}

/*
DeleteInstanceInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteInstanceInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete instance internal server error response has a 2xx status code
func (o *DeleteInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete instance internal server error response has a 3xx status code
func (o *DeleteInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete instance internal server error response has a 4xx status code
func (o *DeleteInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete instance internal server error response has a 5xx status code
func (o *DeleteInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete instance internal server error response a status code equal to that given
func (o *DeleteInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *DeleteInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteInstanceInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /p2p/preheat/instances/{preheat_instance_name}][%d] deleteInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteInstanceInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
