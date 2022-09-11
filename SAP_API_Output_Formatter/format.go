package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-company-code-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToCompanyCode(raw []byte, l *logger.Logger) ([]CompanyCode, error) {
	pm := &responses.CompanyCode{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to CompanyCode. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}
	companyCode := make([]CompanyCode, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		companyCode = append(companyCode, CompanyCode{
			CompanyCode:   					data.CompanyCode,
			CompanyCodeName:			    data.CompanyCodeName,
			CityName:						data.CityName,
			Country:						data.Country,
			Currency:						data.Currency,
			Language:						data.Language,
			ChartOfAccounts:				data.ChartOfAccounts,
			FiscalYearVariant:				data.FiscalYearVariant,
			Company:						data.Company,
			CreditControlArea:				data.CreditControlArea,
			CountryChartOfAccounts:			data.CountryChartOfAccounts,
			FinancialManagementArea:		data.FinancialManagementArea,
			AddressID:						data.AddressID,
			TaxableEntity:					data.TaxableEntity,
			VATRegistration:				data.VATRegistration,
			ExtendedWhldgTaxIsActive: 		data.ExtendedWhldgTaxIsActive,
			ControllingArea:				data.ControllingArea,
			FieldStatusVariant:				data.FieldStatusVariant,
			NonTaxableTransactionTaxCode:	data.NonTaxableTransactionTaxCode,
			DocDateIsUsedForTaxDetn:		data.DocDateIsUsedForTaxDetn,
			TaxRptgDateIsActive:			data.TaxRptgDateIsActive,
		})
	}

	return companyCode, nil
}
