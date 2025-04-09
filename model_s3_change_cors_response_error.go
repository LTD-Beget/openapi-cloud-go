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

// S3ChangeCorsResponseError struct for S3ChangeCorsResponseError
type S3ChangeCorsResponseError struct {
	Message *string `json:"message,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewS3ChangeCorsResponseError instantiates a new S3ChangeCorsResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3ChangeCorsResponseError() *S3ChangeCorsResponseError {
	this := S3ChangeCorsResponseError{}
	return &this
}

// NewS3ChangeCorsResponseErrorWithDefaults instantiates a new S3ChangeCorsResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ChangeCorsResponseErrorWithDefaults() *S3ChangeCorsResponseError {
	this := S3ChangeCorsResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *S3ChangeCorsResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ChangeCorsResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *S3ChangeCorsResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *S3ChangeCorsResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *S3ChangeCorsResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ChangeCorsResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *S3ChangeCorsResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *S3ChangeCorsResponseError) SetCode(v string) {
	o.Code = &v
}

func (o S3ChangeCorsResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableS3ChangeCorsResponseError struct {
	value *S3ChangeCorsResponseError
	isSet bool
}

func (v NullableS3ChangeCorsResponseError) Get() *S3ChangeCorsResponseError {
	return v.value
}

func (v *NullableS3ChangeCorsResponseError) Set(val *S3ChangeCorsResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableS3ChangeCorsResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableS3ChangeCorsResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3ChangeCorsResponseError(val *S3ChangeCorsResponseError) *NullableS3ChangeCorsResponseError {
	return &NullableS3ChangeCorsResponseError{value: val, isSet: true}
}

func (v NullableS3ChangeCorsResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3ChangeCorsResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


