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

// checks if the FirewallAttachmentModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallAttachmentModel{}

// FirewallAttachmentModel struct for FirewallAttachmentModel
type FirewallAttachmentModel struct {
	CreatedAt *CustomTime                `json:"created_at,omitempty"`
	Id        *int32                     `json:"id,omitempty"`
	Status    *string                    `json:"status,omitempty"`
	Vm        *FirewallAttachmentVMModel `json:"vm,omitempty"`
}

// NewFirewallAttachmentModel instantiates a new FirewallAttachmentModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallAttachmentModel() *FirewallAttachmentModel {
	this := FirewallAttachmentModel{}
	return &this
}

// NewFirewallAttachmentModelWithDefaults instantiates a new FirewallAttachmentModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallAttachmentModelWithDefaults() *FirewallAttachmentModel {
	this := FirewallAttachmentModel{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FirewallAttachmentModel) GetCreatedAt() CustomTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret CustomTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentModel) GetCreatedAtOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FirewallAttachmentModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given CustomTime and assigns it to the CreatedAt field.
func (o *FirewallAttachmentModel) SetCreatedAt(v CustomTime) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FirewallAttachmentModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FirewallAttachmentModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FirewallAttachmentModel) SetId(v int32) {
	o.Id = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FirewallAttachmentModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FirewallAttachmentModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FirewallAttachmentModel) SetStatus(v string) {
	o.Status = &v
}

// GetVm returns the Vm field value if set, zero value otherwise.
func (o *FirewallAttachmentModel) GetVm() FirewallAttachmentVMModel {
	if o == nil || IsNil(o.Vm) {
		var ret FirewallAttachmentVMModel
		return ret
	}
	return *o.Vm
}

// GetVmOk returns a tuple with the Vm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentModel) GetVmOk() (*FirewallAttachmentVMModel, bool) {
	if o == nil || IsNil(o.Vm) {
		return nil, false
	}
	return o.Vm, true
}

// HasVm returns a boolean if a field has been set.
func (o *FirewallAttachmentModel) HasVm() bool {
	if o != nil && !IsNil(o.Vm) {
		return true
	}

	return false
}

// SetVm gets a reference to the given FirewallAttachmentVMModel and assigns it to the Vm field.
func (o *FirewallAttachmentModel) SetVm(v FirewallAttachmentVMModel) {
	o.Vm = &v
}

func (o FirewallAttachmentModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallAttachmentModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Vm) {
		toSerialize["vm"] = o.Vm
	}
	return toSerialize, nil
}

type NullableFirewallAttachmentModel struct {
	value *FirewallAttachmentModel
	isSet bool
}

func (v NullableFirewallAttachmentModel) Get() *FirewallAttachmentModel {
	return v.value
}

func (v *NullableFirewallAttachmentModel) Set(val *FirewallAttachmentModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallAttachmentModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallAttachmentModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallAttachmentModel(val *FirewallAttachmentModel) *NullableFirewallAttachmentModel {
	return &NullableFirewallAttachmentModel{value: val, isSet: true}
}

func (v NullableFirewallAttachmentModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallAttachmentModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
