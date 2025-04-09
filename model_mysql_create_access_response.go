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

// MysqlCreateAccessResponse struct for MysqlCreateAccessResponse
type MysqlCreateAccessResponse struct {
	Access *MysqlDbAccess `json:"access,omitempty"`
	Error *MysqlCreateAccessResponseError `json:"error,omitempty"`
}

// NewMysqlCreateAccessResponse instantiates a new MysqlCreateAccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlCreateAccessResponse() *MysqlCreateAccessResponse {
	this := MysqlCreateAccessResponse{}
	return &this
}

// NewMysqlCreateAccessResponseWithDefaults instantiates a new MysqlCreateAccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlCreateAccessResponseWithDefaults() *MysqlCreateAccessResponse {
	this := MysqlCreateAccessResponse{}
	return &this
}

// GetAccess returns the Access field value if set, zero value otherwise.
func (o *MysqlCreateAccessResponse) GetAccess() MysqlDbAccess {
	if o == nil || isNil(o.Access) {
		var ret MysqlDbAccess
		return ret
	}
	return *o.Access
}

// GetAccessOk returns a tuple with the Access field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateAccessResponse) GetAccessOk() (*MysqlDbAccess, bool) {
	if o == nil || isNil(o.Access) {
    return nil, false
	}
	return o.Access, true
}

// HasAccess returns a boolean if a field has been set.
func (o *MysqlCreateAccessResponse) HasAccess() bool {
	if o != nil && !isNil(o.Access) {
		return true
	}

	return false
}

// SetAccess gets a reference to the given MysqlDbAccess and assigns it to the Access field.
func (o *MysqlCreateAccessResponse) SetAccess(v MysqlDbAccess) {
	o.Access = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MysqlCreateAccessResponse) GetError() MysqlCreateAccessResponseError {
	if o == nil || isNil(o.Error) {
		var ret MysqlCreateAccessResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlCreateAccessResponse) GetErrorOk() (*MysqlCreateAccessResponseError, bool) {
	if o == nil || isNil(o.Error) {
    return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MysqlCreateAccessResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given MysqlCreateAccessResponseError and assigns it to the Error field.
func (o *MysqlCreateAccessResponse) SetError(v MysqlCreateAccessResponseError) {
	o.Error = &v
}

func (o MysqlCreateAccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Access) {
		toSerialize["access"] = o.Access
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlCreateAccessResponse struct {
	value *MysqlCreateAccessResponse
	isSet bool
}

func (v NullableMysqlCreateAccessResponse) Get() *MysqlCreateAccessResponse {
	return v.value
}

func (v *NullableMysqlCreateAccessResponse) Set(val *MysqlCreateAccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlCreateAccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlCreateAccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlCreateAccessResponse(val *MysqlCreateAccessResponse) *NullableMysqlCreateAccessResponse {
	return &NullableMysqlCreateAccessResponse{value: val, isSet: true}
}

func (v NullableMysqlCreateAccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlCreateAccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


