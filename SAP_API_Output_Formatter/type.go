package sap_api_output_formatter

type CompanyCodeReads struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Product       string `json:"Product"`
	APISchema     string `json:"api_schema"`
	CompanyCode   string `json:"company_code"`
	Deleted       string `json:"deleted"`
}

type CompanyCode struct {
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
}
