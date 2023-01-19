# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEventManagementAgent**](EventManagementAgentsApi.md#CreateEventManagementAgent) | **Post** /api/v2/architecture/eventManagementAgents | (Beta) Creates an EMA object
[**DeleteEventManagementAgent**](EventManagementAgentsApi.md#DeleteEventManagementAgent) | **Delete** /api/v2/architecture/eventManagementAgents/{id} | (Beta) Deletes an EMA object
[**GetEventManagementAgent**](EventManagementAgentsApi.md#GetEventManagementAgent) | **Get** /api/v2/architecture/eventManagementAgents/{id} | (Beta) Retrieves an EMA object
[**GetEventManagementAgentConfigFile**](EventManagementAgentsApi.md#GetEventManagementAgentConfigFile) | **Get** /api/v2/architecture/eventManagementAgents/{id}/configuration/file | (Beta) Retrieves the raw configs in file format for an EMA object
[**GetEventManagementAgentConfigRaw**](EventManagementAgentsApi.md#GetEventManagementAgentConfigRaw) | **Get** /api/v2/architecture/eventManagementAgents/{id}/configuration/raw | (Beta) Retrieves the raw configs in string format for an EMA object
[**GetEventManagementAgents**](EventManagementAgentsApi.md#GetEventManagementAgents) | **Get** /api/v2/architecture/eventManagementAgents | (Beta) Retrieves a list of EMAs
[**UpdateEventManagementAgent**](EventManagementAgentsApi.md#UpdateEventManagementAgent) | **Patch** /api/v2/architecture/eventManagementAgents/{id} | (Beta) Updates an EMA object

# **CreateEventManagementAgent**
> EventManagementAgentResponse CreateEventManagementAgent(ctx, body)
(Beta) Creates an EMA object

Use this API to create an EMA object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventManagementAgent**](EventManagementAgent.md)| The EMA object. | 

### Return type

[**EventManagementAgentResponse**](EventManagementAgentResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventManagementAgent**
> DeleteEventManagementAgent(ctx, id)
(Beta) Deletes an EMA object

Use this API to delete an EMA.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the EMA object. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventManagementAgent**
> EventManagementAgentResponse GetEventManagementAgent(ctx, id, optional)
(Beta) Retrieves an EMA object

Use this API to retrieve a single EMA by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the EMA object. | 
 **optional** | ***EventManagementAgentsApiGetEventManagementAgentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventManagementAgentsApiGetEventManagementAgentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**| Specify extra data to be included, options are: referencedByMessagingServiceIds | 

### Return type

[**EventManagementAgentResponse**](EventManagementAgentResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventManagementAgentConfigFile**
> string GetEventManagementAgentConfigFile(ctx, id)
(Beta) Retrieves the raw configs in file format for an EMA object

Use this API to retrieve the raw configs for a single EMA by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the EMA object. | 

### Return type

**string**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/octet-stream, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventManagementAgentConfigRaw**
> string GetEventManagementAgentConfigRaw(ctx, id)
(Beta) Retrieves the raw configs in string format for an EMA object

Use this API to retrieve the raw configs for a single EMA by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the EMA object. | 

### Return type

**string**

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventManagementAgents**
> EventManagementAgentsResponse GetEventManagementAgents(ctx, optional)
(Beta) Retrieves a list of EMAs

Use this API to retrieve a list of EMAs that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventManagementAgentsApiGetEventManagementAgentsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventManagementAgentsApiGetEventManagementAgentsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of EMAs to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 
 **ids** | [**optional.Interface of []string**](string.md)| The IDs of the EMAs. | 
 **createdBy** | **optional.String**| Match only EMAs created by this user | 
 **eventManagementAgentRegionId** | **optional.String**| Match only EMAs in the given EMA-Region | 
 **include** | **optional.String**| Specify extra data to be included, options are: referencedByMessagingServiceIds | 

### Return type

[**EventManagementAgentsResponse**](EventManagementAgentsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventManagementAgent**
> EventManagementAgentResponse UpdateEventManagementAgent(ctx, body, id)
(Beta) Updates an EMA object

Use this API to update an EMA. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventManagementAgent**](EventManagementAgent.md)| The EMA object. | 
  **id** | **string**| The ID of the EMA object to update. | 

### Return type

[**EventManagementAgentResponse**](EventManagementAgentResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

