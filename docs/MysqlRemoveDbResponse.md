# MysqlRemoveDbResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to [**MysqlRemoveDbResponseError**](MysqlRemoveDbResponseError.md) |  | [optional] 

## Methods

### NewMysqlRemoveDbResponse

`func NewMysqlRemoveDbResponse() *MysqlRemoveDbResponse`

NewMysqlRemoveDbResponse instantiates a new MysqlRemoveDbResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlRemoveDbResponseWithDefaults

`func NewMysqlRemoveDbResponseWithDefaults() *MysqlRemoveDbResponse`

NewMysqlRemoveDbResponseWithDefaults instantiates a new MysqlRemoveDbResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MysqlRemoveDbResponse) GetSuccess() map[string]interface{}`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MysqlRemoveDbResponse) GetSuccessOk() (*map[string]interface{}, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MysqlRemoveDbResponse) SetSuccess(v map[string]interface{})`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *MysqlRemoveDbResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *MysqlRemoveDbResponse) GetError() MysqlRemoveDbResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MysqlRemoveDbResponse) GetErrorOk() (*MysqlRemoveDbResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MysqlRemoveDbResponse) SetError(v MysqlRemoveDbResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MysqlRemoveDbResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


