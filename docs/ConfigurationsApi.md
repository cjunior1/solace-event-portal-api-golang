# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfiguration**](ConfigurationsApi.md#GetConfiguration) | **Get** /api/v2/architecture/configurations/{id} | (Beta) Retrieves a configuration object
[**GetConfigurations**](ConfigurationsApi.md#GetConfigurations) | **Get** /api/v2/architecture/configurations | (Beta) Gets the configuration objects

# **GetConfiguration**
> ConfigurationResponse GetConfiguration(ctx, id)
(Beta) Retrieves a configuration object

Use this API to retrieve a single configuration by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the configuration object. | 

### Return type

[**ConfigurationResponse**](ConfigurationResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetConfigurations**
> ConfigurationsResponse GetConfigurations(ctx, optional)
(Beta) Gets the configuration objects

Use this API to retrieve a list of configurations that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigurationsApiGetConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigurationsApiGetConfigurationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of configurations to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **messagingServiceIds** | [**optional.Interface of []string**](string.md)| Match only configurations with the given messaging service IDs separated by commas. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only configurations with the given IDs separated by commas. | 
 **configurationTypeIds** | [**optional.Interface of []string**](string.md)| Match only configurations with the given configuration type IDs separated by commas.&lt;br&gt;Refer &lt;a href&#x3D;\&quot;#/Configuration%20Types/getConfigurationTypes\&quot;&gt;here&lt;/a&gt; for details on configuration types. | 
 **entityTypes** | [**optional.Interface of []string**](string.md)| Match only configurations with the given entity type values separated by commas. | 
 **entityIds** | [**optional.Interface of []string**](string.md)| Match only configurations with the given entity IDs separated by commas. | 
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 

### Return type

[**ConfigurationsResponse**](ConfigurationsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

