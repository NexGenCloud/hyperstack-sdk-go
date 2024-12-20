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

// checks if the OrganizationObjectResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationObjectResponse{}

// OrganizationObjectResponse struct for OrganizationObjectResponse
type OrganizationObjectResponse struct {
	OrgId     *int32                           `json:"org_id,omitempty"`
	Resources []InfrahubResourceObjectResponse `json:"resources,omitempty"`
}

// NewOrganizationObjectResponse instantiates a new OrganizationObjectResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationObjectResponse() *OrganizationObjectResponse {
	this := OrganizationObjectResponse{}
	return &this
}

// NewOrganizationObjectResponseWithDefaults instantiates a new OrganizationObjectResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationObjectResponseWithDefaults() *OrganizationObjectResponse {
	this := OrganizationObjectResponse{}
	return &this
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *OrganizationObjectResponse) GetOrgId() int32 {
	if o == nil || IsNil(o.OrgId) {
		var ret int32
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationObjectResponse) GetOrgIdOk() (*int32, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *OrganizationObjectResponse) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given int32 and assigns it to the OrgId field.
func (o *OrganizationObjectResponse) SetOrgId(v int32) {
	o.OrgId = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *OrganizationObjectResponse) GetResources() []InfrahubResourceObjectResponse {
	if o == nil || IsNil(o.Resources) {
		var ret []InfrahubResourceObjectResponse
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationObjectResponse) GetResourcesOk() ([]InfrahubResourceObjectResponse, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *OrganizationObjectResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []InfrahubResourceObjectResponse and assigns it to the Resources field.
func (o *OrganizationObjectResponse) SetResources(v []InfrahubResourceObjectResponse) {
	o.Resources = v
}

func (o OrganizationObjectResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationObjectResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableOrganizationObjectResponse struct {
	value *OrganizationObjectResponse
	isSet bool
}

func (v NullableOrganizationObjectResponse) Get() *OrganizationObjectResponse {
	return v.value
}

func (v *NullableOrganizationObjectResponse) Set(val *OrganizationObjectResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationObjectResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationObjectResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationObjectResponse(val *OrganizationObjectResponse) *NullableOrganizationObjectResponse {
	return &NullableOrganizationObjectResponse{value: val, isSet: true}
}

func (v NullableOrganizationObjectResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationObjectResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
