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

// ProjectParticipationUpdateTimeframe struct for ProjectParticipationUpdateTimeframe
type ProjectParticipationUpdateTimeframe struct {
	Startdate *string `json:"startdate,omitempty"`
	Enddate *string `json:"enddate,omitempty"`
}

// NewProjectParticipationUpdateTimeframe instantiates a new ProjectParticipationUpdateTimeframe object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectParticipationUpdateTimeframe() *ProjectParticipationUpdateTimeframe {
	this := ProjectParticipationUpdateTimeframe{}
	return &this
}

// NewProjectParticipationUpdateTimeframeWithDefaults instantiates a new ProjectParticipationUpdateTimeframe object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectParticipationUpdateTimeframeWithDefaults() *ProjectParticipationUpdateTimeframe {
	this := ProjectParticipationUpdateTimeframe{}
	return &this
}

// GetStartdate returns the Startdate field value if set, zero value otherwise.
func (o *ProjectParticipationUpdateTimeframe) GetStartdate() string {
	if o == nil || o.Startdate == nil {
		var ret string
		return ret
	}
	return *o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationUpdateTimeframe) GetStartdateOk() (*string, bool) {
	if o == nil || o.Startdate == nil {
		return nil, false
	}
	return o.Startdate, true
}

// HasStartdate returns a boolean if a field has been set.
func (o *ProjectParticipationUpdateTimeframe) HasStartdate() bool {
	if o != nil && o.Startdate != nil {
		return true
	}

	return false
}

// SetStartdate gets a reference to the given string and assigns it to the Startdate field.
func (o *ProjectParticipationUpdateTimeframe) SetStartdate(v string) {
	o.Startdate = &v
}

// GetEnddate returns the Enddate field value if set, zero value otherwise.
func (o *ProjectParticipationUpdateTimeframe) GetEnddate() string {
	if o == nil || o.Enddate == nil {
		var ret string
		return ret
	}
	return *o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectParticipationUpdateTimeframe) GetEnddateOk() (*string, bool) {
	if o == nil || o.Enddate == nil {
		return nil, false
	}
	return o.Enddate, true
}

// HasEnddate returns a boolean if a field has been set.
func (o *ProjectParticipationUpdateTimeframe) HasEnddate() bool {
	if o != nil && o.Enddate != nil {
		return true
	}

	return false
}

// SetEnddate gets a reference to the given string and assigns it to the Enddate field.
func (o *ProjectParticipationUpdateTimeframe) SetEnddate(v string) {
	o.Enddate = &v
}

func (o ProjectParticipationUpdateTimeframe) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Startdate != nil {
		toSerialize["startdate"] = o.Startdate
	}
	if o.Enddate != nil {
		toSerialize["enddate"] = o.Enddate
	}
	return json.Marshal(toSerialize)
}

type NullableProjectParticipationUpdateTimeframe struct {
	value *ProjectParticipationUpdateTimeframe
	isSet bool
}

func (v NullableProjectParticipationUpdateTimeframe) Get() *ProjectParticipationUpdateTimeframe {
	return v.value
}

func (v *NullableProjectParticipationUpdateTimeframe) Set(val *ProjectParticipationUpdateTimeframe) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectParticipationUpdateTimeframe) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectParticipationUpdateTimeframe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectParticipationUpdateTimeframe(val *ProjectParticipationUpdateTimeframe) *NullableProjectParticipationUpdateTimeframe {
	return &NullableProjectParticipationUpdateTimeframe{value: val, isSet: true}
}

func (v NullableProjectParticipationUpdateTimeframe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectParticipationUpdateTimeframe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


