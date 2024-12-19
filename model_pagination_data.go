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

// checks if the PaginationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaginationData{}

// PaginationData struct for PaginationData
type PaginationData struct {
	Page    *int32 `json:"page,omitempty"`
	Pages   *int32 `json:"pages,omitempty"`
	PerPage *int32 `json:"per_page,omitempty"`
}

// NewPaginationData instantiates a new PaginationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaginationData() *PaginationData {
	this := PaginationData{}
	return &this
}

// NewPaginationDataWithDefaults instantiates a new PaginationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaginationDataWithDefaults() *PaginationData {
	this := PaginationData{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *PaginationData) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *PaginationData) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *PaginationData) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *PaginationData) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *PaginationData) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *PaginationData) SetPages(v int32) {
	o.Pages = &v
}

// GetPerPage returns the PerPage field value if set, zero value otherwise.
func (o *PaginationData) GetPerPage() int32 {
	if o == nil || IsNil(o.PerPage) {
		var ret int32
		return ret
	}
	return *o.PerPage
}

// GetPerPageOk returns a tuple with the PerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaginationData) GetPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.PerPage) {
		return nil, false
	}
	return o.PerPage, true
}

// HasPerPage returns a boolean if a field has been set.
func (o *PaginationData) HasPerPage() bool {
	if o != nil && !IsNil(o.PerPage) {
		return true
	}

	return false
}

// SetPerPage gets a reference to the given int32 and assigns it to the PerPage field.
func (o *PaginationData) SetPerPage(v int32) {
	o.PerPage = &v
}

func (o PaginationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaginationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	if !IsNil(o.PerPage) {
		toSerialize["per_page"] = o.PerPage
	}
	return toSerialize, nil
}

type NullablePaginationData struct {
	value *PaginationData
	isSet bool
}

func (v NullablePaginationData) Get() *PaginationData {
	return v.value
}

func (v *NullablePaginationData) Set(val *PaginationData) {
	v.value = val
	v.isSet = true
}

func (v NullablePaginationData) IsSet() bool {
	return v.isSet
}

func (v *NullablePaginationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaginationData(val *PaginationData) *NullablePaginationData {
	return &NullablePaginationData{value: val, isSet: true}
}

func (v NullablePaginationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaginationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}