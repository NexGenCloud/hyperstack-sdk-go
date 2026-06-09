# ConsentEventsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consent** | Pointer to [**UserConsent**](UserConsent.md) |  | [optional] 
**Events** | Pointer to [**[]UserConsentEvent**](UserConsentEvent.md) |  | [optional] 

## Methods

### NewConsentEventsResponse

`func NewConsentEventsResponse() *ConsentEventsResponse`

NewConsentEventsResponse instantiates a new ConsentEventsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentEventsResponseWithDefaults

`func NewConsentEventsResponseWithDefaults() *ConsentEventsResponse`

NewConsentEventsResponseWithDefaults instantiates a new ConsentEventsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsent

`func (o *ConsentEventsResponse) GetConsent() UserConsent`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *ConsentEventsResponse) GetConsentOk() (*UserConsent, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *ConsentEventsResponse) SetConsent(v UserConsent)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *ConsentEventsResponse) HasConsent() bool`

HasConsent returns a boolean if a field has been set.

### GetEvents

`func (o *ConsentEventsResponse) GetEvents() []UserConsentEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ConsentEventsResponse) GetEventsOk() (*[]UserConsentEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ConsentEventsResponse) SetEvents(v []UserConsentEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ConsentEventsResponse) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


