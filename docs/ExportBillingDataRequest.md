# ExportBillingDataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndDate** | **time.Time** |  | 
**OrgId** | Pointer to **int32** |  | [optional] 
**RequiredAttributes** | **[]string** |  | 
**RequiredMetrics** | **[]string** |  | 
**ResourceType** | **string** |  | 
**StartDate** | **time.Time** |  | 

## Methods

### NewExportBillingDataRequest

`func NewExportBillingDataRequest(endDate time.Time, requiredAttributes []string, requiredMetrics []string, resourceType string, startDate time.Time, ) *ExportBillingDataRequest`

NewExportBillingDataRequest instantiates a new ExportBillingDataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportBillingDataRequestWithDefaults

`func NewExportBillingDataRequestWithDefaults() *ExportBillingDataRequest`

NewExportBillingDataRequestWithDefaults instantiates a new ExportBillingDataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndDate

`func (o *ExportBillingDataRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ExportBillingDataRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ExportBillingDataRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.


### GetOrgId

`func (o *ExportBillingDataRequest) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ExportBillingDataRequest) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ExportBillingDataRequest) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *ExportBillingDataRequest) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetRequiredAttributes

`func (o *ExportBillingDataRequest) GetRequiredAttributes() []string`

GetRequiredAttributes returns the RequiredAttributes field if non-nil, zero value otherwise.

### GetRequiredAttributesOk

`func (o *ExportBillingDataRequest) GetRequiredAttributesOk() (*[]string, bool)`

GetRequiredAttributesOk returns a tuple with the RequiredAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAttributes

`func (o *ExportBillingDataRequest) SetRequiredAttributes(v []string)`

SetRequiredAttributes sets RequiredAttributes field to given value.


### GetRequiredMetrics

`func (o *ExportBillingDataRequest) GetRequiredMetrics() []string`

GetRequiredMetrics returns the RequiredMetrics field if non-nil, zero value otherwise.

### GetRequiredMetricsOk

`func (o *ExportBillingDataRequest) GetRequiredMetricsOk() (*[]string, bool)`

GetRequiredMetricsOk returns a tuple with the RequiredMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredMetrics

`func (o *ExportBillingDataRequest) SetRequiredMetrics(v []string)`

SetRequiredMetrics sets RequiredMetrics field to given value.


### GetResourceType

`func (o *ExportBillingDataRequest) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *ExportBillingDataRequest) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *ExportBillingDataRequest) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.


### GetStartDate

`func (o *ExportBillingDataRequest) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ExportBillingDataRequest) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ExportBillingDataRequest) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


