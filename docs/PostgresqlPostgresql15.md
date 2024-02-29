# PostgresqlPostgresql15

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**PostgresqlPostgresql15Configuration**](PostgresqlPostgresql15Configuration.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**AddressInfo** | Pointer to [**StructuresAddressInfo**](StructuresAddressInfo.md) |  | [optional] 
**WebuiUrl** | Pointer to **string** |  | [optional] 
**DiskUsed** | Pointer to **string** |  | [optional] 
**DiskLeft** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**WebuiInstalling** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostgresqlPostgresql15

`func NewPostgresqlPostgresql15() *PostgresqlPostgresql15`

NewPostgresqlPostgresql15 instantiates a new PostgresqlPostgresql15 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPostgresql15WithDefaults

`func NewPostgresqlPostgresql15WithDefaults() *PostgresqlPostgresql15`

NewPostgresqlPostgresql15WithDefaults instantiates a new PostgresqlPostgresql15 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *PostgresqlPostgresql15) GetConfiguration() PostgresqlPostgresql15Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostgresqlPostgresql15) GetConfigurationOk() (*PostgresqlPostgresql15Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostgresqlPostgresql15) SetConfiguration(v PostgresqlPostgresql15Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostgresqlPostgresql15) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetHost

`func (o *PostgresqlPostgresql15) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PostgresqlPostgresql15) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PostgresqlPostgresql15) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PostgresqlPostgresql15) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *PostgresqlPostgresql15) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostgresqlPostgresql15) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostgresqlPostgresql15) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PostgresqlPostgresql15) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAddressInfo

`func (o *PostgresqlPostgresql15) GetAddressInfo() StructuresAddressInfo`

GetAddressInfo returns the AddressInfo field if non-nil, zero value otherwise.

### GetAddressInfoOk

`func (o *PostgresqlPostgresql15) GetAddressInfoOk() (*StructuresAddressInfo, bool)`

GetAddressInfoOk returns a tuple with the AddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInfo

`func (o *PostgresqlPostgresql15) SetAddressInfo(v StructuresAddressInfo)`

SetAddressInfo sets AddressInfo field to given value.

### HasAddressInfo

`func (o *PostgresqlPostgresql15) HasAddressInfo() bool`

HasAddressInfo returns a boolean if a field has been set.

### GetWebuiUrl

`func (o *PostgresqlPostgresql15) GetWebuiUrl() string`

GetWebuiUrl returns the WebuiUrl field if non-nil, zero value otherwise.

### GetWebuiUrlOk

`func (o *PostgresqlPostgresql15) GetWebuiUrlOk() (*string, bool)`

GetWebuiUrlOk returns a tuple with the WebuiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiUrl

`func (o *PostgresqlPostgresql15) SetWebuiUrl(v string)`

SetWebuiUrl sets WebuiUrl field to given value.

### HasWebuiUrl

`func (o *PostgresqlPostgresql15) HasWebuiUrl() bool`

HasWebuiUrl returns a boolean if a field has been set.

### GetDiskUsed

`func (o *PostgresqlPostgresql15) GetDiskUsed() string`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *PostgresqlPostgresql15) GetDiskUsedOk() (*string, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *PostgresqlPostgresql15) SetDiskUsed(v string)`

SetDiskUsed sets DiskUsed field to given value.

### HasDiskUsed

`func (o *PostgresqlPostgresql15) HasDiskUsed() bool`

HasDiskUsed returns a boolean if a field has been set.

### GetDiskLeft

`func (o *PostgresqlPostgresql15) GetDiskLeft() string`

GetDiskLeft returns the DiskLeft field if non-nil, zero value otherwise.

### GetDiskLeftOk

`func (o *PostgresqlPostgresql15) GetDiskLeftOk() (*string, bool)`

GetDiskLeftOk returns a tuple with the DiskLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLeft

`func (o *PostgresqlPostgresql15) SetDiskLeft(v string)`

SetDiskLeft sets DiskLeft field to given value.

### HasDiskLeft

`func (o *PostgresqlPostgresql15) HasDiskLeft() bool`

HasDiskLeft returns a boolean if a field has been set.

### GetReadOnly

`func (o *PostgresqlPostgresql15) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PostgresqlPostgresql15) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PostgresqlPostgresql15) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PostgresqlPostgresql15) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetWebuiInstalling

`func (o *PostgresqlPostgresql15) GetWebuiInstalling() bool`

GetWebuiInstalling returns the WebuiInstalling field if non-nil, zero value otherwise.

### GetWebuiInstallingOk

`func (o *PostgresqlPostgresql15) GetWebuiInstallingOk() (*bool, bool)`

GetWebuiInstallingOk returns a tuple with the WebuiInstalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiInstalling

`func (o *PostgresqlPostgresql15) SetWebuiInstalling(v bool)`

SetWebuiInstalling sets WebuiInstalling field to given value.

### HasWebuiInstalling

`func (o *PostgresqlPostgresql15) HasWebuiInstalling() bool`

HasWebuiInstalling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


