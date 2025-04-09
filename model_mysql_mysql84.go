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

// MysqlMysql84 struct for MysqlMysql84
type MysqlMysql84 struct {
	Configuration *MysqlMysql84Configuration `json:"configuration,omitempty"`
	Host *string `json:"host,omitempty"`
	Port *int32 `json:"port,omitempty"`
	AddressInfo *StructuresAddressInfo `json:"address_info,omitempty"`
	PmaUrl *string `json:"pma_url,omitempty"`
	DiskUsed *string `json:"disk_used,omitempty"`
	DiskLeft *string `json:"disk_left,omitempty"`
	ReadOnly *bool `json:"read_only,omitempty"`
	PmaInstalling *bool `json:"pma_installing,omitempty"`
}

// NewMysqlMysql84 instantiates a new MysqlMysql84 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlMysql84() *MysqlMysql84 {
	this := MysqlMysql84{}
	return &this
}

// NewMysqlMysql84WithDefaults instantiates a new MysqlMysql84 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlMysql84WithDefaults() *MysqlMysql84 {
	this := MysqlMysql84{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *MysqlMysql84) GetConfiguration() MysqlMysql84Configuration {
	if o == nil || isNil(o.Configuration) {
		var ret MysqlMysql84Configuration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql84) GetConfigurationOk() (*MysqlMysql84Configuration, bool) {
	if o == nil || isNil(o.Configuration) {
    return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *MysqlMysql84) HasConfiguration() bool {
	if o != nil && !isNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given MysqlMysql84Configuration and assigns it to the Configuration field.
func (o *MysqlMysql84) SetConfiguration(v MysqlMysql84Configuration) {
	o.Configuration = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *MysqlMysql84) GetHost() string {
	if o == nil || isNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql84) GetHostOk() (*string, bool) {
	if o == nil || isNil(o.Host) {
    return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *MysqlMysql84) HasHost() bool {
	if o != nil && !isNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *MysqlMysql84) SetHost(v string) {
	o.Host = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *MysqlMysql84) GetPort() int32 {
	if o == nil || isNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql84) GetPortOk() (*int32, bool) {
	if o == nil || isNil(o.Port) {
    return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *MysqlMysql84) HasPort() bool {
	if o != nil && !isNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *MysqlMysql84) SetPort(v int32) {
	o.Port = &v
}

// GetAddressInfo returns the AddressInfo field value if set, zero value otherwise.
func (o *MysqlMysql84) GetAddressInfo() StructuresAddressInfo {
	if o == nil || isNil(o.AddressInfo) {
		var ret StructuresAddressInfo
		return ret
	}
	return *o.AddressInfo
}

// GetAddressInfoOk returns a tuple with the AddressInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql84) GetAddressInfoOk() (*StructuresAddressInfo, bool) {
	if o == nil || isNil(o.AddressInfo) {
    return nil, false
	}
	return o.AddressInfo, true
}

// HasAddressInfo returns a boolean if a field has been set.
func (o *MysqlMysql84) HasAddressInfo() bool {
	if o != nil && !isNil(o.AddressInfo) {
		return true
	}

	return false
}

// SetAddressInfo gets a reference to the given StructuresAddressInfo and assigns it to the AddressInfo field.
func (o *MysqlMysql84) SetAddressInfo(v StructuresAddressInfo) {
	o.AddressInfo = &v
}

// GetPmaUrl returns the PmaUrl field value if set, zero value otherwise.
func (o *MysqlMysql84) GetPmaUrl() string {
	if o == nil || isNil(o.PmaUrl) {
		var ret string
		return ret
	}
	return *o.PmaUrl
}

// GetPmaUrlOk returns a tuple with the PmaUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql84) GetPmaUrlOk() (*string, bool) {
	if o == nil || isNil(o.PmaUrl) {
    return nil, false
	}
	return o.PmaUrl, true
}

// HasPmaUrl returns a boolean if a field has been set.
func (o *MysqlMysql84) HasPmaUrl() bool {
	if o != nil && !isNil(o.PmaUrl) {
		return true
	}

	return false
}

// SetPmaUrl gets a reference to the given string and assigns it to the PmaUrl field.
func (o *MysqlMysql84) SetPmaUrl(v string) {
	o.PmaUrl = &v
}

// GetDiskUsed returns the DiskUsed field value if set, zero value otherwise.
func (o *MysqlMysql84) GetDiskUsed() string {
	if o == nil || isNil(o.DiskUsed) {
		var ret string
		return ret
	}
	return *o.DiskUsed
}

// GetDiskUsedOk returns a tuple with the DiskUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql84) GetDiskUsedOk() (*string, bool) {
	if o == nil || isNil(o.DiskUsed) {
    return nil, false
	}
	return o.DiskUsed, true
}

// HasDiskUsed returns a boolean if a field has been set.
func (o *MysqlMysql84) HasDiskUsed() bool {
	if o != nil && !isNil(o.DiskUsed) {
		return true
	}

	return false
}

// SetDiskUsed gets a reference to the given string and assigns it to the DiskUsed field.
func (o *MysqlMysql84) SetDiskUsed(v string) {
	o.DiskUsed = &v
}

// GetDiskLeft returns the DiskLeft field value if set, zero value otherwise.
func (o *MysqlMysql84) GetDiskLeft() string {
	if o == nil || isNil(o.DiskLeft) {
		var ret string
		return ret
	}
	return *o.DiskLeft
}

// GetDiskLeftOk returns a tuple with the DiskLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql84) GetDiskLeftOk() (*string, bool) {
	if o == nil || isNil(o.DiskLeft) {
    return nil, false
	}
	return o.DiskLeft, true
}

// HasDiskLeft returns a boolean if a field has been set.
func (o *MysqlMysql84) HasDiskLeft() bool {
	if o != nil && !isNil(o.DiskLeft) {
		return true
	}

	return false
}

// SetDiskLeft gets a reference to the given string and assigns it to the DiskLeft field.
func (o *MysqlMysql84) SetDiskLeft(v string) {
	o.DiskLeft = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *MysqlMysql84) GetReadOnly() bool {
	if o == nil || isNil(o.ReadOnly) {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql84) GetReadOnlyOk() (*bool, bool) {
	if o == nil || isNil(o.ReadOnly) {
    return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *MysqlMysql84) HasReadOnly() bool {
	if o != nil && !isNil(o.ReadOnly) {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *MysqlMysql84) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

// GetPmaInstalling returns the PmaInstalling field value if set, zero value otherwise.
func (o *MysqlMysql84) GetPmaInstalling() bool {
	if o == nil || isNil(o.PmaInstalling) {
		var ret bool
		return ret
	}
	return *o.PmaInstalling
}

// GetPmaInstallingOk returns a tuple with the PmaInstalling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlMysql84) GetPmaInstallingOk() (*bool, bool) {
	if o == nil || isNil(o.PmaInstalling) {
    return nil, false
	}
	return o.PmaInstalling, true
}

// HasPmaInstalling returns a boolean if a field has been set.
func (o *MysqlMysql84) HasPmaInstalling() bool {
	if o != nil && !isNil(o.PmaInstalling) {
		return true
	}

	return false
}

// SetPmaInstalling gets a reference to the given bool and assigns it to the PmaInstalling field.
func (o *MysqlMysql84) SetPmaInstalling(v bool) {
	o.PmaInstalling = &v
}

func (o MysqlMysql84) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !isNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !isNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !isNil(o.AddressInfo) {
		toSerialize["address_info"] = o.AddressInfo
	}
	if !isNil(o.PmaUrl) {
		toSerialize["pma_url"] = o.PmaUrl
	}
	if !isNil(o.DiskUsed) {
		toSerialize["disk_used"] = o.DiskUsed
	}
	if !isNil(o.DiskLeft) {
		toSerialize["disk_left"] = o.DiskLeft
	}
	if !isNil(o.ReadOnly) {
		toSerialize["read_only"] = o.ReadOnly
	}
	if !isNil(o.PmaInstalling) {
		toSerialize["pma_installing"] = o.PmaInstalling
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlMysql84 struct {
	value *MysqlMysql84
	isSet bool
}

func (v NullableMysqlMysql84) Get() *MysqlMysql84 {
	return v.value
}

func (v *NullableMysqlMysql84) Set(val *MysqlMysql84) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlMysql84) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlMysql84) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlMysql84(val *MysqlMysql84) *NullableMysqlMysql84 {
	return &NullableMysqlMysql84{value: val, isSet: true}
}

func (v NullableMysqlMysql84) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlMysql84) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


