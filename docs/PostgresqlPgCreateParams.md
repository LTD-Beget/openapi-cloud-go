# PostgresqlPgCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Param** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPostgresqlPgCreateParams

`func NewPostgresqlPgCreateParams() *PostgresqlPgCreateParams`

NewPostgresqlPgCreateParams instantiates a new PostgresqlPgCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgCreateParamsWithDefaults

`func NewPostgresqlPgCreateParamsWithDefaults() *PostgresqlPgCreateParams`

NewPostgresqlPgCreateParamsWithDefaults instantiates a new PostgresqlPgCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *PostgresqlPgCreateParams) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *PostgresqlPgCreateParams) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *PostgresqlPgCreateParams) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *PostgresqlPgCreateParams) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetRoleName

`func (o *PostgresqlPgCreateParams) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *PostgresqlPgCreateParams) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *PostgresqlPgCreateParams) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *PostgresqlPgCreateParams) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetPassword

`func (o *PostgresqlPgCreateParams) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostgresqlPgCreateParams) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostgresqlPgCreateParams) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PostgresqlPgCreateParams) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetParam

`func (o *PostgresqlPgCreateParams) GetParam() map[string]string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *PostgresqlPgCreateParams) GetParamOk() (*map[string]string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *PostgresqlPgCreateParams) SetParam(v map[string]string)`

SetParam sets Param field to given value.

### HasParam

`func (o *PostgresqlPgCreateParams) HasParam() bool`

HasParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


