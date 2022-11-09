# Hava.io golang SDK

This is an early alpha version of our sdk, it's not ready to be used


## Generating a new version

When generating a new version of this API usign the open api generator, so manual intervention is required. There is currently a bug in the openapi generator that results in the generator not being able to handle arrays of error messages returned from the API.

Update the api_*.go code to select the first element of the array

Find the line matching the original below and update to to the updated line

Original:

```go
newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
```

Updated:

```go
newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v[0])
```
