# CdnGetPriceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdn** | Pointer to [**[]CdnGetPriceResponseCdn**](CdnGetPriceResponseCdn.md) |  | [optional] 
**RequiredTopupAmount** | Pointer to **float64** |  | [optional] 

## Methods

### NewCdnGetPriceResponse

`func NewCdnGetPriceResponse() *CdnGetPriceResponse`

NewCdnGetPriceResponse instantiates a new CdnGetPriceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCdnGetPriceResponseWithDefaults

`func NewCdnGetPriceResponseWithDefaults() *CdnGetPriceResponse`

NewCdnGetPriceResponseWithDefaults instantiates a new CdnGetPriceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdn

`func (o *CdnGetPriceResponse) GetCdn() []CdnGetPriceResponseCdn`

GetCdn returns the Cdn field if non-nil, zero value otherwise.

### GetCdnOk

`func (o *CdnGetPriceResponse) GetCdnOk() (*[]CdnGetPriceResponseCdn, bool)`

GetCdnOk returns a tuple with the Cdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdn

`func (o *CdnGetPriceResponse) SetCdn(v []CdnGetPriceResponseCdn)`

SetCdn sets Cdn field to given value.

### HasCdn

`func (o *CdnGetPriceResponse) HasCdn() bool`

HasCdn returns a boolean if a field has been set.

### GetRequiredTopupAmount

`func (o *CdnGetPriceResponse) GetRequiredTopupAmount() float64`

GetRequiredTopupAmount returns the RequiredTopupAmount field if non-nil, zero value otherwise.

### GetRequiredTopupAmountOk

`func (o *CdnGetPriceResponse) GetRequiredTopupAmountOk() (*float64, bool)`

GetRequiredTopupAmountOk returns a tuple with the RequiredTopupAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredTopupAmount

`func (o *CdnGetPriceResponse) SetRequiredTopupAmount(v float64)`

SetRequiredTopupAmount sets RequiredTopupAmount field to given value.

### HasRequiredTopupAmount

`func (o *CdnGetPriceResponse) HasRequiredTopupAmount() bool`

HasRequiredTopupAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


