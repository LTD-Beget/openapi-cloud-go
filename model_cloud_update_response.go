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

// CloudUpdateResponse struct for CloudUpdateResponse
type CloudUpdateResponse struct {
	Service *CloudService `json:"service,omitempty"`
	Error *CloudUpdateResponseError `json:"error,omitempty"`
}

// NewCloudUpdateResponse instantiates a new CloudUpdateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudUpdateResponse() *CloudUpdateResponse {
	this := CloudUpdateResponse{}
	return &this
}

// NewCloudUpdateResponseWithDefaults instantiates a new CloudUpdateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudUpdateResponseWithDefaults() *CloudUpdateResponse {
	this := CloudUpdateResponse{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CloudUpdateResponse) GetService() CloudService {
	if o == nil || isNil(o.Service) {
		var ret CloudService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudUpdateResponse) GetServiceOk() (*CloudService, bool) {
	if o == nil || isNil(o.Service) {
    return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CloudUpdateResponse) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given CloudService and assigns it to the Service field.
func (o *CloudUpdateResponse) SetService(v CloudService) {
	o.Service = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CloudUpdateResponse) GetError() CloudUpdateResponseError {
	if o == nil || isNil(o.Error) {
		var ret CloudUpdateResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudUpdateResponse) GetErrorOk() (*CloudUpdateResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CloudUpdateResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CloudUpdateResponseError and assigns it to the Error field.
func (o *CloudUpdateResponse) SetError(v CloudUpdateResponseError) {
	o.Error = &v
}

func (o CloudUpdateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableCloudUpdateResponse struct {
	value *CloudUpdateResponse
	isSet bool
}

func (v NullableCloudUpdateResponse) Get() *CloudUpdateResponse {
	return v.value
}

func (v *NullableCloudUpdateResponse) Set(val *CloudUpdateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudUpdateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudUpdateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudUpdateResponse(val *CloudUpdateResponse) *NullableCloudUpdateResponse {
	return &NullableCloudUpdateResponse{value: val, isSet: true}
}

func (v NullableCloudUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudUpdateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


