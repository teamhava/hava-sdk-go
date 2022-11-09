# Go API client for havaclient

Hava API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.2
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import havaclient "github.com/teamhava/hava-sdk-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), havaclient.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), havaclient.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), havaclient.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), havaclient.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.hava.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*SourcesApi* | [**SourcesCreate**](docs/SourcesApi.md#sourcescreate) | **Post** /sources | Create a source
*SourcesApi* | [**SourcesDestroy**](docs/SourcesApi.md#sourcesdestroy) | **Delete** /sources/{source_id} | Delete a source and all associated resources and environments
*SourcesApi* | [**SourcesIndex**](docs/SourcesApi.md#sourcesindex) | **Get** /sources | List all sources
*SourcesApi* | [**SourcesShow**](docs/SourcesApi.md#sourcesshow) | **Get** /sources/{source_id} | Get a source
*SourcesApi* | [**SourcesSync**](docs/SourcesApi.md#sourcessync) | **Post** /sources/{source_id}/sync | Import the latest resources for this source
*SourcesApi* | [**SourcesUpdate**](docs/SourcesApi.md#sourcesupdate) | **Put** /sources/{source_id} | Update a source


## Documentation For Models

 - [ApiTokensCreateRequest](docs/ApiTokensCreateRequest.md)
 - [ApiTokensIndex200Response](docs/ApiTokensIndex200Response.md)
 - [Environment](docs/Environment.md)
 - [EnvironmentCurrentRevision](docs/EnvironmentCurrentRevision.md)
 - [EnvironmentFacet](docs/EnvironmentFacet.md)
 - [EnvironmentLatestRevisions](docs/EnvironmentLatestRevisions.md)
 - [EnvironmentRevision](docs/EnvironmentRevision.md)
 - [EnvironmentRevisionsIndex200Response](docs/EnvironmentRevisionsIndex200Response.md)
 - [EnvironmentShare](docs/EnvironmentShare.md)
 - [EnvironmentSharesCreateRequest](docs/EnvironmentSharesCreateRequest.md)
 - [EnvironmentSharesCreateRequestConfig](docs/EnvironmentSharesCreateRequestConfig.md)
 - [EnvironmentSharesList200Response](docs/EnvironmentSharesList200Response.md)
 - [EnvironmentSharesUpdateRequest](docs/EnvironmentSharesUpdateRequest.md)
 - [EnvironmentSummary](docs/EnvironmentSummary.md)
 - [EnvironmentSummarySourcesInner](docs/EnvironmentSummarySourcesInner.md)
 - [EnvironmentSummaryViewsInner](docs/EnvironmentSummaryViewsInner.md)
 - [EnvironmentsCreateRequest](docs/EnvironmentsCreateRequest.md)
 - [EnvironmentsIndex200Response](docs/EnvironmentsIndex200Response.md)
 - [EnvironmentsUpdateRequest](docs/EnvironmentsUpdateRequest.md)
 - [ErrorInner](docs/ErrorInner.md)
 - [JobsShow200Response](docs/JobsShow200Response.md)
 - [Project](docs/Project.md)
 - [ProjectMembersInner](docs/ProjectMembersInner.md)
 - [ProjectsCreateRequest](docs/ProjectsCreateRequest.md)
 - [ProjectsIndex200Response](docs/ProjectsIndex200Response.md)
 - [ProjectsUpdateRequest](docs/ProjectsUpdateRequest.md)
 - [Report](docs/Report.md)
 - [ReportSummary](docs/ReportSummary.md)
 - [ReportSummarySource](docs/ReportSummarySource.md)
 - [ReportsIndex200Response](docs/ReportsIndex200Response.md)
 - [ReportsReportIdExportPostRequest](docs/ReportsReportIdExportPostRequest.md)
 - [Resource](docs/Resource.md)
 - [ResourceConnectionsInner](docs/ResourceConnectionsInner.md)
 - [ResourceTagsInner](docs/ResourceTagsInner.md)
 - [Source](docs/Source.md)
 - [SourcesAWSCAR](docs/SourcesAWSCAR.md)
 - [SourcesAWSKey](docs/SourcesAWSKey.md)
 - [SourcesAzureCredentials](docs/SourcesAzureCredentials.md)
 - [SourcesCreateRequest](docs/SourcesCreateRequest.md)
 - [SourcesGCPServiceAccountCredentials](docs/SourcesGCPServiceAccountCredentials.md)
 - [SourcesIndex200Response](docs/SourcesIndex200Response.md)
 - [SourcesUpdateRequest](docs/SourcesUpdateRequest.md)
 - [Team](docs/Team.md)
 - [TeamMembersInner](docs/TeamMembersInner.md)
 - [TeamsCreateRequest](docs/TeamsCreateRequest.md)
 - [TeamsIndex200Response](docs/TeamsIndex200Response.md)
 - [TeamsUpdateRequest](docs/TeamsUpdateRequest.md)
 - [Token](docs/Token.md)
 - [TokenSummary](docs/TokenSummary.md)
 - [View](docs/View.md)
 - [ViewResourcesInner](docs/ViewResourcesInner.md)
 - [ViewsExportRequest](docs/ViewsExportRequest.md)


## Documentation For Authorization



### api_key

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@hava.io

