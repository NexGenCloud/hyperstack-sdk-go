# AttachVolumes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**VolumeAttachments** | Pointer to [**[]AttachVolumeFields**](AttachVolumeFields.md) |  | [optional] 

## Methods

### NewAttachVolumes

`func NewAttachVolumes() *AttachVolumes`

NewAttachVolumes instantiates a new AttachVolumes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachVolumesWithDefaults

`func NewAttachVolumesWithDefaults() *AttachVolumes`

NewAttachVolumesWithDefaults instantiates a new AttachVolumes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *AttachVolumes) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AttachVolumes) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AttachVolumes) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AttachVolumes) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *AttachVolumes) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AttachVolumes) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AttachVolumes) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AttachVolumes) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVolumeAttachments

`func (o *AttachVolumes) GetVolumeAttachments() []AttachVolumeFields`

GetVolumeAttachments returns the VolumeAttachments field if non-nil, zero value otherwise.

### GetVolumeAttachmentsOk

`func (o *AttachVolumes) GetVolumeAttachmentsOk() (*[]AttachVolumeFields, bool)`

GetVolumeAttachmentsOk returns a tuple with the VolumeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttachments

`func (o *AttachVolumes) SetVolumeAttachments(v []AttachVolumeFields)`

SetVolumeAttachments sets VolumeAttachments field to given value.

### HasVolumeAttachments

`func (o *AttachVolumes) HasVolumeAttachments() bool`

HasVolumeAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


