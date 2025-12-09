# ProjectCreateProjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**ProjectProject**](ProjectProject.md) |  | [optional] 
**Error** | Pointer to [**ProjectCreateProjectResponseError**](ProjectCreateProjectResponseError.md) |  | [optional] 

## Methods

### NewProjectCreateProjectResponse

`func NewProjectCreateProjectResponse() *ProjectCreateProjectResponse`

NewProjectCreateProjectResponse instantiates a new ProjectCreateProjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateProjectResponseWithDefaults

`func NewProjectCreateProjectResponseWithDefaults() *ProjectCreateProjectResponse`

NewProjectCreateProjectResponseWithDefaults instantiates a new ProjectCreateProjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ProjectCreateProjectResponse) GetProject() ProjectProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectCreateProjectResponse) GetProjectOk() (*ProjectProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectCreateProjectResponse) SetProject(v ProjectProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectCreateProjectResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetError

`func (o *ProjectCreateProjectResponse) GetError() ProjectCreateProjectResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ProjectCreateProjectResponse) GetErrorOk() (*ProjectCreateProjectResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ProjectCreateProjectResponse) SetError(v ProjectCreateProjectResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ProjectCreateProjectResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


