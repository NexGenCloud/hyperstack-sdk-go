# UpdateConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ConsentMethod** | Pointer to **string** |  | [optional] [default to "web_checkbox"]
**ConsentVersion** | Pointer to **string** | Version identifier. Defaults to the current version for the consent type | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Consent-type-specific metadata | [optional] 

## Methods

### NewUpdateConsentRequest

`func NewUpdateConsentRequest(action string, ) *UpdateConsentRequest`

NewUpdateConsentRequest instantiates a new UpdateConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConsentRequestWithDefaults

`func NewUpdateConsentRequestWithDefaults() *UpdateConsentRequest`

NewUpdateConsentRequestWithDefaults instantiates a new UpdateConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *UpdateConsentRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *UpdateConsentRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *UpdateConsentRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetConsentMethod

`func (o *UpdateConsentRequest) GetConsentMethod() string`

GetConsentMethod returns the ConsentMethod field if non-nil, zero value otherwise.

### GetConsentMethodOk

`func (o *UpdateConsentRequest) GetConsentMethodOk() (*string, bool)`

GetConsentMethodOk returns a tuple with the ConsentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentMethod

`func (o *UpdateConsentRequest) SetConsentMethod(v string)`

SetConsentMethod sets ConsentMethod field to given value.

### HasConsentMethod

`func (o *UpdateConsentRequest) HasConsentMethod() bool`

HasConsentMethod returns a boolean if a field has been set.

### GetConsentVersion

`func (o *UpdateConsentRequest) GetConsentVersion() string`

GetConsentVersion returns the ConsentVersion field if non-nil, zero value otherwise.

### GetConsentVersionOk

`func (o *UpdateConsentRequest) GetConsentVersionOk() (*string, bool)`

GetConsentVersionOk returns a tuple with the ConsentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentVersion

`func (o *UpdateConsentRequest) SetConsentVersion(v string)`

SetConsentVersion sets ConsentVersion field to given value.

### HasConsentVersion

`func (o *UpdateConsentRequest) HasConsentVersion() bool`

HasConsentVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdateConsentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateConsentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateConsentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateConsentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


