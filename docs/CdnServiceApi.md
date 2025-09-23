# \CdnServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CdnServiceChangeResourceDomains**](CdnServiceApi.md#CdnServiceChangeResourceDomains) | **Post** /v1/cloud/cdn/{service_id}/resource-domains | 
[**CdnServiceChangeSetting**](CdnServiceApi.md#CdnServiceChangeSetting) | **Post** /v1/cloud/cdn/{service_id}/setting | 
[**CdnServiceGetPrice**](CdnServiceApi.md#CdnServiceGetPrice) | **Get** /v1/cloud/cdn/price | 
[**CdnServiceGetSourceDomains**](CdnServiceApi.md#CdnServiceGetSourceDomains) | **Get** /v1/cloud/cdn/source-domains | 
[**CdnServicePreloadCacheByPaths**](CdnServiceApi.md#CdnServicePreloadCacheByPaths) | **Post** /v1/cloud/cdn/{service_id}/preload-cache-by-paths | 
[**CdnServicePurgeAllCache**](CdnServiceApi.md#CdnServicePurgeAllCache) | **Get** /v1/cloud/cdn/{service_id}/purge-all-cache | 
[**CdnServicePurgeCacheByPaths**](CdnServiceApi.md#CdnServicePurgeCacheByPaths) | **Post** /v1/cloud/cdn/{service_id}/purge-cache-by-paths | 



## CdnServiceChangeResourceDomains

> CdnChangeResourceDomainsResponse CdnServiceChangeResourceDomains(ctx, serviceId).CdnChangeResourceDomainsRequest(cdnChangeResourceDomainsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | 
    cdnChangeResourceDomainsRequest := *openapiclient.NewCdnChangeResourceDomainsRequest() // CdnChangeResourceDomainsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnServiceApi.CdnServiceChangeResourceDomains(context.Background(), serviceId).CdnChangeResourceDomainsRequest(cdnChangeResourceDomainsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnServiceApi.CdnServiceChangeResourceDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnServiceChangeResourceDomains`: CdnChangeResourceDomainsResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnServiceApi.CdnServiceChangeResourceDomains`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCdnServiceChangeResourceDomainsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnChangeResourceDomainsRequest** | [**CdnChangeResourceDomainsRequest**](CdnChangeResourceDomainsRequest.md) |  | 

### Return type

[**CdnChangeResourceDomainsResponse**](CdnChangeResourceDomainsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CdnServiceChangeSetting

> CdnChangeSettingResponse CdnServiceChangeSetting(ctx, serviceId).CdnChangeSettingRequest(cdnChangeSettingRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | 
    cdnChangeSettingRequest := *openapiclient.NewCdnChangeSettingRequest() // CdnChangeSettingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnServiceApi.CdnServiceChangeSetting(context.Background(), serviceId).CdnChangeSettingRequest(cdnChangeSettingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnServiceApi.CdnServiceChangeSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnServiceChangeSetting`: CdnChangeSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnServiceApi.CdnServiceChangeSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCdnServiceChangeSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnChangeSettingRequest** | [**CdnChangeSettingRequest**](CdnChangeSettingRequest.md) |  | 

### Return type

[**CdnChangeSettingResponse**](CdnChangeSettingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CdnServiceGetPrice

> CdnGetPriceResponse CdnServiceGetPrice(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnServiceApi.CdnServiceGetPrice(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnServiceApi.CdnServiceGetPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnServiceGetPrice`: CdnGetPriceResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnServiceApi.CdnServiceGetPrice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCdnServiceGetPriceRequest struct via the builder pattern


### Return type

[**CdnGetPriceResponse**](CdnGetPriceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CdnServiceGetSourceDomains

> CdnGetSourceDomainsResponse CdnServiceGetSourceDomains(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnServiceApi.CdnServiceGetSourceDomains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnServiceApi.CdnServiceGetSourceDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnServiceGetSourceDomains`: CdnGetSourceDomainsResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnServiceApi.CdnServiceGetSourceDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCdnServiceGetSourceDomainsRequest struct via the builder pattern


### Return type

[**CdnGetSourceDomainsResponse**](CdnGetSourceDomainsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CdnServicePreloadCacheByPaths

> CdnPreloadCacheByPathsResponse CdnServicePreloadCacheByPaths(ctx, serviceId).CdnPreloadCacheByPathsRequest(cdnPreloadCacheByPathsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | 
    cdnPreloadCacheByPathsRequest := *openapiclient.NewCdnPreloadCacheByPathsRequest() // CdnPreloadCacheByPathsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnServiceApi.CdnServicePreloadCacheByPaths(context.Background(), serviceId).CdnPreloadCacheByPathsRequest(cdnPreloadCacheByPathsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnServiceApi.CdnServicePreloadCacheByPaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnServicePreloadCacheByPaths`: CdnPreloadCacheByPathsResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnServiceApi.CdnServicePreloadCacheByPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCdnServicePreloadCacheByPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnPreloadCacheByPathsRequest** | [**CdnPreloadCacheByPathsRequest**](CdnPreloadCacheByPathsRequest.md) |  | 

### Return type

[**CdnPreloadCacheByPathsResponse**](CdnPreloadCacheByPathsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CdnServicePurgeAllCache

> CdnPurgeAllCacheResponse CdnServicePurgeAllCache(ctx, serviceId).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnServiceApi.CdnServicePurgeAllCache(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnServiceApi.CdnServicePurgeAllCache``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnServicePurgeAllCache`: CdnPurgeAllCacheResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnServiceApi.CdnServicePurgeAllCache`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCdnServicePurgeAllCacheRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CdnPurgeAllCacheResponse**](CdnPurgeAllCacheResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CdnServicePurgeCacheByPaths

> CdnPurgeCacheByPathsResponse CdnServicePurgeCacheByPaths(ctx, serviceId).CdnPurgeCacheByPathsRequest(cdnPurgeCacheByPathsRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceId := "serviceId_example" // string | 
    cdnPurgeCacheByPathsRequest := *openapiclient.NewCdnPurgeCacheByPathsRequest() // CdnPurgeCacheByPathsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CdnServiceApi.CdnServicePurgeCacheByPaths(context.Background(), serviceId).CdnPurgeCacheByPathsRequest(cdnPurgeCacheByPathsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnServiceApi.CdnServicePurgeCacheByPaths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnServicePurgeCacheByPaths`: CdnPurgeCacheByPathsResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnServiceApi.CdnServicePurgeCacheByPaths`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCdnServicePurgeCacheByPathsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cdnPurgeCacheByPathsRequest** | [**CdnPurgeCacheByPathsRequest**](CdnPurgeCacheByPathsRequest.md) |  | 

### Return type

[**CdnPurgeCacheByPathsResponse**](CdnPurgeCacheByPathsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

