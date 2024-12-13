// Code generated by go-swagger; DO NOT EDIT.

package jobservice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// ActionGetJobLogReader is a Reader for the ActionGetJobLog structure.
type ActionGetJobLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActionGetJobLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActionGetJobLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewActionGetJobLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewActionGetJobLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewActionGetJobLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewActionGetJobLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewActionGetJobLogOK creates a ActionGetJobLogOK with default headers values
func NewActionGetJobLogOK() *ActionGetJobLogOK {
	return &ActionGetJobLogOK{}
}

/*
ActionGetJobLogOK describes a response with status code 200, with default header values.

Get job log successfully.
*/
type ActionGetJobLogOK struct {

	/* The content type of response body
	 */
	ContentType string

	Payload string
}

// IsSuccess returns true when this action get job log o k response has a 2xx status code
func (o *ActionGetJobLogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this action get job log o k response has a 3xx status code
func (o *ActionGetJobLogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action get job log o k response has a 4xx status code
func (o *ActionGetJobLogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this action get job log o k response has a 5xx status code
func (o *ActionGetJobLogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this action get job log o k response a status code equal to that given
func (o *ActionGetJobLogOK) IsCode(code int) bool {
	return code == 200
}

func (o *ActionGetJobLogOK) Error() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogOK  %+v", 200, o.Payload)
}

func (o *ActionGetJobLogOK) String() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogOK  %+v", 200, o.Payload)
}

func (o *ActionGetJobLogOK) GetPayload() string {
	return o.Payload
}

func (o *ActionGetJobLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Type
	hdrContentType := response.GetHeader("Content-Type")

	if hdrContentType != "" {
		o.ContentType = hdrContentType
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActionGetJobLogUnauthorized creates a ActionGetJobLogUnauthorized with default headers values
func NewActionGetJobLogUnauthorized() *ActionGetJobLogUnauthorized {
	return &ActionGetJobLogUnauthorized{}
}

/*
ActionGetJobLogUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ActionGetJobLogUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this action get job log unauthorized response has a 2xx status code
func (o *ActionGetJobLogUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action get job log unauthorized response has a 3xx status code
func (o *ActionGetJobLogUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action get job log unauthorized response has a 4xx status code
func (o *ActionGetJobLogUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this action get job log unauthorized response has a 5xx status code
func (o *ActionGetJobLogUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this action get job log unauthorized response a status code equal to that given
func (o *ActionGetJobLogUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ActionGetJobLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogUnauthorized  %+v", 401, o.Payload)
}

func (o *ActionGetJobLogUnauthorized) String() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogUnauthorized  %+v", 401, o.Payload)
}

func (o *ActionGetJobLogUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ActionGetJobLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewActionGetJobLogForbidden creates a ActionGetJobLogForbidden with default headers values
func NewActionGetJobLogForbidden() *ActionGetJobLogForbidden {
	return &ActionGetJobLogForbidden{}
}

/*
ActionGetJobLogForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ActionGetJobLogForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this action get job log forbidden response has a 2xx status code
func (o *ActionGetJobLogForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action get job log forbidden response has a 3xx status code
func (o *ActionGetJobLogForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action get job log forbidden response has a 4xx status code
func (o *ActionGetJobLogForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this action get job log forbidden response has a 5xx status code
func (o *ActionGetJobLogForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this action get job log forbidden response a status code equal to that given
func (o *ActionGetJobLogForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ActionGetJobLogForbidden) Error() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogForbidden  %+v", 403, o.Payload)
}

func (o *ActionGetJobLogForbidden) String() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogForbidden  %+v", 403, o.Payload)
}

func (o *ActionGetJobLogForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ActionGetJobLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewActionGetJobLogNotFound creates a ActionGetJobLogNotFound with default headers values
func NewActionGetJobLogNotFound() *ActionGetJobLogNotFound {
	return &ActionGetJobLogNotFound{}
}

/*
ActionGetJobLogNotFound describes a response with status code 404, with default header values.

Not found
*/
type ActionGetJobLogNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this action get job log not found response has a 2xx status code
func (o *ActionGetJobLogNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action get job log not found response has a 3xx status code
func (o *ActionGetJobLogNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action get job log not found response has a 4xx status code
func (o *ActionGetJobLogNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this action get job log not found response has a 5xx status code
func (o *ActionGetJobLogNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this action get job log not found response a status code equal to that given
func (o *ActionGetJobLogNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ActionGetJobLogNotFound) Error() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogNotFound  %+v", 404, o.Payload)
}

func (o *ActionGetJobLogNotFound) String() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogNotFound  %+v", 404, o.Payload)
}

func (o *ActionGetJobLogNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ActionGetJobLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewActionGetJobLogInternalServerError creates a ActionGetJobLogInternalServerError with default headers values
func NewActionGetJobLogInternalServerError() *ActionGetJobLogInternalServerError {
	return &ActionGetJobLogInternalServerError{}
}

/*
ActionGetJobLogInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ActionGetJobLogInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this action get job log internal server error response has a 2xx status code
func (o *ActionGetJobLogInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this action get job log internal server error response has a 3xx status code
func (o *ActionGetJobLogInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this action get job log internal server error response has a 4xx status code
func (o *ActionGetJobLogInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this action get job log internal server error response has a 5xx status code
func (o *ActionGetJobLogInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this action get job log internal server error response a status code equal to that given
func (o *ActionGetJobLogInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ActionGetJobLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionGetJobLogInternalServerError) String() string {
	return fmt.Sprintf("[GET /jobservice/jobs/{job_id}/log][%d] actionGetJobLogInternalServerError  %+v", 500, o.Payload)
}

func (o *ActionGetJobLogInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ActionGetJobLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
