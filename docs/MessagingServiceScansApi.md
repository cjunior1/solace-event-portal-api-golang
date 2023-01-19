# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMessagingServiceScan**](MessagingServiceScansApi.md#DeleteMessagingServiceScan) | **Delete** /api/v2/architecture/messagingServiceScans/{id} | (Beta) Deletes a messaging service scan object
[**GetMessagingServiceScan**](MessagingServiceScansApi.md#GetMessagingServiceScan) | **Get** /api/v2/architecture/messagingServiceScans/{id} | (Beta) Retrieves a messaging service scan object
[**GetMessagingServiceScans**](MessagingServiceScansApi.md#GetMessagingServiceScans) | **Get** /api/v2/architecture/messagingServiceScans | (Beta) Retrieves a list of messaging service scans

# **DeleteMessagingServiceScan**
> DeleteMessagingServiceScan(ctx, id)
(Beta) Deletes a messaging service scan object

Use this API to delete a messaging service scan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the messaging service scan object. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagingServiceScan**
> MessagingServiceScanResponse GetMessagingServiceScan(ctx, id)
(Beta) Retrieves a messaging service scan object

Use this API to retrieve a single messaging service scan by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the messaging service scan object. | 

### Return type

[**MessagingServiceScanResponse**](MessagingServiceScanResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagingServiceScans**
> MessagingServiceScansResponse GetMessagingServiceScans(ctx, optional)
(Beta) Retrieves a list of messaging service scans

Use this API to retrieve a list of messaging service scans that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MessagingServiceScansApiGetMessagingServiceScansOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingServiceScansApiGetMessagingServiceScansOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of messaging service scans to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 
 **ids** | [**optional.Interface of []string**](string.md)| The IDs of the messaging service scans. | 
 **messagingServiceId** | **optional.String**| Match only messaging service scans in the given messagingService | 
 **eventMeshId** | **optional.String**| Match only messaging service scans in the given eventMeshId | 

### Return type

[**MessagingServiceScansResponse**](MessagingServiceScansResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

