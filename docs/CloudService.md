# CloudService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ConfigurationId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDetails** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**PriceDay** | Pointer to **float64** |  | [optional] 
**PriceMonth** | Pointer to **float64** |  | [optional] 
**Mysql5** | Pointer to [**MysqlMysql5**](MysqlMysql5.md) |  | [optional] 
**Mysql8** | Pointer to [**MysqlMysql8**](MysqlMysql8.md) |  | [optional] 
**Mysql84** | Pointer to [**MysqlMysql84**](MysqlMysql84.md) |  | [optional] 
**Postgresql15** | Pointer to [**PostgresqlPostgresql15**](PostgresqlPostgresql15.md) |  | [optional] 
**Postgresql14** | Pointer to [**PostgresqlPostgresql14**](PostgresqlPostgresql14.md) |  | [optional] 
**Postgresql164** | Pointer to [**PostgresqlPostgresql164**](PostgresqlPostgresql164.md) |  | [optional] 
**S3** | Pointer to [**S3S3**](S3S3.md) |  | [optional] 
**Cdn** | Pointer to [**CdnCdn**](CdnCdn.md) |  | [optional] 
**ManageEnabled** | Pointer to **bool** |  | [optional] 
**Slug** | Pointer to **string** |  | [optional] 
**MonitorableResources** | Pointer to **[]string** |  | [optional] 
**Unblocking** | Pointer to **bool** |  | [optional] 
**Migrating** | Pointer to **bool** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Restoring** | Pointer to **bool** |  | [optional] 

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

### GetStatusDetails

`func (o *CloudService) GetStatusDetails() string`

GetStatusDetails returns the StatusDetails field if non-nil, zero value otherwise.

### GetStatusDetailsOk

`func (o *CloudService) GetStatusDetailsOk() (*string, bool)`

GetStatusDetailsOk returns a tuple with the StatusDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDetails

`func (o *CloudService) SetStatusDetails(v string)`

SetStatusDetails sets StatusDetails field to given value.

### HasStatusDetails

`func (o *CloudService) HasStatusDetails() bool`

HasStatusDetails returns a boolean if a field has been set.

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

### GetMysql84

`func (o *CloudService) GetMysql84() MysqlMysql84`

GetMysql84 returns the Mysql84 field if non-nil, zero value otherwise.

### GetMysql84Ok

`func (o *CloudService) GetMysql84Ok() (*MysqlMysql84, bool)`

GetMysql84Ok returns a tuple with the Mysql84 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql84

`func (o *CloudService) SetMysql84(v MysqlMysql84)`

SetMysql84 sets Mysql84 field to given value.

### HasMysql84

`func (o *CloudService) HasMysql84() bool`

HasMysql84 returns a boolean if a field has been set.

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

### GetPostgresql164

`func (o *CloudService) GetPostgresql164() PostgresqlPostgresql164`

GetPostgresql164 returns the Postgresql164 field if non-nil, zero value otherwise.

### GetPostgresql164Ok

`func (o *CloudService) GetPostgresql164Ok() (*PostgresqlPostgresql164, bool)`

GetPostgresql164Ok returns a tuple with the Postgresql164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresql164

`func (o *CloudService) SetPostgresql164(v PostgresqlPostgresql164)`

SetPostgresql164 sets Postgresql164 field to given value.

### HasPostgresql164

`func (o *CloudService) HasPostgresql164() bool`

HasPostgresql164 returns a boolean if a field has been set.

### GetS3

`func (o *CloudService) GetS3() S3S3`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *CloudService) GetS3Ok() (*S3S3, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *CloudService) SetS3(v S3S3)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *CloudService) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetCdn

`func (o *CloudService) GetCdn() CdnCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *CloudService) GetCdnOk() (*CdnCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *CloudService) SetCdn(v CdnCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *CloudService) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

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

### GetType

`func (o *CloudService) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudService) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudService) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudService) HasType() bool`

HasType returns a boolean if a field has been set.

### GetRestoring

`func (o *CloudService) GetRestoring() bool`

GetRestoring returns the Restoring field if non-nil, zero value otherwise.

### GetRestoringOk

`func (o *CloudService) GetRestoringOk() (*bool, bool)`

GetRestoringOk returns a tuple with the Restoring field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoring

`func (o *CloudService) SetRestoring(v bool)`

SetRestoring sets Restoring field to given value.

### HasRestoring

`func (o *CloudService) HasRestoring() bool`

HasRestoring returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


