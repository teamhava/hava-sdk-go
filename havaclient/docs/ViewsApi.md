# \ViewsApi

All URIs are relative to *https://api.hava.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ViewsExport**](ViewsApi.md#ViewsExport) | **Post** /views/{view_id}/export | Export a view



## ViewsExport

> ViewsExport(ctx, viewId).ViewIdExportBody(viewIdExportBody).Execute()

Export a view



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
    viewId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    viewIdExportBody := *openapiclient.NewViewIdExportBody() // ViewIdExportBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ViewsApi.ViewsExport(context.Background(), viewId).ViewIdExportBody(viewIdExportBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ViewsApi.ViewsExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**viewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiViewsExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **viewIdExportBody** | [**ViewIdExportBody**](ViewIdExportBody.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

