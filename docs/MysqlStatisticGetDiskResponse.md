# MysqlStatisticGetDiskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataRead** | Pointer to [**StructuresStatisticSeriesData**](StructuresStatisticSeriesData.md) |  | [optional] 
**DataWrite** | Pointer to [**StructuresStatisticSeriesData**](StructuresStatisticSeriesData.md) |  | [optional] 

## Methods

### NewMysqlStatisticGetDiskResponse

`func NewMysqlStatisticGetDiskResponse() *MysqlStatisticGetDiskResponse`

NewMysqlStatisticGetDiskResponse instantiates a new MysqlStatisticGetDiskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlStatisticGetDiskResponseWithDefaults

`func NewMysqlStatisticGetDiskResponseWithDefaults() *MysqlStatisticGetDiskResponse`

NewMysqlStatisticGetDiskResponseWithDefaults instantiates a new MysqlStatisticGetDiskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataRead

`func (o *MysqlStatisticGetDiskResponse) GetDataRead() StructuresStatisticSeriesData`

GetDataRead returns the DataRead field if non-nil, zero value otherwise.

### GetDataReadOk

`func (o *MysqlStatisticGetDiskResponse) GetDataReadOk() (*StructuresStatisticSeriesData, bool)`

GetDataReadOk returns a tuple with the DataRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRead

`func (o *MysqlStatisticGetDiskResponse) SetDataRead(v StructuresStatisticSeriesData)`

SetDataRead sets DataRead field to given value.

### HasDataRead

`func (o *MysqlStatisticGetDiskResponse) HasDataRead() bool`

HasDataRead returns a boolean if a field has been set.

### GetDataWrite

`func (o *MysqlStatisticGetDiskResponse) GetDataWrite() StructuresStatisticSeriesData`

GetDataWrite returns the DataWrite field if non-nil, zero value otherwise.

### GetDataWriteOk

`func (o *MysqlStatisticGetDiskResponse) GetDataWriteOk() (*StructuresStatisticSeriesData, bool)`

GetDataWriteOk returns a tuple with the DataWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrite

`func (o *MysqlStatisticGetDiskResponse) SetDataWrite(v StructuresStatisticSeriesData)`

SetDataWrite sets DataWrite field to given value.

### HasDataWrite

`func (o *MysqlStatisticGetDiskResponse) HasDataWrite() bool`

HasDataWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


