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

// PostgresqlPgUpdateRoleRequest struct for PostgresqlPgUpdateRoleRequest
type PostgresqlPgUpdateRoleRequest struct {
	ServiceId *string `json:"service_id,omitempty"`
	RoleName *string `json:"role_name,omitempty"`
	Password *string `json:"password,omitempty"`
	ParentRole []string `json:"parent_role,omitempty"`
}

// NewPostgresqlPgUpdateRoleRequest instantiates a new PostgresqlPgUpdateRoleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgUpdateRoleRequest() *PostgresqlPgUpdateRoleRequest {
	this := PostgresqlPgUpdateRoleRequest{}
	return &this
}

// NewPostgresqlPgUpdateRoleRequestWithDefaults instantiates a new PostgresqlPgUpdateRoleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgUpdateRoleRequestWithDefaults() *PostgresqlPgUpdateRoleRequest {
	this := PostgresqlPgUpdateRoleRequest{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *PostgresqlPgUpdateRoleRequest) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgUpdateRoleRequest) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *PostgresqlPgUpdateRoleRequest) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *PostgresqlPgUpdateRoleRequest) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetRoleName returns the RoleName field value if set, zero value otherwise.
func (o *PostgresqlPgUpdateRoleRequest) GetRoleName() string {
	if o == nil || isNil(o.RoleName) {
		var ret string
		return ret
	}
	return *o.RoleName
}

// GetRoleNameOk returns a tuple with the RoleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgUpdateRoleRequest) GetRoleNameOk() (*string, bool) {
	if o == nil || isNil(o.RoleName) {
    return nil, false
	}
	return o.RoleName, true
}

// HasRoleName returns a boolean if a field has been set.
func (o *PostgresqlPgUpdateRoleRequest) HasRoleName() bool {
	if o != nil && !isNil(o.RoleName) {
		return true
	}

	return false
}

// SetRoleName gets a reference to the given string and assigns it to the RoleName field.
func (o *PostgresqlPgUpdateRoleRequest) SetRoleName(v string) {
	o.RoleName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PostgresqlPgUpdateRoleRequest) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgUpdateRoleRequest) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PostgresqlPgUpdateRoleRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PostgresqlPgUpdateRoleRequest) SetPassword(v string) {
	o.Password = &v
}

// GetParentRole returns the ParentRole field value if set, zero value otherwise.
func (o *PostgresqlPgUpdateRoleRequest) GetParentRole() []string {
	if o == nil || isNil(o.ParentRole) {
		var ret []string
		return ret
	}
	return o.ParentRole
}

// GetParentRoleOk returns a tuple with the ParentRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgUpdateRoleRequest) GetParentRoleOk() ([]string, bool) {
	if o == nil || isNil(o.ParentRole) {
    return nil, false
	}
	return o.ParentRole, true
}

// HasParentRole returns a boolean if a field has been set.
func (o *PostgresqlPgUpdateRoleRequest) HasParentRole() bool {
	if o != nil && !isNil(o.ParentRole) {
		return true
	}

	return false
}

// SetParentRole gets a reference to the given []string and assigns it to the ParentRole field.
func (o *PostgresqlPgUpdateRoleRequest) SetParentRole(v []string) {
	o.ParentRole = v
}

func (o PostgresqlPgUpdateRoleRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !isNil(o.RoleName) {
		toSerialize["role_name"] = o.RoleName
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.ParentRole) {
		toSerialize["parent_role"] = o.ParentRole
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgUpdateRoleRequest struct {
	value *PostgresqlPgUpdateRoleRequest
	isSet bool
}

func (v NullablePostgresqlPgUpdateRoleRequest) Get() *PostgresqlPgUpdateRoleRequest {
	return v.value
}

func (v *NullablePostgresqlPgUpdateRoleRequest) Set(val *PostgresqlPgUpdateRoleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgUpdateRoleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgUpdateRoleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgUpdateRoleRequest(val *PostgresqlPgUpdateRoleRequest) *NullablePostgresqlPgUpdateRoleRequest {
	return &NullablePostgresqlPgUpdateRoleRequest{value: val, isSet: true}
}

func (v NullablePostgresqlPgUpdateRoleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgUpdateRoleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

