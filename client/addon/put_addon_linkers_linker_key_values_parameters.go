// Code generated by go-swagger; DO NOT EDIT.

package addon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPutAddonLinkersLinkerKeyValuesParams creates a new PutAddonLinkersLinkerKeyValuesParams object
// with the default values initialized.
func NewPutAddonLinkersLinkerKeyValuesParams() *PutAddonLinkersLinkerKeyValuesParams {
	var ()
	return &PutAddonLinkersLinkerKeyValuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAddonLinkersLinkerKeyValuesParamsWithTimeout creates a new PutAddonLinkersLinkerKeyValuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAddonLinkersLinkerKeyValuesParamsWithTimeout(timeout time.Duration) *PutAddonLinkersLinkerKeyValuesParams {
	var ()
	return &PutAddonLinkersLinkerKeyValuesParams{

		timeout: timeout,
	}
}

// NewPutAddonLinkersLinkerKeyValuesParamsWithContext creates a new PutAddonLinkersLinkerKeyValuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAddonLinkersLinkerKeyValuesParamsWithContext(ctx context.Context) *PutAddonLinkersLinkerKeyValuesParams {
	var ()
	return &PutAddonLinkersLinkerKeyValuesParams{

		Context: ctx,
	}
}

// NewPutAddonLinkersLinkerKeyValuesParamsWithHTTPClient creates a new PutAddonLinkersLinkerKeyValuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAddonLinkersLinkerKeyValuesParamsWithHTTPClient(client *http.Client) *PutAddonLinkersLinkerKeyValuesParams {
	var ()
	return &PutAddonLinkersLinkerKeyValuesParams{
		HTTPClient: client,
	}
}

/*PutAddonLinkersLinkerKeyValuesParams contains all the parameters to send to the API endpoint
for the put addon linkers linker key values operation typically these are written to a http.Request
*/
type PutAddonLinkersLinkerKeyValuesParams struct {

	/*LinkerKey*/
	LinkerKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put addon linkers linker key values params
func (o *PutAddonLinkersLinkerKeyValuesParams) WithTimeout(timeout time.Duration) *PutAddonLinkersLinkerKeyValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put addon linkers linker key values params
func (o *PutAddonLinkersLinkerKeyValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put addon linkers linker key values params
func (o *PutAddonLinkersLinkerKeyValuesParams) WithContext(ctx context.Context) *PutAddonLinkersLinkerKeyValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put addon linkers linker key values params
func (o *PutAddonLinkersLinkerKeyValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put addon linkers linker key values params
func (o *PutAddonLinkersLinkerKeyValuesParams) WithHTTPClient(client *http.Client) *PutAddonLinkersLinkerKeyValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put addon linkers linker key values params
func (o *PutAddonLinkersLinkerKeyValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLinkerKey adds the linkerKey to the put addon linkers linker key values params
func (o *PutAddonLinkersLinkerKeyValuesParams) WithLinkerKey(linkerKey string) *PutAddonLinkersLinkerKeyValuesParams {
	o.SetLinkerKey(linkerKey)
	return o
}

// SetLinkerKey adds the linkerKey to the put addon linkers linker key values params
func (o *PutAddonLinkersLinkerKeyValuesParams) SetLinkerKey(linkerKey string) {
	o.LinkerKey = linkerKey
}

// WriteToRequest writes these params to a swagger request
func (o *PutAddonLinkersLinkerKeyValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param linker_key
	if err := r.SetPathParam("linker_key", o.LinkerKey); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
