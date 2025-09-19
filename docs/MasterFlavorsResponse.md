# MasterFlavorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flavors** | Pointer to [**[]ClusterFlavorFields**](ClusterFlavorFields.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 

## Methods

### NewMasterFlavorsResponse

`func NewMasterFlavorsResponse() *MasterFlavorsResponse`

NewMasterFlavorsResponse instantiates a new MasterFlavorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMasterFlavorsResponseWithDefaults

`func NewMasterFlavorsResponseWithDefaults() *MasterFlavorsResponse`

NewMasterFlavorsResponseWithDefaults instantiates a new MasterFlavorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlavors

`func (o *MasterFlavorsResponse) GetFlavors() []ClusterFlavorFields`

GetFlavors returns the Flavors field if non-nil, zero value otherwise.

### GetFlavorsOk

`func (o *MasterFlavorsResponse) GetFlavorsOk() (*[]ClusterFlavorFields, bool)`

GetFlavorsOk returns a tuple with the Flavors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavors

`func (o *MasterFlavorsResponse) SetFlavors(v []ClusterFlavorFields)`

SetFlavors sets Flavors field to given value.

### HasFlavors

`func (o *MasterFlavorsResponse) HasFlavors() bool`

HasFlavors returns a boolean if a field has been set.

### GetMessage

`func (o *MasterFlavorsResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MasterFlavorsResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MasterFlavorsResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *MasterFlavorsResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *MasterFlavorsResponse) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MasterFlavorsResponse) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MasterFlavorsResponse) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MasterFlavorsResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


