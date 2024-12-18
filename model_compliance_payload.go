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

// checks if the CompliancePayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompliancePayload{}

// CompliancePayload struct for CompliancePayload
type CompliancePayload struct {
	BaseValue     int32  `json:"base_value"`
	GpuModel      string `json:"gpu_model"`
	ResourceType  string `json:"resource_type"`
	VariationMax  int32  `json:"variation_max"`
	VariationMin  int32  `json:"variation_min"`
	VariationUnit int32  `json:"variation_unit"`
}

type _CompliancePayload CompliancePayload

// NewCompliancePayload instantiates a new CompliancePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompliancePayload(baseValue int32, gpuModel string, resourceType string, variationMax int32, variationMin int32, variationUnit int32) *CompliancePayload {
	this := CompliancePayload{}
	this.BaseValue = baseValue
	this.GpuModel = gpuModel
	this.ResourceType = resourceType
	this.VariationMax = variationMax
	this.VariationMin = variationMin
	this.VariationUnit = variationUnit
	return &this
}

// NewCompliancePayloadWithDefaults instantiates a new CompliancePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompliancePayloadWithDefaults() *CompliancePayload {
	this := CompliancePayload{}
	return &this
}

// GetBaseValue returns the BaseValue field value
func (o *CompliancePayload) GetBaseValue() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BaseValue
}

// GetBaseValueOk returns a tuple with the BaseValue field value
// and a boolean to check if the value has been set.
func (o *CompliancePayload) GetBaseValueOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BaseValue, true
}

// SetBaseValue sets field value
func (o *CompliancePayload) SetBaseValue(v int32) {
	o.BaseValue = v
}

// GetGpuModel returns the GpuModel field value
func (o *CompliancePayload) GetGpuModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GpuModel
}

// GetGpuModelOk returns a tuple with the GpuModel field value
// and a boolean to check if the value has been set.
func (o *CompliancePayload) GetGpuModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GpuModel, true
}

// SetGpuModel sets field value
func (o *CompliancePayload) SetGpuModel(v string) {
	o.GpuModel = v
}

// GetResourceType returns the ResourceType field value
func (o *CompliancePayload) GetResourceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value
// and a boolean to check if the value has been set.
func (o *CompliancePayload) GetResourceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceType, true
}

// SetResourceType sets field value
func (o *CompliancePayload) SetResourceType(v string) {
	o.ResourceType = v
}

// GetVariationMax returns the VariationMax field value
func (o *CompliancePayload) GetVariationMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VariationMax
}

// GetVariationMaxOk returns a tuple with the VariationMax field value
// and a boolean to check if the value has been set.
func (o *CompliancePayload) GetVariationMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariationMax, true
}

// SetVariationMax sets field value
func (o *CompliancePayload) SetVariationMax(v int32) {
	o.VariationMax = v
}

// GetVariationMin returns the VariationMin field value
func (o *CompliancePayload) GetVariationMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VariationMin
}

// GetVariationMinOk returns a tuple with the VariationMin field value
// and a boolean to check if the value has been set.
func (o *CompliancePayload) GetVariationMinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariationMin, true
}

// SetVariationMin sets field value
func (o *CompliancePayload) SetVariationMin(v int32) {
	o.VariationMin = v
}

// GetVariationUnit returns the VariationUnit field value
func (o *CompliancePayload) GetVariationUnit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VariationUnit
}

// GetVariationUnitOk returns a tuple with the VariationUnit field value
// and a boolean to check if the value has been set.
func (o *CompliancePayload) GetVariationUnitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariationUnit, true
}

// SetVariationUnit sets field value
func (o *CompliancePayload) SetVariationUnit(v int32) {
	o.VariationUnit = v
}

func (o CompliancePayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompliancePayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["base_value"] = o.BaseValue
	toSerialize["gpu_model"] = o.GpuModel
	toSerialize["resource_type"] = o.ResourceType
	toSerialize["variation_max"] = o.VariationMax
	toSerialize["variation_min"] = o.VariationMin
	toSerialize["variation_unit"] = o.VariationUnit
	return toSerialize, nil
}

func (o *CompliancePayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"base_value",
		"gpu_model",
		"resource_type",
		"variation_max",
		"variation_min",
		"variation_unit",
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

	varCompliancePayload := _CompliancePayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCompliancePayload)

	if err != nil {
		return err
	}

	*o = CompliancePayload(varCompliancePayload)

	return err
}

type NullableCompliancePayload struct {
	value *CompliancePayload
	isSet bool
}

func (v NullableCompliancePayload) Get() *CompliancePayload {
	return v.value
}

func (v *NullableCompliancePayload) Set(val *CompliancePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCompliancePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCompliancePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompliancePayload(val *CompliancePayload) *NullableCompliancePayload {
	return &NullableCompliancePayload{value: val, isSet: true}
}

func (v NullableCompliancePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompliancePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
