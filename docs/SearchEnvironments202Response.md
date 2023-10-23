# SearchEnvironments202Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The token to pass through to get the next page of results, if there are any. | [optional] 
**TotalSize** | Pointer to **int32** | The total amount of results to be paginated | [optional] 
**Results** | Pointer to [**[]EnvironmentSearchResult**](EnvironmentSearchResult.md) | The results for the current page | [optional] 

## Methods

### NewSearchEnvironments202Response

`func NewSearchEnvironments202Response() *SearchEnvironments202Response`

NewSearchEnvironments202Response instantiates a new SearchEnvironments202Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchEnvironments202ResponseWithDefaults

`func NewSearchEnvironments202ResponseWithDefaults() *SearchEnvironments202Response`

NewSearchEnvironments202ResponseWithDefaults instantiates a new SearchEnvironments202Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *SearchEnvironments202Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *SearchEnvironments202Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *SearchEnvironments202Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *SearchEnvironments202Response) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTotalSize

`func (o *SearchEnvironments202Response) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *SearchEnvironments202Response) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *SearchEnvironments202Response) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *SearchEnvironments202Response) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetResults

`func (o *SearchEnvironments202Response) GetResults() []EnvironmentSearchResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SearchEnvironments202Response) GetResultsOk() (*[]EnvironmentSearchResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SearchEnvironments202Response) SetResults(v []EnvironmentSearchResult)`

SetResults sets Results field to given value.

### HasResults

`func (o *SearchEnvironments202Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


