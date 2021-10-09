package common

const (
	SetuVersion = "1.0"
	SetuHandle  = "setu-aa"

	SetuBaseUrl                 = "https://aa-sandbox.setu.co"
	SetuRahasyaBaseUrl          = "https://rahasya.setu.co"
	SetuRahasyaGenerateKeyRoute = "ecc/v1/generateKey"
	SetuRahasyaDecryptDataRoute = "ecc/v1/decrypt"

	SetuDataConsumerId = "fiu-zen-pfm" //identifier for the entity that’s requesting for the data

	ConsentMode_View   = "VIEW"
	ConsentMode_Store  = "STORE"
	ConsentMode_Query  = "QUERY"
	ConsentMode_Stream = "STREAM"

	FetchType_OneTime  = "ONETIME"
	FetchType_Periodic = "PERIODIC"

	DataLife_Infinity = "INF"

	SetuClientApiKey = "e2fceb3a-3e55-4df1-98fe-a3147573202f"
)

const (
	ZenServerBaseUrl = "http://localhost:8080"
)

var FIDataTypes = []string{
	"deposit",
	"recurring_deposit",
	"term_deposit",
	"credit_card",
	"cd",
	"idr",
	"mutual_funds",
	"bonds",
	"debentures",
	"equities",
	"ppf",
	"epf",
	"etf",
	"nps",
	"cp",
	"govt_securities",
	"insurance_policies",
	"ulip",
	"reit",
	"invit",
	"aif",
	"sip",
	"cis",
}
