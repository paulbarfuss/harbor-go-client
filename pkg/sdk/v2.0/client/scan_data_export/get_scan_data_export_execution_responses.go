// Code generated by go-swagger; DO NOT EDIT.

package scan_data_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// GetScanDataExportExecutionReader is a Reader for the GetScanDataExportExecution structure.
type GetScanDataExportExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanDataExportExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScanDataExportExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetScanDataExportExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScanDataExportExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScanDataExportExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScanDataExportExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScanDataExportExecutionOK creates a GetScanDataExportExecutionOK with default headers values
func NewGetScanDataExportExecutionOK() *GetScanDataExportExecutionOK {
	return &GetScanDataExportExecutionOK{}
}

/*
GetScanDataExportExecutionOK describes a response with status code 200, with default header values.

Success
*/
type GetScanDataExportExecutionOK struct {
	Payload *models.ScanDataExportExecution
}

// IsSuccess returns true when this get scan data export execution o k response has a 2xx status code
func (o *GetScanDataExportExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scan data export execution o k response has a 3xx status code
func (o *GetScanDataExportExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution o k response has a 4xx status code
func (o *GetScanDataExportExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scan data export execution o k response has a 5xx status code
func (o *GetScanDataExportExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan data export execution o k response a status code equal to that given
func (o *GetScanDataExportExecutionOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScanDataExportExecutionOK) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionOK  %+v", 200, o.Payload)
}

func (o *GetScanDataExportExecutionOK) String() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionOK  %+v", 200, o.Payload)
}

func (o *GetScanDataExportExecutionOK) GetPayload() *models.ScanDataExportExecution {
	return o.Payload
}

func (o *GetScanDataExportExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScanDataExportExecution)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanDataExportExecutionUnauthorized creates a GetScanDataExportExecutionUnauthorized with default headers values
func NewGetScanDataExportExecutionUnauthorized() *GetScanDataExportExecutionUnauthorized {
	return &GetScanDataExportExecutionUnauthorized{}
}

/*
GetScanDataExportExecutionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetScanDataExportExecutionUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan data export execution unauthorized response has a 2xx status code
func (o *GetScanDataExportExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan data export execution unauthorized response has a 3xx status code
func (o *GetScanDataExportExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution unauthorized response has a 4xx status code
func (o *GetScanDataExportExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scan data export execution unauthorized response has a 5xx status code
func (o *GetScanDataExportExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan data export execution unauthorized response a status code equal to that given
func (o *GetScanDataExportExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScanDataExportExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanDataExportExecutionUnauthorized) String() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanDataExportExecutionUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScanDataExportExecutionForbidden creates a GetScanDataExportExecutionForbidden with default headers values
func NewGetScanDataExportExecutionForbidden() *GetScanDataExportExecutionForbidden {
	return &GetScanDataExportExecutionForbidden{}
}

/*
GetScanDataExportExecutionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetScanDataExportExecutionForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan data export execution forbidden response has a 2xx status code
func (o *GetScanDataExportExecutionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan data export execution forbidden response has a 3xx status code
func (o *GetScanDataExportExecutionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution forbidden response has a 4xx status code
func (o *GetScanDataExportExecutionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scan data export execution forbidden response has a 5xx status code
func (o *GetScanDataExportExecutionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan data export execution forbidden response a status code equal to that given
func (o *GetScanDataExportExecutionForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScanDataExportExecutionForbidden) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionForbidden  %+v", 403, o.Payload)
}

func (o *GetScanDataExportExecutionForbidden) String() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionForbidden  %+v", 403, o.Payload)
}

func (o *GetScanDataExportExecutionForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScanDataExportExecutionNotFound creates a GetScanDataExportExecutionNotFound with default headers values
func NewGetScanDataExportExecutionNotFound() *GetScanDataExportExecutionNotFound {
	return &GetScanDataExportExecutionNotFound{}
}

/*
GetScanDataExportExecutionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetScanDataExportExecutionNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan data export execution not found response has a 2xx status code
func (o *GetScanDataExportExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan data export execution not found response has a 3xx status code
func (o *GetScanDataExportExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution not found response has a 4xx status code
func (o *GetScanDataExportExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scan data export execution not found response has a 5xx status code
func (o *GetScanDataExportExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scan data export execution not found response a status code equal to that given
func (o *GetScanDataExportExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetScanDataExportExecutionNotFound) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionNotFound  %+v", 404, o.Payload)
}

func (o *GetScanDataExportExecutionNotFound) String() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionNotFound  %+v", 404, o.Payload)
}

func (o *GetScanDataExportExecutionNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetScanDataExportExecutionInternalServerError creates a GetScanDataExportExecutionInternalServerError with default headers values
func NewGetScanDataExportExecutionInternalServerError() *GetScanDataExportExecutionInternalServerError {
	return &GetScanDataExportExecutionInternalServerError{}
}

/*
GetScanDataExportExecutionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetScanDataExportExecutionInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get scan data export execution internal server error response has a 2xx status code
func (o *GetScanDataExportExecutionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scan data export execution internal server error response has a 3xx status code
func (o *GetScanDataExportExecutionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scan data export execution internal server error response has a 4xx status code
func (o *GetScanDataExportExecutionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scan data export execution internal server error response has a 5xx status code
func (o *GetScanDataExportExecutionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scan data export execution internal server error response a status code equal to that given
func (o *GetScanDataExportExecutionInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScanDataExportExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScanDataExportExecutionInternalServerError) String() string {
	return fmt.Sprintf("[GET /export/cve/execution/{execution_id}][%d] getScanDataExportExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScanDataExportExecutionInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetScanDataExportExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
