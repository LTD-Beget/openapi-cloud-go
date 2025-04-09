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

// PostgresqlPgCreateDbResponse struct for PostgresqlPgCreateDbResponse
type PostgresqlPgCreateDbResponse struct {
	Db *PostgresqlPgDb `json:"db,omitempty"`
	Error *PostgresqlPgCreateDbResponseError `json:"error,omitempty"`
}

// NewPostgresqlPgCreateDbResponse instantiates a new PostgresqlPgCreateDbResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgCreateDbResponse() *PostgresqlPgCreateDbResponse {
	this := PostgresqlPgCreateDbResponse{}
	return &this
}

// NewPostgresqlPgCreateDbResponseWithDefaults instantiates a new PostgresqlPgCreateDbResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgCreateDbResponseWithDefaults() *PostgresqlPgCreateDbResponse {
	this := PostgresqlPgCreateDbResponse{}
	return &this
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *PostgresqlPgCreateDbResponse) GetDb() PostgresqlPgDb {
	if o == nil || isNil(o.Db) {
		var ret PostgresqlPgDb
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgCreateDbResponse) GetDbOk() (*PostgresqlPgDb, bool) {
	if o == nil || isNil(o.Db) {
    return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *PostgresqlPgCreateDbResponse) HasDb() bool {
	if o != nil && !isNil(o.Db) {
		return true
	}

	return false
}

// SetDb gets a reference to the given PostgresqlPgDb and assigns it to the Db field.
func (o *PostgresqlPgCreateDbResponse) SetDb(v PostgresqlPgDb) {
	o.Db = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PostgresqlPgCreateDbResponse) GetError() PostgresqlPgCreateDbResponseError {
	if o == nil || isNil(o.Error) {
		var ret PostgresqlPgCreateDbResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgCreateDbResponse) GetErrorOk() (*PostgresqlPgCreateDbResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PostgresqlPgCreateDbResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given PostgresqlPgCreateDbResponseError and assigns it to the Error field.
func (o *PostgresqlPgCreateDbResponse) SetError(v PostgresqlPgCreateDbResponseError) {
	o.Error = &v
}

func (o PostgresqlPgCreateDbResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Db) {
		toSerialize["db"] = o.Db
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgCreateDbResponse struct {
	value *PostgresqlPgCreateDbResponse
	isSet bool
}

func (v NullablePostgresqlPgCreateDbResponse) Get() *PostgresqlPgCreateDbResponse {
	return v.value
}

func (v *NullablePostgresqlPgCreateDbResponse) Set(val *PostgresqlPgCreateDbResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgCreateDbResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgCreateDbResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgCreateDbResponse(val *PostgresqlPgCreateDbResponse) *NullablePostgresqlPgCreateDbResponse {
	return &NullablePostgresqlPgCreateDbResponse{value: val, isSet: true}
}

func (v NullablePostgresqlPgCreateDbResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgCreateDbResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


