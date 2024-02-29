# PostgresqlPgUpdateDbResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | Pointer to [**PostgresqlPgDb**](PostgresqlPgDb.md) |  | [optional] 
**Error** | Pointer to [**PostgresqlPgUpdateDbResponseError**](PostgresqlPgUpdateDbResponseError.md) |  | [optional] 

## Methods

### NewPostgresqlPgUpdateDbResponse

`func NewPostgresqlPgUpdateDbResponse() *PostgresqlPgUpdateDbResponse`

NewPostgresqlPgUpdateDbResponse instantiates a new PostgresqlPgUpdateDbResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgUpdateDbResponseWithDefaults

`func NewPostgresqlPgUpdateDbResponseWithDefaults() *PostgresqlPgUpdateDbResponse`

NewPostgresqlPgUpdateDbResponseWithDefaults instantiates a new PostgresqlPgUpdateDbResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *PostgresqlPgUpdateDbResponse) GetDb() PostgresqlPgDb`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *PostgresqlPgUpdateDbResponse) GetDbOk() (*PostgresqlPgDb, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *PostgresqlPgUpdateDbResponse) SetDb(v PostgresqlPgDb)`

SetDb sets Db field to given value.

### HasDb

`func (o *PostgresqlPgUpdateDbResponse) HasDb() bool`

HasDb returns a boolean if a field has been set.

### GetError

`func (o *PostgresqlPgUpdateDbResponse) GetError() PostgresqlPgUpdateDbResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostgresqlPgUpdateDbResponse) GetErrorOk() (*PostgresqlPgUpdateDbResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostgresqlPgUpdateDbResponse) SetError(v PostgresqlPgUpdateDbResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *PostgresqlPgUpdateDbResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


