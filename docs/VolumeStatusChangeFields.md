# VolumeStatusChangeFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CurrentStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**PreviousStatus** | Pointer to **string** |  | [optional] 
**VolumeId** | Pointer to **int32** |  | [optional] 

## Methods

### NewVolumeStatusChangeFields

`func NewVolumeStatusChangeFields() *VolumeStatusChangeFields`

NewVolumeStatusChangeFields instantiates a new VolumeStatusChangeFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVolumeStatusChangeFieldsWithDefaults

`func NewVolumeStatusChangeFieldsWithDefaults() *VolumeStatusChangeFields`

NewVolumeStatusChangeFieldsWithDefaults instantiates a new VolumeStatusChangeFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *VolumeStatusChangeFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VolumeStatusChangeFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VolumeStatusChangeFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VolumeStatusChangeFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentStatus

`func (o *VolumeStatusChangeFields) GetCurrentStatus() string`

GetCurrentStatus returns the CurrentStatus field if non-nil, zero value otherwise.

### GetCurrentStatusOk

`func (o *VolumeStatusChangeFields) GetCurrentStatusOk() (*string, bool)`

GetCurrentStatusOk returns a tuple with the CurrentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentStatus

`func (o *VolumeStatusChangeFields) SetCurrentStatus(v string)`

SetCurrentStatus sets CurrentStatus field to given value.

### HasCurrentStatus

`func (o *VolumeStatusChangeFields) HasCurrentStatus() bool`

HasCurrentStatus returns a boolean if a field has been set.

### GetId

`func (o *VolumeStatusChangeFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VolumeStatusChangeFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VolumeStatusChangeFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *VolumeStatusChangeFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPreviousStatus

`func (o *VolumeStatusChangeFields) GetPreviousStatus() string`

GetPreviousStatus returns the PreviousStatus field if non-nil, zero value otherwise.

### GetPreviousStatusOk

`func (o *VolumeStatusChangeFields) GetPreviousStatusOk() (*string, bool)`

GetPreviousStatusOk returns a tuple with the PreviousStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousStatus

`func (o *VolumeStatusChangeFields) SetPreviousStatus(v string)`

SetPreviousStatus sets PreviousStatus field to given value.

### HasPreviousStatus

`func (o *VolumeStatusChangeFields) HasPreviousStatus() bool`

HasPreviousStatus returns a boolean if a field has been set.

### GetVolumeId

`func (o *VolumeStatusChangeFields) GetVolumeId() int32`

GetVolumeId returns the VolumeId field if non-nil, zero value otherwise.

### GetVolumeIdOk

`func (o *VolumeStatusChangeFields) GetVolumeIdOk() (*int32, bool)`

GetVolumeIdOk returns a tuple with the VolumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeId

`func (o *VolumeStatusChangeFields) SetVolumeId(v int32)`

SetVolumeId sets VolumeId field to given value.

### HasVolumeId

`func (o *VolumeStatusChangeFields) HasVolumeId() bool`

HasVolumeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


