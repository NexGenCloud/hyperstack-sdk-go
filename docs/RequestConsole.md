# RequestConsole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewRequestConsole

`func NewRequestConsole() *RequestConsole`

NewRequestConsole instantiates a new RequestConsole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestConsoleWithDefaults

`func NewRequestConsoleWithDefaults() *RequestConsole`

NewRequestConsoleWithDefaults instantiates a new RequestConsole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *RequestConsole) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *RequestConsole) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *RequestConsole) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *RequestConsole) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *RequestConsole) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestConsole) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestConsole) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RequestConsole) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUrl

`func (o *RequestConsole) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RequestConsole) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RequestConsole) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RequestConsole) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


