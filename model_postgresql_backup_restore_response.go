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

// PostgresqlBackupRestoreResponse struct for PostgresqlBackupRestoreResponse
type PostgresqlBackupRestoreResponse struct {
	Order *StructuresPostgresqlRestoreOrder `json:"order,omitempty"`
	Error *PostgresqlBackupRestoreResponseError `json:"error,omitempty"`
}

// NewPostgresqlBackupRestoreResponse instantiates a new PostgresqlBackupRestoreResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlBackupRestoreResponse() *PostgresqlBackupRestoreResponse {
	this := PostgresqlBackupRestoreResponse{}
	return &this
}

// NewPostgresqlBackupRestoreResponseWithDefaults instantiates a new PostgresqlBackupRestoreResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlBackupRestoreResponseWithDefaults() *PostgresqlBackupRestoreResponse {
	this := PostgresqlBackupRestoreResponse{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *PostgresqlBackupRestoreResponse) GetOrder() StructuresPostgresqlRestoreOrder {
	if o == nil || isNil(o.Order) {
		var ret StructuresPostgresqlRestoreOrder
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlBackupRestoreResponse) GetOrderOk() (*StructuresPostgresqlRestoreOrder, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *PostgresqlBackupRestoreResponse) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given StructuresPostgresqlRestoreOrder and assigns it to the Order field.
func (o *PostgresqlBackupRestoreResponse) SetOrder(v StructuresPostgresqlRestoreOrder) {
	o.Order = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *PostgresqlBackupRestoreResponse) GetError() PostgresqlBackupRestoreResponseError {
	if o == nil || isNil(o.Error) {
		var ret PostgresqlBackupRestoreResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlBackupRestoreResponse) GetErrorOk() (*PostgresqlBackupRestoreResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *PostgresqlBackupRestoreResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given PostgresqlBackupRestoreResponseError and assigns it to the Error field.
func (o *PostgresqlBackupRestoreResponse) SetError(v PostgresqlBackupRestoreResponseError) {
	o.Error = &v
}

func (o PostgresqlBackupRestoreResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlBackupRestoreResponse struct {
	value *PostgresqlBackupRestoreResponse
	isSet bool
}

func (v NullablePostgresqlBackupRestoreResponse) Get() *PostgresqlBackupRestoreResponse {
	return v.value
}

func (v *NullablePostgresqlBackupRestoreResponse) Set(val *PostgresqlBackupRestoreResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlBackupRestoreResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlBackupRestoreResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlBackupRestoreResponse(val *PostgresqlBackupRestoreResponse) *NullablePostgresqlBackupRestoreResponse {
	return &NullablePostgresqlBackupRestoreResponse{value: val, isSet: true}
}

func (v NullablePostgresqlBackupRestoreResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlBackupRestoreResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

