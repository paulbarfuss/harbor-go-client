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

// CreateInstanceReader is a Reader for the CreateInstance structure.
type CreateInstanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInstanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateInstanceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateInstanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateInstanceUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateInstanceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateInstanceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateInstanceConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateInstanceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateInstanceCreated creates a CreateInstanceCreated with default headers values
func NewCreateInstanceCreated() *CreateInstanceCreated {
	return &CreateInstanceCreated{}
}

/*
CreateInstanceCreated describes a response with status code 201, with default header values.

Created
*/
type CreateInstanceCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this create instance created response has a 2xx status code
func (o *CreateInstanceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create instance created response has a 3xx status code
func (o *CreateInstanceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create instance created response has a 4xx status code
func (o *CreateInstanceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create instance created response has a 5xx status code
func (o *CreateInstanceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create instance created response a status code equal to that given
func (o *CreateInstanceCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateInstanceCreated) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceCreated ", 201)
}

func (o *CreateInstanceCreated) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceCreated ", 201)
}

func (o *CreateInstanceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewCreateInstanceBadRequest creates a CreateInstanceBadRequest with default headers values
func NewCreateInstanceBadRequest() *CreateInstanceBadRequest {
	return &CreateInstanceBadRequest{}
}

/*
CreateInstanceBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateInstanceBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create instance bad request response has a 2xx status code
func (o *CreateInstanceBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create instance bad request response has a 3xx status code
func (o *CreateInstanceBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create instance bad request response has a 4xx status code
func (o *CreateInstanceBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create instance bad request response has a 5xx status code
func (o *CreateInstanceBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create instance bad request response a status code equal to that given
func (o *CreateInstanceBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateInstanceBadRequest) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInstanceBadRequest) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceBadRequest  %+v", 400, o.Payload)
}

func (o *CreateInstanceBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateInstanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateInstanceUnauthorized creates a CreateInstanceUnauthorized with default headers values
func NewCreateInstanceUnauthorized() *CreateInstanceUnauthorized {
	return &CreateInstanceUnauthorized{}
}

/*
CreateInstanceUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateInstanceUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create instance unauthorized response has a 2xx status code
func (o *CreateInstanceUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create instance unauthorized response has a 3xx status code
func (o *CreateInstanceUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create instance unauthorized response has a 4xx status code
func (o *CreateInstanceUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create instance unauthorized response has a 5xx status code
func (o *CreateInstanceUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create instance unauthorized response a status code equal to that given
func (o *CreateInstanceUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateInstanceUnauthorized) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateInstanceUnauthorized) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateInstanceUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateInstanceUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateInstanceForbidden creates a CreateInstanceForbidden with default headers values
func NewCreateInstanceForbidden() *CreateInstanceForbidden {
	return &CreateInstanceForbidden{}
}

/*
CreateInstanceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateInstanceForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create instance forbidden response has a 2xx status code
func (o *CreateInstanceForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create instance forbidden response has a 3xx status code
func (o *CreateInstanceForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create instance forbidden response has a 4xx status code
func (o *CreateInstanceForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create instance forbidden response has a 5xx status code
func (o *CreateInstanceForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create instance forbidden response a status code equal to that given
func (o *CreateInstanceForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateInstanceForbidden) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceForbidden  %+v", 403, o.Payload)
}

func (o *CreateInstanceForbidden) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceForbidden  %+v", 403, o.Payload)
}

func (o *CreateInstanceForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateInstanceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateInstanceNotFound creates a CreateInstanceNotFound with default headers values
func NewCreateInstanceNotFound() *CreateInstanceNotFound {
	return &CreateInstanceNotFound{}
}

/*
CreateInstanceNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateInstanceNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create instance not found response has a 2xx status code
func (o *CreateInstanceNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create instance not found response has a 3xx status code
func (o *CreateInstanceNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create instance not found response has a 4xx status code
func (o *CreateInstanceNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create instance not found response has a 5xx status code
func (o *CreateInstanceNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create instance not found response a status code equal to that given
func (o *CreateInstanceNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateInstanceNotFound) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceNotFound  %+v", 404, o.Payload)
}

func (o *CreateInstanceNotFound) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceNotFound  %+v", 404, o.Payload)
}

func (o *CreateInstanceNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateInstanceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateInstanceConflict creates a CreateInstanceConflict with default headers values
func NewCreateInstanceConflict() *CreateInstanceConflict {
	return &CreateInstanceConflict{}
}

/*
CreateInstanceConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreateInstanceConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create instance conflict response has a 2xx status code
func (o *CreateInstanceConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create instance conflict response has a 3xx status code
func (o *CreateInstanceConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create instance conflict response has a 4xx status code
func (o *CreateInstanceConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create instance conflict response has a 5xx status code
func (o *CreateInstanceConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create instance conflict response a status code equal to that given
func (o *CreateInstanceConflict) IsCode(code int) bool {
	return code == 409
}

func (o *CreateInstanceConflict) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceConflict  %+v", 409, o.Payload)
}

func (o *CreateInstanceConflict) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceConflict  %+v", 409, o.Payload)
}

func (o *CreateInstanceConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateInstanceConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateInstanceInternalServerError creates a CreateInstanceInternalServerError with default headers values
func NewCreateInstanceInternalServerError() *CreateInstanceInternalServerError {
	return &CreateInstanceInternalServerError{}
}

/*
CreateInstanceInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateInstanceInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create instance internal server error response has a 2xx status code
func (o *CreateInstanceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create instance internal server error response has a 3xx status code
func (o *CreateInstanceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create instance internal server error response has a 4xx status code
func (o *CreateInstanceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create instance internal server error response has a 5xx status code
func (o *CreateInstanceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create instance internal server error response a status code equal to that given
func (o *CreateInstanceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateInstanceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInstanceInternalServerError) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances][%d] createInstanceInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateInstanceInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateInstanceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
