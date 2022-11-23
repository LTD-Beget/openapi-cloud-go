/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// MysqlSetConfigResponse struct for MysqlSetConfigResponse
type MysqlSetConfigResponse struct {
	Config *MysqlConfig `json:"config,omitempty"`
	Error *MysqlSetConfigResponseError `json:"error,omitempty"`
}

// NewMysqlSetConfigResponse instantiates a new MysqlSetConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlSetConfigResponse() *MysqlSetConfigResponse {
	this := MysqlSetConfigResponse{}
	return &this
}

// NewMysqlSetConfigResponseWithDefaults instantiates a new MysqlSetConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlSetConfigResponseWithDefaults() *MysqlSetConfigResponse {
	this := MysqlSetConfigResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *MysqlSetConfigResponse) GetConfig() MysqlConfig {
	if o == nil || isNil(o.Config) {
		var ret MysqlConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSetConfigResponse) GetConfigOk() (*MysqlConfig, bool) {
	if o == nil || isNil(o.Config) {
    return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *MysqlSetConfigResponse) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given MysqlConfig and assigns it to the Config field.
func (o *MysqlSetConfigResponse) SetConfig(v MysqlConfig) {
	o.Config = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MysqlSetConfigResponse) GetError() MysqlSetConfigResponseError {
	if o == nil || isNil(o.Error) {
		var ret MysqlSetConfigResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlSetConfigResponse) GetErrorOk() (*MysqlSetConfigResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MysqlSetConfigResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given MysqlSetConfigResponseError and assigns it to the Error field.
func (o *MysqlSetConfigResponse) SetError(v MysqlSetConfigResponseError) {
	o.Error = &v
}

func (o MysqlSetConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlSetConfigResponse struct {
	value *MysqlSetConfigResponse
	isSet bool
}

func (v NullableMysqlSetConfigResponse) Get() *MysqlSetConfigResponse {
	return v.value
}

func (v *NullableMysqlSetConfigResponse) Set(val *MysqlSetConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlSetConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlSetConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlSetConfigResponse(val *MysqlSetConfigResponse) *NullableMysqlSetConfigResponse {
	return &NullableMysqlSetConfigResponse{value: val, isSet: true}
}

func (v NullableMysqlSetConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlSetConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


