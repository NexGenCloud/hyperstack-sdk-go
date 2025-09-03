# LastDayCostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**LastDayCostFields**](LastDayCostFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewLastDayCostResponse

`func NewLastDayCostResponse() *LastDayCostResponse`

NewLastDayCostResponse instantiates a new LastDayCostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLastDayCostResponseWithDefaults

`func NewLastDayCostResponseWithDefaults() *LastDayCostResponse`

NewLastDayCostResponseWithDefaults instantiates a new LastDayCostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *LastDayCostResponse) GetData() LastDayCostFields`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LastDayCostResponse) GetDataOk() (*LastDayCostFields, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LastDayCostResponse) SetData(v LastDayCostFields)`

SetData sets Data field to given value.

### HasData

`func (o *LastDayCostResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *LastDayCostResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LastDayCostResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LastDayCostResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LastDayCostResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *LastDayCostResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LastDayCostResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LastDayCostResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LastDayCostResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


