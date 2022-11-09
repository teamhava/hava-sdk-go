# SourcesGCPServiceAccountCredentials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** | Must be set to GCP::ServiceAccountCredentials | [optional] 
**EncodedFile** | Pointer to **string** | Base64 encoded credentials file | [optional] 

## Methods

### NewSourcesGCPServiceAccountCredentials

`func NewSourcesGCPServiceAccountCredentials() *SourcesGCPServiceAccountCredentials`

NewSourcesGCPServiceAccountCredentials instantiates a new SourcesGCPServiceAccountCredentials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourcesGCPServiceAccountCredentialsWithDefaults

`func NewSourcesGCPServiceAccountCredentialsWithDefaults() *SourcesGCPServiceAccountCredentials`

NewSourcesGCPServiceAccountCredentialsWithDefaults instantiates a new SourcesGCPServiceAccountCredentials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SourcesGCPServiceAccountCredentials) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SourcesGCPServiceAccountCredentials) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SourcesGCPServiceAccountCredentials) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SourcesGCPServiceAccountCredentials) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *SourcesGCPServiceAccountCredentials) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SourcesGCPServiceAccountCredentials) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SourcesGCPServiceAccountCredentials) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SourcesGCPServiceAccountCredentials) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEncodedFile

`func (o *SourcesGCPServiceAccountCredentials) GetEncodedFile() string`

GetEncodedFile returns the EncodedFile field if non-nil, zero value otherwise.

### GetEncodedFileOk

`func (o *SourcesGCPServiceAccountCredentials) GetEncodedFileOk() (*string, bool)`

GetEncodedFileOk returns a tuple with the EncodedFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodedFile

`func (o *SourcesGCPServiceAccountCredentials) SetEncodedFile(v string)`

SetEncodedFile sets EncodedFile field to given value.

### HasEncodedFile

`func (o *SourcesGCPServiceAccountCredentials) HasEncodedFile() bool`

HasEncodedFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


