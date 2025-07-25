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

// MysqlCreateAccessRequest struct for MysqlCreateAccessRequest
type MysqlCreateAccessRequest struct {
	ServiceId *string `json:"service_id,omitempty"`
	DbName *string `json:"db_name,omitempty"`
	Host *string `json:"host,omitempty"`
	Password *string `json:"password,omitempty"`
	SavePmaPassword *bool `json:"save_pma_password,omitempty"`
}

// NewMysqlCreateAccessRequest instantiates a new MysqlCreateAccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlCreateAccessRequest() *MysqlCreateAccessRequest {
	this := MysqlCreateAccessRequest{}
	return &this
}

// NewMysqlCreateAccessRequestWithDefaults instantiates a new MysqlCreateAccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlCreateAccessRequestWithDefaults() *MysqlCreateAccessRequest {
	this := MysqlCreateAccessRequest{}
	return &this
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *MysqlCreateAccessRequest) GetServiceId() string {
	if o == nil || isNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateAccessRequest) GetServiceIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceId) {
    return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *MysqlCreateAccessRequest) HasServiceId() bool {
	if o != nil && !isNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *MysqlCreateAccessRequest) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetDbName returns the DbName field value if set, zero value otherwise.
func (o *MysqlCreateAccessRequest) GetDbName() string {
	if o == nil || isNil(o.DbName) {
		var ret string
		return ret
	}
	return *o.DbName
}

// GetDbNameOk returns a tuple with the DbName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateAccessRequest) GetDbNameOk() (*string, bool) {
	if o == nil || isNil(o.DbName) {
    return nil, false
	}
	return o.DbName, true
}

// HasDbName returns a boolean if a field has been set.
func (o *MysqlCreateAccessRequest) HasDbName() bool {
	if o != nil && !isNil(o.DbName) {
		return true
	}

	return false
}

// SetDbName gets a reference to the given string and assigns it to the DbName field.
func (o *MysqlCreateAccessRequest) SetDbName(v string) {
	o.DbName = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *MysqlCreateAccessRequest) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateAccessRequest) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *MysqlCreateAccessRequest) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *MysqlCreateAccessRequest) SetHost(v string) {
	o.Host = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *MysqlCreateAccessRequest) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateAccessRequest) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
    return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *MysqlCreateAccessRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *MysqlCreateAccessRequest) SetPassword(v string) {
	o.Password = &v
}

// GetSavePmaPassword returns the SavePmaPassword field value if set, zero value otherwise.
func (o *MysqlCreateAccessRequest) GetSavePmaPassword() bool {
	if o == nil || isNil(o.SavePmaPassword) {
		var ret bool
		return ret
	}
	return *o.SavePmaPassword
}

// GetSavePmaPasswordOk returns a tuple with the SavePmaPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateAccessRequest) GetSavePmaPasswordOk() (*bool, bool) {
	if o == nil || isNil(o.SavePmaPassword) {
    return nil, false
	}
	return o.SavePmaPassword, true
}

// HasSavePmaPassword returns a boolean if a field has been set.
func (o *MysqlCreateAccessRequest) HasSavePmaPassword() bool {
	if o != nil && !isNil(o.SavePmaPassword) {
		return true
	}

	return false
}

// SetSavePmaPassword gets a reference to the given bool and assigns it to the SavePmaPassword field.
func (o *MysqlCreateAccessRequest) SetSavePmaPassword(v bool) {
	o.SavePmaPassword = &v
}

func (o MysqlCreateAccessRequest) MarshalJSON() ([]byte, error) {
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

type NullableMysqlCreateAccessRequest struct {
	value *MysqlCreateAccessRequest
	isSet bool
}

func (v NullableMysqlCreateAccessRequest) Get() *MysqlCreateAccessRequest {
	return v.value
}

func (v *NullableMysqlCreateAccessRequest) Set(val *MysqlCreateAccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlCreateAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlCreateAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlCreateAccessRequest(val *MysqlCreateAccessRequest) *NullableMysqlCreateAccessRequest {
	return &NullableMysqlCreateAccessRequest{value: val, isSet: true}
}

func (v NullableMysqlCreateAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlCreateAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


