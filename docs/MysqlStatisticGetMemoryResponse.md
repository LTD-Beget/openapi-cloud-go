# MysqlStatisticGetMemoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | Pointer to [**StructuresStatisticSeriesData**](StructuresStatisticSeriesData.md) |  | [optional] 

## Methods

### NewMysqlStatisticGetMemoryResponse

`func NewMysqlStatisticGetMemoryResponse() *MysqlStatisticGetMemoryResponse`

NewMysqlStatisticGetMemoryResponse instantiates a new MysqlStatisticGetMemoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlStatisticGetMemoryResponseWithDefaults

`func NewMysqlStatisticGetMemoryResponseWithDefaults() *MysqlStatisticGetMemoryResponse`

NewMysqlStatisticGetMemoryResponseWithDefaults instantiates a new MysqlStatisticGetMemoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *MysqlStatisticGetMemoryResponse) GetMemory() StructuresStatisticSeriesData`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *MysqlStatisticGetMemoryResponse) GetMemoryOk() (*StructuresStatisticSeriesData, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *MysqlStatisticGetMemoryResponse) SetMemory(v StructuresStatisticSeriesData)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *MysqlStatisticGetMemoryResponse) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


