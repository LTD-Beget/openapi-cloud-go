# S3ChangePublicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | Pointer to **string** |  | [optional] 
**Public** | Pointer to **bool** |  | [optional] 

## Methods

### NewS3ChangePublicRequest

`func NewS3ChangePublicRequest() *S3ChangePublicRequest`

NewS3ChangePublicRequest instantiates a new S3ChangePublicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ChangePublicRequestWithDefaults

`func NewS3ChangePublicRequestWithDefaults() *S3ChangePublicRequest`

NewS3ChangePublicRequestWithDefaults instantiates a new S3ChangePublicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *S3ChangePublicRequest) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *S3ChangePublicRequest) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *S3ChangePublicRequest) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.

### HasServiceId

`func (o *S3ChangePublicRequest) HasServiceId() bool`

HasServiceId returns a boolean if a field has been set.

### GetPublic

`func (o *S3ChangePublicRequest) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *S3ChangePublicRequest) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *S3ChangePublicRequest) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *S3ChangePublicRequest) HasPublic() bool`

HasPublic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


