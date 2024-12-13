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

// CreateRobotReader is a Reader for the CreateRobot structure.
type CreateRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRobotCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRobotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRobotCreated creates a CreateRobotCreated with default headers values
func NewCreateRobotCreated() *CreateRobotCreated {
	return &CreateRobotCreated{}
}

/*
CreateRobotCreated describes a response with status code 201, with default header values.

Created
*/
type CreateRobotCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.RobotCreated
}

// IsSuccess returns true when this create robot created response has a 2xx status code
func (o *CreateRobotCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create robot created response has a 3xx status code
func (o *CreateRobotCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create robot created response has a 4xx status code
func (o *CreateRobotCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create robot created response has a 5xx status code
func (o *CreateRobotCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create robot created response a status code equal to that given
func (o *CreateRobotCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateRobotCreated) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotCreated  %+v", 201, o.Payload)
}

func (o *CreateRobotCreated) String() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotCreated  %+v", 201, o.Payload)
}

func (o *CreateRobotCreated) GetPayload() *models.RobotCreated {
	return o.Payload
}

func (o *CreateRobotCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RobotCreated)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRobotBadRequest creates a CreateRobotBadRequest with default headers values
func NewCreateRobotBadRequest() *CreateRobotBadRequest {
	return &CreateRobotBadRequest{}
}

/*
CreateRobotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateRobotBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create robot bad request response has a 2xx status code
func (o *CreateRobotBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create robot bad request response has a 3xx status code
func (o *CreateRobotBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create robot bad request response has a 4xx status code
func (o *CreateRobotBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create robot bad request response has a 5xx status code
func (o *CreateRobotBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create robot bad request response a status code equal to that given
func (o *CreateRobotBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateRobotBadRequest) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRobotBadRequest) String() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRobotBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRobotUnauthorized creates a CreateRobotUnauthorized with default headers values
func NewCreateRobotUnauthorized() *CreateRobotUnauthorized {
	return &CreateRobotUnauthorized{}
}

/*
CreateRobotUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateRobotUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create robot unauthorized response has a 2xx status code
func (o *CreateRobotUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create robot unauthorized response has a 3xx status code
func (o *CreateRobotUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create robot unauthorized response has a 4xx status code
func (o *CreateRobotUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create robot unauthorized response has a 5xx status code
func (o *CreateRobotUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create robot unauthorized response a status code equal to that given
func (o *CreateRobotUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateRobotUnauthorized) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRobotUnauthorized) String() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRobotUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRobotForbidden creates a CreateRobotForbidden with default headers values
func NewCreateRobotForbidden() *CreateRobotForbidden {
	return &CreateRobotForbidden{}
}

/*
CreateRobotForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateRobotForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create robot forbidden response has a 2xx status code
func (o *CreateRobotForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create robot forbidden response has a 3xx status code
func (o *CreateRobotForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create robot forbidden response has a 4xx status code
func (o *CreateRobotForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create robot forbidden response has a 5xx status code
func (o *CreateRobotForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create robot forbidden response a status code equal to that given
func (o *CreateRobotForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateRobotForbidden) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotForbidden  %+v", 403, o.Payload)
}

func (o *CreateRobotForbidden) String() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotForbidden  %+v", 403, o.Payload)
}

func (o *CreateRobotForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRobotNotFound creates a CreateRobotNotFound with default headers values
func NewCreateRobotNotFound() *CreateRobotNotFound {
	return &CreateRobotNotFound{}
}

/*
CreateRobotNotFound describes a response with status code 404, with default header values.

Not found
*/
type CreateRobotNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create robot not found response has a 2xx status code
func (o *CreateRobotNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create robot not found response has a 3xx status code
func (o *CreateRobotNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create robot not found response has a 4xx status code
func (o *CreateRobotNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create robot not found response has a 5xx status code
func (o *CreateRobotNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create robot not found response a status code equal to that given
func (o *CreateRobotNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateRobotNotFound) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotNotFound  %+v", 404, o.Payload)
}

func (o *CreateRobotNotFound) String() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotNotFound  %+v", 404, o.Payload)
}

func (o *CreateRobotNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateRobotInternalServerError creates a CreateRobotInternalServerError with default headers values
func NewCreateRobotInternalServerError() *CreateRobotInternalServerError {
	return &CreateRobotInternalServerError{}
}

/*
CreateRobotInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateRobotInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create robot internal server error response has a 2xx status code
func (o *CreateRobotInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create robot internal server error response has a 3xx status code
func (o *CreateRobotInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create robot internal server error response has a 4xx status code
func (o *CreateRobotInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create robot internal server error response has a 5xx status code
func (o *CreateRobotInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create robot internal server error response a status code equal to that given
func (o *CreateRobotInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateRobotInternalServerError) Error() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRobotInternalServerError) String() string {
	return fmt.Sprintf("[POST /robots][%d] createRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRobotInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRobotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
