# \TeamsApi

All URIs are relative to *https://api.hava.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamMembersAdd**](TeamsApi.md#TeamMembersAdd) | **Post** /accounts/{account_id}/teams/{team_id}/members | Add member to team
[**TeamMembersInvite**](TeamsApi.md#TeamMembersInvite) | **Post** /accounts/{account_id}/teams/{team_id}/members/invite | Invite a user to Hava and add them to the team
[**TeamMembersRemove**](TeamsApi.md#TeamMembersRemove) | **Delete** /accounts/{account_id}/teams/{team_id}/members/{user_id} | Remove member from team
[**TeamsCreate**](TeamsApi.md#TeamsCreate) | **Post** /accounts/{account_id}/teams | Create a team
[**TeamsDestroy**](TeamsApi.md#TeamsDestroy) | **Delete** /accounts/{account_id}/teams/{team_id} | Delete a team and remove all members from your account
[**TeamsIndex**](TeamsApi.md#TeamsIndex) | **Get** /accounts/{account_id}/teams | List all teams
[**TeamsShow**](TeamsApi.md#TeamsShow) | **Get** /accounts/{account_id}/teams/{team_id} | Get a team
[**TeamsUpdate**](TeamsApi.md#TeamsUpdate) | **Put** /accounts/{account_id}/teams/{team_id} | Update a team



## TeamMembersAdd

> Team TeamMembersAdd(ctx, accountId, teamId).Email(email).Execute()

Add member to team

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the team is in
    teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the team to add the user to
    email := "email_example" // string | Email address of the user to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamMembersAdd(context.Background(), accountId, teamId).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamMembersAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamMembersAdd`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamMembersAdd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the team is in | 
**teamId** | **string** | The unique ID of the team to add the user to | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamMembersAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **email** | **string** | Email address of the user to add | 

### Return type

[**Team**](Team.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamMembersInvite

> Team TeamMembersInvite(ctx, accountId, teamId).Email(email).Execute()

Invite a user to Hava and add them to the team

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the team is in
    teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the team to invite the user to
    email := "email_example" // string | Email address of the user to add

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamMembersInvite(context.Background(), accountId, teamId).Email(email).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamMembersInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamMembersInvite`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamMembersInvite`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the team is in | 
**teamId** | **string** | The unique ID of the team to invite the user to | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamMembersInviteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **email** | **string** | Email address of the user to add | 

### Return type

[**Team**](Team.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamMembersRemove

> Team TeamMembersRemove(ctx, accountId, teamId, userId).Execute()

Remove member from team

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the team is in
    teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the team to remove the member from
    userId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the user to remove from the team

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamMembersRemove(context.Background(), accountId, teamId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamMembersRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamMembersRemove`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamMembersRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the team is in | 
**teamId** | **string** | The unique ID of the team to remove the member from | 
**userId** | **string** | The unique ID of the user to remove from the team | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamMembersRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**Team**](Team.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsCreate

> Team TeamsCreate(ctx, accountId).AccountIdTeamsBody(accountIdTeamsBody).Execute()

Create a team



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
    accountIdTeamsBody := *openapiclient.NewAccountIdTeamsBody() // AccountIdTeamsBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsCreate(context.Background(), accountId).AccountIdTeamsBody(accountIdTeamsBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsCreate`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account to create a team in | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accountIdTeamsBody** | [**AccountIdTeamsBody**](AccountIdTeamsBody.md) |  | 

### Return type

[**Team**](Team.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsDestroy

> TeamsDestroy(ctx, accountId, teamId).Execute()

Delete a team and remove all members from your account

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the team is in
    teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the team to delete

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsDestroy(context.Background(), accountId, teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsDestroy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the team is in | 
**teamId** | **string** | The unique ID of the team to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsDestroyRequest struct via the builder pattern


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


## TeamsIndex

> InlineResponse2001 TeamsIndex(ctx, accountId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()

List all teams

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account to list teams for
    page := int32(56) // int32 | The page of results to display. Based on page size and total results. (optional)
    pageSize := int32(56) // int32 | The number of results to return per page (optional)
    sort := "sort_example" // string | The field to sort by (optional)
    sortDir := "sortDir_example" // string | The direction to sort by (optional)
    token := "token_example" // string | The next page token returned from a previous pagination request. If this is provided no other fields are required (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsIndex(context.Background(), accountId).Page(page).PageSize(pageSize).Sort(sort).SortDir(sortDir).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsIndex``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsIndex`: InlineResponse2001
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsIndex`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account to list teams for | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsIndexRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of results to display. Based on page size and total results. | 
 **pageSize** | **int32** | The number of results to return per page | 
 **sort** | **string** | The field to sort by | 
 **sortDir** | **string** | The direction to sort by | 
 **token** | **string** | The next page token returned from a previous pagination request. If this is provided no other fields are required | 

### Return type

[**InlineResponse2001**](InlineResponse2001.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsShow

> Team TeamsShow(ctx, accountId, teamId).Execute()

Get a team

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the team is in
    teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the team to return

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsShow(context.Background(), accountId, teamId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsShow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsShow`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsShow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the team is in | 
**teamId** | **string** | The unique ID of the team to return | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsShowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Team**](Team.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TeamsUpdate

> Team TeamsUpdate(ctx, accountId, teamId).TeamsTeamIdBody(teamsTeamIdBody).Execute()

Update a team

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
    accountId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the account the team is in
    teamId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The unique ID of the team to update
    teamsTeamIdBody := *openapiclient.NewTeamsTeamIdBody() // TeamsTeamIdBody | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamsApi.TeamsUpdate(context.Background(), accountId, teamId).TeamsTeamIdBody(teamsTeamIdBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamsApi.TeamsUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TeamsUpdate`: Team
    fmt.Fprintf(os.Stdout, "Response from `TeamsApi.TeamsUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique ID of the account the team is in | 
**teamId** | **string** | The unique ID of the team to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiTeamsUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **teamsTeamIdBody** | [**TeamsTeamIdBody**](TeamsTeamIdBody.md) |  | 

### Return type

[**Team**](Team.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

