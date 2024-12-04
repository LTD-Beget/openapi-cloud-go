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

// PostgresqlStatisticGetCpuResponse struct for PostgresqlStatisticGetCpuResponse
type PostgresqlStatisticGetCpuResponse struct {
	Cpu *StructuresStatisticSeriesData `json:"cpu,omitempty"`
}

// NewPostgresqlStatisticGetCpuResponse instantiates a new PostgresqlStatisticGetCpuResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlStatisticGetCpuResponse() *PostgresqlStatisticGetCpuResponse {
	this := PostgresqlStatisticGetCpuResponse{}
	return &this
}

// NewPostgresqlStatisticGetCpuResponseWithDefaults instantiates a new PostgresqlStatisticGetCpuResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlStatisticGetCpuResponseWithDefaults() *PostgresqlStatisticGetCpuResponse {
	this := PostgresqlStatisticGetCpuResponse{}
	return &this
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *PostgresqlStatisticGetCpuResponse) GetCpu() StructuresStatisticSeriesData {
	if o == nil || isNil(o.Cpu) {
		var ret StructuresStatisticSeriesData
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlStatisticGetCpuResponse) GetCpuOk() (*StructuresStatisticSeriesData, bool) {
	if o == nil || isNil(o.Cpu) {
    return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *PostgresqlStatisticGetCpuResponse) HasCpu() bool {
	if o != nil && !isNil(o.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given StructuresStatisticSeriesData and assigns it to the Cpu field.
func (o *PostgresqlStatisticGetCpuResponse) SetCpu(v StructuresStatisticSeriesData) {
	o.Cpu = &v
}

func (o PostgresqlStatisticGetCpuResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlStatisticGetCpuResponse struct {
	value *PostgresqlStatisticGetCpuResponse
	isSet bool
}

func (v NullablePostgresqlStatisticGetCpuResponse) Get() *PostgresqlStatisticGetCpuResponse {
	return v.value
}

func (v *NullablePostgresqlStatisticGetCpuResponse) Set(val *PostgresqlStatisticGetCpuResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlStatisticGetCpuResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlStatisticGetCpuResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlStatisticGetCpuResponse(val *PostgresqlStatisticGetCpuResponse) *NullablePostgresqlStatisticGetCpuResponse {
	return &NullablePostgresqlStatisticGetCpuResponse{value: val, isSet: true}
}

func (v NullablePostgresqlStatisticGetCpuResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlStatisticGetCpuResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


