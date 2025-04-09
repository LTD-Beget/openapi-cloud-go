/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// MysqlRemoveDbResponse struct for MysqlRemoveDbResponse
type MysqlRemoveDbResponse struct {
	Success map[string]interface{} `json:"success,omitempty"`
	Error *MysqlRemoveDbResponseError `json:"error,omitempty"`
}

// NewMysqlRemoveDbResponse instantiates a new MysqlRemoveDbResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlRemoveDbResponse() *MysqlRemoveDbResponse {
	this := MysqlRemoveDbResponse{}
	return &this
}

// NewMysqlRemoveDbResponseWithDefaults instantiates a new MysqlRemoveDbResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlRemoveDbResponseWithDefaults() *MysqlRemoveDbResponse {
	this := MysqlRemoveDbResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *MysqlRemoveDbResponse) GetSuccess() map[string]interface{} {
	if o == nil || isNil(o.Success) {
		var ret map[string]interface{}
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlRemoveDbResponse) GetSuccessOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Success) {
    return map[string]interface{}{}, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *MysqlRemoveDbResponse) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given map[string]interface{} and assigns it to the Success field.
func (o *MysqlRemoveDbResponse) SetSuccess(v map[string]interface{}) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MysqlRemoveDbResponse) GetError() MysqlRemoveDbResponseError {
	if o == nil || isNil(o.Error) {
		var ret MysqlRemoveDbResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlRemoveDbResponse) GetErrorOk() (*MysqlRemoveDbResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MysqlRemoveDbResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given MysqlRemoveDbResponseError and assigns it to the Error field.
func (o *MysqlRemoveDbResponse) SetError(v MysqlRemoveDbResponseError) {
	o.Error = &v
}

func (o MysqlRemoveDbResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlRemoveDbResponse struct {
	value *MysqlRemoveDbResponse
	isSet bool
}

func (v NullableMysqlRemoveDbResponse) Get() *MysqlRemoveDbResponse {
	return v.value
}

func (v *NullableMysqlRemoveDbResponse) Set(val *MysqlRemoveDbResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlRemoveDbResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlRemoveDbResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlRemoveDbResponse(val *MysqlRemoveDbResponse) *NullableMysqlRemoveDbResponse {
	return &NullableMysqlRemoveDbResponse{value: val, isSet: true}
}

func (v NullableMysqlRemoveDbResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlRemoveDbResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


