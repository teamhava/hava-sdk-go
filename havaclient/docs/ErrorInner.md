# ErrorInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The error code relevant to this error | [optional] 
**Title** | Pointer to **string** | A short desription for the error type | [optional] 
**Detail** | Pointer to **string** | A detailed message of the error encountered | [optional] 
**Path** | Pointer to **string** | The path to the value that caused the error | [optional] 

## Methods

### NewErrorInner

`func NewErrorInner() *ErrorInner`

NewErrorInner instantiates a new ErrorInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorInnerWithDefaults

`func NewErrorInnerWithDefaults() *ErrorInner`

NewErrorInnerWithDefaults instantiates a new ErrorInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTitle

`func (o *ErrorInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ErrorInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ErrorInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ErrorInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDetail

`func (o *ErrorInner) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorInner) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorInner) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorInner) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetPath

`func (o *ErrorInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ErrorInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ErrorInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ErrorInner) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


