package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-company-code-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetCompanyCode(companyCode, companyCodeName string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "CompanyCode":
			func() {
				c.CompanyCode(companyCode)
				wg.Done()
			}()
		case "CompanyCodeName":
			func() {
				c.CompanyCodeName(companyCodeName)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) CompanyCode(companyCode string) {
	companyCodeData, err := c.callCompanyCodeSrvAPIRequirementCompanyCode("A_CompanyCode", companyCode)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(companyCodeData)
}

func (c *SAPAPICaller) callCompanyCodeSrvAPIRequirementCompanyCode(api, companyCode string) ([]sap_api_output_formatter.CompanyCode, error) {
	url := strings.Join([]string{c.baseURL, "API_COMPANYCODE_SRV", api}, "/")
	param := c.getQueryWithCompanyCode(map[string]string{}, companyCode)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompanyCode(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToCompanyCode(url string) ([]sap_api_output_formatter.CompanyCode, error) {
	resp, err := c.requestClient.Request("GET", url, map[string]string{}, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompanyCode(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) CompanyCodeName(companyCodeName string) {
	data, err := c.callCompanyCodeSrvAPIRequirementCompanyCodeName("A_CompanyCode", companyCodeName)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callCompanyCodeSrvAPIRequirementCompanyCodeName(api, companyCodeName string) ([]sap_api_output_formatter.CompanyCode, error) {
	url := strings.Join([]string{c.baseURL, "API_COMPANYCODE_SRV", api}, "/")

	param := c.getQueryWithCompanyCodeName(map[string]string{}, companyCodeName)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToCompanyCode(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}


func (c *SAPAPICaller) getQueryWithCompanyCode(params map[string]string, companyCode string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("CompanyCode eq '%s'", companyCode)
	return params
}

func (c *SAPAPICaller) getQueryWithCompanyCodeName(params map[string]string, companyCodeName string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("substringof('%s', CompanyCodeName)", companyCodeName)
	return params
}
