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

// Timeframed struct for Timeframed
type Timeframed struct {
	Startdate *string `json:"startdate,omitempty"`
	Enddate *string `json:"enddate,omitempty"`
}

// NewTimeframed instantiates a new Timeframed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeframed() *Timeframed {
	this := Timeframed{}
	return &this
}

// NewTimeframedWithDefaults instantiates a new Timeframed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeframedWithDefaults() *Timeframed {
	this := Timeframed{}
	return &this
}

// GetStartdate returns the Startdate field value if set, zero value otherwise.
func (o *Timeframed) GetStartdate() string {
	if o == nil || o.Startdate == nil {
		var ret string
		return ret
	}
	return *o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeframed) GetStartdateOk() (*string, bool) {
	if o == nil || o.Startdate == nil {
		return nil, false
	}
	return o.Startdate, true
}

// HasStartdate returns a boolean if a field has been set.
func (o *Timeframed) HasStartdate() bool {
	if o != nil && o.Startdate != nil {
		return true
	}

	return false
}

// SetStartdate gets a reference to the given string and assigns it to the Startdate field.
func (o *Timeframed) SetStartdate(v string) {
	o.Startdate = &v
}

// GetEnddate returns the Enddate field value if set, zero value otherwise.
func (o *Timeframed) GetEnddate() string {
	if o == nil || o.Enddate == nil {
		var ret string
		return ret
	}
	return *o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Timeframed) GetEnddateOk() (*string, bool) {
	if o == nil || o.Enddate == nil {
		return nil, false
	}
	return o.Enddate, true
}

// HasEnddate returns a boolean if a field has been set.
func (o *Timeframed) HasEnddate() bool {
	if o != nil && o.Enddate != nil {
		return true
	}

	return false
}

// SetEnddate gets a reference to the given string and assigns it to the Enddate field.
func (o *Timeframed) SetEnddate(v string) {
	o.Enddate = &v
}

func (o Timeframed) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Startdate != nil {
		toSerialize["startdate"] = o.Startdate
	}
	if o.Enddate != nil {
		toSerialize["enddate"] = o.Enddate
	}
	return json.Marshal(toSerialize)
}

type NullableTimeframed struct {
	value *Timeframed
	isSet bool
}

func (v NullableTimeframed) Get() *Timeframed {
	return v.value
}

func (v *NullableTimeframed) Set(val *Timeframed) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeframed) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeframed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeframed(val *Timeframed) *NullableTimeframed {
	return &NullableTimeframed{value: val, isSet: true}
}

func (v NullableTimeframed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeframed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


