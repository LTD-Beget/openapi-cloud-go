# CdnChangeSettingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**Settings** | Pointer to [**CdnSettings**](CdnSettings.md) |  | [optional] 

## Methods

### NewCdnChangeSettingRequest

`func NewCdnChangeSettingRequest() *CdnChangeSettingRequest`

NewCdnChangeSettingRequest instantiates a new CdnChangeSettingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnChangeSettingRequestWithDefaults

`func NewCdnChangeSettingRequestWithDefaults() *CdnChangeSettingRequest`

NewCdnChangeSettingRequestWithDefaults instantiates a new CdnChangeSettingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CdnChangeSettingRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CdnChangeSettingRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CdnChangeSettingRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *CdnChangeSettingRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetSettings

`func (o *CdnChangeSettingRequest) GetSettings() CdnSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CdnChangeSettingRequest) GetSettingsOk() (*CdnSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CdnChangeSettingRequest) SetSettings(v CdnSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CdnChangeSettingRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


