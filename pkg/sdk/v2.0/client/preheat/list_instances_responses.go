// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// ListInstancesReader is a Reader for the ListInstances structure.
type ListInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListInstancesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListInstancesOK creates a ListInstancesOK with default headers values
func NewListInstancesOK() *ListInstancesOK {
	return &ListInstancesOK{}
}

/*
ListInstancesOK describes a response with status code 200, with default header values.

Success
*/
type ListInstancesOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of preheating provider instances
	 */
	XTotalCount int64

	Payload []*models.Instance
}

// IsSuccess returns true when this list instances o k response has a 2xx status code
func (o *ListInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list instances o k response has a 3xx status code
func (o *ListInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list instances o k response has a 4xx status code
func (o *ListInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list instances o k response has a 5xx status code
func (o *ListInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list instances o k response a status code equal to that given
func (o *ListInstancesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ListInstancesOK) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesOK  %+v", 200, o.Payload)
}

func (o *ListInstancesOK) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesOK  %+v", 200, o.Payload)
}

func (o *ListInstancesOK) GetPayload() []*models.Instance {
	return o.Payload
}

func (o *ListInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListInstancesBadRequest creates a ListInstancesBadRequest with default headers values
func NewListInstancesBadRequest() *ListInstancesBadRequest {
	return &ListInstancesBadRequest{}
}

/*
ListInstancesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListInstancesBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list instances bad request response has a 2xx status code
func (o *ListInstancesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list instances bad request response has a 3xx status code
func (o *ListInstancesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list instances bad request response has a 4xx status code
func (o *ListInstancesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list instances bad request response has a 5xx status code
func (o *ListInstancesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list instances bad request response a status code equal to that given
func (o *ListInstancesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ListInstancesBadRequest) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *ListInstancesBadRequest) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *ListInstancesBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListInstancesUnauthorized creates a ListInstancesUnauthorized with default headers values
func NewListInstancesUnauthorized() *ListInstancesUnauthorized {
	return &ListInstancesUnauthorized{}
}

/*
ListInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListInstancesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list instances unauthorized response has a 2xx status code
func (o *ListInstancesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list instances unauthorized response has a 3xx status code
func (o *ListInstancesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list instances unauthorized response has a 4xx status code
func (o *ListInstancesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list instances unauthorized response has a 5xx status code
func (o *ListInstancesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list instances unauthorized response a status code equal to that given
func (o *ListInstancesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *ListInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListInstancesUnauthorized) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListInstancesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListInstancesForbidden creates a ListInstancesForbidden with default headers values
func NewListInstancesForbidden() *ListInstancesForbidden {
	return &ListInstancesForbidden{}
}

/*
ListInstancesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListInstancesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list instances forbidden response has a 2xx status code
func (o *ListInstancesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list instances forbidden response has a 3xx status code
func (o *ListInstancesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list instances forbidden response has a 4xx status code
func (o *ListInstancesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list instances forbidden response has a 5xx status code
func (o *ListInstancesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list instances forbidden response a status code equal to that given
func (o *ListInstancesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *ListInstancesForbidden) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesForbidden  %+v", 403, o.Payload)
}

func (o *ListInstancesForbidden) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesForbidden  %+v", 403, o.Payload)
}

func (o *ListInstancesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListInstancesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListInstancesNotFound creates a ListInstancesNotFound with default headers values
func NewListInstancesNotFound() *ListInstancesNotFound {
	return &ListInstancesNotFound{}
}

/*
ListInstancesNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListInstancesNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list instances not found response has a 2xx status code
func (o *ListInstancesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list instances not found response has a 3xx status code
func (o *ListInstancesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list instances not found response has a 4xx status code
func (o *ListInstancesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list instances not found response has a 5xx status code
func (o *ListInstancesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list instances not found response a status code equal to that given
func (o *ListInstancesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *ListInstancesNotFound) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesNotFound  %+v", 404, o.Payload)
}

func (o *ListInstancesNotFound) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesNotFound  %+v", 404, o.Payload)
}

func (o *ListInstancesNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListInstancesInternalServerError creates a ListInstancesInternalServerError with default headers values
func NewListInstancesInternalServerError() *ListInstancesInternalServerError {
	return &ListInstancesInternalServerError{}
}

/*
ListInstancesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListInstancesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list instances internal server error response has a 2xx status code
func (o *ListInstancesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list instances internal server error response has a 3xx status code
func (o *ListInstancesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list instances internal server error response has a 4xx status code
func (o *ListInstancesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list instances internal server error response has a 5xx status code
func (o *ListInstancesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list instances internal server error response a status code equal to that given
func (o *ListInstancesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *ListInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListInstancesInternalServerError) String() string {
	return fmt.Sprintf("[GET /p2p/preheat/instances][%d] listInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListInstancesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
