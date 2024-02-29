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

// PostgresqlPgRemoteAccessState struct for PostgresqlPgRemoteAccessState
type PostgresqlPgRemoteAccessState struct {
	SubnetAddress []string `json:"subnet_address,omitempty"`
}

// NewPostgresqlPgRemoteAccessState instantiates a new PostgresqlPgRemoteAccessState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgRemoteAccessState() *PostgresqlPgRemoteAccessState {
	this := PostgresqlPgRemoteAccessState{}
	return &this
}

// NewPostgresqlPgRemoteAccessStateWithDefaults instantiates a new PostgresqlPgRemoteAccessState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgRemoteAccessStateWithDefaults() *PostgresqlPgRemoteAccessState {
	this := PostgresqlPgRemoteAccessState{}
	return &this
}

// GetSubnetAddress returns the SubnetAddress field value if set, zero value otherwise.
func (o *PostgresqlPgRemoteAccessState) GetSubnetAddress() []string {
	if o == nil || isNil(o.SubnetAddress) {
		var ret []string
		return ret
	}
	return o.SubnetAddress
}

// GetSubnetAddressOk returns a tuple with the SubnetAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgRemoteAccessState) GetSubnetAddressOk() ([]string, bool) {
	if o == nil || isNil(o.SubnetAddress) {
    return nil, false
	}
	return o.SubnetAddress, true
}

// HasSubnetAddress returns a boolean if a field has been set.
func (o *PostgresqlPgRemoteAccessState) HasSubnetAddress() bool {
	if o != nil && !isNil(o.SubnetAddress) {
		return true
	}

	return false
}

// SetSubnetAddress gets a reference to the given []string and assigns it to the SubnetAddress field.
func (o *PostgresqlPgRemoteAccessState) SetSubnetAddress(v []string) {
	o.SubnetAddress = v
}

func (o PostgresqlPgRemoteAccessState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubnetAddress) {
		toSerialize["subnet_address"] = o.SubnetAddress
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgRemoteAccessState struct {
	value *PostgresqlPgRemoteAccessState
	isSet bool
}

func (v NullablePostgresqlPgRemoteAccessState) Get() *PostgresqlPgRemoteAccessState {
	return v.value
}

func (v *NullablePostgresqlPgRemoteAccessState) Set(val *PostgresqlPgRemoteAccessState) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgRemoteAccessState) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgRemoteAccessState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgRemoteAccessState(val *PostgresqlPgRemoteAccessState) *NullablePostgresqlPgRemoteAccessState {
	return &NullablePostgresqlPgRemoteAccessState{value: val, isSet: true}
}

func (v NullablePostgresqlPgRemoteAccessState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgRemoteAccessState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


