# CreateFirewallPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the firewall. | [optional] 
**EnvironmentId** | **int32** | ID of the environment. | 
**Name** | **string** | Name of the firewall. | 

## Methods

### NewCreateFirewallPayload

`func NewCreateFirewallPayload(environmentId int32, name string, ) *CreateFirewallPayload`

NewCreateFirewallPayload instantiates a new CreateFirewallPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFirewallPayloadWithDefaults

`func NewCreateFirewallPayloadWithDefaults() *CreateFirewallPayload`

NewCreateFirewallPayloadWithDefaults instantiates a new CreateFirewallPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateFirewallPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateFirewallPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateFirewallPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateFirewallPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *CreateFirewallPayload) GetEnvironmentId() int32`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CreateFirewallPayload) GetEnvironmentIdOk() (*int32, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CreateFirewallPayload) SetEnvironmentId(v int32)`

SetEnvironmentId sets EnvironmentId field to given value.


### GetName

`func (o *CreateFirewallPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateFirewallPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateFirewallPayload) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


