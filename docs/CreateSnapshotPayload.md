# CreateSnapshotPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | description | 
**Labels** | Pointer to **[]string** | Labels associated with snapshot | [optional] 
**Name** | **string** | Snapshot name | 

## Methods

### NewCreateSnapshotPayload

`func NewCreateSnapshotPayload(description string, name string, ) *CreateSnapshotPayload`

NewCreateSnapshotPayload instantiates a new CreateSnapshotPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotPayloadWithDefaults

`func NewCreateSnapshotPayloadWithDefaults() *CreateSnapshotPayload`

NewCreateSnapshotPayloadWithDefaults instantiates a new CreateSnapshotPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateSnapshotPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSnapshotPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSnapshotPayload) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetLabels

`func (o *CreateSnapshotPayload) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateSnapshotPayload) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateSnapshotPayload) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateSnapshotPayload) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *CreateSnapshotPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSnapshotPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSnapshotPayload) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


