# PostgresqlPostgresql164

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**PostgresqlPostgresql164Configuration**](PostgresqlPostgresql164Configuration.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**AddressInfo** | Pointer to [**StructuresAddressInfo**](StructuresAddressInfo.md) |  | [optional] 
**WebuiUrl** | Pointer to **string** |  | [optional] 
**DiskUsed** | Pointer to **string** |  | [optional] 
**DiskLeft** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**WebuiInstalling** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostgresqlPostgresql164

`func NewPostgresqlPostgresql164() *PostgresqlPostgresql164`

NewPostgresqlPostgresql164 instantiates a new PostgresqlPostgresql164 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPostgresql164WithDefaults

`func NewPostgresqlPostgresql164WithDefaults() *PostgresqlPostgresql164`

NewPostgresqlPostgresql164WithDefaults instantiates a new PostgresqlPostgresql164 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *PostgresqlPostgresql164) GetConfiguration() PostgresqlPostgresql164Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostgresqlPostgresql164) GetConfigurationOk() (*PostgresqlPostgresql164Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostgresqlPostgresql164) SetConfiguration(v PostgresqlPostgresql164Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostgresqlPostgresql164) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetHost

`func (o *PostgresqlPostgresql164) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PostgresqlPostgresql164) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PostgresqlPostgresql164) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PostgresqlPostgresql164) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *PostgresqlPostgresql164) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostgresqlPostgresql164) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostgresqlPostgresql164) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PostgresqlPostgresql164) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAddressInfo

`func (o *PostgresqlPostgresql164) GetAddressInfo() StructuresAddressInfo`

GetAddressInfo returns the AddressInfo field if non-nil, zero value otherwise.

### GetAddressInfoOk

`func (o *PostgresqlPostgresql164) GetAddressInfoOk() (*StructuresAddressInfo, bool)`

GetAddressInfoOk returns a tuple with the AddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInfo

`func (o *PostgresqlPostgresql164) SetAddressInfo(v StructuresAddressInfo)`

SetAddressInfo sets AddressInfo field to given value.

### HasAddressInfo

`func (o *PostgresqlPostgresql164) HasAddressInfo() bool`

HasAddressInfo returns a boolean if a field has been set.

### GetWebuiUrl

`func (o *PostgresqlPostgresql164) GetWebuiUrl() string`

GetWebuiUrl returns the WebuiUrl field if non-nil, zero value otherwise.

### GetWebuiUrlOk

`func (o *PostgresqlPostgresql164) GetWebuiUrlOk() (*string, bool)`

GetWebuiUrlOk returns a tuple with the WebuiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiUrl

`func (o *PostgresqlPostgresql164) SetWebuiUrl(v string)`

SetWebuiUrl sets WebuiUrl field to given value.

### HasWebuiUrl

`func (o *PostgresqlPostgresql164) HasWebuiUrl() bool`

HasWebuiUrl returns a boolean if a field has been set.

### GetDiskUsed

`func (o *PostgresqlPostgresql164) GetDiskUsed() string`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *PostgresqlPostgresql164) GetDiskUsedOk() (*string, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *PostgresqlPostgresql164) SetDiskUsed(v string)`

SetDiskUsed sets DiskUsed field to given value.

### HasDiskUsed

`func (o *PostgresqlPostgresql164) HasDiskUsed() bool`

HasDiskUsed returns a boolean if a field has been set.

### GetDiskLeft

`func (o *PostgresqlPostgresql164) GetDiskLeft() string`

GetDiskLeft returns the DiskLeft field if non-nil, zero value otherwise.

### GetDiskLeftOk

`func (o *PostgresqlPostgresql164) GetDiskLeftOk() (*string, bool)`

GetDiskLeftOk returns a tuple with the DiskLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLeft

`func (o *PostgresqlPostgresql164) SetDiskLeft(v string)`

SetDiskLeft sets DiskLeft field to given value.

### HasDiskLeft

`func (o *PostgresqlPostgresql164) HasDiskLeft() bool`

HasDiskLeft returns a boolean if a field has been set.

### GetReadOnly

`func (o *PostgresqlPostgresql164) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PostgresqlPostgresql164) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PostgresqlPostgresql164) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PostgresqlPostgresql164) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetWebuiInstalling

`func (o *PostgresqlPostgresql164) GetWebuiInstalling() bool`

GetWebuiInstalling returns the WebuiInstalling field if non-nil, zero value otherwise.

### GetWebuiInstallingOk

`func (o *PostgresqlPostgresql164) GetWebuiInstallingOk() (*bool, bool)`

GetWebuiInstallingOk returns a tuple with the WebuiInstalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiInstalling

`func (o *PostgresqlPostgresql164) SetWebuiInstalling(v bool)`

SetWebuiInstalling sets WebuiInstalling field to given value.

### HasWebuiInstalling

`func (o *PostgresqlPostgresql164) HasWebuiInstalling() bool`

HasWebuiInstalling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


