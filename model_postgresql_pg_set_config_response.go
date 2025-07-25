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

// PostgresqlPgSetConfigResponse struct for PostgresqlPgSetConfigResponse
type PostgresqlPgSetConfigResponse struct {
	Config *PostgresqlPgConfigInfo `json:"config,omitempty"`
	Error *PostgresqlPgSetConfigResponseError `json:"error,omitempty"`
}

// NewPostgresqlPgSetConfigResponse instantiates a new PostgresqlPgSetConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgSetConfigResponse() *PostgresqlPgSetConfigResponse {
	this := PostgresqlPgSetConfigResponse{}
	return &this
}

// NewPostgresqlPgSetConfigResponseWithDefaults instantiates a new PostgresqlPgSetConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgSetConfigResponseWithDefaults() *PostgresqlPgSetConfigResponse {
	this := PostgresqlPgSetConfigResponse{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *PostgresqlPgSetConfigResponse) GetConfig() PostgresqlPgConfigInfo {
	if o == nil || isNil(o.Config) {
		var ret PostgresqlPgConfigInfo
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgSetConfigResponse) GetConfigOk() (*PostgresqlPgConfigInfo, bool) {
	if o == nil || isNil(o.Config) {
    return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *PostgresqlPgSetConfigResponse) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given PostgresqlPgConfigInfo and assigns it to the Config field.
func (o *PostgresqlPgSetConfigResponse) SetConfig(v PostgresqlPgConfigInfo) {
	o.Config = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PostgresqlPgSetConfigResponse) GetError() PostgresqlPgSetConfigResponseError {
	if o == nil || isNil(o.Error) {
		var ret PostgresqlPgSetConfigResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgSetConfigResponse) GetErrorOk() (*PostgresqlPgSetConfigResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PostgresqlPgSetConfigResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given PostgresqlPgSetConfigResponseError and assigns it to the Error field.
func (o *PostgresqlPgSetConfigResponse) SetError(v PostgresqlPgSetConfigResponseError) {
	o.Error = &v
}

func (o PostgresqlPgSetConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgSetConfigResponse struct {
	value *PostgresqlPgSetConfigResponse
	isSet bool
}

func (v NullablePostgresqlPgSetConfigResponse) Get() *PostgresqlPgSetConfigResponse {
	return v.value
}

func (v *NullablePostgresqlPgSetConfigResponse) Set(val *PostgresqlPgSetConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgSetConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgSetConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgSetConfigResponse(val *PostgresqlPgSetConfigResponse) *NullablePostgresqlPgSetConfigResponse {
	return &NullablePostgresqlPgSetConfigResponse{value: val, isSet: true}
}

func (v NullablePostgresqlPgSetConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgSetConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


