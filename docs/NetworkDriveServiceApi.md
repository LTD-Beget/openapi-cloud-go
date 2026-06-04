# \NetworkDriveServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkDriveServiceGetLimits**](NetworkDriveServiceApi.md#NetworkDriveServiceGetLimits) | **Get** /v1/cloud/network-drive/limits | 
[**NetworkDriveServiceGetMounts**](NetworkDriveServiceApi.md#NetworkDriveServiceGetMounts) | **Get** /v1/cloud/network-drive/mounts | 
[**NetworkDriveServiceGetPrice**](NetworkDriveServiceApi.md#NetworkDriveServiceGetPrice) | **Get** /v1/cloud/network-drive/price | 
[**NetworkDriveServiceMount**](NetworkDriveServiceApi.md#NetworkDriveServiceMount) | **Post** /v1/cloud/network-drive/{service_id}/mount | 
[**NetworkDriveServiceResize**](NetworkDriveServiceApi.md#NetworkDriveServiceResize) | **Post** /v1/cloud/network-drive/{service_id}/resize | 
[**NetworkDriveServiceSetMounts**](NetworkDriveServiceApi.md#NetworkDriveServiceSetMounts) | **Post** /v1/cloud/network-drive/mounts | 
[**NetworkDriveServiceUnmount**](NetworkDriveServiceApi.md#NetworkDriveServiceUnmount) | **Post** /v1/cloud/network-drive/{service_id}/unmount | 



## NetworkDriveServiceGetLimits

> NetworkDriveGetLimitsResponse NetworkDriveServiceGetLimits(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDriveServiceApi.NetworkDriveServiceGetLimits(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDriveServiceApi.NetworkDriveServiceGetLimits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkDriveServiceGetLimits`: NetworkDriveGetLimitsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkDriveServiceApi.NetworkDriveServiceGetLimits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDriveServiceGetLimitsRequest struct via the builder pattern


### Return type

[**NetworkDriveGetLimitsResponse**](NetworkDriveGetLimitsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkDriveServiceGetMounts

> NetworkDriveGetMountsResponse NetworkDriveServiceGetMounts(ctx).ResourceId(resourceId).ResourceType(resourceType).Execute()



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
    resourceId := "resourceId_example" // string |  (optional)
    resourceType := "resourceType_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDriveServiceApi.NetworkDriveServiceGetMounts(context.Background()).ResourceId(resourceId).ResourceType(resourceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDriveServiceApi.NetworkDriveServiceGetMounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkDriveServiceGetMounts`: NetworkDriveGetMountsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkDriveServiceApi.NetworkDriveServiceGetMounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDriveServiceGetMountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceId** | **string** |  | 
 **resourceType** | **string** |  | 

### Return type

[**NetworkDriveGetMountsResponse**](NetworkDriveGetMountsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkDriveServiceGetPrice

> NetworkDriveGetPriceResponse NetworkDriveServiceGetPrice(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDriveServiceApi.NetworkDriveServiceGetPrice(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDriveServiceApi.NetworkDriveServiceGetPrice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkDriveServiceGetPrice`: NetworkDriveGetPriceResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkDriveServiceApi.NetworkDriveServiceGetPrice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDriveServiceGetPriceRequest struct via the builder pattern


### Return type

[**NetworkDriveGetPriceResponse**](NetworkDriveGetPriceResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkDriveServiceMount

> NetworkDriveMountResponse NetworkDriveServiceMount(ctx, serviceId).NetworkDriveMountRequest(networkDriveMountRequest).Execute()



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
    networkDriveMountRequest := *openapiclient.NewNetworkDriveMountRequest() // NetworkDriveMountRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDriveServiceApi.NetworkDriveServiceMount(context.Background(), serviceId).NetworkDriveMountRequest(networkDriveMountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDriveServiceApi.NetworkDriveServiceMount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkDriveServiceMount`: NetworkDriveMountResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkDriveServiceApi.NetworkDriveServiceMount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDriveServiceMountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkDriveMountRequest** | [**NetworkDriveMountRequest**](NetworkDriveMountRequest.md) |  | 

### Return type

[**NetworkDriveMountResponse**](NetworkDriveMountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkDriveServiceResize

> NetworkDriveResizeResponse NetworkDriveServiceResize(ctx, serviceId).NetworkDriveResizeRequest(networkDriveResizeRequest).Execute()



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
    networkDriveResizeRequest := *openapiclient.NewNetworkDriveResizeRequest() // NetworkDriveResizeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDriveServiceApi.NetworkDriveServiceResize(context.Background(), serviceId).NetworkDriveResizeRequest(networkDriveResizeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDriveServiceApi.NetworkDriveServiceResize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkDriveServiceResize`: NetworkDriveResizeResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkDriveServiceApi.NetworkDriveServiceResize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDriveServiceResizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkDriveResizeRequest** | [**NetworkDriveResizeRequest**](NetworkDriveResizeRequest.md) |  | 

### Return type

[**NetworkDriveResizeResponse**](NetworkDriveResizeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkDriveServiceSetMounts

> NetworkDriveSetMountsResponse NetworkDriveServiceSetMounts(ctx).NetworkDriveSetMountsRequest(networkDriveSetMountsRequest).Execute()



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
    networkDriveSetMountsRequest := *openapiclient.NewNetworkDriveSetMountsRequest() // NetworkDriveSetMountsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDriveServiceApi.NetworkDriveServiceSetMounts(context.Background()).NetworkDriveSetMountsRequest(networkDriveSetMountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDriveServiceApi.NetworkDriveServiceSetMounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkDriveServiceSetMounts`: NetworkDriveSetMountsResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkDriveServiceApi.NetworkDriveServiceSetMounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDriveServiceSetMountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **networkDriveSetMountsRequest** | [**NetworkDriveSetMountsRequest**](NetworkDriveSetMountsRequest.md) |  | 

### Return type

[**NetworkDriveSetMountsResponse**](NetworkDriveSetMountsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## NetworkDriveServiceUnmount

> NetworkDriveUnmountResponse NetworkDriveServiceUnmount(ctx, serviceId).NetworkDriveUnmountRequest(networkDriveUnmountRequest).Execute()



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
    networkDriveUnmountRequest := *openapiclient.NewNetworkDriveUnmountRequest() // NetworkDriveUnmountRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkDriveServiceApi.NetworkDriveServiceUnmount(context.Background(), serviceId).NetworkDriveUnmountRequest(networkDriveUnmountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkDriveServiceApi.NetworkDriveServiceUnmount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `NetworkDriveServiceUnmount`: NetworkDriveUnmountResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkDriveServiceApi.NetworkDriveServiceUnmount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworkDriveServiceUnmountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **networkDriveUnmountRequest** | [**NetworkDriveUnmountRequest**](NetworkDriveUnmountRequest.md) |  | 

### Return type

[**NetworkDriveUnmountResponse**](NetworkDriveUnmountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

