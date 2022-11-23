# MysqlCreateDbResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewMysqlCreateDbResponseError

`func NewMysqlCreateDbResponseError() *MysqlCreateDbResponseError`

NewMysqlCreateDbResponseError instantiates a new MysqlCreateDbResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlCreateDbResponseErrorWithDefaults

`func NewMysqlCreateDbResponseErrorWithDefaults() *MysqlCreateDbResponseError`

NewMysqlCreateDbResponseErrorWithDefaults instantiates a new MysqlCreateDbResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MysqlCreateDbResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MysqlCreateDbResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MysqlCreateDbResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MysqlCreateDbResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *MysqlCreateDbResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MysqlCreateDbResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MysqlCreateDbResponseError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *MysqlCreateDbResponseError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


