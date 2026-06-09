# RecordConsentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**ConsentMethod** | Pointer to **string** |  | [optional] [default to "web_checkbox"]
**ConsentType** | **string** |  | 
**ConsentVersion** | Pointer to **string** | Version identifier. Defaults to the current version for the consent type | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Consent-type-specific metadata | [optional] 

## Methods

### NewRecordConsentRequest

`func NewRecordConsentRequest(action string, consentType string, ) *RecordConsentRequest`

NewRecordConsentRequest instantiates a new RecordConsentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordConsentRequestWithDefaults

`func NewRecordConsentRequestWithDefaults() *RecordConsentRequest`

NewRecordConsentRequestWithDefaults instantiates a new RecordConsentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *RecordConsentRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *RecordConsentRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *RecordConsentRequest) SetAction(v string)`

SetAction sets Action field to given value.


### GetConsentMethod

`func (o *RecordConsentRequest) GetConsentMethod() string`

GetConsentMethod returns the ConsentMethod field if non-nil, zero value otherwise.

### GetConsentMethodOk

`func (o *RecordConsentRequest) GetConsentMethodOk() (*string, bool)`

GetConsentMethodOk returns a tuple with the ConsentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentMethod

`func (o *RecordConsentRequest) SetConsentMethod(v string)`

SetConsentMethod sets ConsentMethod field to given value.

### HasConsentMethod

`func (o *RecordConsentRequest) HasConsentMethod() bool`

HasConsentMethod returns a boolean if a field has been set.

### GetConsentType

`func (o *RecordConsentRequest) GetConsentType() string`

GetConsentType returns the ConsentType field if non-nil, zero value otherwise.

### GetConsentTypeOk

`func (o *RecordConsentRequest) GetConsentTypeOk() (*string, bool)`

GetConsentTypeOk returns a tuple with the ConsentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentType

`func (o *RecordConsentRequest) SetConsentType(v string)`

SetConsentType sets ConsentType field to given value.


### GetConsentVersion

`func (o *RecordConsentRequest) GetConsentVersion() string`

GetConsentVersion returns the ConsentVersion field if non-nil, zero value otherwise.

### GetConsentVersionOk

`func (o *RecordConsentRequest) GetConsentVersionOk() (*string, bool)`

GetConsentVersionOk returns a tuple with the ConsentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsentVersion

`func (o *RecordConsentRequest) SetConsentVersion(v string)`

SetConsentVersion sets ConsentVersion field to given value.

### HasConsentVersion

`func (o *RecordConsentRequest) HasConsentVersion() bool`

HasConsentVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *RecordConsentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RecordConsentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RecordConsentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RecordConsentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


