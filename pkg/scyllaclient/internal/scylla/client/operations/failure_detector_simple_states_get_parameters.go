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

// NewFailureDetectorSimpleStatesGetParams creates a new FailureDetectorSimpleStatesGetParams object
// with the default values initialized.
func NewFailureDetectorSimpleStatesGetParams() *FailureDetectorSimpleStatesGetParams {

	return &FailureDetectorSimpleStatesGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFailureDetectorSimpleStatesGetParamsWithTimeout creates a new FailureDetectorSimpleStatesGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFailureDetectorSimpleStatesGetParamsWithTimeout(timeout time.Duration) *FailureDetectorSimpleStatesGetParams {

	return &FailureDetectorSimpleStatesGetParams{

		timeout: timeout,
	}
}

// NewFailureDetectorSimpleStatesGetParamsWithContext creates a new FailureDetectorSimpleStatesGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewFailureDetectorSimpleStatesGetParamsWithContext(ctx context.Context) *FailureDetectorSimpleStatesGetParams {

	return &FailureDetectorSimpleStatesGetParams{

		Context: ctx,
	}
}

// NewFailureDetectorSimpleStatesGetParamsWithHTTPClient creates a new FailureDetectorSimpleStatesGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFailureDetectorSimpleStatesGetParamsWithHTTPClient(client *http.Client) *FailureDetectorSimpleStatesGetParams {

	return &FailureDetectorSimpleStatesGetParams{
		HTTPClient: client,
	}
}

/*
FailureDetectorSimpleStatesGetParams contains all the parameters to send to the API endpoint
for the failure detector simple states get operation typically these are written to a http.Request
*/
type FailureDetectorSimpleStatesGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the failure detector simple states get params
func (o *FailureDetectorSimpleStatesGetParams) WithTimeout(timeout time.Duration) *FailureDetectorSimpleStatesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the failure detector simple states get params
func (o *FailureDetectorSimpleStatesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the failure detector simple states get params
func (o *FailureDetectorSimpleStatesGetParams) WithContext(ctx context.Context) *FailureDetectorSimpleStatesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the failure detector simple states get params
func (o *FailureDetectorSimpleStatesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the failure detector simple states get params
func (o *FailureDetectorSimpleStatesGetParams) WithHTTPClient(client *http.Client) *FailureDetectorSimpleStatesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the failure detector simple states get params
func (o *FailureDetectorSimpleStatesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FailureDetectorSimpleStatesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
