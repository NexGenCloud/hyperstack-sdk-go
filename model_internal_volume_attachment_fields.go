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

// checks if the InternalVolumeAttachmentFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalVolumeAttachmentFields{}

// InternalVolumeAttachmentFields struct for InternalVolumeAttachmentFields
type InternalVolumeAttachmentFields struct {
	CreatedAt *CustomTime           `json:"created_at,omitempty"`
	Device    *string               `json:"device,omitempty"`
	Status    *string               `json:"status,omitempty"`
	Volume    *InternalVolumeFields `json:"volume,omitempty"`
}

// NewInternalVolumeAttachmentFields instantiates a new InternalVolumeAttachmentFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalVolumeAttachmentFields() *InternalVolumeAttachmentFields {
	this := InternalVolumeAttachmentFields{}
	return &this
}

// NewInternalVolumeAttachmentFieldsWithDefaults instantiates a new InternalVolumeAttachmentFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalVolumeAttachmentFieldsWithDefaults() *InternalVolumeAttachmentFields {
	this := InternalVolumeAttachmentFields{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *InternalVolumeAttachmentFields) GetCreatedAt() CustomTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret CustomTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalVolumeAttachmentFields) GetCreatedAtOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *InternalVolumeAttachmentFields) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given CustomTime and assigns it to the CreatedAt field.
func (o *InternalVolumeAttachmentFields) SetCreatedAt(v CustomTime) {
	o.CreatedAt = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *InternalVolumeAttachmentFields) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalVolumeAttachmentFields) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *InternalVolumeAttachmentFields) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *InternalVolumeAttachmentFields) SetDevice(v string) {
	o.Device = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *InternalVolumeAttachmentFields) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalVolumeAttachmentFields) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *InternalVolumeAttachmentFields) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *InternalVolumeAttachmentFields) SetStatus(v string) {
	o.Status = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *InternalVolumeAttachmentFields) GetVolume() InternalVolumeFields {
	if o == nil || IsNil(o.Volume) {
		var ret InternalVolumeFields
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalVolumeAttachmentFields) GetVolumeOk() (*InternalVolumeFields, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *InternalVolumeAttachmentFields) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given InternalVolumeFields and assigns it to the Volume field.
func (o *InternalVolumeAttachmentFields) SetVolume(v InternalVolumeFields) {
	o.Volume = &v
}

func (o InternalVolumeAttachmentFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalVolumeAttachmentFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	return toSerialize, nil
}

type NullableInternalVolumeAttachmentFields struct {
	value *InternalVolumeAttachmentFields
	isSet bool
}

func (v NullableInternalVolumeAttachmentFields) Get() *InternalVolumeAttachmentFields {
	return v.value
}

func (v *NullableInternalVolumeAttachmentFields) Set(val *InternalVolumeAttachmentFields) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalVolumeAttachmentFields) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalVolumeAttachmentFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalVolumeAttachmentFields(val *InternalVolumeAttachmentFields) *NullableInternalVolumeAttachmentFields {
	return &NullableInternalVolumeAttachmentFields{value: val, isSet: true}
}

func (v NullableInternalVolumeAttachmentFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalVolumeAttachmentFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
