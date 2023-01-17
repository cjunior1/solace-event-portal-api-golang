/*
 * Event Portal Open API
 *
 * REST API Concepts  Solace PubSub+ Event Portal provides REST APIs that you can use to manage your data in PubSub+ Cloud.  The REST APIs allow you to model your event-driven architectures from your own client applications.  The following document describes the public REST APIs available for use in Event Portal 2.0. Objects created in this  version will not be available in Event Portal 1.0. APIs that display (Beta) in their summary are provided as-is and  are subject to change or removal. They may not be of the quality expected for generally available APIs and have no  guaranteed forward compatibility with the generally available version of the API.  
 *
 * API version: 2.0.11
 * Generated by: Swagger Codegen (https://github.com/eventportal-api/eventportal-codegen.git)
 */
package eventportal

type EventManagementAgent struct {
	CreatedTime string `json:"createdTime,omitempty"`
	UpdatedTime string `json:"updatedTime,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	ChangedBy string `json:"changedBy,omitempty"`
	// Primary key set by the server.
	Id string `json:"id,omitempty"`
	// The name of the EMA.
	Name string `json:"name"`
	// The region in which the EMA belongs to, extracted from the EventManagementAgentRegion.
	Region string `json:"region,omitempty"`
	// The SMF username for a customer's EMA to use to communicate to event-portal.
	ClientUsername string `json:"clientUsername,omitempty"`
	// The SMF password for a customer's EMA to use to communicate to event-portal.
	ClientPassword string `json:"clientPassword,omitempty"`
	ReferencedByMessagingServiceIds []string `json:"referencedByMessagingServiceIds,omitempty"`
	// Used by admin APIs to get a list of EMAs against the given orgId
	OrgId string `json:"orgId,omitempty"`
	// The connection status of EP to the actual EMA which this object represents.
	Status string `json:"status,omitempty"`
	// The ID of the associated EventManagementAgentRegion.
	EventManagementAgentRegionId string `json:"eventManagementAgentRegionId"`
	Type_ string `json:"type,omitempty"`
}
