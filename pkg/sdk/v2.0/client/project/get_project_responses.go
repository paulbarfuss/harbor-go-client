// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetProjectReader is a Reader for the GetProject structure.
type GetProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectOK creates a GetProjectOK with default headers values
func NewGetProjectOK() *GetProjectOK {
	return &GetProjectOK{}
}

/*
GetProjectOK describes a response with status code 200, with default header values.

Return matched project information.
*/
type GetProjectOK struct {
	Payload *models.Project
}

// IsSuccess returns true when this get project o k response has a 2xx status code
func (o *GetProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project o k response has a 3xx status code
func (o *GetProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project o k response has a 4xx status code
func (o *GetProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project o k response has a 5xx status code
func (o *GetProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project o k response a status code equal to that given
func (o *GetProjectOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}][%d] getProjectOK  %+v", 200, o.Payload)
}

func (o *GetProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *GetProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectUnauthorized creates a GetProjectUnauthorized with default headers values
func NewGetProjectUnauthorized() *GetProjectUnauthorized {
	return &GetProjectUnauthorized{}
}

/*
GetProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get project unauthorized response has a 2xx status code
func (o *GetProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project unauthorized response has a 3xx status code
func (o *GetProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project unauthorized response has a 4xx status code
func (o *GetProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project unauthorized response has a 5xx status code
func (o *GetProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get project unauthorized response a status code equal to that given
func (o *GetProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}][%d] getProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProjectUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}][%d] getProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProjectUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectInternalServerError creates a GetProjectInternalServerError with default headers values
func NewGetProjectInternalServerError() *GetProjectInternalServerError {
	return &GetProjectInternalServerError{}
}

/*
GetProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetProjectInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get project internal server error response has a 2xx status code
func (o *GetProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project internal server error response has a 3xx status code
func (o *GetProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project internal server error response has a 4xx status code
func (o *GetProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project internal server error response has a 5xx status code
func (o *GetProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get project internal server error response a status code equal to that given
func (o *GetProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}][%d] getProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}][%d] getProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
