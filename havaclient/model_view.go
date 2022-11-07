/*
Hava

Hava API

API version: 1.1.2
Contact: support@hava.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package havaclient

import (
	"encoding/json"
)

// View struct for View
type View struct {
	// The unique ID of the view
	Id *string `json:"id,omitempty"`
	// The type of view
	Type *string `json:"type,omitempty"`
	// The name of the exported thumbnail for this view
	ImageName *string `json:"image_name,omitempty"`
	// The location of the exported thumbnail for this view
	ImageUrl *string `json:"image_url,omitempty"`
	// The creation timestamp for the exported thumbnail
	ExportTimestamp *int32 `json:"export_timestamp,omitempty"`
	// The unique ID of the current version of this view
	RevisionId *string `json:"revision_id,omitempty"`
	// The regions or locations displayed in this view
	Regions []string `json:"regions,omitempty"`
	// Whether the view is considered empty and has no valuable resources to display
	Empty *bool `json:"empty,omitempty"`
	// A list of resource ID's to be displayed. Only returned for list view.
	Resources []ViewResourcesInner `json:"resources,omitempty"`
}

// NewView instantiates a new View object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewView() *View {
	this := View{}
	return &this
}

// NewViewWithDefaults instantiates a new View object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewWithDefaults() *View {
	this := View{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *View) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *View) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *View) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *View) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *View) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *View) SetType(v string) {
	o.Type = &v
}

// GetImageName returns the ImageName field value if set, zero value otherwise.
func (o *View) GetImageName() string {
	if o == nil || isNil(o.ImageName) {
		var ret string
		return ret
	}
	return *o.ImageName
}

// GetImageNameOk returns a tuple with the ImageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetImageNameOk() (*string, bool) {
	if o == nil || isNil(o.ImageName) {
		return nil, false
	}
	return o.ImageName, true
}

// HasImageName returns a boolean if a field has been set.
func (o *View) HasImageName() bool {
	if o != nil && !isNil(o.ImageName) {
		return true
	}

	return false
}

// SetImageName gets a reference to the given string and assigns it to the ImageName field.
func (o *View) SetImageName(v string) {
	o.ImageName = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *View) GetImageUrl() string {
	if o == nil || isNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetImageUrlOk() (*string, bool) {
	if o == nil || isNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *View) HasImageUrl() bool {
	if o != nil && !isNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *View) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetExportTimestamp returns the ExportTimestamp field value if set, zero value otherwise.
func (o *View) GetExportTimestamp() int32 {
	if o == nil || isNil(o.ExportTimestamp) {
		var ret int32
		return ret
	}
	return *o.ExportTimestamp
}

// GetExportTimestampOk returns a tuple with the ExportTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetExportTimestampOk() (*int32, bool) {
	if o == nil || isNil(o.ExportTimestamp) {
		return nil, false
	}
	return o.ExportTimestamp, true
}

// HasExportTimestamp returns a boolean if a field has been set.
func (o *View) HasExportTimestamp() bool {
	if o != nil && !isNil(o.ExportTimestamp) {
		return true
	}

	return false
}

// SetExportTimestamp gets a reference to the given int32 and assigns it to the ExportTimestamp field.
func (o *View) SetExportTimestamp(v int32) {
	o.ExportTimestamp = &v
}

// GetRevisionId returns the RevisionId field value if set, zero value otherwise.
func (o *View) GetRevisionId() string {
	if o == nil || isNil(o.RevisionId) {
		var ret string
		return ret
	}
	return *o.RevisionId
}

// GetRevisionIdOk returns a tuple with the RevisionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetRevisionIdOk() (*string, bool) {
	if o == nil || isNil(o.RevisionId) {
		return nil, false
	}
	return o.RevisionId, true
}

// HasRevisionId returns a boolean if a field has been set.
func (o *View) HasRevisionId() bool {
	if o != nil && !isNil(o.RevisionId) {
		return true
	}

	return false
}

// SetRevisionId gets a reference to the given string and assigns it to the RevisionId field.
func (o *View) SetRevisionId(v string) {
	o.RevisionId = &v
}

// GetRegions returns the Regions field value if set, zero value otherwise.
func (o *View) GetRegions() []string {
	if o == nil || isNil(o.Regions) {
		var ret []string
		return ret
	}
	return o.Regions
}

// GetRegionsOk returns a tuple with the Regions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetRegionsOk() ([]string, bool) {
	if o == nil || isNil(o.Regions) {
		return nil, false
	}
	return o.Regions, true
}

// HasRegions returns a boolean if a field has been set.
func (o *View) HasRegions() bool {
	if o != nil && !isNil(o.Regions) {
		return true
	}

	return false
}

// SetRegions gets a reference to the given []string and assigns it to the Regions field.
func (o *View) SetRegions(v []string) {
	o.Regions = v
}

// GetEmpty returns the Empty field value if set, zero value otherwise.
func (o *View) GetEmpty() bool {
	if o == nil || isNil(o.Empty) {
		var ret bool
		return ret
	}
	return *o.Empty
}

// GetEmptyOk returns a tuple with the Empty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetEmptyOk() (*bool, bool) {
	if o == nil || isNil(o.Empty) {
		return nil, false
	}
	return o.Empty, true
}

// HasEmpty returns a boolean if a field has been set.
func (o *View) HasEmpty() bool {
	if o != nil && !isNil(o.Empty) {
		return true
	}

	return false
}

// SetEmpty gets a reference to the given bool and assigns it to the Empty field.
func (o *View) SetEmpty(v bool) {
	o.Empty = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *View) GetResources() []ViewResourcesInner {
	if o == nil || isNil(o.Resources) {
		var ret []ViewResourcesInner
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *View) GetResourcesOk() ([]ViewResourcesInner, bool) {
	if o == nil || isNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *View) HasResources() bool {
	if o != nil && !isNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ViewResourcesInner and assigns it to the Resources field.
func (o *View) SetResources(v []ViewResourcesInner) {
	o.Resources = v
}

func (o View) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.ImageName) {
		toSerialize["image_name"] = o.ImageName
	}
	if !isNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !isNil(o.ExportTimestamp) {
		toSerialize["export_timestamp"] = o.ExportTimestamp
	}
	if !isNil(o.RevisionId) {
		toSerialize["revision_id"] = o.RevisionId
	}
	if !isNil(o.Regions) {
		toSerialize["regions"] = o.Regions
	}
	if !isNil(o.Empty) {
		toSerialize["empty"] = o.Empty
	}
	if !isNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return json.Marshal(toSerialize)
}

type NullableView struct {
	value *View
	isSet bool
}

func (v NullableView) Get() *View {
	return v.value
}

func (v *NullableView) Set(val *View) {
	v.value = val
	v.isSet = true
}

func (v NullableView) IsSet() bool {
	return v.isSet
}

func (v *NullableView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableView(val *View) *NullableView {
	return &NullableView{value: val, isSet: true}
}

func (v NullableView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


