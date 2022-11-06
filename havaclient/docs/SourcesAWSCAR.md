# SourcesAWSCAR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for this source | [optional] 
**Type** | Pointer to **string** | Must be set to AWS::CrossAccountRole | [optional] 
**ExternalId** | Pointer to **string** | The external ID used to uniquely identify access on the AWS side. This value must be set to an MD5 of your account ID. You can find this value in the &#39;Add Environment&#39; prompt in Hava. | [optional] 
**RoleArn** | Pointer to **string** | The ARN of the role Hava is to assume in your account to import resources | [optional] 

## Methods

### NewSourcesAWSCAR

`func NewSourcesAWSCAR() *SourcesAWSCAR`

NewSourcesAWSCAR instantiates a new SourcesAWSCAR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesAWSCARWithDefaults

`func NewSourcesAWSCARWithDefaults() *SourcesAWSCAR`

NewSourcesAWSCARWithDefaults instantiates a new SourcesAWSCAR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesAWSCAR) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesAWSCAR) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesAWSCAR) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourcesAWSCAR) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SourcesAWSCAR) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourcesAWSCAR) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourcesAWSCAR) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourcesAWSCAR) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExternalId

`func (o *SourcesAWSCAR) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SourcesAWSCAR) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SourcesAWSCAR) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SourcesAWSCAR) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetRoleArn

`func (o *SourcesAWSCAR) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *SourcesAWSCAR) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *SourcesAWSCAR) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *SourcesAWSCAR) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


