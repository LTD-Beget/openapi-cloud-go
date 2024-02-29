# MysqlSetConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**Param** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewMysqlSetConfigRequest

`func NewMysqlSetConfigRequest() *MysqlSetConfigRequest`

NewMysqlSetConfigRequest instantiates a new MysqlSetConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlSetConfigRequestWithDefaults

`func NewMysqlSetConfigRequestWithDefaults() *MysqlSetConfigRequest`

NewMysqlSetConfigRequestWithDefaults instantiates a new MysqlSetConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *MysqlSetConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MysqlSetConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MysqlSetConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MysqlSetConfigRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetParam

`func (o *MysqlSetConfigRequest) GetParam() map[string]string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *MysqlSetConfigRequest) GetParamOk() (*map[string]string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *MysqlSetConfigRequest) SetParam(v map[string]string)`

SetParam sets Param field to given value.

### HasParam

`func (o *MysqlSetConfigRequest) HasParam() bool`

HasParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


