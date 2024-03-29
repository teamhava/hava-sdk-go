/*
Hava

Hava API

API version: 1.1.3
Contact: support@hava.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package havaclient

import (
	"encoding/json"
)

// EnvironmentSummary struct for EnvironmentSummary
type EnvironmentSummary struct {
	// The unique ID of the environment
	Id *string `json:"id,omitempty"`
	// The name of the environment
	Name *string `json:"name,omitempty"`
	// The state of the environment
	State *string `json:"state,omitempty"`
	// The query string used to create the environment
	Query *string `json:"query,omitempty"`
	// Identifies whether the environment was imported or custom made using search
	EnvironmentType *string `json:"environment_type,omitempty"`
	// The different views for the environment
	Views []EnvironmentSummaryViewsInner `json:"views,omitempty"`
	// The sources that contain the resources this environment covers.
	Sources []EnvironmentSummarySourcesInner `json:"sources,omitempty"`
}

// NewEnvironmentSummary instantiates a new EnvironmentSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentSummary() *EnvironmentSummary {
	this := EnvironmentSummary{}
	return &this
}

// NewEnvironmentSummaryWithDefaults instantiates a new EnvironmentSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentSummaryWithDefaults() *EnvironmentSummary {
	this := EnvironmentSummary{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentSummary) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSummary) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentSummary) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentSummary) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentSummary) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSummary) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentSummary) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentSummary) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *EnvironmentSummary) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSummary) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *EnvironmentSummary) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *EnvironmentSummary) SetState(v string) {
	o.State = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *EnvironmentSummary) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSummary) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *EnvironmentSummary) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *EnvironmentSummary) SetQuery(v string) {
	o.Query = &v
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *EnvironmentSummary) GetEnvironmentType() string {
	if o == nil || isNil(o.EnvironmentType) {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSummary) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || isNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// HasEnvironmentType returns a boolean if a field has been set.
func (o *EnvironmentSummary) HasEnvironmentType() bool {
	if o != nil && !isNil(o.EnvironmentType) {
		return true
	}

	return false
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *EnvironmentSummary) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *EnvironmentSummary) GetViews() []EnvironmentSummaryViewsInner {
	if o == nil || isNil(o.Views) {
		var ret []EnvironmentSummaryViewsInner
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSummary) GetViewsOk() ([]EnvironmentSummaryViewsInner, bool) {
	if o == nil || isNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *EnvironmentSummary) HasViews() bool {
	if o != nil && !isNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given []EnvironmentSummaryViewsInner and assigns it to the Views field.
func (o *EnvironmentSummary) SetViews(v []EnvironmentSummaryViewsInner) {
	o.Views = v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *EnvironmentSummary) GetSources() []EnvironmentSummarySourcesInner {
	if o == nil || isNil(o.Sources) {
		var ret []EnvironmentSummarySourcesInner
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentSummary) GetSourcesOk() ([]EnvironmentSummarySourcesInner, bool) {
	if o == nil || isNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *EnvironmentSummary) HasSources() bool {
	if o != nil && !isNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []EnvironmentSummarySourcesInner and assigns it to the Sources field.
func (o *EnvironmentSummary) SetSources(v []EnvironmentSummarySourcesInner) {
	o.Sources = v
}

func (o EnvironmentSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	if !isNil(o.EnvironmentType) {
		toSerialize["environment_type"] = o.EnvironmentType
	}
	if !isNil(o.Views) {
		toSerialize["views"] = o.Views
	}
	if !isNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentSummary struct {
	value *EnvironmentSummary
	isSet bool
}

func (v NullableEnvironmentSummary) Get() *EnvironmentSummary {
	return v.value
}

func (v *NullableEnvironmentSummary) Set(val *EnvironmentSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentSummary(val *EnvironmentSummary) *NullableEnvironmentSummary {
	return &NullableEnvironmentSummary{value: val, isSet: true}
}

func (v NullableEnvironmentSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
