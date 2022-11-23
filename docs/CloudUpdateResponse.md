# CloudUpdateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**CloudService**](CloudService.md) |  | [optional] 
**Error** | Pointer to [**CloudUpdateResponseError**](CloudUpdateResponseError.md) |  | [optional] 

## Methods

### NewCloudUpdateResponse

`func NewCloudUpdateResponse() *CloudUpdateResponse`

NewCloudUpdateResponse instantiates a new CloudUpdateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudUpdateResponseWithDefaults

`func NewCloudUpdateResponseWithDefaults() *CloudUpdateResponse`

NewCloudUpdateResponseWithDefaults instantiates a new CloudUpdateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CloudUpdateResponse) GetService() CloudService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudUpdateResponse) GetServiceOk() (*CloudService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudUpdateResponse) SetService(v CloudService)`

SetService sets Service field to given value.

### HasService

`func (o *CloudUpdateResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetError

`func (o *CloudUpdateResponse) GetError() CloudUpdateResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudUpdateResponse) GetErrorOk() (*CloudUpdateResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudUpdateResponse) SetError(v CloudUpdateResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *CloudUpdateResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


