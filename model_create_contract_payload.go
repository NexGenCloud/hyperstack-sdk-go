/*
Infrahub-API

Leverage the Infrahub API and Hyperstack platform to easily create, manage, and scale powerful GPU virtual machines and their associated resources.   Access this SDK to automate the deployment of your workloads and streamline your infrastructure management.  To contribute, please raise an issue with a bug report, feature request, feedback, or general inquiry.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hyperstack

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CreateContractPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateContractPayload{}

// CreateContractPayload struct for CreateContractPayload
type CreateContractPayload struct {
	Description       *string                   `json:"description,omitempty"`
	DiscountResources []ContractResourcePayload `json:"discount_resources"`
	EndDate           *CustomTime               `json:"end_date,omitempty"`
	ExpirationPolicy  int32                     `json:"expiration_policy"`
	OrgId             int32                     `json:"org_id"`
	StartDate         *CustomTime               `json:"start_date,omitempty"`
}

type _CreateContractPayload CreateContractPayload

// NewCreateContractPayload instantiates a new CreateContractPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateContractPayload(discountResources []ContractResourcePayload, expirationPolicy int32, orgId int32) *CreateContractPayload {
	this := CreateContractPayload{}
	this.DiscountResources = discountResources
	this.ExpirationPolicy = expirationPolicy
	this.OrgId = orgId
	return &this
}

// NewCreateContractPayloadWithDefaults instantiates a new CreateContractPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateContractPayloadWithDefaults() *CreateContractPayload {
	this := CreateContractPayload{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateContractPayload) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContractPayload) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateContractPayload) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateContractPayload) SetDescription(v string) {
	o.Description = &v
}

// GetDiscountResources returns the DiscountResources field value
func (o *CreateContractPayload) GetDiscountResources() []ContractResourcePayload {
	if o == nil {
		var ret []ContractResourcePayload
		return ret
	}

	return o.DiscountResources
}

// GetDiscountResourcesOk returns a tuple with the DiscountResources field value
// and a boolean to check if the value has been set.
func (o *CreateContractPayload) GetDiscountResourcesOk() ([]ContractResourcePayload, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiscountResources, true
}

// SetDiscountResources sets field value
func (o *CreateContractPayload) SetDiscountResources(v []ContractResourcePayload) {
	o.DiscountResources = v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *CreateContractPayload) GetEndDate() CustomTime {
	if o == nil || IsNil(o.EndDate) {
		var ret CustomTime
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContractPayload) GetEndDateOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *CreateContractPayload) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given CustomTime and assigns it to the EndDate field.
func (o *CreateContractPayload) SetEndDate(v CustomTime) {
	o.EndDate = &v
}

// GetExpirationPolicy returns the ExpirationPolicy field value
func (o *CreateContractPayload) GetExpirationPolicy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpirationPolicy
}

// GetExpirationPolicyOk returns a tuple with the ExpirationPolicy field value
// and a boolean to check if the value has been set.
func (o *CreateContractPayload) GetExpirationPolicyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationPolicy, true
}

// SetExpirationPolicy sets field value
func (o *CreateContractPayload) SetExpirationPolicy(v int32) {
	o.ExpirationPolicy = v
}

// GetOrgId returns the OrgId field value
func (o *CreateContractPayload) GetOrgId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *CreateContractPayload) GetOrgIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *CreateContractPayload) SetOrgId(v int32) {
	o.OrgId = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *CreateContractPayload) GetStartDate() CustomTime {
	if o == nil || IsNil(o.StartDate) {
		var ret CustomTime
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateContractPayload) GetStartDateOk() (*CustomTime, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *CreateContractPayload) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given CustomTime and assigns it to the StartDate field.
func (o *CreateContractPayload) SetStartDate(v CustomTime) {
	o.StartDate = &v
}

func (o CreateContractPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateContractPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["discount_resources"] = o.DiscountResources
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	toSerialize["expiration_policy"] = o.ExpirationPolicy
	toSerialize["org_id"] = o.OrgId
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	return toSerialize, nil
}

func (o *CreateContractPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"discount_resources",
		"expiration_policy",
		"org_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateContractPayload := _CreateContractPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateContractPayload)

	if err != nil {
		return err
	}

	*o = CreateContractPayload(varCreateContractPayload)

	return err
}

type NullableCreateContractPayload struct {
	value *CreateContractPayload
	isSet bool
}

func (v NullableCreateContractPayload) Get() *CreateContractPayload {
	return v.value
}

func (v *NullableCreateContractPayload) Set(val *CreateContractPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateContractPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateContractPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateContractPayload(val *CreateContractPayload) *NullableCreateContractPayload {
	return &NullableCreateContractPayload{value: val, isSet: true}
}

func (v NullableCreateContractPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateContractPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
