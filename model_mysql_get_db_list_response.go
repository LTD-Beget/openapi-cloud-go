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

// MysqlGetDbListResponse struct for MysqlGetDbListResponse
type MysqlGetDbListResponse struct {
	Db []MysqlDb `json:"db,omitempty"`
}

// NewMysqlGetDbListResponse instantiates a new MysqlGetDbListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlGetDbListResponse() *MysqlGetDbListResponse {
	this := MysqlGetDbListResponse{}
	return &this
}

// NewMysqlGetDbListResponseWithDefaults instantiates a new MysqlGetDbListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlGetDbListResponseWithDefaults() *MysqlGetDbListResponse {
	this := MysqlGetDbListResponse{}
	return &this
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *MysqlGetDbListResponse) GetDb() []MysqlDb {
	if o == nil || isNil(o.Db) {
		var ret []MysqlDb
		return ret
	}
	return o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlGetDbListResponse) GetDbOk() ([]MysqlDb, bool) {
	if o == nil || isNil(o.Db) {
    return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *MysqlGetDbListResponse) HasDb() bool {
	if o != nil && !isNil(o.Db) {
		return true
	}

	return false
}

// SetDb gets a reference to the given []MysqlDb and assigns it to the Db field.
func (o *MysqlGetDbListResponse) SetDb(v []MysqlDb) {
	o.Db = v
}

func (o MysqlGetDbListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Db) {
		toSerialize["db"] = o.Db
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlGetDbListResponse struct {
	value *MysqlGetDbListResponse
	isSet bool
}

func (v NullableMysqlGetDbListResponse) Get() *MysqlGetDbListResponse {
	return v.value
}

func (v *NullableMysqlGetDbListResponse) Set(val *MysqlGetDbListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlGetDbListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlGetDbListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlGetDbListResponse(val *MysqlGetDbListResponse) *NullableMysqlGetDbListResponse {
	return &NullableMysqlGetDbListResponse{value: val, isSet: true}
}

func (v NullableMysqlGetDbListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlGetDbListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


