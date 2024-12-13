// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// CopyArtifactReader is a Reader for the CopyArtifact structure.
type CopyArtifactReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CopyArtifactReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCopyArtifactCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCopyArtifactBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCopyArtifactUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCopyArtifactForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCopyArtifactNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewCopyArtifactMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCopyArtifactInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCopyArtifactCreated creates a CopyArtifactCreated with default headers values
func NewCopyArtifactCreated() *CopyArtifactCreated {
	return &CopyArtifactCreated{}
}

/*
CopyArtifactCreated describes a response with status code 201, with default header values.

Created
*/
type CopyArtifactCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this copy artifact created response has a 2xx status code
func (o *CopyArtifactCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this copy artifact created response has a 3xx status code
func (o *CopyArtifactCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy artifact created response has a 4xx status code
func (o *CopyArtifactCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this copy artifact created response has a 5xx status code
func (o *CopyArtifactCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this copy artifact created response a status code equal to that given
func (o *CopyArtifactCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CopyArtifactCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactCreated ", 201)
}

func (o *CopyArtifactCreated) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactCreated ", 201)
}

func (o *CopyArtifactCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCopyArtifactBadRequest creates a CopyArtifactBadRequest with default headers values
func NewCopyArtifactBadRequest() *CopyArtifactBadRequest {
	return &CopyArtifactBadRequest{}
}

/*
CopyArtifactBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CopyArtifactBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this copy artifact bad request response has a 2xx status code
func (o *CopyArtifactBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy artifact bad request response has a 3xx status code
func (o *CopyArtifactBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy artifact bad request response has a 4xx status code
func (o *CopyArtifactBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy artifact bad request response has a 5xx status code
func (o *CopyArtifactBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this copy artifact bad request response a status code equal to that given
func (o *CopyArtifactBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CopyArtifactBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactBadRequest  %+v", 400, o.Payload)
}

func (o *CopyArtifactBadRequest) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactBadRequest  %+v", 400, o.Payload)
}

func (o *CopyArtifactBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CopyArtifactBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCopyArtifactUnauthorized creates a CopyArtifactUnauthorized with default headers values
func NewCopyArtifactUnauthorized() *CopyArtifactUnauthorized {
	return &CopyArtifactUnauthorized{}
}

/*
CopyArtifactUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CopyArtifactUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this copy artifact unauthorized response has a 2xx status code
func (o *CopyArtifactUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy artifact unauthorized response has a 3xx status code
func (o *CopyArtifactUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy artifact unauthorized response has a 4xx status code
func (o *CopyArtifactUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy artifact unauthorized response has a 5xx status code
func (o *CopyArtifactUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this copy artifact unauthorized response a status code equal to that given
func (o *CopyArtifactUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CopyArtifactUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *CopyArtifactUnauthorized) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactUnauthorized  %+v", 401, o.Payload)
}

func (o *CopyArtifactUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CopyArtifactUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCopyArtifactForbidden creates a CopyArtifactForbidden with default headers values
func NewCopyArtifactForbidden() *CopyArtifactForbidden {
	return &CopyArtifactForbidden{}
}

/*
CopyArtifactForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CopyArtifactForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this copy artifact forbidden response has a 2xx status code
func (o *CopyArtifactForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy artifact forbidden response has a 3xx status code
func (o *CopyArtifactForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy artifact forbidden response has a 4xx status code
func (o *CopyArtifactForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy artifact forbidden response has a 5xx status code
func (o *CopyArtifactForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this copy artifact forbidden response a status code equal to that given
func (o *CopyArtifactForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CopyArtifactForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactForbidden  %+v", 403, o.Payload)
}

func (o *CopyArtifactForbidden) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactForbidden  %+v", 403, o.Payload)
}

func (o *CopyArtifactForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CopyArtifactForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCopyArtifactNotFound creates a CopyArtifactNotFound with default headers values
func NewCopyArtifactNotFound() *CopyArtifactNotFound {
	return &CopyArtifactNotFound{}
}

/*
CopyArtifactNotFound describes a response with status code 404, with default header values.

Not found
*/
type CopyArtifactNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this copy artifact not found response has a 2xx status code
func (o *CopyArtifactNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy artifact not found response has a 3xx status code
func (o *CopyArtifactNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy artifact not found response has a 4xx status code
func (o *CopyArtifactNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy artifact not found response has a 5xx status code
func (o *CopyArtifactNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this copy artifact not found response a status code equal to that given
func (o *CopyArtifactNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CopyArtifactNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactNotFound  %+v", 404, o.Payload)
}

func (o *CopyArtifactNotFound) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactNotFound  %+v", 404, o.Payload)
}

func (o *CopyArtifactNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CopyArtifactNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCopyArtifactMethodNotAllowed creates a CopyArtifactMethodNotAllowed with default headers values
func NewCopyArtifactMethodNotAllowed() *CopyArtifactMethodNotAllowed {
	return &CopyArtifactMethodNotAllowed{}
}

/*
CopyArtifactMethodNotAllowed describes a response with status code 405, with default header values.

Method not allowed
*/
type CopyArtifactMethodNotAllowed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this copy artifact method not allowed response has a 2xx status code
func (o *CopyArtifactMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy artifact method not allowed response has a 3xx status code
func (o *CopyArtifactMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy artifact method not allowed response has a 4xx status code
func (o *CopyArtifactMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this copy artifact method not allowed response has a 5xx status code
func (o *CopyArtifactMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this copy artifact method not allowed response a status code equal to that given
func (o *CopyArtifactMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *CopyArtifactMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CopyArtifactMethodNotAllowed) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *CopyArtifactMethodNotAllowed) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CopyArtifactMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCopyArtifactInternalServerError creates a CopyArtifactInternalServerError with default headers values
func NewCopyArtifactInternalServerError() *CopyArtifactInternalServerError {
	return &CopyArtifactInternalServerError{}
}

/*
CopyArtifactInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CopyArtifactInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this copy artifact internal server error response has a 2xx status code
func (o *CopyArtifactInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this copy artifact internal server error response has a 3xx status code
func (o *CopyArtifactInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this copy artifact internal server error response has a 4xx status code
func (o *CopyArtifactInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this copy artifact internal server error response has a 5xx status code
func (o *CopyArtifactInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this copy artifact internal server error response a status code equal to that given
func (o *CopyArtifactInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CopyArtifactInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *CopyArtifactInternalServerError) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/repositories/{repository_name}/artifacts][%d] copyArtifactInternalServerError  %+v", 500, o.Payload)
}

func (o *CopyArtifactInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CopyArtifactInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
