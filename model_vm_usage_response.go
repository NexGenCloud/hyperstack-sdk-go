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

// checks if the VMUsageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VMUsageResponse{}

// VMUsageResponse struct for VMUsageResponse
type VMUsageResponse struct {
	Message         *string               `json:"message,omitempty"`
	OrgId           *int32                `json:"org_id,omitempty"`
	Status          *bool                 `json:"status,omitempty"`
	VirtualMachines []VirtualMachineUsage `json:"virtual_machines,omitempty"`
}

// NewVMUsageResponse instantiates a new VMUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVMUsageResponse() *VMUsageResponse {
	this := VMUsageResponse{}
	return &this
}

// NewVMUsageResponseWithDefaults instantiates a new VMUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVMUsageResponseWithDefaults() *VMUsageResponse {
	this := VMUsageResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *VMUsageResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsageResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *VMUsageResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *VMUsageResponse) SetMessage(v string) {
	o.Message = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *VMUsageResponse) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsageResponse) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *VMUsageResponse) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *VMUsageResponse) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VMUsageResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsageResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VMUsageResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *VMUsageResponse) SetStatus(v bool) {
	o.Status = &v
}

// GetVirtualMachines returns the VirtualMachines field value if set, zero value otherwise.
func (o *VMUsageResponse) GetVirtualMachines() []VirtualMachineUsage {
	if o == nil || IsNil(o.VirtualMachines) {
		var ret []VirtualMachineUsage
		return ret
	}
	return o.VirtualMachines
}

// GetVirtualMachinesOk returns a tuple with the VirtualMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VMUsageResponse) GetVirtualMachinesOk() ([]VirtualMachineUsage, bool) {
	if o == nil || IsNil(o.VirtualMachines) {
		return nil, false
	}
	return o.VirtualMachines, true
}

// HasVirtualMachines returns a boolean if a field has been set.
func (o *VMUsageResponse) HasVirtualMachines() bool {
	if o != nil && !IsNil(o.VirtualMachines) {
		return true
	}

	return false
}

// SetVirtualMachines gets a reference to the given []VirtualMachineUsage and assigns it to the VirtualMachines field.
func (o *VMUsageResponse) SetVirtualMachines(v []VirtualMachineUsage) {
	o.VirtualMachines = v
}

func (o VMUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VMUsageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.VirtualMachines) {
		toSerialize["virtual_machines"] = o.VirtualMachines
	}
	return toSerialize, nil
}

type NullableVMUsageResponse struct {
	value *VMUsageResponse
	isSet bool
}

func (v NullableVMUsageResponse) Get() *VMUsageResponse {
	return v.value
}

func (v *NullableVMUsageResponse) Set(val *VMUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVMUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVMUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVMUsageResponse(val *VMUsageResponse) *NullableVMUsageResponse {
	return &NullableVMUsageResponse{value: val, isSet: true}
}

func (v NullableVMUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVMUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
