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

// NewStorageServiceRemovalStatusGetParams creates a new StorageServiceRemovalStatusGetParams object
// with the default values initialized.
func NewStorageServiceRemovalStatusGetParams() *StorageServiceRemovalStatusGetParams {

	return &StorageServiceRemovalStatusGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceRemovalStatusGetParamsWithTimeout creates a new StorageServiceRemovalStatusGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceRemovalStatusGetParamsWithTimeout(timeout time.Duration) *StorageServiceRemovalStatusGetParams {

	return &StorageServiceRemovalStatusGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceRemovalStatusGetParamsWithContext creates a new StorageServiceRemovalStatusGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceRemovalStatusGetParamsWithContext(ctx context.Context) *StorageServiceRemovalStatusGetParams {

	return &StorageServiceRemovalStatusGetParams{

		Context: ctx,
	}
}

// NewStorageServiceRemovalStatusGetParamsWithHTTPClient creates a new StorageServiceRemovalStatusGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceRemovalStatusGetParamsWithHTTPClient(client *http.Client) *StorageServiceRemovalStatusGetParams {

	return &StorageServiceRemovalStatusGetParams{
		HTTPClient: client,
	}
}

/*
StorageServiceRemovalStatusGetParams contains all the parameters to send to the API endpoint
for the storage service removal status get operation typically these are written to a http.Request
*/
type StorageServiceRemovalStatusGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service removal status get params
func (o *StorageServiceRemovalStatusGetParams) WithTimeout(timeout time.Duration) *StorageServiceRemovalStatusGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service removal status get params
func (o *StorageServiceRemovalStatusGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service removal status get params
func (o *StorageServiceRemovalStatusGetParams) WithContext(ctx context.Context) *StorageServiceRemovalStatusGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service removal status get params
func (o *StorageServiceRemovalStatusGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service removal status get params
func (o *StorageServiceRemovalStatusGetParams) WithHTTPClient(client *http.Client) *StorageServiceRemovalStatusGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service removal status get params
func (o *StorageServiceRemovalStatusGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceRemovalStatusGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
