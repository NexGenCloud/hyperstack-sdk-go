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

// checks if the VolumeOverviewFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VolumeOverviewFields{}

// VolumeOverviewFields struct for VolumeOverviewFields
type VolumeOverviewFields struct {
	CostPerHour *float32 `json:"cost_per_hour,omitempty"`
	Count       *int32   `json:"count,omitempty"`
	Using       *int32   `json:"using,omitempty"`
}

// NewVolumeOverviewFields instantiates a new VolumeOverviewFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolumeOverviewFields() *VolumeOverviewFields {
	this := VolumeOverviewFields{}
	return &this
}

// NewVolumeOverviewFieldsWithDefaults instantiates a new VolumeOverviewFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeOverviewFieldsWithDefaults() *VolumeOverviewFields {
	this := VolumeOverviewFields{}
	return &this
}

// GetCostPerHour returns the CostPerHour field value if set, zero value otherwise.
func (o *VolumeOverviewFields) GetCostPerHour() float32 {
	if o == nil || IsNil(o.CostPerHour) {
		var ret float32
		return ret
	}
	return *o.CostPerHour
}

// GetCostPerHourOk returns a tuple with the CostPerHour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeOverviewFields) GetCostPerHourOk() (*float32, bool) {
	if o == nil || IsNil(o.CostPerHour) {
		return nil, false
	}
	return o.CostPerHour, true
}

// HasCostPerHour returns a boolean if a field has been set.
func (o *VolumeOverviewFields) HasCostPerHour() bool {
	if o != nil && !IsNil(o.CostPerHour) {
		return true
	}

	return false
}

// SetCostPerHour gets a reference to the given float32 and assigns it to the CostPerHour field.
func (o *VolumeOverviewFields) SetCostPerHour(v float32) {
	o.CostPerHour = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *VolumeOverviewFields) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeOverviewFields) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *VolumeOverviewFields) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *VolumeOverviewFields) SetCount(v int32) {
	o.Count = &v
}

// GetUsing returns the Using field value if set, zero value otherwise.
func (o *VolumeOverviewFields) GetUsing() int32 {
	if o == nil || IsNil(o.Using) {
		var ret int32
		return ret
	}
	return *o.Using
}

// GetUsingOk returns a tuple with the Using field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeOverviewFields) GetUsingOk() (*int32, bool) {
	if o == nil || IsNil(o.Using) {
		return nil, false
	}
	return o.Using, true
}

// HasUsing returns a boolean if a field has been set.
func (o *VolumeOverviewFields) HasUsing() bool {
	if o != nil && !IsNil(o.Using) {
		return true
	}

	return false
}

// SetUsing gets a reference to the given int32 and assigns it to the Using field.
func (o *VolumeOverviewFields) SetUsing(v int32) {
	o.Using = &v
}

func (o VolumeOverviewFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeOverviewFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CostPerHour) {
		toSerialize["cost_per_hour"] = o.CostPerHour
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Using) {
		toSerialize["using"] = o.Using
	}
	return toSerialize, nil
}

type NullableVolumeOverviewFields struct {
	value *VolumeOverviewFields
	isSet bool
}

func (v NullableVolumeOverviewFields) Get() *VolumeOverviewFields {
	return v.value
}

func (v *NullableVolumeOverviewFields) Set(val *VolumeOverviewFields) {
	v.value = val
	v.isSet = true
}

func (v NullableVolumeOverviewFields) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeOverviewFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeOverviewFields(val *VolumeOverviewFields) *NullableVolumeOverviewFields {
	return &NullableVolumeOverviewFields{value: val, isSet: true}
}

func (v NullableVolumeOverviewFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeOverviewFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
