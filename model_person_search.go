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

// PersonSearch struct for PersonSearch
type PersonSearch struct {
	PersonIds []string `json:"personIds,omitempty"`
	OfficeIds []string `json:"officeIds,omitempty"`
	Availability *AvailabilityFilter `json:"availability,omitempty"`
	OnsiteRatio *MinMaxPercent `json:"onsiteRatio,omitempty"`
	Seniority []Seniority `json:"seniority,omitempty"`
	Skills []PersonSearchSkillsInner `json:"skills,omitempty"`
	Projects []PersonSearchProjectsInner `json:"projects,omitempty"`
	Organizations []PersonSearchOrganizationsInner `json:"organizations,omitempty"`
	Industries []PersonSearchIndustriesInner `json:"industries,omitempty"`
	Certifications []PersonSearchCertificationsInner `json:"certifications,omitempty"`
}

// NewPersonSearch instantiates a new PersonSearch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPersonSearch() *PersonSearch {
	this := PersonSearch{}
	return &this
}

// NewPersonSearchWithDefaults instantiates a new PersonSearch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPersonSearchWithDefaults() *PersonSearch {
	this := PersonSearch{}
	return &this
}

// GetPersonIds returns the PersonIds field value if set, zero value otherwise.
func (o *PersonSearch) GetPersonIds() []string {
	if o == nil || o.PersonIds == nil {
		var ret []string
		return ret
	}
	return o.PersonIds
}

// GetPersonIdsOk returns a tuple with the PersonIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetPersonIdsOk() ([]string, bool) {
	if o == nil || o.PersonIds == nil {
		return nil, false
	}
	return o.PersonIds, true
}

// HasPersonIds returns a boolean if a field has been set.
func (o *PersonSearch) HasPersonIds() bool {
	if o != nil && o.PersonIds != nil {
		return true
	}

	return false
}

// SetPersonIds gets a reference to the given []string and assigns it to the PersonIds field.
func (o *PersonSearch) SetPersonIds(v []string) {
	o.PersonIds = v
}

// GetOfficeIds returns the OfficeIds field value if set, zero value otherwise.
func (o *PersonSearch) GetOfficeIds() []string {
	if o == nil || o.OfficeIds == nil {
		var ret []string
		return ret
	}
	return o.OfficeIds
}

// GetOfficeIdsOk returns a tuple with the OfficeIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetOfficeIdsOk() ([]string, bool) {
	if o == nil || o.OfficeIds == nil {
		return nil, false
	}
	return o.OfficeIds, true
}

// HasOfficeIds returns a boolean if a field has been set.
func (o *PersonSearch) HasOfficeIds() bool {
	if o != nil && o.OfficeIds != nil {
		return true
	}

	return false
}

// SetOfficeIds gets a reference to the given []string and assigns it to the OfficeIds field.
func (o *PersonSearch) SetOfficeIds(v []string) {
	o.OfficeIds = v
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *PersonSearch) GetAvailability() AvailabilityFilter {
	if o == nil || o.Availability == nil {
		var ret AvailabilityFilter
		return ret
	}
	return *o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetAvailabilityOk() (*AvailabilityFilter, bool) {
	if o == nil || o.Availability == nil {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *PersonSearch) HasAvailability() bool {
	if o != nil && o.Availability != nil {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given AvailabilityFilter and assigns it to the Availability field.
func (o *PersonSearch) SetAvailability(v AvailabilityFilter) {
	o.Availability = &v
}

// GetOnsiteRatio returns the OnsiteRatio field value if set, zero value otherwise.
func (o *PersonSearch) GetOnsiteRatio() MinMaxPercent {
	if o == nil || o.OnsiteRatio == nil {
		var ret MinMaxPercent
		return ret
	}
	return *o.OnsiteRatio
}

// GetOnsiteRatioOk returns a tuple with the OnsiteRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetOnsiteRatioOk() (*MinMaxPercent, bool) {
	if o == nil || o.OnsiteRatio == nil {
		return nil, false
	}
	return o.OnsiteRatio, true
}

// HasOnsiteRatio returns a boolean if a field has been set.
func (o *PersonSearch) HasOnsiteRatio() bool {
	if o != nil && o.OnsiteRatio != nil {
		return true
	}

	return false
}

// SetOnsiteRatio gets a reference to the given MinMaxPercent and assigns it to the OnsiteRatio field.
func (o *PersonSearch) SetOnsiteRatio(v MinMaxPercent) {
	o.OnsiteRatio = &v
}

// GetSeniority returns the Seniority field value if set, zero value otherwise.
func (o *PersonSearch) GetSeniority() []Seniority {
	if o == nil || o.Seniority == nil {
		var ret []Seniority
		return ret
	}
	return o.Seniority
}

// GetSeniorityOk returns a tuple with the Seniority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetSeniorityOk() ([]Seniority, bool) {
	if o == nil || o.Seniority == nil {
		return nil, false
	}
	return o.Seniority, true
}

// HasSeniority returns a boolean if a field has been set.
func (o *PersonSearch) HasSeniority() bool {
	if o != nil && o.Seniority != nil {
		return true
	}

	return false
}

// SetSeniority gets a reference to the given []Seniority and assigns it to the Seniority field.
func (o *PersonSearch) SetSeniority(v []Seniority) {
	o.Seniority = v
}

// GetSkills returns the Skills field value if set, zero value otherwise.
func (o *PersonSearch) GetSkills() []PersonSearchSkillsInner {
	if o == nil || o.Skills == nil {
		var ret []PersonSearchSkillsInner
		return ret
	}
	return o.Skills
}

// GetSkillsOk returns a tuple with the Skills field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetSkillsOk() ([]PersonSearchSkillsInner, bool) {
	if o == nil || o.Skills == nil {
		return nil, false
	}
	return o.Skills, true
}

// HasSkills returns a boolean if a field has been set.
func (o *PersonSearch) HasSkills() bool {
	if o != nil && o.Skills != nil {
		return true
	}

	return false
}

// SetSkills gets a reference to the given []PersonSearchSkillsInner and assigns it to the Skills field.
func (o *PersonSearch) SetSkills(v []PersonSearchSkillsInner) {
	o.Skills = v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *PersonSearch) GetProjects() []PersonSearchProjectsInner {
	if o == nil || o.Projects == nil {
		var ret []PersonSearchProjectsInner
		return ret
	}
	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetProjectsOk() ([]PersonSearchProjectsInner, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *PersonSearch) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given []PersonSearchProjectsInner and assigns it to the Projects field.
func (o *PersonSearch) SetProjects(v []PersonSearchProjectsInner) {
	o.Projects = v
}

// GetOrganizations returns the Organizations field value if set, zero value otherwise.
func (o *PersonSearch) GetOrganizations() []PersonSearchOrganizationsInner {
	if o == nil || o.Organizations == nil {
		var ret []PersonSearchOrganizationsInner
		return ret
	}
	return o.Organizations
}

// GetOrganizationsOk returns a tuple with the Organizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetOrganizationsOk() ([]PersonSearchOrganizationsInner, bool) {
	if o == nil || o.Organizations == nil {
		return nil, false
	}
	return o.Organizations, true
}

// HasOrganizations returns a boolean if a field has been set.
func (o *PersonSearch) HasOrganizations() bool {
	if o != nil && o.Organizations != nil {
		return true
	}

	return false
}

// SetOrganizations gets a reference to the given []PersonSearchOrganizationsInner and assigns it to the Organizations field.
func (o *PersonSearch) SetOrganizations(v []PersonSearchOrganizationsInner) {
	o.Organizations = v
}

// GetIndustries returns the Industries field value if set, zero value otherwise.
func (o *PersonSearch) GetIndustries() []PersonSearchIndustriesInner {
	if o == nil || o.Industries == nil {
		var ret []PersonSearchIndustriesInner
		return ret
	}
	return o.Industries
}

// GetIndustriesOk returns a tuple with the Industries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetIndustriesOk() ([]PersonSearchIndustriesInner, bool) {
	if o == nil || o.Industries == nil {
		return nil, false
	}
	return o.Industries, true
}

