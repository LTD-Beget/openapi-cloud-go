# S3EnableFtpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 

## Methods

### NewS3EnableFtpRequest

`func NewS3EnableFtpRequest() *S3EnableFtpRequest`

NewS3EnableFtpRequest instantiates a new S3EnableFtpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3EnableFtpRequestWithDefaults

`func NewS3EnableFtpRequestWithDefaults() *S3EnableFtpRequest`

NewS3EnableFtpRequestWithDefaults instantiates a new S3EnableFtpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *S3EnableFtpRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *S3EnableFtpRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *S3EnableFtpRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *S3EnableFtpRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetEnable

`func (o *S3EnableFtpRequest) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *S3EnableFtpRequest) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *S3EnableFtpRequest) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *S3EnableFtpRequest) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


