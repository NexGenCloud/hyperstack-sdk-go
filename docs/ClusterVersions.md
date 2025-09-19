# ClusterVersions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to [**[]ClusterVersion**](ClusterVersion.md) |  | [optional] 

## Methods

### NewClusterVersions

`func NewClusterVersions() *ClusterVersions`

NewClusterVersions instantiates a new ClusterVersions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterVersionsWithDefaults

`func NewClusterVersionsWithDefaults() *ClusterVersions`

NewClusterVersionsWithDefaults instantiates a new ClusterVersions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ClusterVersions) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterVersions) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterVersions) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterVersions) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterVersions) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterVersions) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterVersions) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterVersions) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersions

`func (o *ClusterVersions) GetVersions() []ClusterVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ClusterVersions) GetVersionsOk() (*[]ClusterVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ClusterVersions) SetVersions(v []ClusterVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ClusterVersions) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


