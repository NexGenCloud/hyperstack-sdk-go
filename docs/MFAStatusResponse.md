# MFAStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Mfa** | Pointer to [**MFAStatusFields**](MFAStatusFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewMFAStatusResponse

`func NewMFAStatusResponse() *MFAStatusResponse`

NewMFAStatusResponse instantiates a new MFAStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAStatusResponseWithDefaults

`func NewMFAStatusResponseWithDefaults() *MFAStatusResponse`

NewMFAStatusResponseWithDefaults instantiates a new MFAStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *MFAStatusResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MFAStatusResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MFAStatusResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MFAStatusResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMfa

`func (o *MFAStatusResponse) GetMfa() MFAStatusFields`

GetMfa returns the Mfa field if non-nil, zero value otherwise.

### GetMfaOk

`func (o *MFAStatusResponse) GetMfaOk() (*MFAStatusFields, bool)`

GetMfaOk returns a tuple with the Mfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfa

`func (o *MFAStatusResponse) SetMfa(v MFAStatusFields)`

SetMfa sets Mfa field to given value.

### HasMfa

`func (o *MFAStatusResponse) HasMfa() bool`

HasMfa returns a boolean if a field has been set.

### GetStatus

`func (o *MFAStatusResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MFAStatusResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MFAStatusResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MFAStatusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


