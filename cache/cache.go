package cache

type fiDate struct {
	start string
	end   string
}

var consentHandleUserIdCache = make(map[string]string)
var consentIdConsentHandleCache = make(map[string]string)
var sessionIdConsentIdCache = make(map[string]string)
var sessionIdPrivateKeyCache = make(map[string]string)
var sessionIdKeyMaterialCache = make(map[string]interface{})
var consentHandlerFiDateCache = make(map[string]fiDate)

func InitCache() {

}

func AddConsentHandlerUserId(consentHandler, userId string) {
	consentHandleUserIdCache[consentHandler] = userId
}

func AddConsentIdConsentHandle(consentId, consentHandler string) {
	consentIdConsentHandleCache[consentId] = consentHandler
}

func AddSessionIdConsentId(sessionId, consentId string) {
	sessionIdConsentIdCache[sessionId] = consentId
}

func AddSessionIdPrivateKey(sessionId, privateKey string) {
	sessionIdPrivateKeyCache[sessionId] = privateKey
}

func GetPrivateKeyForSessionId(sessionId string) string {
	return sessionIdPrivateKeyCache[sessionId]
}

func AddSessionIdKeyMaterial(sessionId string, keyMaterial interface{}) {
	sessionIdKeyMaterialCache[sessionId] = keyMaterial
}

func GetKeyMaterialForSessionId(sessionId string) interface{} {
	return sessionIdKeyMaterialCache[sessionId]
}

func AddConsentHandlerConsentTimes(consentHandler, fiDateStartTime, fiDateEndTime string) {
	consentHandlerFiDateCache[consentHandler] = fiDate{
		start: fiDateStartTime,
		end:   fiDateEndTime,
	}
}

func GetFiDateForCosentHandle(consentHandle string) (string, string) {
	dateFromCache := consentHandlerFiDateCache[consentHandle]
	return dateFromCache.start, dateFromCache.end
}

func GetUserIdForSessionId(sessionId string) string {
	return consentHandleUserIdCache[consentIdConsentHandleCache[sessionIdConsentIdCache[sessionId]]]
}

/* func GetConsentHandlerForUser(userId string) (string, error) {
	if val, ok := userConsentCache[userId]; ok {
		consentInfo := val.(map[string]string)
		return consentInfo["consent_handler"], nil
	} else {
		return "", errors.New("no-consent-handler-for-user-in-cache")
	}
}

func AddConsentIdForUser(userId, consentId string) {
	var consentInfo map[string]string
	if val, ok := userConsentCache[userId]; ok {
		consentInfo = val.(map[string]string)
		consentInfo["consent_id"] = consentId
	} else {
		consentInfo = make(map[string]string)
		consentInfo["consent_id"] = consentId
	}
	userConsentCache[userId] = consentInfo
} */
