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

// checks if the FutureNodeStockModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FutureNodeStockModel{}

// FutureNodeStockModel struct for FutureNodeStockModel
type FutureNodeStockModel struct {
	ExpectedAmount int32   `json:"expected_amount"`
	Id             *int32  `json:"id,omitempty"`
	Name           *string `json:"name,omitempty"`
}

type _FutureNodeStockModel FutureNodeStockModel

// NewFutureNodeStockModel instantiates a new FutureNodeStockModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFutureNodeStockModel(expectedAmount int32) *FutureNodeStockModel {
	this := FutureNodeStockModel{}
	this.ExpectedAmount = expectedAmount
	return &this
}

// NewFutureNodeStockModelWithDefaults instantiates a new FutureNodeStockModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFutureNodeStockModelWithDefaults() *FutureNodeStockModel {
	this := FutureNodeStockModel{}
	return &this
}

// GetExpectedAmount returns the ExpectedAmount field value
func (o *FutureNodeStockModel) GetExpectedAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpectedAmount
}

// GetExpectedAmountOk returns a tuple with the ExpectedAmount field value
// and a boolean to check if the value has been set.
func (o *FutureNodeStockModel) GetExpectedAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpectedAmount, true
}

// SetExpectedAmount sets field value
func (o *FutureNodeStockModel) SetExpectedAmount(v int32) {
	o.ExpectedAmount = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *FutureNodeStockModel) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FutureNodeStockModel) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *FutureNodeStockModel) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *FutureNodeStockModel) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FutureNodeStockModel) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FutureNodeStockModel) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FutureNodeStockModel) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FutureNodeStockModel) SetName(v string) {
	o.Name = &v
}

func (o FutureNodeStockModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FutureNodeStockModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["expected_amount"] = o.ExpectedAmount
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

func (o *FutureNodeStockModel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"expected_amount",
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

	varFutureNodeStockModel := _FutureNodeStockModel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFutureNodeStockModel)

	if err != nil {
		return err
	}

	*o = FutureNodeStockModel(varFutureNodeStockModel)

	return err
}

type NullableFutureNodeStockModel struct {
	value *FutureNodeStockModel
	isSet bool
}

func (v NullableFutureNodeStockModel) Get() *FutureNodeStockModel {
	return v.value
}

func (v *NullableFutureNodeStockModel) Set(val *FutureNodeStockModel) {
	v.value = val
	v.isSet = true
}

func (v NullableFutureNodeStockModel) IsSet() bool {
	return v.isSet
}

func (v *NullableFutureNodeStockModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFutureNodeStockModel(val *FutureNodeStockModel) *NullableFutureNodeStockModel {
	return &NullableFutureNodeStockModel{value: val, isSet: true}
}

func (v NullableFutureNodeStockModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFutureNodeStockModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
