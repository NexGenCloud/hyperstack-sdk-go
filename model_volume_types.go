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

// checks if the VolumeTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeTypes{}

// VolumeTypes struct for VolumeTypes
type VolumeTypes struct {
	Message     *string  `json:"message,omitempty"`
	Status      *bool    `json:"status,omitempty"`
	VolumeTypes []string `json:"volume_types,omitempty"`
}

// NewVolumeTypes instantiates a new VolumeTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeTypes() *VolumeTypes {
	this := VolumeTypes{}
	return &this
}

// NewVolumeTypesWithDefaults instantiates a new VolumeTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeTypesWithDefaults() *VolumeTypes {
	this := VolumeTypes{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *VolumeTypes) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeTypes) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *VolumeTypes) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *VolumeTypes) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VolumeTypes) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeTypes) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VolumeTypes) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *VolumeTypes) SetStatus(v bool) {
	o.Status = &v
}

// GetVolumeTypes returns the VolumeTypes field value if set, zero value otherwise.
func (o *VolumeTypes) GetVolumeTypes() []string {
	if o == nil || IsNil(o.VolumeTypes) {
		var ret []string
		return ret
	}
	return o.VolumeTypes
}

// GetVolumeTypesOk returns a tuple with the VolumeTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeTypes) GetVolumeTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.VolumeTypes) {
		return nil, false
	}
	return o.VolumeTypes, true
}

// HasVolumeTypes returns a boolean if a field has been set.
func (o *VolumeTypes) HasVolumeTypes() bool {
	if o != nil && !IsNil(o.VolumeTypes) {
		return true
	}

	return false
}

// SetVolumeTypes gets a reference to the given []string and assigns it to the VolumeTypes field.
func (o *VolumeTypes) SetVolumeTypes(v []string) {
	o.VolumeTypes = v
}

func (o VolumeTypes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.VolumeTypes) {
		toSerialize["volume_types"] = o.VolumeTypes
	}
	return toSerialize, nil
}

type NullableVolumeTypes struct {
	value *VolumeTypes
	isSet bool
}

func (v NullableVolumeTypes) Get() *VolumeTypes {
	return v.value
}

func (v *NullableVolumeTypes) Set(val *VolumeTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeTypes(val *VolumeTypes) *NullableVolumeTypes {
	return &NullableVolumeTypes{value: val, isSet: true}
}

func (v NullableVolumeTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}