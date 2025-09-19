# DetachVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**VolumeAttachments** | Pointer to [**[]AttachVolumeFields**](AttachVolumeFields.md) |  | [optional] 

## Methods

### NewDetachVolumes

`func NewDetachVolumes() *DetachVolumes`

NewDetachVolumes instantiates a new DetachVolumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetachVolumesWithDefaults

`func NewDetachVolumesWithDefaults() *DetachVolumes`

NewDetachVolumesWithDefaults instantiates a new DetachVolumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *DetachVolumes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *DetachVolumes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *DetachVolumes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *DetachVolumes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *DetachVolumes) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DetachVolumes) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DetachVolumes) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DetachVolumes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolumeAttachments

`func (o *DetachVolumes) GetVolumeAttachments() []AttachVolumeFields`

GetVolumeAttachments returns the VolumeAttachments field if non-nil, zero value otherwise.

### GetVolumeAttachmentsOk

`func (o *DetachVolumes) GetVolumeAttachmentsOk() (*[]AttachVolumeFields, bool)`

GetVolumeAttachmentsOk returns a tuple with the VolumeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttachments

`func (o *DetachVolumes) SetVolumeAttachments(v []AttachVolumeFields)`

SetVolumeAttachments sets VolumeAttachments field to given value.

### HasVolumeAttachments

`func (o *DetachVolumes) HasVolumeAttachments() bool`

HasVolumeAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


