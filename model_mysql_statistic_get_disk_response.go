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

// MysqlStatisticGetDiskResponse struct for MysqlStatisticGetDiskResponse
type MysqlStatisticGetDiskResponse struct {
	DataRead *StructuresStatisticSeriesData `json:"data_read,omitempty"`
	DataWrite *StructuresStatisticSeriesData `json:"data_write,omitempty"`
}

// NewMysqlStatisticGetDiskResponse instantiates a new MysqlStatisticGetDiskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlStatisticGetDiskResponse() *MysqlStatisticGetDiskResponse {
	this := MysqlStatisticGetDiskResponse{}
	return &this
}

// NewMysqlStatisticGetDiskResponseWithDefaults instantiates a new MysqlStatisticGetDiskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlStatisticGetDiskResponseWithDefaults() *MysqlStatisticGetDiskResponse {
	this := MysqlStatisticGetDiskResponse{}
	return &this
}

// GetDataRead returns the DataRead field value if set, zero value otherwise.
func (o *MysqlStatisticGetDiskResponse) GetDataRead() StructuresStatisticSeriesData {
	if o == nil || isNil(o.DataRead) {
		var ret StructuresStatisticSeriesData
		return ret
	}
	return *o.DataRead
}

// GetDataReadOk returns a tuple with the DataRead field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetDiskResponse) GetDataReadOk() (*StructuresStatisticSeriesData, bool) {
	if o == nil || isNil(o.DataRead) {
    return nil, false
	}
	return o.DataRead, true
}

// HasDataRead returns a boolean if a field has been set.
func (o *MysqlStatisticGetDiskResponse) HasDataRead() bool {
	if o != nil && !isNil(o.DataRead) {
		return true
	}

	return false
}

// SetDataRead gets a reference to the given StructuresStatisticSeriesData and assigns it to the DataRead field.
func (o *MysqlStatisticGetDiskResponse) SetDataRead(v StructuresStatisticSeriesData) {
	o.DataRead = &v
}

// GetDataWrite returns the DataWrite field value if set, zero value otherwise.
func (o *MysqlStatisticGetDiskResponse) GetDataWrite() StructuresStatisticSeriesData {
	if o == nil || isNil(o.DataWrite) {
		var ret StructuresStatisticSeriesData
		return ret
	}
	return *o.DataWrite
}

// GetDataWriteOk returns a tuple with the DataWrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetDiskResponse) GetDataWriteOk() (*StructuresStatisticSeriesData, bool) {
	if o == nil || isNil(o.DataWrite) {
    return nil, false
	}
	return o.DataWrite, true
}

// HasDataWrite returns a boolean if a field has been set.
func (o *MysqlStatisticGetDiskResponse) HasDataWrite() bool {
	if o != nil && !isNil(o.DataWrite) {
		return true
	}

	return false
}

// SetDataWrite gets a reference to the given StructuresStatisticSeriesData and assigns it to the DataWrite field.
func (o *MysqlStatisticGetDiskResponse) SetDataWrite(v StructuresStatisticSeriesData) {
	o.DataWrite = &v
}

func (o MysqlStatisticGetDiskResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DataRead) {
		toSerialize["data_read"] = o.DataRead
	}
	if !isNil(o.DataWrite) {
		toSerialize["data_write"] = o.DataWrite
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlStatisticGetDiskResponse struct {
	value *MysqlStatisticGetDiskResponse
	isSet bool
}

func (v NullableMysqlStatisticGetDiskResponse) Get() *MysqlStatisticGetDiskResponse {
	return v.value
}

func (v *NullableMysqlStatisticGetDiskResponse) Set(val *MysqlStatisticGetDiskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlStatisticGetDiskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlStatisticGetDiskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlStatisticGetDiskResponse(val *MysqlStatisticGetDiskResponse) *NullableMysqlStatisticGetDiskResponse {
	return &NullableMysqlStatisticGetDiskResponse{value: val, isSet: true}
}

func (v NullableMysqlStatisticGetDiskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlStatisticGetDiskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


