// Code generated by go-swagger; DO NOT EDIT.

package icon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetIconReader is a Reader for the GetIcon structure.
type GetIconReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIconReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIconOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIconBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIconNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIconInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIconOK creates a GetIconOK with default headers values
func NewGetIconOK() *GetIconOK {
	return &GetIconOK{}
}

/*
GetIconOK describes a response with status code 200, with default header values.

Success
*/
type GetIconOK struct {
	Payload *models.Icon
}

// IsSuccess returns true when this get icon o k response has a 2xx status code
func (o *GetIconOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get icon o k response has a 3xx status code
func (o *GetIconOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get icon o k response has a 4xx status code
func (o *GetIconOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get icon o k response has a 5xx status code
func (o *GetIconOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get icon o k response a status code equal to that given
func (o *GetIconOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIconOK) Error() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconOK  %+v", 200, o.Payload)
}

func (o *GetIconOK) String() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconOK  %+v", 200, o.Payload)
}

func (o *GetIconOK) GetPayload() *models.Icon {
	return o.Payload
}

func (o *GetIconOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Icon)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIconBadRequest creates a GetIconBadRequest with default headers values
func NewGetIconBadRequest() *GetIconBadRequest {
	return &GetIconBadRequest{}
}

/*
GetIconBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetIconBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get icon bad request response has a 2xx status code
func (o *GetIconBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get icon bad request response has a 3xx status code
func (o *GetIconBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get icon bad request response has a 4xx status code
func (o *GetIconBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get icon bad request response has a 5xx status code
func (o *GetIconBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get icon bad request response a status code equal to that given
func (o *GetIconBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIconBadRequest) Error() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconBadRequest  %+v", 400, o.Payload)
}

func (o *GetIconBadRequest) String() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconBadRequest  %+v", 400, o.Payload)
}

func (o *GetIconBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetIconBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIconNotFound creates a GetIconNotFound with default headers values
func NewGetIconNotFound() *GetIconNotFound {
	return &GetIconNotFound{}
}

/*
GetIconNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetIconNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get icon not found response has a 2xx status code
func (o *GetIconNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get icon not found response has a 3xx status code
func (o *GetIconNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get icon not found response has a 4xx status code
func (o *GetIconNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get icon not found response has a 5xx status code
func (o *GetIconNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get icon not found response a status code equal to that given
func (o *GetIconNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIconNotFound) Error() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconNotFound  %+v", 404, o.Payload)
}

func (o *GetIconNotFound) String() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconNotFound  %+v", 404, o.Payload)
}

func (o *GetIconNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetIconNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIconInternalServerError creates a GetIconInternalServerError with default headers values
func NewGetIconInternalServerError() *GetIconInternalServerError {
	return &GetIconInternalServerError{}
}

/*
GetIconInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetIconInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get icon internal server error response has a 2xx status code
func (o *GetIconInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get icon internal server error response has a 3xx status code
func (o *GetIconInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get icon internal server error response has a 4xx status code
func (o *GetIconInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get icon internal server error response has a 5xx status code
func (o *GetIconInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get icon internal server error response a status code equal to that given
func (o *GetIconInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIconInternalServerError) Error() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIconInternalServerError) String() string {
	return fmt.Sprintf("[GET /icons/{digest}][%d] getIconInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIconInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetIconInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
