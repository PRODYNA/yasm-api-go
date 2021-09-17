/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.8.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// PagedOrganizationsAllOf struct for PagedOrganizationsAllOf
type PagedOrganizationsAllOf struct {
	Organizations *[]Organization `json:"organizations,omitempty"`
}

// NewPagedOrganizationsAllOf instantiates a new PagedOrganizationsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedOrganizationsAllOf() *PagedOrganizationsAllOf {
	this := PagedOrganizationsAllOf{}
	return &this
}

// NewPagedOrganizationsAllOfWithDefaults instantiates a new PagedOrganizationsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedOrganizationsAllOfWithDefaults() *PagedOrganizationsAllOf {
	this := PagedOrganizationsAllOf{}
	return &this
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *PagedOrganizationsAllOf) GetOrganizations() []Organization {
	if o == nil || o.Organizations == nil {
		var ret []Organization
		return ret
	}
	return *o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedOrganizationsAllOf) GetOrganizationsOk() (*[]Organization, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *PagedOrganizationsAllOf) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []Organization and assigns it to the Organizations field.
func (o *PagedOrganizationsAllOf) SetOrganizations(v []Organization) {
	o.Organizations = &v
}

func (o PagedOrganizationsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	return json.Marshal(toSerialize)
}

type NullablePagedOrganizationsAllOf struct {
	value *PagedOrganizationsAllOf
	isSet bool
}

func (v NullablePagedOrganizationsAllOf) Get() *PagedOrganizationsAllOf {
	return v.value
}

func (v *NullablePagedOrganizationsAllOf) Set(val *PagedOrganizationsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedOrganizationsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedOrganizationsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedOrganizationsAllOf(val *PagedOrganizationsAllOf) *NullablePagedOrganizationsAllOf {
	return &NullablePagedOrganizationsAllOf{value: val, isSet: true}
}

func (v NullablePagedOrganizationsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedOrganizationsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


