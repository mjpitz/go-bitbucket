// Code generated by go-swagger; DO NOT EDIT.

package source

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

// NewPostRepositoriesWorkspaceRepoSlugSrcParams creates a new PostRepositoriesWorkspaceRepoSlugSrcParams object
// with the default values initialized.
func NewPostRepositoriesWorkspaceRepoSlugSrcParams() *PostRepositoriesWorkspaceRepoSlugSrcParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugSrcParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugSrcParamsWithTimeout creates a new PostRepositoriesWorkspaceRepoSlugSrcParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRepositoriesWorkspaceRepoSlugSrcParamsWithTimeout(timeout time.Duration) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugSrcParams{

		timeout: timeout,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugSrcParamsWithContext creates a new PostRepositoriesWorkspaceRepoSlugSrcParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRepositoriesWorkspaceRepoSlugSrcParamsWithContext(ctx context.Context) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugSrcParams{

		Context: ctx,
	}
}

// NewPostRepositoriesWorkspaceRepoSlugSrcParamsWithHTTPClient creates a new PostRepositoriesWorkspaceRepoSlugSrcParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRepositoriesWorkspaceRepoSlugSrcParamsWithHTTPClient(client *http.Client) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	var ()
	return &PostRepositoriesWorkspaceRepoSlugSrcParams{
		HTTPClient: client,
	}
}

/*PostRepositoriesWorkspaceRepoSlugSrcParams contains all the parameters to send to the API endpoint
for the post repositories workspace repo slug src operation typically these are written to a http.Request
*/
type PostRepositoriesWorkspaceRepoSlugSrcParams struct {

	/*Author

	The raw string to be used as the new commit's author.
	This string follows the format
	`Erik van Zijst <evzijst@atlassian.com>`.

	When omitted, Bitbucket uses the authenticated user's
	full/display name and primary email address. Commits cannot
	be created anonymously.

	*/
	Author *string
	/*Branch

	The name of the branch that the new commit should be
	created on. When omitted, the commit will be created on top
	of the main branch and will become the main branch's new
	head.

	When a branch name is provided that already exists in the
	repo, then the commit will be created on top of that
	branch. In this case, *if* a parent SHA1 was also provided,
	then it is asserted that the parent is the branch's
	tip/HEAD at the time the request is made. When this is not
	the case, a 409 is returned.

	This API cannot be used to create new anonymous heads in
	Mercurial repositories.

	When a new branch name is specified (that does not already
	exist in the repo), and no parent SHA1s are provided, then
	the new commit will inherit from the current main branch's
	tip/HEAD commit, but not advance the main branch. The new
	commit will be the new branch. When the request *also*
	specifies a parent SHA1, then the new commit and branch
	are created directly on top of the parent commit,
	regardless of the state of the main branch.

	When a branch name is not specified, but a parent SHA1 is
	provided, then Bitbucket asserts that it represents the
	main branch's current HEAD/tip, or a 409 is returned.

	When a branch name is not specified and the repo is empty,
	the new commit will become the repo's root commit and will
	be on the main branch.

	When a branch name is specified and the repo is empty, the
	new commit will become the repo's root commit and also
	define the repo's main branch going forward.

	This API cannot be used to create additional root commits
	in non-empty repos.

	The branch field cannot be repeated.

	As a side effect, this API can be used to create a new
	branch without modifying any files, by specifying a new
	branch name in this field, together with `parents`, but
	omitting the `files` fields, while not sending any files.
	This will create a new commit and branch with the same
	contents as the first parent. The diff of this commit
	against its first parent will be empty.


	*/
	Branch *string
	/*Files

	Optional field that declares the files that the request is
	manipulating. When adding a new file to a repo, or when
	overwriting an existing file, the client can just upload
	the full contents of the file in a normal form field and
	the use of this `files` meta data field is redundant.
	However, when the `files` field contains a file path that
	does not have a corresponding, identically-named form
	field, then Bitbucket interprets that as the client wanting
	to replace the named file with the null set and the file is
	deleted instead.

	Paths in the repo that are referenced in neither files nor
	an individual file field, remain unchanged and carry over
	from the parent to the new commit.

	This API does not support renaming as an explicit feature.
	To rename a file, simply delete it and recreate it under
	the new name in the same commit.


	*/
	Files *string
	/*Message
	  The commit message. When omitted, Bitbucket uses a canned string.

	*/
	Message *string
	/*Parents

	A comma-separated list of SHA1s of the commits that should
	be the parents of the newly created commit.

	When omitted, the new commit will inherit from and become
	a child of the main branch's tip/HEAD commit.

	When more than one SHA1 is provided, the first SHA1
	identifies the commit from which the content will be
	inherited.

	When more than 2 parents are provided on a Mercurial repo,
	a 400 is returned as Mercurial does not support "octopus
	merges".

	*/
	Parents *string
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

// WithTimeout adds the timeout to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithTimeout(timeout time.Duration) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithContext(ctx context.Context) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithHTTPClient(client *http.Client) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthor adds the author to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithAuthor(author *string) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetAuthor(author)
	return o
}

