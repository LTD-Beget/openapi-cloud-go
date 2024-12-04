# S3Cors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**HttpMethod** | Pointer to **[]string** |  | [optional] 
**HttpHeader** | Pointer to **[]string** |  | [optional] 
**CacheTtl** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 

## Methods

### NewS3Cors

`func NewS3Cors() *S3Cors`

NewS3Cors instantiates a new S3Cors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3CorsWithDefaults

`func NewS3CorsWithDefaults() *S3Cors`

NewS3CorsWithDefaults instantiates a new S3Cors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *S3Cors) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *S3Cors) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *S3Cors) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *S3Cors) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHttpMethod

`func (o *S3Cors) GetHttpMethod() []string`

GetHttpMethod returns the HttpMethod field if non-nil, zero value otherwise.

### GetHttpMethodOk

`func (o *S3Cors) GetHttpMethodOk() (*[]string, bool)`

GetHttpMethodOk returns a tuple with the HttpMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMethod

`func (o *S3Cors) SetHttpMethod(v []string)`

SetHttpMethod sets HttpMethod field to given value.

### HasHttpMethod

`func (o *S3Cors) HasHttpMethod() bool`

HasHttpMethod returns a boolean if a field has been set.

### GetHttpHeader

`func (o *S3Cors) GetHttpHeader() []string`

GetHttpHeader returns the HttpHeader field if non-nil, zero value otherwise.

### GetHttpHeaderOk

`func (o *S3Cors) GetHttpHeaderOk() (*[]string, bool)`

GetHttpHeaderOk returns a tuple with the HttpHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpHeader

`func (o *S3Cors) SetHttpHeader(v []string)`

SetHttpHeader sets HttpHeader field to given value.

### HasHttpHeader

`func (o *S3Cors) HasHttpHeader() bool`

HasHttpHeader returns a boolean if a field has been set.

### GetCacheTtl

`func (o *S3Cors) GetCacheTtl() string`

GetCacheTtl returns the CacheTtl field if non-nil, zero value otherwise.

### GetCacheTtlOk

`func (o *S3Cors) GetCacheTtlOk() (*string, bool)`

GetCacheTtlOk returns a tuple with the CacheTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheTtl

`func (o *S3Cors) SetCacheTtl(v string)`

SetCacheTtl sets CacheTtl field to given value.

### HasCacheTtl

`func (o *S3Cors) HasCacheTtl() bool`

HasCacheTtl returns a boolean if a field has been set.

### GetFqdn

`func (o *S3Cors) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *S3Cors) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *S3Cors) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *S3Cors) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


