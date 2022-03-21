/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 1.0.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// OrganizationAllOf struct for OrganizationAllOf
type OrganizationAllOf struct {
	Partner *bool `json:"partner,omitempty"`
	Customer *bool `json:"customer,omitempty"`
	// base64 encoded image
	Picture *string `json:"picture,omitempty"`
}

// NewOrganizationAllOf instantiates a new OrganizationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationAllOf() *OrganizationAllOf {
	this := OrganizationAllOf{}
	var partner bool = false
	this.Partner = &partner
	var customer bool = false
	this.Customer = &customer
	return &this
}

// NewOrganizationAllOfWithDefaults instantiates a new OrganizationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationAllOfWithDefaults() *OrganizationAllOf {
	this := OrganizationAllOf{}
	var partner bool = false
	this.Partner = &partner
	var customer bool = false
	this.Customer = &customer
	return &this
}

// GetPartner returns the Partner field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetPartner() bool {
	if o == nil || o.Partner == nil {
		var ret bool
		return ret
	}
	return *o.Partner
}

// GetPartnerOk returns a tuple with the Partner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetPartnerOk() (*bool, bool) {
	if o == nil || o.Partner == nil {
		return nil, false
	}
	return o.Partner, true
}

// HasPartner returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasPartner() bool {
	if o != nil && o.Partner != nil {
		return true
	}

	return false
}

// SetPartner gets a reference to the given bool and assigns it to the Partner field.
func (o *OrganizationAllOf) SetPartner(v bool) {
	o.Partner = &v
}

// GetCustomer returns the Customer field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetCustomer() bool {
	if o == nil || o.Customer == nil {
		var ret bool
		return ret
	}
	return *o.Customer
}

// GetCustomerOk returns a tuple with the Customer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetCustomerOk() (*bool, bool) {
	if o == nil || o.Customer == nil {
		return nil, false
	}
	return o.Customer, true
}

// HasCustomer returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasCustomer() bool {
	if o != nil && o.Customer != nil {
		return true
	}

	return false
}

// SetCustomer gets a reference to the given bool and assigns it to the Customer field.
func (o *OrganizationAllOf) SetCustomer(v bool) {
	o.Customer = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *OrganizationAllOf) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationAllOf) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *OrganizationAllOf) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *OrganizationAllOf) SetPicture(v string) {
	o.Picture = &v
}

func (o OrganizationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Partner != nil {
		toSerialize["partner"] = o.Partner
	}
	if o.Customer != nil {
		toSerialize["customer"] = o.Customer
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	return json.Marshal(toSerialize)
}

type NullableOrganizationAllOf struct {
	value *OrganizationAllOf
	isSet bool
}

func (v NullableOrganizationAllOf) Get() *OrganizationAllOf {
	return v.value
}

func (v *NullableOrganizationAllOf) Set(val *OrganizationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationAllOf(val *OrganizationAllOf) *NullableOrganizationAllOf {
	return &NullableOrganizationAllOf{value: val, isSet: true}
}

func (v NullableOrganizationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


