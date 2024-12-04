/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// MysqlBackupRestoreRequest struct for MysqlBackupRestoreRequest
type MysqlBackupRestoreRequest struct {
	CopyId *string `json:"copy_id,omitempty"`
	ServiceUuid *string `json:"service_uuid,omitempty"`
}

// NewMysqlBackupRestoreRequest instantiates a new MysqlBackupRestoreRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlBackupRestoreRequest() *MysqlBackupRestoreRequest {
	this := MysqlBackupRestoreRequest{}
	return &this
}

// NewMysqlBackupRestoreRequestWithDefaults instantiates a new MysqlBackupRestoreRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlBackupRestoreRequestWithDefaults() *MysqlBackupRestoreRequest {
	this := MysqlBackupRestoreRequest{}
	return &this
}

// GetCopyId returns the CopyId field value if set, zero value otherwise.
func (o *MysqlBackupRestoreRequest) GetCopyId() string {
	if o == nil || isNil(o.CopyId) {
		var ret string
		return ret
	}
	return *o.CopyId
}

// GetCopyIdOk returns a tuple with the CopyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlBackupRestoreRequest) GetCopyIdOk() (*string, bool) {
	if o == nil || isNil(o.CopyId) {
    return nil, false
	}
	return o.CopyId, true
}

// HasCopyId returns a boolean if a field has been set.
func (o *MysqlBackupRestoreRequest) HasCopyId() bool {
	if o != nil && !isNil(o.CopyId) {
		return true
	}

	return false
}

// SetCopyId gets a reference to the given string and assigns it to the CopyId field.
func (o *MysqlBackupRestoreRequest) SetCopyId(v string) {
	o.CopyId = &v
}

// GetServiceUuid returns the ServiceUuid field value if set, zero value otherwise.
func (o *MysqlBackupRestoreRequest) GetServiceUuid() string {
	if o == nil || isNil(o.ServiceUuid) {
		var ret string
		return ret
	}
	return *o.ServiceUuid
}

// GetServiceUuidOk returns a tuple with the ServiceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlBackupRestoreRequest) GetServiceUuidOk() (*string, bool) {
	if o == nil || isNil(o.ServiceUuid) {
    return nil, false
	}
	return o.ServiceUuid, true
}

// HasServiceUuid returns a boolean if a field has been set.
func (o *MysqlBackupRestoreRequest) HasServiceUuid() bool {
	if o != nil && !isNil(o.ServiceUuid) {
		return true
	}

	return false
}

// SetServiceUuid gets a reference to the given string and assigns it to the ServiceUuid field.
func (o *MysqlBackupRestoreRequest) SetServiceUuid(v string) {
	o.ServiceUuid = &v
}

func (o MysqlBackupRestoreRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CopyId) {
		toSerialize["copy_id"] = o.CopyId
	}
	if !isNil(o.ServiceUuid) {
		toSerialize["service_uuid"] = o.ServiceUuid
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlBackupRestoreRequest struct {
	value *MysqlBackupRestoreRequest
	isSet bool
}

func (v NullableMysqlBackupRestoreRequest) Get() *MysqlBackupRestoreRequest {
	return v.value
}

func (v *NullableMysqlBackupRestoreRequest) Set(val *MysqlBackupRestoreRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlBackupRestoreRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlBackupRestoreRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlBackupRestoreRequest(val *MysqlBackupRestoreRequest) *NullableMysqlBackupRestoreRequest {
	return &NullableMysqlBackupRestoreRequest{value: val, isSet: true}
}

func (v NullableMysqlBackupRestoreRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlBackupRestoreRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


