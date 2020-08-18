// Code generated by go-swagger; DO NOT EDIT.

package deployments

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

// NewGetDeploymentsForRepositoryParams creates a new GetDeploymentsForRepositoryParams object
// with the default values initialized.
func NewGetDeploymentsForRepositoryParams() *GetDeploymentsForRepositoryParams {
	var ()
	return &GetDeploymentsForRepositoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentsForRepositoryParamsWithTimeout creates a new GetDeploymentsForRepositoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentsForRepositoryParamsWithTimeout(timeout time.Duration) *GetDeploymentsForRepositoryParams {
	var ()
	return &GetDeploymentsForRepositoryParams{

		timeout: timeout,
	}
}

// NewGetDeploymentsForRepositoryParamsWithContext creates a new GetDeploymentsForRepositoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentsForRepositoryParamsWithContext(ctx context.Context) *GetDeploymentsForRepositoryParams {
	var ()
	return &GetDeploymentsForRepositoryParams{

		Context: ctx,
	}
}

// NewGetDeploymentsForRepositoryParamsWithHTTPClient creates a new GetDeploymentsForRepositoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentsForRepositoryParamsWithHTTPClient(client *http.Client) *GetDeploymentsForRepositoryParams {
	var ()
	return &GetDeploymentsForRepositoryParams{
		HTTPClient: client,
	}
}

/*GetDeploymentsForRepositoryParams contains all the parameters to send to the API endpoint
for the get deployments for repository operation typically these are written to a http.Request
*/
type GetDeploymentsForRepositoryParams struct {

	/*RepoSlug
	  The repository.

	*/
	RepoSlug string
	/*Username
	  The account.

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) WithTimeout(timeout time.Duration) *GetDeploymentsForRepositoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) WithContext(ctx context.Context) *GetDeploymentsForRepositoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) WithHTTPClient(client *http.Client) *GetDeploymentsForRepositoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRepoSlug adds the repoSlug to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) WithRepoSlug(repoSlug string) *GetDeploymentsForRepositoryParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithUsername adds the username to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) WithUsername(username string) *GetDeploymentsForRepositoryParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get deployments for repository params
func (o *GetDeploymentsForRepositoryParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentsForRepositoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}