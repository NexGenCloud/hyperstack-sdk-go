# CreateAutoTopupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoTopup** | Pointer to [**AutoTopup**](AutoTopup.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateAutoTopupResponse

`func NewCreateAutoTopupResponse() *CreateAutoTopupResponse`

NewCreateAutoTopupResponse instantiates a new CreateAutoTopupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutoTopupResponseWithDefaults

`func NewCreateAutoTopupResponseWithDefaults() *CreateAutoTopupResponse`

NewCreateAutoTopupResponseWithDefaults instantiates a new CreateAutoTopupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoTopup

`func (o *CreateAutoTopupResponse) GetAutoTopup() AutoTopup`

GetAutoTopup returns the AutoTopup field if non-nil, zero value otherwise.

### GetAutoTopupOk

`func (o *CreateAutoTopupResponse) GetAutoTopupOk() (*AutoTopup, bool)`

GetAutoTopupOk returns a tuple with the AutoTopup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoTopup

`func (o *CreateAutoTopupResponse) SetAutoTopup(v AutoTopup)`

SetAutoTopup sets AutoTopup field to given value.

### HasAutoTopup

`func (o *CreateAutoTopupResponse) HasAutoTopup() bool`

HasAutoTopup returns a boolean if a field has been set.

### GetMessage

`func (o *CreateAutoTopupResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateAutoTopupResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateAutoTopupResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateAutoTopupResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *CreateAutoTopupResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateAutoTopupResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateAutoTopupResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateAutoTopupResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


