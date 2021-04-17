/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: VERSION
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectParticipation struct for ProjectParticipation
type ProjectParticipation struct {
	Id string `json:"id"`
	Startdate Date `json:"startdate"`
	Enddate *Date `json:"enddate,omitempty"`
	Project Project `json:"project"`
	Description *string `json:"description,omitempty"`
	Experiences []Experience `json:"experiences"`
}

// NewProjectParticipation instantiates a new ProjectParticipation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipation(id string, startdate Date, project Project, experiences []Experience) *ProjectParticipation {
	this := ProjectParticipation{}
	this.Id = id
	this.Startdate = startdate
	this.Project = project
	this.Experiences = experiences
	return &this
}

// NewProjectParticipationWithDefaults instantiates a new ProjectParticipation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationWithDefaults() *ProjectParticipation {
	this := ProjectParticipation{}
	return &this
}

// GetId returns the Id field value
func (o *ProjectParticipation) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ProjectParticipation) SetId(v string) {
	o.Id = v
}

// GetStartdate returns the Startdate field value
func (o *ProjectParticipation) GetStartdate() Date {
	if o == nil {
		var ret Date
		return ret
	}

	return o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetStartdateOk() (*Date, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Startdate, true
}

// SetStartdate sets field value
func (o *ProjectParticipation) SetStartdate(v Date) {
	o.Startdate = v
}

// GetEnddate returns the Enddate field value if set, zero value otherwise.
func (o *ProjectParticipation) GetEnddate() Date {
	if o == nil || o.Enddate == nil {
		var ret Date
		return ret
	}
	return *o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetEnddateOk() (*Date, bool) {
	if o == nil || o.Enddate == nil {
		return nil, false
	}
	return o.Enddate, true
}

// HasEnddate returns a boolean if a field has been set.
func (o *ProjectParticipation) HasEnddate() bool {
	if o != nil && o.Enddate != nil {
		return true
	}

	return false
}

// SetEnddate gets a reference to the given Date and assigns it to the Enddate field.
func (o *ProjectParticipation) SetEnddate(v Date) {
	o.Enddate = &v
}

// GetProject returns the Project field value
func (o *ProjectParticipation) GetProject() Project {
	if o == nil {
		var ret Project
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetProjectOk() (*Project, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ProjectParticipation) SetProject(v Project) {
	o.Project = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectParticipation) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectParticipation) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectParticipation) SetDescription(v string) {
	o.Description = &v
}

// GetExperiences returns the Experiences field value
func (o *ProjectParticipation) GetExperiences() []Experience {
	if o == nil {
		var ret []Experience
		return ret
	}

	return o.Experiences
}

// GetExperiencesOk returns a tuple with the Experiences field value
// and a boolean to check if the value has been set.
func (o *ProjectParticipation) GetExperiencesOk() (*[]Experience, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Experiences, true
}

// SetExperiences sets field value
func (o *ProjectParticipation) SetExperiences(v []Experience) {
	o.Experiences = v
}

func (o ProjectParticipation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
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

type NullableProjectParticipation struct {
	value *ProjectParticipation
	isSet bool
}

func (v NullableProjectParticipation) Get() *ProjectParticipation {
	return v.value
}

func (v *NullableProjectParticipation) Set(val *ProjectParticipation) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectParticipation) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectParticipation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectParticipation(val *ProjectParticipation) *NullableProjectParticipation {
	return &NullableProjectParticipation{value: val, isSet: true}
}

func (v NullableProjectParticipation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectParticipation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


