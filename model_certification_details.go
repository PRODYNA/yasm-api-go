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

// CertificationDetails struct for CertificationDetails
type CertificationDetails struct {
	Certification *Certification `json:"certification,omitempty"`
	Organization *Organization `json:"organization,omitempty"`
	Skills []SkillLevel `json:"skills,omitempty"`
}

// NewCertificationDetails instantiates a new CertificationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificationDetails() *CertificationDetails {
	this := CertificationDetails{}
	return &this
}

// NewCertificationDetailsWithDefaults instantiates a new CertificationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationDetailsWithDefaults() *CertificationDetails {
	this := CertificationDetails{}
	return &this
}

// GetCertification returns the Certification field value if set, zero value otherwise.
func (o *CertificationDetails) GetCertification() Certification {
	if o == nil || o.Certification == nil {
		var ret Certification
		return ret
	}
	return *o.Certification
}

// GetCertificationOk returns a tuple with the Certification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationDetails) GetCertificationOk() (*Certification, bool) {
	if o == nil || o.Certification == nil {
		return nil, false
	}
	return o.Certification, true
}

// HasCertification returns a boolean if a field has been set.
func (o *CertificationDetails) HasCertification() bool {
	if o != nil && o.Certification != nil {
		return true
	}

	return false
}

// SetCertification gets a reference to the given Certification and assigns it to the Certification field.
func (o *CertificationDetails) SetCertification(v Certification) {
	o.Certification = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *CertificationDetails) GetOrganization() Organization {
	if o == nil || o.Organization == nil {
		var ret Organization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationDetails) GetOrganizationOk() (*Organization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *CertificationDetails) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given Organization and assigns it to the Organization field.
func (o *CertificationDetails) SetOrganization(v Organization) {
	o.Organization = &v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *CertificationDetails) GetSkills() []SkillLevel {
	if o == nil || o.Skills == nil {
		var ret []SkillLevel
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificationDetails) GetSkillsOk() ([]SkillLevel, bool) {
	if o == nil || o.Skills == nil {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *CertificationDetails) HasSkills() bool {
	if o != nil && o.Skills != nil {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []SkillLevel and assigns it to the Skills field.
func (o *CertificationDetails) SetSkills(v []SkillLevel) {
	o.Skills = v
}

func (o CertificationDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certification != nil {
		toSerialize["certification"] = o.Certification
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Skills != nil {
		toSerialize["skills"] = o.Skills
	}
	return json.Marshal(toSerialize)
}

type NullableCertificationDetails struct {
	value *CertificationDetails
	isSet bool
}

func (v NullableCertificationDetails) Get() *CertificationDetails {
	return v.value
}

func (v *NullableCertificationDetails) Set(val *CertificationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificationDetails(val *CertificationDetails) *NullableCertificationDetails {
	return &NullableCertificationDetails{value: val, isSet: true}
}

func (v NullableCertificationDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


