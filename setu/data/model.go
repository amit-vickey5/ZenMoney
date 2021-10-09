package data

type DHPublicKey struct {
	Expiry     string `json:"expiry"`
	Parameters string `json:"Parameters"`
	KeyValue   string `json:"KeyValue"`
}

type KeyMaterial struct {
	CryptoAlg   string       `json:"cryptoAlg"`
	Curve       string       `json:"curve"`
	Params      string       `json:"params"`
	DHPublicKey *DHPublicKey `json:"DHPublicKey"`
	Nonce       string       `json:"Nonce"`
}

type RahasyaGenerateKeyMaterialResponse struct {
	PrivateKey  string       `json:"privateKey"`
	ErrorInfo   interface{}  `json:"errorInfo"`
	KeyMaterial *KeyMaterial `json:"KeyMaterial"`
}

type FIDataRequest_FataRange struct {
	From string `json:"from"`
	To   string `json:"to"`
}

type FIDataRequest_Consent struct {
	Id               string `json:"id"`
	DigitalSignature string `json:"digitalSignature"`
}

type SetuFIDataRequestRequest struct {
	Version       string                   `json:"ver"`
	Timestamp     string                   `json:"timestamp"`
	TransactionId string                   `json:"txnid"`
	FIDataRange   *FIDataRequest_FataRange `json:"FIDataRange"`
	Consent       *FIDataRequest_Consent   `json:"Consent"`
	KeyMaterial   *KeyMaterial             `json:"KeyMaterial"`
}

type SetuFIDataRequestResponse struct {
	Version       string `json:"ver"`
	Timestamp     string `json:"timestamp"`
	TransactionId string `json:"txnid"`
	ConsentId     string `json:"consentId"`
	SessionId     string `json:"sessionId"`
}

type SetuFIDataResponse struct {
	Version       string    `json:"ver"`
	Timestamp     string    `json:"timestamp"`
	TransactionId string    `json:"txnid"`
	FIData        []*FIData `json:"FI"`
}

type FIData struct {
	FipID         string       `json:"fipID"`
	Timestamp     string       `json:"timestamp"`
	TransactionId string       `json:"txnid"`
	Data          []*DataInfo  `json:"data"`
	KeyMaterial   *KeyMaterial `json:"KeyMaterial"`
}

type DataInfo struct {
	LinkRefNumber   string `json:"linkRefNumber"`
	MaskedAccNumber string `json:"maskedAccNumber"`
	EncryptedFI     string `json:"encryptedFI"`
}

type DecryptFIDataRequest struct {
	Base64Data        string       `json:"base64Data"`
	Base64RemoteNonce string       `json:"base64RemoteNonce"`
	Base64YourNonce   string       `json:"base64YourNonce"`
	OurPrivateKey     string       `json:"ourPrivateKey"`
	RemoteKeyMaterial *KeyMaterial `json:"remoteKeyMaterial"`
}

type DecryptFIDataResponse struct {
	Base64Data string      `json:"base64Data"`
	ErrorInfo  interface{} `json:"errorInfo"`
}

type FIDataStructure struct {
	Account map[string]interface{} `json:"account"`
}
