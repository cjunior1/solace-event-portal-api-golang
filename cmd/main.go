package main

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	eventportal "github.com/cjunior1/solace-event-portal-api-golang"
)

func main() {

	//	apiKey := swagger.APIKey{Key: "eyJhbGciOiJSUzI1NiIsImtpZCI6Im1hYXNfcHJvZF8yMDIwMDMyNiIsInR5cCI6IkpXVCJ9.eyJvcmciOiJyMndmMDJ3ZGVmOSIsIm9yZ1R5cGUiOiJFTlRFUlBSSVNFIiwic3ViIjoidTlidDE3YTNnaGciLCJwZXJtaXNzaW9ucyI6IkFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBUUFBRUFJQUFBSkFLcmdaUWZBUVFBQUNJQlNwSXFKZz09IiwiYXBpVG9rZW5JZCI6IjgwdGpza3pjYmg1IiwiaXNzIjoiU29sYWNlIENvcnBvcmF0aW9uIiwiaWF0IjoxNjczOTIxMzY0fQ.bNlMRlqKuF4QUCto-0TEf0HkCk4kk0-j6Moe4tv6AADx3rFJF9gMMSvjnVjCDW1pgeBfcTgkoztkelWju5wSWXB154k71W0bu-614QuEK2ecR-_7KibVzzwRyEGYM7eVNef-5hqSHrcBOZlcI6ZBzHt-PuYOLupuEEwM8I7c6awJYTPpjloJx6aijGq38nUbaZkVrQuyFMboZuBQf8rKrpd4_fq2hd7pV13ddfInWSxN--pAuxMr13QxThAXElM6CwgxaJBM-1vSZVLWEczOP8fO7CT85BoEpw3BJxXsRlwamsCKew_Z23kqtuH2jfVVCvGv-SEDRHjKTlnuFsGDXQ", Prefix: "Bearer"}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = context.WithValue(ctx, eventportal.ContextAccessToken, "eyJhbGciOiJSUzI1NiIsImtpZCI6Im1hYXNfcHJvZF8yMDIwMDMyNiIsInR5cCI6IkpXVCJ9.eyJvcmciOiJyMndmMDJ3ZGVmOSIsIm9yZ1R5cGUiOiJFTlRFUlBSSVNFIiwic3ViIjoidTlidDE3YTNnaGciLCJwZXJtaXNzaW9ucyI6IkFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBQUFBUUFBRUFJQUFBSkFLcmdaUWZBUVFBQUNJQlNwSXFKZz09IiwiYXBpVG9rZW5JZCI6IjgwdGpza3pjYmg1IiwiaXNzIjoiU29sYWNlIENvcnBvcmF0aW9uIiwiaWF0IjoxNjczOTIxMzY0fQ.bNlMRlqKuF4QUCto-0TEf0HkCk4kk0-j6Moe4tv6AADx3rFJF9gMMSvjnVjCDW1pgeBfcTgkoztkelWju5wSWXB154k71W0bu-614QuEK2ecR-_7KibVzzwRyEGYM7eVNef-5hqSHrcBOZlcI6ZBzHt-PuYOLupuEEwM8I7c6awJYTPpjloJx6aijGq38nUbaZkVrQuyFMboZuBQf8rKrpd4_fq2hd7pV13ddfInWSxN--pAuxMr13QxThAXElM6CwgxaJBM-1vSZVLWEczOP8fO7CT85BoEpw3BJxXsRlwamsCKew_Z23kqtuH2jfVVCvGv-SEDRHjKTlnuFsGDXQ")

	cfg := eventportal.NewConfiguration()

	client := eventportal.NewAPIClient(cfg)

	ca := optional.NewString("eventName==demo")
	resp, _, err := client.SchemasApi.GetSchemaVersions(ctx, &eventportal.SchemasApiGetSchemaVersionsOpts{CustomAttributes: ca})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Data[0].Content)

}
