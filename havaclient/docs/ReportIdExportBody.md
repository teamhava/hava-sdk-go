# ReportIdExportBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportFormat** | **string** | The format to export the report to | 
**ExportOption** | Pointer to **string** | Toggle to show or hide the details in the findings section | [optional] 

## Methods

### NewReportIdExportBody

`func NewReportIdExportBody(exportFormat string, ) *ReportIdExportBody`

NewReportIdExportBody instantiates a new ReportIdExportBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportIdExportBodyWithDefaults

`func NewReportIdExportBodyWithDefaults() *ReportIdExportBody`

NewReportIdExportBodyWithDefaults instantiates a new ReportIdExportBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportFormat

`func (o *ReportIdExportBody) GetExportFormat() string`

GetExportFormat returns the ExportFormat field if non-nil, zero value otherwise.

### GetExportFormatOk

`func (o *ReportIdExportBody) GetExportFormatOk() (*string, bool)`

GetExportFormatOk returns a tuple with the ExportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFormat

`func (o *ReportIdExportBody) SetExportFormat(v string)`

SetExportFormat sets ExportFormat field to given value.


### GetExportOption

`func (o *ReportIdExportBody) GetExportOption() string`

GetExportOption returns the ExportOption field if non-nil, zero value otherwise.

### GetExportOptionOk

`func (o *ReportIdExportBody) GetExportOptionOk() (*string, bool)`

GetExportOptionOk returns a tuple with the ExportOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportOption

`func (o *ReportIdExportBody) SetExportOption(v string)`

SetExportOption sets ExportOption field to given value.

### HasExportOption

`func (o *ReportIdExportBody) HasExportOption() bool`

HasExportOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


