# EnvironmentSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the environment | [optional] 
**Name** | Pointer to **string** | The name of the environment | [optional] 
**State** | Pointer to **string** | The state of the environment | [optional] 
**Query** | Pointer to **string** | The query string used to create the environment | [optional] 
**EnvironmentType** | Pointer to **string** | Identifies whether the environment was imported or custom made using search | [optional] 
**Views** | Pointer to [**[]EnvironmentSummaryViews**](EnvironmentSummaryViews.md) | The different views for the environment | [optional] 
**Sources** | Pointer to [**[]EnvironmentSummarySources**](EnvironmentSummarySources.md) | The sources that contain the resources this environment covers. | [optional] 

## Methods

### NewEnvironmentSummary

`func NewEnvironmentSummary() *EnvironmentSummary`

NewEnvironmentSummary instantiates a new EnvironmentSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentSummaryWithDefaults

`func NewEnvironmentSummaryWithDefaults() *EnvironmentSummary`

NewEnvironmentSummaryWithDefaults instantiates a new EnvironmentSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentSummary) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EnvironmentSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EnvironmentSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *EnvironmentSummary) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *EnvironmentSummary) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *EnvironmentSummary) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *EnvironmentSummary) HasState() bool`

HasState returns a boolean if a field has been set.

### GetQuery

`func (o *EnvironmentSummary) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *EnvironmentSummary) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *EnvironmentSummary) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *EnvironmentSummary) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *EnvironmentSummary) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *EnvironmentSummary) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *EnvironmentSummary) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *EnvironmentSummary) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetViews

`func (o *EnvironmentSummary) GetViews() []EnvironmentSummaryViews`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *EnvironmentSummary) GetViewsOk() (*[]EnvironmentSummaryViews, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *EnvironmentSummary) SetViews(v []EnvironmentSummaryViews)`

SetViews sets Views field to given value.

### HasViews

`func (o *EnvironmentSummary) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetSources

`func (o *EnvironmentSummary) GetSources() []EnvironmentSummarySources`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *EnvironmentSummary) GetSourcesOk() (*[]EnvironmentSummarySources, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *EnvironmentSummary) SetSources(v []EnvironmentSummarySources)`

SetSources sets Sources field to given value.

### HasSources

`func (o *EnvironmentSummary) HasSources() bool`

HasSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


