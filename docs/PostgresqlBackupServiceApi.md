# \PostgresqlBackupServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostgresqlBackupServiceGetList**](PostgresqlBackupServiceApi.md#PostgresqlBackupServiceGetList) | **Get** /v1/cloud/postgresql/backup | 
[**PostgresqlBackupServiceGetOrders**](PostgresqlBackupServiceApi.md#PostgresqlBackupServiceGetOrders) | **Get** /v1/cloud/postgresql/backup/orders | 
[**PostgresqlBackupServiceRestore**](PostgresqlBackupServiceApi.md#PostgresqlBackupServiceRestore) | **Post** /v1/cloud/postgresql/backup/{copy_id} | 



## PostgresqlBackupServiceGetList

> PostgresqlBackupGetListResponse PostgresqlBackupServiceGetList(ctx).Filter(filter).Execute()



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
    filter := "filter_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlBackupServiceApi.PostgresqlBackupServiceGetList(context.Background()).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlBackupServiceApi.PostgresqlBackupServiceGetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlBackupServiceGetList`: PostgresqlBackupGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlBackupServiceApi.PostgresqlBackupServiceGetList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlBackupServiceGetListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** |  | 

### Return type

[**PostgresqlBackupGetListResponse**](PostgresqlBackupGetListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlBackupServiceGetOrders

> PostgresqlBackupGetOrdersResponse PostgresqlBackupServiceGetOrders(ctx).Limit(limit).Offset(offset).Execute()



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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlBackupServiceApi.PostgresqlBackupServiceGetOrders(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlBackupServiceApi.PostgresqlBackupServiceGetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlBackupServiceGetOrders`: PostgresqlBackupGetOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlBackupServiceApi.PostgresqlBackupServiceGetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlBackupServiceGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**PostgresqlBackupGetOrdersResponse**](PostgresqlBackupGetOrdersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostgresqlBackupServiceRestore

> PostgresqlBackupRestoreResponse PostgresqlBackupServiceRestore(ctx, copyId).PostgresqlBackupRestoreRequest(postgresqlBackupRestoreRequest).Execute()



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
    copyId := "copyId_example" // string | 
    postgresqlBackupRestoreRequest := *openapiclient.NewPostgresqlBackupRestoreRequest() // PostgresqlBackupRestoreRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PostgresqlBackupServiceApi.PostgresqlBackupServiceRestore(context.Background(), copyId).PostgresqlBackupRestoreRequest(postgresqlBackupRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PostgresqlBackupServiceApi.PostgresqlBackupServiceRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostgresqlBackupServiceRestore`: PostgresqlBackupRestoreResponse
    fmt.Fprintf(os.Stdout, "Response from `PostgresqlBackupServiceApi.PostgresqlBackupServiceRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**copyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostgresqlBackupServiceRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postgresqlBackupRestoreRequest** | [**PostgresqlBackupRestoreRequest**](PostgresqlBackupRestoreRequest.md) |  | 

### Return type

[**PostgresqlBackupRestoreResponse**](PostgresqlBackupRestoreResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

