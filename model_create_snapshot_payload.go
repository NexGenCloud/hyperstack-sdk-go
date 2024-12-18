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

// checks if the CreateSnapshotPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSnapshotPayload{}

// CreateSnapshotPayload struct for CreateSnapshotPayload
type CreateSnapshotPayload struct {
	// description
	Description string `json:"description"`
	// Indicates if the snapshot is an image
	IsImage bool `json:"is_image"`
	// Labels associated with snapshot
	Labels []string `json:"labels,omitempty"`
	// Snapshot name
	Name string `json:"name"`
}

type _CreateSnapshotPayload CreateSnapshotPayload

// NewCreateSnapshotPayload instantiates a new CreateSnapshotPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSnapshotPayload(description string, isImage bool, name string) *CreateSnapshotPayload {
	this := CreateSnapshotPayload{}
	this.Description = description
	this.IsImage = isImage
	this.Name = name
	return &this
}

// NewCreateSnapshotPayloadWithDefaults instantiates a new CreateSnapshotPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSnapshotPayloadWithDefaults() *CreateSnapshotPayload {
	this := CreateSnapshotPayload{}
	return &this
}

// GetDescription returns the Description field value
func (o *CreateSnapshotPayload) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CreateSnapshotPayload) SetDescription(v string) {
	o.Description = v
}

// GetIsImage returns the IsImage field value
func (o *CreateSnapshotPayload) GetIsImage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsImage
}

// GetIsImageOk returns a tuple with the IsImage field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetIsImageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsImage, true
}

// SetIsImage sets field value
func (o *CreateSnapshotPayload) SetIsImage(v bool) {
	o.IsImage = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CreateSnapshotPayload) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CreateSnapshotPayload) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *CreateSnapshotPayload) SetLabels(v []string) {
	o.Labels = v
}

// GetName returns the Name field value
func (o *CreateSnapshotPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSnapshotPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSnapshotPayload) SetName(v string) {
	o.Name = v
}

func (o CreateSnapshotPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSnapshotPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["is_image"] = o.IsImage
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

func (o *CreateSnapshotPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"is_image",
		"name",
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

	varCreateSnapshotPayload := _CreateSnapshotPayload{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSnapshotPayload)

	if err != nil {
		return err
	}

	*o = CreateSnapshotPayload(varCreateSnapshotPayload)

	return err
}

type NullableCreateSnapshotPayload struct {
	value *CreateSnapshotPayload
	isSet bool
}

func (v NullableCreateSnapshotPayload) Get() *CreateSnapshotPayload {
	return v.value
}

func (v *NullableCreateSnapshotPayload) Set(val *CreateSnapshotPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSnapshotPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSnapshotPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSnapshotPayload(val *CreateSnapshotPayload) *NullableCreateSnapshotPayload {
	return &NullableCreateSnapshotPayload{value: val, isSet: true}
}

func (v NullableCreateSnapshotPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSnapshotPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
