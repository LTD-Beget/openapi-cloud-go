# PostgresqlStatisticGetMemoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Memory** | Pointer to [**StructuresStatisticSeriesData**](StructuresStatisticSeriesData.md) |  | [optional] 

## Methods

### NewPostgresqlStatisticGetMemoryResponse

`func NewPostgresqlStatisticGetMemoryResponse() *PostgresqlStatisticGetMemoryResponse`

NewPostgresqlStatisticGetMemoryResponse instantiates a new PostgresqlStatisticGetMemoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlStatisticGetMemoryResponseWithDefaults

`func NewPostgresqlStatisticGetMemoryResponseWithDefaults() *PostgresqlStatisticGetMemoryResponse`

NewPostgresqlStatisticGetMemoryResponseWithDefaults instantiates a new PostgresqlStatisticGetMemoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMemory

`func (o *PostgresqlStatisticGetMemoryResponse) GetMemory() StructuresStatisticSeriesData`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *PostgresqlStatisticGetMemoryResponse) GetMemoryOk() (*StructuresStatisticSeriesData, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *PostgresqlStatisticGetMemoryResponse) SetMemory(v StructuresStatisticSeriesData)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *PostgresqlStatisticGetMemoryResponse) HasMemory() bool`

HasMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


