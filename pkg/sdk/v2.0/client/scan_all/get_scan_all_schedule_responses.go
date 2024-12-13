// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetScanAllScheduleReader is a Reader for the GetScanAllSchedule structure.
type GetScanAllScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanAllScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScanAllScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScanAllScheduleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScanAllScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewGetScanAllSchedulePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScanAllScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScanAllScheduleOK creates a GetScanAllScheduleOK with default headers values
func NewGetScanAllScheduleOK() *GetScanAllScheduleOK {
	return &GetScanAllScheduleOK{}
}

/*
GetScanAllScheduleOK describes a response with status code 200, with default header values.

Get a schedule for the scan all job, which scans all of images in Harbor.
*/
type GetScanAllScheduleOK struct {
	Payload *models.Schedule
}

// IsSuccess returns true when this get scan all schedule o k response has a 2xx status code
func (o *GetScanAllScheduleOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scan all schedule o k response has a 3xx status code
func (o *GetScanAllScheduleOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan all schedule o k response has a 4xx status code
func (o *GetScanAllScheduleOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scan all schedule o k response has a 5xx status code
func (o *GetScanAllScheduleOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan all schedule o k response a status code equal to that given
func (o *GetScanAllScheduleOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScanAllScheduleOK) Error() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllScheduleOK  %+v", 200, o.Payload)
}

func (o *GetScanAllScheduleOK) String() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllScheduleOK  %+v", 200, o.Payload)
}

func (o *GetScanAllScheduleOK) GetPayload() *models.Schedule {
	return o.Payload
}

func (o *GetScanAllScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Schedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanAllScheduleUnauthorized creates a GetScanAllScheduleUnauthorized with default headers values
func NewGetScanAllScheduleUnauthorized() *GetScanAllScheduleUnauthorized {
	return &GetScanAllScheduleUnauthorized{}
}

/*
GetScanAllScheduleUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetScanAllScheduleUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan all schedule unauthorized response has a 2xx status code
func (o *GetScanAllScheduleUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan all schedule unauthorized response has a 3xx status code
func (o *GetScanAllScheduleUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan all schedule unauthorized response has a 4xx status code
func (o *GetScanAllScheduleUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scan all schedule unauthorized response has a 5xx status code
func (o *GetScanAllScheduleUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan all schedule unauthorized response a status code equal to that given
func (o *GetScanAllScheduleUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScanAllScheduleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanAllScheduleUnauthorized) String() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllScheduleUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanAllScheduleUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanAllScheduleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScanAllScheduleForbidden creates a GetScanAllScheduleForbidden with default headers values
func NewGetScanAllScheduleForbidden() *GetScanAllScheduleForbidden {
	return &GetScanAllScheduleForbidden{}
}

/*
GetScanAllScheduleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetScanAllScheduleForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan all schedule forbidden response has a 2xx status code
func (o *GetScanAllScheduleForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan all schedule forbidden response has a 3xx status code
func (o *GetScanAllScheduleForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan all schedule forbidden response has a 4xx status code
func (o *GetScanAllScheduleForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scan all schedule forbidden response has a 5xx status code
func (o *GetScanAllScheduleForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan all schedule forbidden response a status code equal to that given
func (o *GetScanAllScheduleForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScanAllScheduleForbidden) Error() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetScanAllScheduleForbidden) String() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllScheduleForbidden  %+v", 403, o.Payload)
}

func (o *GetScanAllScheduleForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanAllScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScanAllSchedulePreconditionFailed creates a GetScanAllSchedulePreconditionFailed with default headers values
func NewGetScanAllSchedulePreconditionFailed() *GetScanAllSchedulePreconditionFailed {
	return &GetScanAllSchedulePreconditionFailed{}
}

/*
GetScanAllSchedulePreconditionFailed describes a response with status code 412, with default header values.

Precondition failed
*/
type GetScanAllSchedulePreconditionFailed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan all schedule precondition failed response has a 2xx status code
func (o *GetScanAllSchedulePreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan all schedule precondition failed response has a 3xx status code
func (o *GetScanAllSchedulePreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan all schedule precondition failed response has a 4xx status code
func (o *GetScanAllSchedulePreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scan all schedule precondition failed response has a 5xx status code
func (o *GetScanAllSchedulePreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan all schedule precondition failed response a status code equal to that given
func (o *GetScanAllSchedulePreconditionFailed) IsCode(code int) bool {
	return code == 412
}

func (o *GetScanAllSchedulePreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllSchedulePreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetScanAllSchedulePreconditionFailed) String() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllSchedulePreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetScanAllSchedulePreconditionFailed) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanAllSchedulePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScanAllScheduleInternalServerError creates a GetScanAllScheduleInternalServerError with default headers values
func NewGetScanAllScheduleInternalServerError() *GetScanAllScheduleInternalServerError {
	return &GetScanAllScheduleInternalServerError{}
}

/*
GetScanAllScheduleInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetScanAllScheduleInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan all schedule internal server error response has a 2xx status code
func (o *GetScanAllScheduleInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan all schedule internal server error response has a 3xx status code
func (o *GetScanAllScheduleInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan all schedule internal server error response has a 4xx status code
func (o *GetScanAllScheduleInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scan all schedule internal server error response has a 5xx status code
func (o *GetScanAllScheduleInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scan all schedule internal server error response a status code equal to that given
func (o *GetScanAllScheduleInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScanAllScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScanAllScheduleInternalServerError) String() string {
	return fmt.Sprintf("[GET /system/scanAll/schedule][%d] getScanAllScheduleInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScanAllScheduleInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanAllScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
