# MetricsFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**MetricItemFields**](MetricItemFields.md) |  | [optional] 
**DiskRead** | Pointer to [**MetricItemFields**](MetricItemFields.md) |  | [optional] 
**DiskWrite** | Pointer to [**MetricItemFields**](MetricItemFields.md) |  | [optional] 
**MemoryUsages** | Pointer to [**MetricItemFields**](MetricItemFields.md) |  | [optional] 
**NetworkIn** | Pointer to [**MetricItemFields**](MetricItemFields.md) |  | [optional] 
**NetworkOut** | Pointer to [**MetricItemFields**](MetricItemFields.md) |  | [optional] 

## Methods

### NewMetricsFields

`func NewMetricsFields() *MetricsFields`

NewMetricsFields instantiates a new MetricsFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsFieldsWithDefaults

`func NewMetricsFieldsWithDefaults() *MetricsFields`

NewMetricsFieldsWithDefaults instantiates a new MetricsFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *MetricsFields) GetCpu() MetricItemFields`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *MetricsFields) GetCpuOk() (*MetricItemFields, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *MetricsFields) SetCpu(v MetricItemFields)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *MetricsFields) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDiskRead

`func (o *MetricsFields) GetDiskRead() MetricItemFields`

GetDiskRead returns the DiskRead field if non-nil, zero value otherwise.

### GetDiskReadOk

`func (o *MetricsFields) GetDiskReadOk() (*MetricItemFields, bool)`

GetDiskReadOk returns a tuple with the DiskRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskRead

`func (o *MetricsFields) SetDiskRead(v MetricItemFields)`

SetDiskRead sets DiskRead field to given value.

### HasDiskRead

`func (o *MetricsFields) HasDiskRead() bool`

HasDiskRead returns a boolean if a field has been set.

### GetDiskWrite

`func (o *MetricsFields) GetDiskWrite() MetricItemFields`

GetDiskWrite returns the DiskWrite field if non-nil, zero value otherwise.

### GetDiskWriteOk

`func (o *MetricsFields) GetDiskWriteOk() (*MetricItemFields, bool)`

GetDiskWriteOk returns a tuple with the DiskWrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWrite

`func (o *MetricsFields) SetDiskWrite(v MetricItemFields)`

SetDiskWrite sets DiskWrite field to given value.

### HasDiskWrite

`func (o *MetricsFields) HasDiskWrite() bool`

HasDiskWrite returns a boolean if a field has been set.

### GetMemoryUsages

`func (o *MetricsFields) GetMemoryUsages() MetricItemFields`

GetMemoryUsages returns the MemoryUsages field if non-nil, zero value otherwise.

### GetMemoryUsagesOk

`func (o *MetricsFields) GetMemoryUsagesOk() (*MetricItemFields, bool)`

GetMemoryUsagesOk returns a tuple with the MemoryUsages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsages

`func (o *MetricsFields) SetMemoryUsages(v MetricItemFields)`

SetMemoryUsages sets MemoryUsages field to given value.

### HasMemoryUsages

`func (o *MetricsFields) HasMemoryUsages() bool`

HasMemoryUsages returns a boolean if a field has been set.

### GetNetworkIn

`func (o *MetricsFields) GetNetworkIn() MetricItemFields`

GetNetworkIn returns the NetworkIn field if non-nil, zero value otherwise.

### GetNetworkInOk

`func (o *MetricsFields) GetNetworkInOk() (*MetricItemFields, bool)`

GetNetworkInOk returns a tuple with the NetworkIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIn

`func (o *MetricsFields) SetNetworkIn(v MetricItemFields)`

SetNetworkIn sets NetworkIn field to given value.

### HasNetworkIn

`func (o *MetricsFields) HasNetworkIn() bool`

HasNetworkIn returns a boolean if a field has been set.

### GetNetworkOut

`func (o *MetricsFields) GetNetworkOut() MetricItemFields`

GetNetworkOut returns the NetworkOut field if non-nil, zero value otherwise.

### GetNetworkOutOk

`func (o *MetricsFields) GetNetworkOutOk() (*MetricItemFields, bool)`

GetNetworkOutOk returns a tuple with the NetworkOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOut

`func (o *MetricsFields) SetNetworkOut(v MetricItemFields)`

SetNetworkOut sets NetworkOut field to given value.

### HasNetworkOut

`func (o *MetricsFields) HasNetworkOut() bool`

HasNetworkOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


