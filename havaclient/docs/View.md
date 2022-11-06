# View

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the view | [optional] 
**Type** | Pointer to **string** | The type of view | [optional] 
**ImageName** | Pointer to **string** | The name of the exported thumbnail for this view | [optional] 
**ImageUrl** | Pointer to **string** | The location of the exported thumbnail for this view | [optional] 
**ExportTimestamp** | Pointer to **int32** | The creation timestamp for the exported thumbnail | [optional] 
**RevisionId** | Pointer to **string** | The unique ID of the current version of this view | [optional] 
**Regions** | Pointer to **[]string** | The regions or locations displayed in this view | [optional] 
**Empty** | Pointer to **bool** | Whether the view is considered empty and has no valuable resources to display | [optional] 
**Resources** | Pointer to [**[]ViewResources**](ViewResources.md) | A list of resource ID&#39;s to be displayed. Only returned for list view. | [optional] 

## Methods

### NewView

`func NewView() *View`

NewView instantiates a new View object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewWithDefaults

`func NewViewWithDefaults() *View`

NewViewWithDefaults instantiates a new View object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *View) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *View) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *View) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *View) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *View) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *View) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *View) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *View) HasType() bool`

HasType returns a boolean if a field has been set.

### GetImageName

`func (o *View) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *View) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *View) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *View) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetImageUrl

`func (o *View) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *View) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *View) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *View) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetExportTimestamp

`func (o *View) GetExportTimestamp() int32`

GetExportTimestamp returns the ExportTimestamp field if non-nil, zero value otherwise.

### GetExportTimestampOk

`func (o *View) GetExportTimestampOk() (*int32, bool)`

GetExportTimestampOk returns a tuple with the ExportTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportTimestamp

`func (o *View) SetExportTimestamp(v int32)`

SetExportTimestamp sets ExportTimestamp field to given value.

### HasExportTimestamp

`func (o *View) HasExportTimestamp() bool`

HasExportTimestamp returns a boolean if a field has been set.

### GetRevisionId

`func (o *View) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *View) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *View) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *View) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.

### GetRegions

`func (o *View) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *View) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *View) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *View) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetEmpty

`func (o *View) GetEmpty() bool`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *View) GetEmptyOk() (*bool, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *View) SetEmpty(v bool)`

SetEmpty sets Empty field to given value.

### HasEmpty

`func (o *View) HasEmpty() bool`

HasEmpty returns a boolean if a field has been set.

### GetResources

`func (o *View) GetResources() []ViewResources`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *View) GetResourcesOk() (*[]ViewResources, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *View) SetResources(v []ViewResources)`

SetResources sets Resources field to given value.

### HasResources

`func (o *View) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


