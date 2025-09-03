# CreateInstancesPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssignFloatingIp** | Pointer to **bool** | When this field is set to &#x60;true&#x60;, it attaches a [public IP address](https://docs.hyperstack.cloud/docs/api-reference/core-resources/virtual-machines/floating-ip/) to the virtual machine, enabling internet accessibility. | [optional] 
**CallbackUrl** | Pointer to **string** | An optional URL where actions performed on the virtual machine will be sent. For additional information on event callbacks, [**click here**](https://docs.hyperstack.cloud/docs/api-reference/core-resources/virtual-machines/callbacks-vms/). | [optional] 
**Count** | **int32** | The number of virtual machines to be created. | 
**CreateBootableVolume** | Pointer to **bool** | Indicates whether to create a bootable volume for the virtual machine. When set to &#x60;true&#x60;, a bootable volume will be created; the default value is &#x60;false&#x60;. | [optional] 
**EnablePortRandomization** | Pointer to **bool** | Indicates whether to enable port randomization.This setting is only effective if &#39;assign_floating_ip&#39; is true. Defaults to true. | [optional] [default to true]
**EnvironmentName** | **string** | The name of the [environment](https://docs.hyperstack.cloud/docs/api-reference/core-resources/environments/) in which the virtual machine is to be created. | 
**Flavor** | Pointer to [**FlavorObjectFields**](FlavorObjectFields.md) |  | [optional] 
**FlavorName** | **string** | The name of the GPU hardware configuration ([flavor](https://docs.hyperstack.cloud/docs/hardware/flavors)) for the virtual machines being created. | 
**ImageName** | Pointer to **string** | The [operating system (OS) image](https://docs.hyperstack.cloud/docs/virtual-machines/images) name designated for installation on the virtual machine.It also accepts custom, private images, created from [existing snapshots](https://docs.hyperstack.cloud/docs/virtual-machines/custom-images). | [optional] 
**KeyName** | **string** | The name of the existing SSH key pair to be used for secure access to the virtual machine. For additional information on SSH key pairs, [**click here**](https://docs.hyperstack.cloud/docs/api-reference/core-resources/keypairs/). | 
**Labels** | Pointer to **[]string** |  | [optional] 
**Name** | **string** | The name of the virtual machine being created. | 
**Profile** | Pointer to [**ProfileObjectFields**](ProfileObjectFields.md) |  | [optional] 
**SecurityRules** | Pointer to [**[]CreateSecurityRulePayload**](CreateSecurityRulePayload.md) |  | [optional] 
**UserData** | Pointer to **string** | Optional initialization configuration commands to manage the configuration of a virtual machine at launch using cloud-init scripts. For more information about custom VM configuration using cloud-init, [**click here**](https://docs.hyperstack.cloud/docs/virtual-machines/initialization-configuration). | [optional] 
**VolumeName** | Pointer to **string** | The names of the volume(s) to be attached to the virtual machine being created. | [optional] 

## Methods

### NewCreateInstancesPayload

`func NewCreateInstancesPayload(count int32, environmentName string, flavorName string, keyName string, name string, ) *CreateInstancesPayload`

NewCreateInstancesPayload instantiates a new CreateInstancesPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstancesPayloadWithDefaults

`func NewCreateInstancesPayloadWithDefaults() *CreateInstancesPayload`

NewCreateInstancesPayloadWithDefaults instantiates a new CreateInstancesPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignFloatingIp

`func (o *CreateInstancesPayload) GetAssignFloatingIp() bool`

GetAssignFloatingIp returns the AssignFloatingIp field if non-nil, zero value otherwise.

### GetAssignFloatingIpOk

`func (o *CreateInstancesPayload) GetAssignFloatingIpOk() (*bool, bool)`

GetAssignFloatingIpOk returns a tuple with the AssignFloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignFloatingIp

`func (o *CreateInstancesPayload) SetAssignFloatingIp(v bool)`

SetAssignFloatingIp sets AssignFloatingIp field to given value.

### HasAssignFloatingIp

`func (o *CreateInstancesPayload) HasAssignFloatingIp() bool`

HasAssignFloatingIp returns a boolean if a field has been set.

### GetCallbackUrl

`func (o *CreateInstancesPayload) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *CreateInstancesPayload) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *CreateInstancesPayload) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.

### HasCallbackUrl

`func (o *CreateInstancesPayload) HasCallbackUrl() bool`

HasCallbackUrl returns a boolean if a field has been set.

### GetCount

`func (o *CreateInstancesPayload) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateInstancesPayload) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateInstancesPayload) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCreateBootableVolume

`func (o *CreateInstancesPayload) GetCreateBootableVolume() bool`

GetCreateBootableVolume returns the CreateBootableVolume field if non-nil, zero value otherwise.

### GetCreateBootableVolumeOk

`func (o *CreateInstancesPayload) GetCreateBootableVolumeOk() (*bool, bool)`

GetCreateBootableVolumeOk returns a tuple with the CreateBootableVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBootableVolume

`func (o *CreateInstancesPayload) SetCreateBootableVolume(v bool)`

SetCreateBootableVolume sets CreateBootableVolume field to given value.

### HasCreateBootableVolume

`func (o *CreateInstancesPayload) HasCreateBootableVolume() bool`

HasCreateBootableVolume returns a boolean if a field has been set.

### GetEnablePortRandomization

`func (o *CreateInstancesPayload) GetEnablePortRandomization() bool`

GetEnablePortRandomization returns the EnablePortRandomization field if non-nil, zero value otherwise.

### GetEnablePortRandomizationOk

`func (o *CreateInstancesPayload) GetEnablePortRandomizationOk() (*bool, bool)`

GetEnablePortRandomizationOk returns a tuple with the EnablePortRandomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePortRandomization

`func (o *CreateInstancesPayload) SetEnablePortRandomization(v bool)`

SetEnablePortRandomization sets EnablePortRandomization field to given value.

### HasEnablePortRandomization

`func (o *CreateInstancesPayload) HasEnablePortRandomization() bool`

HasEnablePortRandomization returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *CreateInstancesPayload) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *CreateInstancesPayload) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *CreateInstancesPayload) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetFlavor

`func (o *CreateInstancesPayload) GetFlavor() FlavorObjectFields`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *CreateInstancesPayload) GetFlavorOk() (*FlavorObjectFields, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *CreateInstancesPayload) SetFlavor(v FlavorObjectFields)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *CreateInstancesPayload) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetFlavorName

