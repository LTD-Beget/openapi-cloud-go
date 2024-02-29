# MysqlCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DbName** | Pointer to **string** |  | [optional] 
**AccessPassword** | Pointer to **string** |  | [optional] 
**Param** | Pointer to **map[string]string** |  | [optional] 
**SavePmaPassword** | Pointer to **bool** |  | [optional] 

## Methods

### NewMysqlCreateParams

`func NewMysqlCreateParams() *MysqlCreateParams`

NewMysqlCreateParams instantiates a new MysqlCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlCreateParamsWithDefaults

`func NewMysqlCreateParamsWithDefaults() *MysqlCreateParams`

NewMysqlCreateParamsWithDefaults instantiates a new MysqlCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDbName

`func (o *MysqlCreateParams) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MysqlCreateParams) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MysqlCreateParams) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *MysqlCreateParams) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetAccessPassword

`func (o *MysqlCreateParams) GetAccessPassword() string`

GetAccessPassword returns the AccessPassword field if non-nil, zero value otherwise.

### GetAccessPasswordOk

`func (o *MysqlCreateParams) GetAccessPasswordOk() (*string, bool)`

GetAccessPasswordOk returns a tuple with the AccessPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessPassword

`func (o *MysqlCreateParams) SetAccessPassword(v string)`

SetAccessPassword sets AccessPassword field to given value.

### HasAccessPassword

`func (o *MysqlCreateParams) HasAccessPassword() bool`

HasAccessPassword returns a boolean if a field has been set.

### GetParam

`func (o *MysqlCreateParams) GetParam() map[string]string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *MysqlCreateParams) GetParamOk() (*map[string]string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *MysqlCreateParams) SetParam(v map[string]string)`

SetParam sets Param field to given value.

### HasParam

`func (o *MysqlCreateParams) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetSavePmaPassword

`func (o *MysqlCreateParams) GetSavePmaPassword() bool`

GetSavePmaPassword returns the SavePmaPassword field if non-nil, zero value otherwise.

### GetSavePmaPasswordOk

`func (o *MysqlCreateParams) GetSavePmaPasswordOk() (*bool, bool)`

GetSavePmaPasswordOk returns a tuple with the SavePmaPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavePmaPassword

`func (o *MysqlCreateParams) SetSavePmaPassword(v bool)`

SetSavePmaPassword sets SavePmaPassword field to given value.

### HasSavePmaPassword

`func (o *MysqlCreateParams) HasSavePmaPassword() bool`

HasSavePmaPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


