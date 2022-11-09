# ReportSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the report | [optional] 
**Name** | Pointer to **string** | The name of the report | [optional] 
**State** | Pointer to **string** | The state of the report | [optional] 
**CreatedAt** | Pointer to **string** | The date and time the report was generated | [optional] 
**Type** | Pointer to **string** | The type of report | [optional] 
**Source** | Pointer to [**ReportSummarySource**](ReportSummarySource.md) |  | [optional] 

## Methods

### NewReportSummary

`func NewReportSummary() *ReportSummary`

NewReportSummary instantiates a new ReportSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportSummaryWithDefaults

`func NewReportSummaryWithDefaults() *ReportSummary`

NewReportSummaryWithDefaults instantiates a new ReportSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReportSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ReportSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *ReportSummary) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ReportSummary) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ReportSummary) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ReportSummary) HasState() bool`

HasState returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReportSummary) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReportSummary) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReportSummary) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReportSummary) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetType

`func (o *ReportSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ReportSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ReportSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ReportSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSource

`func (o *ReportSummary) GetSource() ReportSummarySource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ReportSummary) GetSourceOk() (*ReportSummarySource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ReportSummary) SetSource(v ReportSummarySource)`

SetSource sets Source field to given value.

### HasSource

`func (o *ReportSummary) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


