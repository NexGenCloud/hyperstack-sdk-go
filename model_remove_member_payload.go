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

// checks if the RemoveMemberPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RemoveMemberPayload{}

// RemoveMemberPayload struct for RemoveMemberPayload
type RemoveMemberPayload struct {
	// The email of the user to be removed.
	Email string `json:"email"`
}

type _RemoveMemberPayload RemoveMemberPayload

// NewRemoveMemberPayload instantiates a new RemoveMemberPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRemoveMemberPayload(email string) *RemoveMemberPayload {
	this := RemoveMemberPayload{}
	this.Email = email
	return &this
}

// NewRemoveMemberPayloadWithDefaults instantiates a new RemoveMemberPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRemoveMemberPayloadWithDefaults() *RemoveMemberPayload {
	this := RemoveMemberPayload{}
	return &this
}

// GetEmail returns the Email field value
func (o *RemoveMemberPayload) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *RemoveMemberPayload) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *RemoveMemberPayload) SetEmail(v string) {
	o.Email = v
}

func (o RemoveMemberPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RemoveMemberPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email"] = o.Email
	return toSerialize, nil
}

func (o *RemoveMemberPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email",
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

	varRemoveMemberPayload := _RemoveMemberPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRemoveMemberPayload)

	if err != nil {
		return err
	}

	*o = RemoveMemberPayload(varRemoveMemberPayload)

	return err
}

type NullableRemoveMemberPayload struct {
	value *RemoveMemberPayload
	isSet bool
}

func (v NullableRemoveMemberPayload) Get() *RemoveMemberPayload {
	return v.value
}

func (v *NullableRemoveMemberPayload) Set(val *RemoveMemberPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableRemoveMemberPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableRemoveMemberPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRemoveMemberPayload(val *RemoveMemberPayload) *NullableRemoveMemberPayload {
	return &NullableRemoveMemberPayload{value: val, isSet: true}
}

func (v NullableRemoveMemberPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRemoveMemberPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
