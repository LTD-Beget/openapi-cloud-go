# CdnPurgeCacheByPathsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdn** | Pointer to [**CdnCdn**](CdnCdn.md) |  | [optional] 
**Error** | Pointer to [**CdnPurgeCacheByPathsResponseError**](CdnPurgeCacheByPathsResponseError.md) |  | [optional] 

## Methods

### NewCdnPurgeCacheByPathsResponse

`func NewCdnPurgeCacheByPathsResponse() *CdnPurgeCacheByPathsResponse`

NewCdnPurgeCacheByPathsResponse instantiates a new CdnPurgeCacheByPathsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnPurgeCacheByPathsResponseWithDefaults

`func NewCdnPurgeCacheByPathsResponseWithDefaults() *CdnPurgeCacheByPathsResponse`

NewCdnPurgeCacheByPathsResponseWithDefaults instantiates a new CdnPurgeCacheByPathsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdn

`func (o *CdnPurgeCacheByPathsResponse) GetCdn() CdnCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *CdnPurgeCacheByPathsResponse) GetCdnOk() (*CdnCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *CdnPurgeCacheByPathsResponse) SetCdn(v CdnCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *CdnPurgeCacheByPathsResponse) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetError

`func (o *CdnPurgeCacheByPathsResponse) GetError() CdnPurgeCacheByPathsResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CdnPurgeCacheByPathsResponse) GetErrorOk() (*CdnPurgeCacheByPathsResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CdnPurgeCacheByPathsResponse) SetError(v CdnPurgeCacheByPathsResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *CdnPurgeCacheByPathsResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


