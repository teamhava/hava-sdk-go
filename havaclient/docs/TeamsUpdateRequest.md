# TeamsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The updated name for the team | [optional] 
**Roles** | Pointer to **[]string** | The updated roles to assign the users in this team | [optional] 

## Methods

### NewTeamsUpdateRequest

`func NewTeamsUpdateRequest() *TeamsUpdateRequest`

NewTeamsUpdateRequest instantiates a new TeamsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsUpdateRequestWithDefaults

`func NewTeamsUpdateRequestWithDefaults() *TeamsUpdateRequest`

NewTeamsUpdateRequestWithDefaults instantiates a new TeamsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TeamsUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamsUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamsUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TeamsUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *TeamsUpdateRequest) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *TeamsUpdateRequest) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *TeamsUpdateRequest) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *TeamsUpdateRequest) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


