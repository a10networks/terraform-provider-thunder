package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type ServerSSL struct {
	UUID ServerSSLInstance `json:"server-ssl,omitempty"`
}
type CrlCerts struct {
	Crl string `json:"crl,omitempty"`
}
type EcList struct {
	Ec string `json:"ec,omitempty"`
}
type CipherWithoutPrioList struct {
	CipherWoPrio string `json:"cipher-wo-prio,omitempty"`
}
type CaCerts struct {
	CaCert                string `json:"ca-cert,omitempty"`
	CaCertPartitionShared int    `json:"ca-cert-partition-shared,omitempty"`
	ServerOcspSg          string `json:"server-ocsp-sg,omitempty"`
	ServerOcspSrvr        string `json:"server-ocsp-srvr,omitempty"`
}
type ServerCertificateError struct {
	ErrorType string `json:"error-type,omitempty"`
}
type ServerSSLInstance struct {
	SessionCacheTimeout           int                      `json:"session-cache-timeout,omitempty"`
	CipherTemplate                string                   `json:"cipher-template,omitempty"`
	Sslilogging                   string                   `json:"sslilogging,omitempty"`
	UserTag                       string                   `json:"user-tag,omitempty"`
	Passphrase                    string                   `json:"passphrase,omitempty"`
	OcspStapling                  int                      `json:"ocsp-stapling,omitempty"`
	Crl                           []CrlCerts               `json:"crl-certs,omitempty"`
	UUID                          string                   `json:"uuid,omitempty"`
	KeySharedStr                  string                   `json:"key-shared-str,omitempty"`
	TemplateCipherShared          string                   `json:"template-cipher-shared,omitempty"`
	Dgversion                     int                      `json:"dgversion,omitempty"`
	CertSharedStr                 string                   `json:"cert-shared-str,omitempty"`
	Version                       int                      `json:"version,omitempty"`
	Ec                            []EcList                 `json:"ec-list,omitempty"`
	Encrypted                     string                   `json:"encrypted,omitempty"`
	SsliLogging                   int                      `json:"ssli-logging,omitempty"`
	SessionCacheSize              int                      `json:"session-cache-size,omitempty"`
	DhType                        string                   `json:"dh-type,omitempty"`
	UseClientSni                  int                      `json:"use-client-sni,omitempty"`
	ForwardProxyEnable            int                      `json:"forward-proxy-enable,omitempty"`
	Key                           string                   `json:"key,omitempty"`
	KeySharedEncrypted            string                   `json:"key-shared-encrypted,omitempty"`
	CipherWoPrio                  []CipherWithoutPrioList  `json:"cipher-without-prio-list,omitempty"`
	CaCert                        []CaCerts                `json:"ca-certs,omitempty"`
	Name                          string                   `json:"name,omitempty"`
	SharedPartitionCipherTemplate int                      `json:"shared-partition-cipher-template,omitempty"`
	EnableTLSAlertLogging         int                      `json:"enable-tls-alert-logging,omitempty"`
	SessionTicketEnable           int                      `json:"session-ticket-enable,omitempty"`
	AlertType                     string                   `json:"alert-type,omitempty"`
	Cert                          string                   `json:"cert,omitempty"`
	HandshakeLoggingEnable        int                      `json:"handshake-logging-enable,omitempty"`
	RenegotiationDisable          int                      `json:"renegotiation-disable,omitempty"`
	ErrorType                     []ServerCertificateError `json:"server-certificate-error,omitempty"`
	CloseNotify                   int                      `json:"close-notify,omitempty"`
	KeySharedPassphrase           string                   `json:"key-shared-passphrase,omitempty"`
}

func PostSlbTemplateServerSSL(id string, inst ServerSSL, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostSlbTemplateServerSSL")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/template/server-ssl", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServerSSL
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostSlbTemplateServerSSL REQ RES..........................", m)
			err := check_api_status("PostSlbTemplateServerSSL", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetSlbTemplateServerSSL(id string, name string, host string) (*ServerSSL, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetSlbTemplateServerSSL")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/template/server-ssl/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServerSSL
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetSlbTemplateServerSSL REQ RES..........................", m)
			err := check_api_status("GetSlbTemplateServerSSL", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}

func PutSlbTemplateServerSSL(id string, name string, inst ServerSSL, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PutSlbTemplateServerSSL")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("PUT", "https://"+host+"/axapi/v3/slb/template/server-ssl/"+name, bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServerSSL
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PutSlbTemplateServerSSL REQ RES..........................", m)
			err := check_api_status("PutSlbTemplateServerSSL", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func DeleteSlbTemplateServerSSL(id string, name string, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside DeleteSlbTemplateServerSSL")

	resp, err := DoHttp("DELETE", "https://"+host+"/axapi/v3/slb/template/server-ssl/"+name, nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err
		return err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m ServerSSL
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
