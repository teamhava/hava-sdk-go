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


// TeamsApiService TeamsApi service
type TeamsApiService service

type ApiTeamMembersAddRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	accountId string
	teamId string
	email *string
}

// Email address of the user to add
func (r ApiTeamMembersAddRequest) Email(email string) ApiTeamMembersAddRequest {
	r.email = &email
	return r
}

func (r ApiTeamMembersAddRequest) Execute() (*Team, *http.Response, error) {
	return r.ApiService.TeamMembersAddExecute(r)
}

/*
TeamMembersAdd Add member to team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account the team is in
 @param teamId The unique ID of the team to add the user to
 @return ApiTeamMembersAddRequest
*/
func (a *TeamsApiService) TeamMembersAdd(ctx context.Context, accountId string, teamId string) ApiTeamMembersAddRequest {
	return ApiTeamMembersAddRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return Team
func (a *TeamsApiService) TeamMembersAddExecute(r ApiTeamMembersAddRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.TeamMembersAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/teams/{team_id}/members"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"team_id"+"}", url.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
	}

	localVarQueryParams.Add("email", parameterToString(*r.email, ""))
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

type ApiTeamMembersInviteRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	accountId string
	teamId string
	email *string
}

// Email address of the user to add
func (r ApiTeamMembersInviteRequest) Email(email string) ApiTeamMembersInviteRequest {
	r.email = &email
	return r
}

func (r ApiTeamMembersInviteRequest) Execute() (*Team, *http.Response, error) {
	return r.ApiService.TeamMembersInviteExecute(r)
}

/*
TeamMembersInvite Invite a user to Hava and add them to the team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account the team is in
 @param teamId The unique ID of the team to invite the user to
 @return ApiTeamMembersInviteRequest
*/
func (a *TeamsApiService) TeamMembersInvite(ctx context.Context, accountId string, teamId string) ApiTeamMembersInviteRequest {
	return ApiTeamMembersInviteRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return Team
func (a *TeamsApiService) TeamMembersInviteExecute(r ApiTeamMembersInviteRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.TeamMembersInvite")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/teams/{team_id}/members/invite"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"team_id"+"}", url.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.email == nil {
		return localVarReturnValue, nil, reportError("email is required and must be specified")
	}

	localVarQueryParams.Add("email", parameterToString(*r.email, ""))
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

type ApiTeamMembersRemoveRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	accountId string
	teamId string
	userId string
}

func (r ApiTeamMembersRemoveRequest) Execute() (*Team, *http.Response, error) {
	return r.ApiService.TeamMembersRemoveExecute(r)
}

/*
TeamMembersRemove Remove member from team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account the team is in
 @param teamId The unique ID of the team to remove the member from
 @param userId The unique ID of the user to remove from the team
 @return ApiTeamMembersRemoveRequest
*/
func (a *TeamsApiService) TeamMembersRemove(ctx context.Context, accountId string, teamId string, userId string) ApiTeamMembersRemoveRequest {
	return ApiTeamMembersRemoveRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		teamId: teamId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Team
func (a *TeamsApiService) TeamMembersRemoveExecute(r ApiTeamMembersRemoveRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.TeamMembersRemove")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/teams/{team_id}/members/{user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"team_id"+"}", url.PathEscape(parameterToString(r.teamId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user_id"+"}", url.PathEscape(parameterToString(r.userId, "")), -1)

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

type ApiTeamsCreateRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	accountId string
	accountIdTeamsBody *AccountIdTeamsBody
}

func (r ApiTeamsCreateRequest) AccountIdTeamsBody(accountIdTeamsBody AccountIdTeamsBody) ApiTeamsCreateRequest {
	r.accountIdTeamsBody = &accountIdTeamsBody
	return r
}

func (r ApiTeamsCreateRequest) Execute() (*Team, *http.Response, error) {
	return r.ApiService.TeamsCreateExecute(r)
}

/*
TeamsCreate Create a team

By default created teams have read only access. You can pass 'Admin' in the roles array to give them permission to create and modify sources and environments.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account to create a team in
 @return ApiTeamsCreateRequest
*/
func (a *TeamsApiService) TeamsCreate(ctx context.Context, accountId string) ApiTeamsCreateRequest {
	return ApiTeamsCreateRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return Team
func (a *TeamsApiService) TeamsCreateExecute(r ApiTeamsCreateRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.TeamsCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/teams"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.accountIdTeamsBody == nil {
		return localVarReturnValue, nil, reportError("accountIdTeamsBody is required and must be specified")
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
	localVarPostBody = r.accountIdTeamsBody
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

type ApiTeamsDestroyRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	accountId string
	teamId string
}

func (r ApiTeamsDestroyRequest) Execute() (*http.Response, error) {
	return r.ApiService.TeamsDestroyExecute(r)
}

/*
TeamsDestroy Delete a team and remove all members from your account

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account the team is in
 @param teamId The unique ID of the team to delete
 @return ApiTeamsDestroyRequest
*/
func (a *TeamsApiService) TeamsDestroy(ctx context.Context, accountId string, teamId string) ApiTeamsDestroyRequest {
	return ApiTeamsDestroyRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		teamId: teamId,
	}
}

