package consent

type SetuWebhookRequest_Notifier struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type SetuWebhookRequest_ConsentStatusNotification struct {
	ConsentId     string `json:"consentId"`
	ConsentHandle string `json:"consentHandle"`
	ConsentStatus string `json:"consentStatus"`
}

type FIPAccounts struct {
	LinkRefNumber string `json:"linkRefNumber"`
	FIStatus      string `json:"FIStatus"`
	Description   string `json:"description"`
}

type FIStatusResponse struct {
	FipId     string         `json:"fipID"`
	Acccounts []*FIPAccounts `json:"Accounts"`
}

type SetuWebhookRequest_FIStatusNotification struct {
	SessionId        string              `json:"sessionId"`
	SessionStatus    string              `json:"sessionStatus"`
	FIStatusResponse []*FIStatusResponse `json:"FIStatusResponse"`
}

type SetuWebhookRequest struct {
	Version                   string                                        `json:"ver"`
	Timestamp                 string                                        `json:"timestamp"`
	TransactionId             string                                        `json:"txnid"`
	Notifier                  *SetuWebhookRequest_Notifier                  `json:"Notifier"`
	ConsentStatusNotification *SetuWebhookRequest_ConsentStatusNotification `json:"ConsentStatusNotification"`
	FIStatusNotification      *SetuWebhookRequest_FIStatusNotification      `json:"FIStatusNotification"`
}

type SetuWebhookResponse struct {
	Version       string `json:"ver"`
	Timestamp     string `json:"timestamp"`
	TransactionId string `json:"txnid"`
	Response      string `json:"response"`
}

type Consent_DataConsumer struct {
	Id string `json:"id"`
}

type Consent_Customer struct {
	Id string `json:"id"`
}

type Consent_Purpose struct {
	Code     string            `json:"code"`
	RefUri   string            `json:"refUri"`
	Text     string            `json:"text"`
	Category map[string]string `json:"Category"`
}

type Consent_DataLife struct {
	Unit  string `json:"unit"`
	Value int    `json:"value"`
}

type Consent_Frequency struct {
	Unit  string `json:"unit"`
	Value int    `json:"value"`
}

type Consent_DataFilter struct {
	Type     string `json:"type"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Consent_FIDataRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type SetuCreateConsentRequest_Consent struct {
	ConsentStart  string                `json:"consentStart"`
	ConsentExpiry string                `json:"consentExpiry"`
	ConsentMode   string                `json:"consentMode"`
	FetchType     string                `json:"fetchType"`
	ConsentTypes  []string              `json:"consentTypes"`
	FiTypes       []string              `json:"fiTypes"`
	DataConsumer  *Consent_DataConsumer `json:"DataConsumer"`
	Customer      *Consent_Customer     `json:"Customer"`
	Purpose       *Consent_Purpose      `json:"Purpose"`
	FIDataRange   *Consent_FIDataRange  `json:"FIDataRange"`
	DataLife      *Consent_DataLife     `json:"DataLife"`
	Frequency     *Consent_Frequency    `json:"Frequency"`
	DataFilter    []*Consent_DataFilter `json:"DataFilter,omitempty"`
}

type SetuCreateConsentRequest struct {
	Version       string                            `json:"ver"`
	Timestamp     string                            `json:"timestamp"`
	TransactionId string                            `json:"txnid"`
	ConsentDetail *SetuCreateConsentRequest_Consent `json:"ConsentDetail"`
}

type SetuCreateConsentResponse struct {
	Version       string            `json:"ver"`
	Timestamp     string            `json:"timestamp"`
	TransactionId string            `json:"txnid"`
	Customer      *Consent_Customer `json:Customer`
	ConsentHandle string            `json:"ConsentHandle"`
}

type FetchConsent_ConsentUse struct {
	LogUri          string `json:"logUri"`
	Count           int    `json:"count"`
	LastUseDateTime string `json:"lastUseDateTime"`
}

type SetuFetchConsentResponse struct {
	Version         string                   `json:"ver"`
	TransactionId   string                   `json:"txnid"`
	ConsentId       string                   `json:"consentId"`
	Status          string                   `json:"status"`
	CreateTimestamp string                   `json:"createTimestamp"`
	SignedConsent   string                   `json:"signedConsent"`
	ConsentUse      *FetchConsent_ConsentUse `json:ConsentUse`
}

type CreateConsentRequest struct {
	ConsentTypes []string `json:"consent_types"`
	FiTypes      []string `json:"fi_types"`
	PhoneNumber  string   `json:"phone_number"`
}
