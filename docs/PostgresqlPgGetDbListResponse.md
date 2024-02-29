# PostgresqlPgGetDbListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | Pointer to [**[]PostgresqlPgDb**](PostgresqlPgDb.md) |  | [optional] 

## Methods

### NewPostgresqlPgGetDbListResponse

`func NewPostgresqlPgGetDbListResponse() *PostgresqlPgGetDbListResponse`

NewPostgresqlPgGetDbListResponse instantiates a new PostgresqlPgGetDbListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostgresqlPgGetDbListResponseWithDefaults

`func NewPostgresqlPgGetDbListResponseWithDefaults() *PostgresqlPgGetDbListResponse`

NewPostgresqlPgGetDbListResponseWithDefaults instantiates a new PostgresqlPgGetDbListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *PostgresqlPgGetDbListResponse) GetDb() []PostgresqlPgDb`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *PostgresqlPgGetDbListResponse) GetDbOk() (*[]PostgresqlPgDb, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *PostgresqlPgGetDbListResponse) SetDb(v []PostgresqlPgDb)`

SetDb sets Db field to given value.

### HasDb

`func (o *PostgresqlPgGetDbListResponse) HasDb() bool`

HasDb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


