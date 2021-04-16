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

// PersonDetails struct for PersonDetails
type PersonDetails struct {
	Person *Person `json:"person,omitempty"`
	Projects *[]ProjectParticipation `json:"projects,omitempty"`
	Interests *[]Skill `json:"interests,omitempty"`
	Certifications *[]CertificationDetails `json:"certifications,omitempty"`
	Languages *[]LanguageLevel `json:"languages,omitempty"`
	Location *string `json:"location,omitempty"`
	Office *Office `json:"office,omitempty"`
}

// NewPersonDetails instantiates a new PersonDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonDetails() *PersonDetails {
	this := PersonDetails{}
	return &this
}

// NewPersonDetailsWithDefaults instantiates a new PersonDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonDetailsWithDefaults() *PersonDetails {
	this := PersonDetails{}
	return &this
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *PersonDetails) GetPerson() Person {
	if o == nil || o.Person == nil {
		var ret Person
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetPersonOk() (*Person, bool) {
	if o == nil || o.Person == nil {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *PersonDetails) HasPerson() bool {
	if o != nil && o.Person != nil {
		return true
	}

	return false
}

// SetPerson gets a reference to the given Person and assigns it to the Person field.
func (o *PersonDetails) SetPerson(v Person) {
	o.Person = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PersonDetails) GetProjects() []ProjectParticipation {
	if o == nil || o.Projects == nil {
		var ret []ProjectParticipation
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetProjectsOk() (*[]ProjectParticipation, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PersonDetails) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []ProjectParticipation and assigns it to the Projects field.
func (o *PersonDetails) SetProjects(v []ProjectParticipation) {
	o.Projects = &v
}

// GetInterests returns the Interests field value if set, zero value otherwise.
func (o *PersonDetails) GetInterests() []Skill {
	if o == nil || o.Interests == nil {
		var ret []Skill
		return ret
	}
	return *o.Interests
}

// GetInterestsOk returns a tuple with the Interests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetInterestsOk() (*[]Skill, bool) {
	if o == nil || o.Interests == nil {
		return nil, false
	}
	return o.Interests, true
}

// HasInterests returns a boolean if a field has been set.
func (o *PersonDetails) HasInterests() bool {
	if o != nil && o.Interests != nil {
		return true
	}

	return false
}

// SetInterests gets a reference to the given []Skill and assigns it to the Interests field.
func (o *PersonDetails) SetInterests(v []Skill) {
	o.Interests = &v
}

// GetCertifications returns the Certifications field value if set, zero value otherwise.
func (o *PersonDetails) GetCertifications() []CertificationDetails {
	if o == nil || o.Certifications == nil {
		var ret []CertificationDetails
		return ret
	}
	return *o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetCertificationsOk() (*[]CertificationDetails, bool) {
	if o == nil || o.Certifications == nil {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *PersonDetails) HasCertifications() bool {
	if o != nil && o.Certifications != nil {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []CertificationDetails and assigns it to the Certifications field.
func (o *PersonDetails) SetCertifications(v []CertificationDetails) {
	o.Certifications = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *PersonDetails) GetLanguages() []LanguageLevel {
	if o == nil || o.Languages == nil {
		var ret []LanguageLevel
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetLanguagesOk() (*[]LanguageLevel, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *PersonDetails) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []LanguageLevel and assigns it to the Languages field.
func (o *PersonDetails) SetLanguages(v []LanguageLevel) {
	o.Languages = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PersonDetails) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PersonDetails) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *PersonDetails) SetLocation(v string) {
	o.Location = &v
}

// GetOffice returns the Office field value if set, zero value otherwise.
func (o *PersonDetails) GetOffice() Office {
	if o == nil || o.Office == nil {
		var ret Office
		return ret
	}
	return *o.Office
}

// GetOfficeOk returns a tuple with the Office field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonDetails) GetOfficeOk() (*Office, bool) {
	if o == nil || o.Office == nil {
		return nil, false
	}
	return o.Office, true
}

// HasOffice returns a boolean if a field has been set.
func (o *PersonDetails) HasOffice() bool {
	if o != nil && o.Office != nil {
		return true
	}

	return false
}

// SetOffice gets a reference to the given Office and assigns it to the Office field.
func (o *PersonDetails) SetOffice(v Office) {
	o.Office = &v
}

func (o PersonDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Person != nil {
		toSerialize["person"] = o.Person
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Interests != nil {
		toSerialize["interests"] = o.Interests
	}
	if o.Certifications != nil {
		toSerialize["certifications"] = o.Certifications
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Office != nil {
		toSerialize["office"] = o.Office
	}
	return json.Marshal(toSerialize)
}

type NullablePersonDetails struct {
	value *PersonDetails
	isSet bool
}

func (v NullablePersonDetails) Get() *PersonDetails {
	return v.value
}

func (v *NullablePersonDetails) Set(val *PersonDetails) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonDetails) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonDetails(val *PersonDetails) *NullablePersonDetails {
	return &NullablePersonDetails{value: val, isSet: true}
}

func (v NullablePersonDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


