// Code generated by go-swagger; DO NOT EDIT.

package issue_tracker

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
	"github.com/go-openapi/swag"
)

// NewGetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams creates a new GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams object
// with the default values initialized.
func NewGetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams() *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParamsWithTimeout creates a new GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParamsWithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams{

		timeout: timeout,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParamsWithContext creates a new GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParamsWithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams{

		Context: ctx,
	}
}

// NewGetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParamsWithHTTPClient creates a new GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParamsWithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	var ()
	return &GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams{
		HTTPClient: client,
	}
}

/*GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams contains all the parameters to send to the API endpoint
for the get repositories workspace repo slug milestones milestone ID operation typically these are written to a http.Request
*/
type GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams struct {

	/*MilestoneID
	  The milestone's id

	*/
	MilestoneID int64
	/*RepoSlug
	  This can either be the repository slug or the UUID of the repository,
	surrounded by curly-braces, for example: `{repository UUID}`.


	*/
	RepoSlug string
	/*Workspace
	  This can either be the workspace ID (slug) or the workspace UUID
	surrounded by curly-braces, for example: `{workspace UUID}`.


	*/
	Workspace string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) WithTimeout(timeout time.Duration) *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) WithContext(ctx context.Context) *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) WithHTTPClient(client *http.Client) *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMilestoneID adds the milestoneID to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) WithMilestoneID(milestoneID int64) *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	o.SetMilestoneID(milestoneID)
	return o
}

// SetMilestoneID adds the milestoneId to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) SetMilestoneID(milestoneID int64) {
	o.MilestoneID = milestoneID
}

// WithRepoSlug adds the repoSlug to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) WithRepoSlug(repoSlug string) *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) WithWorkspace(workspace string) *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the get repositories workspace repo slug milestones milestone ID params
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *GetRepositoriesWorkspaceRepoSlugMilestonesMilestoneIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param milestone_id
	if err := r.SetPathParam("milestone_id", swag.FormatInt64(o.MilestoneID)); err != nil {
		return err
	}

	// path param repo_slug
	if err := r.SetPathParam("repo_slug", o.RepoSlug); err != nil {
		return err
	}

	// path param workspace
	if err := r.SetPathParam("workspace", o.Workspace); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
