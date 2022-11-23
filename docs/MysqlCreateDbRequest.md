# MysqlCreateDbRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SavePmaPassword** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewMysqlCreateDbRequest

`func NewMysqlCreateDbRequest() *MysqlCreateDbRequest`

NewMysqlCreateDbRequest instantiates a new MysqlCreateDbRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlCreateDbRequestWithDefaults

`func NewMysqlCreateDbRequestWithDefaults() *MysqlCreateDbRequest`

NewMysqlCreateDbRequestWithDefaults instantiates a new MysqlCreateDbRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *MysqlCreateDbRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MysqlCreateDbRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MysqlCreateDbRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MysqlCreateDbRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDbName

`func (o *MysqlCreateDbRequest) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MysqlCreateDbRequest) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MysqlCreateDbRequest) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *MysqlCreateDbRequest) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetPassword

`func (o *MysqlCreateDbRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MysqlCreateDbRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MysqlCreateDbRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MysqlCreateDbRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSavePmaPassword

`func (o *MysqlCreateDbRequest) GetSavePmaPassword() bool`

GetSavePmaPassword returns the SavePmaPassword field if non-nil, zero value otherwise.

### GetSavePmaPasswordOk

`func (o *MysqlCreateDbRequest) GetSavePmaPasswordOk() (*bool, bool)`

GetSavePmaPasswordOk returns a tuple with the SavePmaPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavePmaPassword

`func (o *MysqlCreateDbRequest) SetSavePmaPassword(v bool)`

SetSavePmaPassword sets SavePmaPassword field to given value.

### HasSavePmaPassword

`func (o *MysqlCreateDbRequest) HasSavePmaPassword() bool`

HasSavePmaPassword returns a boolean if a field has been set.

### GetDescription

`func (o *MysqlCreateDbRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MysqlCreateDbRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MysqlCreateDbRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MysqlCreateDbRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


