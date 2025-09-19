# UpdateKeypairNameResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Keypair** | Pointer to [**KeypairFields**](KeypairFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateKeypairNameResponse

`func NewUpdateKeypairNameResponse() *UpdateKeypairNameResponse`

NewUpdateKeypairNameResponse instantiates a new UpdateKeypairNameResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateKeypairNameResponseWithDefaults

`func NewUpdateKeypairNameResponseWithDefaults() *UpdateKeypairNameResponse`

NewUpdateKeypairNameResponseWithDefaults instantiates a new UpdateKeypairNameResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeypair

`func (o *UpdateKeypairNameResponse) GetKeypair() KeypairFields`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *UpdateKeypairNameResponse) GetKeypairOk() (*KeypairFields, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *UpdateKeypairNameResponse) SetKeypair(v KeypairFields)`

SetKeypair sets Keypair field to given value.

### HasKeypair

`func (o *UpdateKeypairNameResponse) HasKeypair() bool`

HasKeypair returns a boolean if a field has been set.

### GetMessage

`func (o *UpdateKeypairNameResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateKeypairNameResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateKeypairNameResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateKeypairNameResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateKeypairNameResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateKeypairNameResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateKeypairNameResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateKeypairNameResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


