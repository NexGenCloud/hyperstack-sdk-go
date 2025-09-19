# NodeStocksPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InUse** | **int32** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Total** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewNodeStocksPayload

`func NewNodeStocksPayload(inUse int32, total int32, type_ string, ) *NodeStocksPayload`

NewNodeStocksPayload instantiates a new NodeStocksPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeStocksPayloadWithDefaults

`func NewNodeStocksPayloadWithDefaults() *NodeStocksPayload`

NewNodeStocksPayloadWithDefaults instantiates a new NodeStocksPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInUse

`func (o *NodeStocksPayload) GetInUse() int32`

GetInUse returns the InUse field if non-nil, zero value otherwise.

### GetInUseOk

`func (o *NodeStocksPayload) GetInUseOk() (*int32, bool)`

GetInUseOk returns a tuple with the InUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInUse

`func (o *NodeStocksPayload) SetInUse(v int32)`

SetInUse sets InUse field to given value.


### GetName

`func (o *NodeStocksPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NodeStocksPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NodeStocksPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NodeStocksPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTotal

`func (o *NodeStocksPayload) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NodeStocksPayload) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NodeStocksPayload) SetTotal(v int32)`

SetTotal sets Total field to given value.


### GetType

`func (o *NodeStocksPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeStocksPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeStocksPayload) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


