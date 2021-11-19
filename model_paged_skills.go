/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.9.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedSkills struct for PagedSkills
type PagedSkills struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Skills *[]Skill `json:"skills,omitempty"`
}

// NewPagedSkills instantiates a new PagedSkills object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedSkills() *PagedSkills {
	this := PagedSkills{}
	return &this
}

// NewPagedSkillsWithDefaults instantiates a new PagedSkills object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedSkillsWithDefaults() *PagedSkills {
	this := PagedSkills{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *PagedSkills) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedSkills) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *PagedSkills) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *PagedSkills) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PagedSkills) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedSkills) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PagedSkills) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PagedSkills) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PagedSkills) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedSkills) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PagedSkills) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PagedSkills) SetCount(v int32) {
	o.Count = &v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *PagedSkills) GetSkills() []Skill {
	if o == nil || o.Skills == nil {
		var ret []Skill
		return ret
	}
	return *o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedSkills) GetSkillsOk() (*[]Skill, bool) {
	if o == nil || o.Skills == nil {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *PagedSkills) HasSkills() bool {
	if o != nil && o.Skills != nil {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []Skill and assigns it to the Skills field.
func (o *PagedSkills) SetSkills(v []Skill) {
	o.Skills = &v
}

func (o PagedSkills) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Skip != nil {
		toSerialize["skip"] = o.Skip
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Skills != nil {
		toSerialize["skills"] = o.Skills
	}
	return json.Marshal(toSerialize)
}

type NullablePagedSkills struct {
	value *PagedSkills
	isSet bool
}

func (v NullablePagedSkills) Get() *PagedSkills {
	return v.value
}

func (v *NullablePagedSkills) Set(val *PagedSkills) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedSkills) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedSkills) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedSkills(val *PagedSkills) *NullablePagedSkills {
	return &NullablePagedSkills{value: val, isSet: true}
}

func (v NullablePagedSkills) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedSkills) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


