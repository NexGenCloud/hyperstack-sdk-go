# UpdateVolumeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Updates** | Pointer to **map[string]interface{}** | Summary of fields that were updated | [optional] 
**Volume** | Pointer to [**VolumeFields**](VolumeFields.md) |  | [optional] 

## Methods

### NewUpdateVolumeResponse

`func NewUpdateVolumeResponse() *UpdateVolumeResponse`

NewUpdateVolumeResponse instantiates a new UpdateVolumeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVolumeResponseWithDefaults

`func NewUpdateVolumeResponseWithDefaults() *UpdateVolumeResponse`

NewUpdateVolumeResponseWithDefaults instantiates a new UpdateVolumeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateVolumeResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateVolumeResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateVolumeResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateVolumeResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateVolumeResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateVolumeResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateVolumeResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateVolumeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdates

`func (o *UpdateVolumeResponse) GetUpdates() map[string]interface{}`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *UpdateVolumeResponse) GetUpdatesOk() (*map[string]interface{}, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *UpdateVolumeResponse) SetUpdates(v map[string]interface{})`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *UpdateVolumeResponse) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetVolume

`func (o *UpdateVolumeResponse) GetVolume() VolumeFields`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *UpdateVolumeResponse) GetVolumeOk() (*VolumeFields, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *UpdateVolumeResponse) SetVolume(v VolumeFields)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *UpdateVolumeResponse) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


