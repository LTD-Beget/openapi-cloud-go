# MysqlMysql5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**MysqlMysql5Configuration**](MysqlMysql5Configuration.md) |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**AddressInfo** | Pointer to [**StructuresAddressInfo**](StructuresAddressInfo.md) |  | [optional] 
**PmaUrl** | Pointer to **string** |  | [optional] 
**DiskUsed** | Pointer to **string** |  | [optional] 
**DiskLeft** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**PmaInstalling** | Pointer to **bool** |  | [optional] 

## Methods

### NewMysqlMysql5

`func NewMysqlMysql5() *MysqlMysql5`

NewMysqlMysql5 instantiates a new MysqlMysql5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlMysql5WithDefaults

`func NewMysqlMysql5WithDefaults() *MysqlMysql5`

NewMysqlMysql5WithDefaults instantiates a new MysqlMysql5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *MysqlMysql5) GetConfiguration() MysqlMysql5Configuration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *MysqlMysql5) GetConfigurationOk() (*MysqlMysql5Configuration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *MysqlMysql5) SetConfiguration(v MysqlMysql5Configuration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *MysqlMysql5) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetHost

`func (o *MysqlMysql5) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MysqlMysql5) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MysqlMysql5) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MysqlMysql5) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *MysqlMysql5) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *MysqlMysql5) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *MysqlMysql5) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *MysqlMysql5) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetAddressInfo

`func (o *MysqlMysql5) GetAddressInfo() StructuresAddressInfo`

GetAddressInfo returns the AddressInfo field if non-nil, zero value otherwise.

### GetAddressInfoOk

`func (o *MysqlMysql5) GetAddressInfoOk() (*StructuresAddressInfo, bool)`

GetAddressInfoOk returns a tuple with the AddressInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressInfo

`func (o *MysqlMysql5) SetAddressInfo(v StructuresAddressInfo)`

SetAddressInfo sets AddressInfo field to given value.

### HasAddressInfo

`func (o *MysqlMysql5) HasAddressInfo() bool`

HasAddressInfo returns a boolean if a field has been set.

### GetPmaUrl

`func (o *MysqlMysql5) GetPmaUrl() string`

GetPmaUrl returns the PmaUrl field if non-nil, zero value otherwise.

### GetPmaUrlOk

`func (o *MysqlMysql5) GetPmaUrlOk() (*string, bool)`

GetPmaUrlOk returns a tuple with the PmaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmaUrl

`func (o *MysqlMysql5) SetPmaUrl(v string)`

SetPmaUrl sets PmaUrl field to given value.

### HasPmaUrl

`func (o *MysqlMysql5) HasPmaUrl() bool`

HasPmaUrl returns a boolean if a field has been set.

### GetDiskUsed

`func (o *MysqlMysql5) GetDiskUsed() string`

GetDiskUsed returns the DiskUsed field if non-nil, zero value otherwise.

### GetDiskUsedOk

`func (o *MysqlMysql5) GetDiskUsedOk() (*string, bool)`

GetDiskUsedOk returns a tuple with the DiskUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsed

`func (o *MysqlMysql5) SetDiskUsed(v string)`

SetDiskUsed sets DiskUsed field to given value.

### HasDiskUsed

`func (o *MysqlMysql5) HasDiskUsed() bool`

HasDiskUsed returns a boolean if a field has been set.

### GetDiskLeft

`func (o *MysqlMysql5) GetDiskLeft() string`

GetDiskLeft returns the DiskLeft field if non-nil, zero value otherwise.

### GetDiskLeftOk

`func (o *MysqlMysql5) GetDiskLeftOk() (*string, bool)`

GetDiskLeftOk returns a tuple with the DiskLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskLeft

`func (o *MysqlMysql5) SetDiskLeft(v string)`

SetDiskLeft sets DiskLeft field to given value.

### HasDiskLeft

`func (o *MysqlMysql5) HasDiskLeft() bool`

HasDiskLeft returns a boolean if a field has been set.

### GetReadOnly

`func (o *MysqlMysql5) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *MysqlMysql5) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *MysqlMysql5) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *MysqlMysql5) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetPmaInstalling

`func (o *MysqlMysql5) GetPmaInstalling() bool`

GetPmaInstalling returns the PmaInstalling field if non-nil, zero value otherwise.

### GetPmaInstallingOk

`func (o *MysqlMysql5) GetPmaInstallingOk() (*bool, bool)`

GetPmaInstallingOk returns a tuple with the PmaInstalling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmaInstalling

`func (o *MysqlMysql5) SetPmaInstalling(v bool)`

SetPmaInstalling sets PmaInstalling field to given value.

### HasPmaInstalling

`func (o *MysqlMysql5) HasPmaInstalling() bool`

HasPmaInstalling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


