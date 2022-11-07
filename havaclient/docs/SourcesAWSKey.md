# SourcesAWSKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name for this source | [optional] 
**Type** | Pointer to **string** | Must be set to AWS::Keys | [optional] 
**AccessKey** | Pointer to **string** | The access key for your AWS account | [optional] 
**SecretKey** | Pointer to **string** | The secret key for your AWS account | [optional] 
**ProjectId** | Pointer to **string** | The ID of the project the source will be added to. If not set the source will get added to the Default project | [optional] 

## Methods

### NewSourcesAWSKey

`func NewSourcesAWSKey() *SourcesAWSKey`

NewSourcesAWSKey instantiates a new SourcesAWSKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesAWSKeyWithDefaults

`func NewSourcesAWSKeyWithDefaults() *SourcesAWSKey`

NewSourcesAWSKeyWithDefaults instantiates a new SourcesAWSKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesAWSKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesAWSKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesAWSKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourcesAWSKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SourcesAWSKey) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourcesAWSKey) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourcesAWSKey) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourcesAWSKey) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccessKey

`func (o *SourcesAWSKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *SourcesAWSKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *SourcesAWSKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *SourcesAWSKey) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *SourcesAWSKey) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *SourcesAWSKey) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *SourcesAWSKey) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *SourcesAWSKey) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetProjectId

`func (o *SourcesAWSKey) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SourcesAWSKey) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SourcesAWSKey) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SourcesAWSKey) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


