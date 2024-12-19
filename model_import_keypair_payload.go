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

// checks if the ImportKeypairPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportKeypairPayload{}

// ImportKeypairPayload struct for ImportKeypairPayload
type ImportKeypairPayload struct {
	// The name of the environment where the key pair is being created.
	EnvironmentName string `json:"environment_name"`
	// The name of the key pair that is being created.
	Name string `json:"name"`
	// The public key that is being used to import an SSH key pair.
	PublicKey string `json:"public_key"`
}

type _ImportKeypairPayload ImportKeypairPayload

// NewImportKeypairPayload instantiates a new ImportKeypairPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportKeypairPayload(environmentName string, name string, publicKey string) *ImportKeypairPayload {
	this := ImportKeypairPayload{}
	this.EnvironmentName = environmentName
	this.Name = name
	this.PublicKey = publicKey
	return &this
}

// NewImportKeypairPayloadWithDefaults instantiates a new ImportKeypairPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportKeypairPayloadWithDefaults() *ImportKeypairPayload {
	this := ImportKeypairPayload{}
	return &this
}

// GetEnvironmentName returns the EnvironmentName field value
func (o *ImportKeypairPayload) GetEnvironmentName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EnvironmentName
}

// GetEnvironmentNameOk returns a tuple with the EnvironmentName field value
// and a boolean to check if the value has been set.
func (o *ImportKeypairPayload) GetEnvironmentNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnvironmentName, true
}

// SetEnvironmentName sets field value
func (o *ImportKeypairPayload) SetEnvironmentName(v string) {
	o.EnvironmentName = v
}

// GetName returns the Name field value
func (o *ImportKeypairPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportKeypairPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImportKeypairPayload) SetName(v string) {
	o.Name = v
}

// GetPublicKey returns the PublicKey field value
func (o *ImportKeypairPayload) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *ImportKeypairPayload) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *ImportKeypairPayload) SetPublicKey(v string) {
	o.PublicKey = v
}

func (o ImportKeypairPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportKeypairPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment_name"] = o.EnvironmentName
	toSerialize["name"] = o.Name
	toSerialize["public_key"] = o.PublicKey
	return toSerialize, nil
}

func (o *ImportKeypairPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment_name",
		"name",
		"public_key",
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

	varImportKeypairPayload := _ImportKeypairPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varImportKeypairPayload)

	if err != nil {
		return err
	}

	*o = ImportKeypairPayload(varImportKeypairPayload)

	return err
}

type NullableImportKeypairPayload struct {
	value *ImportKeypairPayload
	isSet bool
}

func (v NullableImportKeypairPayload) Get() *ImportKeypairPayload {
	return v.value
}

func (v *NullableImportKeypairPayload) Set(val *ImportKeypairPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableImportKeypairPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableImportKeypairPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportKeypairPayload(val *ImportKeypairPayload) *NullableImportKeypairPayload {
	return &NullableImportKeypairPayload{value: val, isSet: true}
}

func (v NullableImportKeypairPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportKeypairPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}