# SourcesBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for this source | [optional] 
**Type** | Pointer to **string** | Must be set to Azure::Credentials | [optional] 
**AccessKey** | Pointer to **string** | The access key for your AWS account | [optional] 
**SecretKey** | Pointer to **string** | The Client Secret for your Service Principle | [optional] 
**ExternalId** | Pointer to **string** | The external ID used to uniquely identify access on the AWS side. This value must be set to an MD5 of your account ID. You can find this value in the &#39;Add Environment&#39; prompt in Hava. | [optional] 
**RoleArn** | Pointer to **string** | The ARN of the role Hava is to assume in your account to import resources | [optional] 
**SubscriptionId** | Pointer to **string** | The ID of the Azure subscription to import from | [optional] 
**TenantId** | Pointer to **string** | The GUID representing the Active Directory Tenant | [optional] 
**ClientId** | Pointer to **string** | The Client ID for your Service Principle | [optional] 

## Methods

### NewSourcesBody

`func NewSourcesBody() *SourcesBody`

NewSourcesBody instantiates a new SourcesBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesBodyWithDefaults

`func NewSourcesBodyWithDefaults() *SourcesBody`

NewSourcesBodyWithDefaults instantiates a new SourcesBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourcesBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SourcesBody) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourcesBody) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourcesBody) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourcesBody) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessKey

`func (o *SourcesBody) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *SourcesBody) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *SourcesBody) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *SourcesBody) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *SourcesBody) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *SourcesBody) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *SourcesBody) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *SourcesBody) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetExternalId

`func (o *SourcesBody) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SourcesBody) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SourcesBody) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SourcesBody) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRoleArn

`func (o *SourcesBody) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *SourcesBody) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *SourcesBody) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *SourcesBody) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SourcesBody) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SourcesBody) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SourcesBody) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SourcesBody) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *SourcesBody) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SourcesBody) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SourcesBody) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SourcesBody) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *SourcesBody) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SourcesBody) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SourcesBody) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SourcesBody) HasClientId() bool`

HasClientId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


