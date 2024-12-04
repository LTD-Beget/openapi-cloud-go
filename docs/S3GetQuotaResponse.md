# S3GetQuotaResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsedSize** | Pointer to **int32** |  | [optional] 
**AvailableSize** | Pointer to **int32** |  | [optional] 

## Methods

### NewS3GetQuotaResponse

`func NewS3GetQuotaResponse() *S3GetQuotaResponse`

NewS3GetQuotaResponse instantiates a new S3GetQuotaResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3GetQuotaResponseWithDefaults

`func NewS3GetQuotaResponseWithDefaults() *S3GetQuotaResponse`

NewS3GetQuotaResponseWithDefaults instantiates a new S3GetQuotaResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsedSize

`func (o *S3GetQuotaResponse) GetUsedSize() int32`

GetUsedSize returns the UsedSize field if non-nil, zero value otherwise.

### GetUsedSizeOk

`func (o *S3GetQuotaResponse) GetUsedSizeOk() (*int32, bool)`

GetUsedSizeOk returns a tuple with the UsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedSize

`func (o *S3GetQuotaResponse) SetUsedSize(v int32)`

SetUsedSize sets UsedSize field to given value.

### HasUsedSize

`func (o *S3GetQuotaResponse) HasUsedSize() bool`

HasUsedSize returns a boolean if a field has been set.

### GetAvailableSize

`func (o *S3GetQuotaResponse) GetAvailableSize() int32`

GetAvailableSize returns the AvailableSize field if non-nil, zero value otherwise.

### GetAvailableSizeOk

`func (o *S3GetQuotaResponse) GetAvailableSizeOk() (*int32, bool)`

GetAvailableSizeOk returns a tuple with the AvailableSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableSize

`func (o *S3GetQuotaResponse) SetAvailableSize(v int32)`

SetAvailableSize sets AvailableSize field to given value.

### HasAvailableSize

`func (o *S3GetQuotaResponse) HasAvailableSize() bool`

HasAvailableSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


