package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type CSV struct {
	UUID CSVInstance `json:"csv,omitempty"`
}

type MultipleFields struct {
	Field   int    `json:"field,omitempty"`
	CsvType string `json:"csv-type,omitempty"`
}
type CSVInstance struct {
	UUID       string           `json:"uuid,omitempty"`
	CsvName    string           `json:"csv-name,omitempty"`
	UserTag    string           `json:"user-tag,omitempty"`
	Ipv6Enable int              `json:"ipv6-enable,omitempty"`
	DelimNum   int              `json:"delim-num,omitempty"`
	DelimChar  string           `json:"delim-char,omitempty"`
	Field      []MultipleFields `json:"multiple-fields,omitempty"`
}

func PostSlbTemplateCSV(id string, inst CSV, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateCSV")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/csv", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m CSV
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateCSV REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateCSV", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateCSV(id string, name string, host string) (*CSV, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateCSV")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/csv/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m CSV
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateCSV REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateCSV", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateCSV(id string, name string, inst CSV, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateCSV")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/csv/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m CSV
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateCSV REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateCSV", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateCSV(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateCSV")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/csv/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m CSV
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] GET REQ RES..........................", m)

		}
	}
	return nil
}
