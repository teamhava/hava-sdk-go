# EnvironmentFacet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the facet | [optional] 
**Resources** | Pointer to [**[]Resource**](Resource.md) | All the resources that are included in the environment. Resources can return different attributes based on their type. | [optional] 

## Methods

### NewEnvironmentFacet

`func NewEnvironmentFacet() *EnvironmentFacet`

NewEnvironmentFacet instantiates a new EnvironmentFacet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentFacetWithDefaults

`func NewEnvironmentFacetWithDefaults() *EnvironmentFacet`

NewEnvironmentFacetWithDefaults instantiates a new EnvironmentFacet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentFacet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentFacet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentFacet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentFacet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResources

`func (o *EnvironmentFacet) GetResources() []Resource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *EnvironmentFacet) GetResourcesOk() (*[]Resource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *EnvironmentFacet) SetResources(v []Resource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *EnvironmentFacet) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


