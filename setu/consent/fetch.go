package consent

import (
	"encoding/json"
	"fmt"

	"amit.vickey/ZenMoney/setu/apicaller"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func FetchSignedConsent(ctx *gin.Context, consentId string) (string, error) {
	route := fmt.Sprintf("Consent/%s", consentId)

	status, response, err := apicaller.CallSetuAAApi(ctx, nil, route, "GET")

	if err != nil {
		return "", errors.Wrap(err, "error-calling-setu-api")
	}
	if status == 200 {
		var setuFetchConsentResponse SetuFetchConsentResponse
		if err := json.Unmarshal(response, &setuFetchConsentResponse); err != nil {
			return "", errors.Wrap(err, "error-unmarshaling-fetch-consent-response")
		}
		return setuFetchConsentResponse.SignedConsent, nil
	} else {
		return "", fmt.Errorf("status: %v, body: %v", status, string(response))
	}
}
