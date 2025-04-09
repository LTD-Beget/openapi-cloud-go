/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.1
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


// CloudServiceApiService CloudServiceApi service
type CloudServiceApiService service

type ApiCloudServiceChangeConfigurationRequest struct {
	ctx context.Context
	ApiService *CloudServiceApiService
	serviceId string
	cloudChangeConfigurationRequest *CloudChangeConfigurationRequest
}

func (r ApiCloudServiceChangeConfigurationRequest) CloudChangeConfigurationRequest(cloudChangeConfigurationRequest CloudChangeConfigurationRequest) ApiCloudServiceChangeConfigurationRequest {
	r.cloudChangeConfigurationRequest = &cloudChangeConfigurationRequest
	return r
}

func (r ApiCloudServiceChangeConfigurationRequest) Execute() (*CloudChangeConfigurationResponse, *http.Response, error) {
	return r.ApiService.CloudServiceChangeConfigurationExecute(r)
}

/*
CloudServiceChangeConfiguration Method for CloudServiceChangeConfiguration

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiCloudServiceChangeConfigurationRequest
*/
func (a *CloudServiceApiService) CloudServiceChangeConfiguration(ctx context.Context, serviceId string) ApiCloudServiceChangeConfigurationRequest {
	return ApiCloudServiceChangeConfigurationRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return CloudChangeConfigurationResponse
func (a *CloudServiceApiService) CloudServiceChangeConfigurationExecute(r ApiCloudServiceChangeConfigurationRequest) (*CloudChangeConfigurationResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CloudChangeConfigurationResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudServiceApiService.CloudServiceChangeConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/{service_id}/configuration"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudChangeConfigurationRequest == nil {
		return localVarReturnValue, nil, reportError("cloudChangeConfigurationRequest is required and must be specified")
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
	localVarPostBody = r.cloudChangeConfigurationRequest
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

type ApiCloudServiceCreateRequest struct {
	ctx context.Context
	ApiService *CloudServiceApiService
	cloudCreateRequest *CloudCreateRequest
}

func (r ApiCloudServiceCreateRequest) CloudCreateRequest(cloudCreateRequest CloudCreateRequest) ApiCloudServiceCreateRequest {
	r.cloudCreateRequest = &cloudCreateRequest
	return r
}

func (r ApiCloudServiceCreateRequest) Execute() (*CloudCreateResponse, *http.Response, error) {
	return r.ApiService.CloudServiceCreateExecute(r)
}

/*
CloudServiceCreate Method for CloudServiceCreate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCloudServiceCreateRequest
*/
func (a *CloudServiceApiService) CloudServiceCreate(ctx context.Context) ApiCloudServiceCreateRequest {
	return ApiCloudServiceCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CloudCreateResponse
func (a *CloudServiceApiService) CloudServiceCreateExecute(r ApiCloudServiceCreateRequest) (*CloudCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CloudCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudServiceApiService.CloudServiceCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudCreateRequest == nil {
		return localVarReturnValue, nil, reportError("cloudCreateRequest is required and must be specified")
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
	localVarPostBody = r.cloudCreateRequest
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

type ApiCloudServiceGetRequest struct {
	ctx context.Context
	ApiService *CloudServiceApiService
	serviceId string
}

func (r ApiCloudServiceGetRequest) Execute() (*CloudGetResponse, *http.Response, error) {
	return r.ApiService.CloudServiceGetExecute(r)
}

/*
CloudServiceGet Method for CloudServiceGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiCloudServiceGetRequest
*/
func (a *CloudServiceApiService) CloudServiceGet(ctx context.Context, serviceId string) ApiCloudServiceGetRequest {
	return ApiCloudServiceGetRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return CloudGetResponse
func (a *CloudServiceApiService) CloudServiceGetExecute(r ApiCloudServiceGetRequest) (*CloudGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CloudGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudServiceApiService.CloudServiceGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/{service_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

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

type ApiCloudServiceGetConfigurationListRequest struct {
	ctx context.Context
	ApiService *CloudServiceApiService
}

func (r ApiCloudServiceGetConfigurationListRequest) Execute() (*CloudGetConfigurationListResponse, *http.Response, error) {
	return r.ApiService.CloudServiceGetConfigurationListExecute(r)
}

/*
CloudServiceGetConfigurationList Method for CloudServiceGetConfigurationList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCloudServiceGetConfigurationListRequest
*/
func (a *CloudServiceApiService) CloudServiceGetConfigurationList(ctx context.Context) ApiCloudServiceGetConfigurationListRequest {
	return ApiCloudServiceGetConfigurationListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CloudGetConfigurationListResponse
func (a *CloudServiceApiService) CloudServiceGetConfigurationListExecute(r ApiCloudServiceGetConfigurationListRequest) (*CloudGetConfigurationListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CloudGetConfigurationListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudServiceApiService.CloudServiceGetConfigurationList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/configuration"

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

type ApiCloudServiceGetListRequest struct {
	ctx context.Context
	ApiService *CloudServiceApiService
}

func (r ApiCloudServiceGetListRequest) Execute() (*CloudGetListResponse, *http.Response, error) {
	return r.ApiService.CloudServiceGetListExecute(r)
}

/*
CloudServiceGetList Method for CloudServiceGetList

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCloudServiceGetListRequest
*/
func (a *CloudServiceApiService) CloudServiceGetList(ctx context.Context) ApiCloudServiceGetListRequest {
	return ApiCloudServiceGetListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CloudGetListResponse
func (a *CloudServiceApiService) CloudServiceGetListExecute(r ApiCloudServiceGetListRequest) (*CloudGetListResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CloudGetListResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudServiceApiService.CloudServiceGetList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud"

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

type ApiCloudServiceRemoveRequest struct {
	ctx context.Context
	ApiService *CloudServiceApiService
	serviceId string
}

func (r ApiCloudServiceRemoveRequest) Execute() (*CloudRemoveResponse, *http.Response, error) {
	return r.ApiService.CloudServiceRemoveExecute(r)
}

/*
CloudServiceRemove Method for CloudServiceRemove

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiCloudServiceRemoveRequest
*/
func (a *CloudServiceApiService) CloudServiceRemove(ctx context.Context, serviceId string) ApiCloudServiceRemoveRequest {
	return ApiCloudServiceRemoveRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return CloudRemoveResponse
func (a *CloudServiceApiService) CloudServiceRemoveExecute(r ApiCloudServiceRemoveRequest) (*CloudRemoveResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CloudRemoveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudServiceApiService.CloudServiceRemove")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/{service_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

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

type ApiCloudServiceUpdateRequest struct {
	ctx context.Context
	ApiService *CloudServiceApiService
	serviceId string
	cloudUpdateRequest *CloudUpdateRequest
}

func (r ApiCloudServiceUpdateRequest) CloudUpdateRequest(cloudUpdateRequest CloudUpdateRequest) ApiCloudServiceUpdateRequest {
	r.cloudUpdateRequest = &cloudUpdateRequest
	return r
}

func (r ApiCloudServiceUpdateRequest) Execute() (*CloudUpdateResponse, *http.Response, error) {
	return r.ApiService.CloudServiceUpdateExecute(r)
}

/*
CloudServiceUpdate Method for CloudServiceUpdate

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiCloudServiceUpdateRequest
*/
func (a *CloudServiceApiService) CloudServiceUpdate(ctx context.Context, serviceId string) ApiCloudServiceUpdateRequest {
	return ApiCloudServiceUpdateRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return CloudUpdateResponse
func (a *CloudServiceApiService) CloudServiceUpdateExecute(r ApiCloudServiceUpdateRequest) (*CloudUpdateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CloudUpdateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CloudServiceApiService.CloudServiceUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/{service_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.cloudUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("cloudUpdateRequest is required and must be specified")
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
	localVarPostBody = r.cloudUpdateRequest
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
