# ProjectUpdateProjectResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to [**ProjectProject**](ProjectProject.md) |  | [optional] 
**Error** | Pointer to [**ProjectUpdateProjectResponseError**](ProjectUpdateProjectResponseError.md) |  | [optional] 

## Methods

### NewProjectUpdateProjectResponse

`func NewProjectUpdateProjectResponse() *ProjectUpdateProjectResponse`

NewProjectUpdateProjectResponse instantiates a new ProjectUpdateProjectResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUpdateProjectResponseWithDefaults

`func NewProjectUpdateProjectResponseWithDefaults() *ProjectUpdateProjectResponse`

NewProjectUpdateProjectResponseWithDefaults instantiates a new ProjectUpdateProjectResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *ProjectUpdateProjectResponse) GetProject() ProjectProject`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *ProjectUpdateProjectResponse) GetProjectOk() (*ProjectProject, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *ProjectUpdateProjectResponse) SetProject(v ProjectProject)`

SetProject sets Project field to given value.

### HasProject

`func (o *ProjectUpdateProjectResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### GetError

`func (o *ProjectUpdateProjectResponse) GetError() ProjectUpdateProjectResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ProjectUpdateProjectResponse) GetErrorOk() (*ProjectUpdateProjectResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ProjectUpdateProjectResponse) SetError(v ProjectUpdateProjectResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *ProjectUpdateProjectResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


