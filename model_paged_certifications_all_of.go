/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.0.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedCertificationsAllOf struct for PagedCertificationsAllOf
type PagedCertificationsAllOf struct {
	Certifications *[]Certification `json:"certifications,omitempty"`
}

// NewPagedCertificationsAllOf instantiates a new PagedCertificationsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedCertificationsAllOf() *PagedCertificationsAllOf {
	this := PagedCertificationsAllOf{}
	return &this
}

// NewPagedCertificationsAllOfWithDefaults instantiates a new PagedCertificationsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedCertificationsAllOfWithDefaults() *PagedCertificationsAllOf {
	this := PagedCertificationsAllOf{}
	return &this
}

// GetCertifications returns the Certifications field value if set, zero value otherwise.
func (o *PagedCertificationsAllOf) GetCertifications() []Certification {
	if o == nil || o.Certifications == nil {
		var ret []Certification
		return ret
	}
	return *o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedCertificationsAllOf) GetCertificationsOk() (*[]Certification, bool) {
	if o == nil || o.Certifications == nil {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *PagedCertificationsAllOf) HasCertifications() bool {
	if o != nil && o.Certifications != nil {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []Certification and assigns it to the Certifications field.
func (o *PagedCertificationsAllOf) SetCertifications(v []Certification) {
	o.Certifications = &v
}

func (o PagedCertificationsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certifications != nil {
		toSerialize["certifications"] = o.Certifications
	}
	return json.Marshal(toSerialize)
}

type NullablePagedCertificationsAllOf struct {
	value *PagedCertificationsAllOf
	isSet bool
}

func (v NullablePagedCertificationsAllOf) Get() *PagedCertificationsAllOf {
	return v.value
}

func (v *NullablePagedCertificationsAllOf) Set(val *PagedCertificationsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedCertificationsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedCertificationsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedCertificationsAllOf(val *PagedCertificationsAllOf) *NullablePagedCertificationsAllOf {
	return &NullablePagedCertificationsAllOf{value: val, isSet: true}
}

func (v NullablePagedCertificationsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedCertificationsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


