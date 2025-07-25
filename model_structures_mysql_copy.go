/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// StructuresMysqlCopy struct for StructuresMysqlCopy
type StructuresMysqlCopy struct {
	Id *string `json:"id,omitempty"`
	ServiceUuid *string `json:"service_uuid,omitempty"`
	ServiceName *string `json:"service_name,omitempty"`
	Version *string `json:"version,omitempty"`
	Date *string `json:"date,omitempty"`
	Size *string `json:"size,omitempty"`
	Region *string `json:"region,omitempty"`
}

// NewStructuresMysqlCopy instantiates a new StructuresMysqlCopy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStructuresMysqlCopy() *StructuresMysqlCopy {
	this := StructuresMysqlCopy{}
	return &this
}

// NewStructuresMysqlCopyWithDefaults instantiates a new StructuresMysqlCopy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStructuresMysqlCopyWithDefaults() *StructuresMysqlCopy {
	this := StructuresMysqlCopy{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StructuresMysqlCopy) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresMysqlCopy) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StructuresMysqlCopy) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StructuresMysqlCopy) SetId(v string) {
	o.Id = &v
}

// GetServiceUuid returns the ServiceUuid field value if set, zero value otherwise.
func (o *StructuresMysqlCopy) GetServiceUuid() string {
	if o == nil || isNil(o.ServiceUuid) {
		var ret string
		return ret
	}
	return *o.ServiceUuid
}

// GetServiceUuidOk returns a tuple with the ServiceUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresMysqlCopy) GetServiceUuidOk() (*string, bool) {
	if o == nil || isNil(o.ServiceUuid) {
    return nil, false
	}
	return o.ServiceUuid, true
}

// HasServiceUuid returns a boolean if a field has been set.
func (o *StructuresMysqlCopy) HasServiceUuid() bool {
	if o != nil && !isNil(o.ServiceUuid) {
		return true
	}

	return false
}

// SetServiceUuid gets a reference to the given string and assigns it to the ServiceUuid field.
func (o *StructuresMysqlCopy) SetServiceUuid(v string) {
	o.ServiceUuid = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *StructuresMysqlCopy) GetServiceName() string {
	if o == nil || isNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresMysqlCopy) GetServiceNameOk() (*string, bool) {
	if o == nil || isNil(o.ServiceName) {
    return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *StructuresMysqlCopy) HasServiceName() bool {
	if o != nil && !isNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *StructuresMysqlCopy) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *StructuresMysqlCopy) GetVersion() string {
	if o == nil || isNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresMysqlCopy) GetVersionOk() (*string, bool) {
	if o == nil || isNil(o.Version) {
    return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *StructuresMysqlCopy) HasVersion() bool {
	if o != nil && !isNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *StructuresMysqlCopy) SetVersion(v string) {
	o.Version = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *StructuresMysqlCopy) GetDate() string {
	if o == nil || isNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresMysqlCopy) GetDateOk() (*string, bool) {
	if o == nil || isNil(o.Date) {
    return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *StructuresMysqlCopy) HasDate() bool {
	if o != nil && !isNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *StructuresMysqlCopy) SetDate(v string) {
	o.Date = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *StructuresMysqlCopy) GetSize() string {
	if o == nil || isNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresMysqlCopy) GetSizeOk() (*string, bool) {
	if o == nil || isNil(o.Size) {
    return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *StructuresMysqlCopy) HasSize() bool {
	if o != nil && !isNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *StructuresMysqlCopy) SetSize(v string) {
	o.Size = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *StructuresMysqlCopy) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StructuresMysqlCopy) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
    return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *StructuresMysqlCopy) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *StructuresMysqlCopy) SetRegion(v string) {
	o.Region = &v
}

func (o StructuresMysqlCopy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.ServiceUuid) {
		toSerialize["service_uuid"] = o.ServiceUuid
	}
	if !isNil(o.ServiceName) {
		toSerialize["service_name"] = o.ServiceName
	}
	if !isNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !isNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !isNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableStructuresMysqlCopy struct {
	value *StructuresMysqlCopy
	isSet bool
}

func (v NullableStructuresMysqlCopy) Get() *StructuresMysqlCopy {
	return v.value
}

func (v *NullableStructuresMysqlCopy) Set(val *StructuresMysqlCopy) {
	v.value = val
	v.isSet = true
}

func (v NullableStructuresMysqlCopy) IsSet() bool {
	return v.isSet
}

func (v *NullableStructuresMysqlCopy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStructuresMysqlCopy(val *StructuresMysqlCopy) *NullableStructuresMysqlCopy {
	return &NullableStructuresMysqlCopy{value: val, isSet: true}
}

func (v NullableStructuresMysqlCopy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStructuresMysqlCopy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


