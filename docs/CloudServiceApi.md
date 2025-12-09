# \CloudServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CloudServiceBindProject**](CloudServiceApi.md#CloudServiceBindProject) | **Put** /v1/cloud/{service_id}/project | 
[**CloudServiceChangeConfiguration**](CloudServiceApi.md#CloudServiceChangeConfiguration) | **Patch** /v1/cloud/{service_id}/configuration | 
[**CloudServiceChangePinned**](CloudServiceApi.md#CloudServiceChangePinned) | **Put** /v1/cloud/{service_id}/pin | 
[**CloudServiceCreate**](CloudServiceApi.md#CloudServiceCreate) | **Post** /v1/cloud | 
[**CloudServiceGet**](CloudServiceApi.md#CloudServiceGet) | **Get** /v1/cloud/{service_id} | 
[**CloudServiceGetConfigurationList**](CloudServiceApi.md#CloudServiceGetConfigurationList) | **Get** /v1/cloud/configuration | 
[**CloudServiceGetList**](CloudServiceApi.md#CloudServiceGetList) | **Get** /v1/cloud | 
[**CloudServiceRemove**](CloudServiceApi.md#CloudServiceRemove) | **Delete** /v1/cloud/{service_id} | 
[**CloudServiceUpdate**](CloudServiceApi.md#CloudServiceUpdate) | **Patch** /v1/cloud/{service_id} | 



## CloudServiceBindProject

> CloudBindProjectResponse CloudServiceBindProject(ctx, serviceId).CloudBindProjectRequest(cloudBindProjectRequest).Execute()



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
    cloudBindProjectRequest := *openapiclient.NewCloudBindProjectRequest() // CloudBindProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudServiceApi.CloudServiceBindProject(context.Background(), serviceId).CloudBindProjectRequest(cloudBindProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudServiceApi.CloudServiceBindProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudServiceBindProject`: CloudBindProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudServiceApi.CloudServiceBindProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServiceBindProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudBindProjectRequest** | [**CloudBindProjectRequest**](CloudBindProjectRequest.md) |  | 

### Return type

[**CloudBindProjectResponse**](CloudBindProjectResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServiceChangeConfiguration

> CloudChangeConfigurationResponse CloudServiceChangeConfiguration(ctx, serviceId).CloudChangeConfigurationRequest(cloudChangeConfigurationRequest).Execute()



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
    cloudChangeConfigurationRequest := *openapiclient.NewCloudChangeConfigurationRequest() // CloudChangeConfigurationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudServiceApi.CloudServiceChangeConfiguration(context.Background(), serviceId).CloudChangeConfigurationRequest(cloudChangeConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudServiceApi.CloudServiceChangeConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudServiceChangeConfiguration`: CloudChangeConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudServiceApi.CloudServiceChangeConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServiceChangeConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudChangeConfigurationRequest** | [**CloudChangeConfigurationRequest**](CloudChangeConfigurationRequest.md) |  | 

### Return type

[**CloudChangeConfigurationResponse**](CloudChangeConfigurationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServiceChangePinned

> CloudChangePinnedResponse CloudServiceChangePinned(ctx, serviceId).CloudChangePinnedRequest(cloudChangePinnedRequest).Execute()



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
    cloudChangePinnedRequest := *openapiclient.NewCloudChangePinnedRequest() // CloudChangePinnedRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudServiceApi.CloudServiceChangePinned(context.Background(), serviceId).CloudChangePinnedRequest(cloudChangePinnedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudServiceApi.CloudServiceChangePinned``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudServiceChangePinned`: CloudChangePinnedResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudServiceApi.CloudServiceChangePinned`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServiceChangePinnedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudChangePinnedRequest** | [**CloudChangePinnedRequest**](CloudChangePinnedRequest.md) |  | 

### Return type

[**CloudChangePinnedResponse**](CloudChangePinnedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServiceCreate

> CloudCreateResponse CloudServiceCreate(ctx).CloudCreateRequest(cloudCreateRequest).Execute()



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
    cloudCreateRequest := *openapiclient.NewCloudCreateRequest() // CloudCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudServiceApi.CloudServiceCreate(context.Background()).CloudCreateRequest(cloudCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudServiceApi.CloudServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudServiceCreate`: CloudCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudServiceApi.CloudServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cloudCreateRequest** | [**CloudCreateRequest**](CloudCreateRequest.md) |  | 

### Return type

[**CloudCreateResponse**](CloudCreateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServiceGet

> CloudGetResponse CloudServiceGet(ctx, serviceId).Execute()



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
    resp, r, err := apiClient.CloudServiceApi.CloudServiceGet(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudServiceApi.CloudServiceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudServiceGet`: CloudGetResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudServiceApi.CloudServiceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServiceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudGetResponse**](CloudGetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServiceGetConfigurationList

> CloudGetConfigurationListResponse CloudServiceGetConfigurationList(ctx).Execute()



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
    resp, r, err := apiClient.CloudServiceApi.CloudServiceGetConfigurationList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudServiceApi.CloudServiceGetConfigurationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudServiceGetConfigurationList`: CloudGetConfigurationListResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudServiceApi.CloudServiceGetConfigurationList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServiceGetConfigurationListRequest struct via the builder pattern


### Return type

[**CloudGetConfigurationListResponse**](CloudGetConfigurationListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServiceGetList

> CloudGetListResponse CloudServiceGetList(ctx).Offset(offset).Limit(limit).Sort(sort).Filter(filter).View(view).Execute()



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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    sort := "sort_example" // string |  (optional)
    filter := "filter_example" // string |  (optional)
    view := "view_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudServiceApi.CloudServiceGetList(context.Background()).Offset(offset).Limit(limit).Sort(sort).Filter(filter).View(view).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudServiceApi.CloudServiceGetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudServiceGetList`: CloudGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudServiceApi.CloudServiceGetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCloudServiceGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **sort** | **string** |  | 
 **filter** | **string** |  | 
 **view** | **string** |  | 

### Return type

[**CloudGetListResponse**](CloudGetListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServiceRemove

> CloudRemoveResponse CloudServiceRemove(ctx, serviceId).Execute()



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
    resp, r, err := apiClient.CloudServiceApi.CloudServiceRemove(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudServiceApi.CloudServiceRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudServiceRemove`: CloudRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudServiceApi.CloudServiceRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServiceRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CloudRemoveResponse**](CloudRemoveResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CloudServiceUpdate

> CloudUpdateResponse CloudServiceUpdate(ctx, serviceId).CloudUpdateRequest(cloudUpdateRequest).Execute()



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
    cloudUpdateRequest := *openapiclient.NewCloudUpdateRequest() // CloudUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CloudServiceApi.CloudServiceUpdate(context.Background(), serviceId).CloudUpdateRequest(cloudUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CloudServiceApi.CloudServiceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CloudServiceUpdate`: CloudUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `CloudServiceApi.CloudServiceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCloudServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cloudUpdateRequest** | [**CloudUpdateRequest**](CloudUpdateRequest.md) |  | 

### Return type

[**CloudUpdateResponse**](CloudUpdateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