// HasIndustries returns a boolean if a field has been set.
func (o *PersonSearch) HasIndustries() bool {
	if o != nil && o.Industries != nil {
		return true
	}

	return false
}

// SetIndustries gets a reference to the given []PersonSearchIndustriesInner and assigns it to the Industries field.
func (o *PersonSearch) SetIndustries(v []PersonSearchIndustriesInner) {
	o.Industries = v
}

// GetCertifications returns the Certifications field value if set, zero value otherwise.
func (o *PersonSearch) GetCertifications() []PersonSearchCertificationsInner {
	if o == nil || o.Certifications == nil {
		var ret []PersonSearchCertificationsInner
		return ret
	}
	return o.Certifications
}

// GetCertificationsOk returns a tuple with the Certifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PersonSearch) GetCertificationsOk() ([]PersonSearchCertificationsInner, bool) {
	if o == nil || o.Certifications == nil {
		return nil, false
	}
	return o.Certifications, true
}

// HasCertifications returns a boolean if a field has been set.
func (o *PersonSearch) HasCertifications() bool {
	if o != nil && o.Certifications != nil {
		return true
	}

	return false
}

// SetCertifications gets a reference to the given []PersonSearchCertificationsInner and assigns it to the Certifications field.
func (o *PersonSearch) SetCertifications(v []PersonSearchCertificationsInner) {
	o.Certifications = v
}

func (o PersonSearch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PersonIds != nil {
		toSerialize["personIds"] = o.PersonIds
	}
	if o.OfficeIds != nil {
		toSerialize["officeIds"] = o.OfficeIds
	}
	if o.Availability != nil {
		toSerialize["availability"] = o.Availability
	}
	if o.OnsiteRatio != nil {
		toSerialize["onsiteRatio"] = o.OnsiteRatio
	}
	if o.Seniority != nil {
		toSerialize["seniority"] = o.Seniority
	}
	if o.Skills != nil {
		toSerialize["skills"] = o.Skills
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.Organizations != nil {
		toSerialize["organizations"] = o.Organizations
	}
	if o.Industries != nil {
		toSerialize["industries"] = o.Industries
	}
	if o.Certifications != nil {
		toSerialize["certifications"] = o.Certifications
	}
	return json.Marshal(toSerialize)
}

type NullablePersonSearch struct {
	value *PersonSearch
	isSet bool
}

func (v NullablePersonSearch) Get() *PersonSearch {
	return v.value
}

func (v *NullablePersonSearch) Set(val *PersonSearch) {
	v.value = val
	v.isSet = true
}

func (v NullablePersonSearch) IsSet() bool {
	return v.isSet
}

func (v *NullablePersonSearch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePersonSearch(val *PersonSearch) *NullablePersonSearch {
	return &NullablePersonSearch{value: val, isSet: true}
}

func (v NullablePersonSearch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePersonSearch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


