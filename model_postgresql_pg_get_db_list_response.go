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

// PostgresqlPgGetDbListResponse struct for PostgresqlPgGetDbListResponse
type PostgresqlPgGetDbListResponse struct {
	Db []PostgresqlPgDb `json:"db,omitempty"`
}

// NewPostgresqlPgGetDbListResponse instantiates a new PostgresqlPgGetDbListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgGetDbListResponse() *PostgresqlPgGetDbListResponse {
	this := PostgresqlPgGetDbListResponse{}
	return &this
}

// NewPostgresqlPgGetDbListResponseWithDefaults instantiates a new PostgresqlPgGetDbListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgGetDbListResponseWithDefaults() *PostgresqlPgGetDbListResponse {
	this := PostgresqlPgGetDbListResponse{}
	return &this
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *PostgresqlPgGetDbListResponse) GetDb() []PostgresqlPgDb {
	if o == nil || isNil(o.Db) {
		var ret []PostgresqlPgDb
		return ret
	}
	return o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgGetDbListResponse) GetDbOk() ([]PostgresqlPgDb, bool) {
	if o == nil || isNil(o.Db) {
    return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *PostgresqlPgGetDbListResponse) HasDb() bool {
	if o != nil && !isNil(o.Db) {
		return true
	}

	return false
}

// SetDb gets a reference to the given []PostgresqlPgDb and assigns it to the Db field.
func (o *PostgresqlPgGetDbListResponse) SetDb(v []PostgresqlPgDb) {
	o.Db = v
}

func (o PostgresqlPgGetDbListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Db) {
		toSerialize["db"] = o.Db
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgGetDbListResponse struct {
	value *PostgresqlPgGetDbListResponse
	isSet bool
}

func (v NullablePostgresqlPgGetDbListResponse) Get() *PostgresqlPgGetDbListResponse {
	return v.value
}

func (v *NullablePostgresqlPgGetDbListResponse) Set(val *PostgresqlPgGetDbListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgGetDbListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgGetDbListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgGetDbListResponse(val *PostgresqlPgGetDbListResponse) *NullablePostgresqlPgGetDbListResponse {
	return &NullablePostgresqlPgGetDbListResponse{value: val, isSet: true}
}

func (v NullablePostgresqlPgGetDbListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgGetDbListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


