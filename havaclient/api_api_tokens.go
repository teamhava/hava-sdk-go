/*
Hava

Hava API

API version: 1.1.1
Contact: support@hava.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package havaclient

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// APITokensApiService APITokensApi service
type APITokensApiService service

type ApiApiTokensCreateRequest struct {
	ctx context.Context
	ApiService *APITokensApiService
	accountId string
	accountIdTokensBody *AccountIdTokensBody
}

func (r ApiApiTokensCreateRequest) AccountIdTokensBody(accountIdTokensBody AccountIdTokensBody) ApiApiTokensCreateRequest {
	r.accountIdTokensBody = &accountIdTokensBody
	return r
}

func (r ApiApiTokensCreateRequest) Execute() (*Token, *http.Response, error) {
	return r.ApiService.ApiTokensCreateExecute(r)
}

/*
ApiTokensCreate Generate a new token

This endpoint will generate and return an API token used to access the API. This is the only time the token will be displayed.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account to create a team in
 @return ApiApiTokensCreateRequest
*/
func (a *APITokensApiService) ApiTokensCreate(ctx context.Context, accountId string) ApiApiTokensCreateRequest {
	return ApiApiTokensCreateRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return Token
func (a *APITokensApiService) ApiTokensCreateExecute(r ApiApiTokensCreateRequest) (*Token, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Token
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APITokensApiService.ApiTokensCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountIdTokensBody == nil {
		return localVarReturnValue, nil, reportError("accountIdTokensBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.accountIdTokensBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiTokensDeleteRequest struct {
	ctx context.Context
	ApiService *APITokensApiService
	accountId string
	tokenId string
}

func (r ApiApiTokensDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiTokensDeleteExecute(r)
}

/*
ApiTokensDelete Delete an API token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account the API token is in
 @param tokenId The unique ID of the API token to delete
 @return ApiApiTokensDeleteRequest
*/
func (a *APITokensApiService) ApiTokensDelete(ctx context.Context, accountId string, tokenId string) ApiApiTokensDeleteRequest {
	return ApiApiTokensDeleteRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		tokenId: tokenId,
	}
}

// Execute executes the request
func (a *APITokensApiService) ApiTokensDeleteExecute(r ApiApiTokensDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APITokensApiService.ApiTokensDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/tokens/{token_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"token_id"+"}", url.PathEscape(parameterToString(r.tokenId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiTokensIndexRequest struct {
	ctx context.Context
	ApiService *APITokensApiService
	accountId string
	page *int32
	pageSize *int32
	sort *string
	sortDir *string
	token *string
}

// The page of results to display. Based on page size and total results.
func (r ApiApiTokensIndexRequest) Page(page int32) ApiApiTokensIndexRequest {
	r.page = &page
	return r
}

// The number of results to return per page
func (r ApiApiTokensIndexRequest) PageSize(pageSize int32) ApiApiTokensIndexRequest {
	r.pageSize = &pageSize
	return r
}

// The field to sort by
func (r ApiApiTokensIndexRequest) Sort(sort string) ApiApiTokensIndexRequest {
	r.sort = &sort
	return r
}

// The direction to sort by
func (r ApiApiTokensIndexRequest) SortDir(sortDir string) ApiApiTokensIndexRequest {
	r.sortDir = &sortDir
	return r
}

// The next page token returned from a previous pagination request. If this is provided no other fields are required
func (r ApiApiTokensIndexRequest) Token(token string) ApiApiTokensIndexRequest {
	r.token = &token
	return r
}

func (r ApiApiTokensIndexRequest) Execute() (*InlineResponse2003, *http.Response, error) {
	return r.ApiService.ApiTokensIndexExecute(r)
}

/*
ApiTokensIndex List all API tokens in the specified account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account to list tokens for
 @return ApiApiTokensIndexRequest
*/
func (a *APITokensApiService) ApiTokensIndex(ctx context.Context, accountId string) ApiApiTokensIndexRequest {
	return ApiApiTokensIndexRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return InlineResponse2003
func (a *APITokensApiService) ApiTokensIndexExecute(r ApiApiTokensIndexRequest) (*InlineResponse2003, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse2003
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "APITokensApiService.ApiTokensIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/tokens"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		localVarQueryParams.Add("page", parameterToString(*r.page, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.sort != nil {
		localVarQueryParams.Add("sort", parameterToString(*r.sort, ""))
	}
	if r.sortDir != nil {
		localVarQueryParams.Add("sort_dir", parameterToString(*r.sortDir, ""))
	}
	if r.token != nil {
		localVarQueryParams.Add("token", parameterToString(*r.token, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
