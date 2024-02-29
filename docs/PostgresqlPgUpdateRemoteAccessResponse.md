# PostgresqlPgUpdateRemoteAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | Pointer to [**PostgresqlPgRemoteAccessState**](PostgresqlPgRemoteAccessState.md) |  | [optional] 
**Error** | Pointer to [**PostgresqlPgUpdateRemoteAccessResponseError**](PostgresqlPgUpdateRemoteAccessResponseError.md) |  | [optional] 

## Methods

### NewPostgresqlPgUpdateRemoteAccessResponse

`func NewPostgresqlPgUpdateRemoteAccessResponse() *PostgresqlPgUpdateRemoteAccessResponse`

NewPostgresqlPgUpdateRemoteAccessResponse instantiates a new PostgresqlPgUpdateRemoteAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgUpdateRemoteAccessResponseWithDefaults

`func NewPostgresqlPgUpdateRemoteAccessResponseWithDefaults() *PostgresqlPgUpdateRemoteAccessResponse`

NewPostgresqlPgUpdateRemoteAccessResponseWithDefaults instantiates a new PostgresqlPgUpdateRemoteAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *PostgresqlPgUpdateRemoteAccessResponse) GetState() PostgresqlPgRemoteAccessState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PostgresqlPgUpdateRemoteAccessResponse) GetStateOk() (*PostgresqlPgRemoteAccessState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PostgresqlPgUpdateRemoteAccessResponse) SetState(v PostgresqlPgRemoteAccessState)`

SetState sets State field to given value.

### HasState

`func (o *PostgresqlPgUpdateRemoteAccessResponse) HasState() bool`

HasState returns a boolean if a field has been set.

### GetError

`func (o *PostgresqlPgUpdateRemoteAccessResponse) GetError() PostgresqlPgUpdateRemoteAccessResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostgresqlPgUpdateRemoteAccessResponse) GetErrorOk() (*PostgresqlPgUpdateRemoteAccessResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostgresqlPgUpdateRemoteAccessResponse) SetError(v PostgresqlPgUpdateRemoteAccessResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *PostgresqlPgUpdateRemoteAccessResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


