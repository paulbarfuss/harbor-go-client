// Code generated by go-swagger; DO NOT EDIT.

package purge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/paulbarfuss/harbor-go-client/pkg/sdk/v2.0/models"
)

// NewCreatePurgeScheduleParams creates a new CreatePurgeScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePurgeScheduleParams() *CreatePurgeScheduleParams {
	return &CreatePurgeScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePurgeScheduleParamsWithTimeout creates a new CreatePurgeScheduleParams object
// with the ability to set a timeout on a request.
func NewCreatePurgeScheduleParamsWithTimeout(timeout time.Duration) *CreatePurgeScheduleParams {
	return &CreatePurgeScheduleParams{
		timeout: timeout,
	}
}

// NewCreatePurgeScheduleParamsWithContext creates a new CreatePurgeScheduleParams object
// with the ability to set a context for a request.
func NewCreatePurgeScheduleParamsWithContext(ctx context.Context) *CreatePurgeScheduleParams {
	return &CreatePurgeScheduleParams{
		Context: ctx,
	}
}

// NewCreatePurgeScheduleParamsWithHTTPClient creates a new CreatePurgeScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePurgeScheduleParamsWithHTTPClient(client *http.Client) *CreatePurgeScheduleParams {
	return &CreatePurgeScheduleParams{
		HTTPClient: client,
	}
}

/*
CreatePurgeScheduleParams contains all the parameters to send to the API endpoint

	for the create purge schedule operation.

	Typically these are written to a http.Request.
*/
type CreatePurgeScheduleParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* Schedule.

	     The purge job's schedule, it is a json object. |
	The sample format is |
	{"parameters":{"audit_retention_hour":168,"dry_run":true, "include_operations":"create,delete,pull"},"schedule":{"type":"Hourly","cron":"0 0 * * * *"}} |
	the include_operation should be a comma separated string, e.g. create,delete,pull, if it is empty, no operation will be purged.

	*/
	Schedule *models.Schedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create purge schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePurgeScheduleParams) WithDefaults() *CreatePurgeScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create purge schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePurgeScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create purge schedule params
func (o *CreatePurgeScheduleParams) WithTimeout(timeout time.Duration) *CreatePurgeScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create purge schedule params
func (o *CreatePurgeScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create purge schedule params
func (o *CreatePurgeScheduleParams) WithContext(ctx context.Context) *CreatePurgeScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create purge schedule params
func (o *CreatePurgeScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create purge schedule params
func (o *CreatePurgeScheduleParams) WithHTTPClient(client *http.Client) *CreatePurgeScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create purge schedule params
func (o *CreatePurgeScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create purge schedule params
func (o *CreatePurgeScheduleParams) WithXRequestID(xRequestID *string) *CreatePurgeScheduleParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create purge schedule params
func (o *CreatePurgeScheduleParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSchedule adds the schedule to the create purge schedule params
func (o *CreatePurgeScheduleParams) WithSchedule(schedule *models.Schedule) *CreatePurgeScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the create purge schedule params
func (o *CreatePurgeScheduleParams) SetSchedule(schedule *models.Schedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePurgeScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}
	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
