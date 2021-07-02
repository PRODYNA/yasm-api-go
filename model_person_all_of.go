/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.7.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PersonAllOf struct for PersonAllOf
type PersonAllOf struct {
	Location *string `json:"location,omitempty"`
}

// NewPersonAllOf instantiates a new PersonAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonAllOf() *PersonAllOf {
	this := PersonAllOf{}
	return &this
}

// NewPersonAllOfWithDefaults instantiates a new PersonAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonAllOfWithDefaults() *PersonAllOf {
	this := PersonAllOf{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *PersonAllOf) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonAllOf) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *PersonAllOf) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *PersonAllOf) SetLocation(v string) {
	o.Location = &v
}

func (o PersonAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	return json.Marshal(toSerialize)
}

type NullablePersonAllOf struct {
	value *PersonAllOf
	isSet bool
}

func (v NullablePersonAllOf) Get() *PersonAllOf {
	return v.value
}

func (v *NullablePersonAllOf) Set(val *PersonAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonAllOf(val *PersonAllOf) *NullablePersonAllOf {
	return &NullablePersonAllOf{value: val, isSet: true}
}

func (v NullablePersonAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


