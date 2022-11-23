# MysqlChangeAccessPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**DbName** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SavePmaPassword** | Pointer to **bool** |  | [optional] 

## Methods

### NewMysqlChangeAccessPasswordRequest

`func NewMysqlChangeAccessPasswordRequest() *MysqlChangeAccessPasswordRequest`

NewMysqlChangeAccessPasswordRequest instantiates a new MysqlChangeAccessPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlChangeAccessPasswordRequestWithDefaults

`func NewMysqlChangeAccessPasswordRequestWithDefaults() *MysqlChangeAccessPasswordRequest`

NewMysqlChangeAccessPasswordRequestWithDefaults instantiates a new MysqlChangeAccessPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *MysqlChangeAccessPasswordRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *MysqlChangeAccessPasswordRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *MysqlChangeAccessPasswordRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *MysqlChangeAccessPasswordRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDbName

`func (o *MysqlChangeAccessPasswordRequest) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *MysqlChangeAccessPasswordRequest) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *MysqlChangeAccessPasswordRequest) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *MysqlChangeAccessPasswordRequest) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetHost

`func (o *MysqlChangeAccessPasswordRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *MysqlChangeAccessPasswordRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *MysqlChangeAccessPasswordRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *MysqlChangeAccessPasswordRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPassword

`func (o *MysqlChangeAccessPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *MysqlChangeAccessPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *MysqlChangeAccessPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *MysqlChangeAccessPasswordRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSavePmaPassword

`func (o *MysqlChangeAccessPasswordRequest) GetSavePmaPassword() bool`

GetSavePmaPassword returns the SavePmaPassword field if non-nil, zero value otherwise.

### GetSavePmaPasswordOk

`func (o *MysqlChangeAccessPasswordRequest) GetSavePmaPasswordOk() (*bool, bool)`

GetSavePmaPasswordOk returns a tuple with the SavePmaPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavePmaPassword

`func (o *MysqlChangeAccessPasswordRequest) SetSavePmaPassword(v bool)`

SetSavePmaPassword sets SavePmaPassword field to given value.

### HasSavePmaPassword

`func (o *MysqlChangeAccessPasswordRequest) HasSavePmaPassword() bool`

HasSavePmaPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


