# PostgresqlPgGetRoleListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**[]PostgresqlRole**](PostgresqlRole.md) |  | [optional] 

## Methods

### NewPostgresqlPgGetRoleListResponse

`func NewPostgresqlPgGetRoleListResponse() *PostgresqlPgGetRoleListResponse`

NewPostgresqlPgGetRoleListResponse instantiates a new PostgresqlPgGetRoleListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgGetRoleListResponseWithDefaults

`func NewPostgresqlPgGetRoleListResponseWithDefaults() *PostgresqlPgGetRoleListResponse`

NewPostgresqlPgGetRoleListResponseWithDefaults instantiates a new PostgresqlPgGetRoleListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *PostgresqlPgGetRoleListResponse) GetRole() []PostgresqlRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PostgresqlPgGetRoleListResponse) GetRoleOk() (*[]PostgresqlRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PostgresqlPgGetRoleListResponse) SetRole(v []PostgresqlRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *PostgresqlPgGetRoleListResponse) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


