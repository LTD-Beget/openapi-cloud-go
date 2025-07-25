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

// StructuresAddressInfo struct for StructuresAddressInfo
type StructuresAddressInfo struct {
	Public []StructuresAddressInfoPublicAddress `json:"public,omitempty"`
	Private []StructuresAddressInfoPrivateAddress `json:"private,omitempty"`
}

// NewStructuresAddressInfo instantiates a new StructuresAddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresAddressInfo() *StructuresAddressInfo {
	this := StructuresAddressInfo{}
	return &this
}

// NewStructuresAddressInfoWithDefaults instantiates a new StructuresAddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresAddressInfoWithDefaults() *StructuresAddressInfo {
	this := StructuresAddressInfo{}
	return &this
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *StructuresAddressInfo) GetPublic() []StructuresAddressInfoPublicAddress {
	if o == nil || isNil(o.Public) {
		var ret []StructuresAddressInfoPublicAddress
		return ret
	}
	return o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresAddressInfo) GetPublicOk() ([]StructuresAddressInfoPublicAddress, bool) {
	if o == nil || isNil(o.Public) {
    return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *StructuresAddressInfo) HasPublic() bool {
	if o != nil && !isNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given []StructuresAddressInfoPublicAddress and assigns it to the Public field.
func (o *StructuresAddressInfo) SetPublic(v []StructuresAddressInfoPublicAddress) {
	o.Public = v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *StructuresAddressInfo) GetPrivate() []StructuresAddressInfoPrivateAddress {
	if o == nil || isNil(o.Private) {
		var ret []StructuresAddressInfoPrivateAddress
		return ret
	}
	return o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresAddressInfo) GetPrivateOk() ([]StructuresAddressInfoPrivateAddress, bool) {
	if o == nil || isNil(o.Private) {
    return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *StructuresAddressInfo) HasPrivate() bool {
	if o != nil && !isNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given []StructuresAddressInfoPrivateAddress and assigns it to the Private field.
func (o *StructuresAddressInfo) SetPrivate(v []StructuresAddressInfoPrivateAddress) {
	o.Private = v
}

func (o StructuresAddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !isNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresAddressInfo struct {
	value *StructuresAddressInfo
	isSet bool
}

func (v NullableStructuresAddressInfo) Get() *StructuresAddressInfo {
	return v.value
}

func (v *NullableStructuresAddressInfo) Set(val *StructuresAddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresAddressInfo(val *StructuresAddressInfo) *NullableStructuresAddressInfo {
	return &NullableStructuresAddressInfo{value: val, isSet: true}
}

func (v NullableStructuresAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


