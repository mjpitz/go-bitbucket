// Code generated by go-swagger; DO NOT EDIT.

package webhooks

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

// NewGetHookEventsSubjectTypeParams creates a new GetHookEventsSubjectTypeParams object
// with the default values initialized.
func NewGetHookEventsSubjectTypeParams() *GetHookEventsSubjectTypeParams {
	var ()
	return &GetHookEventsSubjectTypeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetHookEventsSubjectTypeParamsWithTimeout creates a new GetHookEventsSubjectTypeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetHookEventsSubjectTypeParamsWithTimeout(timeout time.Duration) *GetHookEventsSubjectTypeParams {
	var ()
	return &GetHookEventsSubjectTypeParams{

		timeout: timeout,
	}
}

// NewGetHookEventsSubjectTypeParamsWithContext creates a new GetHookEventsSubjectTypeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetHookEventsSubjectTypeParamsWithContext(ctx context.Context) *GetHookEventsSubjectTypeParams {
	var ()
	return &GetHookEventsSubjectTypeParams{

		Context: ctx,
	}
}

// NewGetHookEventsSubjectTypeParamsWithHTTPClient creates a new GetHookEventsSubjectTypeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetHookEventsSubjectTypeParamsWithHTTPClient(client *http.Client) *GetHookEventsSubjectTypeParams {
	var ()
	return &GetHookEventsSubjectTypeParams{
		HTTPClient: client,
	}
}

/*GetHookEventsSubjectTypeParams contains all the parameters to send to the API endpoint
for the get hook events subject type operation typically these are written to a http.Request
*/
type GetHookEventsSubjectTypeParams struct {

	/*SubjectType
	  A resource or subject type.

	*/
	SubjectType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get hook events subject type params
func (o *GetHookEventsSubjectTypeParams) WithTimeout(timeout time.Duration) *GetHookEventsSubjectTypeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get hook events subject type params
func (o *GetHookEventsSubjectTypeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get hook events subject type params
func (o *GetHookEventsSubjectTypeParams) WithContext(ctx context.Context) *GetHookEventsSubjectTypeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get hook events subject type params
func (o *GetHookEventsSubjectTypeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get hook events subject type params
func (o *GetHookEventsSubjectTypeParams) WithHTTPClient(client *http.Client) *GetHookEventsSubjectTypeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get hook events subject type params
func (o *GetHookEventsSubjectTypeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubjectType adds the subjectType to the get hook events subject type params
func (o *GetHookEventsSubjectTypeParams) WithSubjectType(subjectType string) *GetHookEventsSubjectTypeParams {
	o.SetSubjectType(subjectType)
	return o
}

// SetSubjectType adds the subjectType to the get hook events subject type params
func (o *GetHookEventsSubjectTypeParams) SetSubjectType(subjectType string) {
	o.SubjectType = subjectType
}

// WriteToRequest writes these params to a swagger request
func (o *GetHookEventsSubjectTypeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param subject_type
	if err := r.SetPathParam("subject_type", o.SubjectType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
