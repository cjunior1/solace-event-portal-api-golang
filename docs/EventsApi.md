# {{classname}}

All URIs are relative to *https://api.solace.cloud*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEvent**](EventsApi.md#CreateEvent) | **Post** /api/v2/architecture/events | Creates an event object
[**CreateEventVersion**](EventsApi.md#CreateEventVersion) | **Post** /api/v2/architecture/eventVersions | Creates an event version object
[**CreateEventVersionForEvent**](EventsApi.md#CreateEventVersionForEvent) | **Post** /api/v2/architecture/events/{eventId}/versions | Creates an event version object
[**DeleteEvent**](EventsApi.md#DeleteEvent) | **Delete** /api/v2/architecture/events/{id} | Deletes an event object
[**DeleteEventVersion**](EventsApi.md#DeleteEventVersion) | **Delete** /api/v2/architecture/eventVersions/{id} | Deletes an event version object
[**DeleteEventVersionForEvent**](EventsApi.md#DeleteEventVersionForEvent) | **Delete** /api/v2/architecture/events/{eventId}/versions/{id} | Deletes an event version object
[**GetEvent**](EventsApi.md#GetEvent) | **Get** /api/v2/architecture/events/{id} | Retrieves an event object
[**GetEventVersion**](EventsApi.md#GetEventVersion) | **Get** /api/v2/architecture/eventVersions/{id} | Retrieves an event version object
[**GetEventVersionForEvent**](EventsApi.md#GetEventVersionForEvent) | **Get** /api/v2/architecture/events/{eventId}/versions/{id} | Retrieves an event version object
[**GetEventVersions**](EventsApi.md#GetEventVersions) | **Get** /api/v2/architecture/eventVersions | Gets event version objects
[**GetEventVersionsForEvent**](EventsApi.md#GetEventVersionsForEvent) | **Get** /api/v2/architecture/events/{eventId}/versions | Gets the event version objects for an event
[**GetEvents**](EventsApi.md#GetEvents) | **Get** /api/v2/architecture/events | Gets the event objects
[**UpdateEvent**](EventsApi.md#UpdateEvent) | **Patch** /api/v2/architecture/events/{id} | Updates an event object
[**UpdateEventVersion**](EventsApi.md#UpdateEventVersion) | **Patch** /api/v2/architecture/eventVersions/{id} | Updates an event version object
[**UpdateEventVersionForEvent**](EventsApi.md#UpdateEventVersionForEvent) | **Patch** /api/v2/architecture/events/{eventId}/versions/{id} | Updates an event version object
[**UpdateEventVersionState**](EventsApi.md#UpdateEventVersionState) | **Patch** /api/v2/architecture/eventVersions/{id}/state | Updates the state of an event version object
[**UpdateEventVersionStateForEvent**](EventsApi.md#UpdateEventVersionStateForEvent) | **Patch** /api/v2/architecture/events/{eventId}/versions/{id}/state | Updates the state of an event version object
[**UpdateMsgSvcAssociationForEventVersion**](EventsApi.md#UpdateMsgSvcAssociationForEventVersion) | **Put** /api/v2/architecture/eventVersions/{id}/messagingServices | Update messaging service association for an event version object

# **CreateEvent**
> EventResponse CreateEvent(ctx, body)
Creates an event object

Events are the primary building block of an event-driven architecture. Applications publish and subscribe to events and events reference schemas.  In the Event Portal, an event is a type of event as opposed to a specific event instance.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Event**](Event.md)| The event requires a name and an application domain ID. | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventVersion**
> EventVersionResponse CreateEventVersion(ctx, body)
Creates an event version object

Creates an event version object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventVersion**](EventVersion.md)| App version request body description | 

### Return type

[**EventVersionResponse**](EventVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateEventVersionForEvent**
> EventVersionResponse CreateEventVersionForEvent(ctx, body, eventId)
Creates an event version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Events/createEventVersion>another endpoint.</a><br><br>*Creates an event version object

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventVersion**](EventVersion.md)| App version request body description | 
  **eventId** | **string**| The ID of the parent event | 

### Return type

[**EventVersionResponse**](EventVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEvent**
> DeleteEvent(ctx, id)
Deletes an event object

Use this API to delete an event.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event object. | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventVersion**
> DeleteEventVersion(ctx, id)
Deletes an event version object

Use this API to delete an event version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event version | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteEventVersionForEvent**
> DeleteEventVersionForEvent(ctx, eventId, id)
Deletes an event version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Events/deleteEventVersion>another endpoint.</a><br><br>*Use this API to delete an event version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | **string**| The ID of the parent event | 
  **id** | **string**| The ID of the event version | 

### Return type

 (empty response body)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvent**
> EventResponse GetEvent(ctx, id)
Retrieves an event object

Use this API to retrieve a single event by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event object. | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventVersion**
> EventVersionResponse GetEventVersion(ctx, id, optional)
Retrieves an event version object

Use this API to retrieve a single event version by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| The ID of the event version object. | 
 **optional** | ***EventsApiGetEventVersionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetEventVersionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **optional.String**|  | 

### Return type

[**EventVersionResponse**](EventVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventVersionForEvent**
> EventVersionResponse GetEventVersionForEvent(ctx, eventId, id)
Retrieves an event version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Events/getEventVersion>another endpoint.</a><br><br>*Use this API to retrieve a single event version by its ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | **string**| The ID of the parent event. | 
  **id** | **string**| The ID of the event version. | 

### Return type

[**EventVersionResponse**](EventVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventVersions**
> EventVersionsResponse GetEventVersions(ctx, optional)
Gets event version objects

Use this API to retrieve a list of event versions that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventsApiGetEventVersionsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetEventVersionsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of event to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **eventIds** | [**optional.Interface of []string**](string.md)| Match only event versions of these event IDs, separated by commas. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only event versions with the given IDs, separated by commas. | 
 **messagingServiceIds** | [**optional.Interface of []string**](string.md)| Match only event versions with the given messaging service IDs, separated by commas. | 
 **include** | **optional.String**|  | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 

### Return type

[**EventVersionsResponse**](EventVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEventVersionsForEvent**
> EventVersionsResponse GetEventVersionsForEvent(ctx, eventId, optional)
Gets the event version objects for an event

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Events/getEventVersions>another endpoint.</a><br><br>*Use this API to retrieve a list of event versions that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventId** | **string**| The ID of the parent event. | 
 **optional** | ***EventsApiGetEventVersionsForEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetEventVersionsForEventOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 
 **displayName** | **optional.String**| Match event versions with the given display name. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match event versions with the given IDs separated by commas. | 
 **version** | **optional.String**| Match event version objects with the given version. | 

### Return type

[**EventVersionsResponse**](EventVersionsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetEvents**
> EventsResponse GetEvents(ctx, optional)
Gets the event objects

Use this API to retrieve a list of events that match the given parameters.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventsApiGetEventsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiGetEventsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **optional.Int32**| The number of events to get per page. | [default to 20]
 **pageNumber** | **optional.Int32**| The page number to get. | [default to 1]
 **name** | **optional.String**| Name of the event to match on. | 
 **shared** | **optional.Bool**| Match only with shared or unshared events. | 
 **applicationDomainId** | **optional.String**| Match only events in the given application domain. | 
 **applicationDomainIds** | [**optional.Interface of []string**](string.md)| Match only events in the given application domain ids. | 
 **sort** | **optional.String**| Sort based on the provided parameters. &lt;br&gt; The value can either be a standalone field name (&#x60;?sort&#x3D;&lt;field&gt;&#x60;) or a field and direction, which must be delimited by a colon (&#x60;?sort&#x3D;&lt;field&gt;:&lt;asc|desc&gt;&#x60;). If the direction is not specified, the default is ascending. | 
 **customAttributes** | **optional.String**| Returns the entities that match the custom attribute filter.&lt;br&gt;To filter by custom attribute name and value, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x3D;&#x3D;&lt;custom-attribute-value&gt;&#x60;. &lt;br&gt;To filter by custom attribute name, use the format: &#x60;customAttributes&#x3D;&lt;custom-attribute-name&gt;&#x60;. &lt;br&gt;The filter supports the &#x60;AND&#x60; operator for multiple custom attribute definitions (not multiple values for a given definition). Use &#x60;;&#x60; (&#x60;semicolon&#x60;) to separate multiple queries with &#x60;AND&#x60; operation. &lt;br&gt;Note: the filter only supports custom attribute values containing characters in &#x60;[a-zA-Z0-9_\\-\\. ]&#x60;. | 
 **ids** | [**optional.Interface of []string**](string.md)| Match only events with the given IDs separated by commas. | 

### Return type

[**EventsResponse**](EventsResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEvent**
> EventResponse UpdateEvent(ctx, body, id)
Updates an event object

Use this API to update an event. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Event**](Event.md)| The event object. | 
  **id** | **string**| The ID of the event object to update. | 

### Return type

[**EventResponse**](EventResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventVersion**
> EventVersionResponse UpdateEventVersion(ctx, body, id)
Updates an event version object

Use this API to update an event version. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventVersion**](EventVersion.md)| The event version object. | 
  **id** | **string**| The ID of the event version object to update. | 

### Return type

[**EventVersionResponse**](EventVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventVersionForEvent**
> EventVersionResponse UpdateEventVersionForEvent(ctx, body, eventId, id)
Updates an event version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Events/updateEventVersion>another endpoint.</a><br><br>*Use this API to update an event version. You only need to specify the fields that need to be updated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventVersion**](EventVersion.md)| The event version object. | 
  **eventId** | **string**| The ID of the parent event object. | 
  **id** | **string**| The ID of the event version object to update. | 

### Return type

[**EventVersionResponse**](EventVersionResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventVersionState**
> StateChangeRequestResponse UpdateEventVersionState(ctx, body, id)
Updates the state of an event version object

Use this API to update the state of event version. You only need to specify the target stateId field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)| The state object. | 
  **id** | **string**| The ID of the event version object to update. | 

### Return type

[**StateChangeRequestResponse**](StateChangeRequestResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateEventVersionStateForEvent**
> VersionedObjectStateChangeRequest UpdateEventVersionStateForEvent(ctx, body, eventId, id)
Updates the state of an event version object

*deprecationDate: 2022-11-01T00:00:00.000Z<br>removalDate: 2023-02-01T00:00:00.000Z<br>reason: Replaced by <a href=#/Events/updateEventVersionState>another endpoint.</a><br><br>*Use this API to update the state of event version. You only need to specify the target stateId field

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EventVersion**](EventVersion.md)| The event version object. | 
  **eventId** | **string**| The ID of the parent event object. | 
  **id** | **string**| The ID of the event version object to update. | 

### Return type

[**VersionedObjectStateChangeRequest**](VersionedObjectStateChangeRequest.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMsgSvcAssociationForEventVersion**
> MessagingServiceAssociationResponse UpdateMsgSvcAssociationForEventVersion(ctx, body, id)
Update messaging service association for an event version object

Use this API to update the messaging service association for an event version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MessagingServiceAssociationDto**](MessagingServiceAssociationDto.md)| The messaging service association object | 
  **id** | **string**| The ID of the event version | 

### Return type

[**MessagingServiceAssociationResponse**](MessagingServiceAssociationResponse.md)

### Authorization

[APIToken](../README.md#APIToken)

### HTTP request headers

 - **Content-Type**: application/json;charset=UTF-8
 - **Accept**: application/json;charset=UTF-8, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

