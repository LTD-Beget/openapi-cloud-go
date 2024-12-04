# \MysqlBackupServiceApi

All URIs are relative to *https://api.beget.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MysqlBackupServiceGetList**](MysqlBackupServiceApi.md#MysqlBackupServiceGetList) | **Get** /v1/cloud/mysql/backup | 
[**MysqlBackupServiceGetOrders**](MysqlBackupServiceApi.md#MysqlBackupServiceGetOrders) | **Get** /v1/cloud/mysql/backup/orders | 
[**MysqlBackupServiceRestore**](MysqlBackupServiceApi.md#MysqlBackupServiceRestore) | **Post** /v1/cloud/mysql/backup/{copy_id} | 



## MysqlBackupServiceGetList

> MysqlBackupGetListResponse MysqlBackupServiceGetList(ctx).Execute()



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
    resp, r, err := apiClient.MysqlBackupServiceApi.MysqlBackupServiceGetList(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlBackupServiceApi.MysqlBackupServiceGetList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlBackupServiceGetList`: MysqlBackupGetListResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlBackupServiceApi.MysqlBackupServiceGetList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlBackupServiceGetListRequest struct via the builder pattern


### Return type

[**MysqlBackupGetListResponse**](MysqlBackupGetListResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlBackupServiceGetOrders

> MysqlBackupGetOrdersResponse MysqlBackupServiceGetOrders(ctx).Limit(limit).Offset(offset).Execute()



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
    limit := int32(56) // int32 |  (optional)
    offset := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlBackupServiceApi.MysqlBackupServiceGetOrders(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlBackupServiceApi.MysqlBackupServiceGetOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlBackupServiceGetOrders`: MysqlBackupGetOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlBackupServiceApi.MysqlBackupServiceGetOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMysqlBackupServiceGetOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 
 **offset** | **int32** |  | 

### Return type

[**MysqlBackupGetOrdersResponse**](MysqlBackupGetOrdersResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MysqlBackupServiceRestore

> MysqlBackupRestoreResponse MysqlBackupServiceRestore(ctx, copyId).MysqlBackupRestoreRequest(mysqlBackupRestoreRequest).Execute()



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
    copyId := "copyId_example" // string | 
    mysqlBackupRestoreRequest := *openapiclient.NewMysqlBackupRestoreRequest() // MysqlBackupRestoreRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MysqlBackupServiceApi.MysqlBackupServiceRestore(context.Background(), copyId).MysqlBackupRestoreRequest(mysqlBackupRestoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MysqlBackupServiceApi.MysqlBackupServiceRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MysqlBackupServiceRestore`: MysqlBackupRestoreResponse
    fmt.Fprintf(os.Stdout, "Response from `MysqlBackupServiceApi.MysqlBackupServiceRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**copyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMysqlBackupServiceRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mysqlBackupRestoreRequest** | [**MysqlBackupRestoreRequest**](MysqlBackupRestoreRequest.md) |  | 

### Return type

[**MysqlBackupRestoreResponse**](MysqlBackupRestoreResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

