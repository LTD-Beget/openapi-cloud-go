# S3ChangeCorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**S3** | Pointer to [**S3S3**](S3S3.md) |  | [optional] 
**Error** | Pointer to [**S3ChangeCorsResponseError**](S3ChangeCorsResponseError.md) |  | [optional] 

## Methods

### NewS3ChangeCorsResponse

`func NewS3ChangeCorsResponse() *S3ChangeCorsResponse`

NewS3ChangeCorsResponse instantiates a new S3ChangeCorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ChangeCorsResponseWithDefaults

`func NewS3ChangeCorsResponseWithDefaults() *S3ChangeCorsResponse`

NewS3ChangeCorsResponseWithDefaults instantiates a new S3ChangeCorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetS3

`func (o *S3ChangeCorsResponse) GetS3() S3S3`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *S3ChangeCorsResponse) GetS3Ok() (*S3S3, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *S3ChangeCorsResponse) SetS3(v S3S3)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *S3ChangeCorsResponse) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetError

`func (o *S3ChangeCorsResponse) GetError() S3ChangeCorsResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *S3ChangeCorsResponse) GetErrorOk() (*S3ChangeCorsResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *S3ChangeCorsResponse) SetError(v S3ChangeCorsResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *S3ChangeCorsResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


