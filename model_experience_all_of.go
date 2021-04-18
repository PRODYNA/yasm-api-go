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

// ExperienceAllOf struct for ExperienceAllOf
type ExperienceAllOf struct {
	Skill *SkillLevel `json:"skill,omitempty"`
	ConfirmedBy *[]Person `json:"confirmedBy,omitempty"`
}

// NewExperienceAllOf instantiates a new ExperienceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExperienceAllOf() *ExperienceAllOf {
	this := ExperienceAllOf{}
	return &this
}

// NewExperienceAllOfWithDefaults instantiates a new ExperienceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExperienceAllOfWithDefaults() *ExperienceAllOf {
	this := ExperienceAllOf{}
	return &this
}

// GetSkill returns the Skill field value if set, zero value otherwise.
func (o *ExperienceAllOf) GetSkill() SkillLevel {
	if o == nil || o.Skill == nil {
		var ret SkillLevel
		return ret
	}
	return *o.Skill
}

// GetSkillOk returns a tuple with the Skill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperienceAllOf) GetSkillOk() (*SkillLevel, bool) {
	if o == nil || o.Skill == nil {
		return nil, false
	}
	return o.Skill, true
}

// HasSkill returns a boolean if a field has been set.
func (o *ExperienceAllOf) HasSkill() bool {
	if o != nil && o.Skill != nil {
		return true
	}

	return false
}

// SetSkill gets a reference to the given SkillLevel and assigns it to the Skill field.
func (o *ExperienceAllOf) SetSkill(v SkillLevel) {
	o.Skill = &v
}

// GetConfirmedBy returns the ConfirmedBy field value if set, zero value otherwise.
func (o *ExperienceAllOf) GetConfirmedBy() []Person {
	if o == nil || o.ConfirmedBy == nil {
		var ret []Person
		return ret
	}
	return *o.ConfirmedBy
}

// GetConfirmedByOk returns a tuple with the ConfirmedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExperienceAllOf) GetConfirmedByOk() (*[]Person, bool) {
	if o == nil || o.ConfirmedBy == nil {
		return nil, false
	}
	return o.ConfirmedBy, true
}

// HasConfirmedBy returns a boolean if a field has been set.
func (o *ExperienceAllOf) HasConfirmedBy() bool {
	if o != nil && o.ConfirmedBy != nil {
		return true
	}

	return false
}

// SetConfirmedBy gets a reference to the given []Person and assigns it to the ConfirmedBy field.
func (o *ExperienceAllOf) SetConfirmedBy(v []Person) {
	o.ConfirmedBy = &v
}

func (o ExperienceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skill != nil {
		toSerialize["skill"] = o.Skill
	}
	if o.ConfirmedBy != nil {
		toSerialize["confirmedBy"] = o.ConfirmedBy
	}
	return json.Marshal(toSerialize)
}

type NullableExperienceAllOf struct {
	value *ExperienceAllOf
	isSet bool
}

func (v NullableExperienceAllOf) Get() *ExperienceAllOf {
	return v.value
}

func (v *NullableExperienceAllOf) Set(val *ExperienceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExperienceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExperienceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExperienceAllOf(val *ExperienceAllOf) *NullableExperienceAllOf {
	return &NullableExperienceAllOf{value: val, isSet: true}
}

func (v NullableExperienceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExperienceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


