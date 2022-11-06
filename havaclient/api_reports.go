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


// ReportsApiService ReportsApi service
type ReportsApiService service

type ApiReportsIndexRequest struct {
	ctx context.Context
	ApiService *ReportsApiService
	page *int32
	pageSize *int32
	sort *string
	sortDir *string
	token *string
}

// The page of results to display. Based on page size and total results.
func (r ApiReportsIndexRequest) Page(page int32) ApiReportsIndexRequest {
	r.page = &page
	return r
}

// The number of results to return per page
func (r ApiReportsIndexRequest) PageSize(pageSize int32) ApiReportsIndexRequest {
	r.pageSize = &pageSize
	return r
}

// The field to sort by
func (r ApiReportsIndexRequest) Sort(sort string) ApiReportsIndexRequest {
	r.sort = &sort
	return r
}

// The direction to sort by
func (r ApiReportsIndexRequest) SortDir(sortDir string) ApiReportsIndexRequest {
	r.sortDir = &sortDir
	return r
}

// The next page token returned from a previous pagination request. If this is provided no other fields are required
func (r ApiReportsIndexRequest) Token(token string) ApiReportsIndexRequest {
	r.token = &token
	return r
}

func (r ApiReportsIndexRequest) Execute() (*InlineResponse2007, *http.Response, error) {
	return r.ApiService.ReportsIndexExecute(r)
}

/*
ReportsIndex List all reports

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReportsIndexRequest
*/
func (a *ReportsApiService) ReportsIndex(ctx context.Context) ApiReportsIndexRequest {
	return ApiReportsIndexRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return InlineResponse2007
func (a *ReportsApiService) ReportsIndexExecute(r ApiReportsIndexRequest) (*InlineResponse2007, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *InlineResponse2007
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsApiService.ReportsIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports"

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

type ApiReportsReportIdExportPostRequest struct {
	ctx context.Context
	ApiService *ReportsApiService
	reportId string
	exportFormat *string
	reportIdExportBody *ReportIdExportBody
}

// The format to export the report to
func (r ApiReportsReportIdExportPostRequest) ExportFormat(exportFormat string) ApiReportsReportIdExportPostRequest {
	r.exportFormat = &exportFormat
	return r
}

func (r ApiReportsReportIdExportPostRequest) ReportIdExportBody(reportIdExportBody ReportIdExportBody) ApiReportsReportIdExportPostRequest {
	r.reportIdExportBody = &reportIdExportBody
	return r
}

func (r ApiReportsReportIdExportPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ReportsReportIdExportPostExecute(r)
}

/*
ReportsReportIdExportPost Export a report

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param reportId The unique ID of the report to export
 @return ApiReportsReportIdExportPostRequest
*/
func (a *ReportsApiService) ReportsReportIdExportPost(ctx context.Context, reportId string) ApiReportsReportIdExportPostRequest {
	return ApiReportsReportIdExportPostRequest{
		ApiService: a,
		ctx: ctx,
		reportId: reportId,
	}
}

// Execute executes the request
func (a *ReportsApiService) ReportsReportIdExportPostExecute(r ApiReportsReportIdExportPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsApiService.ReportsReportIdExportPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/{report_id}/export"
	localVarPath = strings.Replace(localVarPath, "{"+"report_id"+"}", url.PathEscape(parameterToString(r.reportId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.exportFormat == nil {
		return nil, reportError("exportFormat is required and must be specified")
	}
	if r.reportIdExportBody == nil {
		return nil, reportError("reportIdExportBody is required and must be specified")
	}

	localVarQueryParams.Add("export_format", parameterToString(*r.exportFormat, ""))
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
	localVarPostBody = r.reportIdExportBody
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
		if localVarHTTPResponse.StatusCode == 422 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiReportsShowRequest struct {
	ctx context.Context
	ApiService *ReportsApiService
	reportId string
}

func (r ApiReportsShowRequest) Execute() (*Report, *http.Response, error) {
	return r.ApiService.ReportsShowExecute(r)
}

/*
ReportsShow Get a report

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param reportId
 @return ApiReportsShowRequest
*/
func (a *ReportsApiService) ReportsShow(ctx context.Context, reportId string) ApiReportsShowRequest {
	return ApiReportsShowRequest{
		ApiService: a,
		ctx: ctx,
		reportId: reportId,
	}
}

// Execute executes the request
//  @return Report
func (a *ReportsApiService) ReportsShowExecute(r ApiReportsShowRequest) (*Report, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Report
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ReportsApiService.ReportsShow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/reports/{report_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"report_id"+"}", url.PathEscape(parameterToString(r.reportId, "")), -1)

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