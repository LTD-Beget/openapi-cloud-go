# CdnCdn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceDomain** | Pointer to [**[]CdnResourceDomain**](CdnResourceDomain.md) |  | [optional] 
**SourceDomain** | Pointer to [**CdnSourceDomain**](CdnSourceDomain.md) |  | [optional] 
**Settings** | Pointer to [**CdnSettings**](CdnSettings.md) |  | [optional] 
**Traffic** | Pointer to **string** |  | [optional] 

## Methods

### NewCdnCdn

`func NewCdnCdn() *CdnCdn`

NewCdnCdn instantiates a new CdnCdn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCdnWithDefaults

`func NewCdnCdnWithDefaults() *CdnCdn`

NewCdnCdnWithDefaults instantiates a new CdnCdn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceDomain

`func (o *CdnCdn) GetResourceDomain() []CdnResourceDomain`

GetResourceDomain returns the ResourceDomain field if non-nil, zero value otherwise.

### GetResourceDomainOk

`func (o *CdnCdn) GetResourceDomainOk() (*[]CdnResourceDomain, bool)`

GetResourceDomainOk returns a tuple with the ResourceDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDomain

`func (o *CdnCdn) SetResourceDomain(v []CdnResourceDomain)`

SetResourceDomain sets ResourceDomain field to given value.

### HasResourceDomain

`func (o *CdnCdn) HasResourceDomain() bool`

HasResourceDomain returns a boolean if a field has been set.

### GetSourceDomain

`func (o *CdnCdn) GetSourceDomain() CdnSourceDomain`

GetSourceDomain returns the SourceDomain field if non-nil, zero value otherwise.

### GetSourceDomainOk

`func (o *CdnCdn) GetSourceDomainOk() (*CdnSourceDomain, bool)`

GetSourceDomainOk returns a tuple with the SourceDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDomain

`func (o *CdnCdn) SetSourceDomain(v CdnSourceDomain)`

SetSourceDomain sets SourceDomain field to given value.

### HasSourceDomain

`func (o *CdnCdn) HasSourceDomain() bool`

HasSourceDomain returns a boolean if a field has been set.

### GetSettings

`func (o *CdnCdn) GetSettings() CdnSettings`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *CdnCdn) GetSettingsOk() (*CdnSettings, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *CdnCdn) SetSettings(v CdnSettings)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *CdnCdn) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTraffic

`func (o *CdnCdn) GetTraffic() string`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *CdnCdn) GetTrafficOk() (*string, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *CdnCdn) SetTraffic(v string)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *CdnCdn) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


