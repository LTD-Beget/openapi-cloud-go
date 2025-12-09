# CloudGetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | Pointer to [**[]CloudService**](CloudService.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudGetListResponse

`func NewCloudGetListResponse() *CloudGetListResponse`

NewCloudGetListResponse instantiates a new CloudGetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGetListResponseWithDefaults

`func NewCloudGetListResponseWithDefaults() *CloudGetListResponse`

NewCloudGetListResponseWithDefaults instantiates a new CloudGetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *CloudGetListResponse) GetService() []CloudService`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *CloudGetListResponse) GetServiceOk() (*[]CloudService, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *CloudGetListResponse) SetService(v []CloudService)`

SetService sets Service field to given value.

### HasService

`func (o *CloudGetListResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetTotalCount

`func (o *CloudGetListResponse) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *CloudGetListResponse) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *CloudGetListResponse) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *CloudGetListResponse) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


