# ClusterNodeFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Instance** | Pointer to [**ClusterNodeInstanceFields**](ClusterNodeInstanceFields.md) |  | [optional] 
**IsBastion** | Pointer to **bool** |  | [optional] 
**NodeGroupId** | Pointer to **int32** |  | [optional] 
**NodeGroupName** | Pointer to **string** |  | [optional] 
**RequiresPublicIp** | Pointer to **bool** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusReason** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewClusterNodeFields

`func NewClusterNodeFields() *ClusterNodeFields`

NewClusterNodeFields instantiates a new ClusterNodeFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeFieldsWithDefaults

`func NewClusterNodeFieldsWithDefaults() *ClusterNodeFields`

NewClusterNodeFieldsWithDefaults instantiates a new ClusterNodeFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ClusterNodeFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterNodeFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterNodeFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClusterNodeFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ClusterNodeFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterNodeFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterNodeFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterNodeFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstance

`func (o *ClusterNodeFields) GetInstance() ClusterNodeInstanceFields`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ClusterNodeFields) GetInstanceOk() (*ClusterNodeInstanceFields, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ClusterNodeFields) SetInstance(v ClusterNodeInstanceFields)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ClusterNodeFields) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetIsBastion

`func (o *ClusterNodeFields) GetIsBastion() bool`

GetIsBastion returns the IsBastion field if non-nil, zero value otherwise.

### GetIsBastionOk

`func (o *ClusterNodeFields) GetIsBastionOk() (*bool, bool)`

GetIsBastionOk returns a tuple with the IsBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBastion

`func (o *ClusterNodeFields) SetIsBastion(v bool)`

SetIsBastion sets IsBastion field to given value.

### HasIsBastion

`func (o *ClusterNodeFields) HasIsBastion() bool`

HasIsBastion returns a boolean if a field has been set.

### GetNodeGroupId

`func (o *ClusterNodeFields) GetNodeGroupId() int32`

GetNodeGroupId returns the NodeGroupId field if non-nil, zero value otherwise.

### GetNodeGroupIdOk

`func (o *ClusterNodeFields) GetNodeGroupIdOk() (*int32, bool)`

GetNodeGroupIdOk returns a tuple with the NodeGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroupId

`func (o *ClusterNodeFields) SetNodeGroupId(v int32)`

SetNodeGroupId sets NodeGroupId field to given value.

### HasNodeGroupId

`func (o *ClusterNodeFields) HasNodeGroupId() bool`

HasNodeGroupId returns a boolean if a field has been set.

### GetNodeGroupName

`func (o *ClusterNodeFields) GetNodeGroupName() string`

GetNodeGroupName returns the NodeGroupName field if non-nil, zero value otherwise.

### GetNodeGroupNameOk

`func (o *ClusterNodeFields) GetNodeGroupNameOk() (*string, bool)`

GetNodeGroupNameOk returns a tuple with the NodeGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroupName

`func (o *ClusterNodeFields) SetNodeGroupName(v string)`

SetNodeGroupName sets NodeGroupName field to given value.

### HasNodeGroupName

`func (o *ClusterNodeFields) HasNodeGroupName() bool`

HasNodeGroupName returns a boolean if a field has been set.

### GetRequiresPublicIp

`func (o *ClusterNodeFields) GetRequiresPublicIp() bool`

GetRequiresPublicIp returns the RequiresPublicIp field if non-nil, zero value otherwise.

### GetRequiresPublicIpOk

`func (o *ClusterNodeFields) GetRequiresPublicIpOk() (*bool, bool)`

GetRequiresPublicIpOk returns a tuple with the RequiresPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPublicIp

`func (o *ClusterNodeFields) SetRequiresPublicIp(v bool)`

SetRequiresPublicIp sets RequiresPublicIp field to given value.

### HasRequiresPublicIp

`func (o *ClusterNodeFields) HasRequiresPublicIp() bool`

HasRequiresPublicIp returns a boolean if a field has been set.

### GetRole

`func (o *ClusterNodeFields) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ClusterNodeFields) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ClusterNodeFields) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *ClusterNodeFields) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterNodeFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterNodeFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterNodeFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterNodeFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *ClusterNodeFields) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *ClusterNodeFields) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *ClusterNodeFields) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *ClusterNodeFields) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ClusterNodeFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ClusterNodeFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ClusterNodeFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ClusterNodeFields) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


