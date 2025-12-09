# ProjectGetProjectListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to [**ProjectGetProjectListResponseProjectPage**](ProjectGetProjectListResponseProjectPage.md) |  | [optional] 
**Error** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewProjectGetProjectListResponse

`func NewProjectGetProjectListResponse() *ProjectGetProjectListResponse`

NewProjectGetProjectListResponse instantiates a new ProjectGetProjectListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectGetProjectListResponseWithDefaults

`func NewProjectGetProjectListResponseWithDefaults() *ProjectGetProjectListResponse`

NewProjectGetProjectListResponseWithDefaults instantiates a new ProjectGetProjectListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ProjectGetProjectListResponse) GetPage() ProjectGetProjectListResponseProjectPage`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ProjectGetProjectListResponse) GetPageOk() (*ProjectGetProjectListResponseProjectPage, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ProjectGetProjectListResponse) SetPage(v ProjectGetProjectListResponseProjectPage)`

SetPage sets Page field to given value.

### HasPage

`func (o *ProjectGetProjectListResponse) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetError

`func (o *ProjectGetProjectListResponse) GetError() map[string]interface{}`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ProjectGetProjectListResponse) GetErrorOk() (*map[string]interface{}, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ProjectGetProjectListResponse) SetError(v map[string]interface{})`

SetError sets Error field to given value.

### HasError

`func (o *ProjectGetProjectListResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


