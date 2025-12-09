# ProjectGetProjectListResponseProjectPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]ProjectProject**](ProjectProject.md) |  | [optional] 
**TotalCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewProjectGetProjectListResponseProjectPage

`func NewProjectGetProjectListResponseProjectPage() *ProjectGetProjectListResponseProjectPage`

NewProjectGetProjectListResponseProjectPage instantiates a new ProjectGetProjectListResponseProjectPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectGetProjectListResponseProjectPageWithDefaults

`func NewProjectGetProjectListResponseProjectPageWithDefaults() *ProjectGetProjectListResponseProjectPage`

NewProjectGetProjectListResponseProjectPageWithDefaults instantiates a new ProjectGetProjectListResponseProjectPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *ProjectGetProjectListResponseProjectPage) GetProjects() []ProjectProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectGetProjectListResponseProjectPage) GetProjectsOk() (*[]ProjectProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectGetProjectListResponseProjectPage) SetProjects(v []ProjectProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProjectGetProjectListResponseProjectPage) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetTotalCount

`func (o *ProjectGetProjectListResponseProjectPage) GetTotalCount() int32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ProjectGetProjectListResponseProjectPage) GetTotalCountOk() (*int32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ProjectGetProjectListResponseProjectPage) SetTotalCount(v int32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *ProjectGetProjectListResponseProjectPage) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


