# \CdnStatisticServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CdnStatisticServiceGetNetwork**](CdnStatisticServiceApi.md#CdnStatisticServiceGetNetwork) | **Get** /v1/cloud/cdn/{service_id}/statistic/network | 
[**CdnStatisticServiceGetRequest**](CdnStatisticServiceApi.md#CdnStatisticServiceGetRequest) | **Get** /v1/cloud/cdn/{service_id}/statistic/count-request | 
[**CdnStatisticServiceGetTraffic**](CdnStatisticServiceApi.md#CdnStatisticServiceGetTraffic) | **Get** /v1/cloud/cdn/{service_id}/statistic/traffic-usage | 



## CdnStatisticServiceGetNetwork

> CdnStatisticGetNetworkResponse CdnStatisticServiceGetNetwork(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.CdnStatisticServiceApi.CdnStatisticServiceGetNetwork(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnStatisticServiceApi.CdnStatisticServiceGetNetwork``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnStatisticServiceGetNetwork`: CdnStatisticGetNetworkResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnStatisticServiceApi.CdnStatisticServiceGetNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCdnStatisticServiceGetNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**CdnStatisticGetNetworkResponse**](CdnStatisticGetNetworkResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CdnStatisticServiceGetRequest

> CdnStatisticGetRequestResponse CdnStatisticServiceGetRequest(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.CdnStatisticServiceApi.CdnStatisticServiceGetRequest(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnStatisticServiceApi.CdnStatisticServiceGetRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnStatisticServiceGetRequest`: CdnStatisticGetRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnStatisticServiceApi.CdnStatisticServiceGetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCdnStatisticServiceGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**CdnStatisticGetRequestResponse**](CdnStatisticGetRequestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CdnStatisticServiceGetTraffic

> CdnStatisticGetTrafficResponse CdnStatisticServiceGetTraffic(ctx, serviceId).Period(period).Execute()



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
    resp, r, err := apiClient.CdnStatisticServiceApi.CdnStatisticServiceGetTraffic(context.Background(), serviceId).Period(period).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CdnStatisticServiceApi.CdnStatisticServiceGetTraffic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CdnStatisticServiceGetTraffic`: CdnStatisticGetTrafficResponse
    fmt.Fprintf(os.Stdout, "Response from `CdnStatisticServiceApi.CdnStatisticServiceGetTraffic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCdnStatisticServiceGetTrafficRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **string** |  | 

### Return type

[**CdnStatisticGetTrafficResponse**](CdnStatisticGetTrafficResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

