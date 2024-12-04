# MysqlMysql84

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**MysqlMysql84Configuration**](MysqlMysql84Configuration.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**AddressInfo** | Pointer to [**StructuresAddressInfo**](StructuresAddressInfo.md) |  | [optional] 
**PmaUrl** | Pointer to **string** |  | [optional] 
**DiskUsed** | Pointer to **string** |  | [optional] 
**DiskLeft** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**PmaInstalling** | Pointer to **bool** |  | [optional] 

## Methods

### NewMysqlMysql84

`func NewMysqlMysql84() *MysqlMysql84`

NewMysqlMysql84 instantiates a new MysqlMysql84 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlMysql84WithDefaults

`func NewMysqlMysql84WithDefaults() *MysqlMysql84`

NewMysqlMysql84WithDefaults instantiates a new MysqlMysql84 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *MysqlMysql84) GetConfiguration() MysqlMysql84Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *MysqlMysql84) GetConfigurationOk() (*MysqlMysql84Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *MysqlMysql84) SetConfiguration(v MysqlMysql84Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *MysqlMysql84) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetHost

`func (o *MysqlMysql84) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MysqlMysql84) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MysqlMysql84) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MysqlMysql84) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *MysqlMysql84) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MysqlMysql84) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MysqlMysql84) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MysqlMysql84) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAddressInfo

`func (o *MysqlMysql84) GetAddressInfo() StructuresAddressInfo`

GetAddressInfo returns the AddressInfo field if non-nil, zero value otherwise.

### GetAddressInfoOk

`func (o *MysqlMysql84) GetAddressInfoOk() (*StructuresAddressInfo, bool)`

GetAddressInfoOk returns a tuple with the AddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInfo

`func (o *MysqlMysql84) SetAddressInfo(v StructuresAddressInfo)`

SetAddressInfo sets AddressInfo field to given value.

### HasAddressInfo

`func (o *MysqlMysql84) HasAddressInfo() bool`

HasAddressInfo returns a boolean if a field has been set.

### GetPmaUrl

`func (o *MysqlMysql84) GetPmaUrl() string`

GetPmaUrl returns the PmaUrl field if non-nil, zero value otherwise.

### GetPmaUrlOk

`func (o *MysqlMysql84) GetPmaUrlOk() (*string, bool)`

GetPmaUrlOk returns a tuple with the PmaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmaUrl

`func (o *MysqlMysql84) SetPmaUrl(v string)`

SetPmaUrl sets PmaUrl field to given value.

### HasPmaUrl

`func (o *MysqlMysql84) HasPmaUrl() bool`

HasPmaUrl returns a boolean if a field has been set.

### GetDiskUsed

`func (o *MysqlMysql84) GetDiskUsed() string`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *MysqlMysql84) GetDiskUsedOk() (*string, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *MysqlMysql84) SetDiskUsed(v string)`

SetDiskUsed sets DiskUsed field to given value.

### HasDiskUsed

`func (o *MysqlMysql84) HasDiskUsed() bool`

HasDiskUsed returns a boolean if a field has been set.

### GetDiskLeft

`func (o *MysqlMysql84) GetDiskLeft() string`

GetDiskLeft returns the DiskLeft field if non-nil, zero value otherwise.

### GetDiskLeftOk

`func (o *MysqlMysql84) GetDiskLeftOk() (*string, bool)`

GetDiskLeftOk returns a tuple with the DiskLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLeft

`func (o *MysqlMysql84) SetDiskLeft(v string)`

SetDiskLeft sets DiskLeft field to given value.

### HasDiskLeft

`func (o *MysqlMysql84) HasDiskLeft() bool`

HasDiskLeft returns a boolean if a field has been set.

### GetReadOnly

`func (o *MysqlMysql84) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MysqlMysql84) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *MysqlMysql84) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *MysqlMysql84) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetPmaInstalling

`func (o *MysqlMysql84) GetPmaInstalling() bool`

GetPmaInstalling returns the PmaInstalling field if non-nil, zero value otherwise.

### GetPmaInstallingOk

`func (o *MysqlMysql84) GetPmaInstallingOk() (*bool, bool)`

GetPmaInstallingOk returns a tuple with the PmaInstalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmaInstalling

`func (o *MysqlMysql84) SetPmaInstalling(v bool)`

SetPmaInstalling sets PmaInstalling field to given value.

### HasPmaInstalling

`func (o *MysqlMysql84) HasPmaInstalling() bool`

HasPmaInstalling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


