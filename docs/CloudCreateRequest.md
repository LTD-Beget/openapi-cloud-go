# CloudCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigurationId** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MysqlParams** | Pointer to [**MysqlCreateParams**](MysqlCreateParams.md) |  | [optional] 

## Methods

### NewCloudCreateRequest

`func NewCloudCreateRequest() *CloudCreateRequest`

NewCloudCreateRequest instantiates a new CloudCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateRequestWithDefaults

`func NewCloudCreateRequestWithDefaults() *CloudCreateRequest`

NewCloudCreateRequestWithDefaults instantiates a new CloudCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigurationId

`func (o *CloudCreateRequest) GetConfigurationId() string`

GetConfigurationId returns the ConfigurationId field if non-nil, zero value otherwise.

### GetConfigurationIdOk

`func (o *CloudCreateRequest) GetConfigurationIdOk() (*string, bool)`

GetConfigurationIdOk returns a tuple with the ConfigurationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationId

`func (o *CloudCreateRequest) SetConfigurationId(v string)`

SetConfigurationId sets ConfigurationId field to given value.

### HasConfigurationId

`func (o *CloudCreateRequest) HasConfigurationId() bool`

HasConfigurationId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudCreateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudCreateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudCreateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudCreateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *CloudCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CloudCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CloudCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CloudCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMysqlParams

`func (o *CloudCreateRequest) GetMysqlParams() MysqlCreateParams`

GetMysqlParams returns the MysqlParams field if non-nil, zero value otherwise.

### GetMysqlParamsOk

`func (o *CloudCreateRequest) GetMysqlParamsOk() (*MysqlCreateParams, bool)`

GetMysqlParamsOk returns a tuple with the MysqlParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysqlParams

`func (o *CloudCreateRequest) SetMysqlParams(v MysqlCreateParams)`

SetMysqlParams sets MysqlParams field to given value.

### HasMysqlParams

`func (o *CloudCreateRequest) HasMysqlParams() bool`

HasMysqlParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


