# CdnPreloadCacheByPathsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdn** | Pointer to [**CdnCdn**](CdnCdn.md) |  | [optional] 
**Error** | Pointer to [**CdnPreloadCacheByPathsResponseError**](CdnPreloadCacheByPathsResponseError.md) |  | [optional] 

## Methods

### NewCdnPreloadCacheByPathsResponse

`func NewCdnPreloadCacheByPathsResponse() *CdnPreloadCacheByPathsResponse`

NewCdnPreloadCacheByPathsResponse instantiates a new CdnPreloadCacheByPathsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnPreloadCacheByPathsResponseWithDefaults

`func NewCdnPreloadCacheByPathsResponseWithDefaults() *CdnPreloadCacheByPathsResponse`

NewCdnPreloadCacheByPathsResponseWithDefaults instantiates a new CdnPreloadCacheByPathsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdn

`func (o *CdnPreloadCacheByPathsResponse) GetCdn() CdnCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *CdnPreloadCacheByPathsResponse) GetCdnOk() (*CdnCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *CdnPreloadCacheByPathsResponse) SetCdn(v CdnCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *CdnPreloadCacheByPathsResponse) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetError

`func (o *CdnPreloadCacheByPathsResponse) GetError() CdnPreloadCacheByPathsResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CdnPreloadCacheByPathsResponse) GetErrorOk() (*CdnPreloadCacheByPathsResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CdnPreloadCacheByPathsResponse) SetError(v CdnPreloadCacheByPathsResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *CdnPreloadCacheByPathsResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


