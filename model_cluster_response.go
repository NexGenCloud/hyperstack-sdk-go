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

// checks if the ClusterResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterResponse{}

// ClusterResponse struct for ClusterResponse
type ClusterResponse struct {
	Cluster *ClusterFields `json:"cluster,omitempty"`
	Message *string        `json:"message,omitempty"`
	Status  *bool          `json:"status,omitempty"`
}

// NewClusterResponse instantiates a new ClusterResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterResponse() *ClusterResponse {
	this := ClusterResponse{}
	return &this
}

// NewClusterResponseWithDefaults instantiates a new ClusterResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterResponseWithDefaults() *ClusterResponse {
	this := ClusterResponse{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *ClusterResponse) GetCluster() ClusterFields {
	if o == nil || IsNil(o.Cluster) {
		var ret ClusterFields
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetClusterOk() (*ClusterFields, bool) {
	if o == nil || IsNil(o.Cluster) {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *ClusterResponse) HasCluster() bool {
	if o != nil && !IsNil(o.Cluster) {
		return true
	}

	return false
}

// SetCluster gets a reference to the given ClusterFields and assigns it to the Cluster field.
func (o *ClusterResponse) SetCluster(v ClusterFields) {
	o.Cluster = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ClusterResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ClusterResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ClusterResponse) SetMessage(v string) {
	o.Message = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ClusterResponse) GetStatus() bool {
	if o == nil || IsNil(o.Status) {
		var ret bool
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterResponse) GetStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ClusterResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given bool and assigns it to the Status field.
func (o *ClusterResponse) SetStatus(v bool) {
	o.Status = &v
}

func (o ClusterResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cluster) {
		toSerialize["cluster"] = o.Cluster
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableClusterResponse struct {
	value *ClusterResponse
	isSet bool
}

func (v NullableClusterResponse) Get() *ClusterResponse {
	return v.value
}

func (v *NullableClusterResponse) Set(val *ClusterResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterResponse(val *ClusterResponse) *NullableClusterResponse {
	return &NullableClusterResponse{value: val, isSet: true}
}

func (v NullableClusterResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}