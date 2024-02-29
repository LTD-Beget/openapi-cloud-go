# PostgresqlPgCreateDbRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**OwnerName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewPostgresqlPgCreateDbRequest

`func NewPostgresqlPgCreateDbRequest() *PostgresqlPgCreateDbRequest`

NewPostgresqlPgCreateDbRequest instantiates a new PostgresqlPgCreateDbRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgCreateDbRequestWithDefaults

`func NewPostgresqlPgCreateDbRequestWithDefaults() *PostgresqlPgCreateDbRequest`

NewPostgresqlPgCreateDbRequestWithDefaults instantiates a new PostgresqlPgCreateDbRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *PostgresqlPgCreateDbRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PostgresqlPgCreateDbRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PostgresqlPgCreateDbRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PostgresqlPgCreateDbRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDbName

`func (o *PostgresqlPgCreateDbRequest) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *PostgresqlPgCreateDbRequest) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *PostgresqlPgCreateDbRequest) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *PostgresqlPgCreateDbRequest) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetOwnerName

`func (o *PostgresqlPgCreateDbRequest) GetOwnerName() string`

GetOwnerName returns the OwnerName field if non-nil, zero value otherwise.

### GetOwnerNameOk

`func (o *PostgresqlPgCreateDbRequest) GetOwnerNameOk() (*string, bool)`

GetOwnerNameOk returns a tuple with the OwnerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerName

`func (o *PostgresqlPgCreateDbRequest) SetOwnerName(v string)`

SetOwnerName sets OwnerName field to given value.

### HasOwnerName

`func (o *PostgresqlPgCreateDbRequest) HasOwnerName() bool`

HasOwnerName returns a boolean if a field has been set.

### GetDescription

`func (o *PostgresqlPgCreateDbRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostgresqlPgCreateDbRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostgresqlPgCreateDbRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostgresqlPgCreateDbRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


