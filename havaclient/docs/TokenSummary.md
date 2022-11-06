# TokenSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the API token | [optional] 
**AccountId** | Pointer to **string** | The unique ID of the account this API token is in | [optional] 
**Name** | Pointer to **string** | The name of the API token | [optional] 
**LastUsedAt** | Pointer to **string** | The last time this API token was used. Will be empty if it has never been used. | [optional] 

## Methods

### NewTokenSummary

`func NewTokenSummary() *TokenSummary`

NewTokenSummary instantiates a new TokenSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenSummaryWithDefaults

`func NewTokenSummaryWithDefaults() *TokenSummary`

NewTokenSummaryWithDefaults instantiates a new TokenSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TokenSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *TokenSummary) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TokenSummary) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TokenSummary) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TokenSummary) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *TokenSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *TokenSummary) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *TokenSummary) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *TokenSummary) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *TokenSummary) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


