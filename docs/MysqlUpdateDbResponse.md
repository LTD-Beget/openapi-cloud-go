# MysqlUpdateDbResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | Pointer to [**MysqlDb**](MysqlDb.md) |  | [optional] 
**Error** | Pointer to [**MysqlUpdateDbResponseError**](MysqlUpdateDbResponseError.md) |  | [optional] 

## Methods

### NewMysqlUpdateDbResponse

`func NewMysqlUpdateDbResponse() *MysqlUpdateDbResponse`

NewMysqlUpdateDbResponse instantiates a new MysqlUpdateDbResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlUpdateDbResponseWithDefaults

`func NewMysqlUpdateDbResponseWithDefaults() *MysqlUpdateDbResponse`

NewMysqlUpdateDbResponseWithDefaults instantiates a new MysqlUpdateDbResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *MysqlUpdateDbResponse) GetDb() MysqlDb`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *MysqlUpdateDbResponse) GetDbOk() (*MysqlDb, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *MysqlUpdateDbResponse) SetDb(v MysqlDb)`

SetDb sets Db field to given value.

### HasDb

`func (o *MysqlUpdateDbResponse) HasDb() bool`

HasDb returns a boolean if a field has been set.

### GetError

`func (o *MysqlUpdateDbResponse) GetError() MysqlUpdateDbResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MysqlUpdateDbResponse) GetErrorOk() (*MysqlUpdateDbResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MysqlUpdateDbResponse) SetError(v MysqlUpdateDbResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MysqlUpdateDbResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


