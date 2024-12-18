# InternalInstanceFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootSource** | Pointer to **string** |  | [optional] 
**CallbackUrl** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Environment** | Pointer to [**InternalEnvironmentFields**](InternalEnvironmentFields.md) |  | [optional] 
**FixedIp** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to [**InternalInstanceFlavorFields**](InternalInstanceFlavorFields.md) |  | [optional] 
**FloatingIp** | Pointer to **string** |  | [optional] 
**FloatingIpStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Image** | Pointer to [**InternalInstanceImageFields**](InternalInstanceImageFields.md) |  | [optional] 
**Keypair** | Pointer to [**InternalInstanceKeypairFields**](InternalInstanceKeypairFields.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OpenstackId** | Pointer to **string** |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**SecurityRules** | Pointer to [**[]InternalSecurityRulesFieldsForInstance**](InternalSecurityRulesFieldsForInstance.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UserData** | Pointer to **string** |  | [optional] 
**VmState** | Pointer to **string** |  | [optional] 
**VolumeAttachments** | Pointer to [**[]InternalVolumeAttachmentFields**](InternalVolumeAttachmentFields.md) |  | [optional] 

## Methods

### NewInternalInstanceFields

`func NewInternalInstanceFields() *InternalInstanceFields`

NewInternalInstanceFields instantiates a new InternalInstanceFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalInstanceFieldsWithDefaults

`func NewInternalInstanceFieldsWithDefaults() *InternalInstanceFields`

NewInternalInstanceFieldsWithDefaults instantiates a new InternalInstanceFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootSource

`func (o *InternalInstanceFields) GetBootSource() string`

GetBootSource returns the BootSource field if non-nil, zero value otherwise.

### GetBootSourceOk

`func (o *InternalInstanceFields) GetBootSourceOk() (*string, bool)`

GetBootSourceOk returns a tuple with the BootSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootSource

`func (o *InternalInstanceFields) SetBootSource(v string)`

SetBootSource sets BootSource field to given value.

### HasBootSource

`func (o *InternalInstanceFields) HasBootSource() bool`

HasBootSource returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *InternalInstanceFields) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *InternalInstanceFields) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *InternalInstanceFields) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *InternalInstanceFields) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InternalInstanceFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InternalInstanceFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InternalInstanceFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InternalInstanceFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *InternalInstanceFields) GetEnvironment() InternalEnvironmentFields`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InternalInstanceFields) GetEnvironmentOk() (*InternalEnvironmentFields, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InternalInstanceFields) SetEnvironment(v InternalEnvironmentFields)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InternalInstanceFields) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetFixedIp

`func (o *InternalInstanceFields) GetFixedIp() string`

GetFixedIp returns the FixedIp field if non-nil, zero value otherwise.

### GetFixedIpOk

`func (o *InternalInstanceFields) GetFixedIpOk() (*string, bool)`

GetFixedIpOk returns a tuple with the FixedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIp

`func (o *InternalInstanceFields) SetFixedIp(v string)`

SetFixedIp sets FixedIp field to given value.

### HasFixedIp

`func (o *InternalInstanceFields) HasFixedIp() bool`

HasFixedIp returns a boolean if a field has been set.

### GetFlavor

`func (o *InternalInstanceFields) GetFlavor() InternalInstanceFlavorFields`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *InternalInstanceFields) GetFlavorOk() (*InternalInstanceFlavorFields, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *InternalInstanceFields) SetFlavor(v InternalInstanceFlavorFields)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *InternalInstanceFields) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetFloatingIp

`func (o *InternalInstanceFields) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *InternalInstanceFields) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *InternalInstanceFields) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *InternalInstanceFields) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetFloatingIpStatus

`func (o *InternalInstanceFields) GetFloatingIpStatus() string`

GetFloatingIpStatus returns the FloatingIpStatus field if non-nil, zero value otherwise.

### GetFloatingIpStatusOk

`func (o *InternalInstanceFields) GetFloatingIpStatusOk() (*string, bool)`

GetFloatingIpStatusOk returns a tuple with the FloatingIpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpStatus

