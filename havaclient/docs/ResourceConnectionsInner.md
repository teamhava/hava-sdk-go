# ResourceConnectionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The unique ID of the connection | [optional] 
**ResourceId** | Pointer to **string** | The unique ID of the resource this connection is from | [optional] 
**RemoteResourceId** | Pointer to **string** | The unique ID of the resource this connection is to | [optional] 
**RemoteResourceType** | Pointer to **string** | The type of the resource this connection is to | [optional] 
**Data** | Pointer to **map[string]interface{}** | Additional data related to this connection | [optional] 

## Methods

### NewResourceConnectionsInner

`func NewResourceConnectionsInner() *ResourceConnectionsInner`

NewResourceConnectionsInner instantiates a new ResourceConnectionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConnectionsInnerWithDefaults

`func NewResourceConnectionsInnerWithDefaults() *ResourceConnectionsInner`

NewResourceConnectionsInnerWithDefaults instantiates a new ResourceConnectionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceConnectionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceConnectionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceConnectionsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceConnectionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResourceId

`func (o *ResourceConnectionsInner) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceConnectionsInner) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceConnectionsInner) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *ResourceConnectionsInner) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetRemoteResourceId

`func (o *ResourceConnectionsInner) GetRemoteResourceId() string`

GetRemoteResourceId returns the RemoteResourceId field if non-nil, zero value otherwise.

### GetRemoteResourceIdOk

`func (o *ResourceConnectionsInner) GetRemoteResourceIdOk() (*string, bool)`

GetRemoteResourceIdOk returns a tuple with the RemoteResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteResourceId

`func (o *ResourceConnectionsInner) SetRemoteResourceId(v string)`

SetRemoteResourceId sets RemoteResourceId field to given value.

### HasRemoteResourceId

`func (o *ResourceConnectionsInner) HasRemoteResourceId() bool`

HasRemoteResourceId returns a boolean if a field has been set.

### GetRemoteResourceType

`func (o *ResourceConnectionsInner) GetRemoteResourceType() string`

GetRemoteResourceType returns the RemoteResourceType field if non-nil, zero value otherwise.

### GetRemoteResourceTypeOk

`func (o *ResourceConnectionsInner) GetRemoteResourceTypeOk() (*string, bool)`

GetRemoteResourceTypeOk returns a tuple with the RemoteResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteResourceType

`func (o *ResourceConnectionsInner) SetRemoteResourceType(v string)`

SetRemoteResourceType sets RemoteResourceType field to given value.

### HasRemoteResourceType

`func (o *ResourceConnectionsInner) HasRemoteResourceType() bool`

HasRemoteResourceType returns a boolean if a field has been set.

### GetData

`func (o *ResourceConnectionsInner) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceConnectionsInner) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceConnectionsInner) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ResourceConnectionsInner) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


