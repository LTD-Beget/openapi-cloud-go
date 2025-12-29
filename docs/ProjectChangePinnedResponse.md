# ProjectChangePinnedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**ProjectProject**](ProjectProject.md) |  | [optional] 
**Error** | Pointer to [**ProjectChangePinnedResponseError**](ProjectChangePinnedResponseError.md) |  | [optional] 

## Methods

### NewProjectChangePinnedResponse

`func NewProjectChangePinnedResponse() *ProjectChangePinnedResponse`

NewProjectChangePinnedResponse instantiates a new ProjectChangePinnedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectChangePinnedResponseWithDefaults

`func NewProjectChangePinnedResponseWithDefaults() *ProjectChangePinnedResponse`

NewProjectChangePinnedResponseWithDefaults instantiates a new ProjectChangePinnedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ProjectChangePinnedResponse) GetProject() ProjectProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectChangePinnedResponse) GetProjectOk() (*ProjectProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectChangePinnedResponse) SetProject(v ProjectProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectChangePinnedResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetError

`func (o *ProjectChangePinnedResponse) GetError() ProjectChangePinnedResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ProjectChangePinnedResponse) GetErrorOk() (*ProjectChangePinnedResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ProjectChangePinnedResponse) SetError(v ProjectChangePinnedResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ProjectChangePinnedResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


