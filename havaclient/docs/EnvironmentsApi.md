# \EnvironmentsApi

All URIs are relative to *https://api.hava.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EnvironmentRevisionsIndex**](EnvironmentsApi.md#EnvironmentRevisionsIndex) | **Get** /environments/{environment_id}/revisions | List versions of an environment
[**EnvironmentSharesCreate**](EnvironmentsApi.md#EnvironmentSharesCreate) | **Post** /environments/{environment_id}/shares | Generates an embedded environment token. If the provided view type exists it will return the share details, otherwise it will return a job that will generate the share details.
[**EnvironmentSharesFind**](EnvironmentsApi.md#EnvironmentSharesFind) | **Get** /environments/{environment_id}/shares/{share_id} | Return the details of the specified environment share
[**EnvironmentSharesList**](EnvironmentsApi.md#EnvironmentSharesList) | **Get** /environments/{environment_id}/shares | List shares of the specific environment
[**EnvironmentSharesUpdate**](EnvironmentsApi.md#EnvironmentSharesUpdate) | **Put** /environments/{environment_id}/shares/{share_id} | Update the details of the specified environment share. Used to change the config for PNG shares.
[**EnvironmentsCreate**](EnvironmentsApi.md#EnvironmentsCreate) | **Post** /environments | Create an environment using a query
[**EnvironmentsDestroy**](EnvironmentsApi.md#EnvironmentsDestroy) | **Delete** /environments/{environment_id} | Delete a custom environment
[**EnvironmentsIndex**](EnvironmentsApi.md#EnvironmentsIndex) | **Get** /environments | List environments
[**EnvironmentsShow**](EnvironmentsApi.md#EnvironmentsShow) | **Get** /environments/{environment_id} | Get an environment
[**EnvironmentsUpdate**](EnvironmentsApi.md#EnvironmentsUpdate) | **Put** /environments/{environment_id} | Update a custom environment



## EnvironmentRevisionsIndex

> EnvironmentRevisionsIndex200Response EnvironmentRevisionsIndex(ctx, environmentId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()

List versions of an environment

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    page := int32(56) // int32 | The page of results to display. Based on page size and total results. (optional)
    pageSize := int32(56) // int32 | The number of results to return per page (optional)
    sort := "sort_example" // string | The field to sort by (optional)
    sortDir := "sortDir_example" // string | The direction to sort by (optional)
    token := "token_example" // string | The next page token returned from a previous pagination request. If this is provided no other fields are required (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentRevisionsIndex(context.Background(), environmentId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentRevisionsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentRevisionsIndex`: EnvironmentRevisionsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentRevisionsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentRevisionsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of results to display. Based on page size and total results. | 
 **pageSize** | **int32** | The number of results to return per page | 
 **sort** | **string** | The field to sort by | 
 **sortDir** | **string** | The direction to sort by | 
 **token** | **string** | The next page token returned from a previous pagination request. If this is provided no other fields are required | 

### Return type

[**EnvironmentRevisionsIndex200Response**](EnvironmentRevisionsIndex200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentSharesCreate

> EnvironmentShare EnvironmentSharesCreate(ctx, environmentId).EnvironmentSharesCreateRequest(environmentSharesCreateRequest).Execute()

Generates an embedded environment token. If the provided view type exists it will return the share details, otherwise it will return a job that will generate the share details.

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the environment
    environmentSharesCreateRequest := *openapiclient.NewEnvironmentSharesCreateRequest() // EnvironmentSharesCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentSharesCreate(context.Background(), environmentId).EnvironmentSharesCreateRequest(environmentSharesCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentSharesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentSharesCreate`: EnvironmentShare
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentSharesCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique ID of the environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentSharesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentSharesCreateRequest** | [**EnvironmentSharesCreateRequest**](EnvironmentSharesCreateRequest.md) |  | 

### Return type

[**EnvironmentShare**](EnvironmentShare.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentSharesFind

> EnvironmentShare EnvironmentSharesFind(ctx, environmentId, shareId).Execute()

Return the details of the specified environment share

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the environment
    shareId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the environment share

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentSharesFind(context.Background(), environmentId, shareId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentSharesFind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentSharesFind`: EnvironmentShare
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentSharesFind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique ID of the environment | 
**shareId** | **string** | The unique ID of the environment share | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentSharesFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnvironmentShare**](EnvironmentShare.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentSharesList

> EnvironmentSharesList200Response EnvironmentSharesList(ctx, environmentId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()

List shares of the specific environment

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the environment
    page := int32(56) // int32 | The page of results to display. Based on page size and total results. (optional)
    pageSize := int32(56) // int32 | The number of results to return per page (optional)
    sort := "sort_example" // string | The field to sort by (optional)
    sortDir := "sortDir_example" // string | The direction to sort by (optional)
    token := "token_example" // string | The next page token returned from a previous pagination request. If this is provided no other fields are required (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentSharesList(context.Background(), environmentId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentSharesList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentSharesList`: EnvironmentSharesList200Response
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentSharesList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique ID of the environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentSharesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of results to display. Based on page size and total results. | 
 **pageSize** | **int32** | The number of results to return per page | 
 **sort** | **string** | The field to sort by | 
 **sortDir** | **string** | The direction to sort by | 
 **token** | **string** | The next page token returned from a previous pagination request. If this is provided no other fields are required | 

### Return type

[**EnvironmentSharesList200Response**](EnvironmentSharesList200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentSharesUpdate

> EnvironmentSharesUpdate(ctx, environmentId, shareId).EnvironmentSharesUpdateRequest(environmentSharesUpdateRequest).Execute()

Update the details of the specified environment share. Used to change the config for PNG shares.

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the environment
    shareId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the environment share
    environmentSharesUpdateRequest := *openapiclient.NewEnvironmentSharesUpdateRequest() // EnvironmentSharesUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentSharesUpdate(context.Background(), environmentId, shareId).EnvironmentSharesUpdateRequest(environmentSharesUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentSharesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique ID of the environment | 
**shareId** | **string** | The unique ID of the environment share | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentSharesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **environmentSharesUpdateRequest** | [**EnvironmentSharesUpdateRequest**](EnvironmentSharesUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsCreate

> Environment EnvironmentsCreate(ctx).EnvironmentsCreateRequest(environmentsCreateRequest).Execute()

Create an environment using a query



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
    environmentsCreateRequest := *openapiclient.NewEnvironmentsCreateRequest("(vpc:vpc-1234 OR vpc:vpc-4567) AND CostCenter:DEV") // EnvironmentsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsCreate(context.Background()).EnvironmentsCreateRequest(environmentsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsCreate`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **environmentsCreateRequest** | [**EnvironmentsCreateRequest**](EnvironmentsCreateRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsDestroy

> EnvironmentsDestroy(ctx, environmentId).Execute()

Delete a custom environment



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the environment to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsDestroy(context.Background(), environmentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique ID of the environment to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsDestroyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsIndex

> EnvironmentsIndex200Response EnvironmentsIndex(ctx).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()

List environments

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
    page := int32(56) // int32 | The page of results to display. Based on page size and total results. (optional)
    pageSize := int32(56) // int32 | The number of results to return per page (optional)
    sort := "sort_example" // string | The field to sort by (optional)
    sortDir := "sortDir_example" // string | The direction to sort by (optional)
    token := "token_example" // string | The next page token returned from a previous pagination request. If this is provided no other fields are required (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsIndex(context.Background()).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsIndex`: EnvironmentsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsIndex`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The page of results to display. Based on page size and total results. | 
 **pageSize** | **int32** | The number of results to return per page | 
 **sort** | **string** | The field to sort by | 
 **sortDir** | **string** | The direction to sort by | 
 **token** | **string** | The next page token returned from a previous pagination request. If this is provided no other fields are required | 

### Return type

[**EnvironmentsIndex200Response**](EnvironmentsIndex200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsShow

> Environment EnvironmentsShow(ctx, environmentId).RevisionId(revisionId).Execute()

Get an environment

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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the environment
    revisionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the environment revision to view (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsShow(context.Background(), environmentId).RevisionId(revisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsShow`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique ID of the environment | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **revisionId** | **string** | The ID of the environment revision to view | 

### Return type

[**Environment**](Environment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnvironmentsUpdate

> Environment EnvironmentsUpdate(ctx, environmentId).EnvironmentsUpdateRequest(environmentsUpdateRequest).Execute()

Update a custom environment



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
    environmentId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the environment to update
    environmentsUpdateRequest := *openapiclient.NewEnvironmentsUpdateRequest("vpc:vpc-1234 AND CostCenter:DEV") // EnvironmentsUpdateRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EnvironmentsApi.EnvironmentsUpdate(context.Background(), environmentId).EnvironmentsUpdateRequest(environmentsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.EnvironmentsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnvironmentsUpdate`: Environment
    fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.EnvironmentsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentId** | **string** | The unique ID of the environment to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnvironmentsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentsUpdateRequest** | [**EnvironmentsUpdateRequest**](EnvironmentsUpdateRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

