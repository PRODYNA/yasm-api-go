/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.8.11
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedLanguages struct for PagedLanguages
type PagedLanguages struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Languages *[]Language `json:"languages,omitempty"`
}

// NewPagedLanguages instantiates a new PagedLanguages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedLanguages() *PagedLanguages {
	this := PagedLanguages{}
	return &this
}

// NewPagedLanguagesWithDefaults instantiates a new PagedLanguages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedLanguagesWithDefaults() *PagedLanguages {
	this := PagedLanguages{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *PagedLanguages) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedLanguages) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *PagedLanguages) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *PagedLanguages) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PagedLanguages) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedLanguages) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PagedLanguages) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PagedLanguages) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PagedLanguages) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedLanguages) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PagedLanguages) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PagedLanguages) SetCount(v int32) {
	o.Count = &v
}

// GetLanguages returns the Languages field value if set, zero value otherwise.
func (o *PagedLanguages) GetLanguages() []Language {
	if o == nil || o.Languages == nil {
		var ret []Language
		return ret
	}
	return *o.Languages
}

// GetLanguagesOk returns a tuple with the Languages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedLanguages) GetLanguagesOk() (*[]Language, bool) {
	if o == nil || o.Languages == nil {
		return nil, false
	}
	return o.Languages, true
}

// HasLanguages returns a boolean if a field has been set.
func (o *PagedLanguages) HasLanguages() bool {
	if o != nil && o.Languages != nil {
		return true
	}

	return false
}

// SetLanguages gets a reference to the given []Language and assigns it to the Languages field.
func (o *PagedLanguages) SetLanguages(v []Language) {
	o.Languages = &v
}

func (o PagedLanguages) MarshalJSON() ([]byte, error) {
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
	if o.Languages != nil {
		toSerialize["languages"] = o.Languages
	}
	return json.Marshal(toSerialize)
}

type NullablePagedLanguages struct {
	value *PagedLanguages
	isSet bool
}

func (v NullablePagedLanguages) Get() *PagedLanguages {
	return v.value
}

func (v *NullablePagedLanguages) Set(val *PagedLanguages) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedLanguages) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedLanguages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedLanguages(val *PagedLanguages) *NullablePagedLanguages {
	return &NullablePagedLanguages{value: val, isSet: true}
}

func (v NullablePagedLanguages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedLanguages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


