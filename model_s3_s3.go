/*
API Управляемых сервисов

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.4.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package begetOpenapiCloud

import (
	"encoding/json"
)

// S3S3 struct for S3S3
type S3S3 struct {
	Public *bool `json:"public,omitempty"`
	AccessKey *string `json:"access_key,omitempty"`
	SecretKey *string `json:"secret_key,omitempty"`
	Fqdn *string `json:"fqdn,omitempty"`
	Cors []S3Cors `json:"cors,omitempty"`
	QuotaUsedSize *int32 `json:"quota_used_size,omitempty"`
	Ftp *S3Ftp `json:"ftp,omitempty"`
	Sftp *S3Sftp `json:"sftp,omitempty"`
}

// NewS3S3 instantiates a new S3S3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewS3S3() *S3S3 {
	this := S3S3{}
	return &this
}

// NewS3S3WithDefaults instantiates a new S3S3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewS3S3WithDefaults() *S3S3 {
	this := S3S3{}
	return &this
}

// GetPublic returns the Public field value if set, zero value otherwise.
func (o *S3S3) GetPublic() bool {
	if o == nil || isNil(o.Public) {
		var ret bool
		return ret
	}
	return *o.Public
}

// GetPublicOk returns a tuple with the Public field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3S3) GetPublicOk() (*bool, bool) {
	if o == nil || isNil(o.Public) {
    return nil, false
	}
	return o.Public, true
}

// HasPublic returns a boolean if a field has been set.
func (o *S3S3) HasPublic() bool {
	if o != nil && !isNil(o.Public) {
		return true
	}

	return false
}

// SetPublic gets a reference to the given bool and assigns it to the Public field.
func (o *S3S3) SetPublic(v bool) {
	o.Public = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *S3S3) GetAccessKey() string {
	if o == nil || isNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3S3) GetAccessKeyOk() (*string, bool) {
	if o == nil || isNil(o.AccessKey) {
    return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *S3S3) HasAccessKey() bool {
	if o != nil && !isNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *S3S3) SetAccessKey(v string) {
	o.AccessKey = &v
}

// GetSecretKey returns the SecretKey field value if set, zero value otherwise.
func (o *S3S3) GetSecretKey() string {
	if o == nil || isNil(o.SecretKey) {
		var ret string
		return ret
	}
	return *o.SecretKey
}

// GetSecretKeyOk returns a tuple with the SecretKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3S3) GetSecretKeyOk() (*string, bool) {
	if o == nil || isNil(o.SecretKey) {
    return nil, false
	}
	return o.SecretKey, true
}

// HasSecretKey returns a boolean if a field has been set.
func (o *S3S3) HasSecretKey() bool {
	if o != nil && !isNil(o.SecretKey) {
		return true
	}

	return false
}

// SetSecretKey gets a reference to the given string and assigns it to the SecretKey field.
func (o *S3S3) SetSecretKey(v string) {
	o.SecretKey = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *S3S3) GetFqdn() string {
	if o == nil || isNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3S3) GetFqdnOk() (*string, bool) {
	if o == nil || isNil(o.Fqdn) {
    return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *S3S3) HasFqdn() bool {
	if o != nil && !isNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *S3S3) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetCors returns the Cors field value if set, zero value otherwise.
func (o *S3S3) GetCors() []S3Cors {
	if o == nil || isNil(o.Cors) {
		var ret []S3Cors
		return ret
	}
	return o.Cors
}

// GetCorsOk returns a tuple with the Cors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3S3) GetCorsOk() ([]S3Cors, bool) {
	if o == nil || isNil(o.Cors) {
    return nil, false
	}
	return o.Cors, true
}

// HasCors returns a boolean if a field has been set.
func (o *S3S3) HasCors() bool {
	if o != nil && !isNil(o.Cors) {
		return true
	}

	return false
}

// SetCors gets a reference to the given []S3Cors and assigns it to the Cors field.
func (o *S3S3) SetCors(v []S3Cors) {
	o.Cors = v
}

// GetQuotaUsedSize returns the QuotaUsedSize field value if set, zero value otherwise.
func (o *S3S3) GetQuotaUsedSize() int32 {
	if o == nil || isNil(o.QuotaUsedSize) {
		var ret int32
		return ret
	}
	return *o.QuotaUsedSize
}

// GetQuotaUsedSizeOk returns a tuple with the QuotaUsedSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3S3) GetQuotaUsedSizeOk() (*int32, bool) {
	if o == nil || isNil(o.QuotaUsedSize) {
    return nil, false
	}
	return o.QuotaUsedSize, true
}

// HasQuotaUsedSize returns a boolean if a field has been set.
func (o *S3S3) HasQuotaUsedSize() bool {
	if o != nil && !isNil(o.QuotaUsedSize) {
		return true
	}

	return false
}

// SetQuotaUsedSize gets a reference to the given int32 and assigns it to the QuotaUsedSize field.
func (o *S3S3) SetQuotaUsedSize(v int32) {
	o.QuotaUsedSize = &v
}

// GetFtp returns the Ftp field value if set, zero value otherwise.
func (o *S3S3) GetFtp() S3Ftp {
	if o == nil || isNil(o.Ftp) {
		var ret S3Ftp
		return ret
	}
	return *o.Ftp
}

// GetFtpOk returns a tuple with the Ftp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3S3) GetFtpOk() (*S3Ftp, bool) {
	if o == nil || isNil(o.Ftp) {
    return nil, false
	}
	return o.Ftp, true
}

// HasFtp returns a boolean if a field has been set.
func (o *S3S3) HasFtp() bool {
	if o != nil && !isNil(o.Ftp) {
		return true
	}

	return false
}

// SetFtp gets a reference to the given S3Ftp and assigns it to the Ftp field.
func (o *S3S3) SetFtp(v S3Ftp) {
	o.Ftp = &v
}

// GetSftp returns the Sftp field value if set, zero value otherwise.
func (o *S3S3) GetSftp() S3Sftp {
	if o == nil || isNil(o.Sftp) {
		var ret S3Sftp
		return ret
	}
	return *o.Sftp
}

// GetSftpOk returns a tuple with the Sftp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *S3S3) GetSftpOk() (*S3Sftp, bool) {
	if o == nil || isNil(o.Sftp) {
    return nil, false
	}
	return o.Sftp, true
}

// HasSftp returns a boolean if a field has been set.
func (o *S3S3) HasSftp() bool {
	if o != nil && !isNil(o.Sftp) {
		return true
	}

	return false
}

// SetSftp gets a reference to the given S3Sftp and assigns it to the Sftp field.
func (o *S3S3) SetSftp(v S3Sftp) {
	o.Sftp = &v
}

func (o S3S3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Public) {
		toSerialize["public"] = o.Public
	}
	if !isNil(o.AccessKey) {
		toSerialize["access_key"] = o.AccessKey
	}
	if !isNil(o.SecretKey) {
		toSerialize["secret_key"] = o.SecretKey
	}
	if !isNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !isNil(o.Cors) {
		toSerialize["cors"] = o.Cors
	}
	if !isNil(o.QuotaUsedSize) {
		toSerialize["quota_used_size"] = o.QuotaUsedSize
	}
	if !isNil(o.Ftp) {
		toSerialize["ftp"] = o.Ftp
	}
	if !isNil(o.Sftp) {
		toSerialize["sftp"] = o.Sftp
	}
	return json.Marshal(toSerialize)
}

type NullableS3S3 struct {
	value *S3S3
	isSet bool
}

func (v NullableS3S3) Get() *S3S3 {
	return v.value
}

func (v *NullableS3S3) Set(val *S3S3) {
	v.value = val
	v.isSet = true
}

func (v NullableS3S3) IsSet() bool {
	return v.isSet
}

func (v *NullableS3S3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableS3S3(val *S3S3) *NullableS3S3 {
	return &NullableS3S3{value: val, isSet: true}
}

func (v NullableS3S3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableS3S3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


