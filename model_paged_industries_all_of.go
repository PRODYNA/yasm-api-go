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

// PagedIndustriesAllOf struct for PagedIndustriesAllOf
type PagedIndustriesAllOf struct {
	Industries *[]Industry `json:"industries,omitempty"`
}

// NewPagedIndustriesAllOf instantiates a new PagedIndustriesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedIndustriesAllOf() *PagedIndustriesAllOf {
	this := PagedIndustriesAllOf{}
	return &this
}

// NewPagedIndustriesAllOfWithDefaults instantiates a new PagedIndustriesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedIndustriesAllOfWithDefaults() *PagedIndustriesAllOf {
	this := PagedIndustriesAllOf{}
	return &this
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *PagedIndustriesAllOf) GetIndustries() []Industry {
	if o == nil || o.Industries == nil {
		var ret []Industry
		return ret
	}
	return *o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedIndustriesAllOf) GetIndustriesOk() (*[]Industry, bool) {
	if o == nil || o.Industries == nil {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *PagedIndustriesAllOf) HasIndustries() bool {
	if o != nil && o.Industries != nil {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []Industry and assigns it to the Industries field.
func (o *PagedIndustriesAllOf) SetIndustries(v []Industry) {
	o.Industries = &v
}

func (o PagedIndustriesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Industries != nil {
		toSerialize["industries"] = o.Industries
	}
	return json.Marshal(toSerialize)
}

type NullablePagedIndustriesAllOf struct {
	value *PagedIndustriesAllOf
	isSet bool
}

func (v NullablePagedIndustriesAllOf) Get() *PagedIndustriesAllOf {
	return v.value
}

func (v *NullablePagedIndustriesAllOf) Set(val *PagedIndustriesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedIndustriesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedIndustriesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedIndustriesAllOf(val *PagedIndustriesAllOf) *NullablePagedIndustriesAllOf {
	return &NullablePagedIndustriesAllOf{value: val, isSet: true}
}

func (v NullablePagedIndustriesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedIndustriesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

