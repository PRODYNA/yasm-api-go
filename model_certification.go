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

// Certification struct for Certification
type Certification struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Synonyms *[]string `json:"synonyms,omitempty"`
	Validity *string `json:"validity,omitempty"`
}

// NewCertification instantiates a new Certification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertification(id string, name string) *Certification {
	this := Certification{}
	this.Id = id
	this.Name = name
	return &this
}

// NewCertificationWithDefaults instantiates a new Certification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificationWithDefaults() *Certification {
	this := Certification{}
	return &this
}

// GetId returns the Id field value
func (o *Certification) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Certification) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Certification) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Certification) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Certification) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Certification) SetName(v string) {
	o.Name = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Certification) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return *o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetSynonymsOk() (*[]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Certification) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Certification) SetSynonyms(v []string) {
	o.Synonyms = &v
}

// GetValidity returns the Validity field value if set, zero value otherwise.
func (o *Certification) GetValidity() string {
	if o == nil || o.Validity == nil {
		var ret string
		return ret
	}
	return *o.Validity
}

// GetValidityOk returns a tuple with the Validity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Certification) GetValidityOk() (*string, bool) {
	if o == nil || o.Validity == nil {
		return nil, false
	}
	return o.Validity, true
}

// HasValidity returns a boolean if a field has been set.
func (o *Certification) HasValidity() bool {
	if o != nil && o.Validity != nil {
		return true
	}

	return false
}

// SetValidity gets a reference to the given string and assigns it to the Validity field.
func (o *Certification) SetValidity(v string) {
	o.Validity = &v
}

func (o Certification) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Synonyms != nil {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.Validity != nil {
		toSerialize["validity"] = o.Validity
	}
	return json.Marshal(toSerialize)
}

type NullableCertification struct {
	value *Certification
	isSet bool
}

func (v NullableCertification) Get() *Certification {
	return v.value
}

func (v *NullableCertification) Set(val *Certification) {
	v.value = val
	v.isSet = true
}

func (v NullableCertification) IsSet() bool {
	return v.isSet
}

func (v *NullableCertification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertification(val *Certification) *NullableCertification {
	return &NullableCertification{value: val, isSet: true}
}

func (v NullableCertification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

