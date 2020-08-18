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

// NewPostAddonLinkersLinkerKeyValuesParams creates a new PostAddonLinkersLinkerKeyValuesParams object
// with the default values initialized.
func NewPostAddonLinkersLinkerKeyValuesParams() *PostAddonLinkersLinkerKeyValuesParams {
	var ()
	return &PostAddonLinkersLinkerKeyValuesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAddonLinkersLinkerKeyValuesParamsWithTimeout creates a new PostAddonLinkersLinkerKeyValuesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAddonLinkersLinkerKeyValuesParamsWithTimeout(timeout time.Duration) *PostAddonLinkersLinkerKeyValuesParams {
	var ()
	return &PostAddonLinkersLinkerKeyValuesParams{

		timeout: timeout,
	}
}

// NewPostAddonLinkersLinkerKeyValuesParamsWithContext creates a new PostAddonLinkersLinkerKeyValuesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAddonLinkersLinkerKeyValuesParamsWithContext(ctx context.Context) *PostAddonLinkersLinkerKeyValuesParams {
	var ()
	return &PostAddonLinkersLinkerKeyValuesParams{

		Context: ctx,
	}
}

// NewPostAddonLinkersLinkerKeyValuesParamsWithHTTPClient creates a new PostAddonLinkersLinkerKeyValuesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAddonLinkersLinkerKeyValuesParamsWithHTTPClient(client *http.Client) *PostAddonLinkersLinkerKeyValuesParams {
	var ()
	return &PostAddonLinkersLinkerKeyValuesParams{
		HTTPClient: client,
	}
}

/*PostAddonLinkersLinkerKeyValuesParams contains all the parameters to send to the API endpoint
for the post addon linkers linker key values operation typically these are written to a http.Request
*/
type PostAddonLinkersLinkerKeyValuesParams struct {

	/*LinkerKey*/
	LinkerKey string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post addon linkers linker key values params
func (o *PostAddonLinkersLinkerKeyValuesParams) WithTimeout(timeout time.Duration) *PostAddonLinkersLinkerKeyValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post addon linkers linker key values params
func (o *PostAddonLinkersLinkerKeyValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post addon linkers linker key values params
func (o *PostAddonLinkersLinkerKeyValuesParams) WithContext(ctx context.Context) *PostAddonLinkersLinkerKeyValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post addon linkers linker key values params
func (o *PostAddonLinkersLinkerKeyValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post addon linkers linker key values params
func (o *PostAddonLinkersLinkerKeyValuesParams) WithHTTPClient(client *http.Client) *PostAddonLinkersLinkerKeyValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post addon linkers linker key values params
func (o *PostAddonLinkersLinkerKeyValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLinkerKey adds the linkerKey to the post addon linkers linker key values params
func (o *PostAddonLinkersLinkerKeyValuesParams) WithLinkerKey(linkerKey string) *PostAddonLinkersLinkerKeyValuesParams {
	o.SetLinkerKey(linkerKey)
	return o
}

// SetLinkerKey adds the linkerKey to the post addon linkers linker key values params
func (o *PostAddonLinkersLinkerKeyValuesParams) SetLinkerKey(linkerKey string) {
	o.LinkerKey = linkerKey
}

// WriteToRequest writes these params to a swagger request
func (o *PostAddonLinkersLinkerKeyValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
