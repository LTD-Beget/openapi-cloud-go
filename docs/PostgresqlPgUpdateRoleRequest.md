# PostgresqlPgUpdateRoleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ParentRole** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPostgresqlPgUpdateRoleRequest

`func NewPostgresqlPgUpdateRoleRequest() *PostgresqlPgUpdateRoleRequest`

NewPostgresqlPgUpdateRoleRequest instantiates a new PostgresqlPgUpdateRoleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgUpdateRoleRequestWithDefaults

`func NewPostgresqlPgUpdateRoleRequestWithDefaults() *PostgresqlPgUpdateRoleRequest`

NewPostgresqlPgUpdateRoleRequestWithDefaults instantiates a new PostgresqlPgUpdateRoleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *PostgresqlPgUpdateRoleRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PostgresqlPgUpdateRoleRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PostgresqlPgUpdateRoleRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PostgresqlPgUpdateRoleRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetRoleName

`func (o *PostgresqlPgUpdateRoleRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *PostgresqlPgUpdateRoleRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *PostgresqlPgUpdateRoleRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *PostgresqlPgUpdateRoleRequest) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetPassword

`func (o *PostgresqlPgUpdateRoleRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostgresqlPgUpdateRoleRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostgresqlPgUpdateRoleRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PostgresqlPgUpdateRoleRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetParentRole

`func (o *PostgresqlPgUpdateRoleRequest) GetParentRole() []string`

GetParentRole returns the ParentRole field if non-nil, zero value otherwise.

### GetParentRoleOk

`func (o *PostgresqlPgUpdateRoleRequest) GetParentRoleOk() (*[]string, bool)`

GetParentRoleOk returns a tuple with the ParentRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRole

`func (o *PostgresqlPgUpdateRoleRequest) SetParentRole(v []string)`

SetParentRole sets ParentRole field to given value.

### HasParentRole

`func (o *PostgresqlPgUpdateRoleRequest) HasParentRole() bool`

HasParentRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


