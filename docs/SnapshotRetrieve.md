# SnapshotRetrieve

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Snapshot** | Pointer to [**SnapshotRetrieveFields**](SnapshotRetrieveFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewSnapshotRetrieve

`func NewSnapshotRetrieve() *SnapshotRetrieve`

NewSnapshotRetrieve instantiates a new SnapshotRetrieve object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotRetrieveWithDefaults

`func NewSnapshotRetrieveWithDefaults() *SnapshotRetrieve`

NewSnapshotRetrieveWithDefaults instantiates a new SnapshotRetrieve object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SnapshotRetrieve) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SnapshotRetrieve) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SnapshotRetrieve) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SnapshotRetrieve) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSnapshot

`func (o *SnapshotRetrieve) GetSnapshot() SnapshotRetrieveFields`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *SnapshotRetrieve) GetSnapshotOk() (*SnapshotRetrieveFields, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *SnapshotRetrieve) SetSnapshot(v SnapshotRetrieveFields)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *SnapshotRetrieve) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.

### GetStatus

`func (o *SnapshotRetrieve) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotRetrieve) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotRetrieve) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnapshotRetrieve) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


