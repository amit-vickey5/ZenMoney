package consent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"amit.vickey/ZenMoney/cache"
	"amit.vickey/ZenMoney/common"
	"amit.vickey/ZenMoney/repo"
	"amit.vickey/ZenMoney/setu/apicaller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func Create(ctx *gin.Context) {
	// create a struct SetuCreateConsentRequest out of the input
	// create a consent request on Setu i.e. make API call
	// associate the ConsentHandle in response with Cusomer.Id (our user's id)
	// return back ConsentHandle
	var request CreateConsentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		common.ErrorResponse(ctx, errors.Wrap(err, "error-converting-input-to-request-object"))
		return
	}

	consentStartTime := time.Now()
	consentExpiryTime := consentStartTime.Add(30 * time.Minute)

	fiDateStartTime := time.Now().AddDate(0, -6, 0).Format(time.RFC3339)
	fiDateEndTime := time.Now().Format(time.RFC3339)

	setuConsentRequest := SetuCreateConsentRequest{
		Version:       common.SetuVersion,
		Timestamp:     time.Now().Format(time.RFC3339),
		TransactionId: uuid.New().String(),
		ConsentDetail: &SetuCreateConsentRequest_Consent{
			ConsentStart:  consentStartTime.Format(time.RFC3339),
			ConsentExpiry: consentExpiryTime.Format(time.RFC3339),
			ConsentMode:   common.ConsentMode_Store,
			FetchType:     common.FetchType_Periodic,
			ConsentTypes:  request.ConsentTypes, //from FE
			FiTypes:       request.FiTypes,      //from FE
			DataConsumer: &Consent_DataConsumer{
				Id: common.SetuDataConsumerId,
			},
			Customer: &Consent_Customer{
				Id: common.GetSetuAACustomerHandle(request.PhoneNumber), //from FE
			},
			Purpose: &Consent_Purpose{
				Code:   "101",
				RefUri: "https://api.rebit.org.in/aa/purpose/101.xml",
				Text:   "Wealth management service",
				Category: map[string]string{
					"type": "Personal Finance",
				},
			},
			FIDataRange: &Consent_FIDataRange{
				From: fiDateStartTime,
				To:   fiDateEndTime,
			},
			DataLife: &Consent_DataLife{
				Unit:  common.DataLife_Infinity,
				Value: 0,
			},
			Frequency: &Consent_Frequency{
				Unit:  "HOUR",
				Value: 1,
			},
		},
	}

	status, response, err := apicaller.CallSetuAAApi(ctx, setuConsentRequest, "Consent", "POST")
	if err != nil {
		common.ErrorResponse(ctx, errors.Wrap(err, "error-calling-setu-api"))
		return
	}
	if status == 200 {
		var setuConsentResponse SetuCreateConsentResponse
		if err := json.Unmarshal(response, &setuConsentResponse); err != nil {
			common.ErrorResponse(ctx, errors.Wrap(err, "error-unmarshaling-create-consent-response"))
			return
		}
		cache.AddConsentHandlerUserId(setuConsentResponse.ConsentHandle, request.PhoneNumber)
		cache.AddConsentHandlerConsentTimes(setuConsentResponse.ConsentHandle, fiDateStartTime, fiDateEndTime)
		// if err := addConsentForUser(ctx, "", setuConsentRequest, setuConsentResponse); err != nil {
		// 	common.Logger.Write([]byte(err.Error() + "\n"))
		// }
		ctx.JSON(http.StatusOK, gin.H{
			"consent_handle": setuConsentResponse.ConsentHandle,
			"redirect_url":   fmt.Sprintf("%s/%s", common.ZenServerBaseUrl, "ConsentRedirect"),
		})
	} else {
		common.LogInfo(ctx, fmt.Sprintf("status: %v, body: %v", status, string(response)))
	}
}

func addConsentForUser(ctx *gin.Context, userId string, request SetuCreateConsentRequest, response SetuCreateConsentResponse) error {
	fiDataFrom, err := common.GetTimestampFromISO(request.ConsentDetail.FIDataRange.From)
	if err != nil {
		return errors.Wrap(err, "error-parsing-from-fi-data-range")
	}
	fiDataTo, err := common.GetTimestampFromISO(request.ConsentDetail.FIDataRange.To)
	if err != nil {
		return errors.Wrap(err, "error-parsing-to-fi-data-range")
	}
	userConsentInfo := repo.UserConsentInfo{
		TransactionId: request.TransactionId,
		ConsentHandle: response.ConsentHandle,
		FIDataRange: &repo.UserConsent_FIDataRange{
			From: fiDataFrom,
			To:   fiDataTo,
		},
	}
	err = repo.AddConsentForUser(ctx, userId, userConsentInfo)
	if err != nil {
		return errors.Wrap(err, "error-adding-consent-for-user")
	}
	return nil
}
