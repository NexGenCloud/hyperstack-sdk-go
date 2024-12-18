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

// checks if the Billingmetricesresponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Billingmetricesresponse{}

// Billingmetricesresponse struct for Billingmetricesresponse
type Billingmetricesresponse struct {
	Data    []Billingmetricesfields `json:"data,omitempty"`
	Message *string                 `json:"message,omitempty"`
	Status  *bool                   `json:"status,omitempty"`
}

// NewBillingmetricesresponse instantiates a new Billingmetricesresponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingmetricesresponse() *Billingmetricesresponse {
	this := Billingmetricesresponse{}
	return &this
}

// NewBillingmetricesresponseWithDefaults instantiates a new Billingmetricesresponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingmetricesresponseWithDefaults() *Billingmetricesresponse {
	this := Billingmetricesresponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *Billingmetricesresponse) GetData() []Billingmetricesfields {
	if o == nil || IsNil(o.Data) {
		var ret []Billingmetricesfields
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billingmetricesresponse) GetDataOk() ([]Billingmetricesfields, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *Billingmetricesresponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Billingmetricesfields and assigns it to the Data field.
func (o *Billingmetricesresponse) SetData(v []Billingmetricesfields) {
	o.Data = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *Billingmetricesresponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billingmetricesresponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *Billingmetricesresponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *Billingmetricesresponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Billingmetricesresponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billingmetricesresponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Billingmetricesresponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *Billingmetricesresponse) SetStatus(v bool) {
	o.Status = &v
}

func (o Billingmetricesresponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Billingmetricesresponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableBillingmetricesresponse struct {
	value *Billingmetricesresponse
	isSet bool
}

func (v NullableBillingmetricesresponse) Get() *Billingmetricesresponse {
	return v.value
}

func (v *NullableBillingmetricesresponse) Set(val *Billingmetricesresponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBillingmetricesresponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBillingmetricesresponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBillingmetricesresponse(val *Billingmetricesresponse) *NullableBillingmetricesresponse {
	return &NullableBillingmetricesresponse{value: val, isSet: true}
}

func (v NullableBillingmetricesresponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBillingmetricesresponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
