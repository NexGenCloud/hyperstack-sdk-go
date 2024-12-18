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

// checks if the ContractBillingHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContractBillingHistory{}

// ContractBillingHistory struct for ContractBillingHistory
type ContractBillingHistory struct {
	Attributes *ContractBillingHistoryResponseAttributes `json:"attributes,omitempty"`
	Metrics    *ContractlBillingHistoryResponseMetrics   `json:"metrics,omitempty"`
}

// NewContractBillingHistory instantiates a new ContractBillingHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractBillingHistory() *ContractBillingHistory {
	this := ContractBillingHistory{}
	return &this
}

// NewContractBillingHistoryWithDefaults instantiates a new ContractBillingHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractBillingHistoryWithDefaults() *ContractBillingHistory {
	this := ContractBillingHistory{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ContractBillingHistory) GetAttributes() ContractBillingHistoryResponseAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ContractBillingHistoryResponseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractBillingHistory) GetAttributesOk() (*ContractBillingHistoryResponseAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ContractBillingHistory) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ContractBillingHistoryResponseAttributes and assigns it to the Attributes field.
func (o *ContractBillingHistory) SetAttributes(v ContractBillingHistoryResponseAttributes) {
	o.Attributes = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *ContractBillingHistory) GetMetrics() ContractlBillingHistoryResponseMetrics {
	if o == nil || IsNil(o.Metrics) {
		var ret ContractlBillingHistoryResponseMetrics
		return ret
	}
	return *o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractBillingHistory) GetMetricsOk() (*ContractlBillingHistoryResponseMetrics, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *ContractBillingHistory) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given ContractlBillingHistoryResponseMetrics and assigns it to the Metrics field.
func (o *ContractBillingHistory) SetMetrics(v ContractlBillingHistoryResponseMetrics) {
	o.Metrics = &v
}

func (o ContractBillingHistory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContractBillingHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	return toSerialize, nil
}

type NullableContractBillingHistory struct {
	value *ContractBillingHistory
	isSet bool
}

func (v NullableContractBillingHistory) Get() *ContractBillingHistory {
	return v.value
}

func (v *NullableContractBillingHistory) Set(val *ContractBillingHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableContractBillingHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableContractBillingHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractBillingHistory(val *ContractBillingHistory) *NullableContractBillingHistory {
	return &NullableContractBillingHistory{value: val, isSet: true}
}

func (v NullableContractBillingHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractBillingHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
