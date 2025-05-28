# InternalInstancesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]InternalInstanceFields**](InternalInstanceFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewInternalInstancesResponse

`func NewInternalInstancesResponse() *InternalInstancesResponse`

NewInternalInstancesResponse instantiates a new InternalInstancesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalInstancesResponseWithDefaults

`func NewInternalInstancesResponseWithDefaults() *InternalInstancesResponse`

NewInternalInstancesResponseWithDefaults instantiates a new InternalInstancesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *InternalInstancesResponse) GetInstances() []InternalInstanceFields`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *InternalInstancesResponse) GetInstancesOk() (*[]InternalInstanceFields, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *InternalInstancesResponse) SetInstances(v []InternalInstanceFields)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *InternalInstancesResponse) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetMessage

`func (o *InternalInstancesResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InternalInstancesResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InternalInstancesResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InternalInstancesResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *InternalInstancesResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalInstancesResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalInstancesResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalInstancesResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


