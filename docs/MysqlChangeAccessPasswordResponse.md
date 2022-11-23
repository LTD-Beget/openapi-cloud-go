# MysqlChangeAccessPasswordResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**MysqlDbAccess**](MysqlDbAccess.md) |  | [optional] 
**Error** | Pointer to [**MysqlChangeAccessPasswordResponseError**](MysqlChangeAccessPasswordResponseError.md) |  | [optional] 

## Methods

### NewMysqlChangeAccessPasswordResponse

`func NewMysqlChangeAccessPasswordResponse() *MysqlChangeAccessPasswordResponse`

NewMysqlChangeAccessPasswordResponse instantiates a new MysqlChangeAccessPasswordResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlChangeAccessPasswordResponseWithDefaults

`func NewMysqlChangeAccessPasswordResponseWithDefaults() *MysqlChangeAccessPasswordResponse`

NewMysqlChangeAccessPasswordResponseWithDefaults instantiates a new MysqlChangeAccessPasswordResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *MysqlChangeAccessPasswordResponse) GetAccess() MysqlDbAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MysqlChangeAccessPasswordResponse) GetAccessOk() (*MysqlDbAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MysqlChangeAccessPasswordResponse) SetAccess(v MysqlDbAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *MysqlChangeAccessPasswordResponse) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetError

`func (o *MysqlChangeAccessPasswordResponse) GetError() MysqlChangeAccessPasswordResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MysqlChangeAccessPasswordResponse) GetErrorOk() (*MysqlChangeAccessPasswordResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MysqlChangeAccessPasswordResponse) SetError(v MysqlChangeAccessPasswordResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MysqlChangeAccessPasswordResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


