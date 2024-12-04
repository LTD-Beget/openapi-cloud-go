# CloudServiceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PriceDay** | Pointer to **float64** |  | [optional] 
**PriceMonth** | Pointer to **float64** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Mysql5** | Pointer to [**MysqlMysql5Configuration**](MysqlMysql5Configuration.md) |  | [optional] 
**Mysql8** | Pointer to [**MysqlMysql8Configuration**](MysqlMysql8Configuration.md) |  | [optional] 
**Mysql84** | Pointer to [**MysqlMysql84Configuration**](MysqlMysql84Configuration.md) |  | [optional] 
**Postgresql15** | Pointer to [**PostgresqlPostgresql15Configuration**](PostgresqlPostgresql15Configuration.md) |  | [optional] 
**Postgresql14** | Pointer to [**PostgresqlPostgresql14Configuration**](PostgresqlPostgresql14Configuration.md) |  | [optional] 
**Postgresql164** | Pointer to [**PostgresqlPostgresql164Configuration**](PostgresqlPostgresql164Configuration.md) |  | [optional] 
**S3** | Pointer to [**S3S3Configuration**](S3S3Configuration.md) |  | [optional] 

## Methods

### NewCloudServiceConfiguration

`func NewCloudServiceConfiguration() *CloudServiceConfiguration`

NewCloudServiceConfiguration instantiates a new CloudServiceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudServiceConfigurationWithDefaults

`func NewCloudServiceConfigurationWithDefaults() *CloudServiceConfiguration`

NewCloudServiceConfigurationWithDefaults instantiates a new CloudServiceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CloudServiceConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudServiceConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudServiceConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudServiceConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriceDay

`func (o *CloudServiceConfiguration) GetPriceDay() float64`

GetPriceDay returns the PriceDay field if non-nil, zero value otherwise.

### GetPriceDayOk

`func (o *CloudServiceConfiguration) GetPriceDayOk() (*float64, bool)`

GetPriceDayOk returns a tuple with the PriceDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceDay

`func (o *CloudServiceConfiguration) SetPriceDay(v float64)`

SetPriceDay sets PriceDay field to given value.

### HasPriceDay

`func (o *CloudServiceConfiguration) HasPriceDay() bool`

HasPriceDay returns a boolean if a field has been set.

### GetPriceMonth

`func (o *CloudServiceConfiguration) GetPriceMonth() float64`

GetPriceMonth returns the PriceMonth field if non-nil, zero value otherwise.

### GetPriceMonthOk

`func (o *CloudServiceConfiguration) GetPriceMonthOk() (*float64, bool)`

GetPriceMonthOk returns a tuple with the PriceMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMonth

`func (o *CloudServiceConfiguration) SetPriceMonth(v float64)`

SetPriceMonth sets PriceMonth field to given value.

### HasPriceMonth

`func (o *CloudServiceConfiguration) HasPriceMonth() bool`

HasPriceMonth returns a boolean if a field has been set.

### GetRegion

