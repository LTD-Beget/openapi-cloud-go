# S3ChangePublicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3** | Pointer to [**S3S3**](S3S3.md) |  | [optional] 
**Error** | Pointer to [**S3ChangePublicResponseError**](S3ChangePublicResponseError.md) |  | [optional] 

## Methods

### NewS3ChangePublicResponse

`func NewS3ChangePublicResponse() *S3ChangePublicResponse`

NewS3ChangePublicResponse instantiates a new S3ChangePublicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ChangePublicResponseWithDefaults

`func NewS3ChangePublicResponseWithDefaults() *S3ChangePublicResponse`

NewS3ChangePublicResponseWithDefaults instantiates a new S3ChangePublicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3

`func (o *S3ChangePublicResponse) GetS3() S3S3`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *S3ChangePublicResponse) GetS3Ok() (*S3S3, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *S3ChangePublicResponse) SetS3(v S3S3)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *S3ChangePublicResponse) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetError

`func (o *S3ChangePublicResponse) GetError() S3ChangePublicResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *S3ChangePublicResponse) GetErrorOk() (*S3ChangePublicResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *S3ChangePublicResponse) SetError(v S3ChangePublicResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *S3ChangePublicResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


