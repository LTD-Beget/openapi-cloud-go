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

// S3GetPrefixResponse struct for S3GetPrefixResponse
type S3GetPrefixResponse struct {
	Prefix *string `json:"prefix,omitempty"`
}

// NewS3GetPrefixResponse instantiates a new S3GetPrefixResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3GetPrefixResponse() *S3GetPrefixResponse {
	this := S3GetPrefixResponse{}
	return &this
}

// NewS3GetPrefixResponseWithDefaults instantiates a new S3GetPrefixResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3GetPrefixResponseWithDefaults() *S3GetPrefixResponse {
	this := S3GetPrefixResponse{}
	return &this
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *S3GetPrefixResponse) GetPrefix() string {
	if o == nil || isNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3GetPrefixResponse) GetPrefixOk() (*string, bool) {
	if o == nil || isNil(o.Prefix) {
    return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *S3GetPrefixResponse) HasPrefix() bool {
	if o != nil && !isNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *S3GetPrefixResponse) SetPrefix(v string) {
	o.Prefix = &v
}

func (o S3GetPrefixResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableS3GetPrefixResponse struct {
	value *S3GetPrefixResponse
	isSet bool
}

func (v NullableS3GetPrefixResponse) Get() *S3GetPrefixResponse {
	return v.value
}

func (v *NullableS3GetPrefixResponse) Set(val *S3GetPrefixResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableS3GetPrefixResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableS3GetPrefixResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3GetPrefixResponse(val *S3GetPrefixResponse) *NullableS3GetPrefixResponse {
	return &NullableS3GetPrefixResponse{value: val, isSet: true}
}

func (v NullableS3GetPrefixResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3GetPrefixResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


