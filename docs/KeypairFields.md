# KeypairFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Environment** | Pointer to [**KeypairEnvironmentFields**](KeypairEnvironmentFields.md) |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 

## Methods

### NewKeypairFields

`func NewKeypairFields() *KeypairFields`

NewKeypairFields instantiates a new KeypairFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeypairFieldsWithDefaults

`func NewKeypairFieldsWithDefaults() *KeypairFields`

NewKeypairFieldsWithDefaults instantiates a new KeypairFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *KeypairFields) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *KeypairFields) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *KeypairFields) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *KeypairFields) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *KeypairFields) GetEnvironment() KeypairEnvironmentFields`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *KeypairFields) GetEnvironmentOk() (*KeypairEnvironmentFields, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *KeypairFields) SetEnvironment(v KeypairEnvironmentFields)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *KeypairFields) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetFingerprint

`func (o *KeypairFields) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *KeypairFields) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *KeypairFields) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *KeypairFields) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetId

`func (o *KeypairFields) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeypairFields) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeypairFields) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeypairFields) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeypairFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeypairFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeypairFields) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeypairFields) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicKey

`func (o *KeypairFields) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *KeypairFields) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *KeypairFields) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *KeypairFields) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


