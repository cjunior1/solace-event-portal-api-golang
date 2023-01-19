# EventManagementAgent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTime** | **string** |  | [optional] [default to null]
**UpdatedTime** | **string** |  | [optional] [default to null]
**CreatedBy** | **string** |  | [optional] [default to null]
**ChangedBy** | **string** |  | [optional] [default to null]
**Id** | **string** | Primary key set by the server. | [optional] [default to null]
**Name** | **string** | The name of the EMA. | [default to null]
**Region** | **string** | The region in which the EMA belongs to, extracted from the EventManagementAgentRegion. | [optional] [default to null]
**ClientUsername** | **string** | The SMF username for a customer&#x27;s EMA to use to communicate to event-portal. | [optional] [default to null]
**ClientPassword** | **string** | The SMF password for a customer&#x27;s EMA to use to communicate to event-portal. | [optional] [default to null]
**ReferencedByMessagingServiceIds** | **[]string** |  | [optional] [default to null]
**OrgId** | **string** | Used by admin APIs to get a list of EMAs against the given orgId | [optional] [default to null]
**Status** | **string** | The connection status of EP to the actual EMA which this object represents. | [optional] [default to null]
**EventManagementAgentRegionId** | **string** | The ID of the associated EventManagementAgentRegion. | [default to null]
**Type_** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

