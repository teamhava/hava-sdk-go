# \ProjectsApi

All URIs are relative to *https://api.hava.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProjectMembersAdd**](ProjectsApi.md#ProjectMembersAdd) | **Post** /accounts/{account_id}/projects/{project_id}/members | Add member to project
[**ProjectMembersRemove**](ProjectsApi.md#ProjectMembersRemove) | **Delete** /accounts/{account_id}/projects/{project_id}/members/{user_id} | Remove member from project
[**ProjectsCreate**](ProjectsApi.md#ProjectsCreate) | **Post** /accounts/{account_id}/projects | Create a project
[**ProjectsDestroy**](ProjectsApi.md#ProjectsDestroy) | **Delete** /accounts/{account_id}/projects/{project_id} | Delete a project
[**ProjectsIndex**](ProjectsApi.md#ProjectsIndex) | **Get** /accounts/{account_id}/projects | List all projects
[**ProjectsShow**](ProjectsApi.md#ProjectsShow) | **Get** /accounts/{account_id}/projects/{project_id} | Get a project
[**ProjectsUpdate**](ProjectsApi.md#ProjectsUpdate) | **Put** /accounts/{account_id}/projects/{project_id} | Update a project



## ProjectMembersAdd

> Project ProjectMembersAdd(ctx, accountId, projectId).UserId(userId).Execute()

Add member to project

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the project is in
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the project to add the user to
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique user ID of a user in your account to add to this project

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectMembersAdd(context.Background(), accountId, projectId).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectMembersAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectMembersAdd`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectMembersAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the project is in | 
**projectId** | **string** | The unique ID of the project to add the user to | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectMembersAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userId** | **string** | The unique user ID of a user in your account to add to this project | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectMembersRemove

> Project ProjectMembersRemove(ctx, accountId, projectId, userId).Execute()

Remove member from project

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the project is in
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the project to remove the member from
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the user to remove from the project

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectMembersRemove(context.Background(), accountId, projectId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectMembersRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectMembersRemove`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectMembersRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the project is in | 
**projectId** | **string** | The unique ID of the project to remove the member from | 
**userId** | **string** | The unique ID of the user to remove from the project | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectMembersRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsCreate

> Project ProjectsCreate(ctx, accountId).ProjectsCreateRequest(projectsCreateRequest).Execute()

Create a project

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account to create the project in
    projectsCreateRequest := *openapiclient.NewProjectsCreateRequest() // ProjectsCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsCreate(context.Background(), accountId).ProjectsCreateRequest(projectsCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsCreate`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account to create the project in | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectsCreateRequest** | [**ProjectsCreateRequest**](ProjectsCreateRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsDestroy

> ProjectsDestroy(ctx, accountId, projectId).Execute()

Delete a project



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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the project is in
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the project to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsDestroy(context.Background(), accountId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the project is in | 
**projectId** | **string** | The unique ID of the project to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsDestroyRequest struct via the builder pattern


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


## ProjectsIndex

> ProjectsIndex200Response ProjectsIndex(ctx, accountId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()

List all projects

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the projects are in
    page := int32(56) // int32 | The page of results to display. Based on page size and total results. (optional)
    pageSize := int32(56) // int32 | The number of results to return per page (optional)
    sort := "sort_example" // string | The field to sort by (optional)
    sortDir := "sortDir_example" // string | The direction to sort by (optional)
    token := "token_example" // string | The next page token returned from a previous pagination request. If this is provided no other fields are required (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsIndex(context.Background(), accountId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsIndex`: ProjectsIndex200Response
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the projects are in | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of results to display. Based on page size and total results. | 
 **pageSize** | **int32** | The number of results to return per page | 
 **sort** | **string** | The field to sort by | 
 **sortDir** | **string** | The direction to sort by | 
 **token** | **string** | The next page token returned from a previous pagination request. If this is provided no other fields are required | 

### Return type

[**ProjectsIndex200Response**](ProjectsIndex200Response.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsShow

> Project ProjectsShow(ctx, accountId, projectId).Execute()

Get a project

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the project is in
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the project to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsShow(context.Background(), accountId, projectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsShow`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the project is in | 
**projectId** | **string** | The unique ID of the project to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProjectsUpdate

> Project ProjectsUpdate(ctx, accountId, projectId).ProjectsUpdateRequest(projectsUpdateRequest).Execute()

Update a project

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the project is in
    projectId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the project to update
    projectsUpdateRequest := *openapiclient.NewProjectsUpdateRequest() // ProjectsUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.ProjectsUpdate(context.Background(), accountId, projectId).ProjectsUpdateRequest(projectsUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.ProjectsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProjectsUpdate`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.ProjectsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the project is in | 
**projectId** | **string** | The unique ID of the project to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiProjectsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectsUpdateRequest** | [**ProjectsUpdateRequest**](ProjectsUpdateRequest.md) |  | 

### Return type

[**Project**](Project.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

