# ConsentTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocks** | Pointer to [**[]ConsentBlock**](ConsentBlock.md) |  | [optional] 
**Version** | Pointer to **string** | Template version identifier | [optional] 

## Methods

### NewConsentTemplate

`func NewConsentTemplate() *ConsentTemplate`

NewConsentTemplate instantiates a new ConsentTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentTemplateWithDefaults

`func NewConsentTemplateWithDefaults() *ConsentTemplate`

NewConsentTemplateWithDefaults instantiates a new ConsentTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocks

`func (o *ConsentTemplate) GetBlocks() []ConsentBlock`

GetBlocks returns the Blocks field if non-nil, zero value otherwise.

### GetBlocksOk

`func (o *ConsentTemplate) GetBlocksOk() (*[]ConsentBlock, bool)`

GetBlocksOk returns a tuple with the Blocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocks

`func (o *ConsentTemplate) SetBlocks(v []ConsentBlock)`

SetBlocks sets Blocks field to given value.

### HasBlocks

`func (o *ConsentTemplate) HasBlocks() bool`

HasBlocks returns a boolean if a field has been set.

### GetVersion

`func (o *ConsentTemplate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConsentTemplate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConsentTemplate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConsentTemplate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


