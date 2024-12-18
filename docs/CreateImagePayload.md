# CreateImagePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | Pointer to **[]string** | List of labels to attach to the image | [optional] 
**Name** | **string** | Name for the new custom image | 

## Methods

### NewCreateImagePayload

`func NewCreateImagePayload(name string, ) *CreateImagePayload`

NewCreateImagePayload instantiates a new CreateImagePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateImagePayloadWithDefaults

`func NewCreateImagePayloadWithDefaults() *CreateImagePayload`

NewCreateImagePayloadWithDefaults instantiates a new CreateImagePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *CreateImagePayload) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateImagePayload) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateImagePayload) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateImagePayload) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *CreateImagePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateImagePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateImagePayload) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


