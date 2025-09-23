# CdnChangeResourceDomainsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCdnChangeResourceDomainsRequest

`func NewCdnChangeResourceDomainsRequest() *CdnChangeResourceDomainsRequest`

NewCdnChangeResourceDomainsRequest instantiates a new CdnChangeResourceDomainsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnChangeResourceDomainsRequestWithDefaults

`func NewCdnChangeResourceDomainsRequestWithDefaults() *CdnChangeResourceDomainsRequest`

NewCdnChangeResourceDomainsRequestWithDefaults instantiates a new CdnChangeResourceDomainsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *CdnChangeResourceDomainsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *CdnChangeResourceDomainsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *CdnChangeResourceDomainsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *CdnChangeResourceDomainsRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetDomain

`func (o *CdnChangeResourceDomainsRequest) GetDomain() []string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *CdnChangeResourceDomainsRequest) GetDomainOk() (*[]string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *CdnChangeResourceDomainsRequest) SetDomain(v []string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *CdnChangeResourceDomainsRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


