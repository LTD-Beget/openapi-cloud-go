# MysqlDb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **string** |  | [optional] 
**Access** | Pointer to [**[]MysqlDbAccess**](MysqlDbAccess.md) |  | [optional] 
**PmaPasswordSaved** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewMysqlDb

`func NewMysqlDb() *MysqlDb`

NewMysqlDb instantiates a new MysqlDb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlDbWithDefaults

`func NewMysqlDbWithDefaults() *MysqlDb`

NewMysqlDbWithDefaults instantiates a new MysqlDb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MysqlDb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MysqlDb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MysqlDb) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MysqlDb) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSize

`func (o *MysqlDb) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MysqlDb) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MysqlDb) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *MysqlDb) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetAccess

`func (o *MysqlDb) GetAccess() []MysqlDbAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *MysqlDb) GetAccessOk() (*[]MysqlDbAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *MysqlDb) SetAccess(v []MysqlDbAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *MysqlDb) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetPmaPasswordSaved

`func (o *MysqlDb) GetPmaPasswordSaved() bool`

GetPmaPasswordSaved returns the PmaPasswordSaved field if non-nil, zero value otherwise.

### GetPmaPasswordSavedOk

`func (o *MysqlDb) GetPmaPasswordSavedOk() (*bool, bool)`

GetPmaPasswordSavedOk returns a tuple with the PmaPasswordSaved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmaPasswordSaved

`func (o *MysqlDb) SetPmaPasswordSaved(v bool)`

SetPmaPasswordSaved sets PmaPasswordSaved field to given value.

### HasPmaPasswordSaved

`func (o *MysqlDb) HasPmaPasswordSaved() bool`

HasPmaPasswordSaved returns a boolean if a field has been set.

### GetDescription

`func (o *MysqlDb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlDb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlDb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MysqlDb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


