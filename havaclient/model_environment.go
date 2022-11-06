/*
Hava

Hava API

API version: 1.1.1
Contact: support@hava.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package havaclient

import (
	"encoding/json"
)

// Environment struct for Environment
type Environment struct {
	// The unique ID of the environment
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// Environment State
	State *string `json:"state,omitempty"`
	// The query string used to create the environment
	Query *string `json:"query,omitempty"`
	// Identifies whether the environment was imported or custom made using search
	EnvironmentType *string `json:"environment_type,omitempty"`
	Views []View `json:"views,omitempty"`
	Facet *EnvironmentFacet `json:"facet,omitempty"`
	CurrentRevision *EnvironmentCurrentRevision `json:"current_revision,omitempty"`
	// The sources that contain the resources this environment covers.
	Sources []EnvironmentSummarySources `json:"sources,omitempty"`
	LatestRevisions *EnvironmentLatestRevisions `json:"latest_revisions,omitempty"`
}

// NewEnvironment instantiates a new Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironment() *Environment {
	this := Environment{}
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Environment) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Environment) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Environment) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Environment) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Environment) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Environment) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *Environment) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *Environment) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *Environment) SetState(v string) {
	o.State = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *Environment) GetQuery() string {
	if o == nil || isNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetQueryOk() (*string, bool) {
	if o == nil || isNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *Environment) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *Environment) SetQuery(v string) {
	o.Query = &v
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *Environment) GetEnvironmentType() string {
	if o == nil || isNil(o.EnvironmentType) {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || isNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// HasEnvironmentType returns a boolean if a field has been set.
func (o *Environment) HasEnvironmentType() bool {
	if o != nil && !isNil(o.EnvironmentType) {
		return true
	}

	return false
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *Environment) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetViews returns the Views field value if set, zero value otherwise.
func (o *Environment) GetViews() []View {
	if o == nil || isNil(o.Views) {
		var ret []View
		return ret
	}
	return o.Views
}

// GetViewsOk returns a tuple with the Views field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetViewsOk() ([]View, bool) {
	if o == nil || isNil(o.Views) {
		return nil, false
	}
	return o.Views, true
}

// HasViews returns a boolean if a field has been set.
func (o *Environment) HasViews() bool {
	if o != nil && !isNil(o.Views) {
		return true
	}

	return false
}

// SetViews gets a reference to the given []View and assigns it to the Views field.
func (o *Environment) SetViews(v []View) {
	o.Views = v
}

// GetFacet returns the Facet field value if set, zero value otherwise.
func (o *Environment) GetFacet() EnvironmentFacet {
	if o == nil || isNil(o.Facet) {
		var ret EnvironmentFacet
		return ret
	}
	return *o.Facet
}

// GetFacetOk returns a tuple with the Facet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetFacetOk() (*EnvironmentFacet, bool) {
	if o == nil || isNil(o.Facet) {
		return nil, false
	}
	return o.Facet, true
}

// HasFacet returns a boolean if a field has been set.
func (o *Environment) HasFacet() bool {
	if o != nil && !isNil(o.Facet) {
		return true
	}

	return false
}

// SetFacet gets a reference to the given EnvironmentFacet and assigns it to the Facet field.
func (o *Environment) SetFacet(v EnvironmentFacet) {
	o.Facet = &v
}

// GetCurrentRevision returns the CurrentRevision field value if set, zero value otherwise.
func (o *Environment) GetCurrentRevision() EnvironmentCurrentRevision {
	if o == nil || isNil(o.CurrentRevision) {
		var ret EnvironmentCurrentRevision
		return ret
	}
	return *o.CurrentRevision
}

// GetCurrentRevisionOk returns a tuple with the CurrentRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetCurrentRevisionOk() (*EnvironmentCurrentRevision, bool) {
	if o == nil || isNil(o.CurrentRevision) {
		return nil, false
	}
	return o.CurrentRevision, true
}

// HasCurrentRevision returns a boolean if a field has been set.
func (o *Environment) HasCurrentRevision() bool {
	if o != nil && !isNil(o.CurrentRevision) {
		return true
	}

	return false
}

// SetCurrentRevision gets a reference to the given EnvironmentCurrentRevision and assigns it to the CurrentRevision field.
func (o *Environment) SetCurrentRevision(v EnvironmentCurrentRevision) {
	o.CurrentRevision = &v
}

// GetSources returns the Sources field value if set, zero value otherwise.
func (o *Environment) GetSources() []EnvironmentSummarySources {
	if o == nil || isNil(o.Sources) {
		var ret []EnvironmentSummarySources
		return ret
	}
	return o.Sources
}

// GetSourcesOk returns a tuple with the Sources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetSourcesOk() ([]EnvironmentSummarySources, bool) {
	if o == nil || isNil(o.Sources) {
		return nil, false
	}
	return o.Sources, true
}

// HasSources returns a boolean if a field has been set.
func (o *Environment) HasSources() bool {
	if o != nil && !isNil(o.Sources) {
		return true
	}

	return false
}

// SetSources gets a reference to the given []EnvironmentSummarySources and assigns it to the Sources field.
func (o *Environment) SetSources(v []EnvironmentSummarySources) {
	o.Sources = v
}

// GetLatestRevisions returns the LatestRevisions field value if set, zero value otherwise.
func (o *Environment) GetLatestRevisions() EnvironmentLatestRevisions {
	if o == nil || isNil(o.LatestRevisions) {
		var ret EnvironmentLatestRevisions
		return ret
	}
	return *o.LatestRevisions
}

// GetLatestRevisionsOk returns a tuple with the LatestRevisions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetLatestRevisionsOk() (*EnvironmentLatestRevisions, bool) {
	if o == nil || isNil(o.LatestRevisions) {
		return nil, false
	}
	return o.LatestRevisions, true
}

// HasLatestRevisions returns a boolean if a field has been set.
func (o *Environment) HasLatestRevisions() bool {
	if o != nil && !isNil(o.LatestRevisions) {
		return true
	}

	return false
}

// SetLatestRevisions gets a reference to the given EnvironmentLatestRevisions and assigns it to the LatestRevisions field.
func (o *Environment) SetLatestRevisions(v EnvironmentLatestRevisions) {
	o.LatestRevisions = &v
}

func (o Environment) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.Facet) {
		toSerialize["facet"] = o.Facet
	}
	if !isNil(o.CurrentRevision) {
		toSerialize["current_revision"] = o.CurrentRevision
	}
	if !isNil(o.Sources) {
		toSerialize["sources"] = o.Sources
	}
	if !isNil(o.LatestRevisions) {
		toSerialize["latest_revisions"] = o.LatestRevisions
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironment struct {
	value *Environment
	isSet bool
}

func (v NullableEnvironment) Get() *Environment {
	return v.value
}

func (v *NullableEnvironment) Set(val *Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironment(val *Environment) *NullableEnvironment {
	return &NullableEnvironment{value: val, isSet: true}
}

func (v NullableEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


