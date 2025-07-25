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

// MysqlBackupGetOrdersResponse struct for MysqlBackupGetOrdersResponse
type MysqlBackupGetOrdersResponse struct {
	Order []StructuresMysqlRestoreOrder `json:"order,omitempty"`
}

// NewMysqlBackupGetOrdersResponse instantiates a new MysqlBackupGetOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlBackupGetOrdersResponse() *MysqlBackupGetOrdersResponse {
	this := MysqlBackupGetOrdersResponse{}
	return &this
}

// NewMysqlBackupGetOrdersResponseWithDefaults instantiates a new MysqlBackupGetOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlBackupGetOrdersResponseWithDefaults() *MysqlBackupGetOrdersResponse {
	this := MysqlBackupGetOrdersResponse{}
	return &this
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *MysqlBackupGetOrdersResponse) GetOrder() []StructuresMysqlRestoreOrder {
	if o == nil || isNil(o.Order) {
		var ret []StructuresMysqlRestoreOrder
		return ret
	}
	return o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlBackupGetOrdersResponse) GetOrderOk() ([]StructuresMysqlRestoreOrder, bool) {
	if o == nil || isNil(o.Order) {
    return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *MysqlBackupGetOrdersResponse) HasOrder() bool {
	if o != nil && !isNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given []StructuresMysqlRestoreOrder and assigns it to the Order field.
func (o *MysqlBackupGetOrdersResponse) SetOrder(v []StructuresMysqlRestoreOrder) {
	o.Order = v
}

func (o MysqlBackupGetOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlBackupGetOrdersResponse struct {
	value *MysqlBackupGetOrdersResponse
	isSet bool
}

func (v NullableMysqlBackupGetOrdersResponse) Get() *MysqlBackupGetOrdersResponse {
	return v.value
}

func (v *NullableMysqlBackupGetOrdersResponse) Set(val *MysqlBackupGetOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlBackupGetOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlBackupGetOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlBackupGetOrdersResponse(val *MysqlBackupGetOrdersResponse) *NullableMysqlBackupGetOrdersResponse {
	return &NullableMysqlBackupGetOrdersResponse{value: val, isSet: true}
}

func (v NullableMysqlBackupGetOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlBackupGetOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


