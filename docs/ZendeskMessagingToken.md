# ZendeskMessagingToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** | Signed HS256 JWT to pass to the Zendesk Web Widget loginUser API | [optional] 

## Methods

### NewZendeskMessagingToken

`func NewZendeskMessagingToken() *ZendeskMessagingToken`

NewZendeskMessagingToken instantiates a new ZendeskMessagingToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewZendeskMessagingTokenWithDefaults

`func NewZendeskMessagingTokenWithDefaults() *ZendeskMessagingToken`

NewZendeskMessagingTokenWithDefaults instantiates a new ZendeskMessagingToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ZendeskMessagingToken) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ZendeskMessagingToken) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ZendeskMessagingToken) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ZendeskMessagingToken) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ZendeskMessagingToken) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ZendeskMessagingToken) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ZendeskMessagingToken) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ZendeskMessagingToken) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ZendeskMessagingToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ZendeskMessagingToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ZendeskMessagingToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ZendeskMessagingToken) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


