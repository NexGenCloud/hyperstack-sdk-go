# InstanceFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallbackUrl** | Pointer to **string** |  | [optional] 
**ContractId** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Environment** | Pointer to [**InstanceEnvironmentFields**](InstanceEnvironmentFields.md) |  | [optional] 
**Features** | Pointer to **map[string]interface{}** |  | [optional] 
**FixedIp** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to [**InstanceFlavorFields**](InstanceFlavorFields.md) |  | [optional] 
**FloatingIp** | Pointer to **string** |  | [optional] 
**FloatingIpStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Image** | Pointer to [**InstanceImageFields**](InstanceImageFields.md) |  | [optional] 
**Keypair** | Pointer to [**InstanceKeypairFields**](InstanceKeypairFields.md) |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Os** | Pointer to **string** |  | [optional] 
**PortRandomization** | Pointer to **bool** |  | [optional] 
**PortRandomizationStatus** | Pointer to **string** |  | [optional] 
**PowerState** | Pointer to **string** |  | [optional] 
**RequiresPublicIp** | Pointer to **bool** |  | [optional] 
**SecurityRules** | Pointer to [**[]SecurityRulesFieldsForInstance**](SecurityRulesFieldsForInstance.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VmState** | Pointer to **string** |  | [optional] 
**VolumeAttachments** | Pointer to [**[]VolumeAttachmentFields**](VolumeAttachmentFields.md) |  | [optional] 

## Methods

### NewInstanceFields

`func NewInstanceFields() *InstanceFields`

NewInstanceFields instantiates a new InstanceFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceFieldsWithDefaults

`func NewInstanceFieldsWithDefaults() *InstanceFields`

NewInstanceFieldsWithDefaults instantiates a new InstanceFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallbackUrl

`func (o *InstanceFields) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *InstanceFields) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *InstanceFields) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *InstanceFields) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetContractId

`func (o *InstanceFields) GetContractId() int32`

GetContractId returns the ContractId field if non-nil, zero value otherwise.

### GetContractIdOk

`func (o *InstanceFields) GetContractIdOk() (*int32, bool)`

GetContractIdOk returns a tuple with the ContractId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractId

`func (o *InstanceFields) SetContractId(v int32)`

SetContractId sets ContractId field to given value.

### HasContractId

`func (o *InstanceFields) HasContractId() bool`

HasContractId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *InstanceFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstanceFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstanceFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstanceFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *InstanceFields) GetEnvironment() InstanceEnvironmentFields`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InstanceFields) GetEnvironmentOk() (*InstanceEnvironmentFields, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InstanceFields) SetEnvironment(v InstanceEnvironmentFields)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InstanceFields) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetFeatures

`func (o *InstanceFields) GetFeatures() map[string]interface{}`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *InstanceFields) GetFeaturesOk() (*map[string]interface{}, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *InstanceFields) SetFeatures(v map[string]interface{})`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *InstanceFields) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFixedIp

`func (o *InstanceFields) GetFixedIp() string`

GetFixedIp returns the FixedIp field if non-nil, zero value otherwise.

### GetFixedIpOk

`func (o *InstanceFields) GetFixedIpOk() (*string, bool)`

GetFixedIpOk returns a tuple with the FixedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIp

`func (o *InstanceFields) SetFixedIp(v string)`

SetFixedIp sets FixedIp field to given value.

### HasFixedIp

`func (o *InstanceFields) HasFixedIp() bool`

HasFixedIp returns a boolean if a field has been set.

### GetFlavor

`func (o *InstanceFields) GetFlavor() InstanceFlavorFields`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *InstanceFields) GetFlavorOk() (*InstanceFlavorFields, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *InstanceFields) SetFlavor(v InstanceFlavorFields)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *InstanceFields) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetFloatingIp

`func (o *InstanceFields) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *InstanceFields) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *InstanceFields) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *InstanceFields) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetFloatingIpStatus

`func (o *InstanceFields) GetFloatingIpStatus() string`

GetFloatingIpStatus returns the FloatingIpStatus field if non-nil, zero value otherwise.

### GetFloatingIpStatusOk

`func (o *InstanceFields) GetFloatingIpStatusOk() (*string, bool)`

GetFloatingIpStatusOk returns a tuple with the FloatingIpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpStatus

`func (o *InstanceFields) SetFloatingIpStatus(v string)`

SetFloatingIpStatus sets FloatingIpStatus field to given value.

### HasFloatingIpStatus

`func (o *InstanceFields) HasFloatingIpStatus() bool`

HasFloatingIpStatus returns a boolean if a field has been set.

### GetId

`func (o *InstanceFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *InstanceFields) GetImage() InstanceImageFields`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InstanceFields) GetImageOk() (*InstanceImageFields, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InstanceFields) SetImage(v InstanceImageFields)`

SetImage sets Image field to given value.

### HasImage

`func (o *InstanceFields) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetKeypair

