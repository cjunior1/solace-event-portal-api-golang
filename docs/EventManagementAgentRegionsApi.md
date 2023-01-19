# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEventManagementAgentRegions**](EventManagementAgentRegionsApi.md#GetEventManagementAgentRegions) | **Get** /api/v2/architecture/eventManagementAgentRegions | (Beta) Retrieves a list of event management agent regions.

# **GetEventManagementAgentRegions**
> EventManagementAgentRegionsResponse GetEventManagementAgentRegions(ctx, optional)
(Beta) Retrieves a list of event management agent regions.

Use this API to retrieve a list of event management agent regions.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventManagementAgentRegionsApiGetEventManagementAgentRegionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventManagementAgentRegionsApiGetEventManagementAgentRegionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of event management agent regions to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]

### Return type

[**EventManagementAgentRegionsResponse**](EventManagementAgentRegionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

