/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.7.6
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ProjectDetails struct for ProjectDetails
type ProjectDetails struct {
	Project *Project `json:"project,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Persons *[]Person `json:"persons,omitempty"`
}

// NewProjectDetails instantiates a new ProjectDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDetails() *ProjectDetails {
	this := ProjectDetails{}
	return &this
}

// NewProjectDetailsWithDefaults instantiates a new ProjectDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDetailsWithDefaults() *ProjectDetails {
	this := ProjectDetails{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectDetails) GetProject() Project {
	if o == nil || o.Project == nil {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetProjectOk() (*Project, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectDetails) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *ProjectDetails) SetProject(v Project) {
	o.Project = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ProjectDetails) GetOrganization() Organization {
	if o == nil || o.Organization == nil {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetOrganizationOk() (*Organization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ProjectDetails) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *ProjectDetails) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetPersons returns the Persons field value if set, zero value otherwise.
func (o *ProjectDetails) GetPersons() []Person {
	if o == nil || o.Persons == nil {
		var ret []Person
		return ret
	}
	return *o.Persons
}

// GetPersonsOk returns a tuple with the Persons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetails) GetPersonsOk() (*[]Person, bool) {
	if o == nil || o.Persons == nil {
		return nil, false
	}
	return o.Persons, true
}

// HasPersons returns a boolean if a field has been set.
func (o *ProjectDetails) HasPersons() bool {
	if o != nil && o.Persons != nil {
		return true
	}

	return false
}

// SetPersons gets a reference to the given []Person and assigns it to the Persons field.
func (o *ProjectDetails) SetPersons(v []Person) {
	o.Persons = &v
}

func (o ProjectDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Persons != nil {
		toSerialize["persons"] = o.Persons
	}
	return json.Marshal(toSerialize)
}

type NullableProjectDetails struct {
	value *ProjectDetails
	isSet bool
}

func (v NullableProjectDetails) Get() *ProjectDetails {
	return v.value
}

func (v *NullableProjectDetails) Set(val *ProjectDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDetails(val *ProjectDetails) *NullableProjectDetails {
	return &NullableProjectDetails{value: val, isSet: true}
}

func (v NullableProjectDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


