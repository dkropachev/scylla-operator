// Code generated by go-swagger; DO NOT EDIT.

package config

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

// NewFindConfigStoragePortParams creates a new FindConfigStoragePortParams object
// with the default values initialized.
func NewFindConfigStoragePortParams() *FindConfigStoragePortParams {

	return &FindConfigStoragePortParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigStoragePortParamsWithTimeout creates a new FindConfigStoragePortParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigStoragePortParamsWithTimeout(timeout time.Duration) *FindConfigStoragePortParams {

	return &FindConfigStoragePortParams{

		timeout: timeout,
	}
}

// NewFindConfigStoragePortParamsWithContext creates a new FindConfigStoragePortParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigStoragePortParamsWithContext(ctx context.Context) *FindConfigStoragePortParams {

	return &FindConfigStoragePortParams{

		Context: ctx,
	}
}

// NewFindConfigStoragePortParamsWithHTTPClient creates a new FindConfigStoragePortParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigStoragePortParamsWithHTTPClient(client *http.Client) *FindConfigStoragePortParams {

	return &FindConfigStoragePortParams{
		HTTPClient: client,
	}
}

/*
FindConfigStoragePortParams contains all the parameters to send to the API endpoint
for the find config storage port operation typically these are written to a http.Request
*/
type FindConfigStoragePortParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config storage port params
func (o *FindConfigStoragePortParams) WithTimeout(timeout time.Duration) *FindConfigStoragePortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config storage port params
func (o *FindConfigStoragePortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config storage port params
func (o *FindConfigStoragePortParams) WithContext(ctx context.Context) *FindConfigStoragePortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config storage port params
func (o *FindConfigStoragePortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config storage port params
func (o *FindConfigStoragePortParams) WithHTTPClient(client *http.Client) *FindConfigStoragePortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config storage port params
func (o *FindConfigStoragePortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigStoragePortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
