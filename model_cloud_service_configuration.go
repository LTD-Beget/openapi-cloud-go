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

// CloudServiceConfiguration struct for CloudServiceConfiguration
type CloudServiceConfiguration struct {
	Id *string `json:"id,omitempty"`
	PriceDay *float64 `json:"price_day,omitempty"`
	PriceMonth *float64 `json:"price_month,omitempty"`
	Region *string `json:"region,omitempty"`
	Mysql5 *MysqlMysql5Configuration `json:"mysql5,omitempty"`
	Mysql8 *MysqlMysql8Configuration `json:"mysql8,omitempty"`
}

// NewCloudServiceConfiguration instantiates a new CloudServiceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudServiceConfiguration() *CloudServiceConfiguration {
	this := CloudServiceConfiguration{}
	return &this
}

// NewCloudServiceConfigurationWithDefaults instantiates a new CloudServiceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudServiceConfigurationWithDefaults() *CloudServiceConfiguration {
	this := CloudServiceConfiguration{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudServiceConfiguration) SetId(v string) {
	o.Id = &v
}

// GetPriceDay returns the PriceDay field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetPriceDay() float64 {
	if o == nil || isNil(o.PriceDay) {
		var ret float64
		return ret
	}
	return *o.PriceDay
}

// GetPriceDayOk returns a tuple with the PriceDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetPriceDayOk() (*float64, bool) {
	if o == nil || isNil(o.PriceDay) {
    return nil, false
	}
	return o.PriceDay, true
}

// HasPriceDay returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasPriceDay() bool {
	if o != nil && !isNil(o.PriceDay) {
		return true
	}

	return false
}

// SetPriceDay gets a reference to the given float64 and assigns it to the PriceDay field.
func (o *CloudServiceConfiguration) SetPriceDay(v float64) {
	o.PriceDay = &v
}

// GetPriceMonth returns the PriceMonth field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetPriceMonth() float64 {
	if o == nil || isNil(o.PriceMonth) {
		var ret float64
		return ret
	}
	return *o.PriceMonth
}

// GetPriceMonthOk returns a tuple with the PriceMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetPriceMonthOk() (*float64, bool) {
	if o == nil || isNil(o.PriceMonth) {
    return nil, false
	}
	return o.PriceMonth, true
}

// HasPriceMonth returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasPriceMonth() bool {
	if o != nil && !isNil(o.PriceMonth) {
		return true
	}

	return false
}

// SetPriceMonth gets a reference to the given float64 and assigns it to the PriceMonth field.
func (o *CloudServiceConfiguration) SetPriceMonth(v float64) {
	o.PriceMonth = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CloudServiceConfiguration) SetRegion(v string) {
	o.Region = &v
}

// GetMysql5 returns the Mysql5 field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetMysql5() MysqlMysql5Configuration {
	if o == nil || isNil(o.Mysql5) {
		var ret MysqlMysql5Configuration
		return ret
	}
	return *o.Mysql5
}

// GetMysql5Ok returns a tuple with the Mysql5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetMysql5Ok() (*MysqlMysql5Configuration, bool) {
	if o == nil || isNil(o.Mysql5) {
    return nil, false
	}
	return o.Mysql5, true
}

// HasMysql5 returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasMysql5() bool {
	if o != nil && !isNil(o.Mysql5) {
		return true
	}

	return false
}

// SetMysql5 gets a reference to the given MysqlMysql5Configuration and assigns it to the Mysql5 field.
func (o *CloudServiceConfiguration) SetMysql5(v MysqlMysql5Configuration) {
	o.Mysql5 = &v
}

// GetMysql8 returns the Mysql8 field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetMysql8() MysqlMysql8Configuration {
	if o == nil || isNil(o.Mysql8) {
		var ret MysqlMysql8Configuration
		return ret
	}
	return *o.Mysql8
}

// GetMysql8Ok returns a tuple with the Mysql8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetMysql8Ok() (*MysqlMysql8Configuration, bool) {
	if o == nil || isNil(o.Mysql8) {
    return nil, false
	}
	return o.Mysql8, true
}

// HasMysql8 returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasMysql8() bool {
	if o != nil && !isNil(o.Mysql8) {
		return true
	}

	return false
}

// SetMysql8 gets a reference to the given MysqlMysql8Configuration and assigns it to the Mysql8 field.
func (o *CloudServiceConfiguration) SetMysql8(v MysqlMysql8Configuration) {
	o.Mysql8 = &v
}

func (o CloudServiceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.PriceDay) {
		toSerialize["price_day"] = o.PriceDay
	}
	if !isNil(o.PriceMonth) {
		toSerialize["price_month"] = o.PriceMonth
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Mysql5) {
		toSerialize["mysql5"] = o.Mysql5
	}
	if !isNil(o.Mysql8) {
		toSerialize["mysql8"] = o.Mysql8
	}
	return json.Marshal(toSerialize)
}

type NullableCloudServiceConfiguration struct {
	value *CloudServiceConfiguration
	isSet bool
}

func (v NullableCloudServiceConfiguration) Get() *CloudServiceConfiguration {
	return v.value
}

func (v *NullableCloudServiceConfiguration) Set(val *CloudServiceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudServiceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudServiceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudServiceConfiguration(val *CloudServiceConfiguration) *NullableCloudServiceConfiguration {
	return &NullableCloudServiceConfiguration{value: val, isSet: true}
}

func (v NullableCloudServiceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudServiceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


