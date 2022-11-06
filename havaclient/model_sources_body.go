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
	"fmt"
)

// SourcesBody - struct for SourcesBody
type SourcesBody struct {
	SourcesAWSCAR *SourcesAWSCAR
	SourcesAWSKey *SourcesAWSKey
	SourcesAzureCredentials *SourcesAzureCredentials
}

// SourcesAWSCARAsSourcesBody is a convenience function that returns SourcesAWSCAR wrapped in SourcesBody
func SourcesAWSCARAsSourcesBody(v *SourcesAWSCAR) SourcesBody {
	return SourcesBody{
		SourcesAWSCAR: v,
	}
}

// SourcesAWSKeyAsSourcesBody is a convenience function that returns SourcesAWSKey wrapped in SourcesBody
func SourcesAWSKeyAsSourcesBody(v *SourcesAWSKey) SourcesBody {
	return SourcesBody{
		SourcesAWSKey: v,
	}
}

// SourcesAzureCredentialsAsSourcesBody is a convenience function that returns SourcesAzureCredentials wrapped in SourcesBody
func SourcesAzureCredentialsAsSourcesBody(v *SourcesAzureCredentials) SourcesBody {
	return SourcesBody{
		SourcesAzureCredentials: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SourcesBody) UnmarshalJSON(data []byte) error {
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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SourcesAWSCAR = nil
		dst.SourcesAWSKey = nil
		dst.SourcesAzureCredentials = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SourcesBody)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SourcesBody)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SourcesBody) MarshalJSON() ([]byte, error) {
	if src.SourcesAWSCAR != nil {
		return json.Marshal(&src.SourcesAWSCAR)
	}

	if src.SourcesAWSKey != nil {
		return json.Marshal(&src.SourcesAWSKey)
	}

	if src.SourcesAzureCredentials != nil {
		return json.Marshal(&src.SourcesAzureCredentials)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SourcesBody) GetActualInstance() (interface{}) {
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

	// all schemas are nil
	return nil
}

type NullableSourcesBody struct {
	value *SourcesBody
	isSet bool
}

func (v NullableSourcesBody) Get() *SourcesBody {
	return v.value
}

func (v *NullableSourcesBody) Set(val *SourcesBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSourcesBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSourcesBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourcesBody(val *SourcesBody) *NullableSourcesBody {
	return &NullableSourcesBody{value: val, isSet: true}
}

func (v NullableSourcesBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourcesBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

