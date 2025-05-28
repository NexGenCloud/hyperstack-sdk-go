# GetAllDiscountsFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscountResources** | Pointer to [**[]DiscountResourceFields**](DiscountResourceFields.md) |  | [optional] 
**DiscountStatus** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**OrgId** | Pointer to **int32** |  | [optional] 
**OrgName** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetAllDiscountsFields

`func NewGetAllDiscountsFields() *GetAllDiscountsFields`

NewGetAllDiscountsFields instantiates a new GetAllDiscountsFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllDiscountsFieldsWithDefaults

`func NewGetAllDiscountsFieldsWithDefaults() *GetAllDiscountsFields`

NewGetAllDiscountsFieldsWithDefaults instantiates a new GetAllDiscountsFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscountResources

`func (o *GetAllDiscountsFields) GetDiscountResources() []DiscountResourceFields`

GetDiscountResources returns the DiscountResources field if non-nil, zero value otherwise.

### GetDiscountResourcesOk

`func (o *GetAllDiscountsFields) GetDiscountResourcesOk() (*[]DiscountResourceFields, bool)`

GetDiscountResourcesOk returns a tuple with the DiscountResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountResources

`func (o *GetAllDiscountsFields) SetDiscountResources(v []DiscountResourceFields)`

SetDiscountResources sets DiscountResources field to given value.

### HasDiscountResources

`func (o *GetAllDiscountsFields) HasDiscountResources() bool`

HasDiscountResources returns a boolean if a field has been set.

### GetDiscountStatus

`func (o *GetAllDiscountsFields) GetDiscountStatus() string`

GetDiscountStatus returns the DiscountStatus field if non-nil, zero value otherwise.

### GetDiscountStatusOk

`func (o *GetAllDiscountsFields) GetDiscountStatusOk() (*string, bool)`

GetDiscountStatusOk returns a tuple with the DiscountStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountStatus

`func (o *GetAllDiscountsFields) SetDiscountStatus(v string)`

SetDiscountStatus sets DiscountStatus field to given value.

### HasDiscountStatus

`func (o *GetAllDiscountsFields) HasDiscountStatus() bool`

HasDiscountStatus returns a boolean if a field has been set.

### GetEndDate

`func (o *GetAllDiscountsFields) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetAllDiscountsFields) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetAllDiscountsFields) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetAllDiscountsFields) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetOrgId

`func (o *GetAllDiscountsFields) GetOrgId() int32`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *GetAllDiscountsFields) GetOrgIdOk() (*int32, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *GetAllDiscountsFields) SetOrgId(v int32)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *GetAllDiscountsFields) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetOrgName

`func (o *GetAllDiscountsFields) GetOrgName() string`

GetOrgName returns the OrgName field if non-nil, zero value otherwise.

### GetOrgNameOk

`func (o *GetAllDiscountsFields) GetOrgNameOk() (*string, bool)`

GetOrgNameOk returns a tuple with the OrgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgName

`func (o *GetAllDiscountsFields) SetOrgName(v string)`

SetOrgName sets OrgName field to given value.

### HasOrgName

`func (o *GetAllDiscountsFields) HasOrgName() bool`

HasOrgName returns a boolean if a field has been set.

### GetStartDate

`func (o *GetAllDiscountsFields) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetAllDiscountsFields) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetAllDiscountsFields) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetAllDiscountsFields) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


