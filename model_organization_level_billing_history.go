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

// checks if the OrganizationLevelBillingHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationLevelBillingHistory{}

// OrganizationLevelBillingHistory struct for OrganizationLevelBillingHistory
type OrganizationLevelBillingHistory struct {
	BillingHistory []OrganizationLevelBillingHistoryResources `json:"billing_history,omitempty"`
	OrgId          *int32                                     `json:"org_id,omitempty"`
	TotalCount     *int32                                     `json:"total_count,omitempty"`
}

// NewOrganizationLevelBillingHistory instantiates a new OrganizationLevelBillingHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationLevelBillingHistory() *OrganizationLevelBillingHistory {
	this := OrganizationLevelBillingHistory{}
	return &this
}

// NewOrganizationLevelBillingHistoryWithDefaults instantiates a new OrganizationLevelBillingHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationLevelBillingHistoryWithDefaults() *OrganizationLevelBillingHistory {
	this := OrganizationLevelBillingHistory{}
	return &this
}

// GetBillingHistory returns the BillingHistory field value if set, zero value otherwise.
func (o *OrganizationLevelBillingHistory) GetBillingHistory() []OrganizationLevelBillingHistoryResources {
	if o == nil || IsNil(o.BillingHistory) {
		var ret []OrganizationLevelBillingHistoryResources
		return ret
	}
	return o.BillingHistory
}

// GetBillingHistoryOk returns a tuple with the BillingHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLevelBillingHistory) GetBillingHistoryOk() ([]OrganizationLevelBillingHistoryResources, bool) {
	if o == nil || IsNil(o.BillingHistory) {
		return nil, false
	}
	return o.BillingHistory, true
}

// HasBillingHistory returns a boolean if a field has been set.
func (o *OrganizationLevelBillingHistory) HasBillingHistory() bool {
	if o != nil && !IsNil(o.BillingHistory) {
		return true
	}

	return false
}

// SetBillingHistory gets a reference to the given []OrganizationLevelBillingHistoryResources and assigns it to the BillingHistory field.
func (o *OrganizationLevelBillingHistory) SetBillingHistory(v []OrganizationLevelBillingHistoryResources) {
	o.BillingHistory = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrganizationLevelBillingHistory) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLevelBillingHistory) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrganizationLevelBillingHistory) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *OrganizationLevelBillingHistory) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *OrganizationLevelBillingHistory) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationLevelBillingHistory) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *OrganizationLevelBillingHistory) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *OrganizationLevelBillingHistory) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o OrganizationLevelBillingHistory) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationLevelBillingHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingHistory) {
		toSerialize["billing_history"] = o.BillingHistory
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableOrganizationLevelBillingHistory struct {
	value *OrganizationLevelBillingHistory
	isSet bool
}

func (v NullableOrganizationLevelBillingHistory) Get() *OrganizationLevelBillingHistory {
	return v.value
}

func (v *NullableOrganizationLevelBillingHistory) Set(val *OrganizationLevelBillingHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationLevelBillingHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationLevelBillingHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationLevelBillingHistory(val *OrganizationLevelBillingHistory) *NullableOrganizationLevelBillingHistory {
	return &NullableOrganizationLevelBillingHistory{value: val, isSet: true}
}

func (v NullableOrganizationLevelBillingHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationLevelBillingHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
