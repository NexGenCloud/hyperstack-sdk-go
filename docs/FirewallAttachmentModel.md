# FirewallAttachmentModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Vm** | Pointer to [**FirewallAttachmentVMModel**](FirewallAttachmentVMModel.md) |  | [optional] 

## Methods

### NewFirewallAttachmentModel

`func NewFirewallAttachmentModel() *FirewallAttachmentModel`

NewFirewallAttachmentModel instantiates a new FirewallAttachmentModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirewallAttachmentModelWithDefaults

`func NewFirewallAttachmentModelWithDefaults() *FirewallAttachmentModel`

NewFirewallAttachmentModelWithDefaults instantiates a new FirewallAttachmentModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *FirewallAttachmentModel) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FirewallAttachmentModel) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FirewallAttachmentModel) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FirewallAttachmentModel) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *FirewallAttachmentModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirewallAttachmentModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirewallAttachmentModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FirewallAttachmentModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *FirewallAttachmentModel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirewallAttachmentModel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirewallAttachmentModel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirewallAttachmentModel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVm

`func (o *FirewallAttachmentModel) GetVm() FirewallAttachmentVMModel`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *FirewallAttachmentModel) GetVmOk() (*FirewallAttachmentVMModel, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *FirewallAttachmentModel) SetVm(v FirewallAttachmentVMModel)`

SetVm sets Vm field to given value.

### HasVm

`func (o *FirewallAttachmentModel) HasVm() bool`

HasVm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


