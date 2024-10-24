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

// NewCacheServiceInvalidateCounterCachePostParams creates a new CacheServiceInvalidateCounterCachePostParams object
// with the default values initialized.
func NewCacheServiceInvalidateCounterCachePostParams() *CacheServiceInvalidateCounterCachePostParams {

	return &CacheServiceInvalidateCounterCachePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceInvalidateCounterCachePostParamsWithTimeout creates a new CacheServiceInvalidateCounterCachePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceInvalidateCounterCachePostParamsWithTimeout(timeout time.Duration) *CacheServiceInvalidateCounterCachePostParams {

	return &CacheServiceInvalidateCounterCachePostParams{

		timeout: timeout,
	}
}

// NewCacheServiceInvalidateCounterCachePostParamsWithContext creates a new CacheServiceInvalidateCounterCachePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceInvalidateCounterCachePostParamsWithContext(ctx context.Context) *CacheServiceInvalidateCounterCachePostParams {

	return &CacheServiceInvalidateCounterCachePostParams{

		Context: ctx,
	}
}

// NewCacheServiceInvalidateCounterCachePostParamsWithHTTPClient creates a new CacheServiceInvalidateCounterCachePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceInvalidateCounterCachePostParamsWithHTTPClient(client *http.Client) *CacheServiceInvalidateCounterCachePostParams {

	return &CacheServiceInvalidateCounterCachePostParams{
		HTTPClient: client,
	}
}

/*
CacheServiceInvalidateCounterCachePostParams contains all the parameters to send to the API endpoint
for the cache service invalidate counter cache post operation typically these are written to a http.Request
*/
type CacheServiceInvalidateCounterCachePostParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service invalidate counter cache post params
func (o *CacheServiceInvalidateCounterCachePostParams) WithTimeout(timeout time.Duration) *CacheServiceInvalidateCounterCachePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service invalidate counter cache post params
func (o *CacheServiceInvalidateCounterCachePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service invalidate counter cache post params
func (o *CacheServiceInvalidateCounterCachePostParams) WithContext(ctx context.Context) *CacheServiceInvalidateCounterCachePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service invalidate counter cache post params
func (o *CacheServiceInvalidateCounterCachePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service invalidate counter cache post params
func (o *CacheServiceInvalidateCounterCachePostParams) WithHTTPClient(client *http.Client) *CacheServiceInvalidateCounterCachePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service invalidate counter cache post params
func (o *CacheServiceInvalidateCounterCachePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceInvalidateCounterCachePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
