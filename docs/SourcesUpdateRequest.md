# SourcesUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Must be set to GCP::ServiceAccountCredentials | [optional] 
**AccessKey** | Pointer to **string** | The access key for your AWS account | [optional] 
**SecretKey** | Pointer to **string** | The Client Secret for your Service Principle | [optional] 
**ProjectId** | Pointer to **string** | The ID of the project the source will be added to. If not set the source will get added to the Default project | [optional] 
**ExternalId** | Pointer to **string** | The external ID used to uniquely identify access on the AWS side. This value must be set to an MD5 of your account ID. You can find this value in the &#39;Add Environment&#39; prompt in Hava. | [optional] 
**RoleArn** | Pointer to **string** | The ARN of the role Hava is to assume in your account to import resources | [optional] 
**SubscriptionId** | Pointer to **string** | The ID of the Azure subscription to import from | [optional] 
**TenantId** | Pointer to **string** | The GUID representing the Active Directory Tenant | [optional] 
**ClientId** | Pointer to **string** | The Client ID for your Service Principle | [optional] 
**EncodedFile** | Pointer to **string** | Base64 encoded credentials file | [optional] 

## Methods

### NewSourcesUpdateRequest

`func NewSourcesUpdateRequest() *SourcesUpdateRequest`

NewSourcesUpdateRequest instantiates a new SourcesUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesUpdateRequestWithDefaults

`func NewSourcesUpdateRequestWithDefaults() *SourcesUpdateRequest`

NewSourcesUpdateRequestWithDefaults instantiates a new SourcesUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourcesUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SourcesUpdateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourcesUpdateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourcesUpdateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourcesUpdateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessKey

`func (o *SourcesUpdateRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *SourcesUpdateRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *SourcesUpdateRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *SourcesUpdateRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *SourcesUpdateRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *SourcesUpdateRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *SourcesUpdateRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *SourcesUpdateRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetProjectId

`func (o *SourcesUpdateRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SourcesUpdateRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SourcesUpdateRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SourcesUpdateRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetExternalId

`func (o *SourcesUpdateRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SourcesUpdateRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SourcesUpdateRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SourcesUpdateRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRoleArn

`func (o *SourcesUpdateRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *SourcesUpdateRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *SourcesUpdateRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *SourcesUpdateRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SourcesUpdateRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SourcesUpdateRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SourcesUpdateRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SourcesUpdateRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *SourcesUpdateRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SourcesUpdateRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SourcesUpdateRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SourcesUpdateRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *SourcesUpdateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SourcesUpdateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SourcesUpdateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SourcesUpdateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetEncodedFile

`func (o *SourcesUpdateRequest) GetEncodedFile() string`

GetEncodedFile returns the EncodedFile field if non-nil, zero value otherwise.

### GetEncodedFileOk

`func (o *SourcesUpdateRequest) GetEncodedFileOk() (*string, bool)`

GetEncodedFileOk returns a tuple with the EncodedFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedFile

`func (o *SourcesUpdateRequest) SetEncodedFile(v string)`

SetEncodedFile sets EncodedFile field to given value.

### HasEncodedFile

`func (o *SourcesUpdateRequest) HasEncodedFile() bool`

HasEncodedFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


