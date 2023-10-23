# EnvironmentsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The updated name for the environment | [optional] 
**Query** | **string** | The updated query string for the environment | 

## Methods

### NewEnvironmentsUpdateRequest

`func NewEnvironmentsUpdateRequest(query string, ) *EnvironmentsUpdateRequest`

NewEnvironmentsUpdateRequest instantiates a new EnvironmentsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsUpdateRequestWithDefaults

`func NewEnvironmentsUpdateRequestWithDefaults() *EnvironmentsUpdateRequest`

NewEnvironmentsUpdateRequestWithDefaults instantiates a new EnvironmentsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentsUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentsUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentsUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentsUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *EnvironmentsUpdateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EnvironmentsUpdateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EnvironmentsUpdateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


