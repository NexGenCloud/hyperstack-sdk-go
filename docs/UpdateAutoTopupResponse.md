# UpdateAutoTopupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTopup** | Pointer to [**AutoTopup**](AutoTopup.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateAutoTopupResponse

`func NewUpdateAutoTopupResponse() *UpdateAutoTopupResponse`

NewUpdateAutoTopupResponse instantiates a new UpdateAutoTopupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAutoTopupResponseWithDefaults

`func NewUpdateAutoTopupResponseWithDefaults() *UpdateAutoTopupResponse`

NewUpdateAutoTopupResponseWithDefaults instantiates a new UpdateAutoTopupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTopup

`func (o *UpdateAutoTopupResponse) GetAutoTopup() AutoTopup`

GetAutoTopup returns the AutoTopup field if non-nil, zero value otherwise.

### GetAutoTopupOk

`func (o *UpdateAutoTopupResponse) GetAutoTopupOk() (*AutoTopup, bool)`

GetAutoTopupOk returns a tuple with the AutoTopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopup

`func (o *UpdateAutoTopupResponse) SetAutoTopup(v AutoTopup)`

SetAutoTopup sets AutoTopup field to given value.

### HasAutoTopup

`func (o *UpdateAutoTopupResponse) HasAutoTopup() bool`

HasAutoTopup returns a boolean if a field has been set.

### GetMessage

`func (o *UpdateAutoTopupResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateAutoTopupResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateAutoTopupResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateAutoTopupResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateAutoTopupResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAutoTopupResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAutoTopupResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAutoTopupResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


