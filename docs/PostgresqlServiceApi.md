# \PostgresqlServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostgresqlServiceCreateDb**](PostgresqlServiceApi.md#PostgresqlServiceCreateDb) | **Post** /v1/cloud/postgresql/{service_id}/db | 
[**PostgresqlServiceCreateRole**](PostgresqlServiceApi.md#PostgresqlServiceCreateRole) | **Post** /v1/cloud/postgresql/{service_id}/role | 
[**PostgresqlServiceGetConfig**](PostgresqlServiceApi.md#PostgresqlServiceGetConfig) | **Get** /v1/cloud/postgresql/{service_id}/config | 
[**PostgresqlServiceGetDbList**](PostgresqlServiceApi.md#PostgresqlServiceGetDbList) | **Get** /v1/cloud/postgresql/{service_id}/db | 
[**PostgresqlServiceGetRemoteAccess**](PostgresqlServiceApi.md#PostgresqlServiceGetRemoteAccess) | **Get** /v1/cloud/postgresql/{service_id}/remote-access | 
[**PostgresqlServiceGetRoleList**](PostgresqlServiceApi.md#PostgresqlServiceGetRoleList) | **Get** /v1/cloud/postgresql/{service_id}/role | 
[**PostgresqlServiceRemoveDb**](PostgresqlServiceApi.md#PostgresqlServiceRemoveDb) | **Delete** /v1/cloud/postgresql/{service_id}/db/{db_name} | 
[**PostgresqlServiceRemoveRole**](PostgresqlServiceApi.md#PostgresqlServiceRemoveRole) | **Delete** /v1/cloud/postgresql/{service_id}/role/{role_name} | 
[**PostgresqlServiceSetConfig**](PostgresqlServiceApi.md#PostgresqlServiceSetConfig) | **Put** /v1/cloud/postgresql/{service_id}/config | 
[**PostgresqlServiceUpdateDb**](PostgresqlServiceApi.md#PostgresqlServiceUpdateDb) | **Patch** /v1/cloud/postgresql/{service_id}/db/{db_name} | 
[**PostgresqlServiceUpdateRemoteAccess**](PostgresqlServiceApi.md#PostgresqlServiceUpdateRemoteAccess) | **Put** /v1/cloud/postgresql/{service_id}/remote-access | 
[**PostgresqlServiceUpdateRole**](PostgresqlServiceApi.md#PostgresqlServiceUpdateRole) | **Patch** /v1/cloud/postgresql/{service_id}/role/{role_name} | 



## PostgresqlServiceCreateDb

> PostgresqlPgCreateDbResponse PostgresqlServiceCreateDb(ctx, serviceId).PostgresqlPgCreateDbRequest(postgresqlPgCreateDbRequest).Execute()



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
    postgresqlPgCreateDbRequest := *openapiclient.NewPostgresqlPgCreateDbRequest() // PostgresqlPgCreateDbRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceCreateDb(context.Background(), serviceId).PostgresqlPgCreateDbRequest(postgresqlPgCreateDbRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceCreateDb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceCreateDb`: PostgresqlPgCreateDbResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceCreateDb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceCreateDbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postgresqlPgCreateDbRequest** | [**PostgresqlPgCreateDbRequest**](PostgresqlPgCreateDbRequest.md) |  | 

### Return type

[**PostgresqlPgCreateDbResponse**](PostgresqlPgCreateDbResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceCreateRole

> PostgresqlPgCreateRoleResponse PostgresqlServiceCreateRole(ctx, serviceId).PostgresqlPgCreateRoleRequest(postgresqlPgCreateRoleRequest).Execute()



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
    postgresqlPgCreateRoleRequest := *openapiclient.NewPostgresqlPgCreateRoleRequest() // PostgresqlPgCreateRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceCreateRole(context.Background(), serviceId).PostgresqlPgCreateRoleRequest(postgresqlPgCreateRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceCreateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceCreateRole`: PostgresqlPgCreateRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceCreateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceCreateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postgresqlPgCreateRoleRequest** | [**PostgresqlPgCreateRoleRequest**](PostgresqlPgCreateRoleRequest.md) |  | 

### Return type

[**PostgresqlPgCreateRoleResponse**](PostgresqlPgCreateRoleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceGetConfig

> PostgresqlPgGetConfigResponse PostgresqlServiceGetConfig(ctx, serviceId).Execute()



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
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceGetConfig(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceGetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceGetConfig`: PostgresqlPgGetConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceGetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceGetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostgresqlPgGetConfigResponse**](PostgresqlPgGetConfigResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceGetDbList

> PostgresqlPgGetDbListResponse PostgresqlServiceGetDbList(ctx, serviceId).Execute()



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
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceGetDbList(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceGetDbList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceGetDbList`: PostgresqlPgGetDbListResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceGetDbList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceGetDbListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostgresqlPgGetDbListResponse**](PostgresqlPgGetDbListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceGetRemoteAccess

> PostgresqlPgGetRemoteAccessResponse PostgresqlServiceGetRemoteAccess(ctx, serviceId).Execute()



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
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceGetRemoteAccess(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceGetRemoteAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceGetRemoteAccess`: PostgresqlPgGetRemoteAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceGetRemoteAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceGetRemoteAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostgresqlPgGetRemoteAccessResponse**](PostgresqlPgGetRemoteAccessResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceGetRoleList

> PostgresqlPgGetRoleListResponse PostgresqlServiceGetRoleList(ctx, serviceId).Execute()



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
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceGetRoleList(context.Background(), serviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceGetRoleList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceGetRoleList`: PostgresqlPgGetRoleListResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceGetRoleList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceGetRoleListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PostgresqlPgGetRoleListResponse**](PostgresqlPgGetRoleListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceRemoveDb

> PostgresqlPgRemoveDbResponse PostgresqlServiceRemoveDb(ctx, serviceId, dbName).Execute()



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
    dbName := "dbName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceRemoveDb(context.Background(), serviceId, dbName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceRemoveDb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceRemoveDb`: PostgresqlPgRemoveDbResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceRemoveDb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 
**dbName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceRemoveDbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostgresqlPgRemoveDbResponse**](PostgresqlPgRemoveDbResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceRemoveRole

> PostgresqlPgRemoveRoleResponse PostgresqlServiceRemoveRole(ctx, serviceId, roleName).Execute()



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
    roleName := "roleName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceRemoveRole(context.Background(), serviceId, roleName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceRemoveRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceRemoveRole`: PostgresqlPgRemoveRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceRemoveRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceRemoveRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**PostgresqlPgRemoveRoleResponse**](PostgresqlPgRemoveRoleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceSetConfig

> PostgresqlPgSetConfigResponse PostgresqlServiceSetConfig(ctx, serviceId).PostgresqlPgSetConfigRequest(postgresqlPgSetConfigRequest).Execute()



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
    postgresqlPgSetConfigRequest := *openapiclient.NewPostgresqlPgSetConfigRequest() // PostgresqlPgSetConfigRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceSetConfig(context.Background(), serviceId).PostgresqlPgSetConfigRequest(postgresqlPgSetConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceSetConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceSetConfig`: PostgresqlPgSetConfigResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceSetConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceSetConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postgresqlPgSetConfigRequest** | [**PostgresqlPgSetConfigRequest**](PostgresqlPgSetConfigRequest.md) |  | 

### Return type

[**PostgresqlPgSetConfigResponse**](PostgresqlPgSetConfigResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceUpdateDb

> PostgresqlPgUpdateDbResponse PostgresqlServiceUpdateDb(ctx, serviceId, dbName).PostgresqlPgUpdateDbRequest(postgresqlPgUpdateDbRequest).Execute()



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
    dbName := "dbName_example" // string | 
    postgresqlPgUpdateDbRequest := *openapiclient.NewPostgresqlPgUpdateDbRequest() // PostgresqlPgUpdateDbRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceUpdateDb(context.Background(), serviceId, dbName).PostgresqlPgUpdateDbRequest(postgresqlPgUpdateDbRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceUpdateDb``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceUpdateDb`: PostgresqlPgUpdateDbResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceUpdateDb`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 
**dbName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceUpdateDbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postgresqlPgUpdateDbRequest** | [**PostgresqlPgUpdateDbRequest**](PostgresqlPgUpdateDbRequest.md) |  | 

### Return type

[**PostgresqlPgUpdateDbResponse**](PostgresqlPgUpdateDbResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceUpdateRemoteAccess

> PostgresqlPgUpdateRemoteAccessResponse PostgresqlServiceUpdateRemoteAccess(ctx, serviceId).PostgresqlPgUpdateRemoteAccessRequest(postgresqlPgUpdateRemoteAccessRequest).Execute()



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
    postgresqlPgUpdateRemoteAccessRequest := *openapiclient.NewPostgresqlPgUpdateRemoteAccessRequest() // PostgresqlPgUpdateRemoteAccessRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceUpdateRemoteAccess(context.Background(), serviceId).PostgresqlPgUpdateRemoteAccessRequest(postgresqlPgUpdateRemoteAccessRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceUpdateRemoteAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceUpdateRemoteAccess`: PostgresqlPgUpdateRemoteAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceUpdateRemoteAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceUpdateRemoteAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postgresqlPgUpdateRemoteAccessRequest** | [**PostgresqlPgUpdateRemoteAccessRequest**](PostgresqlPgUpdateRemoteAccessRequest.md) |  | 

### Return type

[**PostgresqlPgUpdateRemoteAccessResponse**](PostgresqlPgUpdateRemoteAccessResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlServiceUpdateRole

> PostgresqlPgUpdateRoleResponse PostgresqlServiceUpdateRole(ctx, serviceId, roleName).PostgresqlPgUpdateRoleRequest(postgresqlPgUpdateRoleRequest).Execute()



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
    roleName := "roleName_example" // string | 
    postgresqlPgUpdateRoleRequest := *openapiclient.NewPostgresqlPgUpdateRoleRequest() // PostgresqlPgUpdateRoleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlServiceApi.PostgresqlServiceUpdateRole(context.Background(), serviceId, roleName).PostgresqlPgUpdateRoleRequest(postgresqlPgUpdateRoleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlServiceApi.PostgresqlServiceUpdateRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlServiceUpdateRole`: PostgresqlPgUpdateRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlServiceApi.PostgresqlServiceUpdateRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serviceId** | **string** |  | 
**roleName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlServiceUpdateRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postgresqlPgUpdateRoleRequest** | [**PostgresqlPgUpdateRoleRequest**](PostgresqlPgUpdateRoleRequest.md) |  | 

### Return type

[**PostgresqlPgUpdateRoleResponse**](PostgresqlPgUpdateRoleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

