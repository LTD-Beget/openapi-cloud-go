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

// PostgresqlBackupGetListResponse struct for PostgresqlBackupGetListResponse
type PostgresqlBackupGetListResponse struct {
	Copy []StructuresPostgresqlCopy `json:"copy,omitempty"`
}

// NewPostgresqlBackupGetListResponse instantiates a new PostgresqlBackupGetListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlBackupGetListResponse() *PostgresqlBackupGetListResponse {
	this := PostgresqlBackupGetListResponse{}
	return &this
}

// NewPostgresqlBackupGetListResponseWithDefaults instantiates a new PostgresqlBackupGetListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlBackupGetListResponseWithDefaults() *PostgresqlBackupGetListResponse {
	this := PostgresqlBackupGetListResponse{}
	return &this
}

// GetCopy returns the Copy field value if set, zero value otherwise.
func (o *PostgresqlBackupGetListResponse) GetCopy() []StructuresPostgresqlCopy {
	if o == nil || isNil(o.Copy) {
		var ret []StructuresPostgresqlCopy
		return ret
	}
	return o.Copy
}

// GetCopyOk returns a tuple with the Copy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlBackupGetListResponse) GetCopyOk() ([]StructuresPostgresqlCopy, bool) {
	if o == nil || isNil(o.Copy) {
    return nil, false
	}
	return o.Copy, true
}

// HasCopy returns a boolean if a field has been set.
func (o *PostgresqlBackupGetListResponse) HasCopy() bool {
	if o != nil && !isNil(o.Copy) {
		return true
	}

	return false
}

// SetCopy gets a reference to the given []StructuresPostgresqlCopy and assigns it to the Copy field.
func (o *PostgresqlBackupGetListResponse) SetCopy(v []StructuresPostgresqlCopy) {
	o.Copy = v
}

func (o PostgresqlBackupGetListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Copy) {
		toSerialize["copy"] = o.Copy
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlBackupGetListResponse struct {
	value *PostgresqlBackupGetListResponse
	isSet bool
}

func (v NullablePostgresqlBackupGetListResponse) Get() *PostgresqlBackupGetListResponse {
	return v.value
}

func (v *NullablePostgresqlBackupGetListResponse) Set(val *PostgresqlBackupGetListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlBackupGetListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlBackupGetListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlBackupGetListResponse(val *PostgresqlBackupGetListResponse) *NullablePostgresqlBackupGetListResponse {
	return &NullablePostgresqlBackupGetListResponse{value: val, isSet: true}
}

func (v NullablePostgresqlBackupGetListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlBackupGetListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

