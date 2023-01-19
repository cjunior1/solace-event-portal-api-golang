# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnum**](EnumsApi.md#CreateEnum) | **Post** /api/v2/architecture/enums | Creates an enum object
[**CreateEnumVersion**](EnumsApi.md#CreateEnumVersion) | **Post** /api/v2/architecture/enumVersions | Creates an enum version object
[**CreateEnumVersionForEnum**](EnumsApi.md#CreateEnumVersionForEnum) | **Post** /api/v2/architecture/enums/{enumId}/versions | Creates an enum version object
[**DeleteEnum**](EnumsApi.md#DeleteEnum) | **Delete** /api/v2/architecture/enums/{id} | Deletes an enum
[**DeleteEnumVersion**](EnumsApi.md#DeleteEnumVersion) | **Delete** /api/v2/architecture/enumVersions/{id} | Deletes an enum version
[**DeleteEnumVersionForEnum**](EnumsApi.md#DeleteEnumVersionForEnum) | **Delete** /api/v2/architecture/enums/{enumId}/versions/{id} | Deletes an enum version
[**GetEnum**](EnumsApi.md#GetEnum) | **Get** /api/v2/architecture/enums/{id} | Retrieves an enum object
[**GetEnumVersion**](EnumsApi.md#GetEnumVersion) | **Get** /api/v2/architecture/enumVersions/{versionId} | Retrieves an enumeration version object
[**GetEnumVersionForEnum**](EnumsApi.md#GetEnumVersionForEnum) | **Get** /api/v2/architecture/enums/{enumId}/versions/{id} | Retrieves an enum version object
[**GetEnumVersions**](EnumsApi.md#GetEnumVersions) | **Get** /api/v2/architecture/enumVersions | Gets the enumeration version objects
[**GetEnumVersionsForEnum**](EnumsApi.md#GetEnumVersionsForEnum) | **Get** /api/v2/architecture/enums/{enumId}/versions | Lists enums
[**GetEnums**](EnumsApi.md#GetEnums) | **Get** /api/v2/architecture/enums | Lists enums
[**UpdateEnum**](EnumsApi.md#UpdateEnum) | **Patch** /api/v2/architecture/enums/{id} | Updates an enum object
[**UpdateEnumVersion**](EnumsApi.md#UpdateEnumVersion) | **Patch** /api/v2/architecture/enumVersions/{id} | Updates an enum version object
[**UpdateEnumVersionForEnum**](EnumsApi.md#UpdateEnumVersionForEnum) | **Patch** /api/v2/architecture/enums/{enumId}/versions/{id} | Updates an enum version object
[**UpdateEnumVersionState**](EnumsApi.md#UpdateEnumVersionState) | **Patch** /api/v2/architecture/enumVersions/{id}/state | Updates the state of an enum version object
[**UpdateEnumVersionStateForEnum**](EnumsApi.md#UpdateEnumVersionStateForEnum) | **Patch** /api/v2/architecture/enums/{enumId}/versions/{id}/state | Updates the state of an enum version object

# **CreateEnum**
> TopicAddressEnumResponse CreateEnum(ctx, body)
Creates an enum object

description

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicAddressEnum**](TopicAddressEnum.md)| Enum object description. | 

### Return type

[**TopicAddressEnumResponse**](TopicAddressEnumResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEnumVersion**
> TopicAddressEnumVersionResponse CreateEnumVersion(ctx, body)
Creates an enum version object

description

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicAddressEnumVersion**](TopicAddressEnumVersion.md)| Enum object description with its values. | 

### Return type

[**TopicAddressEnumVersionResponse**](TopicAddressEnumVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEnumVersionForEnum**
> TopicAddressEnumVersionResponse CreateEnumVersionForEnum(ctx, body, enumId)
Creates an enum version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Enums/createEnumVersion>another endpoint.</a><br><br>*description

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicAddressEnumVersion**](TopicAddressEnumVersion.md)| Enum object description with its values. | 
  **enumId** | **string**|  | 

### Return type

[**TopicAddressEnumVersionResponse**](TopicAddressEnumVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnum**
> DeleteEnum(ctx, id)
Deletes an enum

Use this API to delete an enum. The enum must not have any versions or else it cannot be deleted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the enum. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnumVersion**
> DeleteEnumVersion(ctx, id)
Deletes an enum version

Use this API to delete an enum version. The version must not be in use by any events else it cannot be deleted. This also deletes the version's values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the enum version object. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEnumVersionForEnum**
> DeleteEnumVersionForEnum(ctx, enumId, id)
Deletes an enum version

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Enums/deleteEnumVersion>another endpoint.</a><br><br>*Use this API to delete an enum version. The version must not be in use by any events else it cannot be deleted. This also deletes the version's values.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enumId** | **string**| The ID of the enum object. | 
  **id** | **string**| The ID of the enum version object. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnum**
> TopicAddressEnumResponse GetEnum(ctx, id)
Retrieves an enum object

Use this API to retrieve a single enum by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the enum object. | 

### Return type

[**TopicAddressEnumResponse**](TopicAddressEnumResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnumVersion**
> TopicAddressEnumVersionResponse GetEnumVersion(ctx, versionId)
Retrieves an enumeration version object

Use this API to retrieve a single enumeration version by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **versionId** | **string**| The ID of the enumeration version object. | 

### Return type

[**TopicAddressEnumVersionResponse**](TopicAddressEnumVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnumVersionForEnum**
> TopicAddressEnumVersionResponse GetEnumVersionForEnum(ctx, enumId, id)
Retrieves an enum version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Enums/getEnumVersion>another endpoint.</a><br><br>*Use this API to retrieve a single enum version by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enumId** | **string**| The ID of the enum object. | 
  **id** | **string**| The ID of the enum version object. | 

### Return type

[**TopicAddressEnumVersionResponse**](TopicAddressEnumVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnumVersions**
> TopicAddressEnumVersionsResponse GetEnumVersions(ctx, optional)
Gets the enumeration version objects

Use this API to retrieve a list of enumeration versions that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnumsApiGetEnumVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnumsApiGetEnumVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of enumerations to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **enumIds** | [**optional.Interface of []string**](string.md)| Match only enumeration versions of these enum IDs, separated by commas. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only enumeration versions with the given IDs, separated by commas. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**TopicAddressEnumVersionsResponse**](TopicAddressEnumVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnumVersionsForEnum**
> TopicAddressEnumVersionsResponse GetEnumVersionsForEnum(ctx, enumId, optional)
Lists enums

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Enums/getEnumVersions>another endpoint.</a><br><br>*Use this API to list enum versions based on certain criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **enumId** | **string**| The ID of the enum object. | 
 **optional** | ***EnumsApiGetEnumVersionsForEnumOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnumsApiGetEnumVersionsForEnumOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **pageSize** | **optional.Int32**| The number of enum versions to get per page. | 
 **ids** | [**optional.Interface of []string**](string.md)| The ids of the enum versions. | 
 **versions** | [**optional.Interface of []string**](string.md)| The versions of the enum version. | 
 **displayName** | **optional.String**| The display name of the enum versions. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**TopicAddressEnumVersionsResponse**](TopicAddressEnumVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEnums**
> TopicAddressEnumsResponse GetEnums(ctx, optional)
Lists enums

Use this API to list enums based on certain criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EnumsApiGetEnumsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EnumsApiGetEnumsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of enums to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **ids** | [**optional.Interface of []string**](string.md)| The IDs of the enums. | 
 **applicationDomainId** | **optional.String**| The application domain ID of the enums. | 
 **applicationDomainIds** | [**optional.Interface of []string**](string.md)| Match only enums in the given application domain ids. | 
 **names** | [**optional.Interface of []string**](string.md)| The names of the enums. | 
 **shared** | **optional.Bool**| Match only with shared or unshared enums. | 
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**TopicAddressEnumsResponse**](TopicAddressEnumsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnum**
> TopicAddressEnumResponse UpdateEnum(ctx, body, id)
Updates an enum object

description

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicAddressEnum**](TopicAddressEnum.md)| Enum updates. | 
  **id** | **string**| The ID of the enum. | 

### Return type

[**TopicAddressEnumResponse**](TopicAddressEnumResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnumVersion**
> TopicAddressEnumVersionResponse UpdateEnumVersion(ctx, body, id)
Updates an enum version object

Use this API to update an enum version. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicAddressEnumVersion**](TopicAddressEnumVersion.md)| The enum version object. | 
  **id** | **string**| The ID of the enum version object to update. | 

### Return type

[**TopicAddressEnumVersionResponse**](TopicAddressEnumVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnumVersionForEnum**
> TopicAddressEnumVersionResponse UpdateEnumVersionForEnum(ctx, body, enumId, id)
Updates an enum version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Enums/updateEnumVersion>another endpoint.</a><br><br>*Use this API to update an enum version. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicAddressEnumVersion**](TopicAddressEnumVersion.md)| The enum version object. | 
  **enumId** | **string**| The ID of the parent enum object. | 
  **id** | **string**| The ID of the enum version object to update. | 

### Return type

[**TopicAddressEnumVersionResponse**](TopicAddressEnumVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnumVersionState**
> StateChangeRequestResponse UpdateEnumVersionState(ctx, body, id)
Updates the state of an enum version object

Use this API to update the state of an enum version. You only need to specify the target stateId field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)| The state object. | 
  **id** | **string**| The ID of the enum version object to update. | 

### Return type

[**StateChangeRequestResponse**](StateChangeRequestResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEnumVersionStateForEnum**
> VersionedObjectStateChangeRequest UpdateEnumVersionStateForEnum(ctx, body, enumId, id)
Updates the state of an enum version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Enums/updateEnumVersionState>another endpoint.</a><br><br>*Use this API to update the state of an enum version. You only need to specify the target stateId field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TopicAddressEnumVersion**](TopicAddressEnumVersion.md)| The enum version object. | 
  **enumId** | **string**| The ID of the parent enum object. | 
  **id** | **string**| The ID of the enum version object to update. | 

### Return type

[**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

