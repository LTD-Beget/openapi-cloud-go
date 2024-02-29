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

// MysqlRemoveAccessResponse struct for MysqlRemoveAccessResponse
type MysqlRemoveAccessResponse struct {
	Success map[string]interface{} `json:"success,omitempty"`
	Error *MysqlRemoveAccessResponseError `json:"error,omitempty"`
}

// NewMysqlRemoveAccessResponse instantiates a new MysqlRemoveAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlRemoveAccessResponse() *MysqlRemoveAccessResponse {
	this := MysqlRemoveAccessResponse{}
	return &this
}

// NewMysqlRemoveAccessResponseWithDefaults instantiates a new MysqlRemoveAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlRemoveAccessResponseWithDefaults() *MysqlRemoveAccessResponse {
	this := MysqlRemoveAccessResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *MysqlRemoveAccessResponse) GetSuccess() map[string]interface{} {
	if o == nil || isNil(o.Success) {
		var ret map[string]interface{}
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlRemoveAccessResponse) GetSuccessOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Success) {
    return map[string]interface{}{}, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *MysqlRemoveAccessResponse) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given map[string]interface{} and assigns it to the Success field.
func (o *MysqlRemoveAccessResponse) SetSuccess(v map[string]interface{}) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MysqlRemoveAccessResponse) GetError() MysqlRemoveAccessResponseError {
	if o == nil || isNil(o.Error) {
		var ret MysqlRemoveAccessResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlRemoveAccessResponse) GetErrorOk() (*MysqlRemoveAccessResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MysqlRemoveAccessResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given MysqlRemoveAccessResponseError and assigns it to the Error field.
func (o *MysqlRemoveAccessResponse) SetError(v MysqlRemoveAccessResponseError) {
	o.Error = &v
}

func (o MysqlRemoveAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlRemoveAccessResponse struct {
	value *MysqlRemoveAccessResponse
	isSet bool
}

func (v NullableMysqlRemoveAccessResponse) Get() *MysqlRemoveAccessResponse {
	return v.value
}

func (v *NullableMysqlRemoveAccessResponse) Set(val *MysqlRemoveAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlRemoveAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlRemoveAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlRemoveAccessResponse(val *MysqlRemoveAccessResponse) *NullableMysqlRemoveAccessResponse {
	return &NullableMysqlRemoveAccessResponse{value: val, isSet: true}
}

func (v NullableMysqlRemoveAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlRemoveAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


