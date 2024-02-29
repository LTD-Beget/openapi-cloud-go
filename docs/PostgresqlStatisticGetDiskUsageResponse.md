# PostgresqlStatisticGetDiskUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiskUsage** | Pointer to [**StructuresStatisticSeriesData**](StructuresStatisticSeriesData.md) |  | [optional] 

## Methods

### NewPostgresqlStatisticGetDiskUsageResponse

`func NewPostgresqlStatisticGetDiskUsageResponse() *PostgresqlStatisticGetDiskUsageResponse`

NewPostgresqlStatisticGetDiskUsageResponse instantiates a new PostgresqlStatisticGetDiskUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlStatisticGetDiskUsageResponseWithDefaults

`func NewPostgresqlStatisticGetDiskUsageResponseWithDefaults() *PostgresqlStatisticGetDiskUsageResponse`

NewPostgresqlStatisticGetDiskUsageResponseWithDefaults instantiates a new PostgresqlStatisticGetDiskUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiskUsage

`func (o *PostgresqlStatisticGetDiskUsageResponse) GetDiskUsage() StructuresStatisticSeriesData`

GetDiskUsage returns the DiskUsage field if non-nil, zero value otherwise.

### GetDiskUsageOk

`func (o *PostgresqlStatisticGetDiskUsageResponse) GetDiskUsageOk() (*StructuresStatisticSeriesData, bool)`

GetDiskUsageOk returns a tuple with the DiskUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskUsage

`func (o *PostgresqlStatisticGetDiskUsageResponse) SetDiskUsage(v StructuresStatisticSeriesData)`

SetDiskUsage sets DiskUsage field to given value.

### HasDiskUsage

`func (o *PostgresqlStatisticGetDiskUsageResponse) HasDiskUsage() bool`

HasDiskUsage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


