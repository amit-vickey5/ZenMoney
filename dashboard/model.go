package dashboard

type DashboardDataRequest struct {
	UserId string `json:"user_id"`
}

type FIData struct {
	Deposit           []interface{} `json:"DEPOSIT"`
	RecurringDeposit  []interface{} `json:"RECURRING_DEPOSIT"`
	TermDeposit       []interface{} `json:"TERM_DEPOSIT"`
	CreditCard        []interface{} `json:"CREDIT_CARD"`
	CD                []interface{} `json:"CD"`
	CP                []interface{} `json:"CP"`
	IDR               []interface{} `json:"IDR"`
	MutualFunds       []interface{} `json:"MUTUAL_FUNDS"`
	Bonds             []interface{} `json:"BONDS"`
	Equities          []interface{} `json:"EQUITIES"`
	PPF               []interface{} `json:"PPF"`
	EPF               []interface{} `json:"EPF"`
	ETF               []interface{} `json:"ETF"`
	NPS               []interface{} `json:"NPS"`
	GovtSecurities    []interface{} `json:"GOVT_SECURITIES"`
	InsurancePolicies []interface{} `json:"INSURANCE_POLICIES"`
	ULIP              []interface{} `json:"ULIP"`
	AIF               []interface{} `json:"AIF"`
	INVIT             []interface{} `json:"INVIT"`
	SIP               []interface{} `json:"SIP"`
	REIT              []interface{} `json:"REIT"`
	CIS               []interface{} `json:"CIS"`
	Debentures        []interface{} `json:"DEBENTURES"`
}

type DashboardDataResponse struct {
	Id            string
	IsActive      bool
	FirstName     string
	LastName      string
	Email         string
	PrimaryMobile string
	AaHandle      string
	FIData        FIData
	CreatedAt     string
	UpdatedAt     string
}
