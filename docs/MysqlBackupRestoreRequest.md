# MysqlBackupRestoreRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopyId** | Pointer to **string** |  | [optional] 
**ServiceUuid** | Pointer to **string** |  | [optional] 

## Methods

### NewMysqlBackupRestoreRequest

`func NewMysqlBackupRestoreRequest() *MysqlBackupRestoreRequest`

NewMysqlBackupRestoreRequest instantiates a new MysqlBackupRestoreRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlBackupRestoreRequestWithDefaults

`func NewMysqlBackupRestoreRequestWithDefaults() *MysqlBackupRestoreRequest`

NewMysqlBackupRestoreRequestWithDefaults instantiates a new MysqlBackupRestoreRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopyId

`func (o *MysqlBackupRestoreRequest) GetCopyId() string`

GetCopyId returns the CopyId field if non-nil, zero value otherwise.

### GetCopyIdOk

`func (o *MysqlBackupRestoreRequest) GetCopyIdOk() (*string, bool)`

GetCopyIdOk returns a tuple with the CopyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyId

`func (o *MysqlBackupRestoreRequest) SetCopyId(v string)`

SetCopyId sets CopyId field to given value.

### HasCopyId

`func (o *MysqlBackupRestoreRequest) HasCopyId() bool`

HasCopyId returns a boolean if a field has been set.

### GetServiceUuid

`func (o *MysqlBackupRestoreRequest) GetServiceUuid() string`

GetServiceUuid returns the ServiceUuid field if non-nil, zero value otherwise.

### GetServiceUuidOk

`func (o *MysqlBackupRestoreRequest) GetServiceUuidOk() (*string, bool)`

GetServiceUuidOk returns a tuple with the ServiceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUuid

`func (o *MysqlBackupRestoreRequest) SetServiceUuid(v string)`

SetServiceUuid sets ServiceUuid field to given value.

### HasServiceUuid

`func (o *MysqlBackupRestoreRequest) HasServiceUuid() bool`

HasServiceUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


