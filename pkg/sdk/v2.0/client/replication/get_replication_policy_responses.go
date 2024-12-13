// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetReplicationPolicyReader is a Reader for the GetReplicationPolicy structure.
type GetReplicationPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReplicationPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReplicationPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetReplicationPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetReplicationPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReplicationPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReplicationPolicyOK creates a GetReplicationPolicyOK with default headers values
func NewGetReplicationPolicyOK() *GetReplicationPolicyOK {
	return &GetReplicationPolicyOK{}
}

/*
GetReplicationPolicyOK describes a response with status code 200, with default header values.

Success
*/
type GetReplicationPolicyOK struct {
	Payload *models.ReplicationPolicy
}

// IsSuccess returns true when this get replication policy o k response has a 2xx status code
func (o *GetReplicationPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get replication policy o k response has a 3xx status code
func (o *GetReplicationPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get replication policy o k response has a 4xx status code
func (o *GetReplicationPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get replication policy o k response has a 5xx status code
func (o *GetReplicationPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get replication policy o k response a status code equal to that given
func (o *GetReplicationPolicyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetReplicationPolicyOK) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyOK  %+v", 200, o.Payload)
}

func (o *GetReplicationPolicyOK) String() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyOK  %+v", 200, o.Payload)
}

func (o *GetReplicationPolicyOK) GetPayload() *models.ReplicationPolicy {
	return o.Payload
}

func (o *GetReplicationPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReplicationPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReplicationPolicyUnauthorized creates a GetReplicationPolicyUnauthorized with default headers values
func NewGetReplicationPolicyUnauthorized() *GetReplicationPolicyUnauthorized {
	return &GetReplicationPolicyUnauthorized{}
}

/*
GetReplicationPolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetReplicationPolicyUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get replication policy unauthorized response has a 2xx status code
func (o *GetReplicationPolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get replication policy unauthorized response has a 3xx status code
func (o *GetReplicationPolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get replication policy unauthorized response has a 4xx status code
func (o *GetReplicationPolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get replication policy unauthorized response has a 5xx status code
func (o *GetReplicationPolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get replication policy unauthorized response a status code equal to that given
func (o *GetReplicationPolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetReplicationPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReplicationPolicyUnauthorized) String() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *GetReplicationPolicyUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReplicationPolicyForbidden creates a GetReplicationPolicyForbidden with default headers values
func NewGetReplicationPolicyForbidden() *GetReplicationPolicyForbidden {
	return &GetReplicationPolicyForbidden{}
}

/*
GetReplicationPolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetReplicationPolicyForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get replication policy forbidden response has a 2xx status code
func (o *GetReplicationPolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get replication policy forbidden response has a 3xx status code
func (o *GetReplicationPolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get replication policy forbidden response has a 4xx status code
func (o *GetReplicationPolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get replication policy forbidden response has a 5xx status code
func (o *GetReplicationPolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get replication policy forbidden response a status code equal to that given
func (o *GetReplicationPolicyForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetReplicationPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetReplicationPolicyForbidden) String() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyForbidden  %+v", 403, o.Payload)
}

func (o *GetReplicationPolicyForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetReplicationPolicyInternalServerError creates a GetReplicationPolicyInternalServerError with default headers values
func NewGetReplicationPolicyInternalServerError() *GetReplicationPolicyInternalServerError {
	return &GetReplicationPolicyInternalServerError{}
}

/*
GetReplicationPolicyInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetReplicationPolicyInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get replication policy internal server error response has a 2xx status code
func (o *GetReplicationPolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get replication policy internal server error response has a 3xx status code
func (o *GetReplicationPolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get replication policy internal server error response has a 4xx status code
func (o *GetReplicationPolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get replication policy internal server error response has a 5xx status code
func (o *GetReplicationPolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get replication policy internal server error response a status code equal to that given
func (o *GetReplicationPolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetReplicationPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReplicationPolicyInternalServerError) String() string {
	return fmt.Sprintf("[GET /replication/policies/{id}][%d] getReplicationPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *GetReplicationPolicyInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetReplicationPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
