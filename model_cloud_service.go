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

// CloudService struct for CloudService
type CloudService struct {
	Id *string `json:"id,omitempty"`
	ConfigurationId *string `json:"configuration_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Description *string `json:"description,omitempty"`
	Status *string `json:"status,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	PriceDay *float64 `json:"price_day,omitempty"`
	PriceMonth *float64 `json:"price_month,omitempty"`
	Mysql5 *MysqlMysql5 `json:"mysql5,omitempty"`
	Mysql8 *MysqlMysql8 `json:"mysql8,omitempty"`
	Postgresql15 *PostgresqlPostgresql15 `json:"postgresql15,omitempty"`
	Postgresql14 *PostgresqlPostgresql14 `json:"postgresql14,omitempty"`
	ManageEnabled *bool `json:"manage_enabled,omitempty"`
	Slug *string `json:"slug,omitempty"`
	MonitorableResources []string `json:"monitorable_resources,omitempty"`
	Unblocking *bool `json:"unblocking,omitempty"`
	Migrating *bool `json:"migrating,omitempty"`
	Region *string `json:"region,omitempty"`
}

// NewCloudService instantiates a new CloudService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudService() *CloudService {
	this := CloudService{}
	return &this
}

// NewCloudServiceWithDefaults instantiates a new CloudService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudServiceWithDefaults() *CloudService {
	this := CloudService{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CloudService) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CloudService) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CloudService) SetId(v string) {
	o.Id = &v
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *CloudService) GetConfigurationId() string {
	if o == nil || isNil(o.ConfigurationId) {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetConfigurationIdOk() (*string, bool) {
	if o == nil || isNil(o.ConfigurationId) {
    return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *CloudService) HasConfigurationId() bool {
	if o != nil && !isNil(o.ConfigurationId) {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *CloudService) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CloudService) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CloudService) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CloudService) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudService) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudService) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudService) SetDescription(v string) {
	o.Description = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CloudService) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CloudService) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CloudService) SetStatus(v string) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CloudService) GetCreatedAt() string {
	if o == nil || isNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetCreatedAtOk() (*string, bool) {
	if o == nil || isNil(o.CreatedAt) {
    return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CloudService) HasCreatedAt() bool {
	if o != nil && !isNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *CloudService) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetPriceDay returns the PriceDay field value if set, zero value otherwise.
func (o *CloudService) GetPriceDay() float64 {
	if o == nil || isNil(o.PriceDay) {
		var ret float64
		return ret
	}
	return *o.PriceDay
}

// GetPriceDayOk returns a tuple with the PriceDay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetPriceDayOk() (*float64, bool) {
	if o == nil || isNil(o.PriceDay) {
    return nil, false
	}
	return o.PriceDay, true
}

// HasPriceDay returns a boolean if a field has been set.
func (o *CloudService) HasPriceDay() bool {
	if o != nil && !isNil(o.PriceDay) {
		return true
	}

	return false
}

// SetPriceDay gets a reference to the given float64 and assigns it to the PriceDay field.
func (o *CloudService) SetPriceDay(v float64) {
	o.PriceDay = &v
}

// GetPriceMonth returns the PriceMonth field value if set, zero value otherwise.
func (o *CloudService) GetPriceMonth() float64 {
	if o == nil || isNil(o.PriceMonth) {
		var ret float64
		return ret
	}
	return *o.PriceMonth
}

// GetPriceMonthOk returns a tuple with the PriceMonth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetPriceMonthOk() (*float64, bool) {
	if o == nil || isNil(o.PriceMonth) {
    return nil, false
	}
	return o.PriceMonth, true
}

// HasPriceMonth returns a boolean if a field has been set.
func (o *CloudService) HasPriceMonth() bool {
	if o != nil && !isNil(o.PriceMonth) {
		return true
	}

	return false
}

// SetPriceMonth gets a reference to the given float64 and assigns it to the PriceMonth field.
func (o *CloudService) SetPriceMonth(v float64) {
	o.PriceMonth = &v
}

// GetMysql5 returns the Mysql5 field value if set, zero value otherwise.
func (o *CloudService) GetMysql5() MysqlMysql5 {
	if o == nil || isNil(o.Mysql5) {
		var ret MysqlMysql5
		return ret
	}
	return *o.Mysql5
}

// GetMysql5Ok returns a tuple with the Mysql5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetMysql5Ok() (*MysqlMysql5, bool) {
	if o == nil || isNil(o.Mysql5) {
    return nil, false
	}
	return o.Mysql5, true
}

// HasMysql5 returns a boolean if a field has been set.
func (o *CloudService) HasMysql5() bool {
	if o != nil && !isNil(o.Mysql5) {
		return true
	}

	return false
}

