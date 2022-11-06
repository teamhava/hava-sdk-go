# EnvironmentsEnvironmentIdBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The updated name for the environment | [optional] 
**Query** | **string** | The updated query string for the environment | 
**Complete** | Pointer to **bool** | If true the environment will be set to active and display on your environment list | [optional] 

## Methods

### NewEnvironmentsEnvironmentIdBody

`func NewEnvironmentsEnvironmentIdBody(query string, ) *EnvironmentsEnvironmentIdBody`

NewEnvironmentsEnvironmentIdBody instantiates a new EnvironmentsEnvironmentIdBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentsEnvironmentIdBodyWithDefaults

`func NewEnvironmentsEnvironmentIdBodyWithDefaults() *EnvironmentsEnvironmentIdBody`

NewEnvironmentsEnvironmentIdBodyWithDefaults instantiates a new EnvironmentsEnvironmentIdBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentsEnvironmentIdBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentsEnvironmentIdBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentsEnvironmentIdBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentsEnvironmentIdBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQuery

`func (o *EnvironmentsEnvironmentIdBody) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EnvironmentsEnvironmentIdBody) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EnvironmentsEnvironmentIdBody) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetComplete

`func (o *EnvironmentsEnvironmentIdBody) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *EnvironmentsEnvironmentIdBody) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *EnvironmentsEnvironmentIdBody) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *EnvironmentsEnvironmentIdBody) HasComplete() bool`

HasComplete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


