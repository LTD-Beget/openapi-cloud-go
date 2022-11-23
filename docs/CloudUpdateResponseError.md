# CloudUpdateResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudUpdateResponseError

`func NewCloudUpdateResponseError() *CloudUpdateResponseError`

NewCloudUpdateResponseError instantiates a new CloudUpdateResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUpdateResponseErrorWithDefaults

`func NewCloudUpdateResponseErrorWithDefaults() *CloudUpdateResponseError`

NewCloudUpdateResponseErrorWithDefaults instantiates a new CloudUpdateResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CloudUpdateResponseError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudUpdateResponseError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudUpdateResponseError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CloudUpdateResponseError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *CloudUpdateResponseError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CloudUpdateResponseError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CloudUpdateResponseError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CloudUpdateResponseError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


