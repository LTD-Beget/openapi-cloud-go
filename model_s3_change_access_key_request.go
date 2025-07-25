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

// S3ChangeAccessKeyRequest struct for S3ChangeAccessKeyRequest
type S3ChangeAccessKeyRequest struct {
	ServiceId *string `json:"service_id,omitempty"`
}

// NewS3ChangeAccessKeyRequest instantiates a new S3ChangeAccessKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3ChangeAccessKeyRequest() *S3ChangeAccessKeyRequest {
	this := S3ChangeAccessKeyRequest{}
	return &this
}

// NewS3ChangeAccessKeyRequestWithDefaults instantiates a new S3ChangeAccessKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3ChangeAccessKeyRequestWithDefaults() *S3ChangeAccessKeyRequest {
	this := S3ChangeAccessKeyRequest{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *S3ChangeAccessKeyRequest) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3ChangeAccessKeyRequest) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *S3ChangeAccessKeyRequest) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *S3ChangeAccessKeyRequest) SetServiceId(v string) {
	o.ServiceId = &v
}

func (o S3ChangeAccessKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	return json.Marshal(toSerialize)
}

type NullableS3ChangeAccessKeyRequest struct {
	value *S3ChangeAccessKeyRequest
	isSet bool
}

func (v NullableS3ChangeAccessKeyRequest) Get() *S3ChangeAccessKeyRequest {
	return v.value
}

func (v *NullableS3ChangeAccessKeyRequest) Set(val *S3ChangeAccessKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableS3ChangeAccessKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableS3ChangeAccessKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3ChangeAccessKeyRequest(val *S3ChangeAccessKeyRequest) *NullableS3ChangeAccessKeyRequest {
	return &NullableS3ChangeAccessKeyRequest{value: val, isSet: true}
}

func (v NullableS3ChangeAccessKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3ChangeAccessKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


