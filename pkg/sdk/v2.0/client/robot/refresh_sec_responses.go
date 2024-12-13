// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// RefreshSecReader is a Reader for the RefreshSec structure.
type RefreshSecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshSecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRefreshSecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRefreshSecBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRefreshSecUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRefreshSecForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRefreshSecNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRefreshSecInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRefreshSecOK creates a RefreshSecOK with default headers values
func NewRefreshSecOK() *RefreshSecOK {
	return &RefreshSecOK{}
}

/*
RefreshSecOK describes a response with status code 200, with default header values.

Return refreshed robot sec.
*/
type RefreshSecOK struct {
	Payload *models.RobotSec
}

// IsSuccess returns true when this refresh sec o k response has a 2xx status code
func (o *RefreshSecOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this refresh sec o k response has a 3xx status code
func (o *RefreshSecOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh sec o k response has a 4xx status code
func (o *RefreshSecOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh sec o k response has a 5xx status code
func (o *RefreshSecOK) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh sec o k response a status code equal to that given
func (o *RefreshSecOK) IsCode(code int) bool {
	return code == 200
}

func (o *RefreshSecOK) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecOK  %+v", 200, o.Payload)
}

func (o *RefreshSecOK) String() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecOK  %+v", 200, o.Payload)
}

func (o *RefreshSecOK) GetPayload() *models.RobotSec {
	return o.Payload
}

func (o *RefreshSecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RobotSec)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefreshSecBadRequest creates a RefreshSecBadRequest with default headers values
func NewRefreshSecBadRequest() *RefreshSecBadRequest {
	return &RefreshSecBadRequest{}
}

/*
RefreshSecBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type RefreshSecBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this refresh sec bad request response has a 2xx status code
func (o *RefreshSecBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh sec bad request response has a 3xx status code
func (o *RefreshSecBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh sec bad request response has a 4xx status code
func (o *RefreshSecBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this refresh sec bad request response has a 5xx status code
func (o *RefreshSecBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh sec bad request response a status code equal to that given
func (o *RefreshSecBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *RefreshSecBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecBadRequest  %+v", 400, o.Payload)
}

func (o *RefreshSecBadRequest) String() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecBadRequest  %+v", 400, o.Payload)
}

func (o *RefreshSecBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RefreshSecBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRefreshSecUnauthorized creates a RefreshSecUnauthorized with default headers values
func NewRefreshSecUnauthorized() *RefreshSecUnauthorized {
	return &RefreshSecUnauthorized{}
}

/*
RefreshSecUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RefreshSecUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this refresh sec unauthorized response has a 2xx status code
func (o *RefreshSecUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh sec unauthorized response has a 3xx status code
func (o *RefreshSecUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh sec unauthorized response has a 4xx status code
func (o *RefreshSecUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this refresh sec unauthorized response has a 5xx status code
func (o *RefreshSecUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh sec unauthorized response a status code equal to that given
func (o *RefreshSecUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *RefreshSecUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecUnauthorized  %+v", 401, o.Payload)
}

func (o *RefreshSecUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecUnauthorized  %+v", 401, o.Payload)
}

func (o *RefreshSecUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RefreshSecUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRefreshSecForbidden creates a RefreshSecForbidden with default headers values
func NewRefreshSecForbidden() *RefreshSecForbidden {
	return &RefreshSecForbidden{}
}

/*
RefreshSecForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RefreshSecForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this refresh sec forbidden response has a 2xx status code
func (o *RefreshSecForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh sec forbidden response has a 3xx status code
func (o *RefreshSecForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh sec forbidden response has a 4xx status code
func (o *RefreshSecForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this refresh sec forbidden response has a 5xx status code
func (o *RefreshSecForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh sec forbidden response a status code equal to that given
func (o *RefreshSecForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *RefreshSecForbidden) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecForbidden  %+v", 403, o.Payload)
}

func (o *RefreshSecForbidden) String() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecForbidden  %+v", 403, o.Payload)
}

func (o *RefreshSecForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RefreshSecForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRefreshSecNotFound creates a RefreshSecNotFound with default headers values
func NewRefreshSecNotFound() *RefreshSecNotFound {
	return &RefreshSecNotFound{}
}

/*
RefreshSecNotFound describes a response with status code 404, with default header values.

Not found
*/
type RefreshSecNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this refresh sec not found response has a 2xx status code
func (o *RefreshSecNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh sec not found response has a 3xx status code
func (o *RefreshSecNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh sec not found response has a 4xx status code
func (o *RefreshSecNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this refresh sec not found response has a 5xx status code
func (o *RefreshSecNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this refresh sec not found response a status code equal to that given
func (o *RefreshSecNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *RefreshSecNotFound) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecNotFound  %+v", 404, o.Payload)
}

func (o *RefreshSecNotFound) String() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecNotFound  %+v", 404, o.Payload)
}

func (o *RefreshSecNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RefreshSecNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRefreshSecInternalServerError creates a RefreshSecInternalServerError with default headers values
func NewRefreshSecInternalServerError() *RefreshSecInternalServerError {
	return &RefreshSecInternalServerError{}
}

/*
RefreshSecInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type RefreshSecInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this refresh sec internal server error response has a 2xx status code
func (o *RefreshSecInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this refresh sec internal server error response has a 3xx status code
func (o *RefreshSecInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this refresh sec internal server error response has a 4xx status code
func (o *RefreshSecInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this refresh sec internal server error response has a 5xx status code
func (o *RefreshSecInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this refresh sec internal server error response a status code equal to that given
func (o *RefreshSecInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *RefreshSecInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecInternalServerError  %+v", 500, o.Payload)
}

func (o *RefreshSecInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /robots/{robot_id}][%d] refreshSecInternalServerError  %+v", 500, o.Payload)
}

func (o *RefreshSecInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *RefreshSecInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
