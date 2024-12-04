# S3ChangeCorsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**Cors** | Pointer to [**[]S3Cors**](S3Cors.md) |  | [optional] 

## Methods

### NewS3ChangeCorsRequest

`func NewS3ChangeCorsRequest() *S3ChangeCorsRequest`

NewS3ChangeCorsRequest instantiates a new S3ChangeCorsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ChangeCorsRequestWithDefaults

`func NewS3ChangeCorsRequestWithDefaults() *S3ChangeCorsRequest`

NewS3ChangeCorsRequestWithDefaults instantiates a new S3ChangeCorsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *S3ChangeCorsRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *S3ChangeCorsRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *S3ChangeCorsRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *S3ChangeCorsRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetCors

`func (o *S3ChangeCorsRequest) GetCors() []S3Cors`

GetCors returns the Cors field if non-nil, zero value otherwise.

### GetCorsOk

`func (o *S3ChangeCorsRequest) GetCorsOk() (*[]S3Cors, bool)`

GetCorsOk returns a tuple with the Cors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCors

`func (o *S3ChangeCorsRequest) SetCors(v []S3Cors)`

SetCors sets Cors field to given value.

### HasCors

`func (o *S3ChangeCorsRequest) HasCors() bool`

HasCors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


