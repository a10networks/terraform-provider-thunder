package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateSmtp struct {
	SlbTemplateSMTPInstanceName SlbTemplateSMTPInstance `json:"smtp,omitempty"`
}

type SlbTemplateSMTPInstance struct {
	SlbTemplateSMTPInstanceClientDomainSwitchingSwitchingType []SlbTemplateSMTPInstanceClientDomainSwitching `json:"client-domain-switching,omitempty"`
	SlbTemplateSMTPInstanceClientStarttlsType                 string                                         `json:"client-starttls-type,omitempty"`
	SlbTemplateSMTPInstanceCommandDisableDisableType          []SlbTemplateSMTPInstanceCommandDisable        `json:"command-disable,omitempty"`
	SlbTemplateSMTPInstanceErrorCodeToClient                  int                                            `json:"error-code-to-client,omitempty"`
	SlbTemplateSMTPInstanceLFToCRLF                           int                                            `json:"LF-to-CRLF,omitempty"`
	SlbTemplateSMTPInstanceName                               string                                         `json:"name,omitempty"`
	SlbTemplateSMTPInstanceServerDomain                       string                                         `json:"server-domain,omitempty"`
	SlbTemplateSMTPInstanceServerStarttlsType                 string                                         `json:"server-starttls-type,omitempty"`
	SlbTemplateSMTPInstanceServiceReadyMsg                    string                                         `json:"service-ready-msg,omitempty"`
	SlbTemplateSMTPInstanceTemplateLogging                    SlbTemplateSMTPInstanceTemplate                `json:"template,omitempty"`
	SlbTemplateSMTPInstanceUUID                               string                                         `json:"uuid,omitempty"`
	SlbTemplateSMTPInstanceUserTag                            string                                         `json:"user-tag,omitempty"`
}

type SlbTemplateSMTPInstanceClientDomainSwitching struct {
	SlbTemplateSMTPInstanceClientDomainSwitchingMatchString   string `json:"match-string,omitempty"`
	SlbTemplateSMTPInstanceClientDomainSwitchingServiceGroup  string `json:"service-group,omitempty"`
	SlbTemplateSMTPInstanceClientDomainSwitchingSwitchingType string `json:"switching-type,omitempty"`
}

type SlbTemplateSMTPInstanceCommandDisable struct {
	SlbTemplateSMTPInstanceCommandDisableDisableType string `json:"disable-type,omitempty"`
}

type SlbTemplateSMTPInstanceTemplate struct {
	SlbTemplateSMTPInstanceTemplateLogging string `json:"logging,omitempty"`
}

func PostSlbTemplateSmtp(id string, inst SlbTemplateSmtp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateSmtp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/smtp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateSmtp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateSmtp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateSmtp(id string, name1 string, host string) (*SlbTemplateSmtp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateSmtp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/smtp/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateSmtp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateSmtp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateSmtp(id string, name1 string, inst SlbTemplateSmtp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateSmtp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/smtp/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateSmtp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateSmtp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateSmtp(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateSmtp")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/smtp/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateSmtp
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateSmtp", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
