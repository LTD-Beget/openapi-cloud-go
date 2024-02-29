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

// PostgresqlPgRemoveDbResponse struct for PostgresqlPgRemoveDbResponse
type PostgresqlPgRemoveDbResponse struct {
	Success map[string]interface{} `json:"success,omitempty"`
	Error *PostgresqlPgRemoveDbResponseError `json:"error,omitempty"`
}

// NewPostgresqlPgRemoveDbResponse instantiates a new PostgresqlPgRemoveDbResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgRemoveDbResponse() *PostgresqlPgRemoveDbResponse {
	this := PostgresqlPgRemoveDbResponse{}
	return &this
}

// NewPostgresqlPgRemoveDbResponseWithDefaults instantiates a new PostgresqlPgRemoveDbResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgRemoveDbResponseWithDefaults() *PostgresqlPgRemoveDbResponse {
	this := PostgresqlPgRemoveDbResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *PostgresqlPgRemoveDbResponse) GetSuccess() map[string]interface{} {
	if o == nil || isNil(o.Success) {
		var ret map[string]interface{}
		return ret
	}
	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgRemoveDbResponse) GetSuccessOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Success) {
    return map[string]interface{}{}, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *PostgresqlPgRemoveDbResponse) HasSuccess() bool {
	if o != nil && !isNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given map[string]interface{} and assigns it to the Success field.
func (o *PostgresqlPgRemoveDbResponse) SetSuccess(v map[string]interface{}) {
	o.Success = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PostgresqlPgRemoveDbResponse) GetError() PostgresqlPgRemoveDbResponseError {
	if o == nil || isNil(o.Error) {
		var ret PostgresqlPgRemoveDbResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgRemoveDbResponse) GetErrorOk() (*PostgresqlPgRemoveDbResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PostgresqlPgRemoveDbResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given PostgresqlPgRemoveDbResponseError and assigns it to the Error field.
func (o *PostgresqlPgRemoveDbResponse) SetError(v PostgresqlPgRemoveDbResponseError) {
	o.Error = &v
}

func (o PostgresqlPgRemoveDbResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgRemoveDbResponse struct {
	value *PostgresqlPgRemoveDbResponse
	isSet bool
}

func (v NullablePostgresqlPgRemoveDbResponse) Get() *PostgresqlPgRemoveDbResponse {
	return v.value
}

func (v *NullablePostgresqlPgRemoveDbResponse) Set(val *PostgresqlPgRemoveDbResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgRemoveDbResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgRemoveDbResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgRemoveDbResponse(val *PostgresqlPgRemoveDbResponse) *NullablePostgresqlPgRemoveDbResponse {
	return &NullablePostgresqlPgRemoveDbResponse{value: val, isSet: true}
}

func (v NullablePostgresqlPgRemoveDbResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgRemoveDbResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


