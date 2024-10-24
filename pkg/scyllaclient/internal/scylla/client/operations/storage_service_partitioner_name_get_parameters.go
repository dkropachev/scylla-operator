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

// NewStorageServicePartitionerNameGetParams creates a new StorageServicePartitionerNameGetParams object
// with the default values initialized.
func NewStorageServicePartitionerNameGetParams() *StorageServicePartitionerNameGetParams {

	return &StorageServicePartitionerNameGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServicePartitionerNameGetParamsWithTimeout creates a new StorageServicePartitionerNameGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServicePartitionerNameGetParamsWithTimeout(timeout time.Duration) *StorageServicePartitionerNameGetParams {

	return &StorageServicePartitionerNameGetParams{

		timeout: timeout,
	}
}

// NewStorageServicePartitionerNameGetParamsWithContext creates a new StorageServicePartitionerNameGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServicePartitionerNameGetParamsWithContext(ctx context.Context) *StorageServicePartitionerNameGetParams {

	return &StorageServicePartitionerNameGetParams{

		Context: ctx,
	}
}

// NewStorageServicePartitionerNameGetParamsWithHTTPClient creates a new StorageServicePartitionerNameGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServicePartitionerNameGetParamsWithHTTPClient(client *http.Client) *StorageServicePartitionerNameGetParams {

	return &StorageServicePartitionerNameGetParams{
		HTTPClient: client,
	}
}

/*
StorageServicePartitionerNameGetParams contains all the parameters to send to the API endpoint
for the storage service partitioner name get operation typically these are written to a http.Request
*/
type StorageServicePartitionerNameGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service partitioner name get params
func (o *StorageServicePartitionerNameGetParams) WithTimeout(timeout time.Duration) *StorageServicePartitionerNameGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service partitioner name get params
func (o *StorageServicePartitionerNameGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service partitioner name get params
func (o *StorageServicePartitionerNameGetParams) WithContext(ctx context.Context) *StorageServicePartitionerNameGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service partitioner name get params
func (o *StorageServicePartitionerNameGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service partitioner name get params
func (o *StorageServicePartitionerNameGetParams) WithHTTPClient(client *http.Client) *StorageServicePartitionerNameGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service partitioner name get params
func (o *StorageServicePartitionerNameGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServicePartitionerNameGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
