# MysqlBackupGetOrdersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**[]StructuresMysqlRestoreOrder**](StructuresMysqlRestoreOrder.md) |  | [optional] 

## Methods

### NewMysqlBackupGetOrdersResponse

`func NewMysqlBackupGetOrdersResponse() *MysqlBackupGetOrdersResponse`

NewMysqlBackupGetOrdersResponse instantiates a new MysqlBackupGetOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlBackupGetOrdersResponseWithDefaults

`func NewMysqlBackupGetOrdersResponseWithDefaults() *MysqlBackupGetOrdersResponse`

NewMysqlBackupGetOrdersResponseWithDefaults instantiates a new MysqlBackupGetOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *MysqlBackupGetOrdersResponse) GetOrder() []StructuresMysqlRestoreOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MysqlBackupGetOrdersResponse) GetOrderOk() (*[]StructuresMysqlRestoreOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MysqlBackupGetOrdersResponse) SetOrder(v []StructuresMysqlRestoreOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MysqlBackupGetOrdersResponse) HasOrder() bool`

HasOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


