# \S3ServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3ServiceChangeAccessKey**](S3ServiceApi.md#S3ServiceChangeAccessKey) | **Patch** /v1/cloud/s3/{service_id}/access-key | 
[**S3ServiceChangeCors**](S3ServiceApi.md#S3ServiceChangeCors) | **Patch** /v1/cloud/s3/{service_id}/cors | 
[**S3ServiceChangeDomain**](S3ServiceApi.md#S3ServiceChangeDomain) | **Patch** /v1/cloud/s3/{service_id}/domain | 
[**S3ServiceChangePublic**](S3ServiceApi.md#S3ServiceChangePublic) | **Patch** /v1/cloud/s3/{service_id}/public | 
[**S3ServiceGetPrefix**](S3ServiceApi.md#S3ServiceGetPrefix) | **Get** /v1/cloud/s3/prefix | 
[**S3ServiceGetPrice**](S3ServiceApi.md#S3ServiceGetPrice) | **Get** /v1/cloud/s3/price | 
[**S3ServiceGetQuota**](S3ServiceApi.md#S3ServiceGetQuota) | **Get** /v1/cloud/s3/quota | 



## S3ServiceChangeAccessKey

> S3ChangeAccessKeyResponse S3ServiceChangeAccessKey(ctx, serviceId).S3ChangeAccessKeyRequest(s3ChangeAccessKeyRequest).Execute()



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
    s3ChangeAccessKeyRequest := *openapiclient.NewS3ChangeAccessKeyRequest() // S3ChangeAccessKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3ServiceApi.S3ServiceChangeAccessKey(context.Background(), serviceId).S3ChangeAccessKeyRequest(s3ChangeAccessKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3ServiceApi.S3ServiceChangeAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3ServiceChangeAccessKey`: S3ChangeAccessKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `S3ServiceApi.S3ServiceChangeAccessKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3ServiceChangeAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3ChangeAccessKeyRequest** | [**S3ChangeAccessKeyRequest**](S3ChangeAccessKeyRequest.md) |  | 

### Return type

[**S3ChangeAccessKeyResponse**](S3ChangeAccessKeyResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3ServiceChangeCors

> S3ChangeCorsResponse S3ServiceChangeCors(ctx, serviceId).S3ChangeCorsRequest(s3ChangeCorsRequest).Execute()



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
    s3ChangeCorsRequest := *openapiclient.NewS3ChangeCorsRequest() // S3ChangeCorsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3ServiceApi.S3ServiceChangeCors(context.Background(), serviceId).S3ChangeCorsRequest(s3ChangeCorsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3ServiceApi.S3ServiceChangeCors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3ServiceChangeCors`: S3ChangeCorsResponse
    fmt.Fprintf(os.Stdout, "Response from `S3ServiceApi.S3ServiceChangeCors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3ServiceChangeCorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3ChangeCorsRequest** | [**S3ChangeCorsRequest**](S3ChangeCorsRequest.md) |  | 

### Return type

[**S3ChangeCorsResponse**](S3ChangeCorsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3ServiceChangeDomain

> S3ChangeDomainResponse S3ServiceChangeDomain(ctx, serviceId).S3ChangeDomainRequest(s3ChangeDomainRequest).Execute()



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
    s3ChangeDomainRequest := *openapiclient.NewS3ChangeDomainRequest() // S3ChangeDomainRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3ServiceApi.S3ServiceChangeDomain(context.Background(), serviceId).S3ChangeDomainRequest(s3ChangeDomainRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3ServiceApi.S3ServiceChangeDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3ServiceChangeDomain`: S3ChangeDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `S3ServiceApi.S3ServiceChangeDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3ServiceChangeDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3ChangeDomainRequest** | [**S3ChangeDomainRequest**](S3ChangeDomainRequest.md) |  | 

### Return type

[**S3ChangeDomainResponse**](S3ChangeDomainResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3ServiceChangePublic

> S3ChangePublicResponse S3ServiceChangePublic(ctx, serviceId).S3ChangePublicRequest(s3ChangePublicRequest).Execute()



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
    s3ChangePublicRequest := *openapiclient.NewS3ChangePublicRequest() // S3ChangePublicRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3ServiceApi.S3ServiceChangePublic(context.Background(), serviceId).S3ChangePublicRequest(s3ChangePublicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3ServiceApi.S3ServiceChangePublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3ServiceChangePublic`: S3ChangePublicResponse
    fmt.Fprintf(os.Stdout, "Response from `S3ServiceApi.S3ServiceChangePublic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3ServiceChangePublicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **s3ChangePublicRequest** | [**S3ChangePublicRequest**](S3ChangePublicRequest.md) |  | 

### Return type

[**S3ChangePublicResponse**](S3ChangePublicResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3ServiceGetPrefix

> S3GetPrefixResponse S3ServiceGetPrefix(ctx).Execute()



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
    resp, r, err := apiClient.S3ServiceApi.S3ServiceGetPrefix(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3ServiceApi.S3ServiceGetPrefix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3ServiceGetPrefix`: S3GetPrefixResponse
    fmt.Fprintf(os.Stdout, "Response from `S3ServiceApi.S3ServiceGetPrefix`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3ServiceGetPrefixRequest struct via the builder pattern


### Return type

[**S3GetPrefixResponse**](S3GetPrefixResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3ServiceGetPrice

> S3GetPriceResponse S3ServiceGetPrice(ctx).Execute()



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
    resp, r, err := apiClient.S3ServiceApi.S3ServiceGetPrice(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3ServiceApi.S3ServiceGetPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3ServiceGetPrice`: S3GetPriceResponse
    fmt.Fprintf(os.Stdout, "Response from `S3ServiceApi.S3ServiceGetPrice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3ServiceGetPriceRequest struct via the builder pattern


### Return type

[**S3GetPriceResponse**](S3GetPriceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3ServiceGetQuota

> S3GetQuotaResponse S3ServiceGetQuota(ctx).Execute()



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
    resp, r, err := apiClient.S3ServiceApi.S3ServiceGetQuota(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3ServiceApi.S3ServiceGetQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3ServiceGetQuota`: S3GetQuotaResponse
    fmt.Fprintf(os.Stdout, "Response from `S3ServiceApi.S3ServiceGetQuota`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiS3ServiceGetQuotaRequest struct via the builder pattern


### Return type

[**S3GetQuotaResponse**](S3GetQuotaResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

