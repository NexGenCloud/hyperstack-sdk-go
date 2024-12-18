# NodePayloadModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | Pointer to [**[]NodePowerUsageModel**](NodePowerUsageModel.md) |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewNodePayloadModel

`func NewNodePayloadModel() *NodePayloadModel`

NewNodePayloadModel instantiates a new NodePayloadModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodePayloadModelWithDefaults

`func NewNodePayloadModelWithDefaults() *NodePayloadModel`

NewNodePayloadModelWithDefaults instantiates a new NodePayloadModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *NodePayloadModel) GetNodes() []NodePowerUsageModel`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *NodePayloadModel) GetNodesOk() (*[]NodePowerUsageModel, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *NodePayloadModel) SetNodes(v []NodePowerUsageModel)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *NodePayloadModel) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetRegion

`func (o *NodePayloadModel) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NodePayloadModel) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NodePayloadModel) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NodePayloadModel) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


