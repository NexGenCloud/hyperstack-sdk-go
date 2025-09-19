# ImportKeypairPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnvironmentName** | **string** | The name of the environment where the key pair is being created. | 
**Name** | **string** | The name of the key pair that is being created. | 
**PublicKey** | **string** | The public key that is being used to import an SSH key pair. | 

## Methods

### NewImportKeypairPayload

`func NewImportKeypairPayload(environmentName string, name string, publicKey string, ) *ImportKeypairPayload`

NewImportKeypairPayload instantiates a new ImportKeypairPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportKeypairPayloadWithDefaults

`func NewImportKeypairPayloadWithDefaults() *ImportKeypairPayload`

NewImportKeypairPayloadWithDefaults instantiates a new ImportKeypairPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironmentName

`func (o *ImportKeypairPayload) GetEnvironmentName() string`

GetEnvironmentName returns the EnvironmentName field if non-nil, zero value otherwise.

### GetEnvironmentNameOk

`func (o *ImportKeypairPayload) GetEnvironmentNameOk() (*string, bool)`

GetEnvironmentNameOk returns a tuple with the EnvironmentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentName

`func (o *ImportKeypairPayload) SetEnvironmentName(v string)`

SetEnvironmentName sets EnvironmentName field to given value.


### GetName

`func (o *ImportKeypairPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImportKeypairPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImportKeypairPayload) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *ImportKeypairPayload) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ImportKeypairPayload) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ImportKeypairPayload) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


