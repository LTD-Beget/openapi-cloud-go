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


// S3ServiceApiService S3ServiceApi service
type S3ServiceApiService service

type ApiS3ServiceChangeAccessKeyRequest struct {
	ctx context.Context
	ApiService *S3ServiceApiService
	serviceId string
	s3ChangeAccessKeyRequest *S3ChangeAccessKeyRequest
}

func (r ApiS3ServiceChangeAccessKeyRequest) S3ChangeAccessKeyRequest(s3ChangeAccessKeyRequest S3ChangeAccessKeyRequest) ApiS3ServiceChangeAccessKeyRequest {
	r.s3ChangeAccessKeyRequest = &s3ChangeAccessKeyRequest
	return r
}

func (r ApiS3ServiceChangeAccessKeyRequest) Execute() (*S3ChangeAccessKeyResponse, *http.Response, error) {
	return r.ApiService.S3ServiceChangeAccessKeyExecute(r)
}

/*
S3ServiceChangeAccessKey Method for S3ServiceChangeAccessKey

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiS3ServiceChangeAccessKeyRequest
*/
func (a *S3ServiceApiService) S3ServiceChangeAccessKey(ctx context.Context, serviceId string) ApiS3ServiceChangeAccessKeyRequest {
	return ApiS3ServiceChangeAccessKeyRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return S3ChangeAccessKeyResponse
func (a *S3ServiceApiService) S3ServiceChangeAccessKeyExecute(r ApiS3ServiceChangeAccessKeyRequest) (*S3ChangeAccessKeyResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *S3ChangeAccessKeyResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "S3ServiceApiService.S3ServiceChangeAccessKey")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/s3/{service_id}/access-key"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.s3ChangeAccessKeyRequest == nil {
		return localVarReturnValue, nil, reportError("s3ChangeAccessKeyRequest is required and must be specified")
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
	localVarPostBody = r.s3ChangeAccessKeyRequest
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

type ApiS3ServiceChangeCorsRequest struct {
	ctx context.Context
	ApiService *S3ServiceApiService
	serviceId string
	s3ChangeCorsRequest *S3ChangeCorsRequest
}

func (r ApiS3ServiceChangeCorsRequest) S3ChangeCorsRequest(s3ChangeCorsRequest S3ChangeCorsRequest) ApiS3ServiceChangeCorsRequest {
	r.s3ChangeCorsRequest = &s3ChangeCorsRequest
	return r
}

func (r ApiS3ServiceChangeCorsRequest) Execute() (*S3ChangeCorsResponse, *http.Response, error) {
	return r.ApiService.S3ServiceChangeCorsExecute(r)
}

/*
S3ServiceChangeCors Method for S3ServiceChangeCors

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiS3ServiceChangeCorsRequest
*/
func (a *S3ServiceApiService) S3ServiceChangeCors(ctx context.Context, serviceId string) ApiS3ServiceChangeCorsRequest {
	return ApiS3ServiceChangeCorsRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return S3ChangeCorsResponse
func (a *S3ServiceApiService) S3ServiceChangeCorsExecute(r ApiS3ServiceChangeCorsRequest) (*S3ChangeCorsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *S3ChangeCorsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "S3ServiceApiService.S3ServiceChangeCors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/s3/{service_id}/cors"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.s3ChangeCorsRequest == nil {
		return localVarReturnValue, nil, reportError("s3ChangeCorsRequest is required and must be specified")
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
	localVarPostBody = r.s3ChangeCorsRequest
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

type ApiS3ServiceChangeDomainRequest struct {
	ctx context.Context
	ApiService *S3ServiceApiService
	serviceId string
	s3ChangeDomainRequest *S3ChangeDomainRequest
}

func (r ApiS3ServiceChangeDomainRequest) S3ChangeDomainRequest(s3ChangeDomainRequest S3ChangeDomainRequest) ApiS3ServiceChangeDomainRequest {
	r.s3ChangeDomainRequest = &s3ChangeDomainRequest
	return r
}

func (r ApiS3ServiceChangeDomainRequest) Execute() (*S3ChangeDomainResponse, *http.Response, error) {
	return r.ApiService.S3ServiceChangeDomainExecute(r)
}

/*
S3ServiceChangeDomain Method for S3ServiceChangeDomain

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiS3ServiceChangeDomainRequest
*/
func (a *S3ServiceApiService) S3ServiceChangeDomain(ctx context.Context, serviceId string) ApiS3ServiceChangeDomainRequest {
	return ApiS3ServiceChangeDomainRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return S3ChangeDomainResponse
func (a *S3ServiceApiService) S3ServiceChangeDomainExecute(r ApiS3ServiceChangeDomainRequest) (*S3ChangeDomainResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *S3ChangeDomainResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "S3ServiceApiService.S3ServiceChangeDomain")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/s3/{service_id}/domain"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.s3ChangeDomainRequest == nil {
		return localVarReturnValue, nil, reportError("s3ChangeDomainRequest is required and must be specified")
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
	localVarPostBody = r.s3ChangeDomainRequest
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

type ApiS3ServiceChangePublicRequest struct {
	ctx context.Context
	ApiService *S3ServiceApiService
	serviceId string
	s3ChangePublicRequest *S3ChangePublicRequest
}

func (r ApiS3ServiceChangePublicRequest) S3ChangePublicRequest(s3ChangePublicRequest S3ChangePublicRequest) ApiS3ServiceChangePublicRequest {
	r.s3ChangePublicRequest = &s3ChangePublicRequest
	return r
}

func (r ApiS3ServiceChangePublicRequest) Execute() (*S3ChangePublicResponse, *http.Response, error) {
	return r.ApiService.S3ServiceChangePublicExecute(r)
}

/*
S3ServiceChangePublic Method for S3ServiceChangePublic

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiS3ServiceChangePublicRequest
*/
func (a *S3ServiceApiService) S3ServiceChangePublic(ctx context.Context, serviceId string) ApiS3ServiceChangePublicRequest {
	return ApiS3ServiceChangePublicRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return S3ChangePublicResponse
func (a *S3ServiceApiService) S3ServiceChangePublicExecute(r ApiS3ServiceChangePublicRequest) (*S3ChangePublicResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *S3ChangePublicResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "S3ServiceApiService.S3ServiceChangePublic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/s3/{service_id}/public"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.s3ChangePublicRequest == nil {
		return localVarReturnValue, nil, reportError("s3ChangePublicRequest is required and must be specified")
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
	localVarPostBody = r.s3ChangePublicRequest
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

type ApiS3ServiceEnableFtpRequest struct {
	ctx context.Context
	ApiService *S3ServiceApiService
	serviceId string
	s3EnableFtpRequest *S3EnableFtpRequest
}

func (r ApiS3ServiceEnableFtpRequest) S3EnableFtpRequest(s3EnableFtpRequest S3EnableFtpRequest) ApiS3ServiceEnableFtpRequest {
	r.s3EnableFtpRequest = &s3EnableFtpRequest
	return r
}

func (r ApiS3ServiceEnableFtpRequest) Execute() (*S3EnableFtpResponse, *http.Response, error) {
	return r.ApiService.S3ServiceEnableFtpExecute(r)
}

/*
S3ServiceEnableFtp Method for S3ServiceEnableFtp

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param serviceId
 @return ApiS3ServiceEnableFtpRequest
*/
func (a *S3ServiceApiService) S3ServiceEnableFtp(ctx context.Context, serviceId string) ApiS3ServiceEnableFtpRequest {
	return ApiS3ServiceEnableFtpRequest{
		ApiService: a,
		ctx: ctx,
		serviceId: serviceId,
	}
}

// Execute executes the request
//  @return S3EnableFtpResponse
func (a *S3ServiceApiService) S3ServiceEnableFtpExecute(r ApiS3ServiceEnableFtpRequest) (*S3EnableFtpResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *S3EnableFtpResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "S3ServiceApiService.S3ServiceEnableFtp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/s3/{service_id}/enable-ftp"
	localVarPath = strings.Replace(localVarPath, "{"+"service_id"+"}", url.PathEscape(parameterToString(r.serviceId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.s3EnableFtpRequest == nil {
		return localVarReturnValue, nil, reportError("s3EnableFtpRequest is required and must be specified")
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
	localVarPostBody = r.s3EnableFtpRequest
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

type ApiS3ServiceGetPrefixRequest struct {
	ctx context.Context
	ApiService *S3ServiceApiService
}

func (r ApiS3ServiceGetPrefixRequest) Execute() (*S3GetPrefixResponse, *http.Response, error) {
	return r.ApiService.S3ServiceGetPrefixExecute(r)
}

/*
S3ServiceGetPrefix Method for S3ServiceGetPrefix

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiS3ServiceGetPrefixRequest
*/
func (a *S3ServiceApiService) S3ServiceGetPrefix(ctx context.Context) ApiS3ServiceGetPrefixRequest {
	return ApiS3ServiceGetPrefixRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return S3GetPrefixResponse
func (a *S3ServiceApiService) S3ServiceGetPrefixExecute(r ApiS3ServiceGetPrefixRequest) (*S3GetPrefixResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *S3GetPrefixResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "S3ServiceApiService.S3ServiceGetPrefix")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/s3/prefix"

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

type ApiS3ServiceGetPriceRequest struct {
	ctx context.Context
	ApiService *S3ServiceApiService
}

func (r ApiS3ServiceGetPriceRequest) Execute() (*S3GetPriceResponse, *http.Response, error) {
	return r.ApiService.S3ServiceGetPriceExecute(r)
}

/*
S3ServiceGetPrice Method for S3ServiceGetPrice

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiS3ServiceGetPriceRequest
*/
func (a *S3ServiceApiService) S3ServiceGetPrice(ctx context.Context) ApiS3ServiceGetPriceRequest {
	return ApiS3ServiceGetPriceRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return S3GetPriceResponse
func (a *S3ServiceApiService) S3ServiceGetPriceExecute(r ApiS3ServiceGetPriceRequest) (*S3GetPriceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *S3GetPriceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "S3ServiceApiService.S3ServiceGetPrice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/s3/price"

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

type ApiS3ServiceGetQuotaRequest struct {
	ctx context.Context
	ApiService *S3ServiceApiService
}

func (r ApiS3ServiceGetQuotaRequest) Execute() (*S3GetQuotaResponse, *http.Response, error) {
	return r.ApiService.S3ServiceGetQuotaExecute(r)
}

/*
S3ServiceGetQuota Method for S3ServiceGetQuota

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiS3ServiceGetQuotaRequest
*/
func (a *S3ServiceApiService) S3ServiceGetQuota(ctx context.Context) ApiS3ServiceGetQuotaRequest {
	return ApiS3ServiceGetQuotaRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return S3GetQuotaResponse
func (a *S3ServiceApiService) S3ServiceGetQuotaExecute(r ApiS3ServiceGetQuotaRequest) (*S3GetQuotaResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *S3GetQuotaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "S3ServiceApiService.S3ServiceGetQuota")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/cloud/s3/quota"

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
