# MysqlMysql8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**MysqlMysql8Configuration**](MysqlMysql8Configuration.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**AddressInfo** | Pointer to [**StructuresAddressInfo**](StructuresAddressInfo.md) |  | [optional] 
**PmaUrl** | Pointer to **string** |  | [optional] 
**DiskUsed** | Pointer to **string** |  | [optional] 
**DiskLeft** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**PmaInstalling** | Pointer to **bool** |  | [optional] 

## Methods

### NewMysqlMysql8

`func NewMysqlMysql8() *MysqlMysql8`

NewMysqlMysql8 instantiates a new MysqlMysql8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlMysql8WithDefaults

`func NewMysqlMysql8WithDefaults() *MysqlMysql8`

NewMysqlMysql8WithDefaults instantiates a new MysqlMysql8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *MysqlMysql8) GetConfiguration() MysqlMysql8Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *MysqlMysql8) GetConfigurationOk() (*MysqlMysql8Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *MysqlMysql8) SetConfiguration(v MysqlMysql8Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *MysqlMysql8) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetHost

`func (o *MysqlMysql8) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MysqlMysql8) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MysqlMysql8) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MysqlMysql8) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *MysqlMysql8) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MysqlMysql8) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MysqlMysql8) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MysqlMysql8) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAddressInfo

`func (o *MysqlMysql8) GetAddressInfo() StructuresAddressInfo`

GetAddressInfo returns the AddressInfo field if non-nil, zero value otherwise.

### GetAddressInfoOk

`func (o *MysqlMysql8) GetAddressInfoOk() (*StructuresAddressInfo, bool)`

GetAddressInfoOk returns a tuple with the AddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInfo

`func (o *MysqlMysql8) SetAddressInfo(v StructuresAddressInfo)`

SetAddressInfo sets AddressInfo field to given value.

### HasAddressInfo

`func (o *MysqlMysql8) HasAddressInfo() bool`

HasAddressInfo returns a boolean if a field has been set.

### GetPmaUrl

`func (o *MysqlMysql8) GetPmaUrl() string`

GetPmaUrl returns the PmaUrl field if non-nil, zero value otherwise.

### GetPmaUrlOk

`func (o *MysqlMysql8) GetPmaUrlOk() (*string, bool)`

GetPmaUrlOk returns a tuple with the PmaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmaUrl

`func (o *MysqlMysql8) SetPmaUrl(v string)`

SetPmaUrl sets PmaUrl field to given value.

### HasPmaUrl

`func (o *MysqlMysql8) HasPmaUrl() bool`

HasPmaUrl returns a boolean if a field has been set.

### GetDiskUsed

`func (o *MysqlMysql8) GetDiskUsed() string`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *MysqlMysql8) GetDiskUsedOk() (*string, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *MysqlMysql8) SetDiskUsed(v string)`

SetDiskUsed sets DiskUsed field to given value.

### HasDiskUsed

`func (o *MysqlMysql8) HasDiskUsed() bool`

HasDiskUsed returns a boolean if a field has been set.

### GetDiskLeft

`func (o *MysqlMysql8) GetDiskLeft() string`

GetDiskLeft returns the DiskLeft field if non-nil, zero value otherwise.

### GetDiskLeftOk

`func (o *MysqlMysql8) GetDiskLeftOk() (*string, bool)`

GetDiskLeftOk returns a tuple with the DiskLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLeft

`func (o *MysqlMysql8) SetDiskLeft(v string)`

SetDiskLeft sets DiskLeft field to given value.

### HasDiskLeft

`func (o *MysqlMysql8) HasDiskLeft() bool`

HasDiskLeft returns a boolean if a field has been set.

### GetReadOnly

`func (o *MysqlMysql8) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MysqlMysql8) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *MysqlMysql8) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *MysqlMysql8) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetPmaInstalling

`func (o *MysqlMysql8) GetPmaInstalling() bool`

GetPmaInstalling returns the PmaInstalling field if non-nil, zero value otherwise.

### GetPmaInstallingOk

`func (o *MysqlMysql8) GetPmaInstallingOk() (*bool, bool)`

GetPmaInstallingOk returns a tuple with the PmaInstalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmaInstalling

`func (o *MysqlMysql8) SetPmaInstalling(v bool)`

SetPmaInstalling sets PmaInstalling field to given value.

### HasPmaInstalling

`func (o *MysqlMysql8) HasPmaInstalling() bool`

HasPmaInstalling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


