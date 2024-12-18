# Templates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Templates** | Pointer to [**[]TemplateFields**](TemplateFields.md) |  | [optional] 

## Methods

### NewTemplates

`func NewTemplates() *Templates`

NewTemplates instantiates a new Templates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesWithDefaults

`func NewTemplatesWithDefaults() *Templates`

NewTemplatesWithDefaults instantiates a new Templates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *Templates) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Templates) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Templates) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Templates) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *Templates) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Templates) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Templates) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Templates) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplates

`func (o *Templates) GetTemplates() []TemplateFields`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *Templates) GetTemplatesOk() (*[]TemplateFields, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *Templates) SetTemplates(v []TemplateFields)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *Templates) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


