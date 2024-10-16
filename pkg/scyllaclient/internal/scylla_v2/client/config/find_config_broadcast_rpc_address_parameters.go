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

// NewFindConfigBroadcastRPCAddressParams creates a new FindConfigBroadcastRPCAddressParams object
// with the default values initialized.
func NewFindConfigBroadcastRPCAddressParams() *FindConfigBroadcastRPCAddressParams {

	return &FindConfigBroadcastRPCAddressParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigBroadcastRPCAddressParamsWithTimeout creates a new FindConfigBroadcastRPCAddressParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigBroadcastRPCAddressParamsWithTimeout(timeout time.Duration) *FindConfigBroadcastRPCAddressParams {

	return &FindConfigBroadcastRPCAddressParams{

		timeout: timeout,
	}
}

// NewFindConfigBroadcastRPCAddressParamsWithContext creates a new FindConfigBroadcastRPCAddressParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigBroadcastRPCAddressParamsWithContext(ctx context.Context) *FindConfigBroadcastRPCAddressParams {

	return &FindConfigBroadcastRPCAddressParams{

		Context: ctx,
	}
}

// NewFindConfigBroadcastRPCAddressParamsWithHTTPClient creates a new FindConfigBroadcastRPCAddressParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigBroadcastRPCAddressParamsWithHTTPClient(client *http.Client) *FindConfigBroadcastRPCAddressParams {

	return &FindConfigBroadcastRPCAddressParams{
		HTTPClient: client,
	}
}

/*
FindConfigBroadcastRPCAddressParams contains all the parameters to send to the API endpoint
for the find config broadcast rpc address operation typically these are written to a http.Request
*/
type FindConfigBroadcastRPCAddressParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config broadcast rpc address params
func (o *FindConfigBroadcastRPCAddressParams) WithTimeout(timeout time.Duration) *FindConfigBroadcastRPCAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config broadcast rpc address params
func (o *FindConfigBroadcastRPCAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config broadcast rpc address params
func (o *FindConfigBroadcastRPCAddressParams) WithContext(ctx context.Context) *FindConfigBroadcastRPCAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config broadcast rpc address params
func (o *FindConfigBroadcastRPCAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config broadcast rpc address params
func (o *FindConfigBroadcastRPCAddressParams) WithHTTPClient(client *http.Client) *FindConfigBroadcastRPCAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config broadcast rpc address params
func (o *FindConfigBroadcastRPCAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigBroadcastRPCAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
