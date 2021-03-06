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

// PagedOrganizations struct for PagedOrganizations
type PagedOrganizations struct {
	Skip *int32 `json:"skip,omitempty"`
	Limit *int32 `json:"limit,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Organizations []OrganizationDetails `json:"organizations,omitempty"`
}

// NewPagedOrganizations instantiates a new PagedOrganizations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagedOrganizations() *PagedOrganizations {
	this := PagedOrganizations{}
	return &this
}

// NewPagedOrganizationsWithDefaults instantiates a new PagedOrganizations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagedOrganizationsWithDefaults() *PagedOrganizations {
	this := PagedOrganizations{}
	return &this
}

// GetSkip returns the Skip field value if set, zero value otherwise.
func (o *PagedOrganizations) GetSkip() int32 {
	if o == nil || o.Skip == nil {
		var ret int32
		return ret
	}
	return *o.Skip
}

// GetSkipOk returns a tuple with the Skip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedOrganizations) GetSkipOk() (*int32, bool) {
	if o == nil || o.Skip == nil {
		return nil, false
	}
	return o.Skip, true
}

// HasSkip returns a boolean if a field has been set.
func (o *PagedOrganizations) HasSkip() bool {
	if o != nil && o.Skip != nil {
		return true
	}

	return false
}

// SetSkip gets a reference to the given int32 and assigns it to the Skip field.
func (o *PagedOrganizations) SetSkip(v int32) {
	o.Skip = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *PagedOrganizations) GetLimit() int32 {
	if o == nil || o.Limit == nil {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedOrganizations) GetLimitOk() (*int32, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *PagedOrganizations) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *PagedOrganizations) SetLimit(v int32) {
	o.Limit = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *PagedOrganizations) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedOrganizations) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *PagedOrganizations) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *PagedOrganizations) SetCount(v int32) {
	o.Count = &v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *PagedOrganizations) GetOrganizations() []OrganizationDetails {
	if o == nil || o.Organizations == nil {
		var ret []OrganizationDetails
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagedOrganizations) GetOrganizationsOk() ([]OrganizationDetails, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *PagedOrganizations) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []OrganizationDetails and assigns it to the Organizations field.
func (o *PagedOrganizations) SetOrganizations(v []OrganizationDetails) {
	o.Organizations = v
}

func (o PagedOrganizations) MarshalJSON() ([]byte, error) {
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
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	return json.Marshal(toSerialize)
}

type NullablePagedOrganizations struct {
	value *PagedOrganizations
	isSet bool
}

func (v NullablePagedOrganizations) Get() *PagedOrganizations {
	return v.value
}

func (v *NullablePagedOrganizations) Set(val *PagedOrganizations) {
	v.value = val
	v.isSet = true
}

func (v NullablePagedOrganizations) IsSet() bool {
	return v.isSet
}

func (v *NullablePagedOrganizations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagedOrganizations(val *PagedOrganizations) *NullablePagedOrganizations {
	return &NullablePagedOrganizations{value: val, isSet: true}
}

func (v NullablePagedOrganizations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagedOrganizations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