// SetMysql5 gets a reference to the given MysqlMysql5 and assigns it to the Mysql5 field.
func (o *CloudService) SetMysql5(v MysqlMysql5) {
	o.Mysql5 = &v
}

// GetMysql8 returns the Mysql8 field value if set, zero value otherwise.
func (o *CloudService) GetMysql8() MysqlMysql8 {
	if o == nil || isNil(o.Mysql8) {
		var ret MysqlMysql8
		return ret
	}
	return *o.Mysql8
}

// GetMysql8Ok returns a tuple with the Mysql8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetMysql8Ok() (*MysqlMysql8, bool) {
	if o == nil || isNil(o.Mysql8) {
    return nil, false
	}
	return o.Mysql8, true
}

// HasMysql8 returns a boolean if a field has been set.
func (o *CloudService) HasMysql8() bool {
	if o != nil && !isNil(o.Mysql8) {
		return true
	}

	return false
}

// SetMysql8 gets a reference to the given MysqlMysql8 and assigns it to the Mysql8 field.
func (o *CloudService) SetMysql8(v MysqlMysql8) {
	o.Mysql8 = &v
}

// GetPostgresql15 returns the Postgresql15 field value if set, zero value otherwise.
func (o *CloudService) GetPostgresql15() PostgresqlPostgresql15 {
	if o == nil || isNil(o.Postgresql15) {
		var ret PostgresqlPostgresql15
		return ret
	}
	return *o.Postgresql15
}

// GetPostgresql15Ok returns a tuple with the Postgresql15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetPostgresql15Ok() (*PostgresqlPostgresql15, bool) {
	if o == nil || isNil(o.Postgresql15) {
    return nil, false
	}
	return o.Postgresql15, true
}

// HasPostgresql15 returns a boolean if a field has been set.
func (o *CloudService) HasPostgresql15() bool {
	if o != nil && !isNil(o.Postgresql15) {
		return true
	}

	return false
}

// SetPostgresql15 gets a reference to the given PostgresqlPostgresql15 and assigns it to the Postgresql15 field.
func (o *CloudService) SetPostgresql15(v PostgresqlPostgresql15) {
	o.Postgresql15 = &v
}

// GetPostgresql14 returns the Postgresql14 field value if set, zero value otherwise.
func (o *CloudService) GetPostgresql14() PostgresqlPostgresql14 {
	if o == nil || isNil(o.Postgresql14) {
		var ret PostgresqlPostgresql14
		return ret
	}
	return *o.Postgresql14
}

// GetPostgresql14Ok returns a tuple with the Postgresql14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetPostgresql14Ok() (*PostgresqlPostgresql14, bool) {
	if o == nil || isNil(o.Postgresql14) {
    return nil, false
	}
	return o.Postgresql14, true
}

// HasPostgresql14 returns a boolean if a field has been set.
func (o *CloudService) HasPostgresql14() bool {
	if o != nil && !isNil(o.Postgresql14) {
		return true
	}

	return false
}

// SetPostgresql14 gets a reference to the given PostgresqlPostgresql14 and assigns it to the Postgresql14 field.
func (o *CloudService) SetPostgresql14(v PostgresqlPostgresql14) {
	o.Postgresql14 = &v
}

// GetManageEnabled returns the ManageEnabled field value if set, zero value otherwise.
func (o *CloudService) GetManageEnabled() bool {
	if o == nil || isNil(o.ManageEnabled) {
		var ret bool
		return ret
	}
	return *o.ManageEnabled
}

// GetManageEnabledOk returns a tuple with the ManageEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetManageEnabledOk() (*bool, bool) {
	if o == nil || isNil(o.ManageEnabled) {
    return nil, false
	}
	return o.ManageEnabled, true
}

// HasManageEnabled returns a boolean if a field has been set.
func (o *CloudService) HasManageEnabled() bool {
	if o != nil && !isNil(o.ManageEnabled) {
		return true
	}

	return false
}

// SetManageEnabled gets a reference to the given bool and assigns it to the ManageEnabled field.
func (o *CloudService) SetManageEnabled(v bool) {
	o.ManageEnabled = &v
}

// GetSlug returns the Slug field value if set, zero value otherwise.
func (o *CloudService) GetSlug() string {
	if o == nil || isNil(o.Slug) {
		var ret string
		return ret
	}
	return *o.Slug
}

// GetSlugOk returns a tuple with the Slug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetSlugOk() (*string, bool) {
	if o == nil || isNil(o.Slug) {
    return nil, false
	}
	return o.Slug, true
}

