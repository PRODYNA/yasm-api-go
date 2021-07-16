/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.7.7
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonScoreDetail struct for PersonScoreDetail
type PersonScoreDetail struct {
	Person *Person `json:"person,omitempty"`
	Projects *[]ProjectParticipation `json:"projects,omitempty"`
	Experiences *[]Experience `json:"experiences,omitempty"`
	Interests *[]Skill `json:"interests,omitempty"`
	Certifications *[]CertificationDetails `json:"certifications,omitempty"`
	Languages *[]LanguageLevel `json:"languages,omitempty"`
	Location *string `json:"location,omitempty"`
	Office *Office `json:"office,omitempty"`
	Score *float32 `json:"score,omitempty"`
}

// NewPersonScoreDetail instantiates a new PersonScoreDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonScoreDetail() *PersonScoreDetail {
	this := PersonScoreDetail{}
	return &this
}

// NewPersonScoreDetailWithDefaults instantiates a new PersonScoreDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonScoreDetailWithDefaults() *PersonScoreDetail {
	this := PersonScoreDetail{}
	return &this
}

// GetPerson returns the Person field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetPerson() Person {
	if o == nil || o.Person == nil {
		var ret Person
		return ret
	}
	return *o.Person
}

// GetPersonOk returns a tuple with the Person field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetPersonOk() (*Person, bool) {
	if o == nil || o.Person == nil {
		return nil, false
	}
	return o.Person, true
}

// HasPerson returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasPerson() bool {
	if o != nil && o.Person != nil {
		return true
	}

	return false
}

// SetPerson gets a reference to the given Person and assigns it to the Person field.
func (o *PersonScoreDetail) SetPerson(v Person) {
	o.Person = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetProjects() []ProjectParticipation {
	if o == nil || o.Projects == nil {
		var ret []ProjectParticipation
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetProjectsOk() (*[]ProjectParticipation, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []ProjectParticipation and assigns it to the Projects field.
func (o *PersonScoreDetail) SetProjects(v []ProjectParticipation) {
	o.Projects = &v
}

// GetExperiences returns the Experiences field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetExperiences() []Experience {
	if o == nil || o.Experiences == nil {
		var ret []Experience
		return ret
	}
	return *o.Experiences
}

// GetExperiencesOk returns a tuple with the Experiences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetExperiencesOk() (*[]Experience, bool) {
	if o == nil || o.Experiences == nil {
		return nil, false
	}
	return o.Experiences, true
}

// HasExperiences returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasExperiences() bool {
	if o != nil && o.Experiences != nil {
		return true
	}

	return false
}

// SetExperiences gets a reference to the given []Experience and assigns it to the Experiences field.
func (o *PersonScoreDetail) SetExperiences(v []Experience) {
	o.Experiences = &v
}

// GetInterests returns the Interests field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetInterests() []Skill {
	if o == nil || o.Interests == nil {
		var ret []Skill
		return ret
	}
	return *o.Interests
}

// GetInterestsOk returns a tuple with the Interests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetInterestsOk() (*[]Skill, bool) {
	if o == nil || o.Interests == nil {
		return nil, false
	}
	return o.Interests, true
}

// HasInterests returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasInterests() bool {
	if o != nil && o.Interests != nil {
		return true
	}

	return false
}

// SetInterests gets a reference to the given []Skill and assigns it to the Interests field.
func (o *PersonScoreDetail) SetInterests(v []Skill) {
	o.Interests = &v
}

// GetCertifications returns the Certifications field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetCertifications() []CertificationDetails {
	if o == nil || o.Certifications == nil {
		var ret []CertificationDetails
		return ret
	}
	return *o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetCertificationsOk() (*[]CertificationDetails, bool) {
	if o == nil || o.Certifications == nil {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasCertifications() bool {
	if o != nil && o.Certifications != nil {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []CertificationDetails and assigns it to the Certifications field.
func (o *PersonScoreDetail) SetCertifications(v []CertificationDetails) {
	o.Certifications = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetLanguages() []LanguageLevel {
	if o == nil || o.Languages == nil {
		var ret []LanguageLevel
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetLanguagesOk() (*[]LanguageLevel, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []LanguageLevel and assigns it to the Languages field.
func (o *PersonScoreDetail) SetLanguages(v []LanguageLevel) {
	o.Languages = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *PersonScoreDetail) SetLocation(v string) {
	o.Location = &v
}

// GetOffice returns the Office field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetOffice() Office {
	if o == nil || o.Office == nil {
		var ret Office
		return ret
	}
	return *o.Office
}

// GetOfficeOk returns a tuple with the Office field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetOfficeOk() (*Office, bool) {
	if o == nil || o.Office == nil {
		return nil, false
	}
	return o.Office, true
}

// HasOffice returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasOffice() bool {
	if o != nil && o.Office != nil {
		return true
	}

	return false
}

// SetOffice gets a reference to the given Office and assigns it to the Office field.
func (o *PersonScoreDetail) SetOffice(v Office) {
	o.Office = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *PersonScoreDetail) GetScore() float32 {
	if o == nil || o.Score == nil {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonScoreDetail) GetScoreOk() (*float32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *PersonScoreDetail) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *PersonScoreDetail) SetScore(v float32) {
	o.Score = &v
}

func (o PersonScoreDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Person != nil {
		toSerialize["person"] = o.Person
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Experiences != nil {
		toSerialize["experiences"] = o.Experiences
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
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullablePersonScoreDetail struct {
	value *PersonScoreDetail
	isSet bool
}

func (v NullablePersonScoreDetail) Get() *PersonScoreDetail {
	return v.value
}

func (v *NullablePersonScoreDetail) Set(val *PersonScoreDetail) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonScoreDetail) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonScoreDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonScoreDetail(val *PersonScoreDetail) *NullablePersonScoreDetail {
	return &NullablePersonScoreDetail{value: val, isSet: true}
}

func (v NullablePersonScoreDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonScoreDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


