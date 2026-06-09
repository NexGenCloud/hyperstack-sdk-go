# ConsentBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Block text, may contain ${variable} placeholders | [optional] 
**Type** | Pointer to **string** | Block display type | [optional] 

## Methods

### NewConsentBlock

`func NewConsentBlock() *ConsentBlock`

NewConsentBlock instantiates a new ConsentBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentBlockWithDefaults

`func NewConsentBlockWithDefaults() *ConsentBlock`

NewConsentBlockWithDefaults instantiates a new ConsentBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ConsentBlock) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ConsentBlock) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ConsentBlock) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ConsentBlock) HasText() bool`

HasText returns a boolean if a field has been set.

### GetType

`func (o *ConsentBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConsentBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConsentBlock) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ConsentBlock) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


