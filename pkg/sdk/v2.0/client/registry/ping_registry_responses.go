// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// PingRegistryReader is a Reader for the PingRegistry structure.
type PingRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPingRegistryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPingRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPingRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPingRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPingRegistryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPingRegistryOK creates a PingRegistryOK with default headers values
func NewPingRegistryOK() *PingRegistryOK {
	return &PingRegistryOK{}
}

/*
PingRegistryOK describes a response with status code 200, with default header values.

Success
*/
type PingRegistryOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this ping registry o k response has a 2xx status code
func (o *PingRegistryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ping registry o k response has a 3xx status code
func (o *PingRegistryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping registry o k response has a 4xx status code
func (o *PingRegistryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ping registry o k response has a 5xx status code
func (o *PingRegistryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ping registry o k response a status code equal to that given
func (o *PingRegistryOK) IsCode(code int) bool {
	return code == 200
}

func (o *PingRegistryOK) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryOK ", 200)
}

func (o *PingRegistryOK) String() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryOK ", 200)
}

func (o *PingRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewPingRegistryBadRequest creates a PingRegistryBadRequest with default headers values
func NewPingRegistryBadRequest() *PingRegistryBadRequest {
	return &PingRegistryBadRequest{}
}

/*
PingRegistryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PingRegistryBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this ping registry bad request response has a 2xx status code
func (o *PingRegistryBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping registry bad request response has a 3xx status code
func (o *PingRegistryBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping registry bad request response has a 4xx status code
func (o *PingRegistryBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping registry bad request response has a 5xx status code
func (o *PingRegistryBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ping registry bad request response a status code equal to that given
func (o *PingRegistryBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PingRegistryBadRequest) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryBadRequest  %+v", 400, o.Payload)
}

func (o *PingRegistryBadRequest) String() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryBadRequest  %+v", 400, o.Payload)
}

func (o *PingRegistryBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingRegistryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPingRegistryUnauthorized creates a PingRegistryUnauthorized with default headers values
func NewPingRegistryUnauthorized() *PingRegistryUnauthorized {
	return &PingRegistryUnauthorized{}
}

/*
PingRegistryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PingRegistryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this ping registry unauthorized response has a 2xx status code
func (o *PingRegistryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping registry unauthorized response has a 3xx status code
func (o *PingRegistryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping registry unauthorized response has a 4xx status code
func (o *PingRegistryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping registry unauthorized response has a 5xx status code
func (o *PingRegistryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ping registry unauthorized response a status code equal to that given
func (o *PingRegistryUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PingRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryUnauthorized  %+v", 401, o.Payload)
}

func (o *PingRegistryUnauthorized) String() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryUnauthorized  %+v", 401, o.Payload)
}

func (o *PingRegistryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPingRegistryForbidden creates a PingRegistryForbidden with default headers values
func NewPingRegistryForbidden() *PingRegistryForbidden {
	return &PingRegistryForbidden{}
}

/*
PingRegistryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PingRegistryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this ping registry forbidden response has a 2xx status code
func (o *PingRegistryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping registry forbidden response has a 3xx status code
func (o *PingRegistryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping registry forbidden response has a 4xx status code
func (o *PingRegistryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping registry forbidden response has a 5xx status code
func (o *PingRegistryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this ping registry forbidden response a status code equal to that given
func (o *PingRegistryForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PingRegistryForbidden) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryForbidden  %+v", 403, o.Payload)
}

func (o *PingRegistryForbidden) String() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryForbidden  %+v", 403, o.Payload)
}

func (o *PingRegistryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPingRegistryNotFound creates a PingRegistryNotFound with default headers values
func NewPingRegistryNotFound() *PingRegistryNotFound {
	return &PingRegistryNotFound{}
}

/*
PingRegistryNotFound describes a response with status code 404, with default header values.

Not found
*/
type PingRegistryNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this ping registry not found response has a 2xx status code
func (o *PingRegistryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping registry not found response has a 3xx status code
func (o *PingRegistryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping registry not found response has a 4xx status code
func (o *PingRegistryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping registry not found response has a 5xx status code
func (o *PingRegistryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ping registry not found response a status code equal to that given
func (o *PingRegistryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PingRegistryNotFound) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryNotFound  %+v", 404, o.Payload)
}

func (o *PingRegistryNotFound) String() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryNotFound  %+v", 404, o.Payload)
}

func (o *PingRegistryNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewPingRegistryInternalServerError creates a PingRegistryInternalServerError with default headers values
func NewPingRegistryInternalServerError() *PingRegistryInternalServerError {
	return &PingRegistryInternalServerError{}
}

/*
PingRegistryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PingRegistryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this ping registry internal server error response has a 2xx status code
func (o *PingRegistryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping registry internal server error response has a 3xx status code
func (o *PingRegistryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping registry internal server error response has a 4xx status code
func (o *PingRegistryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ping registry internal server error response has a 5xx status code
func (o *PingRegistryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ping registry internal server error response a status code equal to that given
func (o *PingRegistryInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PingRegistryInternalServerError) Error() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryInternalServerError  %+v", 500, o.Payload)
}

func (o *PingRegistryInternalServerError) String() string {
	return fmt.Sprintf("[POST /registries/ping][%d] pingRegistryInternalServerError  %+v", 500, o.Payload)
}

func (o *PingRegistryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingRegistryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
