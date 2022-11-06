# EnvironmentLatestRevisions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitValue** | Pointer to **int32** | The amount of results per page | [optional] 
**CurrentPage** | Pointer to **int32** | The current page | [optional] 
**TotalPages** | Pointer to **int32** | The total pages, based on limit and total versions | [optional] 
**TotalCount** | Pointer to **int32** | The total versions available | [optional] 
**Map** | Pointer to [**[]EnvironmentRevision**](EnvironmentRevision.md) |  | [optional] 

## Methods

### NewEnvironmentLatestRevisions

`func NewEnvironmentLatestRevisions() *EnvironmentLatestRevisions`

NewEnvironmentLatestRevisions instantiates a new EnvironmentLatestRevisions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentLatestRevisionsWithDefaults

`func NewEnvironmentLatestRevisionsWithDefaults() *EnvironmentLatestRevisions`

NewEnvironmentLatestRevisionsWithDefaults instantiates a new EnvironmentLatestRevisions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitValue

`func (o *EnvironmentLatestRevisions) GetLimitValue() int32`

GetLimitValue returns the LimitValue field if non-nil, zero value otherwise.

### GetLimitValueOk

`func (o *EnvironmentLatestRevisions) GetLimitValueOk() (*int32, bool)`

GetLimitValueOk returns a tuple with the LimitValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitValue

`func (o *EnvironmentLatestRevisions) SetLimitValue(v int32)`

SetLimitValue sets LimitValue field to given value.

### HasLimitValue

`func (o *EnvironmentLatestRevisions) HasLimitValue() bool`

HasLimitValue returns a boolean if a field has been set.

### GetCurrentPage

`func (o *EnvironmentLatestRevisions) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *EnvironmentLatestRevisions) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *EnvironmentLatestRevisions) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *EnvironmentLatestRevisions) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetTotalPages

`func (o *EnvironmentLatestRevisions) GetTotalPages() int32`

GetTotalPages returns the TotalPages field if non-nil, zero value otherwise.

### GetTotalPagesOk

`func (o *EnvironmentLatestRevisions) GetTotalPagesOk() (*int32, bool)`

GetTotalPagesOk returns a tuple with the TotalPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPages

`func (o *EnvironmentLatestRevisions) SetTotalPages(v int32)`

SetTotalPages sets TotalPages field to given value.

### HasTotalPages

`func (o *EnvironmentLatestRevisions) HasTotalPages() bool`

HasTotalPages returns a boolean if a field has been set.

### GetTotalCount

`func (o *EnvironmentLatestRevisions) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *EnvironmentLatestRevisions) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *EnvironmentLatestRevisions) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *EnvironmentLatestRevisions) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.

### GetMap

`func (o *EnvironmentLatestRevisions) GetMap() []EnvironmentRevision`

GetMap returns the Map field if non-nil, zero value otherwise.

### GetMapOk

`func (o *EnvironmentLatestRevisions) GetMapOk() (*[]EnvironmentRevision, bool)`

GetMapOk returns a tuple with the Map field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMap

`func (o *EnvironmentLatestRevisions) SetMap(v []EnvironmentRevision)`

SetMap sets Map field to given value.

### HasMap

`func (o *EnvironmentLatestRevisions) HasMap() bool`

HasMap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