`func (o *InternalInstanceFields) SetFloatingIpStatus(v string)`

SetFloatingIpStatus sets FloatingIpStatus field to given value.

### HasFloatingIpStatus

`func (o *InternalInstanceFields) HasFloatingIpStatus() bool`

HasFloatingIpStatus returns a boolean if a field has been set.

### GetId

`func (o *InternalInstanceFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalInstanceFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalInstanceFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InternalInstanceFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *InternalInstanceFields) GetImage() InternalInstanceImageFields`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InternalInstanceFields) GetImageOk() (*InternalInstanceImageFields, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InternalInstanceFields) SetImage(v InternalInstanceImageFields)`

SetImage sets Image field to given value.

### HasImage

`func (o *InternalInstanceFields) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetKeypair

`func (o *InternalInstanceFields) GetKeypair() InternalInstanceKeypairFields`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *InternalInstanceFields) GetKeypairOk() (*InternalInstanceKeypairFields, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *InternalInstanceFields) SetKeypair(v InternalInstanceKeypairFields)`

SetKeypair sets Keypair field to given value.

### HasKeypair

`func (o *InternalInstanceFields) HasKeypair() bool`

HasKeypair returns a boolean if a field has been set.

### GetName

`func (o *InternalInstanceFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InternalInstanceFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InternalInstanceFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InternalInstanceFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOpenstackId

`func (o *InternalInstanceFields) GetOpenstackId() string`

GetOpenstackId returns the OpenstackId field if non-nil, zero value otherwise.

### GetOpenstackIdOk

`func (o *InternalInstanceFields) GetOpenstackIdOk() (*string, bool)`

GetOpenstackIdOk returns a tuple with the OpenstackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenstackId

`func (o *InternalInstanceFields) SetOpenstackId(v string)`

SetOpenstackId sets OpenstackId field to given value.

### HasOpenstackId

`func (o *InternalInstanceFields) HasOpenstackId() bool`

HasOpenstackId returns a boolean if a field has been set.

### GetPowerState

`func (o *InternalInstanceFields) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *InternalInstanceFields) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *InternalInstanceFields) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *InternalInstanceFields) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetSecurityRules

`func (o *InternalInstanceFields) GetSecurityRules() []InternalSecurityRulesFieldsForInstance`

GetSecurityRules returns the SecurityRules field if non-nil, zero value otherwise.

### GetSecurityRulesOk

`func (o *InternalInstanceFields) GetSecurityRulesOk() (*[]InternalSecurityRulesFieldsForInstance, bool)`

GetSecurityRulesOk returns a tuple with the SecurityRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRules

`func (o *InternalInstanceFields) SetSecurityRules(v []InternalSecurityRulesFieldsForInstance)`

SetSecurityRules sets SecurityRules field to given value.

### HasSecurityRules

`func (o *InternalInstanceFields) HasSecurityRules() bool`

HasSecurityRules returns a boolean if a field has been set.

### GetStatus

`func (o *InternalInstanceFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InternalInstanceFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InternalInstanceFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InternalInstanceFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUserData

`func (o *InternalInstanceFields) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *InternalInstanceFields) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *InternalInstanceFields) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *InternalInstanceFields) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetVmState

`func (o *InternalInstanceFields) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *InternalInstanceFields) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *InternalInstanceFields) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *InternalInstanceFields) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### GetVolumeAttachments

`func (o *InternalInstanceFields) GetVolumeAttachments() []InternalVolumeAttachmentFields`

GetVolumeAttachments returns the VolumeAttachments field if non-nil, zero value otherwise.

### GetVolumeAttachmentsOk

`func (o *InternalInstanceFields) GetVolumeAttachmentsOk() (*[]InternalVolumeAttachmentFields, bool)`

GetVolumeAttachmentsOk returns a tuple with the VolumeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttachments

`func (o *InternalInstanceFields) SetVolumeAttachments(v []InternalVolumeAttachmentFields)`

SetVolumeAttachments sets VolumeAttachments field to given value.

### HasVolumeAttachments

`func (o *InternalInstanceFields) HasVolumeAttachments() bool`

HasVolumeAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


