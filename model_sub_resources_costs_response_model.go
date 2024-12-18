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

// checks if the SubResourcesCostsResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubResourcesCostsResponseModel{}

// SubResourcesCostsResponseModel struct for SubResourcesCostsResponseModel
type SubResourcesCostsResponseModel struct {
	BillingHistory []SubResourcesGraphBillingHistoryFields `json:"billing_history,omitempty"`
	Granularity    *int32                                  `json:"granularity,omitempty"`
	OrgId          *int32                                  `json:"org_id,omitempty"`
	TotalCount     *int32                                  `json:"total_count,omitempty"`
}

// NewSubResourcesCostsResponseModel instantiates a new SubResourcesCostsResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubResourcesCostsResponseModel() *SubResourcesCostsResponseModel {
	this := SubResourcesCostsResponseModel{}
	return &this
}

// NewSubResourcesCostsResponseModelWithDefaults instantiates a new SubResourcesCostsResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubResourcesCostsResponseModelWithDefaults() *SubResourcesCostsResponseModel {
	this := SubResourcesCostsResponseModel{}
	return &this
}

// GetBillingHistory returns the BillingHistory field value if set, zero value otherwise.
func (o *SubResourcesCostsResponseModel) GetBillingHistory() []SubResourcesGraphBillingHistoryFields {
	if o == nil || IsNil(o.BillingHistory) {
		var ret []SubResourcesGraphBillingHistoryFields
		return ret
	}
	return o.BillingHistory
}

// GetBillingHistoryOk returns a tuple with the BillingHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubResourcesCostsResponseModel) GetBillingHistoryOk() ([]SubResourcesGraphBillingHistoryFields, bool) {
	if o == nil || IsNil(o.BillingHistory) {
		return nil, false
	}
	return o.BillingHistory, true
}

// HasBillingHistory returns a boolean if a field has been set.
func (o *SubResourcesCostsResponseModel) HasBillingHistory() bool {
	if o != nil && !IsNil(o.BillingHistory) {
		return true
	}

	return false
}

// SetBillingHistory gets a reference to the given []SubResourcesGraphBillingHistoryFields and assigns it to the BillingHistory field.
func (o *SubResourcesCostsResponseModel) SetBillingHistory(v []SubResourcesGraphBillingHistoryFields) {
	o.BillingHistory = v
}

// GetGranularity returns the Granularity field value if set, zero value otherwise.
func (o *SubResourcesCostsResponseModel) GetGranularity() int32 {
	if o == nil || IsNil(o.Granularity) {
		var ret int32
		return ret
	}
	return *o.Granularity
}

// GetGranularityOk returns a tuple with the Granularity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubResourcesCostsResponseModel) GetGranularityOk() (*int32, bool) {
	if o == nil || IsNil(o.Granularity) {
		return nil, false
	}
	return o.Granularity, true
}

// HasGranularity returns a boolean if a field has been set.
func (o *SubResourcesCostsResponseModel) HasGranularity() bool {
	if o != nil && !IsNil(o.Granularity) {
		return true
	}

	return false
}

// SetGranularity gets a reference to the given int32 and assigns it to the Granularity field.
func (o *SubResourcesCostsResponseModel) SetGranularity(v int32) {
	o.Granularity = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *SubResourcesCostsResponseModel) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubResourcesCostsResponseModel) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *SubResourcesCostsResponseModel) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *SubResourcesCostsResponseModel) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *SubResourcesCostsResponseModel) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubResourcesCostsResponseModel) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *SubResourcesCostsResponseModel) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *SubResourcesCostsResponseModel) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o SubResourcesCostsResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubResourcesCostsResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BillingHistory) {
		toSerialize["billing_history"] = o.BillingHistory
	}
	if !IsNil(o.Granularity) {
		toSerialize["granularity"] = o.Granularity
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableSubResourcesCostsResponseModel struct {
	value *SubResourcesCostsResponseModel
	isSet bool
}

func (v NullableSubResourcesCostsResponseModel) Get() *SubResourcesCostsResponseModel {
	return v.value
}

func (v *NullableSubResourcesCostsResponseModel) Set(val *SubResourcesCostsResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableSubResourcesCostsResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableSubResourcesCostsResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubResourcesCostsResponseModel(val *SubResourcesCostsResponseModel) *NullableSubResourcesCostsResponseModel {
	return &NullableSubResourcesCostsResponseModel{value: val, isSet: true}
}

func (v NullableSubResourcesCostsResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubResourcesCostsResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
