# EnvironmentSummaryViews

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the view | [optional] 
**Type** | Pointer to **string** | The type of view | [optional] 
**ImageName** | Pointer to **string** | The name of the exported thumbnail for this view | [optional] 
**ImageUrl** | Pointer to **string** | The location of the exported thumbnail for this view | [optional] 
**ExportTimestamp** | Pointer to **int32** | The creation timestamp for the exported thumnail | [optional] 
**RevisionId** | Pointer to **string** | The unique ID of the current version of this view | [optional] 
**Regions** | Pointer to **[]string** | The regions or locations displayed in this view | [optional] 
**Empty** | Pointer to **bool** | Whether the view is considered empty and has no valuable resources to display | [optional] 

## Methods

### NewEnvironmentSummaryViews

`func NewEnvironmentSummaryViews() *EnvironmentSummaryViews`

NewEnvironmentSummaryViews instantiates a new EnvironmentSummaryViews object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSummaryViewsWithDefaults

`func NewEnvironmentSummaryViewsWithDefaults() *EnvironmentSummaryViews`

NewEnvironmentSummaryViewsWithDefaults instantiates a new EnvironmentSummaryViews object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentSummaryViews) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentSummaryViews) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentSummaryViews) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentSummaryViews) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentSummaryViews) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentSummaryViews) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentSummaryViews) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentSummaryViews) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImageName

`func (o *EnvironmentSummaryViews) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *EnvironmentSummaryViews) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *EnvironmentSummaryViews) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *EnvironmentSummaryViews) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageUrl

`func (o *EnvironmentSummaryViews) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *EnvironmentSummaryViews) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *EnvironmentSummaryViews) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *EnvironmentSummaryViews) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetExportTimestamp

`func (o *EnvironmentSummaryViews) GetExportTimestamp() int32`

GetExportTimestamp returns the ExportTimestamp field if non-nil, zero value otherwise.

### GetExportTimestampOk

`func (o *EnvironmentSummaryViews) GetExportTimestampOk() (*int32, bool)`

GetExportTimestampOk returns a tuple with the ExportTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTimestamp

`func (o *EnvironmentSummaryViews) SetExportTimestamp(v int32)`

SetExportTimestamp sets ExportTimestamp field to given value.

### HasExportTimestamp

`func (o *EnvironmentSummaryViews) HasExportTimestamp() bool`

HasExportTimestamp returns a boolean if a field has been set.

### GetRevisionId

`func (o *EnvironmentSummaryViews) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *EnvironmentSummaryViews) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *EnvironmentSummaryViews) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *EnvironmentSummaryViews) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetRegions

`func (o *EnvironmentSummaryViews) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *EnvironmentSummaryViews) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *EnvironmentSummaryViews) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *EnvironmentSummaryViews) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetEmpty

`func (o *EnvironmentSummaryViews) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *EnvironmentSummaryViews) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *EnvironmentSummaryViews) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *EnvironmentSummaryViews) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


