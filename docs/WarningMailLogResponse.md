# WarningMailLogResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WarningMailLogFields**](WarningMailLogFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewWarningMailLogResponse

`func NewWarningMailLogResponse() *WarningMailLogResponse`

NewWarningMailLogResponse instantiates a new WarningMailLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningMailLogResponseWithDefaults

`func NewWarningMailLogResponseWithDefaults() *WarningMailLogResponse`

NewWarningMailLogResponseWithDefaults instantiates a new WarningMailLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *WarningMailLogResponse) GetData() []WarningMailLogFields`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *WarningMailLogResponse) GetDataOk() (*[]WarningMailLogFields, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *WarningMailLogResponse) SetData(v []WarningMailLogFields)`

SetData sets Data field to given value.

### HasData

`func (o *WarningMailLogResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *WarningMailLogResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WarningMailLogResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WarningMailLogResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WarningMailLogResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *WarningMailLogResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WarningMailLogResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WarningMailLogResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WarningMailLogResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *WarningMailLogResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WarningMailLogResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WarningMailLogResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *WarningMailLogResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


