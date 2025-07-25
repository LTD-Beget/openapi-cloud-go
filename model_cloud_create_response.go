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

// CloudCreateResponse struct for CloudCreateResponse
type CloudCreateResponse struct {
	Service *CloudService `json:"service,omitempty"`
	MysqlError *MysqlCreateError `json:"mysql_error,omitempty"`
	PostgresqlError *PostgresqlPgCreateError `json:"postgresql_error,omitempty"`
	S3Error *S3S3CreateError `json:"s3_error,omitempty"`
}

// NewCloudCreateResponse instantiates a new CloudCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudCreateResponse() *CloudCreateResponse {
	this := CloudCreateResponse{}
	return &this
}

// NewCloudCreateResponseWithDefaults instantiates a new CloudCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudCreateResponseWithDefaults() *CloudCreateResponse {
	this := CloudCreateResponse{}
	return &this
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CloudCreateResponse) GetService() CloudService {
	if o == nil || isNil(o.Service) {
		var ret CloudService
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateResponse) GetServiceOk() (*CloudService, bool) {
	if o == nil || isNil(o.Service) {
    return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CloudCreateResponse) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given CloudService and assigns it to the Service field.
func (o *CloudCreateResponse) SetService(v CloudService) {
	o.Service = &v
}

// GetMysqlError returns the MysqlError field value if set, zero value otherwise.
func (o *CloudCreateResponse) GetMysqlError() MysqlCreateError {
	if o == nil || isNil(o.MysqlError) {
		var ret MysqlCreateError
		return ret
	}
	return *o.MysqlError
}

// GetMysqlErrorOk returns a tuple with the MysqlError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateResponse) GetMysqlErrorOk() (*MysqlCreateError, bool) {
	if o == nil || isNil(o.MysqlError) {
    return nil, false
	}
	return o.MysqlError, true
}

// HasMysqlError returns a boolean if a field has been set.
func (o *CloudCreateResponse) HasMysqlError() bool {
	if o != nil && !isNil(o.MysqlError) {
		return true
	}

	return false
}

// SetMysqlError gets a reference to the given MysqlCreateError and assigns it to the MysqlError field.
func (o *CloudCreateResponse) SetMysqlError(v MysqlCreateError) {
	o.MysqlError = &v
}

// GetPostgresqlError returns the PostgresqlError field value if set, zero value otherwise.
func (o *CloudCreateResponse) GetPostgresqlError() PostgresqlPgCreateError {
	if o == nil || isNil(o.PostgresqlError) {
		var ret PostgresqlPgCreateError
		return ret
	}
	return *o.PostgresqlError
}

// GetPostgresqlErrorOk returns a tuple with the PostgresqlError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateResponse) GetPostgresqlErrorOk() (*PostgresqlPgCreateError, bool) {
	if o == nil || isNil(o.PostgresqlError) {
    return nil, false
	}
	return o.PostgresqlError, true
}

// HasPostgresqlError returns a boolean if a field has been set.
func (o *CloudCreateResponse) HasPostgresqlError() bool {
	if o != nil && !isNil(o.PostgresqlError) {
		return true
	}

	return false
}

// SetPostgresqlError gets a reference to the given PostgresqlPgCreateError and assigns it to the PostgresqlError field.
func (o *CloudCreateResponse) SetPostgresqlError(v PostgresqlPgCreateError) {
	o.PostgresqlError = &v
}

// GetS3Error returns the S3Error field value if set, zero value otherwise.
func (o *CloudCreateResponse) GetS3Error() S3S3CreateError {
	if o == nil || isNil(o.S3Error) {
		var ret S3S3CreateError
		return ret
	}
	return *o.S3Error
}

// GetS3ErrorOk returns a tuple with the S3Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudCreateResponse) GetS3ErrorOk() (*S3S3CreateError, bool) {
	if o == nil || isNil(o.S3Error) {
    return nil, false
	}
	return o.S3Error, true
}

// HasS3Error returns a boolean if a field has been set.
func (o *CloudCreateResponse) HasS3Error() bool {
	if o != nil && !isNil(o.S3Error) {
		return true
	}

	return false
}

// SetS3Error gets a reference to the given S3S3CreateError and assigns it to the S3Error field.
func (o *CloudCreateResponse) SetS3Error(v S3S3CreateError) {
	o.S3Error = &v
}

func (o CloudCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.MysqlError) {
		toSerialize["mysql_error"] = o.MysqlError
	}
	if !isNil(o.PostgresqlError) {
		toSerialize["postgresql_error"] = o.PostgresqlError
	}
	if !isNil(o.S3Error) {
		toSerialize["s3_error"] = o.S3Error
	}
	return json.Marshal(toSerialize)
}

type NullableCloudCreateResponse struct {
	value *CloudCreateResponse
	isSet bool
}

func (v NullableCloudCreateResponse) Get() *CloudCreateResponse {
	return v.value
}

func (v *NullableCloudCreateResponse) Set(val *CloudCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudCreateResponse(val *CloudCreateResponse) *NullableCloudCreateResponse {
	return &NullableCloudCreateResponse{value: val, isSet: true}
}

func (v NullableCloudCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


