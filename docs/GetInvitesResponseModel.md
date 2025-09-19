# GetInvitesResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invites** | Pointer to [**[]InviteFields**](InviteFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetInvitesResponseModel

`func NewGetInvitesResponseModel() *GetInvitesResponseModel`

NewGetInvitesResponseModel instantiates a new GetInvitesResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvitesResponseModelWithDefaults

`func NewGetInvitesResponseModelWithDefaults() *GetInvitesResponseModel`

NewGetInvitesResponseModelWithDefaults instantiates a new GetInvitesResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvites

`func (o *GetInvitesResponseModel) GetInvites() []InviteFields`

GetInvites returns the Invites field if non-nil, zero value otherwise.

### GetInvitesOk

`func (o *GetInvitesResponseModel) GetInvitesOk() (*[]InviteFields, bool)`

GetInvitesOk returns a tuple with the Invites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvites

`func (o *GetInvitesResponseModel) SetInvites(v []InviteFields)`

SetInvites sets Invites field to given value.

### HasInvites

`func (o *GetInvitesResponseModel) HasInvites() bool`

HasInvites returns a boolean if a field has been set.

### GetMessage

`func (o *GetInvitesResponseModel) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetInvitesResponseModel) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetInvitesResponseModel) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetInvitesResponseModel) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *GetInvitesResponseModel) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInvitesResponseModel) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInvitesResponseModel) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInvitesResponseModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


