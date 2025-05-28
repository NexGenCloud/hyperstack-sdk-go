# RegionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Region** | Pointer to [**RegionFields**](RegionFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewRegionResponse

`func NewRegionResponse() *RegionResponse`

NewRegionResponse instantiates a new RegionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegionResponseWithDefaults

`func NewRegionResponseWithDefaults() *RegionResponse`

NewRegionResponseWithDefaults instantiates a new RegionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RegionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RegionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RegionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RegionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRegion

`func (o *RegionResponse) GetRegion() RegionFields`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *RegionResponse) GetRegionOk() (*RegionFields, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *RegionResponse) SetRegion(v RegionFields)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *RegionResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *RegionResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RegionResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RegionResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RegionResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


