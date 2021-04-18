/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.5.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedProjectsAllOf struct for PagedProjectsAllOf
type PagedProjectsAllOf struct {
	Projects *[]Project `json:"projects,omitempty"`
}

// NewPagedProjectsAllOf instantiates a new PagedProjectsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedProjectsAllOf() *PagedProjectsAllOf {
	this := PagedProjectsAllOf{}
	return &this
}

// NewPagedProjectsAllOfWithDefaults instantiates a new PagedProjectsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedProjectsAllOfWithDefaults() *PagedProjectsAllOf {
	this := PagedProjectsAllOf{}
	return &this
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PagedProjectsAllOf) GetProjects() []Project {
	if o == nil || o.Projects == nil {
		var ret []Project
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedProjectsAllOf) GetProjectsOk() (*[]Project, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PagedProjectsAllOf) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []Project and assigns it to the Projects field.
func (o *PagedProjectsAllOf) SetProjects(v []Project) {
	o.Projects = &v
}

func (o PagedProjectsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	return json.Marshal(toSerialize)
}

type NullablePagedProjectsAllOf struct {
	value *PagedProjectsAllOf
	isSet bool
}

func (v NullablePagedProjectsAllOf) Get() *PagedProjectsAllOf {
	return v.value
}

func (v *NullablePagedProjectsAllOf) Set(val *PagedProjectsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedProjectsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedProjectsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedProjectsAllOf(val *PagedProjectsAllOf) *NullablePagedProjectsAllOf {
	return &NullablePagedProjectsAllOf{value: val, isSet: true}
}

func (v NullablePagedProjectsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedProjectsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

