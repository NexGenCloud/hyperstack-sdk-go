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

// checks if the GetPoliciesResponseModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetPoliciesResponseModel{}

// GetPoliciesResponseModel struct for GetPoliciesResponseModel
type GetPoliciesResponseModel struct {
	Message  *string        `json:"message,omitempty"`
	Policies []PolicyFields `json:"policies,omitempty"`
	Status   *bool          `json:"status,omitempty"`
}

// NewGetPoliciesResponseModel instantiates a new GetPoliciesResponseModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetPoliciesResponseModel() *GetPoliciesResponseModel {
	this := GetPoliciesResponseModel{}
	return &this
}

// NewGetPoliciesResponseModelWithDefaults instantiates a new GetPoliciesResponseModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetPoliciesResponseModelWithDefaults() *GetPoliciesResponseModel {
	this := GetPoliciesResponseModel{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetPoliciesResponseModel) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPoliciesResponseModel) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetPoliciesResponseModel) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetPoliciesResponseModel) SetMessage(v string) {
	o.Message = &v
}

// GetPolicies returns the Policies field value if set, zero value otherwise.
func (o *GetPoliciesResponseModel) GetPolicies() []PolicyFields {
	if o == nil || IsNil(o.Policies) {
		var ret []PolicyFields
		return ret
	}
	return o.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPoliciesResponseModel) GetPoliciesOk() ([]PolicyFields, bool) {
	if o == nil || IsNil(o.Policies) {
		return nil, false
	}
	return o.Policies, true
}

// HasPolicies returns a boolean if a field has been set.
func (o *GetPoliciesResponseModel) HasPolicies() bool {
	if o != nil && !IsNil(o.Policies) {
		return true
	}

	return false
}

// SetPolicies gets a reference to the given []PolicyFields and assigns it to the Policies field.
func (o *GetPoliciesResponseModel) SetPolicies(v []PolicyFields) {
	o.Policies = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetPoliciesResponseModel) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetPoliciesResponseModel) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetPoliciesResponseModel) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *GetPoliciesResponseModel) SetStatus(v bool) {
	o.Status = &v
}

func (o GetPoliciesResponseModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetPoliciesResponseModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Policies) {
		toSerialize["policies"] = o.Policies
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableGetPoliciesResponseModel struct {
	value *GetPoliciesResponseModel
	isSet bool
}

func (v NullableGetPoliciesResponseModel) Get() *GetPoliciesResponseModel {
	return v.value
}

func (v *NullableGetPoliciesResponseModel) Set(val *GetPoliciesResponseModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPoliciesResponseModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPoliciesResponseModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPoliciesResponseModel(val *GetPoliciesResponseModel) *NullableGetPoliciesResponseModel {
	return &NullableGetPoliciesResponseModel{value: val, isSet: true}
}

func (v NullableGetPoliciesResponseModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPoliciesResponseModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
