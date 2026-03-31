# UpdateClusterNodeGroupPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallIds** | Pointer to **[]int32** | IDs of the firewalls to apply to all nodes in this node group | [optional] 
**MaxCount** | Pointer to **int32** |  | [optional] 
**MinCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateClusterNodeGroupPayload

`func NewUpdateClusterNodeGroupPayload() *UpdateClusterNodeGroupPayload`

NewUpdateClusterNodeGroupPayload instantiates a new UpdateClusterNodeGroupPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterNodeGroupPayloadWithDefaults

`func NewUpdateClusterNodeGroupPayloadWithDefaults() *UpdateClusterNodeGroupPayload`

NewUpdateClusterNodeGroupPayloadWithDefaults instantiates a new UpdateClusterNodeGroupPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewallIds

`func (o *UpdateClusterNodeGroupPayload) GetFirewallIds() []int32`

GetFirewallIds returns the FirewallIds field if non-nil, zero value otherwise.

### GetFirewallIdsOk

`func (o *UpdateClusterNodeGroupPayload) GetFirewallIdsOk() (*[]int32, bool)`

GetFirewallIdsOk returns a tuple with the FirewallIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallIds

`func (o *UpdateClusterNodeGroupPayload) SetFirewallIds(v []int32)`

SetFirewallIds sets FirewallIds field to given value.

### HasFirewallIds

`func (o *UpdateClusterNodeGroupPayload) HasFirewallIds() bool`

HasFirewallIds returns a boolean if a field has been set.

### GetMaxCount

`func (o *UpdateClusterNodeGroupPayload) GetMaxCount() int32`

GetMaxCount returns the MaxCount field if non-nil, zero value otherwise.

### GetMaxCountOk

`func (o *UpdateClusterNodeGroupPayload) GetMaxCountOk() (*int32, bool)`

GetMaxCountOk returns a tuple with the MaxCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCount

`func (o *UpdateClusterNodeGroupPayload) SetMaxCount(v int32)`

SetMaxCount sets MaxCount field to given value.

### HasMaxCount

`func (o *UpdateClusterNodeGroupPayload) HasMaxCount() bool`

HasMaxCount returns a boolean if a field has been set.

### GetMinCount

`func (o *UpdateClusterNodeGroupPayload) GetMinCount() int32`

GetMinCount returns the MinCount field if non-nil, zero value otherwise.

### GetMinCountOk

`func (o *UpdateClusterNodeGroupPayload) GetMinCountOk() (*int32, bool)`

GetMinCountOk returns a tuple with the MinCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCount

`func (o *UpdateClusterNodeGroupPayload) SetMinCount(v int32)`

SetMinCount sets MinCount field to given value.

### HasMinCount

`func (o *UpdateClusterNodeGroupPayload) HasMinCount() bool`

HasMinCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


