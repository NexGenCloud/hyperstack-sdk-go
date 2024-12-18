# ClusterEvents

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterEvents** | Pointer to [**[]ClusterEventsFields**](ClusterEventsFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterEvents

`func NewClusterEvents() *ClusterEvents`

NewClusterEvents instantiates a new ClusterEvents object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterEventsWithDefaults

`func NewClusterEventsWithDefaults() *ClusterEvents`

NewClusterEventsWithDefaults instantiates a new ClusterEvents object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterEvents

`func (o *ClusterEvents) GetClusterEvents() []ClusterEventsFields`

GetClusterEvents returns the ClusterEvents field if non-nil, zero value otherwise.

### GetClusterEventsOk

`func (o *ClusterEvents) GetClusterEventsOk() (*[]ClusterEventsFields, bool)`

GetClusterEventsOk returns a tuple with the ClusterEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterEvents

`func (o *ClusterEvents) SetClusterEvents(v []ClusterEventsFields)`

SetClusterEvents sets ClusterEvents field to given value.

### HasClusterEvents

`func (o *ClusterEvents) HasClusterEvents() bool`

HasClusterEvents returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterEvents) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterEvents) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterEvents) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterEvents) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterEvents) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterEvents) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterEvents) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterEvents) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


