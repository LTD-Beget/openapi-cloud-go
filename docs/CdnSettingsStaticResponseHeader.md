# CdnSettingsStaticResponseHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**Header** | Pointer to [**[]CdnSettingsStaticResponseHeaderHeader**](CdnSettingsStaticResponseHeaderHeader.md) |  | [optional] 

## Methods

### NewCdnSettingsStaticResponseHeader

`func NewCdnSettingsStaticResponseHeader() *CdnSettingsStaticResponseHeader`

NewCdnSettingsStaticResponseHeader instantiates a new CdnSettingsStaticResponseHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnSettingsStaticResponseHeaderWithDefaults

`func NewCdnSettingsStaticResponseHeaderWithDefaults() *CdnSettingsStaticResponseHeader`

NewCdnSettingsStaticResponseHeaderWithDefaults instantiates a new CdnSettingsStaticResponseHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *CdnSettingsStaticResponseHeader) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CdnSettingsStaticResponseHeader) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CdnSettingsStaticResponseHeader) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CdnSettingsStaticResponseHeader) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetHeader

`func (o *CdnSettingsStaticResponseHeader) GetHeader() []CdnSettingsStaticResponseHeaderHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *CdnSettingsStaticResponseHeader) GetHeaderOk() (*[]CdnSettingsStaticResponseHeaderHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *CdnSettingsStaticResponseHeader) SetHeader(v []CdnSettingsStaticResponseHeaderHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *CdnSettingsStaticResponseHeader) HasHeader() bool`

HasHeader returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


