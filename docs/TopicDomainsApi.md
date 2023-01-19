# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTopicDomain**](TopicDomainsApi.md#CreateTopicDomain) | **Post** /api/v2/architecture/topicDomains | Creates a topic domain object
[**DeleteTopicDomain**](TopicDomainsApi.md#DeleteTopicDomain) | **Delete** /api/v2/architecture/topicDomains/{id} | Deletes a topic domain object
[**GetTopicDomain**](TopicDomainsApi.md#GetTopicDomain) | **Get** /api/v2/architecture/topicDomains/{id} | Retrieves a topic domain object
[**GetTopicDomains**](TopicDomainsApi.md#GetTopicDomains) | **Get** /api/v2/architecture/topicDomains | Gets the topic domain objects

# **CreateTopicDomain**
> TopicDomainResponse CreateTopicDomain(ctx, body)
Creates a topic domain object

Topic Domains govern the format of topic addresses within an application domain

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicDomain**](TopicDomain.md)|  | 

### Return type

[**TopicDomainResponse**](TopicDomainResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTopicDomain**
> TopicDomainResponse DeleteTopicDomain(ctx, id)
Deletes a topic domain object

Use this API to delete a topic domain.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the topic domain | 

### Return type

[**TopicDomainResponse**](TopicDomainResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTopicDomain**
> TopicDomainResponse GetTopicDomain(ctx, id)
Retrieves a topic domain object

Use this API to retrieve a single topic domain by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the topic domain object. | 

### Return type

[**TopicDomainResponse**](TopicDomainResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTopicDomains**
> TopicDomainsResponse GetTopicDomains(ctx, optional)
Gets the topic domain objects

Use this API to retrieve a list of topic domains that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TopicDomainsApiGetTopicDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TopicDomainsApiGetTopicDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of topic domains to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **ids** | [**optional.Interface of []string**](string.md)| Match only topic domains with the given IDs separated by commas. | 
 **brokerType** | **optional.String**| Match only topic domains with the given brokerType. | 
 **applicationDomainIds** | [**optional.Interface of []string**](string.md)| Match only topic domains with the given application domain ids separated by commas. | 
 **applicationDomainId** | **optional.String**|  | 

### Return type

[**TopicDomainsResponse**](TopicDomainsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

