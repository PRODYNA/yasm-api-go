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

// CountryDetails struct for CountryDetails
type CountryDetails struct {
	Country *Country `json:"country,omitempty"`
	Languages *[]Language `json:"languages,omitempty"`
}

// NewCountryDetails instantiates a new CountryDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryDetails() *CountryDetails {
	this := CountryDetails{}
	return &this
}

// NewCountryDetailsWithDefaults instantiates a new CountryDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryDetailsWithDefaults() *CountryDetails {
	this := CountryDetails{}
	return &this
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *CountryDetails) GetCountry() Country {
	if o == nil || o.Country == nil {
		var ret Country
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryDetails) GetCountryOk() (*Country, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *CountryDetails) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given Country and assigns it to the Country field.
func (o *CountryDetails) SetCountry(v Country) {
	o.Country = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *CountryDetails) GetLanguages() []Language {
	if o == nil || o.Languages == nil {
		var ret []Language
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryDetails) GetLanguagesOk() (*[]Language, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *CountryDetails) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *CountryDetails) SetLanguages(v []Language) {
	o.Languages = &v
}

func (o CountryDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	return json.Marshal(toSerialize)
}

type NullableCountryDetails struct {
	value *CountryDetails
	isSet bool
}

func (v NullableCountryDetails) Get() *CountryDetails {
	return v.value
}

func (v *NullableCountryDetails) Set(val *CountryDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryDetails(val *CountryDetails) *NullableCountryDetails {
	return &NullableCountryDetails{value: val, isSet: true}
}

func (v NullableCountryDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

