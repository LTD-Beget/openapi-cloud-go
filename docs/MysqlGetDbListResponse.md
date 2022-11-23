# MysqlGetDbListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Db** | Pointer to [**[]MysqlDb**](MysqlDb.md) |  | [optional] 

## Methods

### NewMysqlGetDbListResponse

`func NewMysqlGetDbListResponse() *MysqlGetDbListResponse`

NewMysqlGetDbListResponse instantiates a new MysqlGetDbListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMysqlGetDbListResponseWithDefaults

`func NewMysqlGetDbListResponseWithDefaults() *MysqlGetDbListResponse`

NewMysqlGetDbListResponseWithDefaults instantiates a new MysqlGetDbListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDb

`func (o *MysqlGetDbListResponse) GetDb() []MysqlDb`

GetDb returns the Db field if non-nil, zero value otherwise.

### GetDbOk

`func (o *MysqlGetDbListResponse) GetDbOk() (*[]MysqlDb, bool)`

GetDbOk returns a tuple with the Db field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDb

`func (o *MysqlGetDbListResponse) SetDb(v []MysqlDb)`

SetDb sets Db field to given value.

### HasDb

`func (o *MysqlGetDbListResponse) HasDb() bool`

HasDb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


