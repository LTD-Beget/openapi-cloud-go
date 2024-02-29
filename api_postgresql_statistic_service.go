/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// PostgresqlStatisticServiceApiService PostgresqlStatisticServiceApi service
type PostgresqlStatisticServiceApiService service

type ApiPostgresqlStatisticServiceGetCpuRequest struct {
	ctx context.Context
	ApiService *PostgresqlStatisticServiceApiService
	serviceId string
	period *string
}

func (r ApiPostgresqlStatisticServiceGetCpuRequest) Period(period string) ApiPostgresqlStatisticServiceGetCpuRequest {
	r.period = &period
	return r
}

func (r ApiPostgresqlStatisticServiceGetCpuRequest) Execute() (*PostgresqlStatisticGetCpuResponse, *http.Response, error) {
	return r.ApiService.PostgresqlStatisticServiceGetCpuExecute(r)
}

/*
PostgresqlStatisticServiceGetCpu Method for PostgresqlStatisticServiceGetCpu

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiPostgresqlStatisticServiceGetCpuRequest
*/
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetCpu(ctx context.Context, serviceId string) ApiPostgresqlStatisticServiceGetCpuRequest {
	return ApiPostgresqlStatisticServiceGetCpuRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return PostgresqlStatisticGetCpuResponse
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetCpuExecute(r ApiPostgresqlStatisticServiceGetCpuRequest) (*PostgresqlStatisticGetCpuResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresqlStatisticGetCpuResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresqlStatisticServiceApiService.PostgresqlStatisticServiceGetCpu")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/postgresql/{service_id}/statistic/cpu"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
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

type ApiPostgresqlStatisticServiceGetCpuDetailsRequest struct {
	ctx context.Context
	ApiService *PostgresqlStatisticServiceApiService
	serviceId string
	period *string
}

func (r ApiPostgresqlStatisticServiceGetCpuDetailsRequest) Period(period string) ApiPostgresqlStatisticServiceGetCpuDetailsRequest {
	r.period = &period
	return r
}

func (r ApiPostgresqlStatisticServiceGetCpuDetailsRequest) Execute() (*PostgresqlStatisticGetCpuDetailsResponse, *http.Response, error) {
	return r.ApiService.PostgresqlStatisticServiceGetCpuDetailsExecute(r)
}

/*
PostgresqlStatisticServiceGetCpuDetails Method for PostgresqlStatisticServiceGetCpuDetails

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiPostgresqlStatisticServiceGetCpuDetailsRequest
*/
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetCpuDetails(ctx context.Context, serviceId string) ApiPostgresqlStatisticServiceGetCpuDetailsRequest {
	return ApiPostgresqlStatisticServiceGetCpuDetailsRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return PostgresqlStatisticGetCpuDetailsResponse
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetCpuDetailsExecute(r ApiPostgresqlStatisticServiceGetCpuDetailsRequest) (*PostgresqlStatisticGetCpuDetailsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresqlStatisticGetCpuDetailsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresqlStatisticServiceApiService.PostgresqlStatisticServiceGetCpuDetails")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/postgresql/{service_id}/statistic/cpu-details"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
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

type ApiPostgresqlStatisticServiceGetDiskRequest struct {
	ctx context.Context
	ApiService *PostgresqlStatisticServiceApiService
	serviceId string
	period *string
}

func (r ApiPostgresqlStatisticServiceGetDiskRequest) Period(period string) ApiPostgresqlStatisticServiceGetDiskRequest {
	r.period = &period
	return r
}

func (r ApiPostgresqlStatisticServiceGetDiskRequest) Execute() (*PostgresqlStatisticGetDiskResponse, *http.Response, error) {
	return r.ApiService.PostgresqlStatisticServiceGetDiskExecute(r)
}

/*
PostgresqlStatisticServiceGetDisk Method for PostgresqlStatisticServiceGetDisk

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiPostgresqlStatisticServiceGetDiskRequest
*/
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetDisk(ctx context.Context, serviceId string) ApiPostgresqlStatisticServiceGetDiskRequest {
	return ApiPostgresqlStatisticServiceGetDiskRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return PostgresqlStatisticGetDiskResponse
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetDiskExecute(r ApiPostgresqlStatisticServiceGetDiskRequest) (*PostgresqlStatisticGetDiskResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresqlStatisticGetDiskResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresqlStatisticServiceApiService.PostgresqlStatisticServiceGetDisk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/postgresql/{service_id}/statistic/disk"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
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

type ApiPostgresqlStatisticServiceGetDiskUsageRequest struct {
	ctx context.Context
	ApiService *PostgresqlStatisticServiceApiService
	serviceId string
	period *string
}

func (r ApiPostgresqlStatisticServiceGetDiskUsageRequest) Period(period string) ApiPostgresqlStatisticServiceGetDiskUsageRequest {
	r.period = &period
	return r
}

func (r ApiPostgresqlStatisticServiceGetDiskUsageRequest) Execute() (*PostgresqlStatisticGetDiskUsageResponse, *http.Response, error) {
	return r.ApiService.PostgresqlStatisticServiceGetDiskUsageExecute(r)
}

/*
PostgresqlStatisticServiceGetDiskUsage Method for PostgresqlStatisticServiceGetDiskUsage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiPostgresqlStatisticServiceGetDiskUsageRequest
*/
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetDiskUsage(ctx context.Context, serviceId string) ApiPostgresqlStatisticServiceGetDiskUsageRequest {
	return ApiPostgresqlStatisticServiceGetDiskUsageRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return PostgresqlStatisticGetDiskUsageResponse
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetDiskUsageExecute(r ApiPostgresqlStatisticServiceGetDiskUsageRequest) (*PostgresqlStatisticGetDiskUsageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresqlStatisticGetDiskUsageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresqlStatisticServiceApiService.PostgresqlStatisticServiceGetDiskUsage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/postgresql/{service_id}/statistic/disk-usage"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
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

type ApiPostgresqlStatisticServiceGetLoadAverageRequest struct {
	ctx context.Context
	ApiService *PostgresqlStatisticServiceApiService
	serviceId string
	period *string
}

func (r ApiPostgresqlStatisticServiceGetLoadAverageRequest) Period(period string) ApiPostgresqlStatisticServiceGetLoadAverageRequest {
	r.period = &period
	return r
}

func (r ApiPostgresqlStatisticServiceGetLoadAverageRequest) Execute() (*PostgresqlStatisticGetLoadAverageResponse, *http.Response, error) {
	return r.ApiService.PostgresqlStatisticServiceGetLoadAverageExecute(r)
}

/*
PostgresqlStatisticServiceGetLoadAverage Method for PostgresqlStatisticServiceGetLoadAverage

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiPostgresqlStatisticServiceGetLoadAverageRequest
*/
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetLoadAverage(ctx context.Context, serviceId string) ApiPostgresqlStatisticServiceGetLoadAverageRequest {
	return ApiPostgresqlStatisticServiceGetLoadAverageRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return PostgresqlStatisticGetLoadAverageResponse
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetLoadAverageExecute(r ApiPostgresqlStatisticServiceGetLoadAverageRequest) (*PostgresqlStatisticGetLoadAverageResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresqlStatisticGetLoadAverageResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresqlStatisticServiceApiService.PostgresqlStatisticServiceGetLoadAverage")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/postgresql/{service_id}/statistic/load-average"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
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

type ApiPostgresqlStatisticServiceGetMemoryRequest struct {
	ctx context.Context
	ApiService *PostgresqlStatisticServiceApiService
	serviceId string
	period *string
}

func (r ApiPostgresqlStatisticServiceGetMemoryRequest) Period(period string) ApiPostgresqlStatisticServiceGetMemoryRequest {
	r.period = &period
	return r
}

func (r ApiPostgresqlStatisticServiceGetMemoryRequest) Execute() (*PostgresqlStatisticGetMemoryResponse, *http.Response, error) {
	return r.ApiService.PostgresqlStatisticServiceGetMemoryExecute(r)
}

/*
PostgresqlStatisticServiceGetMemory Method for PostgresqlStatisticServiceGetMemory

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiPostgresqlStatisticServiceGetMemoryRequest
*/
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetMemory(ctx context.Context, serviceId string) ApiPostgresqlStatisticServiceGetMemoryRequest {
	return ApiPostgresqlStatisticServiceGetMemoryRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return PostgresqlStatisticGetMemoryResponse
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetMemoryExecute(r ApiPostgresqlStatisticServiceGetMemoryRequest) (*PostgresqlStatisticGetMemoryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresqlStatisticGetMemoryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresqlStatisticServiceApiService.PostgresqlStatisticServiceGetMemory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/postgresql/{service_id}/statistic/memory"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
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

type ApiPostgresqlStatisticServiceGetNetworkRequest struct {
	ctx context.Context
	ApiService *PostgresqlStatisticServiceApiService
	serviceId string
	period *string
}

func (r ApiPostgresqlStatisticServiceGetNetworkRequest) Period(period string) ApiPostgresqlStatisticServiceGetNetworkRequest {
	r.period = &period
	return r
}

func (r ApiPostgresqlStatisticServiceGetNetworkRequest) Execute() (*PostgresqlStatisticGetNetworkResponse, *http.Response, error) {
	return r.ApiService.PostgresqlStatisticServiceGetNetworkExecute(r)
}

/*
PostgresqlStatisticServiceGetNetwork Method for PostgresqlStatisticServiceGetNetwork

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiPostgresqlStatisticServiceGetNetworkRequest
*/
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetNetwork(ctx context.Context, serviceId string) ApiPostgresqlStatisticServiceGetNetworkRequest {
	return ApiPostgresqlStatisticServiceGetNetworkRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return PostgresqlStatisticGetNetworkResponse
func (a *PostgresqlStatisticServiceApiService) PostgresqlStatisticServiceGetNetworkExecute(r ApiPostgresqlStatisticServiceGetNetworkRequest) (*PostgresqlStatisticGetNetworkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PostgresqlStatisticGetNetworkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PostgresqlStatisticServiceApiService.PostgresqlStatisticServiceGetNetwork")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/postgresql/{service_id}/statistic/network"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.period != nil {
		localVarQueryParams.Add("period", parameterToString(*r.period, ""))
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
