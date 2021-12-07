/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.9.6
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// ProjectStatus the model 'ProjectStatus'
type ProjectStatus string

// List of ProjectStatus
const (
	PLANNING ProjectStatus = "PLANNING"
	ACTIVE ProjectStatus = "ACTIVE"
	DONE ProjectStatus = "DONE"
)

// All allowed values of ProjectStatus enum
var AllowedProjectStatusEnumValues = []ProjectStatus{
	"PLANNING",
	"ACTIVE",
	"DONE",
}

func (v *ProjectStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProjectStatus(value)
	for _, existing := range AllowedProjectStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProjectStatus", value)
}

// NewProjectStatusFromValue returns a pointer to a valid ProjectStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProjectStatusFromValue(v string) (*ProjectStatus, error) {
	ev := ProjectStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProjectStatus: valid values are %v", v, AllowedProjectStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProjectStatus) IsValid() bool {
	for _, existing := range AllowedProjectStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProjectStatus value
func (v ProjectStatus) Ptr() *ProjectStatus {
	return &v
}

type NullableProjectStatus struct {
	value *ProjectStatus
	isSet bool
}

func (v NullableProjectStatus) Get() *ProjectStatus {
	return v.value
}

func (v *NullableProjectStatus) Set(val *ProjectStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectStatus(val *ProjectStatus) *NullableProjectStatus {
	return &NullableProjectStatus{value: val, isSet: true}
}

func (v NullableProjectStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

