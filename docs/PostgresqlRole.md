# PostgresqlRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**ParentRole** | Pointer to **[]string** |  | [optional] 
**Predefined** | Pointer to **bool** |  | [optional] 

## Methods

### NewPostgresqlRole

`func NewPostgresqlRole() *PostgresqlRole`

NewPostgresqlRole instantiates a new PostgresqlRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlRoleWithDefaults

`func NewPostgresqlRoleWithDefaults() *PostgresqlRole`

NewPostgresqlRoleWithDefaults instantiates a new PostgresqlRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PostgresqlRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PostgresqlRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PostgresqlRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PostgresqlRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *PostgresqlRole) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostgresqlRole) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostgresqlRole) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *PostgresqlRole) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetParentRole

`func (o *PostgresqlRole) GetParentRole() []string`

GetParentRole returns the ParentRole field if non-nil, zero value otherwise.

### GetParentRoleOk

`func (o *PostgresqlRole) GetParentRoleOk() (*[]string, bool)`

GetParentRoleOk returns a tuple with the ParentRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRole

`func (o *PostgresqlRole) SetParentRole(v []string)`

SetParentRole sets ParentRole field to given value.

### HasParentRole

`func (o *PostgresqlRole) HasParentRole() bool`

HasParentRole returns a boolean if a field has been set.

### GetPredefined

`func (o *PostgresqlRole) GetPredefined() bool`

GetPredefined returns the Predefined field if non-nil, zero value otherwise.

### GetPredefinedOk

`func (o *PostgresqlRole) GetPredefinedOk() (*bool, bool)`

GetPredefinedOk returns a tuple with the Predefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefined

`func (o *PostgresqlRole) SetPredefined(v bool)`

SetPredefined sets Predefined field to given value.

### HasPredefined

`func (o *PostgresqlRole) HasPredefined() bool`

HasPredefined returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


