# MysqlCreateAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**MysqlDbAccess**](MysqlDbAccess.md) |  | [optional] 
**Error** | Pointer to [**MysqlCreateAccessResponseError**](MysqlCreateAccessResponseError.md) |  | [optional] 

## Methods

### NewMysqlCreateAccessResponse

`func NewMysqlCreateAccessResponse() *MysqlCreateAccessResponse`

NewMysqlCreateAccessResponse instantiates a new MysqlCreateAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlCreateAccessResponseWithDefaults

`func NewMysqlCreateAccessResponseWithDefaults() *MysqlCreateAccessResponse`

NewMysqlCreateAccessResponseWithDefaults instantiates a new MysqlCreateAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *MysqlCreateAccessResponse) GetAccess() MysqlDbAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MysqlCreateAccessResponse) GetAccessOk() (*MysqlDbAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MysqlCreateAccessResponse) SetAccess(v MysqlDbAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *MysqlCreateAccessResponse) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetError

`func (o *MysqlCreateAccessResponse) GetError() MysqlCreateAccessResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MysqlCreateAccessResponse) GetErrorOk() (*MysqlCreateAccessResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MysqlCreateAccessResponse) SetError(v MysqlCreateAccessResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MysqlCreateAccessResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


