# CdnCdnCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceDomain** | Pointer to **[]string** |  | [optional] 
**SourceDomain** | Pointer to [**CdnSourceDomain**](CdnSourceDomain.md) |  | [optional] 
**Settings** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCdnCdnCreateParams

`func NewCdnCdnCreateParams() *CdnCdnCreateParams`

NewCdnCdnCreateParams instantiates a new CdnCdnCreateParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCdnCreateParamsWithDefaults

`func NewCdnCdnCreateParamsWithDefaults() *CdnCdnCreateParams`

NewCdnCdnCreateParamsWithDefaults instantiates a new CdnCdnCreateParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceDomain

`func (o *CdnCdnCreateParams) GetResourceDomain() []string`

GetResourceDomain returns the ResourceDomain field if non-nil, zero value otherwise.

### GetResourceDomainOk

`func (o *CdnCdnCreateParams) GetResourceDomainOk() (*[]string, bool)`

GetResourceDomainOk returns a tuple with the ResourceDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDomain

`func (o *CdnCdnCreateParams) SetResourceDomain(v []string)`

SetResourceDomain sets ResourceDomain field to given value.

### HasResourceDomain

`func (o *CdnCdnCreateParams) HasResourceDomain() bool`

HasResourceDomain returns a boolean if a field has been set.

### GetSourceDomain

`func (o *CdnCdnCreateParams) GetSourceDomain() CdnSourceDomain`

GetSourceDomain returns the SourceDomain field if non-nil, zero value otherwise.

### GetSourceDomainOk

`func (o *CdnCdnCreateParams) GetSourceDomainOk() (*CdnSourceDomain, bool)`

GetSourceDomainOk returns a tuple with the SourceDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDomain

`func (o *CdnCdnCreateParams) SetSourceDomain(v CdnSourceDomain)`

SetSourceDomain sets SourceDomain field to given value.

### HasSourceDomain

`func (o *CdnCdnCreateParams) HasSourceDomain() bool`

HasSourceDomain returns a boolean if a field has been set.

### GetSettings

`func (o *CdnCdnCreateParams) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CdnCdnCreateParams) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CdnCdnCreateParams) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CdnCdnCreateParams) HasSettings() bool`

HasSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


