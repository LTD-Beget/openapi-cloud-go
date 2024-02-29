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

// MysqlChangeAccessPasswordRequest struct for MysqlChangeAccessPasswordRequest
type MysqlChangeAccessPasswordRequest struct {
	ServiceId *string `json:"service_id,omitempty"`
	DbName *string `json:"db_name,omitempty"`
	Host *string `json:"host,omitempty"`
	Password *string `json:"password,omitempty"`
	SavePmaPassword *bool `json:"save_pma_password,omitempty"`
}

// NewMysqlChangeAccessPasswordRequest instantiates a new MysqlChangeAccessPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlChangeAccessPasswordRequest() *MysqlChangeAccessPasswordRequest {
	this := MysqlChangeAccessPasswordRequest{}
	return &this
}

// NewMysqlChangeAccessPasswordRequestWithDefaults instantiates a new MysqlChangeAccessPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlChangeAccessPasswordRequestWithDefaults() *MysqlChangeAccessPasswordRequest {
	this := MysqlChangeAccessPasswordRequest{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *MysqlChangeAccessPasswordRequest) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlChangeAccessPasswordRequest) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *MysqlChangeAccessPasswordRequest) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *MysqlChangeAccessPasswordRequest) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *MysqlChangeAccessPasswordRequest) GetDbName() string {
	if o == nil || isNil(o.DbName) {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlChangeAccessPasswordRequest) GetDbNameOk() (*string, bool) {
	if o == nil || isNil(o.DbName) {
    return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *MysqlChangeAccessPasswordRequest) HasDbName() bool {
	if o != nil && !isNil(o.DbName) {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *MysqlChangeAccessPasswordRequest) SetDbName(v string) {
	o.DbName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *MysqlChangeAccessPasswordRequest) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlChangeAccessPasswordRequest) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *MysqlChangeAccessPasswordRequest) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *MysqlChangeAccessPasswordRequest) SetHost(v string) {
	o.Host = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *MysqlChangeAccessPasswordRequest) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlChangeAccessPasswordRequest) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *MysqlChangeAccessPasswordRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *MysqlChangeAccessPasswordRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSavePmaPassword returns the SavePmaPassword field value if set, zero value otherwise.
func (o *MysqlChangeAccessPasswordRequest) GetSavePmaPassword() bool {
	if o == nil || isNil(o.SavePmaPassword) {
		var ret bool
		return ret
	}
	return *o.SavePmaPassword
}

// GetSavePmaPasswordOk returns a tuple with the SavePmaPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlChangeAccessPasswordRequest) GetSavePmaPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.SavePmaPassword) {
    return nil, false
	}
	return o.SavePmaPassword, true
}

// HasSavePmaPassword returns a boolean if a field has been set.
func (o *MysqlChangeAccessPasswordRequest) HasSavePmaPassword() bool {
	if o != nil && !isNil(o.SavePmaPassword) {
		return true
	}

	return false
}

// SetSavePmaPassword gets a reference to the given bool and assigns it to the SavePmaPassword field.
func (o *MysqlChangeAccessPasswordRequest) SetSavePmaPassword(v bool) {
	o.SavePmaPassword = &v
}

func (o MysqlChangeAccessPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServiceId) {
		toSerialize["service_id"] = o.ServiceId
	}
	if !isNil(o.DbName) {
		toSerialize["db_name"] = o.DbName
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.SavePmaPassword) {
		toSerialize["save_pma_password"] = o.SavePmaPassword
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlChangeAccessPasswordRequest struct {
	value *MysqlChangeAccessPasswordRequest
	isSet bool
}

func (v NullableMysqlChangeAccessPasswordRequest) Get() *MysqlChangeAccessPasswordRequest {
	return v.value
}

func (v *NullableMysqlChangeAccessPasswordRequest) Set(val *MysqlChangeAccessPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlChangeAccessPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlChangeAccessPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlChangeAccessPasswordRequest(val *MysqlChangeAccessPasswordRequest) *NullableMysqlChangeAccessPasswordRequest {
	return &NullableMysqlChangeAccessPasswordRequest{value: val, isSet: true}
}

func (v NullableMysqlChangeAccessPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlChangeAccessPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


