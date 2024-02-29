# PostgresqlStatisticGetDiskResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataRead** | Pointer to [**StructuresStatisticSeriesData**](StructuresStatisticSeriesData.md) |  | [optional] 
**DataWrite** | Pointer to [**StructuresStatisticSeriesData**](StructuresStatisticSeriesData.md) |  | [optional] 

## Methods

### NewPostgresqlStatisticGetDiskResponse

`func NewPostgresqlStatisticGetDiskResponse() *PostgresqlStatisticGetDiskResponse`

NewPostgresqlStatisticGetDiskResponse instantiates a new PostgresqlStatisticGetDiskResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlStatisticGetDiskResponseWithDefaults

`func NewPostgresqlStatisticGetDiskResponseWithDefaults() *PostgresqlStatisticGetDiskResponse`

NewPostgresqlStatisticGetDiskResponseWithDefaults instantiates a new PostgresqlStatisticGetDiskResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataRead

`func (o *PostgresqlStatisticGetDiskResponse) GetDataRead() StructuresStatisticSeriesData`

GetDataRead returns the DataRead field if non-nil, zero value otherwise.

### GetDataReadOk

`func (o *PostgresqlStatisticGetDiskResponse) GetDataReadOk() (*StructuresStatisticSeriesData, bool)`

GetDataReadOk returns a tuple with the DataRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRead

`func (o *PostgresqlStatisticGetDiskResponse) SetDataRead(v StructuresStatisticSeriesData)`

SetDataRead sets DataRead field to given value.

### HasDataRead

`func (o *PostgresqlStatisticGetDiskResponse) HasDataRead() bool`

HasDataRead returns a boolean if a field has been set.

### GetDataWrite

`func (o *PostgresqlStatisticGetDiskResponse) GetDataWrite() StructuresStatisticSeriesData`

GetDataWrite returns the DataWrite field if non-nil, zero value otherwise.

### GetDataWriteOk

`func (o *PostgresqlStatisticGetDiskResponse) GetDataWriteOk() (*StructuresStatisticSeriesData, bool)`

GetDataWriteOk returns a tuple with the DataWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataWrite

`func (o *PostgresqlStatisticGetDiskResponse) SetDataWrite(v StructuresStatisticSeriesData)`

SetDataWrite sets DataWrite field to given value.

### HasDataWrite

`func (o *PostgresqlStatisticGetDiskResponse) HasDataWrite() bool`

HasDataWrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


