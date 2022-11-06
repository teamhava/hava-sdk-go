# SourcesAzureCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for this source | [optional] 
**Type** | Pointer to **string** | Must be set to Azure::Credentials | [optional] 
**SubscriptionId** | Pointer to **string** | The ID of the Azure subscription to import from | [optional] 
**TenantId** | Pointer to **string** | The GUID representing the Active Directory Tenant | [optional] 
**ClientId** | Pointer to **string** | The Client ID for your Service Principle | [optional] 
**SecretKey** | Pointer to **string** | The Client Secret for your Service Principle | [optional] 

## Methods

### NewSourcesAzureCredentials

`func NewSourcesAzureCredentials() *SourcesAzureCredentials`

NewSourcesAzureCredentials instantiates a new SourcesAzureCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesAzureCredentialsWithDefaults

`func NewSourcesAzureCredentialsWithDefaults() *SourcesAzureCredentials`

NewSourcesAzureCredentialsWithDefaults instantiates a new SourcesAzureCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesAzureCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesAzureCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesAzureCredentials) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourcesAzureCredentials) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SourcesAzureCredentials) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourcesAzureCredentials) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourcesAzureCredentials) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourcesAzureCredentials) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SourcesAzureCredentials) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SourcesAzureCredentials) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SourcesAzureCredentials) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SourcesAzureCredentials) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *SourcesAzureCredentials) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SourcesAzureCredentials) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SourcesAzureCredentials) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SourcesAzureCredentials) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *SourcesAzureCredentials) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SourcesAzureCredentials) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SourcesAzureCredentials) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SourcesAzureCredentials) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecretKey

`func (o *SourcesAzureCredentials) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *SourcesAzureCredentials) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *SourcesAzureCredentials) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *SourcesAzureCredentials) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


