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

// checks if the EnvironmentFieldsforVolume type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentFieldsforVolume{}

// EnvironmentFieldsforVolume struct for EnvironmentFieldsforVolume
type EnvironmentFieldsforVolume struct {
	Name *string `json:"name,omitempty"`
}

// NewEnvironmentFieldsforVolume instantiates a new EnvironmentFieldsforVolume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentFieldsforVolume() *EnvironmentFieldsforVolume {
	this := EnvironmentFieldsforVolume{}
	return &this
}

// NewEnvironmentFieldsforVolumeWithDefaults instantiates a new EnvironmentFieldsforVolume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentFieldsforVolumeWithDefaults() *EnvironmentFieldsforVolume {
	this := EnvironmentFieldsforVolume{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EnvironmentFieldsforVolume) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentFieldsforVolume) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EnvironmentFieldsforVolume) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EnvironmentFieldsforVolume) SetName(v string) {
	o.Name = &v
}

func (o EnvironmentFieldsforVolume) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentFieldsforVolume) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableEnvironmentFieldsforVolume struct {
	value *EnvironmentFieldsforVolume
	isSet bool
}

func (v NullableEnvironmentFieldsforVolume) Get() *EnvironmentFieldsforVolume {
	return v.value
}

func (v *NullableEnvironmentFieldsforVolume) Set(val *EnvironmentFieldsforVolume) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentFieldsforVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentFieldsforVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentFieldsforVolume(val *EnvironmentFieldsforVolume) *NullableEnvironmentFieldsforVolume {
	return &NullableEnvironmentFieldsforVolume{value: val, isSet: true}
}

func (v NullableEnvironmentFieldsforVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentFieldsforVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}