// Code generated by go-swagger; DO NOT EDIT.

package usergroup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// SearchUserGroupsReader is a Reader for the SearchUserGroups structure.
type SearchUserGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchUserGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSearchUserGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSearchUserGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSearchUserGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSearchUserGroupsOK creates a SearchUserGroupsOK with default headers values
func NewSearchUserGroupsOK() *SearchUserGroupsOK {
	return &SearchUserGroupsOK{}
}

/*SearchUserGroupsOK handles this case with default header values.

Search groups successfully.
*/
type SearchUserGroupsOK struct {
	/*Link to previous page and next page
	 */
	Link string
	/*The total count of available items
	 */
	XTotalCount int64

	Payload []*models.UserGroupSearchItem
}

func (o *SearchUserGroupsOK) Error() string {
	return fmt.Sprintf("[GET /usergroups/search][%d] searchUserGroupsOK  %+v", 200, o.Payload)
}

func (o *SearchUserGroupsOK) GetPayload() []*models.UserGroupSearchItem {
	return o.Payload
}

func (o *SearchUserGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Link
	o.Link = response.GetHeader("Link")

	// response header X-Total-Count
	xTotalCount, err := swag.ConvertInt64(response.GetHeader("X-Total-Count"))
	if err != nil {
		return errors.InvalidType("X-Total-Count", "header", "int64", response.GetHeader("X-Total-Count"))
	}
	o.XTotalCount = xTotalCount

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUserGroupsUnauthorized creates a SearchUserGroupsUnauthorized with default headers values
func NewSearchUserGroupsUnauthorized() *SearchUserGroupsUnauthorized {
	return &SearchUserGroupsUnauthorized{}
}

/*SearchUserGroupsUnauthorized handles this case with default header values.

Unauthorized
*/
type SearchUserGroupsUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *SearchUserGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /usergroups/search][%d] searchUserGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *SearchUserGroupsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchUserGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSearchUserGroupsInternalServerError creates a SearchUserGroupsInternalServerError with default headers values
func NewSearchUserGroupsInternalServerError() *SearchUserGroupsInternalServerError {
	return &SearchUserGroupsInternalServerError{}
}

/*SearchUserGroupsInternalServerError handles this case with default header values.

Internal server error
*/
type SearchUserGroupsInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *SearchUserGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /usergroups/search][%d] searchUserGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *SearchUserGroupsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *SearchUserGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}