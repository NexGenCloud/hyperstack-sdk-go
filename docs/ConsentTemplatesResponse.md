# ConsentTemplatesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Templates** | Pointer to **map[string]interface{}** | Templates keyed by consent type | [optional] 

## Methods

### NewConsentTemplatesResponse

`func NewConsentTemplatesResponse() *ConsentTemplatesResponse`

NewConsentTemplatesResponse instantiates a new ConsentTemplatesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsentTemplatesResponseWithDefaults

`func NewConsentTemplatesResponseWithDefaults() *ConsentTemplatesResponse`

NewConsentTemplatesResponseWithDefaults instantiates a new ConsentTemplatesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplates

`func (o *ConsentTemplatesResponse) GetTemplates() map[string]interface{}`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *ConsentTemplatesResponse) GetTemplatesOk() (*map[string]interface{}, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *ConsentTemplatesResponse) SetTemplates(v map[string]interface{})`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *ConsentTemplatesResponse) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


