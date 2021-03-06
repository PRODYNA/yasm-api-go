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

// ProjectScoreDetailAllOf struct for ProjectScoreDetailAllOf
type ProjectScoreDetailAllOf struct {
	Score *float32 `json:"score,omitempty"`
}

// NewProjectScoreDetailAllOf instantiates a new ProjectScoreDetailAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectScoreDetailAllOf() *ProjectScoreDetailAllOf {
	this := ProjectScoreDetailAllOf{}
	return &this
}

// NewProjectScoreDetailAllOfWithDefaults instantiates a new ProjectScoreDetailAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectScoreDetailAllOfWithDefaults() *ProjectScoreDetailAllOf {
	this := ProjectScoreDetailAllOf{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ProjectScoreDetailAllOf) GetScore() float32 {
	if o == nil || o.Score == nil {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectScoreDetailAllOf) GetScoreOk() (*float32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ProjectScoreDetailAllOf) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *ProjectScoreDetailAllOf) SetScore(v float32) {
	o.Score = &v
}

func (o ProjectScoreDetailAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableProjectScoreDetailAllOf struct {
	value *ProjectScoreDetailAllOf
	isSet bool
}

func (v NullableProjectScoreDetailAllOf) Get() *ProjectScoreDetailAllOf {
	return v.value
}

func (v *NullableProjectScoreDetailAllOf) Set(val *ProjectScoreDetailAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectScoreDetailAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectScoreDetailAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectScoreDetailAllOf(val *ProjectScoreDetailAllOf) *NullableProjectScoreDetailAllOf {
	return &NullableProjectScoreDetailAllOf{value: val, isSet: true}
}

func (v NullableProjectScoreDetailAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectScoreDetailAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


