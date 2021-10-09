package data

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"amit.vickey/ZenMoney/cache"
	"amit.vickey/ZenMoney/common"
	"amit.vickey/ZenMoney/repo"
	"amit.vickey/ZenMoney/setu/apicaller"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func FIDataRequest(ctx *gin.Context, consentId, consentHandle, signedConsent string) error {
	keyMaterialResponse, err := generateKeyMaterialViaSetuRahasya(ctx)
	if err != nil {
		return errors.Wrap(err, "error-generating-key-material")
	}

	fiDateStart := time.Now().AddDate(0, -5, 0).Format(time.RFC3339)
	fiDateEnd := time.Now().AddDate(0, 0, -1).Format(time.RFC3339)

	// fiDateStart, fiDateEnd := cache.GetFiDateForCosentHandle(consentHandle)

	request := SetuFIDataRequestRequest{
		Version:       common.SetuVersion,
		Timestamp:     time.Now().Format(time.RFC3339),
		TransactionId: uuid.New().String(),
		FIDataRange: &FIDataRequest_FataRange{
			From: fiDateStart,
			To:   fiDateEnd,
		},
		Consent: &FIDataRequest_Consent{
			Id:               consentId,
			DigitalSignature: strings.Split(signedConsent, ".")[2],
		},
		KeyMaterial: keyMaterialResponse.KeyMaterial,
	}

	status, httpResponse, err := apicaller.CallSetuAAApi(ctx, request, "FI/request", "POST")
	if err != nil {
		return errors.Wrap(err, "error-calling-setu-api")
	}

	if status == 200 {
		var dataResponse SetuFIDataRequestResponse
		if err := json.Unmarshal(httpResponse, &dataResponse); err != nil {
			return errors.Wrap(err, "error-unmarshaling-create-consent-response")
		}
		common.LogInfo(ctx, fmt.Sprintf("sesstion-id: %s", dataResponse.SessionId))
		cache.AddConsentIdConsentHandle(consentId, consentHandle)
		cache.AddSessionIdConsentId(dataResponse.SessionId, consentId)
		cache.AddSessionIdPrivateKey(dataResponse.SessionId, keyMaterialResponse.PrivateKey)
		cache.AddSessionIdKeyMaterial(dataResponse.SessionId, keyMaterialResponse.KeyMaterial)
		ctx.JSON(http.StatusOK, gin.H{
			"session_id": dataResponse.SessionId,
		})
	} else {
		common.LogInfo(ctx, fmt.Sprintf("status: %v, body: %v", status, string(httpResponse)))
		return errors.New(fmt.Sprintf("status: %v, body: %v", status, string(httpResponse)))
	}

	return nil
}

func generateKeyMaterialViaSetuRahasya(ctx *gin.Context) (RahasyaGenerateKeyMaterialResponse, error) {
	var response RahasyaGenerateKeyMaterialResponse
	httpResponse, err := http.Get(fmt.Sprintf("%s/%s", common.SetuRahasyaBaseUrl, common.SetuRahasyaGenerateKeyRoute))
	if err != nil {
		return response, errors.Wrap(err, "error-making-rahasya-call")
	}
	defer httpResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return response, errors.Wrap(err, "error-reading-http-response")
	}

	if httpResponse.StatusCode == 200 {
		if err := json.Unmarshal(responseBody, &response); err != nil {
			return response, errors.Wrap(err, "error-unmarshaling-generate-key-material-response")
		}
		return response, nil
	} else {
		common.LogInfo(ctx, fmt.Sprintf("status: %v, body: %v", httpResponse.StatusCode, string(responseBody)))
		return response, fmt.Errorf("status: %v, body: %v", httpResponse.StatusCode, string(responseBody))
	}
}

func FetchDataFromSetu(ctx *gin.Context, sessionId string) error {
	route := fmt.Sprintf("%s/%s", "FI/fetch", sessionId)
	status, responseBody, err := apicaller.CallSetuAAApi(ctx, nil, route, "GET")
	if err != nil {
		return errors.Wrap(err, "error-making-setu-api-call")
	}
	if status == 200 {
		var dataResponse SetuFIDataResponse
		if err := json.Unmarshal(responseBody, &dataResponse); err != nil {
			return errors.Wrap(err, "error-unmarshaling-fi-data-response")
		}
		decryptFIData(ctx, dataResponse.FIData,
			cache.GetPrivateKeyForSessionId(sessionId),
			cache.GetKeyMaterialForSessionId(sessionId).(*KeyMaterial),
			cache.GetUserIdForSessionId(sessionId))
	}
	return nil
}

func decryptFIData(ctx *gin.Context, fiDataList []*FIData, privateKey string, keyMaterial *KeyMaterial, userId string) error {
	for _, fiData := range fiDataList {
		for _, data := range fiData.Data {
			decryptDataRequest := DecryptFIDataRequest{
				Base64Data:        data.EncryptedFI,
				Base64RemoteNonce: fiData.KeyMaterial.Nonce,
				Base64YourNonce:   keyMaterial.Nonce,
				OurPrivateKey:     privateKey,
				RemoteKeyMaterial: fiData.KeyMaterial,
			}
			requestBuffer := new(bytes.Buffer)
			json.NewEncoder(requestBuffer).Encode(decryptDataRequest)
			httpResponse, err := http.Post(fmt.Sprintf("%s/%s", common.SetuRahasyaBaseUrl, common.SetuRahasyaDecryptDataRoute), "application/json", requestBuffer)
			if err != nil {
				common.LogInfo(ctx, fmt.Sprintf("error-making-rahasya-call-to-decrypt-data: %v", err.Error()))
				continue
			}
			defer httpResponse.Body.Close()

			status := httpResponse.StatusCode
			responseBody, _ := ioutil.ReadAll(httpResponse.Body)
			if status == 200 {
				var decryptResponse DecryptFIDataResponse
				if err := json.Unmarshal(responseBody, &decryptResponse); err != nil {
					common.LogInfo(ctx, fmt.Sprintf("error-unmarshaling-decrypt-data-response: %v", err.Error()))
					continue
				}
				dataBytes, _ := base64.StdEncoding.DecodeString(decryptResponse.Base64Data)
				common.LogInfo(ctx, string(dataBytes))
				addDataToDatabase(ctx, userId, dataBytes)
			} else {
				common.LogInfo(ctx, fmt.Sprintf("status: %v, body: %v", status, string(responseBody)))
			}
		}
	}
	return nil
}

func addDataToDatabase(ctx *gin.Context, userId string, dataBytes []byte) error {
	var dataObj FIDataStructure
	if err := json.Unmarshal(dataBytes, &dataObj); err != nil {
		return errors.Wrap(err, "error-unmarshaling-fi-data-response-to-structure")
	}
	fiDataType := dataObj.Account["type"].(string)
	if err := repo.AddFIDataForUser(ctx, userId, fiDataType, dataObj.Account); err != nil {
		return errors.Wrap(err, "error-saving-data-to-database")
	}
	common.LogInfo(ctx, fmt.Sprintf("inserted-data: user-id: %s, fi-data-type: %s", userId, fiDataType))
	return nil
}
