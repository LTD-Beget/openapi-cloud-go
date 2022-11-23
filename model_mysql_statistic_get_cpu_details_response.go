/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// MysqlStatisticGetCpuDetailsResponse struct for MysqlStatisticGetCpuDetailsResponse
type MysqlStatisticGetCpuDetailsResponse struct {
	Date []string `json:"date,omitempty"`
	User []float64 `json:"user,omitempty"`
	Nice []float64 `json:"nice,omitempty"`
	System []float64 `json:"system,omitempty"`
	Idle []float64 `json:"idle,omitempty"`
	Iowait []float64 `json:"iowait,omitempty"`
	Irq []float64 `json:"irq,omitempty"`
	Softirq []float64 `json:"softirq,omitempty"`
}

// NewMysqlStatisticGetCpuDetailsResponse instantiates a new MysqlStatisticGetCpuDetailsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMysqlStatisticGetCpuDetailsResponse() *MysqlStatisticGetCpuDetailsResponse {
	this := MysqlStatisticGetCpuDetailsResponse{}
	return &this
}

// NewMysqlStatisticGetCpuDetailsResponseWithDefaults instantiates a new MysqlStatisticGetCpuDetailsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMysqlStatisticGetCpuDetailsResponseWithDefaults() *MysqlStatisticGetCpuDetailsResponse {
	this := MysqlStatisticGetCpuDetailsResponse{}
	return &this
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *MysqlStatisticGetCpuDetailsResponse) GetDate() []string {
	if o == nil || isNil(o.Date) {
		var ret []string
		return ret
	}
	return o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) GetDateOk() ([]string, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given []string and assigns it to the Date field.
func (o *MysqlStatisticGetCpuDetailsResponse) SetDate(v []string) {
	o.Date = v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *MysqlStatisticGetCpuDetailsResponse) GetUser() []float64 {
	if o == nil || isNil(o.User) {
		var ret []float64
		return ret
	}
	return o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) GetUserOk() ([]float64, bool) {
	if o == nil || isNil(o.User) {
    return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given []float64 and assigns it to the User field.
func (o *MysqlStatisticGetCpuDetailsResponse) SetUser(v []float64) {
	o.User = v
}

// GetNice returns the Nice field value if set, zero value otherwise.
func (o *MysqlStatisticGetCpuDetailsResponse) GetNice() []float64 {
	if o == nil || isNil(o.Nice) {
		var ret []float64
		return ret
	}
	return o.Nice
}

// GetNiceOk returns a tuple with the Nice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) GetNiceOk() ([]float64, bool) {
	if o == nil || isNil(o.Nice) {
    return nil, false
	}
	return o.Nice, true
}

// HasNice returns a boolean if a field has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) HasNice() bool {
	if o != nil && !isNil(o.Nice) {
		return true
	}

	return false
}

// SetNice gets a reference to the given []float64 and assigns it to the Nice field.
func (o *MysqlStatisticGetCpuDetailsResponse) SetNice(v []float64) {
	o.Nice = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *MysqlStatisticGetCpuDetailsResponse) GetSystem() []float64 {
	if o == nil || isNil(o.System) {
		var ret []float64
		return ret
	}
	return o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) GetSystemOk() ([]float64, bool) {
	if o == nil || isNil(o.System) {
    return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) HasSystem() bool {
	if o != nil && !isNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given []float64 and assigns it to the System field.
func (o *MysqlStatisticGetCpuDetailsResponse) SetSystem(v []float64) {
	o.System = v
}

// GetIdle returns the Idle field value if set, zero value otherwise.
func (o *MysqlStatisticGetCpuDetailsResponse) GetIdle() []float64 {
	if o == nil || isNil(o.Idle) {
		var ret []float64
		return ret
	}
	return o.Idle
}

// GetIdleOk returns a tuple with the Idle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) GetIdleOk() ([]float64, bool) {
	if o == nil || isNil(o.Idle) {
    return nil, false
	}
	return o.Idle, true
}

// HasIdle returns a boolean if a field has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) HasIdle() bool {
	if o != nil && !isNil(o.Idle) {
		return true
	}

	return false
}

// SetIdle gets a reference to the given []float64 and assigns it to the Idle field.
func (o *MysqlStatisticGetCpuDetailsResponse) SetIdle(v []float64) {
	o.Idle = v
}

// GetIowait returns the Iowait field value if set, zero value otherwise.
func (o *MysqlStatisticGetCpuDetailsResponse) GetIowait() []float64 {
	if o == nil || isNil(o.Iowait) {
		var ret []float64
		return ret
	}
	return o.Iowait
}

// GetIowaitOk returns a tuple with the Iowait field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) GetIowaitOk() ([]float64, bool) {
	if o == nil || isNil(o.Iowait) {
    return nil, false
	}
	return o.Iowait, true
}

