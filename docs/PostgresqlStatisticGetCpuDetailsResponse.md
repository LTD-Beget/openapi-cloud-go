# PostgresqlStatisticGetCpuDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **[]string** |  | [optional] 
**User** | Pointer to **[]float64** |  | [optional] 
**Nice** | Pointer to **[]float64** |  | [optional] 
**System** | Pointer to **[]float64** |  | [optional] 
**Idle** | Pointer to **[]float64** |  | [optional] 
**Iowait** | Pointer to **[]float64** |  | [optional] 
**Irq** | Pointer to **[]float64** |  | [optional] 
**Softirq** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewPostgresqlStatisticGetCpuDetailsResponse

`func NewPostgresqlStatisticGetCpuDetailsResponse() *PostgresqlStatisticGetCpuDetailsResponse`

NewPostgresqlStatisticGetCpuDetailsResponse instantiates a new PostgresqlStatisticGetCpuDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlStatisticGetCpuDetailsResponseWithDefaults

`func NewPostgresqlStatisticGetCpuDetailsResponseWithDefaults() *PostgresqlStatisticGetCpuDetailsResponse`

NewPostgresqlStatisticGetCpuDetailsResponseWithDefaults instantiates a new PostgresqlStatisticGetCpuDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetDate() []string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetDateOk() (*[]string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PostgresqlStatisticGetCpuDetailsResponse) SetDate(v []string)`

SetDate sets Date field to given value.

### HasDate

`func (o *PostgresqlStatisticGetCpuDetailsResponse) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetUser

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetUser() []float64`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetUserOk() (*[]float64, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PostgresqlStatisticGetCpuDetailsResponse) SetUser(v []float64)`

SetUser sets User field to given value.

### HasUser

`func (o *PostgresqlStatisticGetCpuDetailsResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetNice

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetNice() []float64`

GetNice returns the Nice field if non-nil, zero value otherwise.

### GetNiceOk

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetNiceOk() (*[]float64, bool)`

GetNiceOk returns a tuple with the Nice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNice

`func (o *PostgresqlStatisticGetCpuDetailsResponse) SetNice(v []float64)`

SetNice sets Nice field to given value.

### HasNice

`func (o *PostgresqlStatisticGetCpuDetailsResponse) HasNice() bool`

HasNice returns a boolean if a field has been set.

### GetSystem

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetSystem() []float64`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetSystemOk() (*[]float64, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *PostgresqlStatisticGetCpuDetailsResponse) SetSystem(v []float64)`

SetSystem sets System field to given value.

### HasSystem

`func (o *PostgresqlStatisticGetCpuDetailsResponse) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetIdle

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetIdle() []float64`

GetIdle returns the Idle field if non-nil, zero value otherwise.

### GetIdleOk

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetIdleOk() (*[]float64, bool)`

GetIdleOk returns a tuple with the Idle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdle

`func (o *PostgresqlStatisticGetCpuDetailsResponse) SetIdle(v []float64)`

SetIdle sets Idle field to given value.

### HasIdle

`func (o *PostgresqlStatisticGetCpuDetailsResponse) HasIdle() bool`

HasIdle returns a boolean if a field has been set.

### GetIowait

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetIowait() []float64`

GetIowait returns the Iowait field if non-nil, zero value otherwise.

### GetIowaitOk

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetIowaitOk() (*[]float64, bool)`

GetIowaitOk returns a tuple with the Iowait field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIowait

`func (o *PostgresqlStatisticGetCpuDetailsResponse) SetIowait(v []float64)`

SetIowait sets Iowait field to given value.

### HasIowait

`func (o *PostgresqlStatisticGetCpuDetailsResponse) HasIowait() bool`

HasIowait returns a boolean if a field has been set.

### GetIrq

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetIrq() []float64`

GetIrq returns the Irq field if non-nil, zero value otherwise.

### GetIrqOk

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetIrqOk() (*[]float64, bool)`

GetIrqOk returns a tuple with the Irq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIrq

`func (o *PostgresqlStatisticGetCpuDetailsResponse) SetIrq(v []float64)`

SetIrq sets Irq field to given value.

### HasIrq

`func (o *PostgresqlStatisticGetCpuDetailsResponse) HasIrq() bool`

HasIrq returns a boolean if a field has been set.

### GetSoftirq

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetSoftirq() []float64`

GetSoftirq returns the Softirq field if non-nil, zero value otherwise.

### GetSoftirqOk

`func (o *PostgresqlStatisticGetCpuDetailsResponse) GetSoftirqOk() (*[]float64, bool)`

GetSoftirqOk returns a tuple with the Softirq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftirq

`func (o *PostgresqlStatisticGetCpuDetailsResponse) SetSoftirq(v []float64)`

SetSoftirq sets Softirq field to given value.

### HasSoftirq

`func (o *PostgresqlStatisticGetCpuDetailsResponse) HasSoftirq() bool`

HasSoftirq returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