// HasSlug returns a boolean if a field has been set.
func (o *CloudService) HasSlug() bool {
	if o != nil && !isNil(o.Slug) {
		return true
	}

	return false
}

// SetSlug gets a reference to the given string and assigns it to the Slug field.
func (o *CloudService) SetSlug(v string) {
	o.Slug = &v
}

// GetMonitorableResources returns the MonitorableResources field value if set, zero value otherwise.
func (o *CloudService) GetMonitorableResources() []string {
	if o == nil || isNil(o.MonitorableResources) {
		var ret []string
		return ret
	}
	return o.MonitorableResources
}

// GetMonitorableResourcesOk returns a tuple with the MonitorableResources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetMonitorableResourcesOk() ([]string, bool) {
	if o == nil || isNil(o.MonitorableResources) {
    return nil, false
	}
	return o.MonitorableResources, true
}

// HasMonitorableResources returns a boolean if a field has been set.
func (o *CloudService) HasMonitorableResources() bool {
	if o != nil && !isNil(o.MonitorableResources) {
		return true
	}

	return false
}

// SetMonitorableResources gets a reference to the given []string and assigns it to the MonitorableResources field.
func (o *CloudService) SetMonitorableResources(v []string) {
	o.MonitorableResources = v
}

// GetUnblocking returns the Unblocking field value if set, zero value otherwise.
func (o *CloudService) GetUnblocking() bool {
	if o == nil || isNil(o.Unblocking) {
		var ret bool
		return ret
	}
	return *o.Unblocking
}

// GetUnblockingOk returns a tuple with the Unblocking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetUnblockingOk() (*bool, bool) {
	if o == nil || isNil(o.Unblocking) {
    return nil, false
	}
	return o.Unblocking, true
}

// HasUnblocking returns a boolean if a field has been set.
func (o *CloudService) HasUnblocking() bool {
	if o != nil && !isNil(o.Unblocking) {
		return true
	}

	return false
}

// SetUnblocking gets a reference to the given bool and assigns it to the Unblocking field.
func (o *CloudService) SetUnblocking(v bool) {
	o.Unblocking = &v
}

// GetMigrating returns the Migrating field value if set, zero value otherwise.
func (o *CloudService) GetMigrating() bool {
	if o == nil || isNil(o.Migrating) {
		var ret bool
		return ret
	}
	return *o.Migrating
}

// GetMigratingOk returns a tuple with the Migrating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetMigratingOk() (*bool, bool) {
	if o == nil || isNil(o.Migrating) {
    return nil, false
	}
	return o.Migrating, true
}

// HasMigrating returns a boolean if a field has been set.
func (o *CloudService) HasMigrating() bool {
	if o != nil && !isNil(o.Migrating) {
		return true
	}

	return false
}

// SetMigrating gets a reference to the given bool and assigns it to the Migrating field.
func (o *CloudService) SetMigrating(v bool) {
	o.Migrating = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CloudService) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudService) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CloudService) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CloudService) SetRegion(v string) {
	o.Region = &v
}

func (o CloudService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ConfigurationId) {
		toSerialize["configuration_id"] = o.ConfigurationId
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !isNil(o.PriceDay) {
		toSerialize["price_day"] = o.PriceDay
	}
	if !isNil(o.PriceMonth) {
		toSerialize["price_month"] = o.PriceMonth
	}
	if !isNil(o.Mysql5) {
		toSerialize["mysql5"] = o.Mysql5
	}
	if !isNil(o.Mysql8) {
		toSerialize["mysql8"] = o.Mysql8
	}
	if !isNil(o.Postgresql15) {
		toSerialize["postgresql15"] = o.Postgresql15
	}
	if !isNil(o.Postgresql14) {
		toSerialize["postgresql14"] = o.Postgresql14
	}
	if !isNil(o.ManageEnabled) {
		toSerialize["manage_enabled"] = o.ManageEnabled
	}
	if !isNil(o.Slug) {
		toSerialize["slug"] = o.Slug
	}
	if !isNil(o.MonitorableResources) {
		toSerialize["monitorable_resources"] = o.MonitorableResources
	}
	if !isNil(o.Unblocking) {
		toSerialize["unblocking"] = o.Unblocking
	}
	if !isNil(o.Migrating) {
		toSerialize["migrating"] = o.Migrating
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableCloudService struct {
	value *CloudService
	isSet bool
}

func (v NullableCloudService) Get() *CloudService {
	return v.value
}

func (v *NullableCloudService) Set(val *CloudService) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudService) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudService(val *CloudService) *NullableCloudService {
	return &NullableCloudService{value: val, isSet: true}
}

func (v NullableCloudService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


