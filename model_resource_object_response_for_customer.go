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

// checks if the ResourceObjectResponseForCustomer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceObjectResponseForCustomer{}

// ResourceObjectResponseForCustomer struct for ResourceObjectResponseForCustomer
type ResourceObjectResponseForCustomer struct {
	OrgId    *int32                                     `json:"org_id,omitempty"`
	Resource *InfrahubResourceObjectResponseForCustomer `json:"resource,omitempty"`
}

// NewResourceObjectResponseForCustomer instantiates a new ResourceObjectResponseForCustomer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceObjectResponseForCustomer() *ResourceObjectResponseForCustomer {
	this := ResourceObjectResponseForCustomer{}
	return &this
}

// NewResourceObjectResponseForCustomerWithDefaults instantiates a new ResourceObjectResponseForCustomer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceObjectResponseForCustomerWithDefaults() *ResourceObjectResponseForCustomer {
	this := ResourceObjectResponseForCustomer{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *ResourceObjectResponseForCustomer) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObjectResponseForCustomer) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *ResourceObjectResponseForCustomer) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *ResourceObjectResponseForCustomer) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourceObjectResponseForCustomer) GetResource() InfrahubResourceObjectResponseForCustomer {
	if o == nil || IsNil(o.Resource) {
		var ret InfrahubResourceObjectResponseForCustomer
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceObjectResponseForCustomer) GetResourceOk() (*InfrahubResourceObjectResponseForCustomer, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourceObjectResponseForCustomer) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given InfrahubResourceObjectResponseForCustomer and assigns it to the Resource field.
func (o *ResourceObjectResponseForCustomer) SetResource(v InfrahubResourceObjectResponseForCustomer) {
	o.Resource = &v
}

func (o ResourceObjectResponseForCustomer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceObjectResponseForCustomer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullableResourceObjectResponseForCustomer struct {
	value *ResourceObjectResponseForCustomer
	isSet bool
}

func (v NullableResourceObjectResponseForCustomer) Get() *ResourceObjectResponseForCustomer {
	return v.value
}

func (v *NullableResourceObjectResponseForCustomer) Set(val *ResourceObjectResponseForCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceObjectResponseForCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceObjectResponseForCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceObjectResponseForCustomer(val *ResourceObjectResponseForCustomer) *NullableResourceObjectResponseForCustomer {
	return &NullableResourceObjectResponseForCustomer{value: val, isSet: true}
}

func (v NullableResourceObjectResponseForCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceObjectResponseForCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}