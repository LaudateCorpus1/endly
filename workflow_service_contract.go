package endly

import (
	"errors"
	"github.com/viant/toolbox/url"
)

//RunRequest represents workflow run request
type RunRequest struct {
	EnableLogging     bool                   `description:"flag to enable logging"`
	LoggingDirectory  string                 `description:"log directory"`
	WorkflowURL               string                 `description:"workflow URL if workflow is not found in the registry, it is loaded"`
	Name              string                 `required:"true" description:"name defined in workflow document"`
	Params            map[string]interface{} `description:"workflow parameters, accessibly by paras.[Key], if PublishParameters is set, all parameters are place in context.state"`
	Tasks             string                 `required:"true" description:"coma separated task list or '*'to run all tasks sequencialy"` //tasks to run with coma separated list or '*', or empty string for all tasks
	TagIDs            string                 `description:"coma separated TagID list, if present in a task, only matched runs, other task run as normal"`
	PublishParameters bool                   `description:"flag to publish parameters directly into context state"`
	Async             bool                   `description:"flag to run it asynchronously. Do not set it your self runner sets the flag for the first workflow"`
}

//RunResponse represents workflow run response
type RunResponse struct {
	Data      map[string]interface{} //  data populated by  .Post variable section.
	SessionID string                 //session id
}

//RegisterRequest represents workflow register request
type RegisterRequest struct {
	*Workflow
}

//RegisterResponse represents workflow register response
type RegisterResponse struct {
	Source *url.Resource
}

// LoadRequest represents workflow load request from the specified source
type LoadRequest struct {
	Source *url.Resource
}

// LoadResponse represents loaded workflow
type LoadResponse struct {
	*Workflow
}

// SwitchCase represent matching candidate case
type SwitchCase struct {
	*ActionRequest `description:"action to run if matched"`
	Task                 string      `description:"task to run if matched"`
	Value                interface{} `required:"true" description:"matching sourceKey value"`
}

// SwitchRequest represent switch action request
type SwitchRequest struct {
	SourceKey string        `required:"true" description:"sourceKey for matching value"`
	Cases     []*SwitchCase `required:"true" description:"matching value cases"`
	Default   *SwitchCase   `description:"in case no value was match case"`
}

//Match matches source with supplied action request.
func (r *SwitchRequest) Match(source interface{}) *SwitchCase {
	for _, switchCase := range r.Cases {
		if switchCase.Value == source {
			return switchCase
		}
	}
	return r.Default
}

// SwitchResponse represents actual action or task response
type SwitchResponse interface{}

//Validate checks if workflow is valid
func (r *SwitchRequest) Validate() error {
	if r.SourceKey == "" {
		return errors.New("sourceKey was empty")
	}
	if len(r.Cases) == 0 {
		return errors.New("cases were empty")
	}
	for _, matchingCase := range r.Cases {
		if matchingCase.Value == nil {
			return errors.New("cases.value was empty")
		}
	}
	return nil
}

// GotoRequest represents goto task action, this request will terminate current task execution to switch to specified task
type GotoRequest struct {
	Task string
}

// GotoResponse represents workflow task response
type GotoResponse interface{}

// ExitRequest represents workflow exit request, to exit a caller workflow
type ExitRequest struct {
	Source *url.Resource
}

// ExitResponse represents workflow exit response
type ExitResponse struct{}

// FailRequest represents fail request
type FailRequest struct {
	Message string
}

// FailResponse represents workflow exit response
type FailResponse struct{}
