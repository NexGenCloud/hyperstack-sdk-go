# Voucher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Voucher code | 
**Id** | **int32** | Voucher ID | 
**MaxRedemptionCount** | Pointer to **int32** | Max redemption count for a General Voucher | [optional] 
**RedemptionCount** | Pointer to **int32** | Current redemption count for a General Voucher | [optional] 
**RemainingRedemptions** | Pointer to **int32** | Remaining redemptions | [optional] 
**Status** | **string** | Voucher status | 

## Methods

### NewVoucher

`func NewVoucher(code string, id int32, status string, ) *Voucher`

NewVoucher instantiates a new Voucher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherWithDefaults

`func NewVoucherWithDefaults() *Voucher`

NewVoucherWithDefaults instantiates a new Voucher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *Voucher) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Voucher) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Voucher) SetCode(v string)`

SetCode sets Code field to given value.


### GetId

`func (o *Voucher) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Voucher) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Voucher) SetId(v int32)`

SetId sets Id field to given value.


### GetMaxRedemptionCount

`func (o *Voucher) GetMaxRedemptionCount() int32`

GetMaxRedemptionCount returns the MaxRedemptionCount field if non-nil, zero value otherwise.

### GetMaxRedemptionCountOk

`func (o *Voucher) GetMaxRedemptionCountOk() (*int32, bool)`

GetMaxRedemptionCountOk returns a tuple with the MaxRedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRedemptionCount

`func (o *Voucher) SetMaxRedemptionCount(v int32)`

SetMaxRedemptionCount sets MaxRedemptionCount field to given value.

### HasMaxRedemptionCount

`func (o *Voucher) HasMaxRedemptionCount() bool`

HasMaxRedemptionCount returns a boolean if a field has been set.

### GetRedemptionCount

`func (o *Voucher) GetRedemptionCount() int32`

GetRedemptionCount returns the RedemptionCount field if non-nil, zero value otherwise.

### GetRedemptionCountOk

`func (o *Voucher) GetRedemptionCountOk() (*int32, bool)`

GetRedemptionCountOk returns a tuple with the RedemptionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionCount

`func (o *Voucher) SetRedemptionCount(v int32)`

SetRedemptionCount sets RedemptionCount field to given value.

### HasRedemptionCount

`func (o *Voucher) HasRedemptionCount() bool`

HasRedemptionCount returns a boolean if a field has been set.

### GetRemainingRedemptions

`func (o *Voucher) GetRemainingRedemptions() int32`

GetRemainingRedemptions returns the RemainingRedemptions field if non-nil, zero value otherwise.

### GetRemainingRedemptionsOk

`func (o *Voucher) GetRemainingRedemptionsOk() (*int32, bool)`

GetRemainingRedemptionsOk returns a tuple with the RemainingRedemptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingRedemptions

`func (o *Voucher) SetRemainingRedemptions(v int32)`

SetRemainingRedemptions sets RemainingRedemptions field to given value.

### HasRemainingRedemptions

`func (o *Voucher) HasRemainingRedemptions() bool`

HasRemainingRedemptions returns a boolean if a field has been set.

### GetStatus

`func (o *Voucher) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Voucher) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Voucher) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


