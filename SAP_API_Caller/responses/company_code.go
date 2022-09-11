package responses

type CompanyCode struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
				Etag string `json:"etag"`
			} `json:"__metadata"`
			CompanyCode   					string `json:"CompanyCode"`
			CompanyCodeName					string `json:"CompanyCodeName"`
			CityName						string `json:"CityName"`
			Country							string `json:"Country"`
			Currency						string `json:"Currency"`
			Language						string `json:"Language"`
			ChartOfAccounts					string `json:"ChartOfAccounts"`
			FiscalYearVariant				string `json:"FiscalYearVariant"`
			Company							string `json:"Company"`
			CreditControlArea				string `json:"CreditControlArea"`
			CountryChartOfAccounts 			string `json:"CountryChartOfAccounts"`
			FinancialManagementArea 		string `json:"FinancialManagementArea"`
			AddressID						string `json:"AddressID"`
			TaxableEntity					string `json:"TaxableEntity"`
			VATRegistration					string `json:"VATRegistration"`
			ExtendedWhldgTaxIsActive 		bool   `json:"ExtendedWhldgTaxIsActive"`
			ControllingArea					string `json:"ControllingArea"`
			FieldStatusVariant				string `json:"FieldStatusVariant"`
			NonTaxableTransactionTaxCode 	string `json:"NonTaxableTransactionTaxCode"`
			DocDateIsUsedForTaxDetn			bool   `json:"DocDateIsUsedForTaxDetn"`
			TaxRptgDateIsActive				bool   `json:"TaxRptgDateIsActive"`
		} `json:"results"`
	} `json:"d"`
}
