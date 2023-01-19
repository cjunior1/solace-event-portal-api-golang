# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEnvironment**](EnvironmentsApi.md#GetEnvironment) | **Get** /api/v2/architecture/environments/{id} | Retrieves an environment object
[**GetEnvironments**](EnvironmentsApi.md#GetEnvironments) | **Get** /api/v2/architecture/environments | Gets the environment objects

# **GetEnvironment**
> EnvironmentResponse GetEnvironment(ctx, id)
Retrieves an environment object

Use this API to retrieve a single environment by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the environment object. | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnvironments**
> EnvironmentsResponse GetEnvironments(ctx, optional)
Gets the environment objects

Use this API to list all environments.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnvironmentsApiGetEnvironmentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnvironmentsApiGetEnvironmentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of events to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **sort** | **optional.String**|  | 
 **like** | **optional.String**|  | 

### Return type

[**EnvironmentsResponse**](EnvironmentsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

