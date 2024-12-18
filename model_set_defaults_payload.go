/*
Infrahub-API

Leverage the Infrahub API and Hyperstack platform to easily create, manage, and scale powerful GPU virtual machines and their associated resources.   Access this SDK to automate the deployment of your workloads and streamline your infrastructure management.  To contribute, please raise an issue with a bug report, feature request, feedback, or general inquiry.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SetDefaultsPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetDefaultsPayload{}

// SetDefaultsPayload struct for SetDefaultsPayload
type SetDefaultsPayload struct {
	Flavors []int32 `json:"flavors"`
	Images  []int32 `json:"images"`
}

type _SetDefaultsPayload SetDefaultsPayload

// NewSetDefaultsPayload instantiates a new SetDefaultsPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetDefaultsPayload(flavors []int32, images []int32) *SetDefaultsPayload {
	this := SetDefaultsPayload{}
	this.Flavors = flavors
	this.Images = images
	return &this
}

// NewSetDefaultsPayloadWithDefaults instantiates a new SetDefaultsPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetDefaultsPayloadWithDefaults() *SetDefaultsPayload {
	this := SetDefaultsPayload{}
	return &this
}

// GetFlavors returns the Flavors field value
func (o *SetDefaultsPayload) GetFlavors() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Flavors
}

// GetFlavorsOk returns a tuple with the Flavors field value
// and a boolean to check if the value has been set.
func (o *SetDefaultsPayload) GetFlavorsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Flavors, true
}

// SetFlavors sets field value
func (o *SetDefaultsPayload) SetFlavors(v []int32) {
	o.Flavors = v
}

// GetImages returns the Images field value
func (o *SetDefaultsPayload) GetImages() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Images
}

// GetImagesOk returns a tuple with the Images field value
// and a boolean to check if the value has been set.
func (o *SetDefaultsPayload) GetImagesOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Images, true
}

// SetImages sets field value
func (o *SetDefaultsPayload) SetImages(v []int32) {
	o.Images = v
}

func (o SetDefaultsPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetDefaultsPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["flavors"] = o.Flavors
	toSerialize["images"] = o.Images
	return toSerialize, nil
}

func (o *SetDefaultsPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"flavors",
		"images",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSetDefaultsPayload := _SetDefaultsPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetDefaultsPayload)

	if err != nil {
		return err
	}

	*o = SetDefaultsPayload(varSetDefaultsPayload)

	return err
}

type NullableSetDefaultsPayload struct {
	value *SetDefaultsPayload
	isSet bool
}

func (v NullableSetDefaultsPayload) Get() *SetDefaultsPayload {
	return v.value
}

func (v *NullableSetDefaultsPayload) Set(val *SetDefaultsPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSetDefaultsPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSetDefaultsPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetDefaultsPayload(val *SetDefaultsPayload) *NullableSetDefaultsPayload {
	return &NullableSetDefaultsPayload{value: val, isSet: true}
}

func (v NullableSetDefaultsPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetDefaultsPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
