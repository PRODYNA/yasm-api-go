/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.7.4
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OfficeAllOf struct for OfficeAllOf
type OfficeAllOf struct {
	Geolocation *Geolocation `json:"geolocation,omitempty"`
}

// NewOfficeAllOf instantiates a new OfficeAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOfficeAllOf() *OfficeAllOf {
	this := OfficeAllOf{}
	return &this
}

// NewOfficeAllOfWithDefaults instantiates a new OfficeAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfficeAllOfWithDefaults() *OfficeAllOf {
	this := OfficeAllOf{}
	return &this
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *OfficeAllOf) GetGeolocation() Geolocation {
	if o == nil || o.Geolocation == nil {
		var ret Geolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OfficeAllOf) GetGeolocationOk() (*Geolocation, bool) {
	if o == nil || o.Geolocation == nil {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *OfficeAllOf) HasGeolocation() bool {
	if o != nil && o.Geolocation != nil {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given Geolocation and assigns it to the Geolocation field.
func (o *OfficeAllOf) SetGeolocation(v Geolocation) {
	o.Geolocation = &v
}

func (o OfficeAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Geolocation != nil {
		toSerialize["geolocation"] = o.Geolocation
	}
	return json.Marshal(toSerialize)
}

type NullableOfficeAllOf struct {
	value *OfficeAllOf
	isSet bool
}

func (v NullableOfficeAllOf) Get() *OfficeAllOf {
	return v.value
}

func (v *NullableOfficeAllOf) Set(val *OfficeAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOfficeAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOfficeAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOfficeAllOf(val *OfficeAllOf) *NullableOfficeAllOf {
	return &NullableOfficeAllOf{value: val, isSet: true}
}

func (v NullableOfficeAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOfficeAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


