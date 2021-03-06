# PersonSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PersonIds** | Pointer to **[]string** |  | [optional] 
**OfficeIds** | Pointer to **[]string** |  | [optional] 
**Availability** | Pointer to [**AvailabilityFilter**](AvailabilityFilter.md) |  | [optional] 
**OnsiteRatio** | Pointer to [**MinMaxPercent**](MinMaxPercent.md) |  | [optional] 
**Seniority** | Pointer to [**[]Seniority**](Seniority.md) |  | [optional] 
**Skills** | Pointer to [**[]PersonSearchSkillsInner**](PersonSearchSkillsInner.md) |  | [optional] 
**Projects** | Pointer to [**[]PersonSearchProjectsInner**](PersonSearchProjectsInner.md) |  | [optional] 
**Organizations** | Pointer to [**[]PersonSearchOrganizationsInner**](PersonSearchOrganizationsInner.md) |  | [optional] 
**Industries** | Pointer to [**[]PersonSearchIndustriesInner**](PersonSearchIndustriesInner.md) |  | [optional] 
**Certifications** | Pointer to [**[]PersonSearchCertificationsInner**](PersonSearchCertificationsInner.md) |  | [optional] 

## Methods

### NewPersonSearch

`func NewPersonSearch() *PersonSearch`

NewPersonSearch instantiates a new PersonSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPersonSearchWithDefaults

`func NewPersonSearchWithDefaults() *PersonSearch`

NewPersonSearchWithDefaults instantiates a new PersonSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPersonIds

`func (o *PersonSearch) GetPersonIds() []string`

GetPersonIds returns the PersonIds field if non-nil, zero value otherwise.

### GetPersonIdsOk

`func (o *PersonSearch) GetPersonIdsOk() (*[]string, bool)`

GetPersonIdsOk returns a tuple with the PersonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonIds

`func (o *PersonSearch) SetPersonIds(v []string)`

SetPersonIds sets PersonIds field to given value.

### HasPersonIds

`func (o *PersonSearch) HasPersonIds() bool`

HasPersonIds returns a boolean if a field has been set.

### GetOfficeIds

`func (o *PersonSearch) GetOfficeIds() []string`

GetOfficeIds returns the OfficeIds field if non-nil, zero value otherwise.

### GetOfficeIdsOk

`func (o *PersonSearch) GetOfficeIdsOk() (*[]string, bool)`

GetOfficeIdsOk returns a tuple with the OfficeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeIds

`func (o *PersonSearch) SetOfficeIds(v []string)`

SetOfficeIds sets OfficeIds field to given value.

### HasOfficeIds

`func (o *PersonSearch) HasOfficeIds() bool`

HasOfficeIds returns a boolean if a field has been set.

### GetAvailability

`func (o *PersonSearch) GetAvailability() AvailabilityFilter`

GetAvailability returns the Availability field if non-nil, zero value otherwise.

### GetAvailabilityOk

`func (o *PersonSearch) GetAvailabilityOk() (*AvailabilityFilter, bool)`

GetAvailabilityOk returns a tuple with the Availability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailability

`func (o *PersonSearch) SetAvailability(v AvailabilityFilter)`

SetAvailability sets Availability field to given value.

### HasAvailability

`func (o *PersonSearch) HasAvailability() bool`

HasAvailability returns a boolean if a field has been set.

### GetOnsiteRatio

`func (o *PersonSearch) GetOnsiteRatio() MinMaxPercent`

GetOnsiteRatio returns the OnsiteRatio field if non-nil, zero value otherwise.

### GetOnsiteRatioOk

`func (o *PersonSearch) GetOnsiteRatioOk() (*MinMaxPercent, bool)`

GetOnsiteRatioOk returns a tuple with the OnsiteRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnsiteRatio

`func (o *PersonSearch) SetOnsiteRatio(v MinMaxPercent)`

SetOnsiteRatio sets OnsiteRatio field to given value.

### HasOnsiteRatio

`func (o *PersonSearch) HasOnsiteRatio() bool`

HasOnsiteRatio returns a boolean if a field has been set.

### GetSeniority

`func (o *PersonSearch) GetSeniority() []Seniority`

GetSeniority returns the Seniority field if non-nil, zero value otherwise.

### GetSeniorityOk

`func (o *PersonSearch) GetSeniorityOk() (*[]Seniority, bool)`

GetSeniorityOk returns a tuple with the Seniority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniority

`func (o *PersonSearch) SetSeniority(v []Seniority)`

SetSeniority sets Seniority field to given value.

### HasSeniority

`func (o *PersonSearch) HasSeniority() bool`

HasSeniority returns a boolean if a field has been set.

### GetSkills

`func (o *PersonSearch) GetSkills() []PersonSearchSkillsInner`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *PersonSearch) GetSkillsOk() (*[]PersonSearchSkillsInner, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *PersonSearch) SetSkills(v []PersonSearchSkillsInner)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *PersonSearch) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### GetProjects

`func (o *PersonSearch) GetProjects() []PersonSearchProjectsInner`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *PersonSearch) GetProjectsOk() (*[]PersonSearchProjectsInner, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *PersonSearch) SetProjects(v []PersonSearchProjectsInner)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *PersonSearch) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetOrganizations

`func (o *PersonSearch) GetOrganizations() []PersonSearchOrganizationsInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *PersonSearch) GetOrganizationsOk() (*[]PersonSearchOrganizationsInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *PersonSearch) SetOrganizations(v []PersonSearchOrganizationsInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *PersonSearch) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.

### GetIndustries

`func (o *PersonSearch) GetIndustries() []PersonSearchIndustriesInner`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *PersonSearch) GetIndustriesOk() (*[]PersonSearchIndustriesInner, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *PersonSearch) SetIndustries(v []PersonSearchIndustriesInner)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *PersonSearch) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetCertifications

`func (o *PersonSearch) GetCertifications() []PersonSearchCertificationsInner`

GetCertifications returns the Certifications field if non-nil, zero value otherwise.

### GetCertificationsOk

`func (o *PersonSearch) GetCertificationsOk() (*[]PersonSearchCertificationsInner, bool)`

GetCertificationsOk returns a tuple with the Certifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifications

`func (o *PersonSearch) SetCertifications(v []PersonSearchCertificationsInner)`

SetCertifications sets Certifications field to given value.

### HasCertifications

`func (o *PersonSearch) HasCertifications() bool`

HasCertifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


