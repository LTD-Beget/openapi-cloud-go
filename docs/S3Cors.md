# S3Cors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**HttpMethod** | Pointer to **[]string** |  | [optional] 
**AccessControlAllowHeaders** | Pointer to **[]string** |  | [optional] 
**AccessControlExposeHeaders** | Pointer to **[]string** |  | [optional] 
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

### GetAccessControlAllowHeaders

`func (o *S3Cors) GetAccessControlAllowHeaders() []string`

GetAccessControlAllowHeaders returns the AccessControlAllowHeaders field if non-nil, zero value otherwise.

### GetAccessControlAllowHeadersOk

`func (o *S3Cors) GetAccessControlAllowHeadersOk() (*[]string, bool)`

GetAccessControlAllowHeadersOk returns a tuple with the AccessControlAllowHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlAllowHeaders

`func (o *S3Cors) SetAccessControlAllowHeaders(v []string)`

SetAccessControlAllowHeaders sets AccessControlAllowHeaders field to given value.

### HasAccessControlAllowHeaders

`func (o *S3Cors) HasAccessControlAllowHeaders() bool`

HasAccessControlAllowHeaders returns a boolean if a field has been set.

### GetAccessControlExposeHeaders

`func (o *S3Cors) GetAccessControlExposeHeaders() []string`

GetAccessControlExposeHeaders returns the AccessControlExposeHeaders field if non-nil, zero value otherwise.

### GetAccessControlExposeHeadersOk

`func (o *S3Cors) GetAccessControlExposeHeadersOk() (*[]string, bool)`

GetAccessControlExposeHeadersOk returns a tuple with the AccessControlExposeHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControlExposeHeaders

`func (o *S3Cors) SetAccessControlExposeHeaders(v []string)`

SetAccessControlExposeHeaders sets AccessControlExposeHeaders field to given value.

### HasAccessControlExposeHeaders

`func (o *S3Cors) HasAccessControlExposeHeaders() bool`

HasAccessControlExposeHeaders returns a boolean if a field has been set.

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


