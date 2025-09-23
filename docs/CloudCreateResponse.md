# CloudCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**CloudService**](CloudService.md) |  | [optional] 
**MysqlError** | Pointer to [**MysqlCreateError**](MysqlCreateError.md) |  | [optional] 
**PostgresqlError** | Pointer to [**PostgresqlPgCreateError**](PostgresqlPgCreateError.md) |  | [optional] 
**S3Error** | Pointer to [**S3S3CreateError**](S3S3CreateError.md) |  | [optional] 
**CdnError** | Pointer to [**CdnCdnCreateError**](CdnCdnCreateError.md) |  | [optional] 

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

### GetS3Error

`func (o *CloudCreateResponse) GetS3Error() S3S3CreateError`

GetS3Error returns the S3Error field if non-nil, zero value otherwise.

### GetS3ErrorOk

`func (o *CloudCreateResponse) GetS3ErrorOk() (*S3S3CreateError, bool)`

GetS3ErrorOk returns a tuple with the S3Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Error

`func (o *CloudCreateResponse) SetS3Error(v S3S3CreateError)`

SetS3Error sets S3Error field to given value.

### HasS3Error

`func (o *CloudCreateResponse) HasS3Error() bool`

HasS3Error returns a boolean if a field has been set.

### GetCdnError

`func (o *CloudCreateResponse) GetCdnError() CdnCdnCreateError`

GetCdnError returns the CdnError field if non-nil, zero value otherwise.

### GetCdnErrorOk

`func (o *CloudCreateResponse) GetCdnErrorOk() (*CdnCdnCreateError, bool)`

GetCdnErrorOk returns a tuple with the CdnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdnError

`func (o *CloudCreateResponse) SetCdnError(v CdnCdnCreateError)`

SetCdnError sets CdnError field to given value.

### HasCdnError

`func (o *CloudCreateResponse) HasCdnError() bool`

HasCdnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


