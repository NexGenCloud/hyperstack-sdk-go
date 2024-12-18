# StartDeploymentPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**Name** | **string** |  | 
**TemplateId** | **int32** |  | 
**Variables** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewStartDeploymentPayload

`func NewStartDeploymentPayload(description string, name string, templateId int32, ) *StartDeploymentPayload`

NewStartDeploymentPayload instantiates a new StartDeploymentPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartDeploymentPayloadWithDefaults

`func NewStartDeploymentPayloadWithDefaults() *StartDeploymentPayload`

NewStartDeploymentPayloadWithDefaults instantiates a new StartDeploymentPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *StartDeploymentPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StartDeploymentPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StartDeploymentPayload) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetName

`func (o *StartDeploymentPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StartDeploymentPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StartDeploymentPayload) SetName(v string)`

SetName sets Name field to given value.


### GetTemplateId

`func (o *StartDeploymentPayload) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *StartDeploymentPayload) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *StartDeploymentPayload) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.


### GetVariables

`func (o *StartDeploymentPayload) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *StartDeploymentPayload) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *StartDeploymentPayload) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *StartDeploymentPayload) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


