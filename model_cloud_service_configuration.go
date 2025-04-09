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

// CloudServiceConfiguration struct for CloudServiceConfiguration
type CloudServiceConfiguration struct {
	Id *string `json:"id,omitempty"`
	PriceDay *float64 `json:"price_day,omitempty"`
	PriceMonth *float64 `json:"price_month,omitempty"`
	Region *string `json:"region,omitempty"`
	Type *string `json:"type,omitempty"`
	Mysql5 *MysqlMysql5Configuration `json:"mysql5,omitempty"`
	Mysql8 *MysqlMysql8Configuration `json:"mysql8,omitempty"`
	Mysql84 *MysqlMysql84Configuration `json:"mysql84,omitempty"`
	Postgresql15 *PostgresqlPostgresql15Configuration `json:"postgresql15,omitempty"`
	Postgresql14 *PostgresqlPostgresql14Configuration `json:"postgresql14,omitempty"`
	Postgresql164 *PostgresqlPostgresql164Configuration `json:"postgresql164,omitempty"`
	S3 *S3S3Configuration `json:"s3,omitempty"`
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

// GetType returns the Type field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CloudServiceConfiguration) SetType(v string) {
	o.Type = &v
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

// GetMysql84 returns the Mysql84 field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetMysql84() MysqlMysql84Configuration {
	if o == nil || isNil(o.Mysql84) {
		var ret MysqlMysql84Configuration
		return ret
	}
	return *o.Mysql84
}

// GetMysql84Ok returns a tuple with the Mysql84 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetMysql84Ok() (*MysqlMysql84Configuration, bool) {
	if o == nil || isNil(o.Mysql84) {
    return nil, false
	}
	return o.Mysql84, true
}

// HasMysql84 returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasMysql84() bool {
	if o != nil && !isNil(o.Mysql84) {
		return true
	}

	return false
}

// SetMysql84 gets a reference to the given MysqlMysql84Configuration and assigns it to the Mysql84 field.
func (o *CloudServiceConfiguration) SetMysql84(v MysqlMysql84Configuration) {
	o.Mysql84 = &v
}

// GetPostgresql15 returns the Postgresql15 field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetPostgresql15() PostgresqlPostgresql15Configuration {
	if o == nil || isNil(o.Postgresql15) {
		var ret PostgresqlPostgresql15Configuration
		return ret
	}
	return *o.Postgresql15
}

// GetPostgresql15Ok returns a tuple with the Postgresql15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetPostgresql15Ok() (*PostgresqlPostgresql15Configuration, bool) {
	if o == nil || isNil(o.Postgresql15) {
    return nil, false
	}
	return o.Postgresql15, true
}

// HasPostgresql15 returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasPostgresql15() bool {
	if o != nil && !isNil(o.Postgresql15) {
		return true
	}

	return false
}

// SetPostgresql15 gets a reference to the given PostgresqlPostgresql15Configuration and assigns it to the Postgresql15 field.
func (o *CloudServiceConfiguration) SetPostgresql15(v PostgresqlPostgresql15Configuration) {
	o.Postgresql15 = &v
}

// GetPostgresql14 returns the Postgresql14 field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetPostgresql14() PostgresqlPostgresql14Configuration {
	if o == nil || isNil(o.Postgresql14) {
		var ret PostgresqlPostgresql14Configuration
		return ret
	}
	return *o.Postgresql14
}

// GetPostgresql14Ok returns a tuple with the Postgresql14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetPostgresql14Ok() (*PostgresqlPostgresql14Configuration, bool) {
	if o == nil || isNil(o.Postgresql14) {
    return nil, false
	}
	return o.Postgresql14, true
}

// HasPostgresql14 returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasPostgresql14() bool {
	if o != nil && !isNil(o.Postgresql14) {
		return true
	}

	return false
}

// SetPostgresql14 gets a reference to the given PostgresqlPostgresql14Configuration and assigns it to the Postgresql14 field.
func (o *CloudServiceConfiguration) SetPostgresql14(v PostgresqlPostgresql14Configuration) {
	o.Postgresql14 = &v
}

// GetPostgresql164 returns the Postgresql164 field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetPostgresql164() PostgresqlPostgresql164Configuration {
	if o == nil || isNil(o.Postgresql164) {
		var ret PostgresqlPostgresql164Configuration
		return ret
	}
	return *o.Postgresql164
}

// GetPostgresql164Ok returns a tuple with the Postgresql164 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetPostgresql164Ok() (*PostgresqlPostgresql164Configuration, bool) {
	if o == nil || isNil(o.Postgresql164) {
    return nil, false
	}
	return o.Postgresql164, true
}

// HasPostgresql164 returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasPostgresql164() bool {
	if o != nil && !isNil(o.Postgresql164) {
		return true
	}

	return false
}

// SetPostgresql164 gets a reference to the given PostgresqlPostgresql164Configuration and assigns it to the Postgresql164 field.
func (o *CloudServiceConfiguration) SetPostgresql164(v PostgresqlPostgresql164Configuration) {
	o.Postgresql164 = &v
}

// GetS3 returns the S3 field value if set, zero value otherwise.
func (o *CloudServiceConfiguration) GetS3() S3S3Configuration {
	if o == nil || isNil(o.S3) {
		var ret S3S3Configuration
		return ret
	}
	return *o.S3
}

// GetS3Ok returns a tuple with the S3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudServiceConfiguration) GetS3Ok() (*S3S3Configuration, bool) {
	if o == nil || isNil(o.S3) {
    return nil, false
	}
	return o.S3, true
}

// HasS3 returns a boolean if a field has been set.
func (o *CloudServiceConfiguration) HasS3() bool {
	if o != nil && !isNil(o.S3) {
		return true
	}

	return false
}

// SetS3 gets a reference to the given S3S3Configuration and assigns it to the S3 field.
func (o *CloudServiceConfiguration) SetS3(v S3S3Configuration) {
	o.S3 = &v
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
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Mysql5) {
		toSerialize["mysql5"] = o.Mysql5
	}
	if !isNil(o.Mysql8) {
		toSerialize["mysql8"] = o.Mysql8
	}
	if !isNil(o.Mysql84) {
		toSerialize["mysql84"] = o.Mysql84
	}
	if !isNil(o.Postgresql15) {
		toSerialize["postgresql15"] = o.Postgresql15
	}
	if !isNil(o.Postgresql14) {
		toSerialize["postgresql14"] = o.Postgresql14
	}
	if !isNil(o.Postgresql164) {
		toSerialize["postgresql164"] = o.Postgresql164
	}
	if !isNil(o.S3) {
		toSerialize["s3"] = o.S3
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


