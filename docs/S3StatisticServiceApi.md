# \S3StatisticServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**S3StatisticServiceGetNetwork**](S3StatisticServiceApi.md#S3StatisticServiceGetNetwork) | **Get** /v1/cloud/s3/{service_id}/statistic/network | 
[**S3StatisticServiceGetQuota**](S3StatisticServiceApi.md#S3StatisticServiceGetQuota) | **Get** /v1/cloud/s3/{service_id}/statistic/quota | 
[**S3StatisticServiceGetRequest**](S3StatisticServiceApi.md#S3StatisticServiceGetRequest) | **Get** /v1/cloud/s3/{service_id}/statistic/count-request | 
[**S3StatisticServiceGetTraffic**](S3StatisticServiceApi.md#S3StatisticServiceGetTraffic) | **Get** /v1/cloud/s3/{service_id}/statistic/traffic-usage | 



## S3StatisticServiceGetNetwork

> S3StatisticGetNetworkResponse S3StatisticServiceGetNetwork(ctx, serviceId).Period(period).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/LTD-Beget/openapi-cloud-go"
)

func main() {
    serviceId := "serviceId_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3StatisticServiceApi.S3StatisticServiceGetNetwork(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3StatisticServiceApi.S3StatisticServiceGetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3StatisticServiceGetNetwork`: S3StatisticGetNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `S3StatisticServiceApi.S3StatisticServiceGetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3StatisticServiceGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**S3StatisticGetNetworkResponse**](S3StatisticGetNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3StatisticServiceGetQuota

> S3StatisticGetQuotaResponse S3StatisticServiceGetQuota(ctx, serviceId).Period(period).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/LTD-Beget/openapi-cloud-go"
)

func main() {
    serviceId := "serviceId_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3StatisticServiceApi.S3StatisticServiceGetQuota(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3StatisticServiceApi.S3StatisticServiceGetQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3StatisticServiceGetQuota`: S3StatisticGetQuotaResponse
    fmt.Fprintf(os.Stdout, "Response from `S3StatisticServiceApi.S3StatisticServiceGetQuota`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3StatisticServiceGetQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**S3StatisticGetQuotaResponse**](S3StatisticGetQuotaResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3StatisticServiceGetRequest

> S3StatisticGetRequestResponse S3StatisticServiceGetRequest(ctx, serviceId).Period(period).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/LTD-Beget/openapi-cloud-go"
)

func main() {
    serviceId := "serviceId_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3StatisticServiceApi.S3StatisticServiceGetRequest(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3StatisticServiceApi.S3StatisticServiceGetRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3StatisticServiceGetRequest`: S3StatisticGetRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `S3StatisticServiceApi.S3StatisticServiceGetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3StatisticServiceGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**S3StatisticGetRequestResponse**](S3StatisticGetRequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## S3StatisticServiceGetTraffic

> S3StatisticGetTrafficResponse S3StatisticServiceGetTraffic(ctx, serviceId).Period(period).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/LTD-Beget/openapi-cloud-go"
)

func main() {
    serviceId := "serviceId_example" // string | 
    period := "period_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.S3StatisticServiceApi.S3StatisticServiceGetTraffic(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `S3StatisticServiceApi.S3StatisticServiceGetTraffic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `S3StatisticServiceGetTraffic`: S3StatisticGetTrafficResponse
    fmt.Fprintf(os.Stdout, "Response from `S3StatisticServiceApi.S3StatisticServiceGetTraffic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiS3StatisticServiceGetTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**S3StatisticGetTrafficResponse**](S3StatisticGetTrafficResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

