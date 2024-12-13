// Code generated by go-swagger; DO NOT EDIT.

package ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// SearchLdapGroupReader is a Reader for the SearchLdapGroup structure.
type SearchLdapGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchLdapGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchLdapGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSearchLdapGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewSearchLdapGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSearchLdapGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchLdapGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchLdapGroupOK creates a SearchLdapGroupOK with default headers values
func NewSearchLdapGroupOK() *SearchLdapGroupOK {
	return &SearchLdapGroupOK{}
}

/*
SearchLdapGroupOK describes a response with status code 200, with default header values.

Search ldap group successfully.
*/
type SearchLdapGroupOK struct {
	Payload []*models.UserGroup
}

// IsSuccess returns true when this search ldap group o k response has a 2xx status code
func (o *SearchLdapGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this search ldap group o k response has a 3xx status code
func (o *SearchLdapGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap group o k response has a 4xx status code
func (o *SearchLdapGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this search ldap group o k response has a 5xx status code
func (o *SearchLdapGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this search ldap group o k response a status code equal to that given
func (o *SearchLdapGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *SearchLdapGroupOK) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupOK  %+v", 200, o.Payload)
}

func (o *SearchLdapGroupOK) String() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupOK  %+v", 200, o.Payload)
}

func (o *SearchLdapGroupOK) GetPayload() []*models.UserGroup {
	return o.Payload
}

func (o *SearchLdapGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchLdapGroupBadRequest creates a SearchLdapGroupBadRequest with default headers values
func NewSearchLdapGroupBadRequest() *SearchLdapGroupBadRequest {
	return &SearchLdapGroupBadRequest{}
}

/*
SearchLdapGroupBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type SearchLdapGroupBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search ldap group bad request response has a 2xx status code
func (o *SearchLdapGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ldap group bad request response has a 3xx status code
func (o *SearchLdapGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap group bad request response has a 4xx status code
func (o *SearchLdapGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ldap group bad request response has a 5xx status code
func (o *SearchLdapGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this search ldap group bad request response a status code equal to that given
func (o *SearchLdapGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *SearchLdapGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupBadRequest  %+v", 400, o.Payload)
}

func (o *SearchLdapGroupBadRequest) String() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupBadRequest  %+v", 400, o.Payload)
}

func (o *SearchLdapGroupBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchLdapGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchLdapGroupUnauthorized creates a SearchLdapGroupUnauthorized with default headers values
func NewSearchLdapGroupUnauthorized() *SearchLdapGroupUnauthorized {
	return &SearchLdapGroupUnauthorized{}
}

/*
SearchLdapGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type SearchLdapGroupUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search ldap group unauthorized response has a 2xx status code
func (o *SearchLdapGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ldap group unauthorized response has a 3xx status code
func (o *SearchLdapGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap group unauthorized response has a 4xx status code
func (o *SearchLdapGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ldap group unauthorized response has a 5xx status code
func (o *SearchLdapGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this search ldap group unauthorized response a status code equal to that given
func (o *SearchLdapGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *SearchLdapGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchLdapGroupUnauthorized) String() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchLdapGroupUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchLdapGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchLdapGroupForbidden creates a SearchLdapGroupForbidden with default headers values
func NewSearchLdapGroupForbidden() *SearchLdapGroupForbidden {
	return &SearchLdapGroupForbidden{}
}

/*
SearchLdapGroupForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SearchLdapGroupForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search ldap group forbidden response has a 2xx status code
func (o *SearchLdapGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ldap group forbidden response has a 3xx status code
func (o *SearchLdapGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap group forbidden response has a 4xx status code
func (o *SearchLdapGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this search ldap group forbidden response has a 5xx status code
func (o *SearchLdapGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this search ldap group forbidden response a status code equal to that given
func (o *SearchLdapGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *SearchLdapGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupForbidden  %+v", 403, o.Payload)
}

func (o *SearchLdapGroupForbidden) String() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupForbidden  %+v", 403, o.Payload)
}

func (o *SearchLdapGroupForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchLdapGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSearchLdapGroupInternalServerError creates a SearchLdapGroupInternalServerError with default headers values
func NewSearchLdapGroupInternalServerError() *SearchLdapGroupInternalServerError {
	return &SearchLdapGroupInternalServerError{}
}

/*
SearchLdapGroupInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type SearchLdapGroupInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this search ldap group internal server error response has a 2xx status code
func (o *SearchLdapGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this search ldap group internal server error response has a 3xx status code
func (o *SearchLdapGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this search ldap group internal server error response has a 4xx status code
func (o *SearchLdapGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this search ldap group internal server error response has a 5xx status code
func (o *SearchLdapGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this search ldap group internal server error response a status code equal to that given
func (o *SearchLdapGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *SearchLdapGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchLdapGroupInternalServerError) String() string {
	return fmt.Sprintf("[GET /ldap/groups/search][%d] searchLdapGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchLdapGroupInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchLdapGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
