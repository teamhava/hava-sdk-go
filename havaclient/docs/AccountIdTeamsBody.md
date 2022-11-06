# AccountIdTeamsBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the team to create | [optional] 
**Roles** | Pointer to **[]string** | An array of roles to assign the users in this team | [optional] 

## Methods

### NewAccountIdTeamsBody

`func NewAccountIdTeamsBody() *AccountIdTeamsBody`

NewAccountIdTeamsBody instantiates a new AccountIdTeamsBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountIdTeamsBodyWithDefaults

`func NewAccountIdTeamsBodyWithDefaults() *AccountIdTeamsBody`

NewAccountIdTeamsBodyWithDefaults instantiates a new AccountIdTeamsBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AccountIdTeamsBody) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AccountIdTeamsBody) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AccountIdTeamsBody) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AccountIdTeamsBody) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRoles

`func (o *AccountIdTeamsBody) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AccountIdTeamsBody) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AccountIdTeamsBody) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *AccountIdTeamsBody) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


