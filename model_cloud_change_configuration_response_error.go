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

// CloudChangeConfigurationResponseError struct for CloudChangeConfigurationResponseError
type CloudChangeConfigurationResponseError struct {
	Message *string `json:"message,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewCloudChangeConfigurationResponseError instantiates a new CloudChangeConfigurationResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudChangeConfigurationResponseError() *CloudChangeConfigurationResponseError {
	this := CloudChangeConfigurationResponseError{}
	return &this
}

// NewCloudChangeConfigurationResponseErrorWithDefaults instantiates a new CloudChangeConfigurationResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudChangeConfigurationResponseErrorWithDefaults() *CloudChangeConfigurationResponseError {
	this := CloudChangeConfigurationResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CloudChangeConfigurationResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudChangeConfigurationResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CloudChangeConfigurationResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CloudChangeConfigurationResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CloudChangeConfigurationResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudChangeConfigurationResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CloudChangeConfigurationResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CloudChangeConfigurationResponseError) SetCode(v string) {
	o.Code = &v
}

func (o CloudChangeConfigurationResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableCloudChangeConfigurationResponseError struct {
	value *CloudChangeConfigurationResponseError
	isSet bool
}

func (v NullableCloudChangeConfigurationResponseError) Get() *CloudChangeConfigurationResponseError {
	return v.value
}

func (v *NullableCloudChangeConfigurationResponseError) Set(val *CloudChangeConfigurationResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudChangeConfigurationResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudChangeConfigurationResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudChangeConfigurationResponseError(val *CloudChangeConfigurationResponseError) *NullableCloudChangeConfigurationResponseError {
	return &NullableCloudChangeConfigurationResponseError{value: val, isSet: true}
}

func (v NullableCloudChangeConfigurationResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudChangeConfigurationResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


