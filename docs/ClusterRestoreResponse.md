# ClusterRestoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailedNodes** | Pointer to [**[]ClusterRestoreFailedNode**](ClusterRestoreFailedNode.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**RestoringNodes** | Pointer to **[]int32** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterRestoreResponse

`func NewClusterRestoreResponse() *ClusterRestoreResponse`

NewClusterRestoreResponse instantiates a new ClusterRestoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRestoreResponseWithDefaults

`func NewClusterRestoreResponseWithDefaults() *ClusterRestoreResponse`

NewClusterRestoreResponseWithDefaults instantiates a new ClusterRestoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailedNodes

`func (o *ClusterRestoreResponse) GetFailedNodes() []ClusterRestoreFailedNode`

GetFailedNodes returns the FailedNodes field if non-nil, zero value otherwise.

### GetFailedNodesOk

`func (o *ClusterRestoreResponse) GetFailedNodesOk() (*[]ClusterRestoreFailedNode, bool)`

GetFailedNodesOk returns a tuple with the FailedNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedNodes

`func (o *ClusterRestoreResponse) SetFailedNodes(v []ClusterRestoreFailedNode)`

SetFailedNodes sets FailedNodes field to given value.

### HasFailedNodes

`func (o *ClusterRestoreResponse) HasFailedNodes() bool`

HasFailedNodes returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterRestoreResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterRestoreResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterRestoreResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterRestoreResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRestoringNodes

`func (o *ClusterRestoreResponse) GetRestoringNodes() []int32`

GetRestoringNodes returns the RestoringNodes field if non-nil, zero value otherwise.

### GetRestoringNodesOk

`func (o *ClusterRestoreResponse) GetRestoringNodesOk() (*[]int32, bool)`

GetRestoringNodesOk returns a tuple with the RestoringNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoringNodes

`func (o *ClusterRestoreResponse) SetRestoringNodes(v []int32)`

SetRestoringNodes sets RestoringNodes field to given value.

### HasRestoringNodes

`func (o *ClusterRestoreResponse) HasRestoringNodes() bool`

HasRestoringNodes returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterRestoreResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterRestoreResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterRestoreResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterRestoreResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


