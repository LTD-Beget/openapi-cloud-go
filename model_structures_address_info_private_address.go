/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// StructuresAddressInfoPrivateAddress struct for StructuresAddressInfoPrivateAddress
type StructuresAddressInfoPrivateAddress struct {
	Ip *string `json:"ip,omitempty"`
	NetworkId *string `json:"network_id,omitempty"`
}

// NewStructuresAddressInfoPrivateAddress instantiates a new StructuresAddressInfoPrivateAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresAddressInfoPrivateAddress() *StructuresAddressInfoPrivateAddress {
	this := StructuresAddressInfoPrivateAddress{}
	return &this
}

// NewStructuresAddressInfoPrivateAddressWithDefaults instantiates a new StructuresAddressInfoPrivateAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresAddressInfoPrivateAddressWithDefaults() *StructuresAddressInfoPrivateAddress {
	this := StructuresAddressInfoPrivateAddress{}
	return &this
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *StructuresAddressInfoPrivateAddress) GetIp() string {
	if o == nil || isNil(o.Ip) {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresAddressInfoPrivateAddress) GetIpOk() (*string, bool) {
	if o == nil || isNil(o.Ip) {
    return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *StructuresAddressInfoPrivateAddress) HasIp() bool {
	if o != nil && !isNil(o.Ip) {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *StructuresAddressInfoPrivateAddress) SetIp(v string) {
	o.Ip = &v
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *StructuresAddressInfoPrivateAddress) GetNetworkId() string {
	if o == nil || isNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresAddressInfoPrivateAddress) GetNetworkIdOk() (*string, bool) {
	if o == nil || isNil(o.NetworkId) {
    return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *StructuresAddressInfoPrivateAddress) HasNetworkId() bool {
	if o != nil && !isNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *StructuresAddressInfoPrivateAddress) SetNetworkId(v string) {
	o.NetworkId = &v
}

func (o StructuresAddressInfoPrivateAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Ip) {
		toSerialize["ip"] = o.Ip
	}
	if !isNil(o.NetworkId) {
		toSerialize["network_id"] = o.NetworkId
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresAddressInfoPrivateAddress struct {
	value *StructuresAddressInfoPrivateAddress
	isSet bool
}

func (v NullableStructuresAddressInfoPrivateAddress) Get() *StructuresAddressInfoPrivateAddress {
	return v.value
}

func (v *NullableStructuresAddressInfoPrivateAddress) Set(val *StructuresAddressInfoPrivateAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresAddressInfoPrivateAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresAddressInfoPrivateAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresAddressInfoPrivateAddress(val *StructuresAddressInfoPrivateAddress) *NullableStructuresAddressInfoPrivateAddress {
	return &NullableStructuresAddressInfoPrivateAddress{value: val, isSet: true}
}

func (v NullableStructuresAddressInfoPrivateAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresAddressInfoPrivateAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


