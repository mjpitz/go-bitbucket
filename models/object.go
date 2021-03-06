// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// Object Base type for most resource objects. It defines the common `type` element that identifies an object's type. It also identifies the element as Swagger's `discriminator`.
//
// swagger:discriminator object type
type Object interface {
	runtime.Validatable

	// type
	// Required: true
	Type() string
	SetType(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type object struct {
	typeField string

	// object additional properties
	ObjectAdditionalProperties map[string]interface{} `json:"-"`
}

// Type gets the type of this polymorphic type
func (m *object) Type() string {
	return "object"
}

// SetType sets the type of this polymorphic type
func (m *object) SetType(val string) {
}

// UnmarshalObjectSlice unmarshals polymorphic slices of Object
func UnmarshalObjectSlice(reader io.Reader, consumer runtime.Consumer) ([]Object, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []Object
	for _, element := range elements {
		obj, err := unmarshalObject(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalObject unmarshals polymorphic Object
func UnmarshalObject(reader io.Reader, consumer runtime.Consumer) (Object, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalObject(data, consumer)
}

func unmarshalObject(data []byte, consumer runtime.Consumer) (Object, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the type property.
	var getType struct {
		Type string `json:"type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "account":
		var result Account
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "author":
		var result Author
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "base_commit":
		var result BaseCommit
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "branching_model":
		var result BranchingModel
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "branching_model_settings":
		var result BranchingModelSettings
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "branchrestriction":
		var result Branchrestriction
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "comment":
		var result Comment
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "commitstatus":
		var result Commitstatus
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "component":
		var result Component
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ddev_report":
		var result DdevReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "deploy_key":
		var result DeployKey
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "deployment":
		var result Deployment
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "deployment_environment":
		var result DeploymentEnvironment
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "deployment_release":
		var result DeploymentRelease
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "deployment_state":
		var result DeploymentState
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "deployment_state_completed_status":
		var result DeploymentStateCompletedStatus
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "deployment_variable":
		var result DeploymentVariable
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "deployments_ddev_deployment_environment":
		var result DeploymentsDdevDeploymentEnvironment
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "deployments_stg_west_deployment_environment":
		var result DeploymentsStgWestDeploymentEnvironment
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "group":
		var result Group
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "issue":
		var result Issue
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "issue_attachment":
		var result IssueAttachment
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "jira_project":
		var result JiraProject
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "jira_site":
		var result JiraSite
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "milestone":
		var result Milestone
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "object":
		var result object
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "participant":
		var result Participant
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline":
		var result Pipeline
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_build_number":
		var result PipelineBuildNumber
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_error":
		var result PipelineError
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_known_host":
		var result PipelineKnownHost
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_schedule":
		var result PipelineSchedule
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_schedule_execution":
		var result PipelineScheduleExecution
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_selector":
		var result PipelineSelector
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_ssh_key_pair":
		var result PipelineSSHKeyPair
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_ssh_public_key":
		var result PipelineSSHPublicKey
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_state":
		var result PipelineState
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_state_completed_result":
		var result PipelineStateCompletedResult
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_state_in_progress_stage":
		var result PipelineStateInProgressStage
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_step":
		var result PipelineStep
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_step_error":
		var result PipelineStepError
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_step_state":
		var result PipelineStepState
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_step_state_completed_result":
		var result PipelineStepStateCompletedResult
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_target":
		var result PipelineTarget
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_trigger":
		var result PipelineTrigger
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipeline_variable":
		var result PipelineVariable
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipelines_config":
		var result PipelinesConfig
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipelines_ddev_pipeline_step":
		var result PipelinesDdevPipelineStep
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pipelines_stg_west_pipeline_step":
		var result PipelinesStgWestPipelineStep
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "project":
		var result Project
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "pullrequest":
		var result Pullrequest
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "report":
		var result Report
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "report_annotation":
		var result ReportAnnotation
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "repository":
		var result Repository
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "snippet":
		var result Snippet
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "snippet_comment":
		var result SnippetComment
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "ssh_key":
		var result SSHKey
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "stg_west_report":
		var result StgWestReport
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "version":
		var result Version
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "webhook_subscription":
		var result WebhookSubscription
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "workspace":
		var result Workspace
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "workspace_membership":
		var result WorkspaceMembership
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid type value: %q", getType.Type)
}

// Validate validates this object
func (m *object) Validate(formats strfmt.Registry) error {
	return nil
}
