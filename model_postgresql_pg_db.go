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

// PostgresqlPgDb struct for PostgresqlPgDb
type PostgresqlPgDb struct {
	Name *string `json:"name,omitempty"`
	Size *string `json:"size,omitempty"`
	Owner *string `json:"owner,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewPostgresqlPgDb instantiates a new PostgresqlPgDb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlPgDb() *PostgresqlPgDb {
	this := PostgresqlPgDb{}
	return &this
}

// NewPostgresqlPgDbWithDefaults instantiates a new PostgresqlPgDb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlPgDbWithDefaults() *PostgresqlPgDb {
	this := PostgresqlPgDb{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PostgresqlPgDb) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgDb) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PostgresqlPgDb) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PostgresqlPgDb) SetName(v string) {
	o.Name = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *PostgresqlPgDb) GetSize() string {
	if o == nil || isNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgDb) GetSizeOk() (*string, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *PostgresqlPgDb) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *PostgresqlPgDb) SetSize(v string) {
	o.Size = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *PostgresqlPgDb) GetOwner() string {
	if o == nil || isNil(o.Owner) {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgDb) GetOwnerOk() (*string, bool) {
	if o == nil || isNil(o.Owner) {
    return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *PostgresqlPgDb) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *PostgresqlPgDb) SetOwner(v string) {
	o.Owner = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PostgresqlPgDb) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlPgDb) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PostgresqlPgDb) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PostgresqlPgDb) SetDescription(v string) {
	o.Description = &v
}

func (o PostgresqlPgDb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullablePostgresqlPgDb struct {
	value *PostgresqlPgDb
	isSet bool
}

func (v NullablePostgresqlPgDb) Get() *PostgresqlPgDb {
	return v.value
}

func (v *NullablePostgresqlPgDb) Set(val *PostgresqlPgDb) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlPgDb) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlPgDb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlPgDb(val *PostgresqlPgDb) *NullablePostgresqlPgDb {
	return &NullablePostgresqlPgDb{value: val, isSet: true}
}

func (v NullablePostgresqlPgDb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlPgDb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


