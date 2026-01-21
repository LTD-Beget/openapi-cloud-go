# CdnChangeResourceDomainsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceDomain** | Pointer to [**CdnChangeResourceDomainsResponseResourceDomains**](CdnChangeResourceDomainsResponseResourceDomains.md) |  | [optional] 
**DomainErrors** | Pointer to [**CdnChangeResourceDomainsResponseDomainErrors**](CdnChangeResourceDomainsResponseDomainErrors.md) |  | [optional] 
**Error** | Pointer to [**CdnChangeResourceDomainsResponseError**](CdnChangeResourceDomainsResponseError.md) |  | [optional] 

## Methods

### NewCdnChangeResourceDomainsResponse

`func NewCdnChangeResourceDomainsResponse() *CdnChangeResourceDomainsResponse`

NewCdnChangeResourceDomainsResponse instantiates a new CdnChangeResourceDomainsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnChangeResourceDomainsResponseWithDefaults

`func NewCdnChangeResourceDomainsResponseWithDefaults() *CdnChangeResourceDomainsResponse`

NewCdnChangeResourceDomainsResponseWithDefaults instantiates a new CdnChangeResourceDomainsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceDomain

`func (o *CdnChangeResourceDomainsResponse) GetResourceDomain() CdnChangeResourceDomainsResponseResourceDomains`

GetResourceDomain returns the ResourceDomain field if non-nil, zero value otherwise.

### GetResourceDomainOk

`func (o *CdnChangeResourceDomainsResponse) GetResourceDomainOk() (*CdnChangeResourceDomainsResponseResourceDomains, bool)`

GetResourceDomainOk returns a tuple with the ResourceDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceDomain

`func (o *CdnChangeResourceDomainsResponse) SetResourceDomain(v CdnChangeResourceDomainsResponseResourceDomains)`

SetResourceDomain sets ResourceDomain field to given value.

### HasResourceDomain

`func (o *CdnChangeResourceDomainsResponse) HasResourceDomain() bool`

HasResourceDomain returns a boolean if a field has been set.

### GetDomainErrors

`func (o *CdnChangeResourceDomainsResponse) GetDomainErrors() CdnChangeResourceDomainsResponseDomainErrors`

GetDomainErrors returns the DomainErrors field if non-nil, zero value otherwise.

### GetDomainErrorsOk

`func (o *CdnChangeResourceDomainsResponse) GetDomainErrorsOk() (*CdnChangeResourceDomainsResponseDomainErrors, bool)`

GetDomainErrorsOk returns a tuple with the DomainErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainErrors

`func (o *CdnChangeResourceDomainsResponse) SetDomainErrors(v CdnChangeResourceDomainsResponseDomainErrors)`

SetDomainErrors sets DomainErrors field to given value.

### HasDomainErrors

`func (o *CdnChangeResourceDomainsResponse) HasDomainErrors() bool`

HasDomainErrors returns a boolean if a field has been set.

### GetError

`func (o *CdnChangeResourceDomainsResponse) GetError() CdnChangeResourceDomainsResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CdnChangeResourceDomainsResponse) GetErrorOk() (*CdnChangeResourceDomainsResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CdnChangeResourceDomainsResponse) SetError(v CdnChangeResourceDomainsResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *CdnChangeResourceDomainsResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


