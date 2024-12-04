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

// PostgresqlPgGetRoleListResponse struct for PostgresqlPgGetRoleListResponse
type PostgresqlPgGetRoleListResponse struct {
	Role []PostgresqlRole `json:"role,omitempty"`
}

// NewPostgresqlPgGetRoleListResponse instantiates a new PostgresqlPgGetRoleListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgGetRoleListResponse() *PostgresqlPgGetRoleListResponse {
	this := PostgresqlPgGetRoleListResponse{}
	return &this
}

// NewPostgresqlPgGetRoleListResponseWithDefaults instantiates a new PostgresqlPgGetRoleListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgGetRoleListResponseWithDefaults() *PostgresqlPgGetRoleListResponse {
	this := PostgresqlPgGetRoleListResponse{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *PostgresqlPgGetRoleListResponse) GetRole() []PostgresqlRole {
	if o == nil || isNil(o.Role) {
		var ret []PostgresqlRole
		return ret
	}
	return o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgGetRoleListResponse) GetRoleOk() ([]PostgresqlRole, bool) {
	if o == nil || isNil(o.Role) {
    return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *PostgresqlPgGetRoleListResponse) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given []PostgresqlRole and assigns it to the Role field.
func (o *PostgresqlPgGetRoleListResponse) SetRole(v []PostgresqlRole) {
	o.Role = v
}

func (o PostgresqlPgGetRoleListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgGetRoleListResponse struct {
	value *PostgresqlPgGetRoleListResponse
	isSet bool
}

func (v NullablePostgresqlPgGetRoleListResponse) Get() *PostgresqlPgGetRoleListResponse {
	return v.value
}

func (v *NullablePostgresqlPgGetRoleListResponse) Set(val *PostgresqlPgGetRoleListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgGetRoleListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgGetRoleListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgGetRoleListResponse(val *PostgresqlPgGetRoleListResponse) *NullablePostgresqlPgGetRoleListResponse {
	return &NullablePostgresqlPgGetRoleListResponse{value: val, isSet: true}
}

func (v NullablePostgresqlPgGetRoleListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgGetRoleListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


