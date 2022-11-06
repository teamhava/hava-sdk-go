# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the resource | [optional] 
**Name** | Pointer to **string** | The name of the resource | [optional] 
**Type** | Pointer to **string** | The type of the resource | [optional] 
**ProviderId** | Pointer to **string** | The ID assigned to this resource by the provider | [optional] 
**Price** | Pointer to **int32** | The estimated price per month assigned by Hava | [optional] 
**Region** | Pointer to **string** | The region or location in which this resource is located | [optional] 
**Data** | Pointer to **map[string]interface{}** | Resource specific data | [optional] 
**Tags** | Pointer to [**[]ResourceTags**](ResourceTags.md) | An array tags associated to this resource | [optional] 
**Connections** | Pointer to [**[]ResourceConnections**](ResourceConnections.md) | A list of connections to other resources | [optional] 

## Methods

### NewResource

`func NewResource() *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Resource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Resource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Resource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Resource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Resource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Resource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Resource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Resource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Resource) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Resource) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Resource) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Resource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProviderId

`func (o *Resource) GetProviderId() string`

GetProviderId returns the ProviderId field if non-nil, zero value otherwise.

### GetProviderIdOk

`func (o *Resource) GetProviderIdOk() (*string, bool)`

GetProviderIdOk returns a tuple with the ProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderId

`func (o *Resource) SetProviderId(v string)`

SetProviderId sets ProviderId field to given value.

### HasProviderId

`func (o *Resource) HasProviderId() bool`

HasProviderId returns a boolean if a field has been set.

### GetPrice

`func (o *Resource) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Resource) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Resource) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Resource) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRegion

`func (o *Resource) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Resource) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Resource) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Resource) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetData

`func (o *Resource) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Resource) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Resource) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *Resource) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTags

`func (o *Resource) GetTags() []ResourceTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Resource) GetTagsOk() (*[]ResourceTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Resource) SetTags(v []ResourceTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Resource) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetConnections

`func (o *Resource) GetConnections() []ResourceConnections`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *Resource) GetConnectionsOk() (*[]ResourceConnections, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *Resource) SetConnections(v []ResourceConnections)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *Resource) HasConnections() bool`

HasConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


