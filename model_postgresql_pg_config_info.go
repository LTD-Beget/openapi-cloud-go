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

// PostgresqlPgConfigInfo struct for PostgresqlPgConfigInfo
type PostgresqlPgConfigInfo struct {
	ParamDefault *map[string]string `json:"param_default,omitempty"`
	ParamConfig []StructuresParamConfig `json:"param_config,omitempty"`
	Param *map[string]string `json:"param,omitempty"`
}

// NewPostgresqlPgConfigInfo instantiates a new PostgresqlPgConfigInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgConfigInfo() *PostgresqlPgConfigInfo {
	this := PostgresqlPgConfigInfo{}
	return &this
}

// NewPostgresqlPgConfigInfoWithDefaults instantiates a new PostgresqlPgConfigInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgConfigInfoWithDefaults() *PostgresqlPgConfigInfo {
	this := PostgresqlPgConfigInfo{}
	return &this
}

// GetParamDefault returns the ParamDefault field value if set, zero value otherwise.
func (o *PostgresqlPgConfigInfo) GetParamDefault() map[string]string {
	if o == nil || isNil(o.ParamDefault) {
		var ret map[string]string
		return ret
	}
	return *o.ParamDefault
}

// GetParamDefaultOk returns a tuple with the ParamDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgConfigInfo) GetParamDefaultOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ParamDefault) {
    return nil, false
	}
	return o.ParamDefault, true
}

// HasParamDefault returns a boolean if a field has been set.
func (o *PostgresqlPgConfigInfo) HasParamDefault() bool {
	if o != nil && !isNil(o.ParamDefault) {
		return true
	}

	return false
}

// SetParamDefault gets a reference to the given map[string]string and assigns it to the ParamDefault field.
func (o *PostgresqlPgConfigInfo) SetParamDefault(v map[string]string) {
	o.ParamDefault = &v
}

// GetParamConfig returns the ParamConfig field value if set, zero value otherwise.
func (o *PostgresqlPgConfigInfo) GetParamConfig() []StructuresParamConfig {
	if o == nil || isNil(o.ParamConfig) {
		var ret []StructuresParamConfig
		return ret
	}
	return o.ParamConfig
}

// GetParamConfigOk returns a tuple with the ParamConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgConfigInfo) GetParamConfigOk() ([]StructuresParamConfig, bool) {
	if o == nil || isNil(o.ParamConfig) {
    return nil, false
	}
	return o.ParamConfig, true
}

// HasParamConfig returns a boolean if a field has been set.
func (o *PostgresqlPgConfigInfo) HasParamConfig() bool {
	if o != nil && !isNil(o.ParamConfig) {
		return true
	}

	return false
}

// SetParamConfig gets a reference to the given []StructuresParamConfig and assigns it to the ParamConfig field.
func (o *PostgresqlPgConfigInfo) SetParamConfig(v []StructuresParamConfig) {
	o.ParamConfig = v
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *PostgresqlPgConfigInfo) GetParam() map[string]string {
	if o == nil || isNil(o.Param) {
		var ret map[string]string
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgConfigInfo) GetParamOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Param) {
    return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *PostgresqlPgConfigInfo) HasParam() bool {
	if o != nil && !isNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given map[string]string and assigns it to the Param field.
func (o *PostgresqlPgConfigInfo) SetParam(v map[string]string) {
	o.Param = &v
}

func (o PostgresqlPgConfigInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ParamDefault) {
		toSerialize["param_default"] = o.ParamDefault
	}
	if !isNil(o.ParamConfig) {
		toSerialize["param_config"] = o.ParamConfig
	}
	if !isNil(o.Param) {
		toSerialize["param"] = o.Param
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgConfigInfo struct {
	value *PostgresqlPgConfigInfo
	isSet bool
}

func (v NullablePostgresqlPgConfigInfo) Get() *PostgresqlPgConfigInfo {
	return v.value
}

func (v *NullablePostgresqlPgConfigInfo) Set(val *PostgresqlPgConfigInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgConfigInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgConfigInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgConfigInfo(val *PostgresqlPgConfigInfo) *NullablePostgresqlPgConfigInfo {
	return &NullablePostgresqlPgConfigInfo{value: val, isSet: true}
}

func (v NullablePostgresqlPgConfigInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgConfigInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


