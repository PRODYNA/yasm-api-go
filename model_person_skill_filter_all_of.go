/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.1.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonSkillFilterAllOf struct for PersonSkillFilterAllOf
type PersonSkillFilterAllOf struct {
	ExperienceInMonth *MinMax `json:"experienceInMonth,omitempty"`
	// filters the last time the employee used the skill in a project
	LastAssignment *string `json:"lastAssignment,omitempty"`
}

// NewPersonSkillFilterAllOf instantiates a new PersonSkillFilterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSkillFilterAllOf() *PersonSkillFilterAllOf {
	this := PersonSkillFilterAllOf{}
	return &this
}

// NewPersonSkillFilterAllOfWithDefaults instantiates a new PersonSkillFilterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSkillFilterAllOfWithDefaults() *PersonSkillFilterAllOf {
	this := PersonSkillFilterAllOf{}
	return &this
}

// GetExperienceInMonth returns the ExperienceInMonth field value if set, zero value otherwise.
func (o *PersonSkillFilterAllOf) GetExperienceInMonth() MinMax {
	if o == nil || o.ExperienceInMonth == nil {
		var ret MinMax
		return ret
	}
	return *o.ExperienceInMonth
}

// GetExperienceInMonthOk returns a tuple with the ExperienceInMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillFilterAllOf) GetExperienceInMonthOk() (*MinMax, bool) {
	if o == nil || o.ExperienceInMonth == nil {
		return nil, false
	}
	return o.ExperienceInMonth, true
}

// HasExperienceInMonth returns a boolean if a field has been set.
func (o *PersonSkillFilterAllOf) HasExperienceInMonth() bool {
	if o != nil && o.ExperienceInMonth != nil {
		return true
	}

	return false
}

// SetExperienceInMonth gets a reference to the given MinMax and assigns it to the ExperienceInMonth field.
func (o *PersonSkillFilterAllOf) SetExperienceInMonth(v MinMax) {
	o.ExperienceInMonth = &v
}

// GetLastAssignment returns the LastAssignment field value if set, zero value otherwise.
func (o *PersonSkillFilterAllOf) GetLastAssignment() string {
	if o == nil || o.LastAssignment == nil {
		var ret string
		return ret
	}
	return *o.LastAssignment
}

// GetLastAssignmentOk returns a tuple with the LastAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSkillFilterAllOf) GetLastAssignmentOk() (*string, bool) {
	if o == nil || o.LastAssignment == nil {
		return nil, false
	}
	return o.LastAssignment, true
}

// HasLastAssignment returns a boolean if a field has been set.
func (o *PersonSkillFilterAllOf) HasLastAssignment() bool {
	if o != nil && o.LastAssignment != nil {
		return true
	}

	return false
}

// SetLastAssignment gets a reference to the given string and assigns it to the LastAssignment field.
func (o *PersonSkillFilterAllOf) SetLastAssignment(v string) {
	o.LastAssignment = &v
}

func (o PersonSkillFilterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExperienceInMonth != nil {
		toSerialize["experienceInMonth"] = o.ExperienceInMonth
	}
	if o.LastAssignment != nil {
		toSerialize["lastAssignment"] = o.LastAssignment
	}
	return json.Marshal(toSerialize)
}

type NullablePersonSkillFilterAllOf struct {
	value *PersonSkillFilterAllOf
	isSet bool
}

func (v NullablePersonSkillFilterAllOf) Get() *PersonSkillFilterAllOf {
	return v.value
}

func (v *NullablePersonSkillFilterAllOf) Set(val *PersonSkillFilterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSkillFilterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSkillFilterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSkillFilterAllOf(val *PersonSkillFilterAllOf) *NullablePersonSkillFilterAllOf {
	return &NullablePersonSkillFilterAllOf{value: val, isSet: true}
}

func (v NullablePersonSkillFilterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSkillFilterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


