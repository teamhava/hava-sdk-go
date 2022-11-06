# EnvironmentRevision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the environment version | [optional] 
**From** | Pointer to **string** | The start of the date range this version covers | [optional] 
**To** | Pointer to **string** | The end date of the range this version covers. Will be empty if it&#39;s the latest version. | [optional] 

## Methods

### NewEnvironmentRevision

`func NewEnvironmentRevision() *EnvironmentRevision`

NewEnvironmentRevision instantiates a new EnvironmentRevision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentRevisionWithDefaults

`func NewEnvironmentRevisionWithDefaults() *EnvironmentRevision`

NewEnvironmentRevisionWithDefaults instantiates a new EnvironmentRevision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentRevision) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentRevision) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentRevision) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentRevision) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFrom

`func (o *EnvironmentRevision) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *EnvironmentRevision) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *EnvironmentRevision) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *EnvironmentRevision) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *EnvironmentRevision) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EnvironmentRevision) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EnvironmentRevision) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *EnvironmentRevision) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


