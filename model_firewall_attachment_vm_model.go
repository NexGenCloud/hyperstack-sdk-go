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

// checks if the FirewallAttachmentVMModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FirewallAttachmentVMModel{}

// FirewallAttachmentVMModel struct for FirewallAttachmentVMModel
type FirewallAttachmentVMModel struct {
	CreatedAt   *CustomTime `json:"created_at,omitempty"`
	Environment *string     `json:"environment,omitempty"`
	Flavor      *string     `json:"flavor,omitempty"`
	Id          *int32      `json:"id,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Status      *string     `json:"status,omitempty"`
}

// NewFirewallAttachmentVMModel instantiates a new FirewallAttachmentVMModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirewallAttachmentVMModel() *FirewallAttachmentVMModel {
	this := FirewallAttachmentVMModel{}
	return &this
}

// NewFirewallAttachmentVMModelWithDefaults instantiates a new FirewallAttachmentVMModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirewallAttachmentVMModelWithDefaults() *FirewallAttachmentVMModel {
	this := FirewallAttachmentVMModel{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *FirewallAttachmentVMModel) GetCreatedAt() CustomTime {
	if o == nil || IsNil(o.CreatedAt) {
		var ret CustomTime
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentVMModel) GetCreatedAtOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *FirewallAttachmentVMModel) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given CustomTime and assigns it to the CreatedAt field.
func (o *FirewallAttachmentVMModel) SetCreatedAt(v CustomTime) {
	o.CreatedAt = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *FirewallAttachmentVMModel) GetEnvironment() string {
	if o == nil || IsNil(o.Environment) {
		var ret string
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentVMModel) GetEnvironmentOk() (*string, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *FirewallAttachmentVMModel) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given string and assigns it to the Environment field.
func (o *FirewallAttachmentVMModel) SetEnvironment(v string) {
	o.Environment = &v
}

// GetFlavor returns the Flavor field value if set, zero value otherwise.
func (o *FirewallAttachmentVMModel) GetFlavor() string {
	if o == nil || IsNil(o.Flavor) {
		var ret string
		return ret
	}
	return *o.Flavor
}

// GetFlavorOk returns a tuple with the Flavor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentVMModel) GetFlavorOk() (*string, bool) {
	if o == nil || IsNil(o.Flavor) {
		return nil, false
	}
	return o.Flavor, true
}

// HasFlavor returns a boolean if a field has been set.
func (o *FirewallAttachmentVMModel) HasFlavor() bool {
	if o != nil && !IsNil(o.Flavor) {
		return true
	}

	return false
}

// SetFlavor gets a reference to the given string and assigns it to the Flavor field.
func (o *FirewallAttachmentVMModel) SetFlavor(v string) {
	o.Flavor = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FirewallAttachmentVMModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentVMModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FirewallAttachmentVMModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FirewallAttachmentVMModel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FirewallAttachmentVMModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentVMModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FirewallAttachmentVMModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FirewallAttachmentVMModel) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *FirewallAttachmentVMModel) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirewallAttachmentVMModel) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *FirewallAttachmentVMModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *FirewallAttachmentVMModel) SetStatus(v string) {
	o.Status = &v
}

func (o FirewallAttachmentVMModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FirewallAttachmentVMModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Flavor) {
		toSerialize["flavor"] = o.Flavor
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableFirewallAttachmentVMModel struct {
	value *FirewallAttachmentVMModel
	isSet bool
}

func (v NullableFirewallAttachmentVMModel) Get() *FirewallAttachmentVMModel {
	return v.value
}

func (v *NullableFirewallAttachmentVMModel) Set(val *FirewallAttachmentVMModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFirewallAttachmentVMModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFirewallAttachmentVMModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirewallAttachmentVMModel(val *FirewallAttachmentVMModel) *NullableFirewallAttachmentVMModel {
	return &NullableFirewallAttachmentVMModel{value: val, isSet: true}
}

func (v NullableFirewallAttachmentVMModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirewallAttachmentVMModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
