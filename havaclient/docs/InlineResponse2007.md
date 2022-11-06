# InlineResponse2007

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | Pointer to **string** | The token to pass through to get the next page of results, if there are any. | [optional] 
**TotalSize** | Pointer to **int32** | The total amount of results to be paginated | [optional] 
**Results** | Pointer to [**[]ReportSummary**](ReportSummary.md) | The results for the current page | [optional] 

## Methods

### NewInlineResponse2007

`func NewInlineResponse2007() *InlineResponse2007`

NewInlineResponse2007 instantiates a new InlineResponse2007 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineResponse2007WithDefaults

`func NewInlineResponse2007WithDefaults() *InlineResponse2007`

NewInlineResponse2007WithDefaults instantiates a new InlineResponse2007 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *InlineResponse2007) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *InlineResponse2007) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *InlineResponse2007) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *InlineResponse2007) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.

### GetTotalSize

`func (o *InlineResponse2007) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *InlineResponse2007) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *InlineResponse2007) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *InlineResponse2007) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetResults

`func (o *InlineResponse2007) GetResults() []ReportSummary`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *InlineResponse2007) GetResultsOk() (*[]ReportSummary, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *InlineResponse2007) SetResults(v []ReportSummary)`

SetResults sets Results field to given value.

### HasResults

`func (o *InlineResponse2007) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


