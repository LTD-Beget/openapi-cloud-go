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

// MysqlConfig struct for MysqlConfig
type MysqlConfig struct {
	Param *map[string]string `json:"param,omitempty"`
	PresetName *string `json:"preset_name,omitempty"`
}

// NewMysqlConfig instantiates a new MysqlConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlConfig() *MysqlConfig {
	this := MysqlConfig{}
	return &this
}

// NewMysqlConfigWithDefaults instantiates a new MysqlConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlConfigWithDefaults() *MysqlConfig {
	this := MysqlConfig{}
	return &this
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *MysqlConfig) GetParam() map[string]string {
	if o == nil || isNil(o.Param) {
		var ret map[string]string
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlConfig) GetParamOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Param) {
    return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *MysqlConfig) HasParam() bool {
	if o != nil && !isNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given map[string]string and assigns it to the Param field.
func (o *MysqlConfig) SetParam(v map[string]string) {
	o.Param = &v
}

// GetPresetName returns the PresetName field value if set, zero value otherwise.
func (o *MysqlConfig) GetPresetName() string {
	if o == nil || isNil(o.PresetName) {
		var ret string
		return ret
	}
	return *o.PresetName
}

// GetPresetNameOk returns a tuple with the PresetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlConfig) GetPresetNameOk() (*string, bool) {
	if o == nil || isNil(o.PresetName) {
    return nil, false
	}
	return o.PresetName, true
}

// HasPresetName returns a boolean if a field has been set.
func (o *MysqlConfig) HasPresetName() bool {
	if o != nil && !isNil(o.PresetName) {
		return true
	}

	return false
}

// SetPresetName gets a reference to the given string and assigns it to the PresetName field.
func (o *MysqlConfig) SetPresetName(v string) {
	o.PresetName = &v
}

func (o MysqlConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Param) {
		toSerialize["param"] = o.Param
	}
	if !isNil(o.PresetName) {
		toSerialize["preset_name"] = o.PresetName
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlConfig struct {
	value *MysqlConfig
	isSet bool
}

func (v NullableMysqlConfig) Get() *MysqlConfig {
	return v.value
}

func (v *NullableMysqlConfig) Set(val *MysqlConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlConfig(val *MysqlConfig) *NullableMysqlConfig {
	return &NullableMysqlConfig{value: val, isSet: true}
}

func (v NullableMysqlConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