// Execute executes the request
func (a *TeamsApiService) TeamsDestroyExecute(r ApiTeamsDestroyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.TeamsDestroy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/teams/{team_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"team_id"+"}", url.PathEscape(parameterToString(r.teamId, "")), -1)

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

type ApiTeamsIndexRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	accountId string
	page *int32
	pageSize *int32
	sort *string
	sortDir *string
	token *string
}

// The page of results to display. Based on page size and total results.
func (r ApiTeamsIndexRequest) Page(page int32) ApiTeamsIndexRequest {
	r.page = &page
	return r
}

// The number of results to return per page
func (r ApiTeamsIndexRequest) PageSize(pageSize int32) ApiTeamsIndexRequest {
	r.pageSize = &pageSize
	return r
}

// The field to sort by
func (r ApiTeamsIndexRequest) Sort(sort string) ApiTeamsIndexRequest {
	r.sort = &sort
	return r
}

// The direction to sort by
func (r ApiTeamsIndexRequest) SortDir(sortDir string) ApiTeamsIndexRequest {
	r.sortDir = &sortDir
	return r
}

// The next page token returned from a previous pagination request. If this is provided no other fields are required
func (r ApiTeamsIndexRequest) Token(token string) ApiTeamsIndexRequest {
	r.token = &token
	return r
}

func (r ApiTeamsIndexRequest) Execute() (*InlineResponse2001, *http.Response, error) {
	return r.ApiService.TeamsIndexExecute(r)
}

/*
TeamsIndex List all teams

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account to list teams for
 @return ApiTeamsIndexRequest
*/
func (a *TeamsApiService) TeamsIndex(ctx context.Context, accountId string) ApiTeamsIndexRequest {
	return ApiTeamsIndexRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
	}
}

// Execute executes the request
//  @return InlineResponse2001
func (a *TeamsApiService) TeamsIndexExecute(r ApiTeamsIndexRequest) (*InlineResponse2001, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse2001
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.TeamsIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/teams"
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

type ApiTeamsShowRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	accountId string
	teamId string
}

func (r ApiTeamsShowRequest) Execute() (*Team, *http.Response, error) {
	return r.ApiService.TeamsShowExecute(r)
}

/*
TeamsShow Get a team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account the team is in
 @param teamId The unique ID of the team to return
 @return ApiTeamsShowRequest
*/
func (a *TeamsApiService) TeamsShow(ctx context.Context, accountId string, teamId string) ApiTeamsShowRequest {
	return ApiTeamsShowRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return Team
func (a *TeamsApiService) TeamsShowExecute(r ApiTeamsShowRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.TeamsShow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/teams/{team_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"team_id"+"}", url.PathEscape(parameterToString(r.teamId, "")), -1)

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

type ApiTeamsUpdateRequest struct {
	ctx context.Context
	ApiService *TeamsApiService
	accountId string
	teamId string
	teamsTeamIdBody *TeamsTeamIdBody
}

func (r ApiTeamsUpdateRequest) TeamsTeamIdBody(teamsTeamIdBody TeamsTeamIdBody) ApiTeamsUpdateRequest {
	r.teamsTeamIdBody = &teamsTeamIdBody
	return r
}

func (r ApiTeamsUpdateRequest) Execute() (*Team, *http.Response, error) {
	return r.ApiService.TeamsUpdateExecute(r)
}

/*
TeamsUpdate Update a team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param accountId The unique ID of the account the team is in
 @param teamId The unique ID of the team to update
 @return ApiTeamsUpdateRequest
*/
func (a *TeamsApiService) TeamsUpdate(ctx context.Context, accountId string, teamId string) ApiTeamsUpdateRequest {
	return ApiTeamsUpdateRequest{
		ApiService: a,
		ctx: ctx,
		accountId: accountId,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return Team
func (a *TeamsApiService) TeamsUpdateExecute(r ApiTeamsUpdateRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsApiService.TeamsUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accounts/{account_id}/teams/{team_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"account_id"+"}", url.PathEscape(parameterToString(r.accountId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"team_id"+"}", url.PathEscape(parameterToString(r.teamId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.teamsTeamIdBody == nil {
		return localVarReturnValue, nil, reportError("teamsTeamIdBody is required and must be specified")
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
	localVarPostBody = r.teamsTeamIdBody
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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
