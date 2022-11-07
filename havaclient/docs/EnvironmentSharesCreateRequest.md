# EnvironmentSharesCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of share to create | [optional] 
**Config** | Pointer to [**EnvironmentSharesCreateRequestConfig**](EnvironmentSharesCreateRequestConfig.md) |  | [optional] 

## Methods

### NewEnvironmentSharesCreateRequest

`func NewEnvironmentSharesCreateRequest() *EnvironmentSharesCreateRequest`

NewEnvironmentSharesCreateRequest instantiates a new EnvironmentSharesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSharesCreateRequestWithDefaults

`func NewEnvironmentSharesCreateRequestWithDefaults() *EnvironmentSharesCreateRequest`

NewEnvironmentSharesCreateRequestWithDefaults instantiates a new EnvironmentSharesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EnvironmentSharesCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentSharesCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentSharesCreateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EnvironmentSharesCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *EnvironmentSharesCreateRequest) GetConfig() EnvironmentSharesCreateRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *EnvironmentSharesCreateRequest) GetConfigOk() (*EnvironmentSharesCreateRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *EnvironmentSharesCreateRequest) SetConfig(v EnvironmentSharesCreateRequestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *EnvironmentSharesCreateRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