// HasIowait returns a boolean if a field has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) HasIowait() bool {
	if o != nil && !isNil(o.Iowait) {
		return true
	}

	return false
}

// SetIowait gets a reference to the given []float64 and assigns it to the Iowait field.
func (o *MysqlStatisticGetCpuDetailsResponse) SetIowait(v []float64) {
	o.Iowait = v
}

// GetIrq returns the Irq field value if set, zero value otherwise.
func (o *MysqlStatisticGetCpuDetailsResponse) GetIrq() []float64 {
	if o == nil || isNil(o.Irq) {
		var ret []float64
		return ret
	}
	return o.Irq
}

// GetIrqOk returns a tuple with the Irq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) GetIrqOk() ([]float64, bool) {
	if o == nil || isNil(o.Irq) {
    return nil, false
	}
	return o.Irq, true
}

// HasIrq returns a boolean if a field has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) HasIrq() bool {
	if o != nil && !isNil(o.Irq) {
		return true
	}

	return false
}

// SetIrq gets a reference to the given []float64 and assigns it to the Irq field.
func (o *MysqlStatisticGetCpuDetailsResponse) SetIrq(v []float64) {
	o.Irq = v
}

// GetSoftirq returns the Softirq field value if set, zero value otherwise.
func (o *MysqlStatisticGetCpuDetailsResponse) GetSoftirq() []float64 {
	if o == nil || isNil(o.Softirq) {
		var ret []float64
		return ret
	}
	return o.Softirq
}

// GetSoftirqOk returns a tuple with the Softirq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) GetSoftirqOk() ([]float64, bool) {
	if o == nil || isNil(o.Softirq) {
    return nil, false
	}
	return o.Softirq, true
}

// HasSoftirq returns a boolean if a field has been set.
func (o *MysqlStatisticGetCpuDetailsResponse) HasSoftirq() bool {
	if o != nil && !isNil(o.Softirq) {
		return true
	}

	return false
}

// SetSoftirq gets a reference to the given []float64 and assigns it to the Softirq field.
func (o *MysqlStatisticGetCpuDetailsResponse) SetSoftirq(v []float64) {
	o.Softirq = v
}

func (o MysqlStatisticGetCpuDetailsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !isNil(o.Nice) {
		toSerialize["nice"] = o.Nice
	}
	if !isNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !isNil(o.Idle) {
		toSerialize["idle"] = o.Idle
	}
	if !isNil(o.Iowait) {
		toSerialize["iowait"] = o.Iowait
	}
	if !isNil(o.Irq) {
		toSerialize["irq"] = o.Irq
	}
	if !isNil(o.Softirq) {
		toSerialize["softirq"] = o.Softirq
	}
	return json.Marshal(toSerialize)
}

type NullableMysqlStatisticGetCpuDetailsResponse struct {
	value *MysqlStatisticGetCpuDetailsResponse
	isSet bool
}

func (v NullableMysqlStatisticGetCpuDetailsResponse) Get() *MysqlStatisticGetCpuDetailsResponse {
	return v.value
}

func (v *NullableMysqlStatisticGetCpuDetailsResponse) Set(val *MysqlStatisticGetCpuDetailsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMysqlStatisticGetCpuDetailsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMysqlStatisticGetCpuDetailsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMysqlStatisticGetCpuDetailsResponse(val *MysqlStatisticGetCpuDetailsResponse) *NullableMysqlStatisticGetCpuDetailsResponse {
	return &NullableMysqlStatisticGetCpuDetailsResponse{value: val, isSet: true}
}

func (v NullableMysqlStatisticGetCpuDetailsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMysqlStatisticGetCpuDetailsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

