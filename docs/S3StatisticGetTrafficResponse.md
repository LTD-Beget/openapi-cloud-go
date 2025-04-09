# S3StatisticGetTrafficResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataRx** | Pointer to [**StructuresStatisticSeriesData**](StructuresStatisticSeriesData.md) |  | [optional] 
**DataTx** | Pointer to [**StructuresStatisticSeriesData**](StructuresStatisticSeriesData.md) |  | [optional] 

## Methods

### NewS3StatisticGetTrafficResponse

`func NewS3StatisticGetTrafficResponse() *S3StatisticGetTrafficResponse`

NewS3StatisticGetTrafficResponse instantiates a new S3StatisticGetTrafficResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3StatisticGetTrafficResponseWithDefaults

`func NewS3StatisticGetTrafficResponseWithDefaults() *S3StatisticGetTrafficResponse`

NewS3StatisticGetTrafficResponseWithDefaults instantiates a new S3StatisticGetTrafficResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataRx

`func (o *S3StatisticGetTrafficResponse) GetDataRx() StructuresStatisticSeriesData`

GetDataRx returns the DataRx field if non-nil, zero value otherwise.

### GetDataRxOk

`func (o *S3StatisticGetTrafficResponse) GetDataRxOk() (*StructuresStatisticSeriesData, bool)`

GetDataRxOk returns a tuple with the DataRx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRx

`func (o *S3StatisticGetTrafficResponse) SetDataRx(v StructuresStatisticSeriesData)`

SetDataRx sets DataRx field to given value.

### HasDataRx

`func (o *S3StatisticGetTrafficResponse) HasDataRx() bool`

HasDataRx returns a boolean if a field has been set.

### GetDataTx

`func (o *S3StatisticGetTrafficResponse) GetDataTx() StructuresStatisticSeriesData`

GetDataTx returns the DataTx field if non-nil, zero value otherwise.

### GetDataTxOk

`func (o *S3StatisticGetTrafficResponse) GetDataTxOk() (*StructuresStatisticSeriesData, bool)`

GetDataTxOk returns a tuple with the DataTx field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataTx

`func (o *S3StatisticGetTrafficResponse) SetDataTx(v StructuresStatisticSeriesData)`

SetDataTx sets DataTx field to given value.

### HasDataTx

`func (o *S3StatisticGetTrafficResponse) HasDataTx() bool`

HasDataTx returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


