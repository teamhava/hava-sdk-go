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
	"fmt"
)

// SourcesCreateRequest - struct for SourcesCreateRequest
type SourcesCreateRequest struct {
	SourcesAWSCAR                       *SourcesAWSCAR
	SourcesAWSKey                       *SourcesAWSKey
	SourcesAzureCredentials             *SourcesAzureCredentials
	SourcesGCPServiceAccountCredentials *SourcesGCPServiceAccountCredentials
}

// SourcesAWSCARAsSourcesCreateRequest is a convenience function that returns SourcesAWSCAR wrapped in SourcesCreateRequest
func SourcesAWSCARAsSourcesCreateRequest(v *SourcesAWSCAR) SourcesCreateRequest {
	return SourcesCreateRequest{
		SourcesAWSCAR: v,
	}
}

// SourcesAWSKeyAsSourcesCreateRequest is a convenience function that returns SourcesAWSKey wrapped in SourcesCreateRequest
func SourcesAWSKeyAsSourcesCreateRequest(v *SourcesAWSKey) SourcesCreateRequest {
	return SourcesCreateRequest{
		SourcesAWSKey: v,
	}
}

// SourcesAzureCredentialsAsSourcesCreateRequest is a convenience function that returns SourcesAzureCredentials wrapped in SourcesCreateRequest
func SourcesAzureCredentialsAsSourcesCreateRequest(v *SourcesAzureCredentials) SourcesCreateRequest {
	return SourcesCreateRequest{
		SourcesAzureCredentials: v,
	}
}

// SourcesGCPServiceAccountCredentialsAsSourcesCreateRequest is a convenience function that returns SourcesGCPServiceAccountCredentials wrapped in SourcesCreateRequest
func SourcesGCPServiceAccountCredentialsAsSourcesCreateRequest(v *SourcesGCPServiceAccountCredentials) SourcesCreateRequest {
	return SourcesCreateRequest{
		SourcesGCPServiceAccountCredentials: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SourcesCreateRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SourcesAWSCAR
	err = newStrictDecoder(data).Decode(&dst.SourcesAWSCAR)
	if err == nil {
		jsonSourcesAWSCAR, _ := json.Marshal(dst.SourcesAWSCAR)
		if string(jsonSourcesAWSCAR) == "{}" { // empty struct
			dst.SourcesAWSCAR = nil
		} else {
			match++
		}
	} else {
		dst.SourcesAWSCAR = nil
	}

	// try to unmarshal data into SourcesAWSKey
	err = newStrictDecoder(data).Decode(&dst.SourcesAWSKey)
	if err == nil {
		jsonSourcesAWSKey, _ := json.Marshal(dst.SourcesAWSKey)
		if string(jsonSourcesAWSKey) == "{}" { // empty struct
			dst.SourcesAWSKey = nil
		} else {
			match++
		}
	} else {
		dst.SourcesAWSKey = nil
	}

	// try to unmarshal data into SourcesAzureCredentials
	err = newStrictDecoder(data).Decode(&dst.SourcesAzureCredentials)
	if err == nil {
		jsonSourcesAzureCredentials, _ := json.Marshal(dst.SourcesAzureCredentials)
		if string(jsonSourcesAzureCredentials) == "{}" { // empty struct
			dst.SourcesAzureCredentials = nil
		} else {
			match++
		}
	} else {
		dst.SourcesAzureCredentials = nil
	}

	// try to unmarshal data into SourcesGCPServiceAccountCredentials
	err = newStrictDecoder(data).Decode(&dst.SourcesGCPServiceAccountCredentials)
	if err == nil {
		jsonSourcesGCPServiceAccountCredentials, _ := json.Marshal(dst.SourcesGCPServiceAccountCredentials)
		if string(jsonSourcesGCPServiceAccountCredentials) == "{}" { // empty struct
			dst.SourcesGCPServiceAccountCredentials = nil
		} else {
			match++
		}
	} else {
		dst.SourcesGCPServiceAccountCredentials = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SourcesAWSCAR = nil
		dst.SourcesAWSKey = nil
		dst.SourcesAzureCredentials = nil
		dst.SourcesGCPServiceAccountCredentials = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SourcesCreateRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SourcesCreateRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SourcesCreateRequest) MarshalJSON() ([]byte, error) {
	if src.SourcesAWSCAR != nil {
		return json.Marshal(&src.SourcesAWSCAR)
	}

	if src.SourcesAWSKey != nil {
		return json.Marshal(&src.SourcesAWSKey)
	}

	if src.SourcesAzureCredentials != nil {
		return json.Marshal(&src.SourcesAzureCredentials)
	}

	if src.SourcesGCPServiceAccountCredentials != nil {
		return json.Marshal(&src.SourcesGCPServiceAccountCredentials)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SourcesCreateRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.SourcesAWSCAR != nil {
		return obj.SourcesAWSCAR
	}

	if obj.SourcesAWSKey != nil {
		return obj.SourcesAWSKey
	}

	if obj.SourcesAzureCredentials != nil {
		return obj.SourcesAzureCredentials
	}

	if obj.SourcesGCPServiceAccountCredentials != nil {
		return obj.SourcesGCPServiceAccountCredentials
	}

	// all schemas are nil
	return nil
}

type NullableSourcesCreateRequest struct {
	value *SourcesCreateRequest
	isSet bool
}

func (v NullableSourcesCreateRequest) Get() *SourcesCreateRequest {
	return v.value
}

func (v *NullableSourcesCreateRequest) Set(val *SourcesCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcesCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcesCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcesCreateRequest(val *SourcesCreateRequest) *NullableSourcesCreateRequest {
	return &NullableSourcesCreateRequest{value: val, isSet: true}
}

func (v NullableSourcesCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcesCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