`func (o *InstanceFields) GetKeypair() InstanceKeypairFields`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *InstanceFields) GetKeypairOk() (*InstanceKeypairFields, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *InstanceFields) SetKeypair(v InstanceKeypairFields)`

SetKeypair sets Keypair field to given value.

### HasKeypair

`func (o *InstanceFields) HasKeypair() bool`

HasKeypair returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceFields) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceFields) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceFields) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceFields) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocked

`func (o *InstanceFields) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *InstanceFields) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *InstanceFields) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *InstanceFields) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetName

`func (o *InstanceFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOs

`func (o *InstanceFields) GetOs() string`

GetOs returns the Os field if non-nil, zero value otherwise.

### GetOsOk

`func (o *InstanceFields) GetOsOk() (*string, bool)`

GetOsOk returns a tuple with the Os field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOs

`func (o *InstanceFields) SetOs(v string)`

SetOs sets Os field to given value.

### HasOs

`func (o *InstanceFields) HasOs() bool`

HasOs returns a boolean if a field has been set.

### GetPortRandomization

`func (o *InstanceFields) GetPortRandomization() bool`

GetPortRandomization returns the PortRandomization field if non-nil, zero value otherwise.

### GetPortRandomizationOk

`func (o *InstanceFields) GetPortRandomizationOk() (*bool, bool)`

GetPortRandomizationOk returns a tuple with the PortRandomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRandomization

`func (o *InstanceFields) SetPortRandomization(v bool)`

SetPortRandomization sets PortRandomization field to given value.

### HasPortRandomization

`func (o *InstanceFields) HasPortRandomization() bool`

HasPortRandomization returns a boolean if a field has been set.

### GetPortRandomizationStatus

`func (o *InstanceFields) GetPortRandomizationStatus() string`

GetPortRandomizationStatus returns the PortRandomizationStatus field if non-nil, zero value otherwise.

### GetPortRandomizationStatusOk

`func (o *InstanceFields) GetPortRandomizationStatusOk() (*string, bool)`

GetPortRandomizationStatusOk returns a tuple with the PortRandomizationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortRandomizationStatus

`func (o *InstanceFields) SetPortRandomizationStatus(v string)`

SetPortRandomizationStatus sets PortRandomizationStatus field to given value.

### HasPortRandomizationStatus

`func (o *InstanceFields) HasPortRandomizationStatus() bool`

HasPortRandomizationStatus returns a boolean if a field has been set.

### GetPowerState

`func (o *InstanceFields) GetPowerState() string`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *InstanceFields) GetPowerStateOk() (*string, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *InstanceFields) SetPowerState(v string)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *InstanceFields) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetRequiresPublicIp

`func (o *InstanceFields) GetRequiresPublicIp() bool`

GetRequiresPublicIp returns the RequiresPublicIp field if non-nil, zero value otherwise.

### GetRequiresPublicIpOk

`func (o *InstanceFields) GetRequiresPublicIpOk() (*bool, bool)`

GetRequiresPublicIpOk returns a tuple with the RequiresPublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresPublicIp

`func (o *InstanceFields) SetRequiresPublicIp(v bool)`

SetRequiresPublicIp sets RequiresPublicIp field to given value.

### HasRequiresPublicIp

`func (o *InstanceFields) HasRequiresPublicIp() bool`

HasRequiresPublicIp returns a boolean if a field has been set.

### GetSecurityRules

`func (o *InstanceFields) GetSecurityRules() []SecurityRulesFieldsForInstance`

GetSecurityRules returns the SecurityRules field if non-nil, zero value otherwise.

### GetSecurityRulesOk

`func (o *InstanceFields) GetSecurityRulesOk() (*[]SecurityRulesFieldsForInstance, bool)`

GetSecurityRulesOk returns a tuple with the SecurityRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRules

`func (o *InstanceFields) SetSecurityRules(v []SecurityRulesFieldsForInstance)`

SetSecurityRules sets SecurityRules field to given value.

### HasSecurityRules

`func (o *InstanceFields) HasSecurityRules() bool`

HasSecurityRules returns a boolean if a field has been set.

### GetStatus

`func (o *InstanceFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVmState

`func (o *InstanceFields) GetVmState() string`

GetVmState returns the VmState field if non-nil, zero value otherwise.

### GetVmStateOk

`func (o *InstanceFields) GetVmStateOk() (*string, bool)`

GetVmStateOk returns a tuple with the VmState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmState

`func (o *InstanceFields) SetVmState(v string)`

SetVmState sets VmState field to given value.

### HasVmState

`func (o *InstanceFields) HasVmState() bool`

HasVmState returns a boolean if a field has been set.

### GetVolumeAttachments

`func (o *InstanceFields) GetVolumeAttachments() []VolumeAttachmentFields`

GetVolumeAttachments returns the VolumeAttachments field if non-nil, zero value otherwise.

### GetVolumeAttachmentsOk

`func (o *InstanceFields) GetVolumeAttachmentsOk() (*[]VolumeAttachmentFields, bool)`

GetVolumeAttachmentsOk returns a tuple with the VolumeAttachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeAttachments

`func (o *InstanceFields) SetVolumeAttachments(v []VolumeAttachmentFields)`

SetVolumeAttachments sets VolumeAttachments field to given value.

### HasVolumeAttachments

`func (o *InstanceFields) HasVolumeAttachments() bool`

HasVolumeAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


