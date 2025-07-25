/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// S3ChangeDomainResponseError struct for S3ChangeDomainResponseError
type S3ChangeDomainResponseError struct {
	Message *string `json:"message,omitempty"`
	Code *string `json:"code,omitempty"`
}

// NewS3ChangeDomainResponseError instantiates a new S3ChangeDomainResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3ChangeDomainResponseError() *S3ChangeDomainResponseError {
	this := S3ChangeDomainResponseError{}
	return &this
}

// NewS3ChangeDomainResponseErrorWithDefaults instantiates a new S3ChangeDomainResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ChangeDomainResponseErrorWithDefaults() *S3ChangeDomainResponseError {
	this := S3ChangeDomainResponseError{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *S3ChangeDomainResponseError) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ChangeDomainResponseError) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *S3ChangeDomainResponseError) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *S3ChangeDomainResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *S3ChangeDomainResponseError) GetCode() string {
	if o == nil || isNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ChangeDomainResponseError) GetCodeOk() (*string, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *S3ChangeDomainResponseError) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *S3ChangeDomainResponseError) SetCode(v string) {
	o.Code = &v
}

func (o S3ChangeDomainResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	return json.Marshal(toSerialize)
}

type NullableS3ChangeDomainResponseError struct {
	value *S3ChangeDomainResponseError
	isSet bool
}

func (v NullableS3ChangeDomainResponseError) Get() *S3ChangeDomainResponseError {
	return v.value
}

func (v *NullableS3ChangeDomainResponseError) Set(val *S3ChangeDomainResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableS3ChangeDomainResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableS3ChangeDomainResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3ChangeDomainResponseError(val *S3ChangeDomainResponseError) *NullableS3ChangeDomainResponseError {
	return &NullableS3ChangeDomainResponseError{value: val, isSet: true}
}

func (v NullableS3ChangeDomainResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3ChangeDomainResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


