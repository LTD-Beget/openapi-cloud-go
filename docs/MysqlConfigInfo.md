# MysqlConfigInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParamDefault** | Pointer to **map[string]string** |  | [optional] 
**ParamConfig** | Pointer to [**[]StructuresParamConfig**](StructuresParamConfig.md) |  | [optional] 
**Param** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMysqlConfigInfo

`func NewMysqlConfigInfo() *MysqlConfigInfo`

NewMysqlConfigInfo instantiates a new MysqlConfigInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlConfigInfoWithDefaults

`func NewMysqlConfigInfoWithDefaults() *MysqlConfigInfo`

NewMysqlConfigInfoWithDefaults instantiates a new MysqlConfigInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParamDefault

`func (o *MysqlConfigInfo) GetParamDefault() map[string]string`

GetParamDefault returns the ParamDefault field if non-nil, zero value otherwise.

### GetParamDefaultOk

`func (o *MysqlConfigInfo) GetParamDefaultOk() (*map[string]string, bool)`

GetParamDefaultOk returns a tuple with the ParamDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamDefault

`func (o *MysqlConfigInfo) SetParamDefault(v map[string]string)`

SetParamDefault sets ParamDefault field to given value.

### HasParamDefault

`func (o *MysqlConfigInfo) HasParamDefault() bool`

HasParamDefault returns a boolean if a field has been set.

### GetParamConfig

`func (o *MysqlConfigInfo) GetParamConfig() []StructuresParamConfig`

GetParamConfig returns the ParamConfig field if non-nil, zero value otherwise.

### GetParamConfigOk

`func (o *MysqlConfigInfo) GetParamConfigOk() (*[]StructuresParamConfig, bool)`

GetParamConfigOk returns a tuple with the ParamConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamConfig

`func (o *MysqlConfigInfo) SetParamConfig(v []StructuresParamConfig)`

SetParamConfig sets ParamConfig field to given value.

### HasParamConfig

`func (o *MysqlConfigInfo) HasParamConfig() bool`

HasParamConfig returns a boolean if a field has been set.

### GetParam

`func (o *MysqlConfigInfo) GetParam() map[string]string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *MysqlConfigInfo) GetParamOk() (*map[string]string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *MysqlConfigInfo) SetParam(v map[string]string)`

SetParam sets Param field to given value.

### HasParam

`func (o *MysqlConfigInfo) HasParam() bool`

HasParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


