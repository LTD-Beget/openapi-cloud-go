# StructuresPostgresqlRestoreOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CopyInfo** | Pointer to [**StructuresPostgresqlCopy**](StructuresPostgresqlCopy.md) |  | [optional] 
**BackupDateCreate** | Pointer to **string** |  | [optional] 
**ServiceId** | Pointer to **string** |  | [optional] 
**ServiceName** | Pointer to **string** |  | [optional] 
**DateCreate** | Pointer to **string** |  | [optional] 
**DateComplete** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewStructuresPostgresqlRestoreOrder

`func NewStructuresPostgresqlRestoreOrder() *StructuresPostgresqlRestoreOrder`

NewStructuresPostgresqlRestoreOrder instantiates a new StructuresPostgresqlRestoreOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuresPostgresqlRestoreOrderWithDefaults

`func NewStructuresPostgresqlRestoreOrderWithDefaults() *StructuresPostgresqlRestoreOrder`

NewStructuresPostgresqlRestoreOrderWithDefaults instantiates a new StructuresPostgresqlRestoreOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StructuresPostgresqlRestoreOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StructuresPostgresqlRestoreOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StructuresPostgresqlRestoreOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StructuresPostgresqlRestoreOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCopyInfo

`func (o *StructuresPostgresqlRestoreOrder) GetCopyInfo() StructuresPostgresqlCopy`

GetCopyInfo returns the CopyInfo field if non-nil, zero value otherwise.

### GetCopyInfoOk

`func (o *StructuresPostgresqlRestoreOrder) GetCopyInfoOk() (*StructuresPostgresqlCopy, bool)`

GetCopyInfoOk returns a tuple with the CopyInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyInfo

`func (o *StructuresPostgresqlRestoreOrder) SetCopyInfo(v StructuresPostgresqlCopy)`

SetCopyInfo sets CopyInfo field to given value.

### HasCopyInfo

`func (o *StructuresPostgresqlRestoreOrder) HasCopyInfo() bool`

HasCopyInfo returns a boolean if a field has been set.

### GetBackupDateCreate

`func (o *StructuresPostgresqlRestoreOrder) GetBackupDateCreate() string`

GetBackupDateCreate returns the BackupDateCreate field if non-nil, zero value otherwise.

### GetBackupDateCreateOk

`func (o *StructuresPostgresqlRestoreOrder) GetBackupDateCreateOk() (*string, bool)`

GetBackupDateCreateOk returns a tuple with the BackupDateCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupDateCreate

`func (o *StructuresPostgresqlRestoreOrder) SetBackupDateCreate(v string)`

SetBackupDateCreate sets BackupDateCreate field to given value.

### HasBackupDateCreate

`func (o *StructuresPostgresqlRestoreOrder) HasBackupDateCreate() bool`

HasBackupDateCreate returns a boolean if a field has been set.

### GetServiceId

`func (o *StructuresPostgresqlRestoreOrder) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *StructuresPostgresqlRestoreOrder) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *StructuresPostgresqlRestoreOrder) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *StructuresPostgresqlRestoreOrder) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetServiceName

`func (o *StructuresPostgresqlRestoreOrder) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *StructuresPostgresqlRestoreOrder) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *StructuresPostgresqlRestoreOrder) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *StructuresPostgresqlRestoreOrder) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetDateCreate

`func (o *StructuresPostgresqlRestoreOrder) GetDateCreate() string`

GetDateCreate returns the DateCreate field if non-nil, zero value otherwise.

### GetDateCreateOk

`func (o *StructuresPostgresqlRestoreOrder) GetDateCreateOk() (*string, bool)`

GetDateCreateOk returns a tuple with the DateCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreate

`func (o *StructuresPostgresqlRestoreOrder) SetDateCreate(v string)`

SetDateCreate sets DateCreate field to given value.

### HasDateCreate

`func (o *StructuresPostgresqlRestoreOrder) HasDateCreate() bool`

HasDateCreate returns a boolean if a field has been set.

### GetDateComplete

`func (o *StructuresPostgresqlRestoreOrder) GetDateComplete() string`

GetDateComplete returns the DateComplete field if non-nil, zero value otherwise.

### GetDateCompleteOk

`func (o *StructuresPostgresqlRestoreOrder) GetDateCompleteOk() (*string, bool)`

GetDateCompleteOk returns a tuple with the DateComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateComplete

`func (o *StructuresPostgresqlRestoreOrder) SetDateComplete(v string)`

SetDateComplete sets DateComplete field to given value.

### HasDateComplete

`func (o *StructuresPostgresqlRestoreOrder) HasDateComplete() bool`

HasDateComplete returns a boolean if a field has been set.

### GetStatus

`func (o *StructuresPostgresqlRestoreOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StructuresPostgresqlRestoreOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StructuresPostgresqlRestoreOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StructuresPostgresqlRestoreOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


