# EnvironmentSummaryViewsInner

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

### NewEnvironmentSummaryViewsInner

`func NewEnvironmentSummaryViewsInner() *EnvironmentSummaryViewsInner`

NewEnvironmentSummaryViewsInner instantiates a new EnvironmentSummaryViewsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSummaryViewsInnerWithDefaults

`func NewEnvironmentSummaryViewsInnerWithDefaults() *EnvironmentSummaryViewsInner`

NewEnvironmentSummaryViewsInnerWithDefaults instantiates a new EnvironmentSummaryViewsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentSummaryViewsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentSummaryViewsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentSummaryViewsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentSummaryViewsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentSummaryViewsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentSummaryViewsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentSummaryViewsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentSummaryViewsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImageName

`func (o *EnvironmentSummaryViewsInner) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *EnvironmentSummaryViewsInner) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *EnvironmentSummaryViewsInner) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *EnvironmentSummaryViewsInner) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageUrl

`func (o *EnvironmentSummaryViewsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *EnvironmentSummaryViewsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *EnvironmentSummaryViewsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *EnvironmentSummaryViewsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetExportTimestamp

`func (o *EnvironmentSummaryViewsInner) GetExportTimestamp() int32`

GetExportTimestamp returns the ExportTimestamp field if non-nil, zero value otherwise.

### GetExportTimestampOk

`func (o *EnvironmentSummaryViewsInner) GetExportTimestampOk() (*int32, bool)`

GetExportTimestampOk returns a tuple with the ExportTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTimestamp

`func (o *EnvironmentSummaryViewsInner) SetExportTimestamp(v int32)`

SetExportTimestamp sets ExportTimestamp field to given value.

### HasExportTimestamp

`func (o *EnvironmentSummaryViewsInner) HasExportTimestamp() bool`

HasExportTimestamp returns a boolean if a field has been set.

### GetRevisionId

`func (o *EnvironmentSummaryViewsInner) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *EnvironmentSummaryViewsInner) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *EnvironmentSummaryViewsInner) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *EnvironmentSummaryViewsInner) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetRegions

`func (o *EnvironmentSummaryViewsInner) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *EnvironmentSummaryViewsInner) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *EnvironmentSummaryViewsInner) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *EnvironmentSummaryViewsInner) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetEmpty

`func (o *EnvironmentSummaryViewsInner) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *EnvironmentSummaryViewsInner) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *EnvironmentSummaryViewsInner) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *EnvironmentSummaryViewsInner) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


