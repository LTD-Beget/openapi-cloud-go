# NetworkDriveGetLimitsResponseLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** |  | [optional] 
**MaxNetworkDrives** | Pointer to **int32** |  | [optional] 
**AvailableNetworkDrives** | Pointer to **int32** |  | [optional] 
**MaxMountsPerResource** | Pointer to **int32** |  | [optional] 
**MaxTotalSize** | Pointer to [**[]NetworkDriveGetLimitsResponseLimitsTotalSize**](NetworkDriveGetLimitsResponseLimitsTotalSize.md) |  | [optional] 
**AvailableTotalSize** | Pointer to [**[]NetworkDriveGetLimitsResponseLimitsTotalSize**](NetworkDriveGetLimitsResponseLimitsTotalSize.md) |  | [optional] 

## Methods

### NewNetworkDriveGetLimitsResponseLimits

`func NewNetworkDriveGetLimitsResponseLimits() *NetworkDriveGetLimitsResponseLimits`

NewNetworkDriveGetLimitsResponseLimits instantiates a new NetworkDriveGetLimitsResponseLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDriveGetLimitsResponseLimitsWithDefaults

`func NewNetworkDriveGetLimitsResponseLimitsWithDefaults() *NetworkDriveGetLimitsResponseLimits`

NewNetworkDriveGetLimitsResponseLimitsWithDefaults instantiates a new NetworkDriveGetLimitsResponseLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *NetworkDriveGetLimitsResponseLimits) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NetworkDriveGetLimitsResponseLimits) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NetworkDriveGetLimitsResponseLimits) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NetworkDriveGetLimitsResponseLimits) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetMaxNetworkDrives

`func (o *NetworkDriveGetLimitsResponseLimits) GetMaxNetworkDrives() int32`

GetMaxNetworkDrives returns the MaxNetworkDrives field if non-nil, zero value otherwise.

### GetMaxNetworkDrivesOk

`func (o *NetworkDriveGetLimitsResponseLimits) GetMaxNetworkDrivesOk() (*int32, bool)`

GetMaxNetworkDrivesOk returns a tuple with the MaxNetworkDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNetworkDrives

`func (o *NetworkDriveGetLimitsResponseLimits) SetMaxNetworkDrives(v int32)`

SetMaxNetworkDrives sets MaxNetworkDrives field to given value.

### HasMaxNetworkDrives

`func (o *NetworkDriveGetLimitsResponseLimits) HasMaxNetworkDrives() bool`

HasMaxNetworkDrives returns a boolean if a field has been set.

### GetAvailableNetworkDrives

`func (o *NetworkDriveGetLimitsResponseLimits) GetAvailableNetworkDrives() int32`

GetAvailableNetworkDrives returns the AvailableNetworkDrives field if non-nil, zero value otherwise.

### GetAvailableNetworkDrivesOk

`func (o *NetworkDriveGetLimitsResponseLimits) GetAvailableNetworkDrivesOk() (*int32, bool)`

GetAvailableNetworkDrivesOk returns a tuple with the AvailableNetworkDrives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableNetworkDrives

`func (o *NetworkDriveGetLimitsResponseLimits) SetAvailableNetworkDrives(v int32)`

SetAvailableNetworkDrives sets AvailableNetworkDrives field to given value.

### HasAvailableNetworkDrives

`func (o *NetworkDriveGetLimitsResponseLimits) HasAvailableNetworkDrives() bool`

HasAvailableNetworkDrives returns a boolean if a field has been set.

### GetMaxMountsPerResource

`func (o *NetworkDriveGetLimitsResponseLimits) GetMaxMountsPerResource() int32`

GetMaxMountsPerResource returns the MaxMountsPerResource field if non-nil, zero value otherwise.

### GetMaxMountsPerResourceOk

`func (o *NetworkDriveGetLimitsResponseLimits) GetMaxMountsPerResourceOk() (*int32, bool)`

GetMaxMountsPerResourceOk returns a tuple with the MaxMountsPerResource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMountsPerResource

`func (o *NetworkDriveGetLimitsResponseLimits) SetMaxMountsPerResource(v int32)`

SetMaxMountsPerResource sets MaxMountsPerResource field to given value.

### HasMaxMountsPerResource

`func (o *NetworkDriveGetLimitsResponseLimits) HasMaxMountsPerResource() bool`

HasMaxMountsPerResource returns a boolean if a field has been set.

### GetMaxTotalSize

`func (o *NetworkDriveGetLimitsResponseLimits) GetMaxTotalSize() []NetworkDriveGetLimitsResponseLimitsTotalSize`

GetMaxTotalSize returns the MaxTotalSize field if non-nil, zero value otherwise.

### GetMaxTotalSizeOk

`func (o *NetworkDriveGetLimitsResponseLimits) GetMaxTotalSizeOk() (*[]NetworkDriveGetLimitsResponseLimitsTotalSize, bool)`

GetMaxTotalSizeOk returns a tuple with the MaxTotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTotalSize

`func (o *NetworkDriveGetLimitsResponseLimits) SetMaxTotalSize(v []NetworkDriveGetLimitsResponseLimitsTotalSize)`

SetMaxTotalSize sets MaxTotalSize field to given value.

### HasMaxTotalSize

`func (o *NetworkDriveGetLimitsResponseLimits) HasMaxTotalSize() bool`

HasMaxTotalSize returns a boolean if a field has been set.

### GetAvailableTotalSize

`func (o *NetworkDriveGetLimitsResponseLimits) GetAvailableTotalSize() []NetworkDriveGetLimitsResponseLimitsTotalSize`

GetAvailableTotalSize returns the AvailableTotalSize field if non-nil, zero value otherwise.

### GetAvailableTotalSizeOk

`func (o *NetworkDriveGetLimitsResponseLimits) GetAvailableTotalSizeOk() (*[]NetworkDriveGetLimitsResponseLimitsTotalSize, bool)`

GetAvailableTotalSizeOk returns a tuple with the AvailableTotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTotalSize

`func (o *NetworkDriveGetLimitsResponseLimits) SetAvailableTotalSize(v []NetworkDriveGetLimitsResponseLimitsTotalSize)`

SetAvailableTotalSize sets AvailableTotalSize field to given value.

### HasAvailableTotalSize

`func (o *NetworkDriveGetLimitsResponseLimits) HasAvailableTotalSize() bool`

HasAvailableTotalSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


