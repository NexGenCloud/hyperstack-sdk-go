# ClusterFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiAddress** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**EnvironmentName** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**KeypairName** | Pointer to **string** |  | [optional] 
**KubeConfig** | Pointer to **string** |  | [optional] 
**KubernetesVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**NodeFlavor** | Pointer to [**InstanceFlavorFields**](InstanceFlavorFields.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusReason** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterFields

`func NewClusterFields() *ClusterFields`

NewClusterFields instantiates a new ClusterFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterFieldsWithDefaults

`func NewClusterFieldsWithDefaults() *ClusterFields`

NewClusterFieldsWithDefaults instantiates a new ClusterFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiAddress

`func (o *ClusterFields) GetApiAddress() string`

GetApiAddress returns the ApiAddress field if non-nil, zero value otherwise.

### GetApiAddressOk

`func (o *ClusterFields) GetApiAddressOk() (*string, bool)`

GetApiAddressOk returns a tuple with the ApiAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiAddress

`func (o *ClusterFields) SetApiAddress(v string)`

SetApiAddress sets ApiAddress field to given value.

### HasApiAddress

`func (o *ClusterFields) HasApiAddress() bool`

HasApiAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ClusterFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusterFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusterFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ClusterFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironmentName

`func (o *ClusterFields) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *ClusterFields) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *ClusterFields) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.

### HasEnvironmentName

`func (o *ClusterFields) HasEnvironmentName() bool`

HasEnvironmentName returns a boolean if a field has been set.

### GetId

`func (o *ClusterFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeypairName

`func (o *ClusterFields) GetKeypairName() string`

GetKeypairName returns the KeypairName field if non-nil, zero value otherwise.

### GetKeypairNameOk

`func (o *ClusterFields) GetKeypairNameOk() (*string, bool)`

GetKeypairNameOk returns a tuple with the KeypairName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeypairName

`func (o *ClusterFields) SetKeypairName(v string)`

SetKeypairName sets KeypairName field to given value.

### HasKeypairName

`func (o *ClusterFields) HasKeypairName() bool`

HasKeypairName returns a boolean if a field has been set.

### GetKubeConfig

`func (o *ClusterFields) GetKubeConfig() string`

GetKubeConfig returns the KubeConfig field if non-nil, zero value otherwise.

### GetKubeConfigOk

`func (o *ClusterFields) GetKubeConfigOk() (*string, bool)`

GetKubeConfigOk returns a tuple with the KubeConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeConfig

`func (o *ClusterFields) SetKubeConfig(v string)`

SetKubeConfig sets KubeConfig field to given value.

### HasKubeConfig

`func (o *ClusterFields) HasKubeConfig() bool`

HasKubeConfig returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *ClusterFields) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *ClusterFields) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *ClusterFields) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *ClusterFields) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetName

`func (o *ClusterFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterFields) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterFields) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterFields) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterFields) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeFlavor

`func (o *ClusterFields) GetNodeFlavor() InstanceFlavorFields`

GetNodeFlavor returns the NodeFlavor field if non-nil, zero value otherwise.

### GetNodeFlavorOk

`func (o *ClusterFields) GetNodeFlavorOk() (*InstanceFlavorFields, bool)`

GetNodeFlavorOk returns a tuple with the NodeFlavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeFlavor

`func (o *ClusterFields) SetNodeFlavor(v InstanceFlavorFields)`

SetNodeFlavor sets NodeFlavor field to given value.

### HasNodeFlavor

`func (o *ClusterFields) HasNodeFlavor() bool`

HasNodeFlavor returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterFields) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterFields) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterFields) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterFields) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusReason

`func (o *ClusterFields) GetStatusReason() string`

GetStatusReason returns the StatusReason field if non-nil, zero value otherwise.

### GetStatusReasonOk

`func (o *ClusterFields) GetStatusReasonOk() (*string, bool)`

GetStatusReasonOk returns a tuple with the StatusReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReason

`func (o *ClusterFields) SetStatusReason(v string)`

SetStatusReason sets StatusReason field to given value.

### HasStatusReason

`func (o *ClusterFields) HasStatusReason() bool`

HasStatusReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


