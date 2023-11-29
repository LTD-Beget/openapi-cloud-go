/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// MysqlStatisticGetCpuResponse struct for MysqlStatisticGetCpuResponse
type MysqlStatisticGetCpuResponse struct {
	Cpu *StructuresStatisticSeriesData `json:"cpu,omitempty"`
}

// NewMysqlStatisticGetCpuResponse instantiates a new MysqlStatisticGetCpuResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlStatisticGetCpuResponse() *MysqlStatisticGetCpuResponse {
	this := MysqlStatisticGetCpuResponse{}
	return &this
}

// NewMysqlStatisticGetCpuResponseWithDefaults instantiates a new MysqlStatisticGetCpuResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlStatisticGetCpuResponseWithDefaults() *MysqlStatisticGetCpuResponse {
	this := MysqlStatisticGetCpuResponse{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *MysqlStatisticGetCpuResponse) GetCpu() StructuresStatisticSeriesData {
	if o == nil || isNil(o.Cpu) {
		var ret StructuresStatisticSeriesData
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetCpuResponse) GetCpuOk() (*StructuresStatisticSeriesData, bool) {
	if o == nil || isNil(o.Cpu) {
    return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *MysqlStatisticGetCpuResponse) HasCpu() bool {
	if o != nil && !isNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given StructuresStatisticSeriesData and assigns it to the Cpu field.
func (o *MysqlStatisticGetCpuResponse) SetCpu(v StructuresStatisticSeriesData) {
	o.Cpu = &v
}

func (o MysqlStatisticGetCpuResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlStatisticGetCpuResponse struct {
	value *MysqlStatisticGetCpuResponse
	isSet bool
}

func (v NullableMysqlStatisticGetCpuResponse) Get() *MysqlStatisticGetCpuResponse {
	return v.value
}

func (v *NullableMysqlStatisticGetCpuResponse) Set(val *MysqlStatisticGetCpuResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlStatisticGetCpuResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlStatisticGetCpuResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlStatisticGetCpuResponse(val *MysqlStatisticGetCpuResponse) *NullableMysqlStatisticGetCpuResponse {
	return &NullableMysqlStatisticGetCpuResponse{value: val, isSet: true}
}

func (v NullableMysqlStatisticGetCpuResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlStatisticGetCpuResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


