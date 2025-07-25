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

// PostgresqlPgGetConfigResponse struct for PostgresqlPgGetConfigResponse
type PostgresqlPgGetConfigResponse struct {
	Config *PostgresqlPgConfigInfo `json:"config,omitempty"`
}

// NewPostgresqlPgGetConfigResponse instantiates a new PostgresqlPgGetConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgGetConfigResponse() *PostgresqlPgGetConfigResponse {
	this := PostgresqlPgGetConfigResponse{}
	return &this
}

// NewPostgresqlPgGetConfigResponseWithDefaults instantiates a new PostgresqlPgGetConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgGetConfigResponseWithDefaults() *PostgresqlPgGetConfigResponse {
	this := PostgresqlPgGetConfigResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *PostgresqlPgGetConfigResponse) GetConfig() PostgresqlPgConfigInfo {
	if o == nil || isNil(o.Config) {
		var ret PostgresqlPgConfigInfo
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgGetConfigResponse) GetConfigOk() (*PostgresqlPgConfigInfo, bool) {
	if o == nil || isNil(o.Config) {
    return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *PostgresqlPgGetConfigResponse) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given PostgresqlPgConfigInfo and assigns it to the Config field.
func (o *PostgresqlPgGetConfigResponse) SetConfig(v PostgresqlPgConfigInfo) {
	o.Config = &v
}

func (o PostgresqlPgGetConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgGetConfigResponse struct {
	value *PostgresqlPgGetConfigResponse
	isSet bool
}

func (v NullablePostgresqlPgGetConfigResponse) Get() *PostgresqlPgGetConfigResponse {
	return v.value
}

func (v *NullablePostgresqlPgGetConfigResponse) Set(val *PostgresqlPgGetConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgGetConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgGetConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgGetConfigResponse(val *PostgresqlPgGetConfigResponse) *NullablePostgresqlPgGetConfigResponse {
	return &NullablePostgresqlPgGetConfigResponse{value: val, isSet: true}
}

func (v NullablePostgresqlPgGetConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgGetConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


