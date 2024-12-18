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

// checks if the ResourceLevelBillingDetailsVolumeAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceLevelBillingDetailsVolumeAttributes{}

// ResourceLevelBillingDetailsVolumeAttributes struct for ResourceLevelBillingDetailsVolumeAttributes
type ResourceLevelBillingDetailsVolumeAttributes struct {
	Id                *string `json:"id,omitempty"`
	InfrahubId        *int32  `json:"infrahub_id,omitempty"`
	ResourceName      *string `json:"resource_name,omitempty"`
	SubresourceAmount *int32  `json:"subresource_amount,omitempty"`
	SubresourceType   *string `json:"subresource_type,omitempty"`
}

// NewResourceLevelBillingDetailsVolumeAttributes instantiates a new ResourceLevelBillingDetailsVolumeAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceLevelBillingDetailsVolumeAttributes() *ResourceLevelBillingDetailsVolumeAttributes {
	this := ResourceLevelBillingDetailsVolumeAttributes{}
	return &this
}

// NewResourceLevelBillingDetailsVolumeAttributesWithDefaults instantiates a new ResourceLevelBillingDetailsVolumeAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceLevelBillingDetailsVolumeAttributesWithDefaults() *ResourceLevelBillingDetailsVolumeAttributes {
	this := ResourceLevelBillingDetailsVolumeAttributes{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceLevelBillingDetailsVolumeAttributes) SetId(v string) {
	o.Id = &v
}

// GetInfrahubId returns the InfrahubId field value if set, zero value otherwise.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetInfrahubId() int32 {
	if o == nil || IsNil(o.InfrahubId) {
		var ret int32
		return ret
	}
	return *o.InfrahubId
}

// GetInfrahubIdOk returns a tuple with the InfrahubId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetInfrahubIdOk() (*int32, bool) {
	if o == nil || IsNil(o.InfrahubId) {
		return nil, false
	}
	return o.InfrahubId, true
}

// HasInfrahubId returns a boolean if a field has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) HasInfrahubId() bool {
	if o != nil && !IsNil(o.InfrahubId) {
		return true
	}

	return false
}

// SetInfrahubId gets a reference to the given int32 and assigns it to the InfrahubId field.
func (o *ResourceLevelBillingDetailsVolumeAttributes) SetInfrahubId(v int32) {
	o.InfrahubId = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *ResourceLevelBillingDetailsVolumeAttributes) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetSubresourceAmount returns the SubresourceAmount field value if set, zero value otherwise.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetSubresourceAmount() int32 {
	if o == nil || IsNil(o.SubresourceAmount) {
		var ret int32
		return ret
	}
	return *o.SubresourceAmount
}

// GetSubresourceAmountOk returns a tuple with the SubresourceAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetSubresourceAmountOk() (*int32, bool) {
	if o == nil || IsNil(o.SubresourceAmount) {
		return nil, false
	}
	return o.SubresourceAmount, true
}

// HasSubresourceAmount returns a boolean if a field has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) HasSubresourceAmount() bool {
	if o != nil && !IsNil(o.SubresourceAmount) {
		return true
	}

	return false
}

// SetSubresourceAmount gets a reference to the given int32 and assigns it to the SubresourceAmount field.
func (o *ResourceLevelBillingDetailsVolumeAttributes) SetSubresourceAmount(v int32) {
	o.SubresourceAmount = &v
}

// GetSubresourceType returns the SubresourceType field value if set, zero value otherwise.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetSubresourceType() string {
	if o == nil || IsNil(o.SubresourceType) {
		var ret string
		return ret
	}
	return *o.SubresourceType
}

// GetSubresourceTypeOk returns a tuple with the SubresourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) GetSubresourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.SubresourceType) {
		return nil, false
	}
	return o.SubresourceType, true
}

// HasSubresourceType returns a boolean if a field has been set.
func (o *ResourceLevelBillingDetailsVolumeAttributes) HasSubresourceType() bool {
	if o != nil && !IsNil(o.SubresourceType) {
		return true
	}

	return false
}

// SetSubresourceType gets a reference to the given string and assigns it to the SubresourceType field.
func (o *ResourceLevelBillingDetailsVolumeAttributes) SetSubresourceType(v string) {
	o.SubresourceType = &v
}

func (o ResourceLevelBillingDetailsVolumeAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceLevelBillingDetailsVolumeAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InfrahubId) {
		toSerialize["infrahub_id"] = o.InfrahubId
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resource_name"] = o.ResourceName
	}
	if !IsNil(o.SubresourceAmount) {
		toSerialize["subresource_amount"] = o.SubresourceAmount
	}
	if !IsNil(o.SubresourceType) {
		toSerialize["subresource_type"] = o.SubresourceType
	}
	return toSerialize, nil
}

type NullableResourceLevelBillingDetailsVolumeAttributes struct {
	value *ResourceLevelBillingDetailsVolumeAttributes
	isSet bool
}

func (v NullableResourceLevelBillingDetailsVolumeAttributes) Get() *ResourceLevelBillingDetailsVolumeAttributes {
	return v.value
}

func (v *NullableResourceLevelBillingDetailsVolumeAttributes) Set(val *ResourceLevelBillingDetailsVolumeAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceLevelBillingDetailsVolumeAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceLevelBillingDetailsVolumeAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceLevelBillingDetailsVolumeAttributes(val *ResourceLevelBillingDetailsVolumeAttributes) *NullableResourceLevelBillingDetailsVolumeAttributes {
	return &NullableResourceLevelBillingDetailsVolumeAttributes{value: val, isSet: true}
}

func (v NullableResourceLevelBillingDetailsVolumeAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceLevelBillingDetailsVolumeAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