`func (o *CreateInstancesPayload) GetFlavorName() string`

GetFlavorName returns the FlavorName field if non-nil, zero value otherwise.

### GetFlavorNameOk

`func (o *CreateInstancesPayload) GetFlavorNameOk() (*string, bool)`

GetFlavorNameOk returns a tuple with the FlavorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorName

`func (o *CreateInstancesPayload) SetFlavorName(v string)`

SetFlavorName sets FlavorName field to given value.


### GetImageName

`func (o *CreateInstancesPayload) GetImageName() string`

GetImageName returns the ImageName field if non-nil, zero value otherwise.

### GetImageNameOk

`func (o *CreateInstancesPayload) GetImageNameOk() (*string, bool)`

GetImageNameOk returns a tuple with the ImageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageName

`func (o *CreateInstancesPayload) SetImageName(v string)`

SetImageName sets ImageName field to given value.

### HasImageName

`func (o *CreateInstancesPayload) HasImageName() bool`

HasImageName returns a boolean if a field has been set.

### GetKeyName

`func (o *CreateInstancesPayload) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *CreateInstancesPayload) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *CreateInstancesPayload) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetLabels

`func (o *CreateInstancesPayload) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateInstancesPayload) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateInstancesPayload) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateInstancesPayload) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetName

`func (o *CreateInstancesPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateInstancesPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateInstancesPayload) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *CreateInstancesPayload) GetProfile() ProfileObjectFields`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *CreateInstancesPayload) GetProfileOk() (*ProfileObjectFields, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *CreateInstancesPayload) SetProfile(v ProfileObjectFields)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *CreateInstancesPayload) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSecurityRules

`func (o *CreateInstancesPayload) GetSecurityRules() []CreateSecurityRulePayload`

GetSecurityRules returns the SecurityRules field if non-nil, zero value otherwise.

### GetSecurityRulesOk

`func (o *CreateInstancesPayload) GetSecurityRulesOk() (*[]CreateSecurityRulePayload, bool)`

GetSecurityRulesOk returns a tuple with the SecurityRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityRules

`func (o *CreateInstancesPayload) SetSecurityRules(v []CreateSecurityRulePayload)`

SetSecurityRules sets SecurityRules field to given value.

### HasSecurityRules

`func (o *CreateInstancesPayload) HasSecurityRules() bool`

HasSecurityRules returns a boolean if a field has been set.

### GetUserData

`func (o *CreateInstancesPayload) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CreateInstancesPayload) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CreateInstancesPayload) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CreateInstancesPayload) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetVolumeName

`func (o *CreateInstancesPayload) GetVolumeName() string`

GetVolumeName returns the VolumeName field if non-nil, zero value otherwise.

### GetVolumeNameOk

`func (o *CreateInstancesPayload) GetVolumeNameOk() (*string, bool)`

GetVolumeNameOk returns a tuple with the VolumeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeName

`func (o *CreateInstancesPayload) SetVolumeName(v string)`

SetVolumeName sets VolumeName field to given value.

### HasVolumeName

`func (o *CreateInstancesPayload) HasVolumeName() bool`

HasVolumeName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


