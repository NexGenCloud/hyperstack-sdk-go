# ClusterNodeGroupsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**NodeGroups** | Pointer to [**[]ClusterNodeGroupFields**](ClusterNodeGroupFields.md) |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterNodeGroupsListResponse

`func NewClusterNodeGroupsListResponse() *ClusterNodeGroupsListResponse`

NewClusterNodeGroupsListResponse instantiates a new ClusterNodeGroupsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterNodeGroupsListResponseWithDefaults

`func NewClusterNodeGroupsListResponseWithDefaults() *ClusterNodeGroupsListResponse`

NewClusterNodeGroupsListResponseWithDefaults instantiates a new ClusterNodeGroupsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ClusterNodeGroupsListResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterNodeGroupsListResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterNodeGroupsListResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterNodeGroupsListResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNodeGroups

`func (o *ClusterNodeGroupsListResponse) GetNodeGroups() []ClusterNodeGroupFields`

GetNodeGroups returns the NodeGroups field if non-nil, zero value otherwise.

### GetNodeGroupsOk

`func (o *ClusterNodeGroupsListResponse) GetNodeGroupsOk() (*[]ClusterNodeGroupFields, bool)`

GetNodeGroupsOk returns a tuple with the NodeGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeGroups

`func (o *ClusterNodeGroupsListResponse) SetNodeGroups(v []ClusterNodeGroupFields)`

SetNodeGroups sets NodeGroups field to given value.

### HasNodeGroups

`func (o *ClusterNodeGroupsListResponse) HasNodeGroups() bool`

HasNodeGroups returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterNodeGroupsListResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterNodeGroupsListResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterNodeGroupsListResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterNodeGroupsListResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


