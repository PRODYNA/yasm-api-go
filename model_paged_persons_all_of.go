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

// PagedPersonsAllOf struct for PagedPersonsAllOf
type PagedPersonsAllOf struct {
	Persons []PersonScoreDetail `json:"persons,omitempty"`
}

// NewPagedPersonsAllOf instantiates a new PagedPersonsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedPersonsAllOf() *PagedPersonsAllOf {
	this := PagedPersonsAllOf{}
	return &this
}

// NewPagedPersonsAllOfWithDefaults instantiates a new PagedPersonsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedPersonsAllOfWithDefaults() *PagedPersonsAllOf {
	this := PagedPersonsAllOf{}
	return &this
}

// GetPersons returns the Persons field value if set, zero value otherwise.
func (o *PagedPersonsAllOf) GetPersons() []PersonScoreDetail {
	if o == nil || o.Persons == nil {
		var ret []PersonScoreDetail
		return ret
	}
	return o.Persons
}

// GetPersonsOk returns a tuple with the Persons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedPersonsAllOf) GetPersonsOk() ([]PersonScoreDetail, bool) {
	if o == nil || o.Persons == nil {
		return nil, false
	}
	return o.Persons, true
}

// HasPersons returns a boolean if a field has been set.
func (o *PagedPersonsAllOf) HasPersons() bool {
	if o != nil && o.Persons != nil {
		return true
	}

	return false
}

// SetPersons gets a reference to the given []PersonScoreDetail and assigns it to the Persons field.
func (o *PagedPersonsAllOf) SetPersons(v []PersonScoreDetail) {
	o.Persons = v
}

func (o PagedPersonsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Persons != nil {
		toSerialize["persons"] = o.Persons
	}
	return json.Marshal(toSerialize)
}

type NullablePagedPersonsAllOf struct {
	value *PagedPersonsAllOf
	isSet bool
}

func (v NullablePagedPersonsAllOf) Get() *PagedPersonsAllOf {
	return v.value
}

func (v *NullablePagedPersonsAllOf) Set(val *PagedPersonsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedPersonsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedPersonsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedPersonsAllOf(val *PagedPersonsAllOf) *NullablePagedPersonsAllOf {
	return &NullablePagedPersonsAllOf{value: val, isSet: true}
}

func (v NullablePagedPersonsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedPersonsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


