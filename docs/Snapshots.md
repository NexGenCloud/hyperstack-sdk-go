# Snapshots

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**PageSize** | Pointer to **int32** |  | [optional] 
**Snapshots** | Pointer to [**[]SnapshotFields**](SnapshotFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewSnapshots

`func NewSnapshots() *Snapshots`

NewSnapshots instantiates a new Snapshots object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotsWithDefaults

`func NewSnapshotsWithDefaults() *Snapshots`

NewSnapshotsWithDefaults instantiates a new Snapshots object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *Snapshots) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Snapshots) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Snapshots) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *Snapshots) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetMessage

`func (o *Snapshots) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Snapshots) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Snapshots) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Snapshots) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPage

`func (o *Snapshots) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *Snapshots) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *Snapshots) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *Snapshots) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *Snapshots) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *Snapshots) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *Snapshots) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *Snapshots) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetSnapshots

`func (o *Snapshots) GetSnapshots() []SnapshotFields`

GetSnapshots returns the Snapshots field if non-nil, zero value otherwise.

### GetSnapshotsOk

`func (o *Snapshots) GetSnapshotsOk() (*[]SnapshotFields, bool)`

GetSnapshotsOk returns a tuple with the Snapshots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshots

`func (o *Snapshots) SetSnapshots(v []SnapshotFields)`

SetSnapshots sets Snapshots field to given value.

### HasSnapshots

`func (o *Snapshots) HasSnapshots() bool`

HasSnapshots returns a boolean if a field has been set.

### GetStatus

`func (o *Snapshots) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Snapshots) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Snapshots) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Snapshots) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


