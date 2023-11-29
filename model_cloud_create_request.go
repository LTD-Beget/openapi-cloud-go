/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// CloudCreateRequest struct for CloudCreateRequest
type CloudCreateRequest struct {
	ConfigurationId *string `json:"configuration_id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Description *string `json:"description,omitempty"`
	MysqlParams *MysqlCreateParams `json:"mysql_params,omitempty"`
	Extra *string `json:"extra,omitempty"`
	Region *string `json:"region,omitempty"`
}

// NewCloudCreateRequest instantiates a new CloudCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCreateRequest() *CloudCreateRequest {
	this := CloudCreateRequest{}
	return &this
}

// NewCloudCreateRequestWithDefaults instantiates a new CloudCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCreateRequestWithDefaults() *CloudCreateRequest {
	this := CloudCreateRequest{}
	return &this
}

// GetConfigurationId returns the ConfigurationId field value if set, zero value otherwise.
func (o *CloudCreateRequest) GetConfigurationId() string {
	if o == nil || isNil(o.ConfigurationId) {
		var ret string
		return ret
	}
	return *o.ConfigurationId
}

// GetConfigurationIdOk returns a tuple with the ConfigurationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateRequest) GetConfigurationIdOk() (*string, bool) {
	if o == nil || isNil(o.ConfigurationId) {
    return nil, false
	}
	return o.ConfigurationId, true
}

// HasConfigurationId returns a boolean if a field has been set.
func (o *CloudCreateRequest) HasConfigurationId() bool {
	if o != nil && !isNil(o.ConfigurationId) {
		return true
	}

	return false
}

// SetConfigurationId gets a reference to the given string and assigns it to the ConfigurationId field.
func (o *CloudCreateRequest) SetConfigurationId(v string) {
	o.ConfigurationId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CloudCreateRequest) GetDisplayName() string {
	if o == nil || isNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.DisplayName) {
    return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CloudCreateRequest) HasDisplayName() bool {
	if o != nil && !isNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CloudCreateRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CloudCreateRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CloudCreateRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CloudCreateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMysqlParams returns the MysqlParams field value if set, zero value otherwise.
func (o *CloudCreateRequest) GetMysqlParams() MysqlCreateParams {
	if o == nil || isNil(o.MysqlParams) {
		var ret MysqlCreateParams
		return ret
	}
	return *o.MysqlParams
}

// GetMysqlParamsOk returns a tuple with the MysqlParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateRequest) GetMysqlParamsOk() (*MysqlCreateParams, bool) {
	if o == nil || isNil(o.MysqlParams) {
    return nil, false
	}
	return o.MysqlParams, true
}

// HasMysqlParams returns a boolean if a field has been set.
func (o *CloudCreateRequest) HasMysqlParams() bool {
	if o != nil && !isNil(o.MysqlParams) {
		return true
	}

	return false
}

// SetMysqlParams gets a reference to the given MysqlCreateParams and assigns it to the MysqlParams field.
func (o *CloudCreateRequest) SetMysqlParams(v MysqlCreateParams) {
	o.MysqlParams = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *CloudCreateRequest) GetExtra() string {
	if o == nil || isNil(o.Extra) {
		var ret string
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateRequest) GetExtraOk() (*string, bool) {
	if o == nil || isNil(o.Extra) {
    return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *CloudCreateRequest) HasExtra() bool {
	if o != nil && !isNil(o.Extra) {
		return true
	}

	return false
}

// SetExtra gets a reference to the given string and assigns it to the Extra field.
func (o *CloudCreateRequest) SetExtra(v string) {
	o.Extra = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CloudCreateRequest) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateRequest) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CloudCreateRequest) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CloudCreateRequest) SetRegion(v string) {
	o.Region = &v
}

func (o CloudCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ConfigurationId) {
		toSerialize["configuration_id"] = o.ConfigurationId
	}
	if !isNil(o.DisplayName) {
		toSerialize["display_name"] = o.DisplayName
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.MysqlParams) {
		toSerialize["mysql_params"] = o.MysqlParams
	}
	if !isNil(o.Extra) {
		toSerialize["extra"] = o.Extra
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableCloudCreateRequest struct {
	value *CloudCreateRequest
	isSet bool
}

func (v NullableCloudCreateRequest) Get() *CloudCreateRequest {
	return v.value
}

func (v *NullableCloudCreateRequest) Set(val *CloudCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCreateRequest(val *CloudCreateRequest) *NullableCloudCreateRequest {
	return &NullableCloudCreateRequest{value: val, isSet: true}
}

func (v NullableCloudCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


