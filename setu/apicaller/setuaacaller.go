package apicaller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"amit.vickey/ZenMoney/common"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func CallSetuAAApi(ctx *gin.Context, request interface{}, route, method string) (int, []byte, error) {
	detachedJWS, err := getDetachedJWS(ctx, request)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error-preparing-jws-token")
	}

	requestBuffer := new(bytes.Buffer)
	json.NewEncoder(requestBuffer).Encode(request)

	url := fmt.Sprintf("%s/%s", common.SetuBaseUrl, route)

	httpRequest, err := http.NewRequest(method, url, requestBuffer)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error-preparing-http-request")
	}
	httpRequest.Header.Add("Content-Type", "application/json")
	httpRequest.Header.Add("client_api_key", common.SetuClientApiKey)
	httpRequest.Header.Add("x-jws-signature", detachedJWS)

	httpClient := http.Client{}
	httpResponse, err := httpClient.Do(httpRequest)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error-making-http-call")
	}
	defer httpResponse.Body.Close()

	responseStatus := httpResponse.StatusCode
	responseBody, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return 0, nil, errors.Wrap(err, "error-reading-http-response")
	}
	return responseStatus, responseBody, nil
}

func getDetachedJWS(ctx *gin.Context, payload interface{}) (string, error) {

	privateKey, err := ioutil.ReadFile("././keys/private_key.pem")
	if err != nil {
		return "", errors.Wrap(err, "error-reading-private-pem-file")
	}

	token := jwt.New(jwt.SigningMethodRS256)
	pkey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", errors.Wrap(err, "error-parsing-rsa-private-key-from-pem")
	}
	jwt, err := token.SignedString(pkey)
	if err != nil {
		return "", errors.Wrap(err, "error-signing-token")
	}

	splittedJwt := strings.Split(jwt, ".")
	splittedJwt[1] = ""

	return strings.Join(splittedJwt, "."), nil
}
