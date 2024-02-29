# CloudCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**CloudService**](CloudService.md) |  | [optional] 
**MysqlError** | Pointer to [**MysqlCreateError**](MysqlCreateError.md) |  | [optional] 
**PostgresqlError** | Pointer to [**PostgresqlPgCreateError**](PostgresqlPgCreateError.md) |  | [optional] 

## Methods

### NewCloudCreateResponse

`func NewCloudCreateResponse() *CloudCreateResponse`

NewCloudCreateResponse instantiates a new CloudCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateResponseWithDefaults

`func NewCloudCreateResponseWithDefaults() *CloudCreateResponse`

NewCloudCreateResponseWithDefaults instantiates a new CloudCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CloudCreateResponse) GetService() CloudService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudCreateResponse) GetServiceOk() (*CloudService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudCreateResponse) SetService(v CloudService)`

SetService sets Service field to given value.

### HasService

`func (o *CloudCreateResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetMysqlError

`func (o *CloudCreateResponse) GetMysqlError() MysqlCreateError`

GetMysqlError returns the MysqlError field if non-nil, zero value otherwise.

### GetMysqlErrorOk

`func (o *CloudCreateResponse) GetMysqlErrorOk() (*MysqlCreateError, bool)`

GetMysqlErrorOk returns a tuple with the MysqlError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlError

`func (o *CloudCreateResponse) SetMysqlError(v MysqlCreateError)`

SetMysqlError sets MysqlError field to given value.

### HasMysqlError

`func (o *CloudCreateResponse) HasMysqlError() bool`

HasMysqlError returns a boolean if a field has been set.

### GetPostgresqlError

`func (o *CloudCreateResponse) GetPostgresqlError() PostgresqlPgCreateError`

GetPostgresqlError returns the PostgresqlError field if non-nil, zero value otherwise.

### GetPostgresqlErrorOk

`func (o *CloudCreateResponse) GetPostgresqlErrorOk() (*PostgresqlPgCreateError, bool)`

GetPostgresqlErrorOk returns a tuple with the PostgresqlError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresqlError

`func (o *CloudCreateResponse) SetPostgresqlError(v PostgresqlPgCreateError)`

SetPostgresqlError sets PostgresqlError field to given value.

### HasPostgresqlError

`func (o *CloudCreateResponse) HasPostgresqlError() bool`

HasPostgresqlError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


