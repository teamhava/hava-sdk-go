# ProjectsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The updated name for the project | [optional] 
**Description** | Pointer to **string** | The updated description of the project | [optional] 

## Methods

### NewProjectsUpdateRequest

`func NewProjectsUpdateRequest() *ProjectsUpdateRequest`

NewProjectsUpdateRequest instantiates a new ProjectsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectsUpdateRequestWithDefaults

`func NewProjectsUpdateRequestWithDefaults() *ProjectsUpdateRequest`

NewProjectsUpdateRequestWithDefaults instantiates a new ProjectsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectsUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectsUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectsUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectsUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectsUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectsUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectsUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectsUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


