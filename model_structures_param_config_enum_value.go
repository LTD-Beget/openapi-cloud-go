/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// StructuresParamConfigEnumValue struct for StructuresParamConfigEnumValue
type StructuresParamConfigEnumValue struct {
	Value []string `json:"value,omitempty"`
}

// NewStructuresParamConfigEnumValue instantiates a new StructuresParamConfigEnumValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresParamConfigEnumValue() *StructuresParamConfigEnumValue {
	this := StructuresParamConfigEnumValue{}
	return &this
}

// NewStructuresParamConfigEnumValueWithDefaults instantiates a new StructuresParamConfigEnumValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresParamConfigEnumValueWithDefaults() *StructuresParamConfigEnumValue {
	this := StructuresParamConfigEnumValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *StructuresParamConfigEnumValue) GetValue() []string {
	if o == nil || isNil(o.Value) {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresParamConfigEnumValue) GetValueOk() ([]string, bool) {
	if o == nil || isNil(o.Value) {
    return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *StructuresParamConfigEnumValue) HasValue() bool {
	if o != nil && !isNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *StructuresParamConfigEnumValue) SetValue(v []string) {
	o.Value = v
}

func (o StructuresParamConfigEnumValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresParamConfigEnumValue struct {
	value *StructuresParamConfigEnumValue
	isSet bool
}

func (v NullableStructuresParamConfigEnumValue) Get() *StructuresParamConfigEnumValue {
	return v.value
}

func (v *NullableStructuresParamConfigEnumValue) Set(val *StructuresParamConfigEnumValue) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresParamConfigEnumValue) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresParamConfigEnumValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresParamConfigEnumValue(val *StructuresParamConfigEnumValue) *NullableStructuresParamConfigEnumValue {
	return &NullableStructuresParamConfigEnumValue{value: val, isSet: true}
}

func (v NullableStructuresParamConfigEnumValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresParamConfigEnumValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

