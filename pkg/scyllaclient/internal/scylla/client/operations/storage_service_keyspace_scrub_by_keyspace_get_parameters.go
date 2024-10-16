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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewStorageServiceKeyspaceScrubByKeyspaceGetParams creates a new StorageServiceKeyspaceScrubByKeyspaceGetParams object
// with the default values initialized.
func NewStorageServiceKeyspaceScrubByKeyspaceGetParams() *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	var ()
	return &StorageServiceKeyspaceScrubByKeyspaceGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStorageServiceKeyspaceScrubByKeyspaceGetParamsWithTimeout creates a new StorageServiceKeyspaceScrubByKeyspaceGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStorageServiceKeyspaceScrubByKeyspaceGetParamsWithTimeout(timeout time.Duration) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	var ()
	return &StorageServiceKeyspaceScrubByKeyspaceGetParams{

		timeout: timeout,
	}
}

// NewStorageServiceKeyspaceScrubByKeyspaceGetParamsWithContext creates a new StorageServiceKeyspaceScrubByKeyspaceGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewStorageServiceKeyspaceScrubByKeyspaceGetParamsWithContext(ctx context.Context) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	var ()
	return &StorageServiceKeyspaceScrubByKeyspaceGetParams{

		Context: ctx,
	}
}

// NewStorageServiceKeyspaceScrubByKeyspaceGetParamsWithHTTPClient creates a new StorageServiceKeyspaceScrubByKeyspaceGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStorageServiceKeyspaceScrubByKeyspaceGetParamsWithHTTPClient(client *http.Client) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	var ()
	return &StorageServiceKeyspaceScrubByKeyspaceGetParams{
		HTTPClient: client,
	}
}

/*
StorageServiceKeyspaceScrubByKeyspaceGetParams contains all the parameters to send to the API endpoint
for the storage service keyspace scrub by keyspace get operation typically these are written to a http.Request
*/
type StorageServiceKeyspaceScrubByKeyspaceGetParams struct {

	/*Cf
	  Comma seperated column family names

	*/
	Cf *string
	/*DisableSnapshot
	  When set to true, disable snapshot

	*/
	DisableSnapshot *bool
	/*Keyspace
	  The keyspace to query about

	*/
	Keyspace string
	/*SkipCorrupted
	  When set to true, skip corrupted

	*/
	SkipCorrupted *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) WithTimeout(timeout time.Duration) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) WithContext(ctx context.Context) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) WithHTTPClient(client *http.Client) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCf adds the cf to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) WithCf(cf *string) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	o.SetCf(cf)
	return o
}

// SetCf adds the cf to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) SetCf(cf *string) {
	o.Cf = cf
}

// WithDisableSnapshot adds the disableSnapshot to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) WithDisableSnapshot(disableSnapshot *bool) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	o.SetDisableSnapshot(disableSnapshot)
	return o
}

// SetDisableSnapshot adds the disableSnapshot to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) SetDisableSnapshot(disableSnapshot *bool) {
	o.DisableSnapshot = disableSnapshot
}

// WithKeyspace adds the keyspace to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) WithKeyspace(keyspace string) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	o.SetKeyspace(keyspace)
	return o
}

// SetKeyspace adds the keyspace to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) SetKeyspace(keyspace string) {
	o.Keyspace = keyspace
}

// WithSkipCorrupted adds the skipCorrupted to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) WithSkipCorrupted(skipCorrupted *bool) *StorageServiceKeyspaceScrubByKeyspaceGetParams {
	o.SetSkipCorrupted(skipCorrupted)
	return o
}

// SetSkipCorrupted adds the skipCorrupted to the storage service keyspace scrub by keyspace get params
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) SetSkipCorrupted(skipCorrupted *bool) {
	o.SkipCorrupted = skipCorrupted
}

// WriteToRequest writes these params to a swagger request
func (o *StorageServiceKeyspaceScrubByKeyspaceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cf != nil {

		// query param cf
		var qrCf string
		if o.Cf != nil {
			qrCf = *o.Cf
		}
		qCf := qrCf
		if qCf != "" {
			if err := r.SetQueryParam("cf", qCf); err != nil {
				return err
			}
		}

	}

	if o.DisableSnapshot != nil {

		// query param disable_snapshot
		var qrDisableSnapshot bool
		if o.DisableSnapshot != nil {
			qrDisableSnapshot = *o.DisableSnapshot
		}
		qDisableSnapshot := swag.FormatBool(qrDisableSnapshot)
		if qDisableSnapshot != "" {
			if err := r.SetQueryParam("disable_snapshot", qDisableSnapshot); err != nil {
				return err
			}
		}

	}

	// path param keyspace
	if err := r.SetPathParam("keyspace", o.Keyspace); err != nil {
		return err
	}

	if o.SkipCorrupted != nil {

		// query param skip_corrupted
		var qrSkipCorrupted bool
		if o.SkipCorrupted != nil {
			qrSkipCorrupted = *o.SkipCorrupted
		}
		qSkipCorrupted := swag.FormatBool(qrSkipCorrupted)
		if qSkipCorrupted != "" {
			if err := r.SetQueryParam("skip_corrupted", qSkipCorrupted); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
