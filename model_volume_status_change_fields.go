/*
Infrahub-API

Leverage the Infrahub API and Hyperstack platform to easily create, manage, and scale powerful GPU virtual machines and their associated resources.   Access this SDK to automate the deployment of your workloads and streamline your infrastructure management.  To contribute, please raise an issue with a bug report, feature request, feedback, or general inquiry.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"encoding/json"
)

// checks if the VolumeStatusChangeFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeStatusChangeFields{}

// VolumeStatusChangeFields struct for VolumeStatusChangeFields
type VolumeStatusChangeFields struct {
	CreatedAt      *CustomTime `json:"created_at,omitempty"`
	CurrentStatus  *string     `json:"current_status,omitempty"`
	Id             *int32      `json:"id,omitempty"`
	PreviousStatus *string     `json:"previous_status,omitempty"`
	VolumeId       *int32      `json:"volume_id,omitempty"`
}

// NewVolumeStatusChangeFields instantiates a new VolumeStatusChangeFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeStatusChangeFields() *VolumeStatusChangeFields {
	this := VolumeStatusChangeFields{}
	return &this
}

// NewVolumeStatusChangeFieldsWithDefaults instantiates a new VolumeStatusChangeFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeStatusChangeFieldsWithDefaults() *VolumeStatusChangeFields {
	this := VolumeStatusChangeFields{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *VolumeStatusChangeFields) GetCreatedAt() CustomTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret CustomTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeStatusChangeFields) GetCreatedAtOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *VolumeStatusChangeFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given CustomTime and assigns it to the CreatedAt field.
func (o *VolumeStatusChangeFields) SetCreatedAt(v CustomTime) {
	o.CreatedAt = &v
}

// GetCurrentStatus returns the CurrentStatus field value if set, zero value otherwise.
func (o *VolumeStatusChangeFields) GetCurrentStatus() string {
	if o == nil || IsNil(o.CurrentStatus) {
		var ret string
		return ret
	}
	return *o.CurrentStatus
}

// GetCurrentStatusOk returns a tuple with the CurrentStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeStatusChangeFields) GetCurrentStatusOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentStatus) {
		return nil, false
	}
	return o.CurrentStatus, true
}

// HasCurrentStatus returns a boolean if a field has been set.
func (o *VolumeStatusChangeFields) HasCurrentStatus() bool {
	if o != nil && !IsNil(o.CurrentStatus) {
		return true
	}

	return false
}

// SetCurrentStatus gets a reference to the given string and assigns it to the CurrentStatus field.
func (o *VolumeStatusChangeFields) SetCurrentStatus(v string) {
	o.CurrentStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VolumeStatusChangeFields) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeStatusChangeFields) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VolumeStatusChangeFields) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *VolumeStatusChangeFields) SetId(v int32) {
	o.Id = &v
}

// GetPreviousStatus returns the PreviousStatus field value if set, zero value otherwise.
func (o *VolumeStatusChangeFields) GetPreviousStatus() string {
	if o == nil || IsNil(o.PreviousStatus) {
		var ret string
		return ret
	}
	return *o.PreviousStatus
}

// GetPreviousStatusOk returns a tuple with the PreviousStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeStatusChangeFields) GetPreviousStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousStatus) {
		return nil, false
	}
	return o.PreviousStatus, true
}

// HasPreviousStatus returns a boolean if a field has been set.
func (o *VolumeStatusChangeFields) HasPreviousStatus() bool {
	if o != nil && !IsNil(o.PreviousStatus) {
		return true
	}

	return false
}

// SetPreviousStatus gets a reference to the given string and assigns it to the PreviousStatus field.
func (o *VolumeStatusChangeFields) SetPreviousStatus(v string) {
	o.PreviousStatus = &v
}

// GetVolumeId returns the VolumeId field value if set, zero value otherwise.
func (o *VolumeStatusChangeFields) GetVolumeId() int32 {
	if o == nil || IsNil(o.VolumeId) {
		var ret int32
		return ret
	}
	return *o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeStatusChangeFields) GetVolumeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VolumeId) {
		return nil, false
	}
	return o.VolumeId, true
}

// HasVolumeId returns a boolean if a field has been set.
func (o *VolumeStatusChangeFields) HasVolumeId() bool {
	if o != nil && !IsNil(o.VolumeId) {
		return true
	}

	return false
}

// SetVolumeId gets a reference to the given int32 and assigns it to the VolumeId field.
func (o *VolumeStatusChangeFields) SetVolumeId(v int32) {
	o.VolumeId = &v
}

func (o VolumeStatusChangeFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeStatusChangeFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CurrentStatus) {
		toSerialize["current_status"] = o.CurrentStatus
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PreviousStatus) {
		toSerialize["previous_status"] = o.PreviousStatus
	}
	if !IsNil(o.VolumeId) {
		toSerialize["volume_id"] = o.VolumeId
	}
	return toSerialize, nil
}

type NullableVolumeStatusChangeFields struct {
	value *VolumeStatusChangeFields
	isSet bool
}

func (v NullableVolumeStatusChangeFields) Get() *VolumeStatusChangeFields {
	return v.value
}

func (v *NullableVolumeStatusChangeFields) Set(val *VolumeStatusChangeFields) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeStatusChangeFields) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeStatusChangeFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeStatusChangeFields(val *VolumeStatusChangeFields) *NullableVolumeStatusChangeFields {
	return &NullableVolumeStatusChangeFields{value: val, isSet: true}
}

func (v NullableVolumeStatusChangeFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeStatusChangeFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}