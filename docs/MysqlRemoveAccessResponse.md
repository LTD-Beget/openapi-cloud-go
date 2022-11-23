# MysqlRemoveAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to [**MysqlRemoveAccessResponseError**](MysqlRemoveAccessResponseError.md) |  | [optional] 

## Methods

### NewMysqlRemoveAccessResponse

`func NewMysqlRemoveAccessResponse() *MysqlRemoveAccessResponse`

NewMysqlRemoveAccessResponse instantiates a new MysqlRemoveAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlRemoveAccessResponseWithDefaults

`func NewMysqlRemoveAccessResponseWithDefaults() *MysqlRemoveAccessResponse`

NewMysqlRemoveAccessResponseWithDefaults instantiates a new MysqlRemoveAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *MysqlRemoveAccessResponse) GetSuccess() map[string]interface{}`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *MysqlRemoveAccessResponse) GetSuccessOk() (*map[string]interface{}, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *MysqlRemoveAccessResponse) SetSuccess(v map[string]interface{})`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *MysqlRemoveAccessResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *MysqlRemoveAccessResponse) GetError() MysqlRemoveAccessResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MysqlRemoveAccessResponse) GetErrorOk() (*MysqlRemoveAccessResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MysqlRemoveAccessResponse) SetError(v MysqlRemoveAccessResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MysqlRemoveAccessResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