// SetAuthor adds the author to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetAuthor(author *string) {
	o.Author = author
}

// WithBranch adds the branch to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithBranch(branch *string) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetBranch(branch *string) {
	o.Branch = branch
}

// WithFiles adds the files to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithFiles(files *string) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetFiles(files)
	return o
}

// SetFiles adds the files to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetFiles(files *string) {
	o.Files = files
}

// WithMessage adds the message to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithMessage(message *string) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetMessage(message)
	return o
}

// SetMessage adds the message to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetMessage(message *string) {
	o.Message = message
}

// WithParents adds the parents to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithParents(parents *string) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetParents(parents)
	return o
}

// SetParents adds the parents to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetParents(parents *string) {
	o.Parents = parents
}

// WithRepoSlug adds the repoSlug to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithRepoSlug(repoSlug string) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetRepoSlug(repoSlug)
	return o
}

// SetRepoSlug adds the repoSlug to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetRepoSlug(repoSlug string) {
	o.RepoSlug = repoSlug
}

// WithWorkspace adds the workspace to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WithWorkspace(workspace string) *PostRepositoriesWorkspaceRepoSlugSrcParams {
	o.SetWorkspace(workspace)
	return o
}

// SetWorkspace adds the workspace to the post repositories workspace repo slug src params
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) SetWorkspace(workspace string) {
	o.Workspace = workspace
}

// WriteToRequest writes these params to a swagger request
func (o *PostRepositoriesWorkspaceRepoSlugSrcParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Author != nil {

		// query param author
		var qrAuthor string
		if o.Author != nil {
			qrAuthor = *o.Author
		}
		qAuthor := qrAuthor
		if qAuthor != "" {
			if err := r.SetQueryParam("author", qAuthor); err != nil {
				return err
			}
		}

	}

	if o.Branch != nil {

		// query param branch
		var qrBranch string
		if o.Branch != nil {
			qrBranch = *o.Branch
		}
		qBranch := qrBranch
		if qBranch != "" {
			if err := r.SetQueryParam("branch", qBranch); err != nil {
				return err
			}
		}

	}

	if o.Files != nil {

		// query param files
		var qrFiles string
		if o.Files != nil {
			qrFiles = *o.Files
		}
		qFiles := qrFiles
		if qFiles != "" {
			if err := r.SetQueryParam("files", qFiles); err != nil {
				return err
			}
		}

	}

	if o.Message != nil {

		// query param message
		var qrMessage string
		if o.Message != nil {
			qrMessage = *o.Message
		}
		qMessage := qrMessage
		if qMessage != "" {
			if err := r.SetQueryParam("message", qMessage); err != nil {
				return err
			}
		}

	}

	if o.Parents != nil {

		// query param parents
		var qrParents string
		if o.Parents != nil {
			qrParents = *o.Parents
		}
		qParents := qrParents
		if qParents != "" {
			if err := r.SetQueryParam("parents", qParents); err != nil {
				return err
			}
		}

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
