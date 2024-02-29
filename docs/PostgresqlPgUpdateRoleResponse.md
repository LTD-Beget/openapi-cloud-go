# PostgresqlPgUpdateRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**PostgresqlRole**](PostgresqlRole.md) |  | [optional] 
**Error** | Pointer to [**PostgresqlPgUpdateRoleResponseError**](PostgresqlPgUpdateRoleResponseError.md) |  | [optional] 

## Methods

### NewPostgresqlPgUpdateRoleResponse

`func NewPostgresqlPgUpdateRoleResponse() *PostgresqlPgUpdateRoleResponse`

NewPostgresqlPgUpdateRoleResponse instantiates a new PostgresqlPgUpdateRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgUpdateRoleResponseWithDefaults

`func NewPostgresqlPgUpdateRoleResponseWithDefaults() *PostgresqlPgUpdateRoleResponse`

NewPostgresqlPgUpdateRoleResponseWithDefaults instantiates a new PostgresqlPgUpdateRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *PostgresqlPgUpdateRoleResponse) GetRole() PostgresqlRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PostgresqlPgUpdateRoleResponse) GetRoleOk() (*PostgresqlRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PostgresqlPgUpdateRoleResponse) SetRole(v PostgresqlRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PostgresqlPgUpdateRoleResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetError

`func (o *PostgresqlPgUpdateRoleResponse) GetError() PostgresqlPgUpdateRoleResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostgresqlPgUpdateRoleResponse) GetErrorOk() (*PostgresqlPgUpdateRoleResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostgresqlPgUpdateRoleResponse) SetError(v PostgresqlPgUpdateRoleResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *PostgresqlPgUpdateRoleResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


