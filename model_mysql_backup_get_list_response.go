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

// MysqlBackupGetListResponse struct for MysqlBackupGetListResponse
type MysqlBackupGetListResponse struct {
	Copy []StructuresMysqlCopy `json:"copy,omitempty"`
}

// NewMysqlBackupGetListResponse instantiates a new MysqlBackupGetListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlBackupGetListResponse() *MysqlBackupGetListResponse {
	this := MysqlBackupGetListResponse{}
	return &this
}

// NewMysqlBackupGetListResponseWithDefaults instantiates a new MysqlBackupGetListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlBackupGetListResponseWithDefaults() *MysqlBackupGetListResponse {
	this := MysqlBackupGetListResponse{}
	return &this
}

// GetCopy returns the Copy field value if set, zero value otherwise.
func (o *MysqlBackupGetListResponse) GetCopy() []StructuresMysqlCopy {
	if o == nil || isNil(o.Copy) {
		var ret []StructuresMysqlCopy
		return ret
	}
	return o.Copy
}

// GetCopyOk returns a tuple with the Copy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlBackupGetListResponse) GetCopyOk() ([]StructuresMysqlCopy, bool) {
	if o == nil || isNil(o.Copy) {
    return nil, false
	}
	return o.Copy, true
}

// HasCopy returns a boolean if a field has been set.
func (o *MysqlBackupGetListResponse) HasCopy() bool {
	if o != nil && !isNil(o.Copy) {
		return true
	}

	return false
}

// SetCopy gets a reference to the given []StructuresMysqlCopy and assigns it to the Copy field.
func (o *MysqlBackupGetListResponse) SetCopy(v []StructuresMysqlCopy) {
	o.Copy = v
}

func (o MysqlBackupGetListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Copy) {
		toSerialize["copy"] = o.Copy
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlBackupGetListResponse struct {
	value *MysqlBackupGetListResponse
	isSet bool
}

func (v NullableMysqlBackupGetListResponse) Get() *MysqlBackupGetListResponse {
	return v.value
}

func (v *NullableMysqlBackupGetListResponse) Set(val *MysqlBackupGetListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlBackupGetListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlBackupGetListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlBackupGetListResponse(val *MysqlBackupGetListResponse) *NullableMysqlBackupGetListResponse {
	return &NullableMysqlBackupGetListResponse{value: val, isSet: true}
}

func (v NullableMysqlBackupGetListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlBackupGetListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


