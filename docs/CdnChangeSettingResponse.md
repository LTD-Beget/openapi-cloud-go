# CdnChangeSettingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdn** | Pointer to [**CdnCdn**](CdnCdn.md) |  | [optional] 
**Error** | Pointer to [**CdnChangeSettingResponseError**](CdnChangeSettingResponseError.md) |  | [optional] 

## Methods

### NewCdnChangeSettingResponse

`func NewCdnChangeSettingResponse() *CdnChangeSettingResponse`

NewCdnChangeSettingResponse instantiates a new CdnChangeSettingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnChangeSettingResponseWithDefaults

`func NewCdnChangeSettingResponseWithDefaults() *CdnChangeSettingResponse`

NewCdnChangeSettingResponseWithDefaults instantiates a new CdnChangeSettingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdn

`func (o *CdnChangeSettingResponse) GetCdn() CdnCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *CdnChangeSettingResponse) GetCdnOk() (*CdnCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *CdnChangeSettingResponse) SetCdn(v CdnCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *CdnChangeSettingResponse) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetError

`func (o *CdnChangeSettingResponse) GetError() CdnChangeSettingResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CdnChangeSettingResponse) GetErrorOk() (*CdnChangeSettingResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CdnChangeSettingResponse) SetError(v CdnChangeSettingResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *CdnChangeSettingResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


