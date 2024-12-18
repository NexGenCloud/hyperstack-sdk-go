# InstanceEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceEvents** | Pointer to [**[]InstanceEventsFields**](InstanceEventsFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewInstanceEvents

`func NewInstanceEvents() *InstanceEvents`

NewInstanceEvents instantiates a new InstanceEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceEventsWithDefaults

`func NewInstanceEventsWithDefaults() *InstanceEvents`

NewInstanceEventsWithDefaults instantiates a new InstanceEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceEvents

`func (o *InstanceEvents) GetInstanceEvents() []InstanceEventsFields`

GetInstanceEvents returns the InstanceEvents field if non-nil, zero value otherwise.

### GetInstanceEventsOk

`func (o *InstanceEvents) GetInstanceEventsOk() (*[]InstanceEventsFields, bool)`

GetInstanceEventsOk returns a tuple with the InstanceEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceEvents

`func (o *InstanceEvents) SetInstanceEvents(v []InstanceEventsFields)`

SetInstanceEvents sets InstanceEvents field to given value.

### HasInstanceEvents

`func (o *InstanceEvents) HasInstanceEvents() bool`

HasInstanceEvents returns a boolean if a field has been set.

### GetMessage

`func (o *InstanceEvents) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InstanceEvents) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InstanceEvents) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InstanceEvents) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *InstanceEvents) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceEvents) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceEvents) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceEvents) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


