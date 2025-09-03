# ClusterNodeGroupsCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**NodeGroup** | Pointer to [**ClusterNodeGroupFields**](ClusterNodeGroupFields.md) |  | [optional] 
**Nodes** | Pointer to [**[]ClusterNodeFields**](ClusterNodeFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterNodeGroupsCreateResponse

`func NewClusterNodeGroupsCreateResponse() *ClusterNodeGroupsCreateResponse`

NewClusterNodeGroupsCreateResponse instantiates a new ClusterNodeGroupsCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeGroupsCreateResponseWithDefaults

`func NewClusterNodeGroupsCreateResponseWithDefaults() *ClusterNodeGroupsCreateResponse`

NewClusterNodeGroupsCreateResponseWithDefaults instantiates a new ClusterNodeGroupsCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ClusterNodeGroupsCreateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterNodeGroupsCreateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterNodeGroupsCreateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterNodeGroupsCreateResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNodeGroup

`func (o *ClusterNodeGroupsCreateResponse) GetNodeGroup() ClusterNodeGroupFields`

GetNodeGroup returns the NodeGroup field if non-nil, zero value otherwise.

### GetNodeGroupOk

`func (o *ClusterNodeGroupsCreateResponse) GetNodeGroupOk() (*ClusterNodeGroupFields, bool)`

GetNodeGroupOk returns a tuple with the NodeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroup

`func (o *ClusterNodeGroupsCreateResponse) SetNodeGroup(v ClusterNodeGroupFields)`

SetNodeGroup sets NodeGroup field to given value.

### HasNodeGroup

`func (o *ClusterNodeGroupsCreateResponse) HasNodeGroup() bool`

HasNodeGroup returns a boolean if a field has been set.

### GetNodes

`func (o *ClusterNodeGroupsCreateResponse) GetNodes() []ClusterNodeFields`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterNodeGroupsCreateResponse) GetNodesOk() (*[]ClusterNodeFields, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterNodeGroupsCreateResponse) SetNodes(v []ClusterNodeFields)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterNodeGroupsCreateResponse) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterNodeGroupsCreateResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterNodeGroupsCreateResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterNodeGroupsCreateResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterNodeGroupsCreateResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


