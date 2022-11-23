# CloudChangeConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**CloudService**](CloudService.md) |  | [optional] 
**Error** | Pointer to [**CloudChangeConfigurationResponseError**](CloudChangeConfigurationResponseError.md) |  | [optional] 

## Methods

### NewCloudChangeConfigurationResponse

`func NewCloudChangeConfigurationResponse() *CloudChangeConfigurationResponse`

NewCloudChangeConfigurationResponse instantiates a new CloudChangeConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudChangeConfigurationResponseWithDefaults

`func NewCloudChangeConfigurationResponseWithDefaults() *CloudChangeConfigurationResponse`

NewCloudChangeConfigurationResponseWithDefaults instantiates a new CloudChangeConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CloudChangeConfigurationResponse) GetService() CloudService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudChangeConfigurationResponse) GetServiceOk() (*CloudService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudChangeConfigurationResponse) SetService(v CloudService)`

SetService sets Service field to given value.

### HasService

`func (o *CloudChangeConfigurationResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetError

`func (o *CloudChangeConfigurationResponse) GetError() CloudChangeConfigurationResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CloudChangeConfigurationResponse) GetErrorOk() (*CloudChangeConfigurationResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CloudChangeConfigurationResponse) SetError(v CloudChangeConfigurationResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *CloudChangeConfigurationResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


