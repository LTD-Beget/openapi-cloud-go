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

// S3StatisticGetQuotaResponse struct for S3StatisticGetQuotaResponse
type S3StatisticGetQuotaResponse struct {
	Size *StructuresStatisticSeriesData `json:"size,omitempty"`
}

// NewS3StatisticGetQuotaResponse instantiates a new S3StatisticGetQuotaResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3StatisticGetQuotaResponse() *S3StatisticGetQuotaResponse {
	this := S3StatisticGetQuotaResponse{}
	return &this
}

// NewS3StatisticGetQuotaResponseWithDefaults instantiates a new S3StatisticGetQuotaResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3StatisticGetQuotaResponseWithDefaults() *S3StatisticGetQuotaResponse {
	this := S3StatisticGetQuotaResponse{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *S3StatisticGetQuotaResponse) GetSize() StructuresStatisticSeriesData {
	if o == nil || isNil(o.Size) {
		var ret StructuresStatisticSeriesData
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3StatisticGetQuotaResponse) GetSizeOk() (*StructuresStatisticSeriesData, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *S3StatisticGetQuotaResponse) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given StructuresStatisticSeriesData and assigns it to the Size field.
func (o *S3StatisticGetQuotaResponse) SetSize(v StructuresStatisticSeriesData) {
	o.Size = &v
}

func (o S3StatisticGetQuotaResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return json.Marshal(toSerialize)
}

type NullableS3StatisticGetQuotaResponse struct {
	value *S3StatisticGetQuotaResponse
	isSet bool
}

func (v NullableS3StatisticGetQuotaResponse) Get() *S3StatisticGetQuotaResponse {
	return v.value
}

func (v *NullableS3StatisticGetQuotaResponse) Set(val *S3StatisticGetQuotaResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableS3StatisticGetQuotaResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableS3StatisticGetQuotaResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3StatisticGetQuotaResponse(val *S3StatisticGetQuotaResponse) *NullableS3StatisticGetQuotaResponse {
	return &NullableS3StatisticGetQuotaResponse{value: val, isSet: true}
}

func (v NullableS3StatisticGetQuotaResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3StatisticGetQuotaResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


