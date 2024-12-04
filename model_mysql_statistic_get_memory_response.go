/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// MysqlStatisticGetMemoryResponse struct for MysqlStatisticGetMemoryResponse
type MysqlStatisticGetMemoryResponse struct {
	Memory *StructuresStatisticSeriesData `json:"memory,omitempty"`
}

// NewMysqlStatisticGetMemoryResponse instantiates a new MysqlStatisticGetMemoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlStatisticGetMemoryResponse() *MysqlStatisticGetMemoryResponse {
	this := MysqlStatisticGetMemoryResponse{}
	return &this
}

// NewMysqlStatisticGetMemoryResponseWithDefaults instantiates a new MysqlStatisticGetMemoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlStatisticGetMemoryResponseWithDefaults() *MysqlStatisticGetMemoryResponse {
	this := MysqlStatisticGetMemoryResponse{}
	return &this
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *MysqlStatisticGetMemoryResponse) GetMemory() StructuresStatisticSeriesData {
	if o == nil || isNil(o.Memory) {
		var ret StructuresStatisticSeriesData
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetMemoryResponse) GetMemoryOk() (*StructuresStatisticSeriesData, bool) {
	if o == nil || isNil(o.Memory) {
    return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *MysqlStatisticGetMemoryResponse) HasMemory() bool {
	if o != nil && !isNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given StructuresStatisticSeriesData and assigns it to the Memory field.
func (o *MysqlStatisticGetMemoryResponse) SetMemory(v StructuresStatisticSeriesData) {
	o.Memory = &v
}

func (o MysqlStatisticGetMemoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlStatisticGetMemoryResponse struct {
	value *MysqlStatisticGetMemoryResponse
	isSet bool
}

func (v NullableMysqlStatisticGetMemoryResponse) Get() *MysqlStatisticGetMemoryResponse {
	return v.value
}

func (v *NullableMysqlStatisticGetMemoryResponse) Set(val *MysqlStatisticGetMemoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlStatisticGetMemoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlStatisticGetMemoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlStatisticGetMemoryResponse(val *MysqlStatisticGetMemoryResponse) *NullableMysqlStatisticGetMemoryResponse {
	return &NullableMysqlStatisticGetMemoryResponse{value: val, isSet: true}
}

func (v NullableMysqlStatisticGetMemoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlStatisticGetMemoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


