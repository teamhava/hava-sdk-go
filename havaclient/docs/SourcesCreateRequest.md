# SourcesCreateRequest

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

### NewSourcesCreateRequest

`func NewSourcesCreateRequest() *SourcesCreateRequest`

NewSourcesCreateRequest instantiates a new SourcesCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesCreateRequestWithDefaults

`func NewSourcesCreateRequestWithDefaults() *SourcesCreateRequest`

NewSourcesCreateRequestWithDefaults instantiates a new SourcesCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourcesCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SourcesCreateRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourcesCreateRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourcesCreateRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourcesCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessKey

`func (o *SourcesCreateRequest) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *SourcesCreateRequest) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *SourcesCreateRequest) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *SourcesCreateRequest) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *SourcesCreateRequest) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *SourcesCreateRequest) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *SourcesCreateRequest) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *SourcesCreateRequest) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetProjectId

`func (o *SourcesCreateRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SourcesCreateRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SourcesCreateRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SourcesCreateRequest) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetExternalId

`func (o *SourcesCreateRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SourcesCreateRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SourcesCreateRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SourcesCreateRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRoleArn

`func (o *SourcesCreateRequest) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *SourcesCreateRequest) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *SourcesCreateRequest) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *SourcesCreateRequest) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetSubscriptionId

`func (o *SourcesCreateRequest) GetSubscriptionId() string`

GetSubscriptionId returns the SubscriptionId field if non-nil, zero value otherwise.

### GetSubscriptionIdOk

`func (o *SourcesCreateRequest) GetSubscriptionIdOk() (*string, bool)`

GetSubscriptionIdOk returns a tuple with the SubscriptionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionId

`func (o *SourcesCreateRequest) SetSubscriptionId(v string)`

SetSubscriptionId sets SubscriptionId field to given value.

### HasSubscriptionId

`func (o *SourcesCreateRequest) HasSubscriptionId() bool`

HasSubscriptionId returns a boolean if a field has been set.

### GetTenantId

`func (o *SourcesCreateRequest) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *SourcesCreateRequest) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *SourcesCreateRequest) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *SourcesCreateRequest) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetClientId

`func (o *SourcesCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SourcesCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SourcesCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SourcesCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetEncodedFile

`func (o *SourcesCreateRequest) GetEncodedFile() string`

GetEncodedFile returns the EncodedFile field if non-nil, zero value otherwise.

### GetEncodedFileOk

`func (o *SourcesCreateRequest) GetEncodedFileOk() (*string, bool)`

GetEncodedFileOk returns a tuple with the EncodedFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedFile

`func (o *SourcesCreateRequest) SetEncodedFile(v string)`

SetEncodedFile sets EncodedFile field to given value.

### HasEncodedFile

`func (o *SourcesCreateRequest) HasEncodedFile() bool`

HasEncodedFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


