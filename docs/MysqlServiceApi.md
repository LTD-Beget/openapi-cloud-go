# \MysqlServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlServiceChangeAccessPassword**](MysqlServiceApi.md#MysqlServiceChangeAccessPassword) | **Patch** /v1/cloud/mysql/{service_id}/db/{db_name}/access/{host} | 
[**MysqlServiceCreateAccess**](MysqlServiceApi.md#MysqlServiceCreateAccess) | **Post** /v1/cloud/mysql/{service_id}/db/{db_name}/access | 
[**MysqlServiceCreateDb**](MysqlServiceApi.md#MysqlServiceCreateDb) | **Post** /v1/cloud/mysql/{service_id}/db | 
[**MysqlServiceGetConfig**](MysqlServiceApi.md#MysqlServiceGetConfig) | **Get** /v1/cloud/mysql/{service_id}/config | 
[**MysqlServiceGetDbList**](MysqlServiceApi.md#MysqlServiceGetDbList) | **Get** /v1/cloud/mysql/{service_id}/db | 
[**MysqlServiceRemoveAccess**](MysqlServiceApi.md#MysqlServiceRemoveAccess) | **Delete** /v1/cloud/mysql/{service_id}/db/{db_name}/access/{host} | 
[**MysqlServiceRemoveDb**](MysqlServiceApi.md#MysqlServiceRemoveDb) | **Delete** /v1/cloud/mysql/{service_id}/db/{db_name} | 
[**MysqlServiceSetConfig**](MysqlServiceApi.md#MysqlServiceSetConfig) | **Put** /v1/cloud/mysql/{service_id}/config | 
[**MysqlServiceUpdateDb**](MysqlServiceApi.md#MysqlServiceUpdateDb) | **Patch** /v1/cloud/mysql/{service_id}/db/{db_name} | 



## MysqlServiceChangeAccessPassword

> MysqlChangeAccessPasswordResponse MysqlServiceChangeAccessPassword(ctx, serviceId, dbName, host).MysqlChangeAccessPasswordRequest(mysqlChangeAccessPasswordRequest).Execute()



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
    dbName := "dbName_example" // string | 
    host := "host_example" // string | 
    mysqlChangeAccessPasswordRequest := *openapiclient.NewMysqlChangeAccessPasswordRequest() // MysqlChangeAccessPasswordRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlServiceApi.MysqlServiceChangeAccessPassword(context.Background(), serviceId, dbName, host).MysqlChangeAccessPasswordRequest(mysqlChangeAccessPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlServiceApi.MysqlServiceChangeAccessPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlServiceChangeAccessPassword`: MysqlChangeAccessPasswordResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlServiceApi.MysqlServiceChangeAccessPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 
**dbName** | **string** |  | 
**host** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlServiceChangeAccessPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **mysqlChangeAccessPasswordRequest** | [**MysqlChangeAccessPasswordRequest**](MysqlChangeAccessPasswordRequest.md) |  | 

### Return type

[**MysqlChangeAccessPasswordResponse**](MysqlChangeAccessPasswordResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlServiceCreateAccess

> MysqlCreateAccessResponse MysqlServiceCreateAccess(ctx, serviceId, dbName).MysqlCreateAccessRequest(mysqlCreateAccessRequest).Execute()



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
    dbName := "dbName_example" // string | 
    mysqlCreateAccessRequest := *openapiclient.NewMysqlCreateAccessRequest() // MysqlCreateAccessRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlServiceApi.MysqlServiceCreateAccess(context.Background(), serviceId, dbName).MysqlCreateAccessRequest(mysqlCreateAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlServiceApi.MysqlServiceCreateAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlServiceCreateAccess`: MysqlCreateAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlServiceApi.MysqlServiceCreateAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 
**dbName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlServiceCreateAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mysqlCreateAccessRequest** | [**MysqlCreateAccessRequest**](MysqlCreateAccessRequest.md) |  | 

### Return type

[**MysqlCreateAccessResponse**](MysqlCreateAccessResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlServiceCreateDb

> MysqlCreateDbResponse MysqlServiceCreateDb(ctx, serviceId).MysqlCreateDbRequest(mysqlCreateDbRequest).Execute()



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
    mysqlCreateDbRequest := *openapiclient.NewMysqlCreateDbRequest() // MysqlCreateDbRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlServiceApi.MysqlServiceCreateDb(context.Background(), serviceId).MysqlCreateDbRequest(mysqlCreateDbRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlServiceApi.MysqlServiceCreateDb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlServiceCreateDb`: MysqlCreateDbResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlServiceApi.MysqlServiceCreateDb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlServiceCreateDbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mysqlCreateDbRequest** | [**MysqlCreateDbRequest**](MysqlCreateDbRequest.md) |  | 

### Return type

[**MysqlCreateDbResponse**](MysqlCreateDbResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlServiceGetConfig

> MysqlGetConfigResponse MysqlServiceGetConfig(ctx, serviceId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlServiceApi.MysqlServiceGetConfig(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlServiceApi.MysqlServiceGetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlServiceGetConfig`: MysqlGetConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlServiceApi.MysqlServiceGetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlServiceGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MysqlGetConfigResponse**](MysqlGetConfigResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlServiceGetDbList

> MysqlGetDbListResponse MysqlServiceGetDbList(ctx, serviceId).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlServiceApi.MysqlServiceGetDbList(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlServiceApi.MysqlServiceGetDbList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlServiceGetDbList`: MysqlGetDbListResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlServiceApi.MysqlServiceGetDbList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlServiceGetDbListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MysqlGetDbListResponse**](MysqlGetDbListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlServiceRemoveAccess

> MysqlRemoveAccessResponse MysqlServiceRemoveAccess(ctx, serviceId, dbName, host).Execute()



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
    dbName := "dbName_example" // string | 
    host := "host_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlServiceApi.MysqlServiceRemoveAccess(context.Background(), serviceId, dbName, host).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlServiceApi.MysqlServiceRemoveAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlServiceRemoveAccess`: MysqlRemoveAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlServiceApi.MysqlServiceRemoveAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 
**dbName** | **string** |  | 
**host** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlServiceRemoveAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**MysqlRemoveAccessResponse**](MysqlRemoveAccessResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlServiceRemoveDb

> MysqlRemoveDbResponse MysqlServiceRemoveDb(ctx, serviceId, dbName).Execute()



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
    dbName := "dbName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlServiceApi.MysqlServiceRemoveDb(context.Background(), serviceId, dbName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlServiceApi.MysqlServiceRemoveDb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlServiceRemoveDb`: MysqlRemoveDbResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlServiceApi.MysqlServiceRemoveDb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 
**dbName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlServiceRemoveDbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MysqlRemoveDbResponse**](MysqlRemoveDbResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlServiceSetConfig

> MysqlSetConfigResponse MysqlServiceSetConfig(ctx, serviceId).MysqlSetConfigRequest(mysqlSetConfigRequest).Execute()



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
    mysqlSetConfigRequest := *openapiclient.NewMysqlSetConfigRequest() // MysqlSetConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlServiceApi.MysqlServiceSetConfig(context.Background(), serviceId).MysqlSetConfigRequest(mysqlSetConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlServiceApi.MysqlServiceSetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlServiceSetConfig`: MysqlSetConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlServiceApi.MysqlServiceSetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlServiceSetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mysqlSetConfigRequest** | [**MysqlSetConfigRequest**](MysqlSetConfigRequest.md) |  | 

### Return type

[**MysqlSetConfigResponse**](MysqlSetConfigResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlServiceUpdateDb

> MysqlUpdateDbResponse MysqlServiceUpdateDb(ctx, serviceId, dbName).MysqlUpdateDbRequest(mysqlUpdateDbRequest).Execute()



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
    dbName := "dbName_example" // string | 
    mysqlUpdateDbRequest := *openapiclient.NewMysqlUpdateDbRequest() // MysqlUpdateDbRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlServiceApi.MysqlServiceUpdateDb(context.Background(), serviceId, dbName).MysqlUpdateDbRequest(mysqlUpdateDbRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlServiceApi.MysqlServiceUpdateDb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlServiceUpdateDb`: MysqlUpdateDbResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlServiceApi.MysqlServiceUpdateDb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 
**dbName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlServiceUpdateDbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **mysqlUpdateDbRequest** | [**MysqlUpdateDbRequest**](MysqlUpdateDbRequest.md) |  | 

### Return type

[**MysqlUpdateDbResponse**](MysqlUpdateDbResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

