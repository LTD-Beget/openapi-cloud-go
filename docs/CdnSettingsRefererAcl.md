# CdnSettingsRefererAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** |  | [optional] 
**PolicyType** | Pointer to **string** |  | [optional] 
**ExceptedDomain** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCdnSettingsRefererAcl

`func NewCdnSettingsRefererAcl() *CdnSettingsRefererAcl`

NewCdnSettingsRefererAcl instantiates a new CdnSettingsRefererAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnSettingsRefererAclWithDefaults

`func NewCdnSettingsRefererAclWithDefaults() *CdnSettingsRefererAcl`

NewCdnSettingsRefererAclWithDefaults instantiates a new CdnSettingsRefererAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *CdnSettingsRefererAcl) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CdnSettingsRefererAcl) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CdnSettingsRefererAcl) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *CdnSettingsRefererAcl) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetPolicyType

`func (o *CdnSettingsRefererAcl) GetPolicyType() string`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *CdnSettingsRefererAcl) GetPolicyTypeOk() (*string, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *CdnSettingsRefererAcl) SetPolicyType(v string)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *CdnSettingsRefererAcl) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetExceptedDomain

`func (o *CdnSettingsRefererAcl) GetExceptedDomain() []string`

GetExceptedDomain returns the ExceptedDomain field if non-nil, zero value otherwise.

### GetExceptedDomainOk

`func (o *CdnSettingsRefererAcl) GetExceptedDomainOk() (*[]string, bool)`

GetExceptedDomainOk returns a tuple with the ExceptedDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptedDomain

`func (o *CdnSettingsRefererAcl) SetExceptedDomain(v []string)`

SetExceptedDomain sets ExceptedDomain field to given value.

### HasExceptedDomain

`func (o *CdnSettingsRefererAcl) HasExceptedDomain() bool`

HasExceptedDomain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


