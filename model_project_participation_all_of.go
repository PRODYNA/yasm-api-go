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

// ProjectParticipationAllOf struct for ProjectParticipationAllOf
type ProjectParticipationAllOf struct {
	Startdate string `json:"startdate"`
	Enddate *string `json:"enddate,omitempty"`
	Project Project `json:"project"`
	Description *string `json:"description,omitempty"`
	Experiences []Experience `json:"experiences"`
}

// NewProjectParticipationAllOf instantiates a new ProjectParticipationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipationAllOf(startdate string, project Project, experiences []Experience) *ProjectParticipationAllOf {
	this := ProjectParticipationAllOf{}
	this.Startdate = startdate
	this.Project = project
	this.Experiences = experiences
	return &this
}

// NewProjectParticipationAllOfWithDefaults instantiates a new ProjectParticipationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationAllOfWithDefaults() *ProjectParticipationAllOf {
	this := ProjectParticipationAllOf{}
	return &this
}

// GetStartdate returns the Startdate field value
func (o *ProjectParticipationAllOf) GetStartdate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetStartdateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Startdate, true
}

// SetStartdate sets field value
func (o *ProjectParticipationAllOf) SetStartdate(v string) {
	o.Startdate = v
}

// GetEnddate returns the Enddate field value if set, zero value otherwise.
func (o *ProjectParticipationAllOf) GetEnddate() string {
	if o == nil || o.Enddate == nil {
		var ret string
		return ret
	}
	return *o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetEnddateOk() (*string, bool) {
	if o == nil || o.Enddate == nil {
		return nil, false
	}
	return o.Enddate, true
}

// HasEnddate returns a boolean if a field has been set.
func (o *ProjectParticipationAllOf) HasEnddate() bool {
	if o != nil && o.Enddate != nil {
		return true
	}

	return false
}

// SetEnddate gets a reference to the given string and assigns it to the Enddate field.
func (o *ProjectParticipationAllOf) SetEnddate(v string) {
	o.Enddate = &v
}

// GetProject returns the Project field value
func (o *ProjectParticipationAllOf) GetProject() Project {
	if o == nil {
		var ret Project
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetProjectOk() (*Project, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ProjectParticipationAllOf) SetProject(v Project) {
	o.Project = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectParticipationAllOf) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectParticipationAllOf) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectParticipationAllOf) SetDescription(v string) {
	o.Description = &v
}

// GetExperiences returns the Experiences field value
func (o *ProjectParticipationAllOf) GetExperiences() []Experience {
	if o == nil {
		var ret []Experience
		return ret
	}

	return o.Experiences
}

// GetExperiencesOk returns a tuple with the Experiences field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipationAllOf) GetExperiencesOk() (*[]Experience, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Experiences, true
}

// SetExperiences sets field value
func (o *ProjectParticipationAllOf) SetExperiences(v []Experience) {
	o.Experiences = v
}

func (o ProjectParticipationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["startdate"] = o.Startdate
	}
	if o.Enddate != nil {
		toSerialize["enddate"] = o.Enddate
	}
	if true {
		toSerialize["project"] = o.Project
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["experiences"] = o.Experiences
	}
	return json.Marshal(toSerialize)
}

type NullableProjectParticipationAllOf struct {
	value *ProjectParticipationAllOf
	isSet bool
}

func (v NullableProjectParticipationAllOf) Get() *ProjectParticipationAllOf {
	return v.value
}

func (v *NullableProjectParticipationAllOf) Set(val *ProjectParticipationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectParticipationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectParticipationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectParticipationAllOf(val *ProjectParticipationAllOf) *NullableProjectParticipationAllOf {
	return &NullableProjectParticipationAllOf{value: val, isSet: true}
}

func (v NullableProjectParticipationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectParticipationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

