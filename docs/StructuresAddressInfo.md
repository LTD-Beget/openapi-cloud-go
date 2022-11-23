# StructuresAddressInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Public** | Pointer to [**[]StructuresAddressInfoPublicAddress**](StructuresAddressInfoPublicAddress.md) |  | [optional] 
**Private** | Pointer to [**[]StructuresAddressInfoPrivateAddress**](StructuresAddressInfoPrivateAddress.md) |  | [optional] 

## Methods

### NewStructuresAddressInfo

`func NewStructuresAddressInfo() *StructuresAddressInfo`

NewStructuresAddressInfo instantiates a new StructuresAddressInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStructuresAddressInfoWithDefaults

`func NewStructuresAddressInfoWithDefaults() *StructuresAddressInfo`

NewStructuresAddressInfoWithDefaults instantiates a new StructuresAddressInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublic

`func (o *StructuresAddressInfo) GetPublic() []StructuresAddressInfoPublicAddress`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *StructuresAddressInfo) GetPublicOk() (*[]StructuresAddressInfoPublicAddress, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *StructuresAddressInfo) SetPublic(v []StructuresAddressInfoPublicAddress)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *StructuresAddressInfo) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetPrivate

`func (o *StructuresAddressInfo) GetPrivate() []StructuresAddressInfoPrivateAddress`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *StructuresAddressInfo) GetPrivateOk() (*[]StructuresAddressInfoPrivateAddress, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *StructuresAddressInfo) SetPrivate(v []StructuresAddressInfoPrivateAddress)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *StructuresAddressInfo) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


