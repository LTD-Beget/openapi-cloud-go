# CdnGetRegionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CountryCode** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**[]CdnGetRegionsResponseRegion**](CdnGetRegionsResponseRegion.md) |  | [optional] 

## Methods

### NewCdnGetRegionsResponse

`func NewCdnGetRegionsResponse() *CdnGetRegionsResponse`

NewCdnGetRegionsResponse instantiates a new CdnGetRegionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnGetRegionsResponseWithDefaults

`func NewCdnGetRegionsResponseWithDefaults() *CdnGetRegionsResponse`

NewCdnGetRegionsResponseWithDefaults instantiates a new CdnGetRegionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountryCode

`func (o *CdnGetRegionsResponse) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *CdnGetRegionsResponse) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *CdnGetRegionsResponse) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *CdnGetRegionsResponse) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetRegion

`func (o *CdnGetRegionsResponse) GetRegion() []CdnGetRegionsResponseRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *CdnGetRegionsResponse) GetRegionOk() (*[]CdnGetRegionsResponseRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *CdnGetRegionsResponse) SetRegion(v []CdnGetRegionsResponseRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *CdnGetRegionsResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


