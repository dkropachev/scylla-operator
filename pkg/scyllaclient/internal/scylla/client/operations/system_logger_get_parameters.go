// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSystemLoggerGetParams creates a new SystemLoggerGetParams object
// with the default values initialized.
func NewSystemLoggerGetParams() *SystemLoggerGetParams {

	return &SystemLoggerGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSystemLoggerGetParamsWithTimeout creates a new SystemLoggerGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSystemLoggerGetParamsWithTimeout(timeout time.Duration) *SystemLoggerGetParams {

	return &SystemLoggerGetParams{

		timeout: timeout,
	}
}

// NewSystemLoggerGetParamsWithContext creates a new SystemLoggerGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewSystemLoggerGetParamsWithContext(ctx context.Context) *SystemLoggerGetParams {

	return &SystemLoggerGetParams{

		Context: ctx,
	}
}

// NewSystemLoggerGetParamsWithHTTPClient creates a new SystemLoggerGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSystemLoggerGetParamsWithHTTPClient(client *http.Client) *SystemLoggerGetParams {

	return &SystemLoggerGetParams{
		HTTPClient: client,
	}
}

/*
SystemLoggerGetParams contains all the parameters to send to the API endpoint
for the system logger get operation typically these are written to a http.Request
*/
type SystemLoggerGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the system logger get params
func (o *SystemLoggerGetParams) WithTimeout(timeout time.Duration) *SystemLoggerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system logger get params
func (o *SystemLoggerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system logger get params
func (o *SystemLoggerGetParams) WithContext(ctx context.Context) *SystemLoggerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system logger get params
func (o *SystemLoggerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system logger get params
func (o *SystemLoggerGetParams) WithHTTPClient(client *http.Client) *SystemLoggerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system logger get params
func (o *SystemLoggerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SystemLoggerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
