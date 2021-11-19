/*
YASM (Yet Another Skill Management) API

This is an example of using OAuth2 Implicit Flow in a specification to describe security to your API.

API version: 0.8.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Person struct for Person
type Person struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Suggestion bool `json:"suggestion"`
	Synonyms *[]string `json:"synonyms,omitempty"`
	Location *string `json:"location,omitempty"`
	Geolocation *Geolocation `json:"geolocation,omitempty"`
	EmployeeId *string `json:"employeeId,omitempty"`
	JobTitle *string `json:"jobTitle,omitempty"`
	Company *string `json:"company,omitempty"`
	Department *string `json:"department,omitempty"`
	Mail *string `json:"mail,omitempty"`
	MobilePhone *string `json:"mobilePhone,omitempty"`
	// base64 encoded image
	Picture *string `json:"picture,omitempty"`
	// Marks persons not working for the company anymore
	Inactive *bool `json:"inactive,omitempty"`
}

// NewPerson instantiates a new Person object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerson(id string, name string, suggestion bool) *Person {
	this := Person{}
	this.Id = id
	this.Name = name
	this.Suggestion = suggestion
	var inactive bool = false
	this.Inactive = &inactive
	return &this
}

// NewPersonWithDefaults instantiates a new Person object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonWithDefaults() *Person {
	this := Person{}
	var suggestion bool = false
	this.Suggestion = suggestion
	var inactive bool = false
	this.Inactive = &inactive
	return &this
}

// GetId returns the Id field value
func (o *Person) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Person) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Person) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Person) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Person) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Person) SetName(v string) {
	o.Name = v
}

// GetSuggestion returns the Suggestion field value
func (o *Person) GetSuggestion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Suggestion
}

// GetSuggestionOk returns a tuple with the Suggestion field value
// and a boolean to check if the value has been set.
func (o *Person) GetSuggestionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Suggestion, true
}

// SetSuggestion sets field value
func (o *Person) SetSuggestion(v bool) {
	o.Suggestion = v
}

// GetSynonyms returns the Synonyms field value if set, zero value otherwise.
func (o *Person) GetSynonyms() []string {
	if o == nil || o.Synonyms == nil {
		var ret []string
		return ret
	}
	return *o.Synonyms
}

// GetSynonymsOk returns a tuple with the Synonyms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetSynonymsOk() (*[]string, bool) {
	if o == nil || o.Synonyms == nil {
		return nil, false
	}
	return o.Synonyms, true
}

// HasSynonyms returns a boolean if a field has been set.
func (o *Person) HasSynonyms() bool {
	if o != nil && o.Synonyms != nil {
		return true
	}

	return false
}

// SetSynonyms gets a reference to the given []string and assigns it to the Synonyms field.
func (o *Person) SetSynonyms(v []string) {
	o.Synonyms = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Person) GetLocation() string {
	if o == nil || o.Location == nil {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetLocationOk() (*string, bool) {
	if o == nil || o.Location == nil {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Person) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Person) SetLocation(v string) {
	o.Location = &v
}

// GetGeolocation returns the Geolocation field value if set, zero value otherwise.
func (o *Person) GetGeolocation() Geolocation {
	if o == nil || o.Geolocation == nil {
		var ret Geolocation
		return ret
	}
	return *o.Geolocation
}

// GetGeolocationOk returns a tuple with the Geolocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetGeolocationOk() (*Geolocation, bool) {
	if o == nil || o.Geolocation == nil {
		return nil, false
	}
	return o.Geolocation, true
}

// HasGeolocation returns a boolean if a field has been set.
func (o *Person) HasGeolocation() bool {
	if o != nil && o.Geolocation != nil {
		return true
	}

	return false
}

// SetGeolocation gets a reference to the given Geolocation and assigns it to the Geolocation field.
func (o *Person) SetGeolocation(v Geolocation) {
	o.Geolocation = &v
}

// GetEmployeeId returns the EmployeeId field value if set, zero value otherwise.
func (o *Person) GetEmployeeId() string {
	if o == nil || o.EmployeeId == nil {
		var ret string
		return ret
	}
	return *o.EmployeeId
}

// GetEmployeeIdOk returns a tuple with the EmployeeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetEmployeeIdOk() (*string, bool) {
	if o == nil || o.EmployeeId == nil {
		return nil, false
	}
	return o.EmployeeId, true
}

// HasEmployeeId returns a boolean if a field has been set.
func (o *Person) HasEmployeeId() bool {
	if o != nil && o.EmployeeId != nil {
		return true
	}

	return false
}

// SetEmployeeId gets a reference to the given string and assigns it to the EmployeeId field.
func (o *Person) SetEmployeeId(v string) {
	o.EmployeeId = &v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *Person) GetJobTitle() string {
	if o == nil || o.JobTitle == nil {
		var ret string
		return ret
	}
	return *o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetJobTitleOk() (*string, bool) {
	if o == nil || o.JobTitle == nil {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *Person) HasJobTitle() bool {
	if o != nil && o.JobTitle != nil {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given string and assigns it to the JobTitle field.
func (o *Person) SetJobTitle(v string) {
	o.JobTitle = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *Person) GetCompany() string {
	if o == nil || o.Company == nil {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetCompanyOk() (*string, bool) {
	if o == nil || o.Company == nil {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *Person) HasCompany() bool {
	if o != nil && o.Company != nil {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *Person) SetCompany(v string) {
	o.Company = &v
}

// GetDepartment returns the Department field value if set, zero value otherwise.
func (o *Person) GetDepartment() string {
	if o == nil || o.Department == nil {
		var ret string
		return ret
	}
	return *o.Department
}

// GetDepartmentOk returns a tuple with the Department field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetDepartmentOk() (*string, bool) {
	if o == nil || o.Department == nil {
		return nil, false
	}
	return o.Department, true
}

// HasDepartment returns a boolean if a field has been set.
func (o *Person) HasDepartment() bool {
	if o != nil && o.Department != nil {
		return true
	}

	return false
}

// SetDepartment gets a reference to the given string and assigns it to the Department field.
func (o *Person) SetDepartment(v string) {
	o.Department = &v
}

// GetMail returns the Mail field value if set, zero value otherwise.
func (o *Person) GetMail() string {
	if o == nil || o.Mail == nil {
		var ret string
		return ret
	}
	return *o.Mail
}

// GetMailOk returns a tuple with the Mail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetMailOk() (*string, bool) {
	if o == nil || o.Mail == nil {
		return nil, false
	}
	return o.Mail, true
}

// HasMail returns a boolean if a field has been set.
func (o *Person) HasMail() bool {
	if o != nil && o.Mail != nil {
		return true
	}

	return false
}

// SetMail gets a reference to the given string and assigns it to the Mail field.
func (o *Person) SetMail(v string) {
	o.Mail = &v
}

// GetMobilePhone returns the MobilePhone field value if set, zero value otherwise.
func (o *Person) GetMobilePhone() string {
	if o == nil || o.MobilePhone == nil {
		var ret string
		return ret
	}
	return *o.MobilePhone
}

// GetMobilePhoneOk returns a tuple with the MobilePhone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetMobilePhoneOk() (*string, bool) {
	if o == nil || o.MobilePhone == nil {
		return nil, false
	}
	return o.MobilePhone, true
}

// HasMobilePhone returns a boolean if a field has been set.
func (o *Person) HasMobilePhone() bool {
	if o != nil && o.MobilePhone != nil {
		return true
	}

	return false
}

// SetMobilePhone gets a reference to the given string and assigns it to the MobilePhone field.
func (o *Person) SetMobilePhone(v string) {
	o.MobilePhone = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *Person) GetPicture() string {
	if o == nil || o.Picture == nil {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetPictureOk() (*string, bool) {
	if o == nil || o.Picture == nil {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *Person) HasPicture() bool {
	if o != nil && o.Picture != nil {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *Person) SetPicture(v string) {
	o.Picture = &v
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *Person) GetInactive() bool {
	if o == nil || o.Inactive == nil {
		var ret bool
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Person) GetInactiveOk() (*bool, bool) {
	if o == nil || o.Inactive == nil {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *Person) HasInactive() bool {
	if o != nil && o.Inactive != nil {
		return true
	}

	return false
}

// SetInactive gets a reference to the given bool and assigns it to the Inactive field.
func (o *Person) SetInactive(v bool) {
	o.Inactive = &v
}

func (o Person) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["suggestion"] = o.Suggestion
	}
	if o.Synonyms != nil {
		toSerialize["synonyms"] = o.Synonyms
	}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Geolocation != nil {
		toSerialize["geolocation"] = o.Geolocation
	}
	if o.EmployeeId != nil {
		toSerialize["employeeId"] = o.EmployeeId
	}
	if o.JobTitle != nil {
		toSerialize["jobTitle"] = o.JobTitle
	}
	if o.Company != nil {
		toSerialize["company"] = o.Company
	}
	if o.Department != nil {
		toSerialize["department"] = o.Department
	}
	if o.Mail != nil {
		toSerialize["mail"] = o.Mail
	}
	if o.MobilePhone != nil {
		toSerialize["mobilePhone"] = o.MobilePhone
	}
	if o.Picture != nil {
		toSerialize["picture"] = o.Picture
	}
	if o.Inactive != nil {
		toSerialize["inactive"] = o.Inactive
	}
	return json.Marshal(toSerialize)
}

type NullablePerson struct {
	value *Person
	isSet bool
}

func (v NullablePerson) Get() *Person {
	return v.value
}

func (v *NullablePerson) Set(val *Person) {
	v.value = val
	v.isSet = true
}

func (v NullablePerson) IsSet() bool {
	return v.isSet
}

func (v *NullablePerson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerson(val *Person) *NullablePerson {
	return &NullablePerson{value: val, isSet: true}
}

func (v NullablePerson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


