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

// checks if the OrganizationLevelBillingHistoryResponseMetrics type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationLevelBillingHistoryResponseMetrics{}

// OrganizationLevelBillingHistoryResponseMetrics struct for OrganizationLevelBillingHistoryResponseMetrics
type OrganizationLevelBillingHistoryResponseMetrics struct {
	ContractCost      *float32 `json:"contract_cost,omitempty"`
	IncurredBill      *float32 `json:"incurred_bill,omitempty"`
	NonDiscountedBill *float32 `json:"non_discounted_bill,omitempty"`
	SnapshotCost      *float32 `json:"snapshot_cost,omitempty"`
	VmCost            *float32 `json:"vm_cost,omitempty"`
	VolumeCost        *float32 `json:"volume_cost,omitempty"`
}

// NewOrganizationLevelBillingHistoryResponseMetrics instantiates a new OrganizationLevelBillingHistoryResponseMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationLevelBillingHistoryResponseMetrics() *OrganizationLevelBillingHistoryResponseMetrics {
	this := OrganizationLevelBillingHistoryResponseMetrics{}
	return &this
}

// NewOrganizationLevelBillingHistoryResponseMetricsWithDefaults instantiates a new OrganizationLevelBillingHistoryResponseMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationLevelBillingHistoryResponseMetricsWithDefaults() *OrganizationLevelBillingHistoryResponseMetrics {
	this := OrganizationLevelBillingHistoryResponseMetrics{}
	return &this
}

// GetContractCost returns the ContractCost field value if set, zero value otherwise.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetContractCost() float32 {
	if o == nil || IsNil(o.ContractCost) {
		var ret float32
		return ret
	}
	return *o.ContractCost
}

// GetContractCostOk returns a tuple with the ContractCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetContractCostOk() (*float32, bool) {
	if o == nil || IsNil(o.ContractCost) {
		return nil, false
	}
	return o.ContractCost, true
}

// HasContractCost returns a boolean if a field has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) HasContractCost() bool {
	if o != nil && !IsNil(o.ContractCost) {
		return true
	}

	return false
}

// SetContractCost gets a reference to the given float32 and assigns it to the ContractCost field.
func (o *OrganizationLevelBillingHistoryResponseMetrics) SetContractCost(v float32) {
	o.ContractCost = &v
}

// GetIncurredBill returns the IncurredBill field value if set, zero value otherwise.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetIncurredBill() float32 {
	if o == nil || IsNil(o.IncurredBill) {
		var ret float32
		return ret
	}
	return *o.IncurredBill
}

// GetIncurredBillOk returns a tuple with the IncurredBill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetIncurredBillOk() (*float32, bool) {
	if o == nil || IsNil(o.IncurredBill) {
		return nil, false
	}
	return o.IncurredBill, true
}

// HasIncurredBill returns a boolean if a field has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) HasIncurredBill() bool {
	if o != nil && !IsNil(o.IncurredBill) {
		return true
	}

	return false
}

// SetIncurredBill gets a reference to the given float32 and assigns it to the IncurredBill field.
func (o *OrganizationLevelBillingHistoryResponseMetrics) SetIncurredBill(v float32) {
	o.IncurredBill = &v
}

// GetNonDiscountedBill returns the NonDiscountedBill field value if set, zero value otherwise.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetNonDiscountedBill() float32 {
	if o == nil || IsNil(o.NonDiscountedBill) {
		var ret float32
		return ret
	}
	return *o.NonDiscountedBill
}

// GetNonDiscountedBillOk returns a tuple with the NonDiscountedBill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetNonDiscountedBillOk() (*float32, bool) {
	if o == nil || IsNil(o.NonDiscountedBill) {
		return nil, false
	}
	return o.NonDiscountedBill, true
}

// HasNonDiscountedBill returns a boolean if a field has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) HasNonDiscountedBill() bool {
	if o != nil && !IsNil(o.NonDiscountedBill) {
		return true
	}

	return false
}

// SetNonDiscountedBill gets a reference to the given float32 and assigns it to the NonDiscountedBill field.
func (o *OrganizationLevelBillingHistoryResponseMetrics) SetNonDiscountedBill(v float32) {
	o.NonDiscountedBill = &v
}

// GetSnapshotCost returns the SnapshotCost field value if set, zero value otherwise.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetSnapshotCost() float32 {
	if o == nil || IsNil(o.SnapshotCost) {
		var ret float32
		return ret
	}
	return *o.SnapshotCost
}

