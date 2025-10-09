# CreateClusterNodeGroupPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** |  | [optional] 
**FlavorName** | **string** |  | 
**MaxCount** | Pointer to **int32** |  | [optional] 
**MinCount** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Role** | **string** |  | [default to "worker"]

## Methods

### NewCreateClusterNodeGroupPayload

`func NewCreateClusterNodeGroupPayload(flavorName string, name string, role string, ) *CreateClusterNodeGroupPayload`

NewCreateClusterNodeGroupPayload instantiates a new CreateClusterNodeGroupPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateClusterNodeGroupPayloadWithDefaults

`func NewCreateClusterNodeGroupPayloadWithDefaults() *CreateClusterNodeGroupPayload`

NewCreateClusterNodeGroupPayloadWithDefaults instantiates a new CreateClusterNodeGroupPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CreateClusterNodeGroupPayload) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateClusterNodeGroupPayload) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateClusterNodeGroupPayload) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *CreateClusterNodeGroupPayload) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFlavorName

`func (o *CreateClusterNodeGroupPayload) GetFlavorName() string`

GetFlavorName returns the FlavorName field if non-nil, zero value otherwise.

### GetFlavorNameOk

`func (o *CreateClusterNodeGroupPayload) GetFlavorNameOk() (*string, bool)`

GetFlavorNameOk returns a tuple with the FlavorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorName

`func (o *CreateClusterNodeGroupPayload) SetFlavorName(v string)`

SetFlavorName sets FlavorName field to given value.


### GetMaxCount

`func (o *CreateClusterNodeGroupPayload) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *CreateClusterNodeGroupPayload) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *CreateClusterNodeGroupPayload) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *CreateClusterNodeGroupPayload) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetMinCount

`func (o *CreateClusterNodeGroupPayload) GetMinCount() int32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *CreateClusterNodeGroupPayload) GetMinCountOk() (*int32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *CreateClusterNodeGroupPayload) SetMinCount(v int32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *CreateClusterNodeGroupPayload) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.

### GetName

`func (o *CreateClusterNodeGroupPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateClusterNodeGroupPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateClusterNodeGroupPayload) SetName(v string)`

SetName sets Name field to given value.


### GetRole

`func (o *CreateClusterNodeGroupPayload) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateClusterNodeGroupPayload) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateClusterNodeGroupPayload) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


