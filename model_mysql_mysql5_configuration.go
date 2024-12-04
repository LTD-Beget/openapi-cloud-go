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

// MysqlMysql5Configuration struct for MysqlMysql5Configuration
type MysqlMysql5Configuration struct {
	CpuCount *int32 `json:"cpu_count,omitempty"`
	DiskSize *int32 `json:"disk_size,omitempty"`
	Memory *int32 `json:"memory,omitempty"`
	Version *string `json:"version,omitempty"`
	DisplayVersion *string `json:"display_version,omitempty"`
	ParamConfig []StructuresParamConfig `json:"param_config,omitempty"`
	ParamDefault *map[string]string `json:"param_default,omitempty"`
}

// NewMysqlMysql5Configuration instantiates a new MysqlMysql5Configuration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlMysql5Configuration() *MysqlMysql5Configuration {
	this := MysqlMysql5Configuration{}
	return &this
}

// NewMysqlMysql5ConfigurationWithDefaults instantiates a new MysqlMysql5Configuration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlMysql5ConfigurationWithDefaults() *MysqlMysql5Configuration {
	this := MysqlMysql5Configuration{}
	return &this
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *MysqlMysql5Configuration) GetCpuCount() int32 {
	if o == nil || isNil(o.CpuCount) {
		var ret int32
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql5Configuration) GetCpuCountOk() (*int32, bool) {
	if o == nil || isNil(o.CpuCount) {
    return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *MysqlMysql5Configuration) HasCpuCount() bool {
	if o != nil && !isNil(o.CpuCount) {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int32 and assigns it to the CpuCount field.
func (o *MysqlMysql5Configuration) SetCpuCount(v int32) {
	o.CpuCount = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *MysqlMysql5Configuration) GetDiskSize() int32 {
	if o == nil || isNil(o.DiskSize) {
		var ret int32
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql5Configuration) GetDiskSizeOk() (*int32, bool) {
	if o == nil || isNil(o.DiskSize) {
    return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *MysqlMysql5Configuration) HasDiskSize() bool {
	if o != nil && !isNil(o.DiskSize) {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int32 and assigns it to the DiskSize field.
func (o *MysqlMysql5Configuration) SetDiskSize(v int32) {
	o.DiskSize = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *MysqlMysql5Configuration) GetMemory() int32 {
	if o == nil || isNil(o.Memory) {
		var ret int32
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql5Configuration) GetMemoryOk() (*int32, bool) {
	if o == nil || isNil(o.Memory) {
    return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *MysqlMysql5Configuration) HasMemory() bool {
	if o != nil && !isNil(o.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int32 and assigns it to the Memory field.
func (o *MysqlMysql5Configuration) SetMemory(v int32) {
	o.Memory = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *MysqlMysql5Configuration) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql5Configuration) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *MysqlMysql5Configuration) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *MysqlMysql5Configuration) SetVersion(v string) {
	o.Version = &v
}

// GetDisplayVersion returns the DisplayVersion field value if set, zero value otherwise.
func (o *MysqlMysql5Configuration) GetDisplayVersion() string {
	if o == nil || isNil(o.DisplayVersion) {
		var ret string
		return ret
	}
	return *o.DisplayVersion
}

// GetDisplayVersionOk returns a tuple with the DisplayVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql5Configuration) GetDisplayVersionOk() (*string, bool) {
	if o == nil || isNil(o.DisplayVersion) {
    return nil, false
	}
	return o.DisplayVersion, true
}

// HasDisplayVersion returns a boolean if a field has been set.
func (o *MysqlMysql5Configuration) HasDisplayVersion() bool {
	if o != nil && !isNil(o.DisplayVersion) {
		return true
	}

	return false
}

// SetDisplayVersion gets a reference to the given string and assigns it to the DisplayVersion field.
func (o *MysqlMysql5Configuration) SetDisplayVersion(v string) {
	o.DisplayVersion = &v
}

// GetParamConfig returns the ParamConfig field value if set, zero value otherwise.
func (o *MysqlMysql5Configuration) GetParamConfig() []StructuresParamConfig {
	if o == nil || isNil(o.ParamConfig) {
		var ret []StructuresParamConfig
		return ret
	}
	return o.ParamConfig
}

// GetParamConfigOk returns a tuple with the ParamConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql5Configuration) GetParamConfigOk() ([]StructuresParamConfig, bool) {
	if o == nil || isNil(o.ParamConfig) {
    return nil, false
	}
	return o.ParamConfig, true
}

// HasParamConfig returns a boolean if a field has been set.
func (o *MysqlMysql5Configuration) HasParamConfig() bool {
	if o != nil && !isNil(o.ParamConfig) {
		return true
	}

	return false
}

// SetParamConfig gets a reference to the given []StructuresParamConfig and assigns it to the ParamConfig field.
func (o *MysqlMysql5Configuration) SetParamConfig(v []StructuresParamConfig) {
	o.ParamConfig = v
}

// GetParamDefault returns the ParamDefault field value if set, zero value otherwise.
func (o *MysqlMysql5Configuration) GetParamDefault() map[string]string {
	if o == nil || isNil(o.ParamDefault) {
		var ret map[string]string
		return ret
	}
	return *o.ParamDefault
}

// GetParamDefaultOk returns a tuple with the ParamDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql5Configuration) GetParamDefaultOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ParamDefault) {
    return nil, false
	}
	return o.ParamDefault, true
}

// HasParamDefault returns a boolean if a field has been set.
func (o *MysqlMysql5Configuration) HasParamDefault() bool {
	if o != nil && !isNil(o.ParamDefault) {
		return true
	}

	return false
}

// SetParamDefault gets a reference to the given map[string]string and assigns it to the ParamDefault field.
func (o *MysqlMysql5Configuration) SetParamDefault(v map[string]string) {
	o.ParamDefault = &v
}

func (o MysqlMysql5Configuration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CpuCount) {
		toSerialize["cpu_count"] = o.CpuCount
	}
	if !isNil(o.DiskSize) {
		toSerialize["disk_size"] = o.DiskSize
	}
	if !isNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.DisplayVersion) {
		toSerialize["display_version"] = o.DisplayVersion
	}
	if !isNil(o.ParamConfig) {
		toSerialize["param_config"] = o.ParamConfig
	}
	if !isNil(o.ParamDefault) {
		toSerialize["param_default"] = o.ParamDefault
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlMysql5Configuration struct {
	value *MysqlMysql5Configuration
	isSet bool
}

func (v NullableMysqlMysql5Configuration) Get() *MysqlMysql5Configuration {
	return v.value
}

func (v *NullableMysqlMysql5Configuration) Set(val *MysqlMysql5Configuration) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlMysql5Configuration) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlMysql5Configuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlMysql5Configuration(val *MysqlMysql5Configuration) *NullableMysqlMysql5Configuration {
	return &NullableMysqlMysql5Configuration{value: val, isSet: true}
}

func (v NullableMysqlMysql5Configuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlMysql5Configuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


