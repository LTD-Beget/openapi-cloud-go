# PostgresqlPostgresql14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**PostgresqlPostgresql14Configuration**](PostgresqlPostgresql14Configuration.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**AddressInfo** | Pointer to [**StructuresAddressInfo**](StructuresAddressInfo.md) |  | [optional] 
**WebuiUrl** | Pointer to **string** |  | [optional] 
**DiskUsed** | Pointer to **string** |  | [optional] 
**DiskLeft** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**WebuiInstalling** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostgresqlPostgresql14

`func NewPostgresqlPostgresql14() *PostgresqlPostgresql14`

NewPostgresqlPostgresql14 instantiates a new PostgresqlPostgresql14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPostgresql14WithDefaults

`func NewPostgresqlPostgresql14WithDefaults() *PostgresqlPostgresql14`

NewPostgresqlPostgresql14WithDefaults instantiates a new PostgresqlPostgresql14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *PostgresqlPostgresql14) GetConfiguration() PostgresqlPostgresql14Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *PostgresqlPostgresql14) GetConfigurationOk() (*PostgresqlPostgresql14Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *PostgresqlPostgresql14) SetConfiguration(v PostgresqlPostgresql14Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *PostgresqlPostgresql14) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetHost

`func (o *PostgresqlPostgresql14) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *PostgresqlPostgresql14) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *PostgresqlPostgresql14) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *PostgresqlPostgresql14) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *PostgresqlPostgresql14) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *PostgresqlPostgresql14) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *PostgresqlPostgresql14) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *PostgresqlPostgresql14) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAddressInfo

`func (o *PostgresqlPostgresql14) GetAddressInfo() StructuresAddressInfo`

GetAddressInfo returns the AddressInfo field if non-nil, zero value otherwise.

### GetAddressInfoOk

`func (o *PostgresqlPostgresql14) GetAddressInfoOk() (*StructuresAddressInfo, bool)`

GetAddressInfoOk returns a tuple with the AddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInfo

`func (o *PostgresqlPostgresql14) SetAddressInfo(v StructuresAddressInfo)`

SetAddressInfo sets AddressInfo field to given value.

### HasAddressInfo

`func (o *PostgresqlPostgresql14) HasAddressInfo() bool`

HasAddressInfo returns a boolean if a field has been set.

### GetWebuiUrl

`func (o *PostgresqlPostgresql14) GetWebuiUrl() string`

GetWebuiUrl returns the WebuiUrl field if non-nil, zero value otherwise.

### GetWebuiUrlOk

`func (o *PostgresqlPostgresql14) GetWebuiUrlOk() (*string, bool)`

GetWebuiUrlOk returns a tuple with the WebuiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiUrl

`func (o *PostgresqlPostgresql14) SetWebuiUrl(v string)`

SetWebuiUrl sets WebuiUrl field to given value.

### HasWebuiUrl

`func (o *PostgresqlPostgresql14) HasWebuiUrl() bool`

HasWebuiUrl returns a boolean if a field has been set.

### GetDiskUsed

`func (o *PostgresqlPostgresql14) GetDiskUsed() string`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *PostgresqlPostgresql14) GetDiskUsedOk() (*string, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *PostgresqlPostgresql14) SetDiskUsed(v string)`

SetDiskUsed sets DiskUsed field to given value.

### HasDiskUsed

`func (o *PostgresqlPostgresql14) HasDiskUsed() bool`

HasDiskUsed returns a boolean if a field has been set.

### GetDiskLeft

`func (o *PostgresqlPostgresql14) GetDiskLeft() string`

GetDiskLeft returns the DiskLeft field if non-nil, zero value otherwise.

### GetDiskLeftOk

`func (o *PostgresqlPostgresql14) GetDiskLeftOk() (*string, bool)`

GetDiskLeftOk returns a tuple with the DiskLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLeft

`func (o *PostgresqlPostgresql14) SetDiskLeft(v string)`

SetDiskLeft sets DiskLeft field to given value.

### HasDiskLeft

`func (o *PostgresqlPostgresql14) HasDiskLeft() bool`

HasDiskLeft returns a boolean if a field has been set.

### GetReadOnly

`func (o *PostgresqlPostgresql14) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *PostgresqlPostgresql14) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *PostgresqlPostgresql14) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *PostgresqlPostgresql14) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetWebuiInstalling

`func (o *PostgresqlPostgresql14) GetWebuiInstalling() bool`

GetWebuiInstalling returns the WebuiInstalling field if non-nil, zero value otherwise.

### GetWebuiInstallingOk

`func (o *PostgresqlPostgresql14) GetWebuiInstallingOk() (*bool, bool)`

GetWebuiInstallingOk returns a tuple with the WebuiInstalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebuiInstalling

`func (o *PostgresqlPostgresql14) SetWebuiInstalling(v bool)`

SetWebuiInstalling sets WebuiInstalling field to given value.

### HasWebuiInstalling

`func (o *PostgresqlPostgresql14) HasWebuiInstalling() bool`

HasWebuiInstalling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


