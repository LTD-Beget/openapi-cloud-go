# \MysqlStatisticServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlStatisticServiceGetCpu**](MysqlStatisticServiceApi.md#MysqlStatisticServiceGetCpu) | **Get** /v1/cloud/mysql/{service_id}/statistic/cpu | 
[**MysqlStatisticServiceGetCpuDetails**](MysqlStatisticServiceApi.md#MysqlStatisticServiceGetCpuDetails) | **Get** /v1/cloud/mysql/{service_id}/statistic/cpu-details | 
[**MysqlStatisticServiceGetDisk**](MysqlStatisticServiceApi.md#MysqlStatisticServiceGetDisk) | **Get** /v1/cloud/mysql/{service_id}/statistic/disk | 
[**MysqlStatisticServiceGetDiskUsage**](MysqlStatisticServiceApi.md#MysqlStatisticServiceGetDiskUsage) | **Get** /v1/cloud/mysql/{service_id}/statistic/disk-usage | 
[**MysqlStatisticServiceGetLoadAverage**](MysqlStatisticServiceApi.md#MysqlStatisticServiceGetLoadAverage) | **Get** /v1/cloud/mysql/{service_id}/statistic/load-average | 
[**MysqlStatisticServiceGetMemory**](MysqlStatisticServiceApi.md#MysqlStatisticServiceGetMemory) | **Get** /v1/cloud/mysql/{service_id}/statistic/memory | 
[**MysqlStatisticServiceGetNetwork**](MysqlStatisticServiceApi.md#MysqlStatisticServiceGetNetwork) | **Get** /v1/cloud/mysql/{service_id}/statistic/network | 



## MysqlStatisticServiceGetCpu

> MysqlStatisticGetCpuResponse MysqlStatisticServiceGetCpu(ctx, serviceId).Period(period).Execute()



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
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetCpu(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlStatisticServiceApi.MysqlStatisticServiceGetCpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlStatisticServiceGetCpu`: MysqlStatisticGetCpuResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlStatisticServiceApi.MysqlStatisticServiceGetCpu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlStatisticServiceGetCpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**MysqlStatisticGetCpuResponse**](MysqlStatisticGetCpuResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlStatisticServiceGetCpuDetails

> MysqlStatisticGetCpuDetailsResponse MysqlStatisticServiceGetCpuDetails(ctx, serviceId).Period(period).Execute()



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
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetCpuDetails(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlStatisticServiceApi.MysqlStatisticServiceGetCpuDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlStatisticServiceGetCpuDetails`: MysqlStatisticGetCpuDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlStatisticServiceApi.MysqlStatisticServiceGetCpuDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlStatisticServiceGetCpuDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**MysqlStatisticGetCpuDetailsResponse**](MysqlStatisticGetCpuDetailsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlStatisticServiceGetDisk

> MysqlStatisticGetDiskResponse MysqlStatisticServiceGetDisk(ctx, serviceId).Period(period).Execute()



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
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetDisk(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlStatisticServiceApi.MysqlStatisticServiceGetDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlStatisticServiceGetDisk`: MysqlStatisticGetDiskResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlStatisticServiceApi.MysqlStatisticServiceGetDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlStatisticServiceGetDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**MysqlStatisticGetDiskResponse**](MysqlStatisticGetDiskResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlStatisticServiceGetDiskUsage

> MysqlStatisticGetDiskUsageResponse MysqlStatisticServiceGetDiskUsage(ctx, serviceId).Period(period).Execute()



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
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetDiskUsage(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlStatisticServiceApi.MysqlStatisticServiceGetDiskUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlStatisticServiceGetDiskUsage`: MysqlStatisticGetDiskUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlStatisticServiceApi.MysqlStatisticServiceGetDiskUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlStatisticServiceGetDiskUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**MysqlStatisticGetDiskUsageResponse**](MysqlStatisticGetDiskUsageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlStatisticServiceGetLoadAverage

> MysqlStatisticGetLoadAverageResponse MysqlStatisticServiceGetLoadAverage(ctx, serviceId).Period(period).Execute()



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
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetLoadAverage(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlStatisticServiceApi.MysqlStatisticServiceGetLoadAverage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlStatisticServiceGetLoadAverage`: MysqlStatisticGetLoadAverageResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlStatisticServiceApi.MysqlStatisticServiceGetLoadAverage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlStatisticServiceGetLoadAverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**MysqlStatisticGetLoadAverageResponse**](MysqlStatisticGetLoadAverageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlStatisticServiceGetMemory

> MysqlStatisticGetMemoryResponse MysqlStatisticServiceGetMemory(ctx, serviceId).Period(period).Execute()



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
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetMemory(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlStatisticServiceApi.MysqlStatisticServiceGetMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlStatisticServiceGetMemory`: MysqlStatisticGetMemoryResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlStatisticServiceApi.MysqlStatisticServiceGetMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlStatisticServiceGetMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**MysqlStatisticGetMemoryResponse**](MysqlStatisticGetMemoryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlStatisticServiceGetNetwork

> MysqlStatisticGetNetworkResponse MysqlStatisticServiceGetNetwork(ctx, serviceId).Period(period).Execute()



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
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlStatisticServiceApi.MysqlStatisticServiceGetNetwork(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlStatisticServiceApi.MysqlStatisticServiceGetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlStatisticServiceGetNetwork`: MysqlStatisticGetNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlStatisticServiceApi.MysqlStatisticServiceGetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlStatisticServiceGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**MysqlStatisticGetNetworkResponse**](MysqlStatisticGetNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

