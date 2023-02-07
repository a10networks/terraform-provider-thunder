package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type SlbTemplateServerSsl struct {
	SlbTemplateServerSslInstanceName SlbTemplateServerSslInstance `json:"server-ssl,omitempty"`
}

type SlbTemplateServerSslInstance struct {
	SlbTemplateServerSslInstanceAlertType                         string                                               `json:"alert-type,omitempty"`
	SlbTemplateServerSslInstanceCaCertsCaCert                     []SlbTemplateServerSslInstanceCaCerts                `json:"ca-certs,omitempty"`
	SlbTemplateServerSslInstanceCertificateCert                   SlbTemplateServerSslInstanceCertificate              `json:"certificate,omitempty"`
	SlbTemplateServerSslInstanceCipherTemplate                    string                                               `json:"cipher-template,omitempty"`
	SlbTemplateServerSslInstanceCipherWithoutPrioListCipherWoPrio []SlbTemplateServerSslInstanceCipherWithoutPrioList  `json:"cipher-without-prio-list,omitempty"`
	SlbTemplateServerSslInstanceCloseNotify                       int                                                  `json:"close-notify,omitempty"`
	SlbTemplateServerSslInstanceCrlCertsCrl                       []SlbTemplateServerSslInstanceCrlCerts               `json:"crl-certs,omitempty"`
	SlbTemplateServerSslInstanceDgversion                         int                                                  `json:"dgversion,omitempty"`
	SlbTemplateServerSslInstanceDhType                            string                                               `json:"dh-type,omitempty"`
	SlbTemplateServerSslInstanceEarlyData                         int                                                  `json:"early-data,omitempty"`
	SlbTemplateServerSslInstanceEcListEc                          []SlbTemplateServerSslInstanceEcList                 `json:"ec-list,omitempty"`
	SlbTemplateServerSslInstanceEnableSsliFtpAlg                  int                                                  `json:"enable-ssli-ftp-alg,omitempty"`
	SlbTemplateServerSslInstanceEnableTLSAlertLogging             int                                                  `json:"enable-tls-alert-logging,omitempty"`
	SlbTemplateServerSslInstanceForwardProxyEnable                int                                                  `json:"forward-proxy-enable,omitempty"`
	SlbTemplateServerSslInstanceHandshakeLoggingEnable            int                                                  `json:"handshake-logging-enable,omitempty"`
	SlbTemplateServerSslInstanceName                              string                                               `json:"name,omitempty"`
	SlbTemplateServerSslInstanceOcspStapling                      int                                                  `json:"ocsp-stapling,omitempty"`
	SlbTemplateServerSslInstanceRenegotiationDisable              int                                                  `json:"renegotiation-disable,omitempty"`
	SlbTemplateServerSslInstanceServerCertificateErrorErrorType   []SlbTemplateServerSslInstanceServerCertificateError `json:"server-certificate-error,omitempty"`
	SlbTemplateServerSslInstanceServerName                        string                                               `json:"server-name,omitempty"`
	SlbTemplateServerSslInstanceSessionCacheSize                  int                                                  `json:"session-cache-size,omitempty"`
	SlbTemplateServerSslInstanceSessionCacheTimeout               int                                                  `json:"session-cache-timeout,omitempty"`
	SlbTemplateServerSslInstanceSessionTicketEnable               int                                                  `json:"session-ticket-enable,omitempty"`
	SlbTemplateServerSslInstanceSharedPartitionCipherTemplate     int                                                  `json:"shared-partition-cipher-template,omitempty"`
	SlbTemplateServerSslInstanceSsliLogging                       int                                                  `json:"ssli-logging,omitempty"`
	SlbTemplateServerSslInstanceSslilogging                       string                                               `json:"sslilogging,omitempty"`
	SlbTemplateServerSslInstanceTemplateCipherShared              string                                               `json:"template-cipher-shared,omitempty"`
	SlbTemplateServerSslInstanceUUID                              string                                               `json:"uuid,omitempty"`
	SlbTemplateServerSslInstanceUseClientSni                      int                                                  `json:"use-client-sni,omitempty"`
	SlbTemplateServerSslInstanceUserTag                           string                                               `json:"user-tag,omitempty"`
	SlbTemplateServerSslInstanceVersion                           int                                                  `json:"version,omitempty"`
}

type SlbTemplateServerSslInstanceCaCerts struct {
	SlbTemplateServerSslInstanceCaCertsCaCert                string `json:"ca-cert,omitempty"`
	SlbTemplateServerSslInstanceCaCertsCaCertPartitionShared int    `json:"ca-cert-partition-shared,omitempty"`
	SlbTemplateServerSslInstanceCaCertsServerOcspSg          string `json:"server-ocsp-sg,omitempty"`
	SlbTemplateServerSslInstanceCaCertsServerOcspSrvr        string `json:"server-ocsp-srvr,omitempty"`
}

type SlbTemplateServerSslInstanceCertificate struct {
	SlbTemplateServerSslInstanceCertificateCert       string `json:"cert,omitempty"`
	SlbTemplateServerSslInstanceCertificateKey        string `json:"key,omitempty"`
	SlbTemplateServerSslInstanceCertificatePassphrase string `json:"passphrase,omitempty"`
	SlbTemplateServerSslInstanceCertificateShared     int    `json:"shared,omitempty"`
	SlbTemplateServerSslInstanceCertificateUUID       string `json:"uuid,omitempty"`
}

type SlbTemplateServerSslInstanceCipherWithoutPrioList struct {
	SlbTemplateServerSslInstanceCipherWithoutPrioListCipherWoPrio string `json:"cipher-wo-prio,omitempty"`
}

type SlbTemplateServerSslInstanceCrlCerts struct {
	SlbTemplateServerSslInstanceCrlCertsCrl                string `json:"crl,omitempty"`
	SlbTemplateServerSslInstanceCrlCertsCrlPartitionShared int    `json:"crl-partition-shared,omitempty"`
}

type SlbTemplateServerSslInstanceEcList struct {
	SlbTemplateServerSslInstanceEcListEc string `json:"ec,omitempty"`
}

type SlbTemplateServerSslInstanceServerCertificateError struct {
	SlbTemplateServerSslInstanceServerCertificateErrorErrorType string `json:"error-type,omitempty"`
}

func PostSlbTemplateServerSsl(id string, inst SlbTemplateServerSsl, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateServerSsl")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/server-ssl", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateServerSsl
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Post REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateServerSsl", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateServerSsl(id string, name1 string, host string) (*SlbTemplateServerSsl, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateServerSsl")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/server-ssl/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateServerSsl
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] Get REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateServerSsl", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateServerSsl(id string, name1 string, inst SlbTemplateServerSsl, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateServerSsl")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
		return err
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/server-ssl/"+name1, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateServerSsl
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Put REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateServerSsl", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateServerSsl(id string, name1 string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateServerSsl")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/server-ssl/"+name1, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m SlbTemplateServerSsl
		err := json.Unmarshal(data, &m)
		if err != nil {
			logger.Println("Unmarshal error ", err)
			return err
		} else {
			logger.Println("[INFO] Delete REQ RES..........................", m)
			err := check_api_status("DeleteSlbTemplateServerSsl", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}
