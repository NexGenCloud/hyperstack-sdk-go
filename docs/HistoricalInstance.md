# HistoricalInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceCount** | Pointer to **int32** |  | [optional] 
**Instances** | Pointer to [**[]HistoricalInstancesFields**](HistoricalInstancesFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewHistoricalInstance

`func NewHistoricalInstance() *HistoricalInstance`

NewHistoricalInstance instantiates a new HistoricalInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHistoricalInstanceWithDefaults

`func NewHistoricalInstanceWithDefaults() *HistoricalInstance`

NewHistoricalInstanceWithDefaults instantiates a new HistoricalInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceCount

`func (o *HistoricalInstance) GetInstanceCount() int32`

GetInstanceCount returns the InstanceCount field if non-nil, zero value otherwise.

### GetInstanceCountOk

`func (o *HistoricalInstance) GetInstanceCountOk() (*int32, bool)`

GetInstanceCountOk returns a tuple with the InstanceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceCount

`func (o *HistoricalInstance) SetInstanceCount(v int32)`

SetInstanceCount sets InstanceCount field to given value.

### HasInstanceCount

`func (o *HistoricalInstance) HasInstanceCount() bool`

HasInstanceCount returns a boolean if a field has been set.

### GetInstances

`func (o *HistoricalInstance) GetInstances() []HistoricalInstancesFields`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *HistoricalInstance) GetInstancesOk() (*[]HistoricalInstancesFields, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *HistoricalInstance) SetInstances(v []HistoricalInstancesFields)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *HistoricalInstance) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMessage

`func (o *HistoricalInstance) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HistoricalInstance) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HistoricalInstance) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HistoricalInstance) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *HistoricalInstance) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HistoricalInstance) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HistoricalInstance) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HistoricalInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


