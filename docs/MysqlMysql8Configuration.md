# MysqlMysql8Configuration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCount** | Pointer to **int32** |  | [optional] 
**DiskSize** | Pointer to **int32** |  | [optional] 
**Memory** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**DisplayVersion** | Pointer to **string** |  | [optional] 
**ParamConfig** | Pointer to [**[]StructuresParamConfig**](StructuresParamConfig.md) |  | [optional] 
**ParamDefault** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMysqlMysql8Configuration

`func NewMysqlMysql8Configuration() *MysqlMysql8Configuration`

NewMysqlMysql8Configuration instantiates a new MysqlMysql8Configuration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlMysql8ConfigurationWithDefaults

`func NewMysqlMysql8ConfigurationWithDefaults() *MysqlMysql8Configuration`

NewMysqlMysql8ConfigurationWithDefaults instantiates a new MysqlMysql8Configuration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCount

`func (o *MysqlMysql8Configuration) GetCpuCount() int32`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *MysqlMysql8Configuration) GetCpuCountOk() (*int32, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *MysqlMysql8Configuration) SetCpuCount(v int32)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *MysqlMysql8Configuration) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetDiskSize

`func (o *MysqlMysql8Configuration) GetDiskSize() int32`

GetDiskSize returns the DiskSize field if non-nil, zero value otherwise.

### GetDiskSizeOk

`func (o *MysqlMysql8Configuration) GetDiskSizeOk() (*int32, bool)`

GetDiskSizeOk returns a tuple with the DiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSize

`func (o *MysqlMysql8Configuration) SetDiskSize(v int32)`

SetDiskSize sets DiskSize field to given value.

### HasDiskSize

`func (o *MysqlMysql8Configuration) HasDiskSize() bool`

HasDiskSize returns a boolean if a field has been set.

### GetMemory

`func (o *MysqlMysql8Configuration) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *MysqlMysql8Configuration) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *MysqlMysql8Configuration) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *MysqlMysql8Configuration) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetVersion

`func (o *MysqlMysql8Configuration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MysqlMysql8Configuration) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MysqlMysql8Configuration) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MysqlMysql8Configuration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDisplayVersion

`func (o *MysqlMysql8Configuration) GetDisplayVersion() string`

GetDisplayVersion returns the DisplayVersion field if non-nil, zero value otherwise.

### GetDisplayVersionOk

`func (o *MysqlMysql8Configuration) GetDisplayVersionOk() (*string, bool)`

GetDisplayVersionOk returns a tuple with the DisplayVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayVersion

`func (o *MysqlMysql8Configuration) SetDisplayVersion(v string)`

SetDisplayVersion sets DisplayVersion field to given value.

### HasDisplayVersion

`func (o *MysqlMysql8Configuration) HasDisplayVersion() bool`

HasDisplayVersion returns a boolean if a field has been set.

### GetParamConfig

`func (o *MysqlMysql8Configuration) GetParamConfig() []StructuresParamConfig`

GetParamConfig returns the ParamConfig field if non-nil, zero value otherwise.

### GetParamConfigOk

`func (o *MysqlMysql8Configuration) GetParamConfigOk() (*[]StructuresParamConfig, bool)`

GetParamConfigOk returns a tuple with the ParamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamConfig

`func (o *MysqlMysql8Configuration) SetParamConfig(v []StructuresParamConfig)`

SetParamConfig sets ParamConfig field to given value.

### HasParamConfig

`func (o *MysqlMysql8Configuration) HasParamConfig() bool`

HasParamConfig returns a boolean if a field has been set.

### GetParamDefault

`func (o *MysqlMysql8Configuration) GetParamDefault() map[string]string`

GetParamDefault returns the ParamDefault field if non-nil, zero value otherwise.

### GetParamDefaultOk

`func (o *MysqlMysql8Configuration) GetParamDefaultOk() (*map[string]string, bool)`

GetParamDefaultOk returns a tuple with the ParamDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamDefault

`func (o *MysqlMysql8Configuration) SetParamDefault(v map[string]string)`

SetParamDefault sets ParamDefault field to given value.

### HasParamDefault

`func (o *MysqlMysql8Configuration) HasParamDefault() bool`

HasParamDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


