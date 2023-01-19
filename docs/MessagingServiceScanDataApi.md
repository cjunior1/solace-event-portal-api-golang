# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMessagingServiceScanData**](MessagingServiceScanDataApi.md#GetMessagingServiceScanData) | **Get** /api/v2/architecture/messagingServiceScans/{scanId}/dataCollection/{id} | (Beta) Retrieves a messaging service scan data object
[**GetMessagingServiceScansData**](MessagingServiceScanDataApi.md#GetMessagingServiceScansData) | **Get** /api/v2/architecture/messagingServiceScans/{scanId}/dataCollection | (Beta) Retrieves a list of messaging service scan data

# **GetMessagingServiceScanData**
> MessagingServiceScanDataResponse GetMessagingServiceScanData(ctx, scanId, id)
(Beta) Retrieves a messaging service scan data object

Use this API to retrieve a single messaging service scan data by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scanId** | **string**| The ID of the messaging service scan object. | 
  **id** | **string**| The ID of the messaging service scan data object. | 

### Return type

[**MessagingServiceScanDataResponse**](MessagingServiceScanDataResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetMessagingServiceScansData**
> MessagingServiceScanDataListResponse GetMessagingServiceScansData(ctx, scanId, optional)
(Beta) Retrieves a list of messaging service scan data

Use this API to retrieve a list of messaging service scan data that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **scanId** | **string**| The ID of the messaging service scan we want data for. | 
 **optional** | ***MessagingServiceScanDataApiGetMessagingServiceScansDataOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MessagingServiceScanDataApiGetMessagingServiceScansDataOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **optional.Int32**| The number of messaging service scan data to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **sort** | **optional.String**| The name of the field to sort on. | 
 **ids** | [**optional.Interface of []string**](string.md)| The IDs of the messaging service scan data. | 
 **collectionTypes** | [**optional.Interface of []string**](string.md)| Match only scan data whose dataCollectionType matches the given list. | 

### Return type

[**MessagingServiceScanDataListResponse**](MessagingServiceScanDataListResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

