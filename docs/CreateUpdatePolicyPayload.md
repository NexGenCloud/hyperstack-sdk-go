# CreateUpdatePolicyPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** |  | 
**IsPublic** | **bool** |  | 
**Name** | **string** |  | 
**Permissions** | **[]int32** |  | 

## Methods

### NewCreateUpdatePolicyPayload

`func NewCreateUpdatePolicyPayload(description string, isPublic bool, name string, permissions []int32, ) *CreateUpdatePolicyPayload`

NewCreateUpdatePolicyPayload instantiates a new CreateUpdatePolicyPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUpdatePolicyPayloadWithDefaults

`func NewCreateUpdatePolicyPayloadWithDefaults() *CreateUpdatePolicyPayload`

NewCreateUpdatePolicyPayloadWithDefaults instantiates a new CreateUpdatePolicyPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateUpdatePolicyPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateUpdatePolicyPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateUpdatePolicyPayload) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIsPublic

`func (o *CreateUpdatePolicyPayload) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *CreateUpdatePolicyPayload) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *CreateUpdatePolicyPayload) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.


### GetName

`func (o *CreateUpdatePolicyPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUpdatePolicyPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUpdatePolicyPayload) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *CreateUpdatePolicyPayload) GetPermissions() []int32`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *CreateUpdatePolicyPayload) GetPermissionsOk() (*[]int32, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *CreateUpdatePolicyPayload) SetPermissions(v []int32)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


