# ViewIdExportBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportFormat** | Pointer to **string** | The format to export the view to | [optional] 
**Connections** | Pointer to **bool** | Display connections on the exported diagram | [optional] 
**Isometric** | Pointer to **bool** | Display diagram in isometric view | [optional] 
**Labels** | Pointer to **bool** | Display labels such as names on the diagram | [optional] 
**RevisionId** | Pointer to **string** | The ID of the view revision to export (optional) | [optional] 

## Methods

### NewViewIdExportBody

`func NewViewIdExportBody() *ViewIdExportBody`

NewViewIdExportBody instantiates a new ViewIdExportBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewIdExportBodyWithDefaults

`func NewViewIdExportBodyWithDefaults() *ViewIdExportBody`

NewViewIdExportBodyWithDefaults instantiates a new ViewIdExportBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportFormat

`func (o *ViewIdExportBody) GetExportFormat() string`

GetExportFormat returns the ExportFormat field if non-nil, zero value otherwise.

### GetExportFormatOk

`func (o *ViewIdExportBody) GetExportFormatOk() (*string, bool)`

GetExportFormatOk returns a tuple with the ExportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFormat

`func (o *ViewIdExportBody) SetExportFormat(v string)`

SetExportFormat sets ExportFormat field to given value.

### HasExportFormat

`func (o *ViewIdExportBody) HasExportFormat() bool`

HasExportFormat returns a boolean if a field has been set.

### GetConnections

`func (o *ViewIdExportBody) GetConnections() bool`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ViewIdExportBody) GetConnectionsOk() (*bool, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ViewIdExportBody) SetConnections(v bool)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *ViewIdExportBody) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetIsometric

`func (o *ViewIdExportBody) GetIsometric() bool`

GetIsometric returns the Isometric field if non-nil, zero value otherwise.

### GetIsometricOk

`func (o *ViewIdExportBody) GetIsometricOk() (*bool, bool)`

GetIsometricOk returns a tuple with the Isometric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsometric

`func (o *ViewIdExportBody) SetIsometric(v bool)`

SetIsometric sets Isometric field to given value.

### HasIsometric

`func (o *ViewIdExportBody) HasIsometric() bool`

HasIsometric returns a boolean if a field has been set.

### GetLabels

`func (o *ViewIdExportBody) GetLabels() bool`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ViewIdExportBody) GetLabelsOk() (*bool, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ViewIdExportBody) SetLabels(v bool)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ViewIdExportBody) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetRevisionId

`func (o *ViewIdExportBody) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *ViewIdExportBody) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *ViewIdExportBody) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *ViewIdExportBody) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


