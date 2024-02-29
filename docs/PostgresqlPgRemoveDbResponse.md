# PostgresqlPgRemoveDbResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to [**PostgresqlPgRemoveDbResponseError**](PostgresqlPgRemoveDbResponseError.md) |  | [optional] 

## Methods

### NewPostgresqlPgRemoveDbResponse

`func NewPostgresqlPgRemoveDbResponse() *PostgresqlPgRemoveDbResponse`

NewPostgresqlPgRemoveDbResponse instantiates a new PostgresqlPgRemoveDbResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgRemoveDbResponseWithDefaults

`func NewPostgresqlPgRemoveDbResponseWithDefaults() *PostgresqlPgRemoveDbResponse`

NewPostgresqlPgRemoveDbResponseWithDefaults instantiates a new PostgresqlPgRemoveDbResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *PostgresqlPgRemoveDbResponse) GetSuccess() map[string]interface{}`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *PostgresqlPgRemoveDbResponse) GetSuccessOk() (*map[string]interface{}, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *PostgresqlPgRemoveDbResponse) SetSuccess(v map[string]interface{})`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *PostgresqlPgRemoveDbResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *PostgresqlPgRemoveDbResponse) GetError() PostgresqlPgRemoveDbResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostgresqlPgRemoveDbResponse) GetErrorOk() (*PostgresqlPgRemoveDbResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostgresqlPgRemoveDbResponse) SetError(v PostgresqlPgRemoveDbResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *PostgresqlPgRemoveDbResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


