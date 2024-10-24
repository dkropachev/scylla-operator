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

// NewCacheServiceMetricsKeyRequestsMovingAvrageGetParams creates a new CacheServiceMetricsKeyRequestsMovingAvrageGetParams object
// with the default values initialized.
func NewCacheServiceMetricsKeyRequestsMovingAvrageGetParams() *CacheServiceMetricsKeyRequestsMovingAvrageGetParams {

	return &CacheServiceMetricsKeyRequestsMovingAvrageGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCacheServiceMetricsKeyRequestsMovingAvrageGetParamsWithTimeout creates a new CacheServiceMetricsKeyRequestsMovingAvrageGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCacheServiceMetricsKeyRequestsMovingAvrageGetParamsWithTimeout(timeout time.Duration) *CacheServiceMetricsKeyRequestsMovingAvrageGetParams {

	return &CacheServiceMetricsKeyRequestsMovingAvrageGetParams{

		timeout: timeout,
	}
}

// NewCacheServiceMetricsKeyRequestsMovingAvrageGetParamsWithContext creates a new CacheServiceMetricsKeyRequestsMovingAvrageGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewCacheServiceMetricsKeyRequestsMovingAvrageGetParamsWithContext(ctx context.Context) *CacheServiceMetricsKeyRequestsMovingAvrageGetParams {

	return &CacheServiceMetricsKeyRequestsMovingAvrageGetParams{

		Context: ctx,
	}
}

// NewCacheServiceMetricsKeyRequestsMovingAvrageGetParamsWithHTTPClient creates a new CacheServiceMetricsKeyRequestsMovingAvrageGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCacheServiceMetricsKeyRequestsMovingAvrageGetParamsWithHTTPClient(client *http.Client) *CacheServiceMetricsKeyRequestsMovingAvrageGetParams {

	return &CacheServiceMetricsKeyRequestsMovingAvrageGetParams{
		HTTPClient: client,
	}
}

/*
CacheServiceMetricsKeyRequestsMovingAvrageGetParams contains all the parameters to send to the API endpoint
for the cache service metrics key requests moving avrage get operation typically these are written to a http.Request
*/
type CacheServiceMetricsKeyRequestsMovingAvrageGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the cache service metrics key requests moving avrage get params
func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetParams) WithTimeout(timeout time.Duration) *CacheServiceMetricsKeyRequestsMovingAvrageGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cache service metrics key requests moving avrage get params
func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cache service metrics key requests moving avrage get params
func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetParams) WithContext(ctx context.Context) *CacheServiceMetricsKeyRequestsMovingAvrageGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cache service metrics key requests moving avrage get params
func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cache service metrics key requests moving avrage get params
func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetParams) WithHTTPClient(client *http.Client) *CacheServiceMetricsKeyRequestsMovingAvrageGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cache service metrics key requests moving avrage get params
func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CacheServiceMetricsKeyRequestsMovingAvrageGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
