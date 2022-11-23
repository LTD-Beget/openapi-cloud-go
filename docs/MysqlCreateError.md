# MysqlCreateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewMysqlCreateError

`func NewMysqlCreateError() *MysqlCreateError`

NewMysqlCreateError instantiates a new MysqlCreateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlCreateErrorWithDefaults

`func NewMysqlCreateErrorWithDefaults() *MysqlCreateError`

NewMysqlCreateErrorWithDefaults instantiates a new MysqlCreateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MysqlCreateError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MysqlCreateError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MysqlCreateError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MysqlCreateError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *MysqlCreateError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MysqlCreateError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MysqlCreateError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MysqlCreateError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


