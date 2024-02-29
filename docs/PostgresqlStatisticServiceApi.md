# \PostgresqlStatisticServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostgresqlStatisticServiceGetCpu**](PostgresqlStatisticServiceApi.md#PostgresqlStatisticServiceGetCpu) | **Get** /v1/cloud/postgresql/{service_id}/statistic/cpu | 
[**PostgresqlStatisticServiceGetCpuDetails**](PostgresqlStatisticServiceApi.md#PostgresqlStatisticServiceGetCpuDetails) | **Get** /v1/cloud/postgresql/{service_id}/statistic/cpu-details | 
[**PostgresqlStatisticServiceGetDisk**](PostgresqlStatisticServiceApi.md#PostgresqlStatisticServiceGetDisk) | **Get** /v1/cloud/postgresql/{service_id}/statistic/disk | 
[**PostgresqlStatisticServiceGetDiskUsage**](PostgresqlStatisticServiceApi.md#PostgresqlStatisticServiceGetDiskUsage) | **Get** /v1/cloud/postgresql/{service_id}/statistic/disk-usage | 
[**PostgresqlStatisticServiceGetLoadAverage**](PostgresqlStatisticServiceApi.md#PostgresqlStatisticServiceGetLoadAverage) | **Get** /v1/cloud/postgresql/{service_id}/statistic/load-average | 
[**PostgresqlStatisticServiceGetMemory**](PostgresqlStatisticServiceApi.md#PostgresqlStatisticServiceGetMemory) | **Get** /v1/cloud/postgresql/{service_id}/statistic/memory | 
[**PostgresqlStatisticServiceGetNetwork**](PostgresqlStatisticServiceApi.md#PostgresqlStatisticServiceGetNetwork) | **Get** /v1/cloud/postgresql/{service_id}/statistic/network | 



## PostgresqlStatisticServiceGetCpu

> PostgresqlStatisticGetCpuResponse PostgresqlStatisticServiceGetCpu(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetCpu(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetCpu``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlStatisticServiceGetCpu`: PostgresqlStatisticGetCpuResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetCpu`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlStatisticServiceGetCpuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**PostgresqlStatisticGetCpuResponse**](PostgresqlStatisticGetCpuResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlStatisticServiceGetCpuDetails

> PostgresqlStatisticGetCpuDetailsResponse PostgresqlStatisticServiceGetCpuDetails(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetCpuDetails(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetCpuDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlStatisticServiceGetCpuDetails`: PostgresqlStatisticGetCpuDetailsResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetCpuDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlStatisticServiceGetCpuDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**PostgresqlStatisticGetCpuDetailsResponse**](PostgresqlStatisticGetCpuDetailsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlStatisticServiceGetDisk

> PostgresqlStatisticGetDiskResponse PostgresqlStatisticServiceGetDisk(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetDisk(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlStatisticServiceGetDisk`: PostgresqlStatisticGetDiskResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlStatisticServiceGetDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**PostgresqlStatisticGetDiskResponse**](PostgresqlStatisticGetDiskResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlStatisticServiceGetDiskUsage

> PostgresqlStatisticGetDiskUsageResponse PostgresqlStatisticServiceGetDiskUsage(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetDiskUsage(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetDiskUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlStatisticServiceGetDiskUsage`: PostgresqlStatisticGetDiskUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetDiskUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlStatisticServiceGetDiskUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**PostgresqlStatisticGetDiskUsageResponse**](PostgresqlStatisticGetDiskUsageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlStatisticServiceGetLoadAverage

> PostgresqlStatisticGetLoadAverageResponse PostgresqlStatisticServiceGetLoadAverage(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetLoadAverage(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetLoadAverage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlStatisticServiceGetLoadAverage`: PostgresqlStatisticGetLoadAverageResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetLoadAverage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlStatisticServiceGetLoadAverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**PostgresqlStatisticGetLoadAverageResponse**](PostgresqlStatisticGetLoadAverageResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlStatisticServiceGetMemory

> PostgresqlStatisticGetMemoryResponse PostgresqlStatisticServiceGetMemory(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetMemory(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetMemory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlStatisticServiceGetMemory`: PostgresqlStatisticGetMemoryResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetMemory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlStatisticServiceGetMemoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**PostgresqlStatisticGetMemoryResponse**](PostgresqlStatisticGetMemoryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlStatisticServiceGetNetwork

> PostgresqlStatisticGetNetworkResponse PostgresqlStatisticServiceGetNetwork(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetNetwork(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlStatisticServiceGetNetwork`: PostgresqlStatisticGetNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlStatisticServiceApi.PostgresqlStatisticServiceGetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlStatisticServiceGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**PostgresqlStatisticGetNetworkResponse**](PostgresqlStatisticGetNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

