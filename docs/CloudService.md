# CloudService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ConfigurationId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**PriceDay** | Pointer to **float64** |  | [optional] 
**PriceMonth** | Pointer to **float64** |  | [optional] 
**Mysql5** | Pointer to [**MysqlMysql5**](MysqlMysql5.md) |  | [optional] 
**Mysql8** | Pointer to [**MysqlMysql8**](MysqlMysql8.md) |  | [optional] 
**Postgresql15** | Pointer to [**PostgresqlPostgresql15**](PostgresqlPostgresql15.md) |  | [optional] 
**Postgresql14** | Pointer to [**PostgresqlPostgresql14**](PostgresqlPostgresql14.md) |  | [optional] 
**ManageEnabled** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**MonitorableResources** | Pointer to **[]string** |  | [optional] 
**Unblocking** | Pointer to **bool** |  | [optional] 
**Migrating** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudService

`func NewCloudService() *CloudService`

NewCloudService instantiates a new CloudService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudServiceWithDefaults

`func NewCloudServiceWithDefaults() *CloudService`

NewCloudServiceWithDefaults instantiates a new CloudService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConfigurationId

`func (o *CloudService) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *CloudService) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *CloudService) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *CloudService) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudService) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudService) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudService) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudService) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *CloudService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *CloudService) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudService) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudService) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudService) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CloudService) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CloudService) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CloudService) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CloudService) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPriceDay

`func (o *CloudService) GetPriceDay() float64`

GetPriceDay returns the PriceDay field if non-nil, zero value otherwise.

### GetPriceDayOk

`func (o *CloudService) GetPriceDayOk() (*float64, bool)`

GetPriceDayOk returns a tuple with the PriceDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDay

`func (o *CloudService) SetPriceDay(v float64)`

SetPriceDay sets PriceDay field to given value.

### HasPriceDay

`func (o *CloudService) HasPriceDay() bool`

HasPriceDay returns a boolean if a field has been set.

### GetPriceMonth

`func (o *CloudService) GetPriceMonth() float64`

GetPriceMonth returns the PriceMonth field if non-nil, zero value otherwise.

### GetPriceMonthOk

`func (o *CloudService) GetPriceMonthOk() (*float64, bool)`

GetPriceMonthOk returns a tuple with the PriceMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonth

`func (o *CloudService) SetPriceMonth(v float64)`

SetPriceMonth sets PriceMonth field to given value.

### HasPriceMonth

`func (o *CloudService) HasPriceMonth() bool`

HasPriceMonth returns a boolean if a field has been set.

### GetMysql5

`func (o *CloudService) GetMysql5() MysqlMysql5`

GetMysql5 returns the Mysql5 field if non-nil, zero value otherwise.

### GetMysql5Ok

`func (o *CloudService) GetMysql5Ok() (*MysqlMysql5, bool)`

GetMysql5Ok returns a tuple with the Mysql5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql5

`func (o *CloudService) SetMysql5(v MysqlMysql5)`

SetMysql5 sets Mysql5 field to given value.

### HasMysql5

`func (o *CloudService) HasMysql5() bool`

HasMysql5 returns a boolean if a field has been set.

### GetMysql8

`func (o *CloudService) GetMysql8() MysqlMysql8`

GetMysql8 returns the Mysql8 field if non-nil, zero value otherwise.

### GetMysql8Ok

`func (o *CloudService) GetMysql8Ok() (*MysqlMysql8, bool)`

GetMysql8Ok returns a tuple with the Mysql8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql8

`func (o *CloudService) SetMysql8(v MysqlMysql8)`

SetMysql8 sets Mysql8 field to given value.

### HasMysql8

`func (o *CloudService) HasMysql8() bool`

HasMysql8 returns a boolean if a field has been set.

### GetPostgresql15

`func (o *CloudService) GetPostgresql15() PostgresqlPostgresql15`

GetPostgresql15 returns the Postgresql15 field if non-nil, zero value otherwise.

### GetPostgresql15Ok

`func (o *CloudService) GetPostgresql15Ok() (*PostgresqlPostgresql15, bool)`

GetPostgresql15Ok returns a tuple with the Postgresql15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresql15

`func (o *CloudService) SetPostgresql15(v PostgresqlPostgresql15)`

SetPostgresql15 sets Postgresql15 field to given value.

### HasPostgresql15

`func (o *CloudService) HasPostgresql15() bool`

HasPostgresql15 returns a boolean if a field has been set.

### GetPostgresql14

`func (o *CloudService) GetPostgresql14() PostgresqlPostgresql14`

GetPostgresql14 returns the Postgresql14 field if non-nil, zero value otherwise.

### GetPostgresql14Ok

`func (o *CloudService) GetPostgresql14Ok() (*PostgresqlPostgresql14, bool)`

GetPostgresql14Ok returns a tuple with the Postgresql14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresql14

`func (o *CloudService) SetPostgresql14(v PostgresqlPostgresql14)`

SetPostgresql14 sets Postgresql14 field to given value.

### HasPostgresql14

`func (o *CloudService) HasPostgresql14() bool`

HasPostgresql14 returns a boolean if a field has been set.

### GetManageEnabled

`func (o *CloudService) GetManageEnabled() bool`

GetManageEnabled returns the ManageEnabled field if non-nil, zero value otherwise.

### GetManageEnabledOk

`func (o *CloudService) GetManageEnabledOk() (*bool, bool)`

GetManageEnabledOk returns a tuple with the ManageEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManageEnabled

`func (o *CloudService) SetManageEnabled(v bool)`

SetManageEnabled sets ManageEnabled field to given value.

### HasManageEnabled

`func (o *CloudService) HasManageEnabled() bool`

HasManageEnabled returns a boolean if a field has been set.

### GetSlug

`func (o *CloudService) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CloudService) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CloudService) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CloudService) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetMonitorableResources

`func (o *CloudService) GetMonitorableResources() []string`

GetMonitorableResources returns the MonitorableResources field if non-nil, zero value otherwise.

### GetMonitorableResourcesOk

`func (o *CloudService) GetMonitorableResourcesOk() (*[]string, bool)`

GetMonitorableResourcesOk returns a tuple with the MonitorableResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorableResources

`func (o *CloudService) SetMonitorableResources(v []string)`

SetMonitorableResources sets MonitorableResources field to given value.

### HasMonitorableResources

`func (o *CloudService) HasMonitorableResources() bool`

HasMonitorableResources returns a boolean if a field has been set.

### GetUnblocking

`func (o *CloudService) GetUnblocking() bool`

GetUnblocking returns the Unblocking field if non-nil, zero value otherwise.

### GetUnblockingOk

`func (o *CloudService) GetUnblockingOk() (*bool, bool)`

GetUnblockingOk returns a tuple with the Unblocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnblocking

`func (o *CloudService) SetUnblocking(v bool)`

SetUnblocking sets Unblocking field to given value.

### HasUnblocking

`func (o *CloudService) HasUnblocking() bool`

HasUnblocking returns a boolean if a field has been set.

### GetMigrating

`func (o *CloudService) GetMigrating() bool`

GetMigrating returns the Migrating field if non-nil, zero value otherwise.

### GetMigratingOk

`func (o *CloudService) GetMigratingOk() (*bool, bool)`

GetMigratingOk returns a tuple with the Migrating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrating

`func (o *CloudService) SetMigrating(v bool)`

SetMigrating sets Migrating field to given value.

### HasMigrating

`func (o *CloudService) HasMigrating() bool`

HasMigrating returns a boolean if a field has been set.

### GetRegion

`func (o *CloudService) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudService) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudService) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudService) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


