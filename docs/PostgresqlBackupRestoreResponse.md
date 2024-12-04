# PostgresqlBackupRestoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**StructuresPostgresqlRestoreOrder**](StructuresPostgresqlRestoreOrder.md) |  | [optional] 
**Error** | Pointer to [**PostgresqlBackupRestoreResponseError**](PostgresqlBackupRestoreResponseError.md) |  | [optional] 

## Methods

### NewPostgresqlBackupRestoreResponse

`func NewPostgresqlBackupRestoreResponse() *PostgresqlBackupRestoreResponse`

NewPostgresqlBackupRestoreResponse instantiates a new PostgresqlBackupRestoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlBackupRestoreResponseWithDefaults

`func NewPostgresqlBackupRestoreResponseWithDefaults() *PostgresqlBackupRestoreResponse`

NewPostgresqlBackupRestoreResponseWithDefaults instantiates a new PostgresqlBackupRestoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *PostgresqlBackupRestoreResponse) GetOrder() StructuresPostgresqlRestoreOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *PostgresqlBackupRestoreResponse) GetOrderOk() (*StructuresPostgresqlRestoreOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *PostgresqlBackupRestoreResponse) SetOrder(v StructuresPostgresqlRestoreOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *PostgresqlBackupRestoreResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetError

`func (o *PostgresqlBackupRestoreResponse) GetError() PostgresqlBackupRestoreResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostgresqlBackupRestoreResponse) GetErrorOk() (*PostgresqlBackupRestoreResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostgresqlBackupRestoreResponse) SetError(v PostgresqlBackupRestoreResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *PostgresqlBackupRestoreResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


