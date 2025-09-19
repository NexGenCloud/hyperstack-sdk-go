# CreateSnapshotResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Snapshot** | Pointer to [**SnapshotFields**](SnapshotFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateSnapshotResponse

`func NewCreateSnapshotResponse() *CreateSnapshotResponse`

NewCreateSnapshotResponse instantiates a new CreateSnapshotResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSnapshotResponseWithDefaults

`func NewCreateSnapshotResponseWithDefaults() *CreateSnapshotResponse`

NewCreateSnapshotResponseWithDefaults instantiates a new CreateSnapshotResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateSnapshotResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateSnapshotResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateSnapshotResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateSnapshotResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSnapshot

`func (o *CreateSnapshotResponse) GetSnapshot() SnapshotFields`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *CreateSnapshotResponse) GetSnapshotOk() (*SnapshotFields, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *CreateSnapshotResponse) SetSnapshot(v SnapshotFields)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *CreateSnapshotResponse) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetStatus

`func (o *CreateSnapshotResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateSnapshotResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateSnapshotResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateSnapshotResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


