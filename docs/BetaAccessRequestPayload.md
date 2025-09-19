# BetaAccessRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]interface{}** | Optional metadata for the request | [optional] 
**Program** | **string** | Name of the beta program | 

## Methods

### NewBetaAccessRequestPayload

`func NewBetaAccessRequestPayload(program string, ) *BetaAccessRequestPayload`

NewBetaAccessRequestPayload instantiates a new BetaAccessRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaAccessRequestPayloadWithDefaults

`func NewBetaAccessRequestPayloadWithDefaults() *BetaAccessRequestPayload`

NewBetaAccessRequestPayloadWithDefaults instantiates a new BetaAccessRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *BetaAccessRequestPayload) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BetaAccessRequestPayload) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BetaAccessRequestPayload) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BetaAccessRequestPayload) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProgram

`func (o *BetaAccessRequestPayload) GetProgram() string`

GetProgram returns the Program field if non-nil, zero value otherwise.

### GetProgramOk

`func (o *BetaAccessRequestPayload) GetProgramOk() (*string, bool)`

GetProgramOk returns a tuple with the Program field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgram

`func (o *BetaAccessRequestPayload) SetProgram(v string)`

SetProgram sets Program field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


