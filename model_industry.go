/*
 * YASM (Yet Another Skill Management) API
 *
 * This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.
 *
 * API version: 0.6.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Industry struct for Industry
type Industry struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Suggestion *bool `json:"suggestion,omitempty"`
}

// NewIndustry instantiates a new Industry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIndustry(id string, name string) *Industry {
	this := Industry{}
	this.Id = id
	this.Name = name
	var suggestion bool = false
	this.Suggestion = &suggestion
	return &this
}

// NewIndustryWithDefaults instantiates a new Industry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIndustryWithDefaults() *Industry {
	this := Industry{}
	var suggestion bool = false
	this.Suggestion = &suggestion
	return &this
}

// GetId returns the Id field value
func (o *Industry) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Industry) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Industry) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Industry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Industry) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Industry) SetName(v string) {
	o.Name = v
}

// GetSuggestion returns the Suggestion field value if set, zero value otherwise.
func (o *Industry) GetSuggestion() bool {
	if o == nil || o.Suggestion == nil {
		var ret bool
		return ret
	}
	return *o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Industry) GetSuggestionOk() (*bool, bool) {
	if o == nil || o.Suggestion == nil {
		return nil, false
	}
	return o.Suggestion, true
}

// HasSuggestion returns a boolean if a field has been set.
func (o *Industry) HasSuggestion() bool {
	if o != nil && o.Suggestion != nil {
		return true
	}

	return false
}

// SetSuggestion gets a reference to the given bool and assigns it to the Suggestion field.
func (o *Industry) SetSuggestion(v bool) {
	o.Suggestion = &v
}

func (o Industry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Suggestion != nil {
		toSerialize["suggestion"] = o.Suggestion
	}
	return json.Marshal(toSerialize)
}

type NullableIndustry struct {
	value *Industry
	isSet bool
}

func (v NullableIndustry) Get() *Industry {
	return v.value
}

func (v *NullableIndustry) Set(val *Industry) {
	v.value = val
	v.isSet = true
}

func (v NullableIndustry) IsSet() bool {
	return v.isSet
}

func (v *NullableIndustry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndustry(val *Industry) *NullableIndustry {
	return &NullableIndustry{value: val, isSet: true}
}

func (v NullableIndustry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndustry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


