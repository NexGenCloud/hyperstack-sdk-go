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

// checks if the ClusterEvents type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterEvents{}

// ClusterEvents struct for ClusterEvents
type ClusterEvents struct {
	ClusterEvents []ClusterEventsFields `json:"cluster_events,omitempty"`
	Message       *string               `json:"message,omitempty"`
	Status        *bool                 `json:"status,omitempty"`
}

// NewClusterEvents instantiates a new ClusterEvents object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterEvents() *ClusterEvents {
	this := ClusterEvents{}
	return &this
}

// NewClusterEventsWithDefaults instantiates a new ClusterEvents object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterEventsWithDefaults() *ClusterEvents {
	this := ClusterEvents{}
	return &this
}

// GetClusterEvents returns the ClusterEvents field value if set, zero value otherwise.
func (o *ClusterEvents) GetClusterEvents() []ClusterEventsFields {
	if o == nil || IsNil(o.ClusterEvents) {
		var ret []ClusterEventsFields
		return ret
	}
	return o.ClusterEvents
}

// GetClusterEventsOk returns a tuple with the ClusterEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEvents) GetClusterEventsOk() ([]ClusterEventsFields, bool) {
	if o == nil || IsNil(o.ClusterEvents) {
		return nil, false
	}
	return o.ClusterEvents, true
}

// HasClusterEvents returns a boolean if a field has been set.
func (o *ClusterEvents) HasClusterEvents() bool {
	if o != nil && !IsNil(o.ClusterEvents) {
		return true
	}

	return false
}

// SetClusterEvents gets a reference to the given []ClusterEventsFields and assigns it to the ClusterEvents field.
func (o *ClusterEvents) SetClusterEvents(v []ClusterEventsFields) {
	o.ClusterEvents = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClusterEvents) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEvents) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterEvents) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClusterEvents) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterEvents) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterEvents) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterEvents) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *ClusterEvents) SetStatus(v bool) {
	o.Status = &v
}

func (o ClusterEvents) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterEvents) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterEvents) {
		toSerialize["cluster_events"] = o.ClusterEvents
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableClusterEvents struct {
	value *ClusterEvents
	isSet bool
}

func (v NullableClusterEvents) Get() *ClusterEvents {
	return v.value
}

func (v *NullableClusterEvents) Set(val *ClusterEvents) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterEvents) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterEvents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterEvents(val *ClusterEvents) *NullableClusterEvents {
	return &NullableClusterEvents{value: val, isSet: true}
}

func (v NullableClusterEvents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterEvents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}