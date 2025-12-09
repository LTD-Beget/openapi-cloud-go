# CloudBindProjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**CloudService**](CloudService.md) |  | [optional] 
**Error** | Pointer to [**CloudBindProjectResponseError**](CloudBindProjectResponseError.md) |  | [optional] 

## Methods

### NewCloudBindProjectResponse

`func NewCloudBindProjectResponse() *CloudBindProjectResponse`

NewCloudBindProjectResponse instantiates a new CloudBindProjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudBindProjectResponseWithDefaults

`func NewCloudBindProjectResponseWithDefaults() *CloudBindProjectResponse`

NewCloudBindProjectResponseWithDefaults instantiates a new CloudBindProjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CloudBindProjectResponse) GetService() CloudService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudBindProjectResponse) GetServiceOk() (*CloudService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudBindProjectResponse) SetService(v CloudService)`

SetService sets Service field to given value.

### HasService

`func (o *CloudBindProjectResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetError

`func (o *CloudBindProjectResponse) GetError() CloudBindProjectResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudBindProjectResponse) GetErrorOk() (*CloudBindProjectResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudBindProjectResponse) SetError(v CloudBindProjectResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *CloudBindProjectResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


