# ClusterNodeGroupsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**NodeGroup** | Pointer to [**ClusterNodeGroupFields**](ClusterNodeGroupFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterNodeGroupsGetResponse

`func NewClusterNodeGroupsGetResponse() *ClusterNodeGroupsGetResponse`

NewClusterNodeGroupsGetResponse instantiates a new ClusterNodeGroupsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeGroupsGetResponseWithDefaults

`func NewClusterNodeGroupsGetResponseWithDefaults() *ClusterNodeGroupsGetResponse`

NewClusterNodeGroupsGetResponseWithDefaults instantiates a new ClusterNodeGroupsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ClusterNodeGroupsGetResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterNodeGroupsGetResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterNodeGroupsGetResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterNodeGroupsGetResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNodeGroup

`func (o *ClusterNodeGroupsGetResponse) GetNodeGroup() ClusterNodeGroupFields`

GetNodeGroup returns the NodeGroup field if non-nil, zero value otherwise.

### GetNodeGroupOk

`func (o *ClusterNodeGroupsGetResponse) GetNodeGroupOk() (*ClusterNodeGroupFields, bool)`

GetNodeGroupOk returns a tuple with the NodeGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroup

`func (o *ClusterNodeGroupsGetResponse) SetNodeGroup(v ClusterNodeGroupFields)`

SetNodeGroup sets NodeGroup field to given value.

### HasNodeGroup

`func (o *ClusterNodeGroupsGetResponse) HasNodeGroup() bool`

HasNodeGroup returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterNodeGroupsGetResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterNodeGroupsGetResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterNodeGroupsGetResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterNodeGroupsGetResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


