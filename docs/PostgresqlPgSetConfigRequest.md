# PostgresqlPgSetConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**Param** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPostgresqlPgSetConfigRequest

`func NewPostgresqlPgSetConfigRequest() *PostgresqlPgSetConfigRequest`

NewPostgresqlPgSetConfigRequest instantiates a new PostgresqlPgSetConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgSetConfigRequestWithDefaults

`func NewPostgresqlPgSetConfigRequestWithDefaults() *PostgresqlPgSetConfigRequest`

NewPostgresqlPgSetConfigRequestWithDefaults instantiates a new PostgresqlPgSetConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *PostgresqlPgSetConfigRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *PostgresqlPgSetConfigRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *PostgresqlPgSetConfigRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *PostgresqlPgSetConfigRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetParam

`func (o *PostgresqlPgSetConfigRequest) GetParam() map[string]string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *PostgresqlPgSetConfigRequest) GetParamOk() (*map[string]string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *PostgresqlPgSetConfigRequest) SetParam(v map[string]string)`

SetParam sets Param field to given value.

### HasParam

`func (o *PostgresqlPgSetConfigRequest) HasParam() bool`

HasParam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


