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

// checks if the ResourceLevelBillingVMDetailsResources type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceLevelBillingVMDetailsResources{}

// ResourceLevelBillingVMDetailsResources struct for ResourceLevelBillingVMDetailsResources
type ResourceLevelBillingVMDetailsResources struct {
	Attributes *ResourceLevelBillingDetailsAttributes `json:"attributes,omitempty"`
	Metrics    *ResourceLevelBillingDetailsMetrics    `json:"metrics,omitempty"`
}

// NewResourceLevelBillingVMDetailsResources instantiates a new ResourceLevelBillingVMDetailsResources object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceLevelBillingVMDetailsResources() *ResourceLevelBillingVMDetailsResources {
	this := ResourceLevelBillingVMDetailsResources{}
	return &this
}

// NewResourceLevelBillingVMDetailsResourcesWithDefaults instantiates a new ResourceLevelBillingVMDetailsResources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceLevelBillingVMDetailsResourcesWithDefaults() *ResourceLevelBillingVMDetailsResources {
	this := ResourceLevelBillingVMDetailsResources{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ResourceLevelBillingVMDetailsResources) GetAttributes() ResourceLevelBillingDetailsAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ResourceLevelBillingDetailsAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLevelBillingVMDetailsResources) GetAttributesOk() (*ResourceLevelBillingDetailsAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ResourceLevelBillingVMDetailsResources) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ResourceLevelBillingDetailsAttributes and assigns it to the Attributes field.
func (o *ResourceLevelBillingVMDetailsResources) SetAttributes(v ResourceLevelBillingDetailsAttributes) {
	o.Attributes = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ResourceLevelBillingVMDetailsResources) GetMetrics() ResourceLevelBillingDetailsMetrics {
	if o == nil || IsNil(o.Metrics) {
		var ret ResourceLevelBillingDetailsMetrics
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLevelBillingVMDetailsResources) GetMetricsOk() (*ResourceLevelBillingDetailsMetrics, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ResourceLevelBillingVMDetailsResources) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given ResourceLevelBillingDetailsMetrics and assigns it to the Metrics field.
func (o *ResourceLevelBillingVMDetailsResources) SetMetrics(v ResourceLevelBillingDetailsMetrics) {
	o.Metrics = &v
}

func (o ResourceLevelBillingVMDetailsResources) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceLevelBillingVMDetailsResources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	return toSerialize, nil
}

type NullableResourceLevelBillingVMDetailsResources struct {
	value *ResourceLevelBillingVMDetailsResources
	isSet bool
}

func (v NullableResourceLevelBillingVMDetailsResources) Get() *ResourceLevelBillingVMDetailsResources {
	return v.value
}

func (v *NullableResourceLevelBillingVMDetailsResources) Set(val *ResourceLevelBillingVMDetailsResources) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceLevelBillingVMDetailsResources) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceLevelBillingVMDetailsResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceLevelBillingVMDetailsResources(val *ResourceLevelBillingVMDetailsResources) *NullableResourceLevelBillingVMDetailsResources {
	return &NullableResourceLevelBillingVMDetailsResources{value: val, isSet: true}
}

func (v NullableResourceLevelBillingVMDetailsResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceLevelBillingVMDetailsResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
