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

// GetPreheatLogReader is a Reader for the GetPreheatLog structure.
type GetPreheatLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPreheatLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPreheatLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPreheatLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPreheatLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPreheatLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPreheatLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPreheatLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPreheatLogOK creates a GetPreheatLogOK with default headers values
func NewGetPreheatLogOK() *GetPreheatLogOK {
	return &GetPreheatLogOK{}
}

/*
GetPreheatLogOK describes a response with status code 200, with default header values.

Get log success
*/
type GetPreheatLogOK struct {

	/* Content type of response
	 */
	ContentType string

	Payload string
}

// IsSuccess returns true when this get preheat log o k response has a 2xx status code
func (o *GetPreheatLogOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get preheat log o k response has a 3xx status code
func (o *GetPreheatLogOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get preheat log o k response has a 4xx status code
func (o *GetPreheatLogOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get preheat log o k response has a 5xx status code
func (o *GetPreheatLogOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get preheat log o k response a status code equal to that given
func (o *GetPreheatLogOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPreheatLogOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogOK  %+v", 200, o.Payload)
}

func (o *GetPreheatLogOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogOK  %+v", 200, o.Payload)
}

func (o *GetPreheatLogOK) GetPayload() string {
	return o.Payload
}

func (o *GetPreheatLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPreheatLogBadRequest creates a GetPreheatLogBadRequest with default headers values
func NewGetPreheatLogBadRequest() *GetPreheatLogBadRequest {
	return &GetPreheatLogBadRequest{}
}

/*
GetPreheatLogBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetPreheatLogBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get preheat log bad request response has a 2xx status code
func (o *GetPreheatLogBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get preheat log bad request response has a 3xx status code
func (o *GetPreheatLogBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get preheat log bad request response has a 4xx status code
func (o *GetPreheatLogBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get preheat log bad request response has a 5xx status code
func (o *GetPreheatLogBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get preheat log bad request response a status code equal to that given
func (o *GetPreheatLogBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPreheatLogBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogBadRequest  %+v", 400, o.Payload)
}

func (o *GetPreheatLogBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogBadRequest  %+v", 400, o.Payload)
}

func (o *GetPreheatLogBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetPreheatLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPreheatLogUnauthorized creates a GetPreheatLogUnauthorized with default headers values
func NewGetPreheatLogUnauthorized() *GetPreheatLogUnauthorized {
	return &GetPreheatLogUnauthorized{}
}

/*
GetPreheatLogUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPreheatLogUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get preheat log unauthorized response has a 2xx status code
func (o *GetPreheatLogUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get preheat log unauthorized response has a 3xx status code
func (o *GetPreheatLogUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get preheat log unauthorized response has a 4xx status code
func (o *GetPreheatLogUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get preheat log unauthorized response has a 5xx status code
func (o *GetPreheatLogUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get preheat log unauthorized response a status code equal to that given
func (o *GetPreheatLogUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetPreheatLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPreheatLogUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPreheatLogUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetPreheatLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPreheatLogForbidden creates a GetPreheatLogForbidden with default headers values
func NewGetPreheatLogForbidden() *GetPreheatLogForbidden {
	return &GetPreheatLogForbidden{}
}

/*
GetPreheatLogForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPreheatLogForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get preheat log forbidden response has a 2xx status code
func (o *GetPreheatLogForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get preheat log forbidden response has a 3xx status code
func (o *GetPreheatLogForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get preheat log forbidden response has a 4xx status code
func (o *GetPreheatLogForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get preheat log forbidden response has a 5xx status code
func (o *GetPreheatLogForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get preheat log forbidden response a status code equal to that given
func (o *GetPreheatLogForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetPreheatLogForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogForbidden  %+v", 403, o.Payload)
}

func (o *GetPreheatLogForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogForbidden  %+v", 403, o.Payload)
}

func (o *GetPreheatLogForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetPreheatLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPreheatLogNotFound creates a GetPreheatLogNotFound with default headers values
func NewGetPreheatLogNotFound() *GetPreheatLogNotFound {
	return &GetPreheatLogNotFound{}
}

/*
GetPreheatLogNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetPreheatLogNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get preheat log not found response has a 2xx status code
func (o *GetPreheatLogNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get preheat log not found response has a 3xx status code
func (o *GetPreheatLogNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get preheat log not found response has a 4xx status code
func (o *GetPreheatLogNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get preheat log not found response has a 5xx status code
func (o *GetPreheatLogNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get preheat log not found response a status code equal to that given
func (o *GetPreheatLogNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPreheatLogNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogNotFound  %+v", 404, o.Payload)
}

func (o *GetPreheatLogNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogNotFound  %+v", 404, o.Payload)
}

func (o *GetPreheatLogNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetPreheatLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPreheatLogInternalServerError creates a GetPreheatLogInternalServerError with default headers values
func NewGetPreheatLogInternalServerError() *GetPreheatLogInternalServerError {
	return &GetPreheatLogInternalServerError{}
}

/*
GetPreheatLogInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetPreheatLogInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get preheat log internal server error response has a 2xx status code
func (o *GetPreheatLogInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get preheat log internal server error response has a 3xx status code
func (o *GetPreheatLogInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get preheat log internal server error response has a 4xx status code
func (o *GetPreheatLogInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get preheat log internal server error response has a 5xx status code
func (o *GetPreheatLogInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get preheat log internal server error response a status code equal to that given
func (o *GetPreheatLogInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetPreheatLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPreheatLogInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}/tasks/{task_id}/logs][%d] getPreheatLogInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPreheatLogInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetPreheatLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
