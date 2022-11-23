# MysqlSetConfigResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**MysqlConfig**](MysqlConfig.md) |  | [optional] 
**Error** | Pointer to [**MysqlSetConfigResponseError**](MysqlSetConfigResponseError.md) |  | [optional] 

## Methods

### NewMysqlSetConfigResponse

`func NewMysqlSetConfigResponse() *MysqlSetConfigResponse`

NewMysqlSetConfigResponse instantiates a new MysqlSetConfigResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlSetConfigResponseWithDefaults

`func NewMysqlSetConfigResponseWithDefaults() *MysqlSetConfigResponse`

NewMysqlSetConfigResponseWithDefaults instantiates a new MysqlSetConfigResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *MysqlSetConfigResponse) GetConfig() MysqlConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *MysqlSetConfigResponse) GetConfigOk() (*MysqlConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *MysqlSetConfigResponse) SetConfig(v MysqlConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *MysqlSetConfigResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetError

`func (o *MysqlSetConfigResponse) GetError() MysqlSetConfigResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *MysqlSetConfigResponse) GetErrorOk() (*MysqlSetConfigResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *MysqlSetConfigResponse) SetError(v MysqlSetConfigResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *MysqlSetConfigResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


