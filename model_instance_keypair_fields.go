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

// checks if the InstanceKeypairFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InstanceKeypairFields{}

// InstanceKeypairFields struct for InstanceKeypairFields
type InstanceKeypairFields struct {
	Name *string `json:"name,omitempty"`
}

// NewInstanceKeypairFields instantiates a new InstanceKeypairFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstanceKeypairFields() *InstanceKeypairFields {
	this := InstanceKeypairFields{}
	return &this
}

// NewInstanceKeypairFieldsWithDefaults instantiates a new InstanceKeypairFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstanceKeypairFieldsWithDefaults() *InstanceKeypairFields {
	this := InstanceKeypairFields{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *InstanceKeypairFields) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InstanceKeypairFields) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *InstanceKeypairFields) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *InstanceKeypairFields) SetName(v string) {
	o.Name = &v
}

func (o InstanceKeypairFields) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InstanceKeypairFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableInstanceKeypairFields struct {
	value *InstanceKeypairFields
	isSet bool
}

func (v NullableInstanceKeypairFields) Get() *InstanceKeypairFields {
	return v.value
}

func (v *NullableInstanceKeypairFields) Set(val *InstanceKeypairFields) {
	v.value = val
	v.isSet = true
}

func (v NullableInstanceKeypairFields) IsSet() bool {
	return v.isSet
}

func (v *NullableInstanceKeypairFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstanceKeypairFields(val *InstanceKeypairFields) *NullableInstanceKeypairFields {
	return &NullableInstanceKeypairFields{value: val, isSet: true}
}

func (v NullableInstanceKeypairFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstanceKeypairFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
