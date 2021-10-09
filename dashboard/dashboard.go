package dashboard

import (
	"fmt"
	"net/http"

	"amit.vickey/ZenMoney/common"
	"amit.vickey/ZenMoney/repo"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func DashboardData(ctx *gin.Context) {
	var request DashboardDataRequest
	var response DashboardDataResponse
	if err := ctx.ShouldBindJSON(&request); err != nil {
		common.ErrorResponse(ctx, errors.Wrap(err, "error-converting-dashboard-input-to-request-object"))
		return
	}
	response.IsActive = true
	response.FirstName = "Test"
	response.LastName = "Test"
	response.Email = "test@gmail.com"
	response.PrimaryMobile = request.UserId
	response.AaHandle = common.GetSetuAACustomerHandle(request.UserId)
	// userMetadata := repo.GetUserMetadata(request.UserId)
	for _, dataType := range common.FIDataTypes {
		fiData, err := repo.GetUserFIDataForType(ctx, request.UserId, dataType)
		if err != nil {
			common.LogInfo(ctx, fmt.Sprintf("error-fetching-data-for-user-for-data-type: %s", dataType))
			continue
		}
		switch dataType {
		case "debit":
			response.FIData.Deposit = fiData
		case "recurring_deposit":
			response.FIData.Deposit = fiData
		case "term_deposit":
			response.FIData.TermDeposit = fiData
		case "credit_card":
			response.FIData.CreditCard = fiData
		case "cd":
			response.FIData.CD = fiData
		case "idr":
			response.FIData.IDR = fiData
		case "mutual_funds":
			response.FIData.MutualFunds = fiData
		case "bonds":
			response.FIData.Bonds = fiData
		case "debentures":
			response.FIData.Debentures = fiData
		case "equities":
			response.FIData.Equities = fiData
		case "ppf":
			response.FIData.PPF = fiData
		case "epf":
			response.FIData.EPF = fiData
		case "etf":
			response.FIData.ETF = fiData
		case "nps":
			response.FIData.NPS = fiData
		case "cp":
			response.FIData.CP = fiData
		case "govt_securities":
			response.FIData.GovtSecurities = fiData
		case "insurance_policies":
			response.FIData.InsurancePolicies = fiData
		case "ulip":
			response.FIData.ULIP = fiData
		case "reit":
			response.FIData.REIT = fiData
		case "invit":
			response.FIData.INVIT = fiData
		case "aif":
			response.FIData.AIF = fiData
		case "sip":
			response.FIData.SIP = fiData
		case "cis":
			response.FIData.CIS = fiData
		}
	}
	ctx.JSON(http.StatusOK, response)
}
