# CloudGetConfigurationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**[]CloudServiceConfiguration**](CloudServiceConfiguration.md) |  | [optional] 

## Methods

### NewCloudGetConfigurationListResponse

`func NewCloudGetConfigurationListResponse() *CloudGetConfigurationListResponse`

NewCloudGetConfigurationListResponse instantiates a new CloudGetConfigurationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGetConfigurationListResponseWithDefaults

`func NewCloudGetConfigurationListResponseWithDefaults() *CloudGetConfigurationListResponse`

NewCloudGetConfigurationListResponseWithDefaults instantiates a new CloudGetConfigurationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *CloudGetConfigurationListResponse) GetConfiguration() []CloudServiceConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CloudGetConfigurationListResponse) GetConfigurationOk() (*[]CloudServiceConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CloudGetConfigurationListResponse) SetConfiguration(v []CloudServiceConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CloudGetConfigurationListResponse) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


