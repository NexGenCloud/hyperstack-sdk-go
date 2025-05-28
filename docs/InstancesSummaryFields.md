# InstancesSummaryFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Environment** | Pointer to **string** |  | [optional] 
**EnvironmentId** | Pointer to **int32** |  | [optional] 
**FixedIp** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to **string** |  | [optional] 
**FlavorId** | Pointer to **int32** |  | [optional] 
**FloatingIp** | Pointer to **string** |  | [optional] 
**FloatingIpStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**ImageId** | Pointer to **int32** |  | [optional] 
**Keypair** | Pointer to **string** |  | [optional] 
**KeypairId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewInstancesSummaryFields

`func NewInstancesSummaryFields() *InstancesSummaryFields`

NewInstancesSummaryFields instantiates a new InstancesSummaryFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstancesSummaryFieldsWithDefaults

`func NewInstancesSummaryFieldsWithDefaults() *InstancesSummaryFields`

NewInstancesSummaryFieldsWithDefaults instantiates a new InstancesSummaryFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *InstancesSummaryFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *InstancesSummaryFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *InstancesSummaryFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *InstancesSummaryFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *InstancesSummaryFields) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InstancesSummaryFields) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InstancesSummaryFields) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InstancesSummaryFields) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *InstancesSummaryFields) GetEnvironmentId() int32`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *InstancesSummaryFields) GetEnvironmentIdOk() (*int32, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *InstancesSummaryFields) SetEnvironmentId(v int32)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *InstancesSummaryFields) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetFixedIp

`func (o *InstancesSummaryFields) GetFixedIp() string`

GetFixedIp returns the FixedIp field if non-nil, zero value otherwise.

### GetFixedIpOk

`func (o *InstancesSummaryFields) GetFixedIpOk() (*string, bool)`

GetFixedIpOk returns a tuple with the FixedIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedIp

`func (o *InstancesSummaryFields) SetFixedIp(v string)`

SetFixedIp sets FixedIp field to given value.

### HasFixedIp

`func (o *InstancesSummaryFields) HasFixedIp() bool`

HasFixedIp returns a boolean if a field has been set.

### GetFlavor

`func (o *InstancesSummaryFields) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *InstancesSummaryFields) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *InstancesSummaryFields) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *InstancesSummaryFields) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetFlavorId

`func (o *InstancesSummaryFields) GetFlavorId() int32`

GetFlavorId returns the FlavorId field if non-nil, zero value otherwise.

### GetFlavorIdOk

`func (o *InstancesSummaryFields) GetFlavorIdOk() (*int32, bool)`

GetFlavorIdOk returns a tuple with the FlavorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavorId

`func (o *InstancesSummaryFields) SetFlavorId(v int32)`

SetFlavorId sets FlavorId field to given value.

### HasFlavorId

`func (o *InstancesSummaryFields) HasFlavorId() bool`

HasFlavorId returns a boolean if a field has been set.

### GetFloatingIp

`func (o *InstancesSummaryFields) GetFloatingIp() string`

GetFloatingIp returns the FloatingIp field if non-nil, zero value otherwise.

### GetFloatingIpOk

`func (o *InstancesSummaryFields) GetFloatingIpOk() (*string, bool)`

GetFloatingIpOk returns a tuple with the FloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIp

`func (o *InstancesSummaryFields) SetFloatingIp(v string)`

SetFloatingIp sets FloatingIp field to given value.

### HasFloatingIp

`func (o *InstancesSummaryFields) HasFloatingIp() bool`

HasFloatingIp returns a boolean if a field has been set.

### GetFloatingIpStatus

`func (o *InstancesSummaryFields) GetFloatingIpStatus() string`

GetFloatingIpStatus returns the FloatingIpStatus field if non-nil, zero value otherwise.

### GetFloatingIpStatusOk

`func (o *InstancesSummaryFields) GetFloatingIpStatusOk() (*string, bool)`

GetFloatingIpStatusOk returns a tuple with the FloatingIpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpStatus

`func (o *InstancesSummaryFields) SetFloatingIpStatus(v string)`

SetFloatingIpStatus sets FloatingIpStatus field to given value.

### HasFloatingIpStatus

`func (o *InstancesSummaryFields) HasFloatingIpStatus() bool`

HasFloatingIpStatus returns a boolean if a field has been set.

### GetId

`func (o *InstancesSummaryFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstancesSummaryFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstancesSummaryFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InstancesSummaryFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *InstancesSummaryFields) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *InstancesSummaryFields) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *InstancesSummaryFields) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *InstancesSummaryFields) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetImageId

`func (o *InstancesSummaryFields) GetImageId() int32`

GetImageId returns the ImageId field if non-nil, zero value otherwise.

### GetImageIdOk

`func (o *InstancesSummaryFields) GetImageIdOk() (*int32, bool)`

GetImageIdOk returns a tuple with the ImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageId

`func (o *InstancesSummaryFields) SetImageId(v int32)`

SetImageId sets ImageId field to given value.

### HasImageId

`func (o *InstancesSummaryFields) HasImageId() bool`

HasImageId returns a boolean if a field has been set.

### GetKeypair

`func (o *InstancesSummaryFields) GetKeypair() string`

GetKeypair returns the Keypair field if non-nil, zero value otherwise.

### GetKeypairOk

`func (o *InstancesSummaryFields) GetKeypairOk() (*string, bool)`

GetKeypairOk returns a tuple with the Keypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypair

`func (o *InstancesSummaryFields) SetKeypair(v string)`

SetKeypair sets Keypair field to given value.

### HasKeypair

`func (o *InstancesSummaryFields) HasKeypair() bool`

HasKeypair returns a boolean if a field has been set.

### GetKeypairId

`func (o *InstancesSummaryFields) GetKeypairId() int32`

GetKeypairId returns the KeypairId field if non-nil, zero value otherwise.

### GetKeypairIdOk

`func (o *InstancesSummaryFields) GetKeypairIdOk() (*int32, bool)`

GetKeypairIdOk returns a tuple with the KeypairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairId

`func (o *InstancesSummaryFields) SetKeypairId(v int32)`

SetKeypairId sets KeypairId field to given value.

### HasKeypairId

`func (o *InstancesSummaryFields) HasKeypairId() bool`

HasKeypairId returns a boolean if a field has been set.

### GetName

`func (o *InstancesSummaryFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstancesSummaryFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstancesSummaryFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstancesSummaryFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrgId

`func (o *InstancesSummaryFields) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *InstancesSummaryFields) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *InstancesSummaryFields) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *InstancesSummaryFields) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetStatus

`func (o *InstancesSummaryFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstancesSummaryFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstancesSummaryFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstancesSummaryFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *InstancesSummaryFields) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *InstancesSummaryFields) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *InstancesSummaryFields) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *InstancesSummaryFields) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


