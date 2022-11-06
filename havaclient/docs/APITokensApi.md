# \APITokensApi

All URIs are relative to *https://api.hava.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiTokensCreate**](APITokensApi.md#ApiTokensCreate) | **Post** /accounts/{account_id}/tokens | Generate a new token
[**ApiTokensDelete**](APITokensApi.md#ApiTokensDelete) | **Delete** /accounts/{account_id}/tokens/{token_id} | Delete an API token
[**ApiTokensIndex**](APITokensApi.md#ApiTokensIndex) | **Get** /accounts/{account_id}/tokens | List all API tokens in the specified account



## ApiTokensCreate

> Token ApiTokensCreate(ctx, accountId).AccountIdTokensBody(accountIdTokensBody).Execute()

Generate a new token



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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account to create a team in
    accountIdTokensBody := *openapiclient.NewAccountIdTokensBody() // AccountIdTokensBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APITokensApi.ApiTokensCreate(context.Background(), accountId).AccountIdTokensBody(accountIdTokensBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensApi.ApiTokensCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTokensCreate`: Token
    fmt.Fprintf(os.Stdout, "Response from `APITokensApi.ApiTokensCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account to create a team in | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTokensCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdTokensBody** | [**AccountIdTokensBody**](AccountIdTokensBody.md) |  | 

### Return type

[**Token**](Token.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTokensDelete

> ApiTokensDelete(ctx, accountId, tokenId).Execute()

Delete an API token

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the API token is in
    tokenId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the API token to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APITokensApi.ApiTokensDelete(context.Background(), accountId, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensApi.ApiTokensDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the API token is in | 
**tokenId** | **string** | The unique ID of the API token to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTokensDeleteRequest struct via the builder pattern


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


## ApiTokensIndex

> InlineResponse2003 ApiTokensIndex(ctx, accountId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()

List all API tokens in the specified account

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account to list tokens for
    page := int32(56) // int32 | The page of results to display. Based on page size and total results. (optional)
    pageSize := int32(56) // int32 | The number of results to return per page (optional)
    sort := "sort_example" // string | The field to sort by (optional)
    sortDir := "sortDir_example" // string | The direction to sort by (optional)
    token := "token_example" // string | The next page token returned from a previous pagination request. If this is provided no other fields are required (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.APITokensApi.ApiTokensIndex(context.Background(), accountId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `APITokensApi.ApiTokensIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApiTokensIndex`: InlineResponse2003
    fmt.Fprintf(os.Stdout, "Response from `APITokensApi.ApiTokensIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account to list tokens for | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiTokensIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of results to display. Based on page size and total results. | 
 **pageSize** | **int32** | The number of results to return per page | 
 **sort** | **string** | The field to sort by | 
 **sortDir** | **string** | The direction to sort by | 
 **token** | **string** | The next page token returned from a previous pagination request. If this is provided no other fields are required | 

### Return type

[**InlineResponse2003**](InlineResponse2003.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

