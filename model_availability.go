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

// Availability struct for Availability
type Availability struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Startdate string `json:"startdate"`
	Enddate string `json:"enddate"`
	WorkHours float32 `json:"workHours"`
	PlannedHours float32 `json:"plannedHours"`
	Descriptions []string `json:"descriptions,omitempty"`
}

// NewAvailability instantiates a new Availability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAvailability(id string, name string, startdate string, enddate string, workHours float32, plannedHours float32) *Availability {
	this := Availability{}
	this.Id = id
	this.Name = name
	this.Startdate = startdate
	this.Enddate = enddate
	this.WorkHours = workHours
	this.PlannedHours = plannedHours
	return &this
}

// NewAvailabilityWithDefaults instantiates a new Availability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAvailabilityWithDefaults() *Availability {
	this := Availability{}
	return &this
}

// GetId returns the Id field value
func (o *Availability) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Availability) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Availability) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Availability) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Availability) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Availability) SetName(v string) {
	o.Name = v
}

// GetStartdate returns the Startdate field value
func (o *Availability) GetStartdate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Startdate
}

// GetStartdateOk returns a tuple with the Startdate field value
// and a boolean to check if the value has been set.
func (o *Availability) GetStartdateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Startdate, true
}

// SetStartdate sets field value
func (o *Availability) SetStartdate(v string) {
	o.Startdate = v
}

// GetEnddate returns the Enddate field value
func (o *Availability) GetEnddate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Enddate
}

// GetEnddateOk returns a tuple with the Enddate field value
// and a boolean to check if the value has been set.
func (o *Availability) GetEnddateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enddate, true
}

// SetEnddate sets field value
func (o *Availability) SetEnddate(v string) {
	o.Enddate = v
}

// GetWorkHours returns the WorkHours field value
func (o *Availability) GetWorkHours() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.WorkHours
}

// GetWorkHoursOk returns a tuple with the WorkHours field value
// and a boolean to check if the value has been set.
func (o *Availability) GetWorkHoursOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkHours, true
}

// SetWorkHours sets field value
func (o *Availability) SetWorkHours(v float32) {
	o.WorkHours = v
}

// GetPlannedHours returns the PlannedHours field value
func (o *Availability) GetPlannedHours() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.PlannedHours
}

// GetPlannedHoursOk returns a tuple with the PlannedHours field value
// and a boolean to check if the value has been set.
func (o *Availability) GetPlannedHoursOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlannedHours, true
}

// SetPlannedHours sets field value
func (o *Availability) SetPlannedHours(v float32) {
	o.PlannedHours = v
}

// GetDescriptions returns the Descriptions field value if set, zero value otherwise.
func (o *Availability) GetDescriptions() []string {
	if o == nil || o.Descriptions == nil {
		var ret []string
		return ret
	}
	return o.Descriptions
}

// GetDescriptionsOk returns a tuple with the Descriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Availability) GetDescriptionsOk() ([]string, bool) {
	if o == nil || o.Descriptions == nil {
		return nil, false
	}
	return o.Descriptions, true
}

// HasDescriptions returns a boolean if a field has been set.
func (o *Availability) HasDescriptions() bool {
	if o != nil && o.Descriptions != nil {
		return true
	}

	return false
}

// SetDescriptions gets a reference to the given []string and assigns it to the Descriptions field.
func (o *Availability) SetDescriptions(v []string) {
	o.Descriptions = v
}

func (o Availability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["startdate"] = o.Startdate
	}
	if true {
		toSerialize["enddate"] = o.Enddate
	}
	if true {
		toSerialize["workHours"] = o.WorkHours
	}
	if true {
		toSerialize["plannedHours"] = o.PlannedHours
	}
	if o.Descriptions != nil {
		toSerialize["descriptions"] = o.Descriptions
	}
	return json.Marshal(toSerialize)
}

type NullableAvailability struct {
	value *Availability
	isSet bool
}

func (v NullableAvailability) Get() *Availability {
	return v.value
}

func (v *NullableAvailability) Set(val *Availability) {
	v.value = val
	v.isSet = true
}

func (v NullableAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvailability(val *Availability) *NullableAvailability {
	return &NullableAvailability{value: val, isSet: true}
}

func (v NullableAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


