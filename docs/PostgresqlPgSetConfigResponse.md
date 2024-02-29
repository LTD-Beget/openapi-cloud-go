# PostgresqlPgSetConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**PostgresqlPgConfigInfo**](PostgresqlPgConfigInfo.md) |  | [optional] 
**Error** | Pointer to [**PostgresqlPgSetConfigResponseError**](PostgresqlPgSetConfigResponseError.md) |  | [optional] 

## Methods

### NewPostgresqlPgSetConfigResponse

`func NewPostgresqlPgSetConfigResponse() *PostgresqlPgSetConfigResponse`

NewPostgresqlPgSetConfigResponse instantiates a new PostgresqlPgSetConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgSetConfigResponseWithDefaults

`func NewPostgresqlPgSetConfigResponseWithDefaults() *PostgresqlPgSetConfigResponse`

NewPostgresqlPgSetConfigResponseWithDefaults instantiates a new PostgresqlPgSetConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *PostgresqlPgSetConfigResponse) GetConfig() PostgresqlPgConfigInfo`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PostgresqlPgSetConfigResponse) GetConfigOk() (*PostgresqlPgConfigInfo, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PostgresqlPgSetConfigResponse) SetConfig(v PostgresqlPgConfigInfo)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *PostgresqlPgSetConfigResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetError

`func (o *PostgresqlPgSetConfigResponse) GetError() PostgresqlPgSetConfigResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PostgresqlPgSetConfigResponse) GetErrorOk() (*PostgresqlPgSetConfigResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PostgresqlPgSetConfigResponse) SetError(v PostgresqlPgSetConfigResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *PostgresqlPgSetConfigResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


