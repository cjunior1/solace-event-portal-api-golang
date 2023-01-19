# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConsumer**](ConsumersApi.md#CreateConsumer) | **Post** /api/v2/architecture/consumers | Creates a consumer object
[**DeleteConsumer**](ConsumersApi.md#DeleteConsumer) | **Delete** /api/v2/architecture/consumers/{id} | Deletes a consumer object
[**GetConsumer**](ConsumersApi.md#GetConsumer) | **Get** /api/v2/architecture/consumers/{id} | Retrieves a consumer object
[**GetConsumers**](ConsumersApi.md#GetConsumers) | **Get** /api/v2/architecture/consumers | Gets the consumer objects
[**UpdateConsumer**](ConsumersApi.md#UpdateConsumer) | **Patch** /api/v2/architecture/consumers/{id} | Update a consumer object

# **CreateConsumer**
> ConsumerResponse CreateConsumer(ctx, body)
Creates a consumer object

Use this API to create a consumer. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Consumer**](Consumer.md)| The consumer object. | 

### Return type

[**ConsumerResponse**](ConsumerResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteConsumer**
> DeleteConsumer(ctx, id)
Deletes a consumer object

Use this API to delete a consumer.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the consumer | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumer**
> ConsumerResponse GetConsumer(ctx, id)
Retrieves a consumer object

Use this API to retrieve a single consumer by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the consumer object. | 

### Return type

[**ConsumerResponse**](ConsumerResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConsumers**
> ConsumersResponse GetConsumers(ctx, optional)
Gets the consumer objects

Use this API to retrieve a list of consumers that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConsumersApiGetConsumersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConsumersApiGetConsumersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of consumers to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **applicationVersionIds** | [**optional.Interface of []string**](string.md)| Match only consumers with the given application version IDs, separated by commas. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only consumers with the given IDs separated by commas. | 

### Return type

[**ConsumersResponse**](ConsumersResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateConsumer**
> ConsumerResponse UpdateConsumer(ctx, body, id)
Update a consumer object

Use this API to update a consumer. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Consumer**](Consumer.md)| The consumer object. | 
  **id** | **string**| The ID of the consumer | 

### Return type

[**ConsumerResponse**](ConsumerResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

