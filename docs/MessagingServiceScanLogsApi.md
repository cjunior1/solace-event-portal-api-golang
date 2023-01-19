# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMessagingServiceScansLogs**](MessagingServiceScanLogsApi.md#GetMessagingServiceScansLogs) | **Get** /api/v2/architecture/messagingServiceScans/{scanId}/logs | (Beta) Retrieves a list of messaging service scan logs

# **GetMessagingServiceScansLogs**
> MessagingServiceScanLogListResponse GetMessagingServiceScansLogs(ctx, scanId, optional)
(Beta) Retrieves a list of messaging service scan logs

Use this API to retrieve a list of messaging service scan logs that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scanId** | **string**| The ID of the messaging service scan we want logs for. | 
 **optional** | ***MessagingServiceScanLogsApiGetMessagingServiceScansLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingServiceScanLogsApiGetMessagingServiceScansLogsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of messaging service scan logs to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 

### Return type

[**MessagingServiceScanLogListResponse**](MessagingServiceScanLogListResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

