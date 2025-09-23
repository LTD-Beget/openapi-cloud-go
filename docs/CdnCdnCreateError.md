# CdnCdnCreateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**CdnCdnCreateErrorError**](CdnCdnCreateErrorError.md) |  | [optional] 
**DomainError** | Pointer to [**[]CdnCdnCreateErrorDomainError**](CdnCdnCreateErrorDomainError.md) |  | [optional] 

## Methods

### NewCdnCdnCreateError

`func NewCdnCdnCreateError() *CdnCdnCreateError`

NewCdnCdnCreateError instantiates a new CdnCdnCreateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnCdnCreateErrorWithDefaults

`func NewCdnCdnCreateErrorWithDefaults() *CdnCdnCreateError`

NewCdnCdnCreateErrorWithDefaults instantiates a new CdnCdnCreateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CdnCdnCreateError) GetError() CdnCdnCreateErrorError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CdnCdnCreateError) GetErrorOk() (*CdnCdnCreateErrorError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CdnCdnCreateError) SetError(v CdnCdnCreateErrorError)`

SetError sets Error field to given value.

### HasError

`func (o *CdnCdnCreateError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDomainError

`func (o *CdnCdnCreateError) GetDomainError() []CdnCdnCreateErrorDomainError`

GetDomainError returns the DomainError field if non-nil, zero value otherwise.

### GetDomainErrorOk

`func (o *CdnCdnCreateError) GetDomainErrorOk() (*[]CdnCdnCreateErrorDomainError, bool)`

GetDomainErrorOk returns a tuple with the DomainError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainError

`func (o *CdnCdnCreateError) SetDomainError(v []CdnCdnCreateErrorDomainError)`

SetDomainError sets DomainError field to given value.

### HasDomainError

`func (o *CdnCdnCreateError) HasDomainError() bool`

HasDomainError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


