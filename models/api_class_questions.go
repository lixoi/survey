// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// APIClassQuestions api class questions
//
// swagger:model apiClassQuestions
type APIClassQuestions string

func NewAPIClassQuestions(value APIClassQuestions) *APIClassQuestions {
	return &value
}

// Pointer returns a pointer to a freshly-allocated APIClassQuestions.
func (m APIClassQuestions) Pointer() *APIClassQuestions {
	return &m
}

const (

	// APIClassQuestionsUNKNOWNQUESTIONSCLASS captures enum value "UNKNOWN_QUESTIONS_CLASS"
	APIClassQuestionsUNKNOWNQUESTIONSCLASS APIClassQuestions = "UNKNOWN_QUESTIONS_CLASS"

	// APIClassQuestionsLINUXQUESTIONS captures enum value "LINUX_QUESTIONS"
	APIClassQuestionsLINUXQUESTIONS APIClassQuestions = "LINUX_QUESTIONS"

	// APIClassQuestionsK8SQUESTIONS captures enum value "K8S_QUESTIONS"
	APIClassQuestionsK8SQUESTIONS APIClassQuestions = "K8S_QUESTIONS"

	// APIClassQuestionsNETWORKQUESTIONS captures enum value "NETWORK_QUESTIONS"
	APIClassQuestionsNETWORKQUESTIONS APIClassQuestions = "NETWORK_QUESTIONS"

	// APIClassQuestionsSECURITYQUESTIONS captures enum value "SECURITY_QUESTIONS"
	APIClassQuestionsSECURITYQUESTIONS APIClassQuestions = "SECURITY_QUESTIONS"

	// APIClassQuestionsCONTAINERQUESTIONS captures enum value "CONTAINER_QUESTIONS"
	APIClassQuestionsCONTAINERQUESTIONS APIClassQuestions = "CONTAINER_QUESTIONS"

	// APIClassQuestionsDEVELOPERUESTIONS captures enum value "DEVELOPER_UESTIONS"
	APIClassQuestionsDEVELOPERUESTIONS APIClassQuestions = "DEVELOPER_UESTIONS"
)

// for schema
var apiClassQuestionsEnum []interface{}

func init() {
	var res []APIClassQuestions
	if err := json.Unmarshal([]byte(`["UNKNOWN_QUESTIONS_CLASS","LINUX_QUESTIONS","K8S_QUESTIONS","NETWORK_QUESTIONS","SECURITY_QUESTIONS","CONTAINER_QUESTIONS","DEVELOPER_UESTIONS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apiClassQuestionsEnum = append(apiClassQuestionsEnum, v)
	}
}

func (m APIClassQuestions) validateAPIClassQuestionsEnum(path, location string, value APIClassQuestions) error {
	if err := validate.EnumCase(path, location, value, apiClassQuestionsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this api class questions
func (m APIClassQuestions) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAPIClassQuestionsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this api class questions based on context it is used
func (m APIClassQuestions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
