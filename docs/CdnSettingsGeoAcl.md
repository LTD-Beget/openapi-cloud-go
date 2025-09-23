# CdnSettingsGeoAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**PolicyType** | Pointer to **string** |  | [optional] 
**ExceptedValues** | Pointer to [**map[string]CdnSettingsGeoAclStringArray**](CdnSettingsGeoAclStringArray.md) |  | [optional] 

## Methods

### NewCdnSettingsGeoAcl

`func NewCdnSettingsGeoAcl() *CdnSettingsGeoAcl`

NewCdnSettingsGeoAcl instantiates a new CdnSettingsGeoAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnSettingsGeoAclWithDefaults

`func NewCdnSettingsGeoAclWithDefaults() *CdnSettingsGeoAcl`

NewCdnSettingsGeoAclWithDefaults instantiates a new CdnSettingsGeoAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *CdnSettingsGeoAcl) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CdnSettingsGeoAcl) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CdnSettingsGeoAcl) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CdnSettingsGeoAcl) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPolicyType

`func (o *CdnSettingsGeoAcl) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *CdnSettingsGeoAcl) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *CdnSettingsGeoAcl) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *CdnSettingsGeoAcl) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetExceptedValues

`func (o *CdnSettingsGeoAcl) GetExceptedValues() map[string]CdnSettingsGeoAclStringArray`

GetExceptedValues returns the ExceptedValues field if non-nil, zero value otherwise.

### GetExceptedValuesOk

`func (o *CdnSettingsGeoAcl) GetExceptedValuesOk() (*map[string]CdnSettingsGeoAclStringArray, bool)`

GetExceptedValuesOk returns a tuple with the ExceptedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptedValues

`func (o *CdnSettingsGeoAcl) SetExceptedValues(v map[string]CdnSettingsGeoAclStringArray)`

SetExceptedValues sets ExceptedValues field to given value.

### HasExceptedValues

`func (o *CdnSettingsGeoAcl) HasExceptedValues() bool`

HasExceptedValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


