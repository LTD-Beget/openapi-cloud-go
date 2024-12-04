/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// MysqlCreateDbResponse struct for MysqlCreateDbResponse
type MysqlCreateDbResponse struct {
	Db *MysqlDb `json:"db,omitempty"`
	Error *MysqlCreateDbResponseError `json:"error,omitempty"`
}

// NewMysqlCreateDbResponse instantiates a new MysqlCreateDbResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlCreateDbResponse() *MysqlCreateDbResponse {
	this := MysqlCreateDbResponse{}
	return &this
}

// NewMysqlCreateDbResponseWithDefaults instantiates a new MysqlCreateDbResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlCreateDbResponseWithDefaults() *MysqlCreateDbResponse {
	this := MysqlCreateDbResponse{}
	return &this
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *MysqlCreateDbResponse) GetDb() MysqlDb {
	if o == nil || isNil(o.Db) {
		var ret MysqlDb
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateDbResponse) GetDbOk() (*MysqlDb, bool) {
	if o == nil || isNil(o.Db) {
    return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *MysqlCreateDbResponse) HasDb() bool {
	if o != nil && !isNil(o.Db) {
		return true
	}

	return false
}

// SetDb gets a reference to the given MysqlDb and assigns it to the Db field.
func (o *MysqlCreateDbResponse) SetDb(v MysqlDb) {
	o.Db = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MysqlCreateDbResponse) GetError() MysqlCreateDbResponseError {
	if o == nil || isNil(o.Error) {
		var ret MysqlCreateDbResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateDbResponse) GetErrorOk() (*MysqlCreateDbResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MysqlCreateDbResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given MysqlCreateDbResponseError and assigns it to the Error field.
func (o *MysqlCreateDbResponse) SetError(v MysqlCreateDbResponseError) {
	o.Error = &v
}

func (o MysqlCreateDbResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Db) {
		toSerialize["db"] = o.Db
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlCreateDbResponse struct {
	value *MysqlCreateDbResponse
	isSet bool
}

func (v NullableMysqlCreateDbResponse) Get() *MysqlCreateDbResponse {
	return v.value
}

func (v *NullableMysqlCreateDbResponse) Set(val *MysqlCreateDbResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlCreateDbResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlCreateDbResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlCreateDbResponse(val *MysqlCreateDbResponse) *NullableMysqlCreateDbResponse {
	return &NullableMysqlCreateDbResponse{value: val, isSet: true}
}

func (v NullableMysqlCreateDbResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlCreateDbResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


