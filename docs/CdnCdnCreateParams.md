# CdnCdnCreateParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceDomain** | Pointer to **[]string** |  | [optional] 
**SourceDomain** | Pointer to [**CdnSourceDomain**](CdnSourceDomain.md) |  | [optional] 
**SourceParams** | Pointer to [**CdnCdnCreateParamsSourceParams**](CdnCdnCreateParamsSourceParams.md) |  | [optional] 

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

### GetSourceParams

`func (o *CdnCdnCreateParams) GetSourceParams() CdnCdnCreateParamsSourceParams`

GetSourceParams returns the SourceParams field if non-nil, zero value otherwise.

### GetSourceParamsOk

`func (o *CdnCdnCreateParams) GetSourceParamsOk() (*CdnCdnCreateParamsSourceParams, bool)`

GetSourceParamsOk returns a tuple with the SourceParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceParams

`func (o *CdnCdnCreateParams) SetSourceParams(v CdnCdnCreateParamsSourceParams)`

SetSourceParams sets SourceParams field to given value.

### HasSourceParams

`func (o *CdnCdnCreateParams) HasSourceParams() bool`

HasSourceParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


