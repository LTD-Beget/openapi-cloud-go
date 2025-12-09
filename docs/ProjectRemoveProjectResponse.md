# ProjectRemoveProjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **map[string]interface{}** |  | [optional] 
**Error** | Pointer to [**ProjectRemoveProjectResponseError**](ProjectRemoveProjectResponseError.md) |  | [optional] 

## Methods

### NewProjectRemoveProjectResponse

`func NewProjectRemoveProjectResponse() *ProjectRemoveProjectResponse`

NewProjectRemoveProjectResponse instantiates a new ProjectRemoveProjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectRemoveProjectResponseWithDefaults

`func NewProjectRemoveProjectResponseWithDefaults() *ProjectRemoveProjectResponse`

NewProjectRemoveProjectResponseWithDefaults instantiates a new ProjectRemoveProjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ProjectRemoveProjectResponse) GetSuccess() map[string]interface{}`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ProjectRemoveProjectResponse) GetSuccessOk() (*map[string]interface{}, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ProjectRemoveProjectResponse) SetSuccess(v map[string]interface{})`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ProjectRemoveProjectResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetError

`func (o *ProjectRemoveProjectResponse) GetError() ProjectRemoveProjectResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ProjectRemoveProjectResponse) GetErrorOk() (*ProjectRemoveProjectResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ProjectRemoveProjectResponse) SetError(v ProjectRemoveProjectResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ProjectRemoveProjectResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


