# \ProjectServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectServiceChangePinned**](ProjectServiceApi.md#ProjectServiceChangePinned) | **Put** /v1/cloud/projects/{project_id}/pin | 
[**ProjectServiceCreate**](ProjectServiceApi.md#ProjectServiceCreate) | **Post** /v1/cloud/projects | 
[**ProjectServiceGetList**](ProjectServiceApi.md#ProjectServiceGetList) | **Get** /v1/cloud/projects/list | 
[**ProjectServiceRemove**](ProjectServiceApi.md#ProjectServiceRemove) | **Delete** /v1/cloud/projects/{project_id} | 
[**ProjectServiceUpdate**](ProjectServiceApi.md#ProjectServiceUpdate) | **Put** /v1/cloud/projects/{project_id} | 



## ProjectServiceChangePinned

> ProjectChangePinnedResponse ProjectServiceChangePinned(ctx, projectId).ProjectChangePinnedRequest(projectChangePinnedRequest).Execute()



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
    projectId := "projectId_example" // string | 
    projectChangePinnedRequest := *openapiclient.NewProjectChangePinnedRequest() // ProjectChangePinnedRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceChangePinned(context.Background(), projectId).ProjectChangePinnedRequest(projectChangePinnedRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceChangePinned``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceChangePinned`: ProjectChangePinnedResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceChangePinned`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceChangePinnedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectChangePinnedRequest** | [**ProjectChangePinnedRequest**](ProjectChangePinnedRequest.md) |  | 

### Return type

[**ProjectChangePinnedResponse**](ProjectChangePinnedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceCreate

> ProjectCreateProjectResponse ProjectServiceCreate(ctx).ProjectCreateProjectRequest(projectCreateProjectRequest).Execute()



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
    projectCreateProjectRequest := *openapiclient.NewProjectCreateProjectRequest() // ProjectCreateProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceCreate(context.Background()).ProjectCreateProjectRequest(projectCreateProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceCreate`: ProjectCreateProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectCreateProjectRequest** | [**ProjectCreateProjectRequest**](ProjectCreateProjectRequest.md) |  | 

### Return type

[**ProjectCreateProjectResponse**](ProjectCreateProjectResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceGetList

> ProjectGetProjectListResponse ProjectServiceGetList(ctx).Offset(offset).Limit(limit).Filter(filter).Sort(sort).Execute()



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
    offset := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)
    filter := "filter_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceGetList(context.Background()).Offset(offset).Limit(limit).Filter(filter).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceGetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceGetList`: ProjectGetProjectListResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceGetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** |  | 
 **limit** | **int32** |  | 
 **filter** | **string** |  | 
 **sort** | **string** |  | 

### Return type

[**ProjectGetProjectListResponse**](ProjectGetProjectListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceRemove

> ProjectRemoveProjectResponse ProjectServiceRemove(ctx, projectId).Execute()



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
    projectId := "projectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceRemove(context.Background(), projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceRemove`: ProjectRemoveProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectRemoveProjectResponse**](ProjectRemoveProjectResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectServiceUpdate

> ProjectUpdateProjectResponse ProjectServiceUpdate(ctx, projectId).ProjectUpdateProjectRequest(projectUpdateProjectRequest).Execute()



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
    projectId := "projectId_example" // string | 
    projectUpdateProjectRequest := *openapiclient.NewProjectUpdateProjectRequest() // ProjectUpdateProjectRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectServiceApi.ProjectServiceUpdate(context.Background(), projectId).ProjectUpdateProjectRequest(projectUpdateProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectServiceApi.ProjectServiceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectServiceUpdate`: ProjectUpdateProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectServiceApi.ProjectServiceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectServiceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectUpdateProjectRequest** | [**ProjectUpdateProjectRequest**](ProjectUpdateProjectRequest.md) |  | 

### Return type

[**ProjectUpdateProjectResponse**](ProjectUpdateProjectResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

