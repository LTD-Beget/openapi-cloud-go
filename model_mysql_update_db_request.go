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

// MysqlUpdateDbRequest struct for MysqlUpdateDbRequest
type MysqlUpdateDbRequest struct {
	ServiceId *string `json:"service_id,omitempty"`
	DbName *string `json:"db_name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewMysqlUpdateDbRequest instantiates a new MysqlUpdateDbRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlUpdateDbRequest() *MysqlUpdateDbRequest {
	this := MysqlUpdateDbRequest{}
	return &this
}

// NewMysqlUpdateDbRequestWithDefaults instantiates a new MysqlUpdateDbRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlUpdateDbRequestWithDefaults() *MysqlUpdateDbRequest {
	this := MysqlUpdateDbRequest{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *MysqlUpdateDbRequest) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlUpdateDbRequest) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *MysqlUpdateDbRequest) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *MysqlUpdateDbRequest) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *MysqlUpdateDbRequest) GetDbName() string {
	if o == nil || isNil(o.DbName) {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlUpdateDbRequest) GetDbNameOk() (*string, bool) {
	if o == nil || isNil(o.DbName) {
    return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *MysqlUpdateDbRequest) HasDbName() bool {
	if o != nil && !isNil(o.DbName) {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *MysqlUpdateDbRequest) SetDbName(v string) {
	o.DbName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *MysqlUpdateDbRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlUpdateDbRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *MysqlUpdateDbRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *MysqlUpdateDbRequest) SetDescription(v string) {
	o.Description = &v
}

func (o MysqlUpdateDbRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !isNil(o.DbName) {
		toSerialize["db_name"] = o.DbName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlUpdateDbRequest struct {
	value *MysqlUpdateDbRequest
	isSet bool
}

func (v NullableMysqlUpdateDbRequest) Get() *MysqlUpdateDbRequest {
	return v.value
}

func (v *NullableMysqlUpdateDbRequest) Set(val *MysqlUpdateDbRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlUpdateDbRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlUpdateDbRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlUpdateDbRequest(val *MysqlUpdateDbRequest) *NullableMysqlUpdateDbRequest {
	return &NullableMysqlUpdateDbRequest{value: val, isSet: true}
}

func (v NullableMysqlUpdateDbRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlUpdateDbRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

