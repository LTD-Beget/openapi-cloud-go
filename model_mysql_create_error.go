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

// MysqlCreateError struct for MysqlCreateError
type MysqlCreateError struct {
	Message *string `json:"message,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewMysqlCreateError instantiates a new MysqlCreateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlCreateError() *MysqlCreateError {
	this := MysqlCreateError{}
	return &this
}

// NewMysqlCreateErrorWithDefaults instantiates a new MysqlCreateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlCreateErrorWithDefaults() *MysqlCreateError {
	this := MysqlCreateError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *MysqlCreateError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *MysqlCreateError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *MysqlCreateError) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *MysqlCreateError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *MysqlCreateError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *MysqlCreateError) SetCode(v string) {
	o.Code = &v
}

func (o MysqlCreateError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlCreateError struct {
	value *MysqlCreateError
	isSet bool
}

func (v NullableMysqlCreateError) Get() *MysqlCreateError {
	return v.value
}

func (v *NullableMysqlCreateError) Set(val *MysqlCreateError) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlCreateError) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlCreateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlCreateError(val *MysqlCreateError) *NullableMysqlCreateError {
	return &NullableMysqlCreateError{value: val, isSet: true}
}

func (v NullableMysqlCreateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlCreateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


