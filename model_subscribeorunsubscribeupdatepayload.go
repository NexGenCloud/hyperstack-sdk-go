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

// checks if the Subscribeorunsubscribeupdatepayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subscribeorunsubscribeupdatepayload{}

// Subscribeorunsubscribeupdatepayload struct for Subscribeorunsubscribeupdatepayload
type Subscribeorunsubscribeupdatepayload struct {
	// `false` indicates that the user will no longer receive notifications for this specific threshold, whereas `true` signifies that the user will receive notification emails.
	Subscribe bool `json:"subscribe"`
}

type _Subscribeorunsubscribeupdatepayload Subscribeorunsubscribeupdatepayload

// NewSubscribeorunsubscribeupdatepayload instantiates a new Subscribeorunsubscribeupdatepayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscribeorunsubscribeupdatepayload(subscribe bool) *Subscribeorunsubscribeupdatepayload {
	this := Subscribeorunsubscribeupdatepayload{}
	this.Subscribe = subscribe
	return &this
}

// NewSubscribeorunsubscribeupdatepayloadWithDefaults instantiates a new Subscribeorunsubscribeupdatepayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscribeorunsubscribeupdatepayloadWithDefaults() *Subscribeorunsubscribeupdatepayload {
	this := Subscribeorunsubscribeupdatepayload{}
	return &this
}

// GetSubscribe returns the Subscribe field value
func (o *Subscribeorunsubscribeupdatepayload) GetSubscribe() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Subscribe
}

// GetSubscribeOk returns a tuple with the Subscribe field value
// and a boolean to check if the value has been set.
func (o *Subscribeorunsubscribeupdatepayload) GetSubscribeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subscribe, true
}

// SetSubscribe sets field value
func (o *Subscribeorunsubscribeupdatepayload) SetSubscribe(v bool) {
	o.Subscribe = v
}

func (o Subscribeorunsubscribeupdatepayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subscribeorunsubscribeupdatepayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscribe"] = o.Subscribe
	return toSerialize, nil
}

func (o *Subscribeorunsubscribeupdatepayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscribe",
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

	varSubscribeorunsubscribeupdatepayload := _Subscribeorunsubscribeupdatepayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscribeorunsubscribeupdatepayload)

	if err != nil {
		return err
	}

	*o = Subscribeorunsubscribeupdatepayload(varSubscribeorunsubscribeupdatepayload)

	return err
}

type NullableSubscribeorunsubscribeupdatepayload struct {
	value *Subscribeorunsubscribeupdatepayload
	isSet bool
}

func (v NullableSubscribeorunsubscribeupdatepayload) Get() *Subscribeorunsubscribeupdatepayload {
	return v.value
}

func (v *NullableSubscribeorunsubscribeupdatepayload) Set(val *Subscribeorunsubscribeupdatepayload) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribeorunsubscribeupdatepayload) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribeorunsubscribeupdatepayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribeorunsubscribeupdatepayload(val *Subscribeorunsubscribeupdatepayload) *NullableSubscribeorunsubscribeupdatepayload {
	return &NullableSubscribeorunsubscribeupdatepayload{value: val, isSet: true}
}

func (v NullableSubscribeorunsubscribeupdatepayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribeorunsubscribeupdatepayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}