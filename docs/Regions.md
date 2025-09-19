# Regions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Regions** | Pointer to [**[]RegionFields**](RegionFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegions

`func NewRegions() *Regions`

NewRegions instantiates a new Regions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionsWithDefaults

`func NewRegionsWithDefaults() *Regions`

NewRegionsWithDefaults instantiates a new Regions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Regions) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Regions) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Regions) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Regions) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRegions

`func (o *Regions) GetRegions() []RegionFields`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *Regions) GetRegionsOk() (*[]RegionFields, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *Regions) SetRegions(v []RegionFields)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *Regions) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetStatus

`func (o *Regions) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Regions) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Regions) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Regions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


