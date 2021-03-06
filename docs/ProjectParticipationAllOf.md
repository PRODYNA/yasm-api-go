# ProjectParticipationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeframe** | Pointer to [**Timeframed**](Timeframed.md) |  | [optional] 
**ProjectDetails** | Pointer to [**ProjectDetails**](ProjectDetails.md) |  | [optional] [readonly] 
**DescriptionOverwrite** | Pointer to **string** |  | [optional] 
**PersonalDescription** | Pointer to **string** |  | [optional] 
**Experiences** | [**[]Experience**](Experience.md) |  | [readonly] 

## Methods

### NewProjectParticipationAllOf

`func NewProjectParticipationAllOf(experiences []Experience, ) *ProjectParticipationAllOf`

NewProjectParticipationAllOf instantiates a new ProjectParticipationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectParticipationAllOfWithDefaults

`func NewProjectParticipationAllOfWithDefaults() *ProjectParticipationAllOf`

NewProjectParticipationAllOfWithDefaults instantiates a new ProjectParticipationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeframe

`func (o *ProjectParticipationAllOf) GetTimeframe() Timeframed`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *ProjectParticipationAllOf) GetTimeframeOk() (*Timeframed, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *ProjectParticipationAllOf) SetTimeframe(v Timeframed)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *ProjectParticipationAllOf) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetProjectDetails

`func (o *ProjectParticipationAllOf) GetProjectDetails() ProjectDetails`

GetProjectDetails returns the ProjectDetails field if non-nil, zero value otherwise.

### GetProjectDetailsOk

`func (o *ProjectParticipationAllOf) GetProjectDetailsOk() (*ProjectDetails, bool)`

GetProjectDetailsOk returns a tuple with the ProjectDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectDetails

`func (o *ProjectParticipationAllOf) SetProjectDetails(v ProjectDetails)`

SetProjectDetails sets ProjectDetails field to given value.

### HasProjectDetails

`func (o *ProjectParticipationAllOf) HasProjectDetails() bool`

HasProjectDetails returns a boolean if a field has been set.

### GetDescriptionOverwrite

`func (o *ProjectParticipationAllOf) GetDescriptionOverwrite() string`

GetDescriptionOverwrite returns the DescriptionOverwrite field if non-nil, zero value otherwise.

### GetDescriptionOverwriteOk

`func (o *ProjectParticipationAllOf) GetDescriptionOverwriteOk() (*string, bool)`

GetDescriptionOverwriteOk returns a tuple with the DescriptionOverwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionOverwrite

`func (o *ProjectParticipationAllOf) SetDescriptionOverwrite(v string)`

SetDescriptionOverwrite sets DescriptionOverwrite field to given value.

### HasDescriptionOverwrite

`func (o *ProjectParticipationAllOf) HasDescriptionOverwrite() bool`

HasDescriptionOverwrite returns a boolean if a field has been set.

### GetPersonalDescription

`func (o *ProjectParticipationAllOf) GetPersonalDescription() string`

GetPersonalDescription returns the PersonalDescription field if non-nil, zero value otherwise.

### GetPersonalDescriptionOk

`func (o *ProjectParticipationAllOf) GetPersonalDescriptionOk() (*string, bool)`

GetPersonalDescriptionOk returns a tuple with the PersonalDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalDescription

`func (o *ProjectParticipationAllOf) SetPersonalDescription(v string)`

SetPersonalDescription sets PersonalDescription field to given value.

### HasPersonalDescription

`func (o *ProjectParticipationAllOf) HasPersonalDescription() bool`

HasPersonalDescription returns a boolean if a field has been set.

### GetExperiences

`func (o *ProjectParticipationAllOf) GetExperiences() []Experience`

GetExperiences returns the Experiences field if non-nil, zero value otherwise.

### GetExperiencesOk

`func (o *ProjectParticipationAllOf) GetExperiencesOk() (*[]Experience, bool)`

GetExperiencesOk returns a tuple with the Experiences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperiences

`func (o *ProjectParticipationAllOf) SetExperiences(v []Experience)`

SetExperiences sets Experiences field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


