package repo

type UserConsent_FIDataRange struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

type UserConsentInfo struct {
	TransactionId string                   `json:"txn_id"`
	ConsentHandle string                   `json:"consent_handle"`
	FIDataRange   *UserConsent_FIDataRange `json:"fi_data_range"`
}
