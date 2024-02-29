# PostgresqlPgCreateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ParentRole** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPostgresqlPgCreateRoleRequest

`func NewPostgresqlPgCreateRoleRequest() *PostgresqlPgCreateRoleRequest`

NewPostgresqlPgCreateRoleRequest instantiates a new PostgresqlPgCreateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgCreateRoleRequestWithDefaults

`func NewPostgresqlPgCreateRoleRequestWithDefaults() *PostgresqlPgCreateRoleRequest`

NewPostgresqlPgCreateRoleRequestWithDefaults instantiates a new PostgresqlPgCreateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *PostgresqlPgCreateRoleRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PostgresqlPgCreateRoleRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PostgresqlPgCreateRoleRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PostgresqlPgCreateRoleRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRoleName

`func (o *PostgresqlPgCreateRoleRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *PostgresqlPgCreateRoleRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *PostgresqlPgCreateRoleRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *PostgresqlPgCreateRoleRequest) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetPassword

`func (o *PostgresqlPgCreateRoleRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostgresqlPgCreateRoleRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostgresqlPgCreateRoleRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PostgresqlPgCreateRoleRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetParentRole

`func (o *PostgresqlPgCreateRoleRequest) GetParentRole() []string`

GetParentRole returns the ParentRole field if non-nil, zero value otherwise.

### GetParentRoleOk

`func (o *PostgresqlPgCreateRoleRequest) GetParentRoleOk() (*[]string, bool)`

GetParentRoleOk returns a tuple with the ParentRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRole

`func (o *PostgresqlPgCreateRoleRequest) SetParentRole(v []string)`

SetParentRole sets ParentRole field to given value.

### HasParentRole

`func (o *PostgresqlPgCreateRoleRequest) HasParentRole() bool`

HasParentRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


