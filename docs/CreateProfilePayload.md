# CreateProfilePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]string** | The data object containing the configuration details of the virtual machine profile being created. | 
**Description** | Pointer to **string** | The optional description for the profile being created. | [optional] 
**Name** | **string** | The name of the profile being created. | 

## Methods

### NewCreateProfilePayload

`func NewCreateProfilePayload(data map[string]string, name string, ) *CreateProfilePayload`

NewCreateProfilePayload instantiates a new CreateProfilePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateProfilePayloadWithDefaults

`func NewCreateProfilePayloadWithDefaults() *CreateProfilePayload`

NewCreateProfilePayloadWithDefaults instantiates a new CreateProfilePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateProfilePayload) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateProfilePayload) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateProfilePayload) SetData(v map[string]string)`

SetData sets Data field to given value.


### GetDescription

`func (o *CreateProfilePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateProfilePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateProfilePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateProfilePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateProfilePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateProfilePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateProfilePayload) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


