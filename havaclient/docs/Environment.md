# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the environment | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** | Environment State | [optional] 
**Query** | Pointer to **string** | The query string used to create the environment | [optional] 
**EnvironmentType** | Pointer to **string** | Identifies whether the environment was imported or custom made using search | [optional] 
**Views** | Pointer to [**[]View**](View.md) |  | [optional] 
**Facet** | Pointer to [**EnvironmentFacet**](EnvironmentFacet.md) |  | [optional] 
**CurrentRevision** | Pointer to [**EnvironmentCurrentRevision**](EnvironmentCurrentRevision.md) |  | [optional] 
**Sources** | Pointer to [**[]EnvironmentSummarySourcesInner**](EnvironmentSummarySourcesInner.md) | The sources that contain the resources this environment covers. | [optional] 
**LatestRevisions** | Pointer to [**EnvironmentLatestRevisions**](EnvironmentLatestRevisions.md) |  | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment() *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Environment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Environment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Environment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Environment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *Environment) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Environment) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Environment) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Environment) HasState() bool`

HasState returns a boolean if a field has been set.

### GetQuery

`func (o *Environment) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Environment) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Environment) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Environment) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetEnvironmentType

`func (o *Environment) GetEnvironmentType() string`

GetEnvironmentType returns the EnvironmentType field if non-nil, zero value otherwise.

### GetEnvironmentTypeOk

`func (o *Environment) GetEnvironmentTypeOk() (*string, bool)`

GetEnvironmentTypeOk returns a tuple with the EnvironmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentType

`func (o *Environment) SetEnvironmentType(v string)`

SetEnvironmentType sets EnvironmentType field to given value.

### HasEnvironmentType

`func (o *Environment) HasEnvironmentType() bool`

HasEnvironmentType returns a boolean if a field has been set.

### GetViews

`func (o *Environment) GetViews() []View`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *Environment) GetViewsOk() (*[]View, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *Environment) SetViews(v []View)`

SetViews sets Views field to given value.

### HasViews

`func (o *Environment) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetFacet

`func (o *Environment) GetFacet() EnvironmentFacet`

GetFacet returns the Facet field if non-nil, zero value otherwise.

### GetFacetOk

`func (o *Environment) GetFacetOk() (*EnvironmentFacet, bool)`

GetFacetOk returns a tuple with the Facet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacet

`func (o *Environment) SetFacet(v EnvironmentFacet)`

SetFacet sets Facet field to given value.

### HasFacet

`func (o *Environment) HasFacet() bool`

HasFacet returns a boolean if a field has been set.

### GetCurrentRevision

`func (o *Environment) GetCurrentRevision() EnvironmentCurrentRevision`

GetCurrentRevision returns the CurrentRevision field if non-nil, zero value otherwise.

### GetCurrentRevisionOk

`func (o *Environment) GetCurrentRevisionOk() (*EnvironmentCurrentRevision, bool)`

GetCurrentRevisionOk returns a tuple with the CurrentRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRevision

`func (o *Environment) SetCurrentRevision(v EnvironmentCurrentRevision)`

SetCurrentRevision sets CurrentRevision field to given value.

### HasCurrentRevision

`func (o *Environment) HasCurrentRevision() bool`

HasCurrentRevision returns a boolean if a field has been set.

### GetSources

`func (o *Environment) GetSources() []EnvironmentSummarySourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *Environment) GetSourcesOk() (*[]EnvironmentSummarySourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *Environment) SetSources(v []EnvironmentSummarySourcesInner)`

SetSources sets Sources field to given value.

### HasSources

`func (o *Environment) HasSources() bool`

HasSources returns a boolean if a field has been set.

### GetLatestRevisions

`func (o *Environment) GetLatestRevisions() EnvironmentLatestRevisions`

GetLatestRevisions returns the LatestRevisions field if non-nil, zero value otherwise.

### GetLatestRevisionsOk

`func (o *Environment) GetLatestRevisionsOk() (*EnvironmentLatestRevisions, bool)`

GetLatestRevisionsOk returns a tuple with the LatestRevisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRevisions

`func (o *Environment) SetLatestRevisions(v EnvironmentLatestRevisions)`

SetLatestRevisions sets LatestRevisions field to given value.

### HasLatestRevisions

`func (o *Environment) HasLatestRevisions() bool`

HasLatestRevisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


