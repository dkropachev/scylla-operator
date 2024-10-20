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

// NewFindConfigCommitlogTotalSpaceInMbParams creates a new FindConfigCommitlogTotalSpaceInMbParams object
// with the default values initialized.
func NewFindConfigCommitlogTotalSpaceInMbParams() *FindConfigCommitlogTotalSpaceInMbParams {

	return &FindConfigCommitlogTotalSpaceInMbParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFindConfigCommitlogTotalSpaceInMbParamsWithTimeout creates a new FindConfigCommitlogTotalSpaceInMbParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFindConfigCommitlogTotalSpaceInMbParamsWithTimeout(timeout time.Duration) *FindConfigCommitlogTotalSpaceInMbParams {

	return &FindConfigCommitlogTotalSpaceInMbParams{

		timeout: timeout,
	}
}

// NewFindConfigCommitlogTotalSpaceInMbParamsWithContext creates a new FindConfigCommitlogTotalSpaceInMbParams object
// with the default values initialized, and the ability to set a context for a request
func NewFindConfigCommitlogTotalSpaceInMbParamsWithContext(ctx context.Context) *FindConfigCommitlogTotalSpaceInMbParams {

	return &FindConfigCommitlogTotalSpaceInMbParams{

		Context: ctx,
	}
}

// NewFindConfigCommitlogTotalSpaceInMbParamsWithHTTPClient creates a new FindConfigCommitlogTotalSpaceInMbParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFindConfigCommitlogTotalSpaceInMbParamsWithHTTPClient(client *http.Client) *FindConfigCommitlogTotalSpaceInMbParams {

	return &FindConfigCommitlogTotalSpaceInMbParams{
		HTTPClient: client,
	}
}

/*
FindConfigCommitlogTotalSpaceInMbParams contains all the parameters to send to the API endpoint
for the find config commitlog total space in mb operation typically these are written to a http.Request
*/
type FindConfigCommitlogTotalSpaceInMbParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the find config commitlog total space in mb params
func (o *FindConfigCommitlogTotalSpaceInMbParams) WithTimeout(timeout time.Duration) *FindConfigCommitlogTotalSpaceInMbParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find config commitlog total space in mb params
func (o *FindConfigCommitlogTotalSpaceInMbParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find config commitlog total space in mb params
func (o *FindConfigCommitlogTotalSpaceInMbParams) WithContext(ctx context.Context) *FindConfigCommitlogTotalSpaceInMbParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find config commitlog total space in mb params
func (o *FindConfigCommitlogTotalSpaceInMbParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find config commitlog total space in mb params
func (o *FindConfigCommitlogTotalSpaceInMbParams) WithHTTPClient(client *http.Client) *FindConfigCommitlogTotalSpaceInMbParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find config commitlog total space in mb params
func (o *FindConfigCommitlogTotalSpaceInMbParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *FindConfigCommitlogTotalSpaceInMbParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
