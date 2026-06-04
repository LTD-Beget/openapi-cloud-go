# NetworkDriveMountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkDrive** | Pointer to [**NetworkDriveNetworkDrive**](NetworkDriveNetworkDrive.md) |  | [optional] 
**Error** | Pointer to [**NetworkDriveMountResponseError**](NetworkDriveMountResponseError.md) |  | [optional] 

## Methods

### NewNetworkDriveMountResponse

`func NewNetworkDriveMountResponse() *NetworkDriveMountResponse`

NewNetworkDriveMountResponse instantiates a new NetworkDriveMountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkDriveMountResponseWithDefaults

`func NewNetworkDriveMountResponseWithDefaults() *NetworkDriveMountResponse`

NewNetworkDriveMountResponseWithDefaults instantiates a new NetworkDriveMountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkDrive

`func (o *NetworkDriveMountResponse) GetNetworkDrive() NetworkDriveNetworkDrive`

GetNetworkDrive returns the NetworkDrive field if non-nil, zero value otherwise.

### GetNetworkDriveOk

`func (o *NetworkDriveMountResponse) GetNetworkDriveOk() (*NetworkDriveNetworkDrive, bool)`

GetNetworkDriveOk returns a tuple with the NetworkDrive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDrive

`func (o *NetworkDriveMountResponse) SetNetworkDrive(v NetworkDriveNetworkDrive)`

SetNetworkDrive sets NetworkDrive field to given value.

### HasNetworkDrive

`func (o *NetworkDriveMountResponse) HasNetworkDrive() bool`

HasNetworkDrive returns a boolean if a field has been set.

### GetError

`func (o *NetworkDriveMountResponse) GetError() NetworkDriveMountResponseError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *NetworkDriveMountResponse) GetErrorOk() (*NetworkDriveMountResponseError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *NetworkDriveMountResponse) SetError(v NetworkDriveMountResponseError)`

SetError sets Error field to given value.

### HasError

`func (o *NetworkDriveMountResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


