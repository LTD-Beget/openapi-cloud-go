/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// CloudChangeConfigurationRequest struct for CloudChangeConfigurationRequest
type CloudChangeConfigurationRequest struct {
	ServiceId *string `json:"service_id,omitempty"`
	ConfigurationId *string `json:"configuration_id,omitempty"`
}

// NewCloudChangeConfigurationRequest instantiates a new CloudChangeConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudChangeConfigurationRequest() *CloudChangeConfigurationRequest {
	this := CloudChangeConfigurationRequest{}
	return &this
}

// NewCloudChangeConfigurationRequestWithDefaults instantiates a new CloudChangeConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudChangeConfigurationRequestWithDefaults() *CloudChangeConfigurationRequest {
	this := CloudChangeConfigurationRequest{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *CloudChangeConfigurationRequest) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudChangeConfigurationRequest) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *CloudChangeConfigurationRequest) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *CloudChangeConfigurationRequest) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *CloudChangeConfigurationRequest) GetConfigurationId() string {
	if o == nil || isNil(o.ConfigurationId) {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudChangeConfigurationRequest) GetConfigurationIdOk() (*string, bool) {
	if o == nil || isNil(o.ConfigurationId) {
    return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *CloudChangeConfigurationRequest) HasConfigurationId() bool {
	if o != nil && !isNil(o.ConfigurationId) {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *CloudChangeConfigurationRequest) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

func (o CloudChangeConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !isNil(o.ConfigurationId) {
		toSerialize["configuration_id"] = o.ConfigurationId
	}
	return json.Marshal(toSerialize)
}

type NullableCloudChangeConfigurationRequest struct {
	value *CloudChangeConfigurationRequest
	isSet bool
}

func (v NullableCloudChangeConfigurationRequest) Get() *CloudChangeConfigurationRequest {
	return v.value
}

func (v *NullableCloudChangeConfigurationRequest) Set(val *CloudChangeConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudChangeConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudChangeConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudChangeConfigurationRequest(val *CloudChangeConfigurationRequest) *NullableCloudChangeConfigurationRequest {
	return &NullableCloudChangeConfigurationRequest{value: val, isSet: true}
}

func (v NullableCloudChangeConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudChangeConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

