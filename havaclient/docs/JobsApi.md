# \JobsApi

All URIs are relative to *https://api.hava.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**JobsShow**](JobsApi.md#JobsShow) | **Get** /jobs/{job_id} | Get a job



## JobsShow

> JobsShow200Response JobsShow(ctx, jobId).Execute()

Get a job

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
    jobId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the job to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JobsApi.JobsShow(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JobsApi.JobsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `JobsShow`: JobsShow200Response
    fmt.Fprintf(os.Stdout, "Response from `JobsApi.JobsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The unique ID of the job to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**JobsShow200Response**](JobsShow200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

