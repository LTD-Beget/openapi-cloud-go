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

// S3ChangeDomainResponse struct for S3ChangeDomainResponse
type S3ChangeDomainResponse struct {
	S3 *S3S3 `json:"s3,omitempty"`
	Error *S3ChangeDomainResponseError `json:"error,omitempty"`
}

// NewS3ChangeDomainResponse instantiates a new S3ChangeDomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3ChangeDomainResponse() *S3ChangeDomainResponse {
	this := S3ChangeDomainResponse{}
	return &this
}

// NewS3ChangeDomainResponseWithDefaults instantiates a new S3ChangeDomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ChangeDomainResponseWithDefaults() *S3ChangeDomainResponse {
	this := S3ChangeDomainResponse{}
	return &this
}

// GetS3 returns the S3 field value if set, zero value otherwise.
func (o *S3ChangeDomainResponse) GetS3() S3S3 {
	if o == nil || isNil(o.S3) {
		var ret S3S3
		return ret
	}
	return *o.S3
}

// GetS3Ok returns a tuple with the S3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ChangeDomainResponse) GetS3Ok() (*S3S3, bool) {
	if o == nil || isNil(o.S3) {
    return nil, false
	}
	return o.S3, true
}

// HasS3 returns a boolean if a field has been set.
func (o *S3ChangeDomainResponse) HasS3() bool {
	if o != nil && !isNil(o.S3) {
		return true
	}

	return false
}

// SetS3 gets a reference to the given S3S3 and assigns it to the S3 field.
func (o *S3ChangeDomainResponse) SetS3(v S3S3) {
	o.S3 = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *S3ChangeDomainResponse) GetError() S3ChangeDomainResponseError {
	if o == nil || isNil(o.Error) {
		var ret S3ChangeDomainResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ChangeDomainResponse) GetErrorOk() (*S3ChangeDomainResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *S3ChangeDomainResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given S3ChangeDomainResponseError and assigns it to the Error field.
func (o *S3ChangeDomainResponse) SetError(v S3ChangeDomainResponseError) {
	o.Error = &v
}

func (o S3ChangeDomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.S3) {
		toSerialize["s3"] = o.S3
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableS3ChangeDomainResponse struct {
	value *S3ChangeDomainResponse
	isSet bool
}

func (v NullableS3ChangeDomainResponse) Get() *S3ChangeDomainResponse {
	return v.value
}

func (v *NullableS3ChangeDomainResponse) Set(val *S3ChangeDomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableS3ChangeDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableS3ChangeDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3ChangeDomainResponse(val *S3ChangeDomainResponse) *NullableS3ChangeDomainResponse {
	return &NullableS3ChangeDomainResponse{value: val, isSet: true}
}

func (v NullableS3ChangeDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3ChangeDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


