# EnvironmentsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for your environment | [optional] 
**Query** | **string** | The query string used to create your environment | 

## Methods

### NewEnvironmentsCreateRequest

`func NewEnvironmentsCreateRequest(query string, ) *EnvironmentsCreateRequest`

NewEnvironmentsCreateRequest instantiates a new EnvironmentsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsCreateRequestWithDefaults

`func NewEnvironmentsCreateRequestWithDefaults() *EnvironmentsCreateRequest`

NewEnvironmentsCreateRequestWithDefaults instantiates a new EnvironmentsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentsCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentsCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentsCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentsCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *EnvironmentsCreateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EnvironmentsCreateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EnvironmentsCreateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


