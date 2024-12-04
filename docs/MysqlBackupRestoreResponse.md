# MysqlBackupRestoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**StructuresMysqlRestoreOrder**](StructuresMysqlRestoreOrder.md) |  | [optional] 
**Error** | Pointer to [**MysqlBackupRestoreResponseError**](MysqlBackupRestoreResponseError.md) |  | [optional] 

## Methods

### NewMysqlBackupRestoreResponse

`func NewMysqlBackupRestoreResponse() *MysqlBackupRestoreResponse`

NewMysqlBackupRestoreResponse instantiates a new MysqlBackupRestoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlBackupRestoreResponseWithDefaults

`func NewMysqlBackupRestoreResponseWithDefaults() *MysqlBackupRestoreResponse`

NewMysqlBackupRestoreResponseWithDefaults instantiates a new MysqlBackupRestoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *MysqlBackupRestoreResponse) GetOrder() StructuresMysqlRestoreOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MysqlBackupRestoreResponse) GetOrderOk() (*StructuresMysqlRestoreOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MysqlBackupRestoreResponse) SetOrder(v StructuresMysqlRestoreOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MysqlBackupRestoreResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetError

`func (o *MysqlBackupRestoreResponse) GetError() MysqlBackupRestoreResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MysqlBackupRestoreResponse) GetErrorOk() (*MysqlBackupRestoreResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MysqlBackupRestoreResponse) SetError(v MysqlBackupRestoreResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MysqlBackupRestoreResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


