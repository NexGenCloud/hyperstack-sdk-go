# ClusterNodesListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]ClusterNodeFields**](ClusterNodeFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterNodesListResponse

`func NewClusterNodesListResponse() *ClusterNodesListResponse`

NewClusterNodesListResponse instantiates a new ClusterNodesListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodesListResponseWithDefaults

`func NewClusterNodesListResponseWithDefaults() *ClusterNodesListResponse`

NewClusterNodesListResponseWithDefaults instantiates a new ClusterNodesListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ClusterNodesListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterNodesListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterNodesListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterNodesListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNodes

`func (o *ClusterNodesListResponse) GetNodes() []ClusterNodeFields`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ClusterNodesListResponse) GetNodesOk() (*[]ClusterNodeFields, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ClusterNodesListResponse) SetNodes(v []ClusterNodeFields)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *ClusterNodesListResponse) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterNodesListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterNodesListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterNodesListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterNodesListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


