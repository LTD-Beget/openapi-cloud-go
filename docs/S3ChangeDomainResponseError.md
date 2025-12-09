# S3ChangeDomainResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**DomainAlreadyUsedDetails** | Pointer to [**S3ChangeDomainResponseErrorDomainAlreadyUsedDetails**](S3ChangeDomainResponseErrorDomainAlreadyUsedDetails.md) |  | [optional] 

## Methods

### NewS3ChangeDomainResponseError

`func NewS3ChangeDomainResponseError() *S3ChangeDomainResponseError`

NewS3ChangeDomainResponseError instantiates a new S3ChangeDomainResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ChangeDomainResponseErrorWithDefaults

`func NewS3ChangeDomainResponseErrorWithDefaults() *S3ChangeDomainResponseError`

NewS3ChangeDomainResponseErrorWithDefaults instantiates a new S3ChangeDomainResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *S3ChangeDomainResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *S3ChangeDomainResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *S3ChangeDomainResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *S3ChangeDomainResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *S3ChangeDomainResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *S3ChangeDomainResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *S3ChangeDomainResponseError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *S3ChangeDomainResponseError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDomainAlreadyUsedDetails

`func (o *S3ChangeDomainResponseError) GetDomainAlreadyUsedDetails() S3ChangeDomainResponseErrorDomainAlreadyUsedDetails`

GetDomainAlreadyUsedDetails returns the DomainAlreadyUsedDetails field if non-nil, zero value otherwise.

### GetDomainAlreadyUsedDetailsOk

`func (o *S3ChangeDomainResponseError) GetDomainAlreadyUsedDetailsOk() (*S3ChangeDomainResponseErrorDomainAlreadyUsedDetails, bool)`

GetDomainAlreadyUsedDetailsOk returns a tuple with the DomainAlreadyUsedDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainAlreadyUsedDetails

`func (o *S3ChangeDomainResponseError) SetDomainAlreadyUsedDetails(v S3ChangeDomainResponseErrorDomainAlreadyUsedDetails)`

SetDomainAlreadyUsedDetails sets DomainAlreadyUsedDetails field to given value.

### HasDomainAlreadyUsedDetails

`func (o *S3ChangeDomainResponseError) HasDomainAlreadyUsedDetails() bool`

HasDomainAlreadyUsedDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


