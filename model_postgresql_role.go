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

// PostgresqlRole struct for PostgresqlRole
type PostgresqlRole struct {
	Name *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	ParentRole []string `json:"parent_role,omitempty"`
	Predefined *bool `json:"predefined,omitempty"`
}

// NewPostgresqlRole instantiates a new PostgresqlRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlRole() *PostgresqlRole {
	this := PostgresqlRole{}
	return &this
}

// NewPostgresqlRoleWithDefaults instantiates a new PostgresqlRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlRoleWithDefaults() *PostgresqlRole {
	this := PostgresqlRole{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostgresqlRole) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlRole) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostgresqlRole) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostgresqlRole) SetName(v string) {
	o.Name = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *PostgresqlRole) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlRole) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *PostgresqlRole) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *PostgresqlRole) SetPassword(v string) {
	o.Password = &v
}

// GetParentRole returns the ParentRole field value if set, zero value otherwise.
func (o *PostgresqlRole) GetParentRole() []string {
	if o == nil || isNil(o.ParentRole) {
		var ret []string
		return ret
	}
	return o.ParentRole
}

// GetParentRoleOk returns a tuple with the ParentRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlRole) GetParentRoleOk() ([]string, bool) {
	if o == nil || isNil(o.ParentRole) {
    return nil, false
	}
	return o.ParentRole, true
}

// HasParentRole returns a boolean if a field has been set.
func (o *PostgresqlRole) HasParentRole() bool {
	if o != nil && !isNil(o.ParentRole) {
		return true
	}

	return false
}

// SetParentRole gets a reference to the given []string and assigns it to the ParentRole field.
func (o *PostgresqlRole) SetParentRole(v []string) {
	o.ParentRole = v
}

// GetPredefined returns the Predefined field value if set, zero value otherwise.
func (o *PostgresqlRole) GetPredefined() bool {
	if o == nil || isNil(o.Predefined) {
		var ret bool
		return ret
	}
	return *o.Predefined
}

// GetPredefinedOk returns a tuple with the Predefined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlRole) GetPredefinedOk() (*bool, bool) {
	if o == nil || isNil(o.Predefined) {
    return nil, false
	}
	return o.Predefined, true
}

// HasPredefined returns a boolean if a field has been set.
func (o *PostgresqlRole) HasPredefined() bool {
	if o != nil && !isNil(o.Predefined) {
		return true
	}

	return false
}

// SetPredefined gets a reference to the given bool and assigns it to the Predefined field.
func (o *PostgresqlRole) SetPredefined(v bool) {
	o.Predefined = &v
}

func (o PostgresqlRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.ParentRole) {
		toSerialize["parent_role"] = o.ParentRole
	}
	if !isNil(o.Predefined) {
		toSerialize["predefined"] = o.Predefined
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlRole struct {
	value *PostgresqlRole
	isSet bool
}

func (v NullablePostgresqlRole) Get() *PostgresqlRole {
	return v.value
}

func (v *NullablePostgresqlRole) Set(val *PostgresqlRole) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlRole) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlRole(val *PostgresqlRole) *NullablePostgresqlRole {
	return &NullablePostgresqlRole{value: val, isSet: true}
}

func (v NullablePostgresqlRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