`func (o *CloudServiceConfiguration) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CloudServiceConfiguration) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CloudServiceConfiguration) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CloudServiceConfiguration) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *CloudServiceConfiguration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CloudServiceConfiguration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CloudServiceConfiguration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CloudServiceConfiguration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMysql5

`func (o *CloudServiceConfiguration) GetMysql5() MysqlMysql5Configuration`

GetMysql5 returns the Mysql5 field if non-nil, zero value otherwise.

### GetMysql5Ok

`func (o *CloudServiceConfiguration) GetMysql5Ok() (*MysqlMysql5Configuration, bool)`

GetMysql5Ok returns a tuple with the Mysql5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql5

`func (o *CloudServiceConfiguration) SetMysql5(v MysqlMysql5Configuration)`

SetMysql5 sets Mysql5 field to given value.

### HasMysql5

`func (o *CloudServiceConfiguration) HasMysql5() bool`

HasMysql5 returns a boolean if a field has been set.

### GetMysql8

`func (o *CloudServiceConfiguration) GetMysql8() MysqlMysql8Configuration`

GetMysql8 returns the Mysql8 field if non-nil, zero value otherwise.

### GetMysql8Ok

`func (o *CloudServiceConfiguration) GetMysql8Ok() (*MysqlMysql8Configuration, bool)`

GetMysql8Ok returns a tuple with the Mysql8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql8

`func (o *CloudServiceConfiguration) SetMysql8(v MysqlMysql8Configuration)`

SetMysql8 sets Mysql8 field to given value.

### HasMysql8

`func (o *CloudServiceConfiguration) HasMysql8() bool`

HasMysql8 returns a boolean if a field has been set.

### GetMysql84

`func (o *CloudServiceConfiguration) GetMysql84() MysqlMysql84Configuration`

GetMysql84 returns the Mysql84 field if non-nil, zero value otherwise.

### GetMysql84Ok

`func (o *CloudServiceConfiguration) GetMysql84Ok() (*MysqlMysql84Configuration, bool)`

GetMysql84Ok returns a tuple with the Mysql84 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql84

`func (o *CloudServiceConfiguration) SetMysql84(v MysqlMysql84Configuration)`

SetMysql84 sets Mysql84 field to given value.

### HasMysql84

`func (o *CloudServiceConfiguration) HasMysql84() bool`

HasMysql84 returns a boolean if a field has been set.

### GetPostgresql15

`func (o *CloudServiceConfiguration) GetPostgresql15() PostgresqlPostgresql15Configuration`

GetPostgresql15 returns the Postgresql15 field if non-nil, zero value otherwise.

### GetPostgresql15Ok

`func (o *CloudServiceConfiguration) GetPostgresql15Ok() (*PostgresqlPostgresql15Configuration, bool)`

GetPostgresql15Ok returns a tuple with the Postgresql15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresql15

`func (o *CloudServiceConfiguration) SetPostgresql15(v PostgresqlPostgresql15Configuration)`

SetPostgresql15 sets Postgresql15 field to given value.

### HasPostgresql15

`func (o *CloudServiceConfiguration) HasPostgresql15() bool`

HasPostgresql15 returns a boolean if a field has been set.

### GetPostgresql14

`func (o *CloudServiceConfiguration) GetPostgresql14() PostgresqlPostgresql14Configuration`

GetPostgresql14 returns the Postgresql14 field if non-nil, zero value otherwise.

### GetPostgresql14Ok

`func (o *CloudServiceConfiguration) GetPostgresql14Ok() (*PostgresqlPostgresql14Configuration, bool)`

GetPostgresql14Ok returns a tuple with the Postgresql14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresql14

`func (o *CloudServiceConfiguration) SetPostgresql14(v PostgresqlPostgresql14Configuration)`

SetPostgresql14 sets Postgresql14 field to given value.

### HasPostgresql14

`func (o *CloudServiceConfiguration) HasPostgresql14() bool`

HasPostgresql14 returns a boolean if a field has been set.

### GetPostgresql164

`func (o *CloudServiceConfiguration) GetPostgresql164() PostgresqlPostgresql164Configuration`

GetPostgresql164 returns the Postgresql164 field if non-nil, zero value otherwise.

### GetPostgresql164Ok

`func (o *CloudServiceConfiguration) GetPostgresql164Ok() (*PostgresqlPostgresql164Configuration, bool)`

GetPostgresql164Ok returns a tuple with the Postgresql164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresql164

`func (o *CloudServiceConfiguration) SetPostgresql164(v PostgresqlPostgresql164Configuration)`

SetPostgresql164 sets Postgresql164 field to given value.

### HasPostgresql164

`func (o *CloudServiceConfiguration) HasPostgresql164() bool`

HasPostgresql164 returns a boolean if a field has been set.

### GetS3

`func (o *CloudServiceConfiguration) GetS3() S3S3Configuration`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *CloudServiceConfiguration) GetS3Ok() (*S3S3Configuration, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *CloudServiceConfiguration) SetS3(v S3S3Configuration)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *CloudServiceConfiguration) HasS3() bool`

HasS3 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