// GetSnapshotCostOk returns a tuple with the SnapshotCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetSnapshotCostOk() (*float32, bool) {
	if o == nil || IsNil(o.SnapshotCost) {
		return nil, false
	}
	return o.SnapshotCost, true
}

// HasSnapshotCost returns a boolean if a field has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) HasSnapshotCost() bool {
	if o != nil && !IsNil(o.SnapshotCost) {
		return true
	}

	return false
}

// SetSnapshotCost gets a reference to the given float32 and assigns it to the SnapshotCost field.
func (o *OrganizationLevelBillingHistoryResponseMetrics) SetSnapshotCost(v float32) {
	o.SnapshotCost = &v
}

// GetVmCost returns the VmCost field value if set, zero value otherwise.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetVmCost() float32 {
	if o == nil || IsNil(o.VmCost) {
		var ret float32
		return ret
	}
	return *o.VmCost
}

// GetVmCostOk returns a tuple with the VmCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetVmCostOk() (*float32, bool) {
	if o == nil || IsNil(o.VmCost) {
		return nil, false
	}
	return o.VmCost, true
}

// HasVmCost returns a boolean if a field has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) HasVmCost() bool {
	if o != nil && !IsNil(o.VmCost) {
		return true
	}

	return false
}

// SetVmCost gets a reference to the given float32 and assigns it to the VmCost field.
func (o *OrganizationLevelBillingHistoryResponseMetrics) SetVmCost(v float32) {
	o.VmCost = &v
}

// GetVolumeCost returns the VolumeCost field value if set, zero value otherwise.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetVolumeCost() float32 {
	if o == nil || IsNil(o.VolumeCost) {
		var ret float32
		return ret
	}
	return *o.VolumeCost
}

// GetVolumeCostOk returns a tuple with the VolumeCost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) GetVolumeCostOk() (*float32, bool) {
	if o == nil || IsNil(o.VolumeCost) {
		return nil, false
	}
	return o.VolumeCost, true
}

// HasVolumeCost returns a boolean if a field has been set.
func (o *OrganizationLevelBillingHistoryResponseMetrics) HasVolumeCost() bool {
	if o != nil && !IsNil(o.VolumeCost) {
		return true
	}

	return false
}

// SetVolumeCost gets a reference to the given float32 and assigns it to the VolumeCost field.
func (o *OrganizationLevelBillingHistoryResponseMetrics) SetVolumeCost(v float32) {
	o.VolumeCost = &v
}

func (o OrganizationLevelBillingHistoryResponseMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationLevelBillingHistoryResponseMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractCost) {
		toSerialize["contract_cost"] = o.ContractCost
	}
	if !IsNil(o.IncurredBill) {
		toSerialize["incurred_bill"] = o.IncurredBill
	}
	if !IsNil(o.NonDiscountedBill) {
		toSerialize["non_discounted_bill"] = o.NonDiscountedBill
	}
	if !IsNil(o.SnapshotCost) {
		toSerialize["snapshot_cost"] = o.SnapshotCost
	}
	if !IsNil(o.VmCost) {
		toSerialize["vm_cost"] = o.VmCost
	}
	if !IsNil(o.VolumeCost) {
		toSerialize["volume_cost"] = o.VolumeCost
	}
	return toSerialize, nil
}

type NullableOrganizationLevelBillingHistoryResponseMetrics struct {
	value *OrganizationLevelBillingHistoryResponseMetrics
	isSet bool
}

func (v NullableOrganizationLevelBillingHistoryResponseMetrics) Get() *OrganizationLevelBillingHistoryResponseMetrics {
	return v.value
}

func (v *NullableOrganizationLevelBillingHistoryResponseMetrics) Set(val *OrganizationLevelBillingHistoryResponseMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationLevelBillingHistoryResponseMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationLevelBillingHistoryResponseMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationLevelBillingHistoryResponseMetrics(val *OrganizationLevelBillingHistoryResponseMetrics) *NullableOrganizationLevelBillingHistoryResponseMetrics {
	return &NullableOrganizationLevelBillingHistoryResponseMetrics{value: val, isSet: true}
}

func (v NullableOrganizationLevelBillingHistoryResponseMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationLevelBillingHistoryResponseMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
