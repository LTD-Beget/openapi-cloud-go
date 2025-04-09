# S3EnableFtpResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3** | Pointer to [**S3S3**](S3S3.md) |  | [optional] 
**Error** | Pointer to [**S3EnableFtpResponseError**](S3EnableFtpResponseError.md) |  | [optional] 

## Methods

### NewS3EnableFtpResponse

`func NewS3EnableFtpResponse() *S3EnableFtpResponse`

NewS3EnableFtpResponse instantiates a new S3EnableFtpResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3EnableFtpResponseWithDefaults

`func NewS3EnableFtpResponseWithDefaults() *S3EnableFtpResponse`

NewS3EnableFtpResponseWithDefaults instantiates a new S3EnableFtpResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3

`func (o *S3EnableFtpResponse) GetS3() S3S3`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *S3EnableFtpResponse) GetS3Ok() (*S3S3, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *S3EnableFtpResponse) SetS3(v S3S3)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *S3EnableFtpResponse) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetError

`func (o *S3EnableFtpResponse) GetError() S3EnableFtpResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *S3EnableFtpResponse) GetErrorOk() (*S3EnableFtpResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *S3EnableFtpResponse) SetError(v S3EnableFtpResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *S3EnableFtpResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


