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

// checks if the CreateDiscountResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDiscountResponse{}

// CreateDiscountResponse struct for CreateDiscountResponse
type CreateDiscountResponse struct {
	DiscountPlan *InsertDiscountPlanFields `json:"discount_plan,omitempty"`
	Message      *string                   `json:"message,omitempty"`
	Status       *bool                     `json:"status,omitempty"`
}

// NewCreateDiscountResponse instantiates a new CreateDiscountResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDiscountResponse() *CreateDiscountResponse {
	this := CreateDiscountResponse{}
	return &this
}

// NewCreateDiscountResponseWithDefaults instantiates a new CreateDiscountResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDiscountResponseWithDefaults() *CreateDiscountResponse {
	this := CreateDiscountResponse{}
	return &this
}

// GetDiscountPlan returns the DiscountPlan field value if set, zero value otherwise.
func (o *CreateDiscountResponse) GetDiscountPlan() InsertDiscountPlanFields {
	if o == nil || IsNil(o.DiscountPlan) {
		var ret InsertDiscountPlanFields
		return ret
	}
	return *o.DiscountPlan
}

// GetDiscountPlanOk returns a tuple with the DiscountPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountResponse) GetDiscountPlanOk() (*InsertDiscountPlanFields, bool) {
	if o == nil || IsNil(o.DiscountPlan) {
		return nil, false
	}
	return o.DiscountPlan, true
}

// HasDiscountPlan returns a boolean if a field has been set.
func (o *CreateDiscountResponse) HasDiscountPlan() bool {
	if o != nil && !IsNil(o.DiscountPlan) {
		return true
	}

	return false
}

// SetDiscountPlan gets a reference to the given InsertDiscountPlanFields and assigns it to the DiscountPlan field.
func (o *CreateDiscountResponse) SetDiscountPlan(v InsertDiscountPlanFields) {
	o.DiscountPlan = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *CreateDiscountResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *CreateDiscountResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *CreateDiscountResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CreateDiscountResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDiscountResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CreateDiscountResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *CreateDiscountResponse) SetStatus(v bool) {
	o.Status = &v
}

func (o CreateDiscountResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDiscountResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DiscountPlan) {
		toSerialize["discount_plan"] = o.DiscountPlan
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCreateDiscountResponse struct {
	value *CreateDiscountResponse
	isSet bool
}

func (v NullableCreateDiscountResponse) Get() *CreateDiscountResponse {
	return v.value
}

func (v *NullableCreateDiscountResponse) Set(val *CreateDiscountResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDiscountResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDiscountResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDiscountResponse(val *CreateDiscountResponse) *NullableCreateDiscountResponse {
	return &NullableCreateDiscountResponse{value: val, isSet: true}
}

func (v NullableCreateDiscountResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDiscountResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
