# ExcludeBillingPostPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exclude** | **bool** | &#x60;true&#x60; excludes the resource from billing while &#x60;false&#x60; does not. | 
**ResourceId** | **int32** | The ID of the resource which is being excluded from billing. | 
**ResourceType** | **string** | The type of the resource which is being excluded from billing. | 

## Methods

### NewExcludeBillingPostPayload

`func NewExcludeBillingPostPayload(exclude bool, resourceId int32, resourceType string, ) *ExcludeBillingPostPayload`

NewExcludeBillingPostPayload instantiates a new ExcludeBillingPostPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExcludeBillingPostPayloadWithDefaults

`func NewExcludeBillingPostPayloadWithDefaults() *ExcludeBillingPostPayload`

NewExcludeBillingPostPayloadWithDefaults instantiates a new ExcludeBillingPostPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExclude

`func (o *ExcludeBillingPostPayload) GetExclude() bool`

GetExclude returns the Exclude field if non-nil, zero value otherwise.

### GetExcludeOk

`func (o *ExcludeBillingPostPayload) GetExcludeOk() (*bool, bool)`

GetExcludeOk returns a tuple with the Exclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExclude

`func (o *ExcludeBillingPostPayload) SetExclude(v bool)`

SetExclude sets Exclude field to given value.


### GetResourceId

`func (o *ExcludeBillingPostPayload) GetResourceId() int32`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ExcludeBillingPostPayload) GetResourceIdOk() (*int32, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ExcludeBillingPostPayload) SetResourceId(v int32)`

SetResourceId sets ResourceId field to given value.


### GetResourceType

`func (o *ExcludeBillingPostPayload) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ExcludeBillingPostPayload) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ExcludeBillingPostPayload) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


