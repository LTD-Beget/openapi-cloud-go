# S3ChangeDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**Fqdn** | Pointer to **string** |  | [optional] 

## Methods

### NewS3ChangeDomainRequest

`func NewS3ChangeDomainRequest() *S3ChangeDomainRequest`

NewS3ChangeDomainRequest instantiates a new S3ChangeDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ChangeDomainRequestWithDefaults

`func NewS3ChangeDomainRequestWithDefaults() *S3ChangeDomainRequest`

NewS3ChangeDomainRequestWithDefaults instantiates a new S3ChangeDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *S3ChangeDomainRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *S3ChangeDomainRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *S3ChangeDomainRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *S3ChangeDomainRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetFqdn

`func (o *S3ChangeDomainRequest) GetFqdn() string`

GetFqdn returns the Fqdn field if non-nil, zero value otherwise.

### GetFqdnOk

`func (o *S3ChangeDomainRequest) GetFqdnOk() (*string, bool)`

GetFqdnOk returns a tuple with the Fqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFqdn

`func (o *S3ChangeDomainRequest) SetFqdn(v string)`

SetFqdn sets Fqdn field to given value.

### HasFqdn

`func (o *S3ChangeDomainRequest) HasFqdn() bool`

HasFqdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


