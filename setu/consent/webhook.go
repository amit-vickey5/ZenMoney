package consent

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"amit.vickey/ZenMoney/common"
	"amit.vickey/ZenMoney/setu/data"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func SetuWebhook(ctx *gin.Context) {
	var request SetuWebhookRequest
	// defer ctx.Request.Body.Close()
	// requestBuf := new(strings.Builder)
	// _, _ = io.Copy(requestBuf, ctx.Request.Body)
	// common.LogInfo(ctx, requestBuf.String())
	if err := ctx.ShouldBindJSON(&request); err != nil {
		common.ErrorResponse(ctx, errors.Wrap(err, "error-converting-input-to-request-object"))
		return
	}

	// if err := validateWebhookRequest(ctx); err != nil {
	// 	common.ErrorResponse(ctx, errors.Wrap(err, "error-validating-webhook-request"))
	// 	return
	// }

	//duplicateCtx := ctx.Copy()

	// different operations based on whether it's a consent notification, or FI notification
	if request.ConsentStatusNotification != nil {
		common.LogInfo(ctx, "received-consent-notification")
		consentWebhookHandler(ctx, request)
	} else if request.FIStatusNotification != nil {
		common.LogInfo(ctx, "received-fi-data-notification")
		fiDataWebhookHandler(ctx, request)
	}

	// send proper response
}

func validateWebhookRequest(ctx *gin.Context) error {
	plan, err := ioutil.ReadFile("././keys/setu_public_key.json")
	if err != nil {
		return errors.Wrap(err, "error-reading-setu-public-key")
	}
	var data interface{}
	if err := json.Unmarshal(plan, &data); err != nil {
		return errors.Wrap(err, "error-unmarshaling-setu-public-key")
	}

	publicKey := ""

	detachedJWS := http.CanonicalHeaderKey("x-jws-signature")
	requestBody := ctx.Request.Body
	defer ctx.Request.Body.Close()

	splittedJws := strings.Split(detachedJWS, ".")

	requestBuf := new(strings.Builder)
	_, err = io.Copy(requestBuf, requestBody)
	if err != nil {
		return errors.Wrap(err, "error-reading-request-body")
	}

	splittedJws[1] = base64.StdEncoding.EncodeToString([]byte(requestBuf.String()))
	requestJwt := strings.Join(splittedJws, ".")

	token := jwt.New(jwt.SigningMethodRS256)
	pkey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
	if err != nil {
		return errors.Wrap(err, "error-parsing-rsa-public-key-from-pem")
	}

	signedString, err := token.Method.Sign(requestJwt, pkey)
	if err != nil {
		return errors.Wrap(err, "error-parsing-rsa-public-key-from-pem")
	}

	common.LogInfo(ctx, signedString)

	return nil
}

func consentWebhookHandler(ctx *gin.Context, request SetuWebhookRequest) error {
	consentId := request.ConsentStatusNotification.ConsentId
	consentHandle := request.ConsentStatusNotification.ConsentHandle
	consentStatus := request.ConsentStatusNotification.ConsentStatus

	if consentStatus != "ACTIVE" {
		common.LogInfo(ctx, fmt.Sprintf("non-active-consent: %v", consentStatus))
		return nil
	}

	signedConsent, err := FetchSignedConsent(ctx, consentId)
	if err != nil {
		return errors.Wrap(err, "error-fetching-signed-consent")
	}

	err = data.FIDataRequest(ctx, consentId, consentHandle, signedConsent)
	if err != nil {
		return errors.Wrap(err, "error-making-data-request")
	}

	//repo.UpdateConsentIdAndSignedConsentForUser(ctx, consentId, signedConsent)

	return nil
}

func fiDataWebhookHandler(ctx *gin.Context, request SetuWebhookRequest) error {
	data.FetchDataFromSetu(ctx, request.FIStatusNotification.SessionId)
	return nil
}
