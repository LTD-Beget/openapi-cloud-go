# PostgresqlPgCreateRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**PostgresqlRole**](PostgresqlRole.md) |  | [optional] 
**Error** | Pointer to [**PostgresqlPgCreateRoleResponseError**](PostgresqlPgCreateRoleResponseError.md) |  | [optional] 

## Methods

### NewPostgresqlPgCreateRoleResponse

`func NewPostgresqlPgCreateRoleResponse() *PostgresqlPgCreateRoleResponse`

NewPostgresqlPgCreateRoleResponse instantiates a new PostgresqlPgCreateRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgCreateRoleResponseWithDefaults

`func NewPostgresqlPgCreateRoleResponseWithDefaults() *PostgresqlPgCreateRoleResponse`

NewPostgresqlPgCreateRoleResponseWithDefaults instantiates a new PostgresqlPgCreateRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *PostgresqlPgCreateRoleResponse) GetRole() PostgresqlRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PostgresqlPgCreateRoleResponse) GetRoleOk() (*PostgresqlRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PostgresqlPgCreateRoleResponse) SetRole(v PostgresqlRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PostgresqlPgCreateRoleResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetError

`func (o *PostgresqlPgCreateRoleResponse) GetError() PostgresqlPgCreateRoleResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostgresqlPgCreateRoleResponse) GetErrorOk() (*PostgresqlPgCreateRoleResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostgresqlPgCreateRoleResponse) SetError(v PostgresqlPgCreateRoleResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *PostgresqlPgCreateRoleResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


