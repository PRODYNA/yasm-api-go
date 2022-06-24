/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectSearchSkillsInner struct for ProjectSearchSkillsInner
type ProjectSearchSkillsInner struct {
	Id string `json:"id"`
}

// NewProjectSearchSkillsInner instantiates a new ProjectSearchSkillsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectSearchSkillsInner(id string) *ProjectSearchSkillsInner {
	this := ProjectSearchSkillsInner{}
	this.Id = id
	return &this
}

// NewProjectSearchSkillsInnerWithDefaults instantiates a new ProjectSearchSkillsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectSearchSkillsInnerWithDefaults() *ProjectSearchSkillsInner {
	this := ProjectSearchSkillsInner{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectSearchSkillsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectSearchSkillsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectSearchSkillsInner) SetId(v string) {
	o.Id = v
}

func (o ProjectSearchSkillsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableProjectSearchSkillsInner struct {
	value *ProjectSearchSkillsInner
	isSet bool
}

func (v NullableProjectSearchSkillsInner) Get() *ProjectSearchSkillsInner {
	return v.value
}

func (v *NullableProjectSearchSkillsInner) Set(val *ProjectSearchSkillsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectSearchSkillsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectSearchSkillsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectSearchSkillsInner(val *ProjectSearchSkillsInner) *NullableProjectSearchSkillsInner {
	return &NullableProjectSearchSkillsInner{value: val, isSet: true}
}

func (v NullableProjectSearchSkillsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectSearchSkillsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

